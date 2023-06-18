package log_activity_controller

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
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

	query := r.URL.Query()
	userID := query.Get("user_id")
	activityIDStr := query.Get("activity_id")
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

	offset := (page - 1) * pageSize

	db := connection.DB

	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if activityIDStr != "" {
		activityID, err := strconv.Atoi(activityIDStr)
		if err != nil {
			ResponseError(w, http.StatusBadRequest, "Invalid value for 'userID' parameter")
			return
		}
		db = db.Where("activity_id = ?", activityID)
	}

	var totalCount int64
	if err := db.Model(&models.LogActivity{}).Count(&totalCount).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var totalPages int64
	if totalCount > 0 {
		totalPages = int64(math.Ceil(float64(totalCount) / float64(pageSize)))
	} else {
		totalPages = 0
	}

	if err := db.Offset(offset).Limit(pageSize).Find(&log_activity).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := make(map[string]interface{})
	meta["page"] = page
	meta["page_size"] = pageSize
	meta["total"] = totalCount
	meta["total_pages"] = totalPages

	responseData := make(map[string]interface{})
	responseData["data"] = log_activity
	responseData["meta"] = meta
	responseData["message"] = "Success Get All Log Activity"

	ResponseJson(w, http.StatusOK, responseData)

}
