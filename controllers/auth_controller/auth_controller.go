package auth_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thriftopia/helper"
	"thriftopia/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GenerateToken(userId int, name string, role string, email string) (string, error) {
	// Buat token baru dengan claims yang sesuai
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"name":    name,
		"role":    role,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 6).Unix(),
	})

	// Tandatangani token menggunakan secret key yang rahasia
	tokenString, err := token.SignedString([]byte("secret-key-thriftopia"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		helper.ResponseJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// Query the database to retrieve the user's credentials
	var user models.User
	err = db.Preload("Role").Where(`email = $1`, creds.Email).First(&user).Error
	if err != nil {
		helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": "Wrong password"})
		return
	}

	token, err := GenerateToken(user.Id, user.Name, user.Role.Name, user.Email)
	if err != nil {
		helper.ResponseJson(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	userData := models.LoginData{
		Email:    user.Email,
		Username: user.Name,
		Role:     user.Role.Name,
		Token:    token,
	}

	data := models.ResponseSuccessLogin{
		Data:    userData,
		Message: "Login success",
	}

	jsonRes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var res interface{}
	err = json.Unmarshal(jsonRes, &res)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	helper.ResponseJson(w, http.StatusOK, res)

}

func Logout(w http.ResponseWriter, r *http.Request) {
		
}
