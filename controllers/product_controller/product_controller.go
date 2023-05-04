package product_controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"thriftopia/connection"
	"thriftopia/helper"
	"thriftopia/models"

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

	if err := connection.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["data"] = products
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
