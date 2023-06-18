package whatsapp_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"thriftopia/helper"

	"github.com/gorilla/mux"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func ValidateNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phone_number := vars["phone_number"]

	payload := []byte(``)

	req, err := http.NewRequest("GET", "https://api.p.2chat.io/open/whatsapp/check-number/+6281232032275/"+phone_number, bytes.NewBuffer(payload))
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Failed to create HTTP request")
		return
	}

	req.Header.Set("X-User-API-Key", "UAK165bcc24-434d-4fd5-9795-2c920703ebf4")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Failed to create HTTP request")
		return
	}
	defer resp.Body.Close()

	responseFromWAAPI := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&responseFromWAAPI)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, "Failed to decode JSON response")
		fmt.Println("Failed to decode JSON response:", err)
		return
	}

	fmt.Println("responseData from WhatsApp API => ", responseFromWAAPI)

	

	// ResponseJson(w, http.StatusCreated, responseData)

	isValid, ok := responseFromWAAPI["is_valid"].(bool)
	if !ok {
		ResponseError(w, http.StatusOK, "Invalid request")
		return
	}

	if isValid {
		onWhatsApp, ok := responseFromWAAPI["on_whatsapp"].(bool)
		if !ok {
			responseData := make(map[string]interface{})
			responseData["is_valid_number"] = isValid

			ResponseJson(w, http.StatusOK, responseData)
			return
		}
		responseData := make(map[string]interface{})
		responseData["is_valid_number"] = isValid
		responseData["on_whatsapp"] = onWhatsApp

		ResponseJson(w, http.StatusOK, responseData)
		fmt.Println("okeeeeeeiiiiiiiiii")
		return
	} else {
		responseData := make(map[string]interface{})
		responseData["is_valid_number"] = isValid

		ResponseJson(w, http.StatusOK, responseData)
		return
	}

	// ResponseJson(w, resp.StatusCode, responseData)
	// fmt.Println("WA VALID. THEN =>", responseFromWAAPI)
}