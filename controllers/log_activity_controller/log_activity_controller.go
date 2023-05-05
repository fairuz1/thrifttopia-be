package log_activity_controller

import (
	"encoding/json"
	"net/http"
	"thriftopia/connection"
	"thriftopia/helper"
	"thriftopia/models"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Create(w http.ResponseWriter, r *http.Request) {

	var log_activity models.LogActivity
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&log_activity); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := connection.DB.Create(&log_activity).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["message"] = "Success Create Log Activity"

	ResponseJson(w, http.StatusCreated, responseData)

}

func GetList(w http.ResponseWriter, r *http.Request) {
	var log_activity []models.LogActivity

	if err := connection.DB.Find(&log_activity).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["data"] = log_activity
	responseData["message"] = "Success Get All Log Activity"

	ResponseJson(w, http.StatusOK, responseData)

}