package models

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ResponseSuccessLogin struct {
	Email           string    `json:"email"`
	SessionToken string `json:"session_token"`
	Message      string `json:"message"`
}
