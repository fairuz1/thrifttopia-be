package pricing_controller

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

	var pricing models.Pricings
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pricing); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()


	if err := connection.DB.Create(&pricing).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["message"] = "Success Create Pricing Plan"

	ResponseJson(w, http.StatusCreated, responseData)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var pricings []models.Pricings

	if err := connection.DB.Find(&pricings).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["data"] = pricings
	responseData["message"] = "Success Get All Pricing Plans"

	ResponseJson(w, http.StatusOK, responseData)

}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var pricing models.Pricings
	if err := connection.DB.First(&pricing, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Pricing plan not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pricing); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := connection.DB.Save(&pricing).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["message"] = "Success Update Pricing Plan with ID " + vars["id"]

	helper.ResponseJson(w, http.StatusOK, responseData)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var pricing models.Pricings
	if err := connection.DB.First(&pricing, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Pricing plan not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if err := connection.DB.Delete(&pricing).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["message"] = "Success Delete Pricing Plan with ID " + vars["id"]

	helper.ResponseJson(w, http.StatusOK, responseData)
}
