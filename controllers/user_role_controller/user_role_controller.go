package user_role_controller

import (
	"net/http"
	"thriftopia/connection"
	"thriftopia/helper"
	"thriftopia/models"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func GetList(w http.ResponseWriter, r *http.Request) {
	var roles []models.Role

	if err := connection.DB.Find(&roles).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["data"] = roles
	responseData["message"] = "Success Get All Roles"

	ResponseJson(w, http.StatusOK, responseData)
}
