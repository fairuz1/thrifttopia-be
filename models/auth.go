package models

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginData struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	SessionToken string `json:"token"`
}
type ResponseSuccessLogin struct {
	Data    LoginData `json:"data"`
	Message string    `json:"message"`
}
