package auth_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"thriftopia/helper"
	"thriftopia/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// user's sessions
var sessions = map[string]session{}

type session struct {
	email string
	expiry   time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
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
	err = db.Where(`email = $1`, creds.Email).First(&user).Error
	if err != nil {
		helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
		return
	}

	if user.Password != creds.Password {
		helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": "Wrong password"})
		return
	}

	// session token
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(24 * time.Hour)

	sessions[sessionToken] = session{
		email: creds.Email,
		expiry:   expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})

	userData := models.LoginData{
		Email: user.Email,
		Username: user.Name,
		SessionToken: sessionToken,
	}

	data := models.ResponseSuccessLogin{
		Data: userData,
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

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mendapat session token dari requests cookies
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// cookie tidak di set
				helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": "Maaf, anda harus login!"})
				return
			}

			helper.ResponseJson(w, http.StatusBadRequest, map[string]string{"message": "Maaf, anda harus login!"})
			return
		}
		sessionToken := c.Value

		// mendapat userSession dari map session
		userSession, exists := sessions[sessionToken]

		if !exists {
			// Session token tidak ada pada map session
			helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": "Maaf, anda harus login!"})
			return
		}

		if userSession.isExpired() {
			delete(sessions, sessionToken)
			helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": "Maaf, anda harus login!"})
			return
		}

		ctx := context.WithValue(r.Context(), "email", userSession.email)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// mendapat session token dari requests cookies
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			helper.ResponseJson(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
			return
		}
		helper.ResponseJson(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	sessionToken := c.Value

	// hapus user's session dari map session
	delete(sessions, sessionToken)

	// hapus cookie dari user & set waktu expired  = sekarang
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})

	helper.ResponseJson(w, http.StatusOK, map[string]string{"message": "Logout success"})
}
