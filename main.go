package main

import (
	"log"
	"net/http"
	"thriftopia/connection"
	"thriftopia/controllers/auth_controller"
	"thriftopia/controllers/user_controller"

	"github.com/gorilla/mux"
)

func main() {
	connection.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/register", user_controller.Register).Methods("POST")
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth_controller.Login(w, r, connection.DB)
	}).Methods("POST")
	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		auth_controller.Logout(w, r)
	}).Methods("GET")
	r.HandleFunc("/users", user_controller.GetList).Methods("GET")
	r.HandleFunc("/user/{id}", user_controller.GetDetail).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
