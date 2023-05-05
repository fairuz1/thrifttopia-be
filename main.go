package main

import (
	"log"
	"net/http"
	"thriftopia/connection"
	"thriftopia/controllers/auth_controller"
	"thriftopia/controllers/history_controller"
	"thriftopia/controllers/log_activity_controller"
	"thriftopia/controllers/pricing_controller"
	"thriftopia/controllers/product_controller"
	"thriftopia/controllers/user_controller"
	"thriftopia/controllers/user_role_controller"

	"github.com/gorilla/mux"
)

func main() {
	connection.ConnectDatabase()
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/v1/register", user_controller.Register).
		Methods("POST")
	r.HandleFunc("/v1/login", func(w http.ResponseWriter, r *http.Request) {
		auth_controller.Login(w, r, connection.DB)
	}).
		Methods("POST")
	r.HandleFunc("/v1/logout", func(w http.ResponseWriter, r *http.Request) {
		auth_controller.Logout(w, r)
	}).
		Methods("GET")
	r.HandleFunc("/v1/users", user_controller.GetList).
		Methods("GET")
	r.HandleFunc("/v1/user/{id}", user_controller.GetDetail).
		Methods("GET")
	r.HandleFunc("/v1/user/{id}", user_controller.Update).
		Methods("PUT")
	r.HandleFunc("/v1/userroles", user_role_controller.GetList).
		Methods("GET")

	// Product routes
	r.HandleFunc("/v1/product", product_controller.Create).
		Methods("POST")
	r.HandleFunc("/v1/products", product_controller.GetList).
		Methods("GET")
	r.HandleFunc("/v1/product/{id}", product_controller.GetDetail).
		Methods("GET")
	r.HandleFunc("/v1/product/{id}", product_controller.Update).
		Methods("PUT")
	r.HandleFunc("/v1/product/publish/{id}", product_controller.Publish).
		Methods("PUT")
	r.HandleFunc("/v1/product/reject/{id}", product_controller.Reject).
		Methods("PUT")
	r.HandleFunc("/v1/product/sold/{id}", product_controller.ChangeToSold).
		Methods("PUT")
	r.HandleFunc("/v1/transaction/history", history_controller.GetTransactions).
		Methods("GET")

	// Pricing Plans routes
	r.HandleFunc("/v1/pricing_plan", pricing_controller.Create).
		Methods("POST")
	r.HandleFunc("/v1/pricing_plans", pricing_controller.GetList).
		Methods("GET")
	r.HandleFunc("/v1/pricing_plan/{id}", pricing_controller.Update).
		Methods("PUT")
	r.HandleFunc("/v1/pricing_plan/{id}", pricing_controller.Delete).
		Methods("DELETE")

	// Log Acitivity routes
	r.HandleFunc("/v1/log_activity", log_activity_controller.Create).
		Methods("POST")
	r.HandleFunc("/v1/log_activities", log_activity_controller.GetList).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
