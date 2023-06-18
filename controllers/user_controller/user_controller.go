package user_controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"thriftopia/connection"
	"thriftopia/helper"
	"thriftopia/models"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = string(hashedPassword)

	query := `INSERT INTO users (role_id, name, email, password, wa_number, created_at, updated_at)
			  VALUES (?, ?, ?, ?, ?, ?, ?)`

	if err := connection.DB.Exec(query, user.RoleId, user.Name, user.Email, user.Password, user.WaNumber, user.CreatedAt, user.UpdatedAt).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	meta := make(map[string]interface{})
	meta["created_at"] = user.CreatedAt
	meta["updated_at"] = user.UpdatedAt

	responseData := make(map[string]interface{})
	responseData["meta"] = meta
	responseData["message"] = "Success Create User"

	helper.ResponseJson(w, http.StatusCreated, responseData)

}

func GetList(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := connection.DB.Preload("Role").Find(&users).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make(map[string]interface{})
	var data []interface{}
	for _, user := range users {
		item := struct {
			ID        int       `json:"id"`
			Role      string    `json:"role"`
			Name      string    `json:"name"`
			Email     string    `json:"email"`
			WaNumber  string    `json:"wa_number"`
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
		}{
			ID:        user.Id,
			Role:      user.Role.Name,
			Name:      user.Name,
			Email:     user.Email,
			WaNumber:  user.WaNumber,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		data = append(data, item)
	}

	responseData["data"] = data
	responseData["message"] = "Success Get All Users"

	helper.ResponseJson(w, http.StatusOK, responseData)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := connection.DB.Preload("Role").First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "User not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	responseData := make(map[string]interface{})
	data := make(map[string]interface{})
	data["id"] = user.Id
	data["role"] = user.Role.Name
	data["name"] = user.Name
	data["email"] = user.Email
	data["wa_number"] = user.WaNumber
	data["created_at"] = user.CreatedAt
	data["updated_at"] = user.UpdatedAt

	responseData["data"] = data
	responseData["message"] = "Success Get Detail Users"

	helper.ResponseJson(w, http.StatusOK, responseData)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := connection.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "User not found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := connection.DB.Save(&user).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.UpdatedAt = time.Now()

	meta := make(map[string]interface{})
	meta["created_at"] = user.CreatedAt
	meta["updated_at"] = user.UpdatedAt

	responseData := make(map[string]interface{})
	responseData["meta"] = meta
	responseData["message"] = "Success Update User with ID " + vars["id"]

	helper.ResponseJson(w, http.StatusCreated, responseData)
}
