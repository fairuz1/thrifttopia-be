package product_controller

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"thriftopia/connection"
	"thriftopia/helper"
	"thriftopia/models"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Create(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := connection.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := make(map[string]interface{})
	meta["created_at"] = product.CreatedAt
	meta["updated_at"] = product.UpdatedAt

	responseData := make(map[string]interface{})
	responseData["meta"] = meta
	responseData["message"] = "Success Create Product"

	ResponseJson(w, http.StatusCreated, responseData)

}

func GetList(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	query := r.URL.Query()
	isSoldStr := query.Get("is_sold")

	// Pagination parameters
	pageStr := query.Get("page")
	pageSizeStr := query.Get("page_size")

	// Default pagination values
	page := 1
	pageSize := 10

	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			ResponseError(w, http.StatusBadRequest, "Invalid value for 'page' parameter")
			return
		}
	}

	if pageSizeStr != "" {
		var err error
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil || pageSize <= 0 {
			ResponseError(w, http.StatusBadRequest, "Invalid value for 'page_size' parameter")
			return
		}
	}

	// Calculate offset based on page number and page size
	offset := (page - 1) * pageSize

	db := connection.DB
	var isSold bool
	if isSoldStr != "" {
		var err error
		isSold, err = strconv.ParseBool(isSoldStr)
		if err != nil {
			ResponseError(w, http.StatusBadRequest, "Invalid value for 'is_sold' parameter")
			return
		}
		db = db.Where("is_sold = ?", isSold)
	}

	var totalCount int64
	if err := db.Model(&models.Product{}).Count(&totalCount).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var totalPages int64
	if totalCount > 0 {
		totalPages = int64(math.Ceil(float64(totalCount) / float64(pageSize)))
	} else {
		totalPages = 0
	}

	if err := db.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := make(map[string]interface{})
	meta["page"] = page
	meta["page_size"] = pageSize
	meta["total"] = totalCount
	meta["total_pages"] = totalPages

	responseData := make(map[string]interface{})
	responseData["data"] = products
	responseData["meta"] = meta
	responseData["message"] = "Success Get All Products"

	ResponseJson(w, http.StatusOK, responseData)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := connection.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	responseData := make(map[string]interface{})
	responseData["data"] = product
	responseData["message"] = "Success Get Detail Product"

	ResponseJson(w, http.StatusOK, responseData)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := connection.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := connection.DB.Save(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	product.UpdatedAt = time.Now()

	meta := make(map[string]interface{})
	meta["created_at"] = product.CreatedAt
	meta["updated_at"] = product.UpdatedAt

	responseData := make(map[string]interface{})
	responseData["meta"] = meta
	responseData["message"] = "Success Update Product with ID " + vars["id"]

	helper.ResponseJson(w, http.StatusOK, responseData)
}

func ChangeToSold(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := connection.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if product.IsSold {
		ResponseError(w, http.StatusBadRequest, "Product already sold")
		return
	}

	query := r.URL.Query()
	buyerID, err := strconv.Atoi(query.Get("buyer_id"))
	if err != nil {
		if query.Get("buyer_id") == "" {
			ResponseError(w, http.StatusBadRequest, "buyer_id is required in query params")
			return
		}
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	transactionHistory := models.TransactionHistories{
		ProductID: product.Id,
		BuyerID:   buyerID,
		CreatedAt: time.Now(),
	}

	if err := connection.DB.Create(&transactionHistory).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	product.IsSold = true

	if err := connection.DB.Save(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := make(map[string]interface{})
	meta["created_at"] = product.CreatedAt
	meta["updated_at"] = product.UpdatedAt

	responseData := make(map[string]interface{})
	responseData["meta"] = meta
	responseData["message"] = "Success Change Product with ID " + vars["id"] + " to Sold"

	ResponseJson(w, http.StatusOK, responseData)
}

func Publish(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := connection.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	product.Status = "published"

	if err := connection.DB.Save(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := make(map[string]interface{})
	meta["created_at"] = product.CreatedAt
	meta["updated_at"] = product.UpdatedAt

	responseData := make(map[string]interface{})
	responseData["meta"] = meta
	responseData["message"] = "Success Publish Product with ID " + vars["id"]

	ResponseJson(w, http.StatusOK, responseData)
}

func Reject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := connection.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	product.Status = "rejected"

	if err := connection.DB.Save(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := make(map[string]interface{})
	meta["created_at"] = product.CreatedAt
	meta["updated_at"] = product.UpdatedAt

	responseData := make(map[string]interface{})
	responseData["meta"] = meta
	responseData["message"] = "Success Reject Product with ID " + vars["id"]

	ResponseJson(w, http.StatusOK, responseData)
}
