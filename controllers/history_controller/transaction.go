package history_controller

import (
	"net/http"
	"thriftopia/connection"
	"thriftopia/helper"
	"thriftopia/models"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.TransactionHistories

	if err := connection.DB.Find(&transactions).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	responseData["data"] = transactions
	responseData["message"] = "Success Get All Transaction Histories"

	ResponseJson(w, http.StatusOK, responseData)

}
