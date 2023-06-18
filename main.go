package main

import (
	"fmt"
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
	"thriftopia/controllers/whatsapp_api"
	"thriftopia/middleware"

	"github.com/gorilla/mux"
)

func main() {
	connection.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Thriftopia API!")
	})

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

	r.Handle("/v1/users", middleware.AdminAuthenticator(http.HandlerFunc(user_controller.GetList))).
		Methods("GET")

	r.Handle("/v1/user/{id}", middleware.Authenticator(http.HandlerFunc(user_controller.GetDetail))).
		Methods("GET")
	r.Handle("/v1/user/{id}", middleware.Authenticator(http.HandlerFunc(user_controller.Update))).
		Methods("PUT")
	r.Handle("/v1/userroles", middleware.Authenticator(http.HandlerFunc(user_role_controller.GetList))).
		Methods("GET")

	// Product routes
	r.Handle("/v1/product", middleware.Authenticator(http.HandlerFunc(product_controller.Create))).
		Methods("POST")
	r.HandleFunc("/v1/products", product_controller.GetList).
		Methods("GET")
	r.HandleFunc("/v1/product/{id}", product_controller.GetDetail).
		Methods("GET")
	r.Handle("/v1/product/{id}", middleware.AdminAuthenticator(http.HandlerFunc(product_controller.Update))).
		Methods("PUT")
	r.Handle("/v1/product/publish/{id}", middleware.AdminAuthenticator(http.HandlerFunc(product_controller.Publish))).
		Methods("PUT")
	r.Handle("/v1/product/reject/{id}", middleware.AdminAuthenticator(http.HandlerFunc(product_controller.Reject))).
		Methods("PUT")
	r.Handle("/v1/product/sold/{id}", middleware.Authenticator(http.HandlerFunc(product_controller.ChangeToSold))).
		Methods("PUT")
	r.HandleFunc("/v1/transaction/history", history_controller.GetTransactions).
		Methods("GET")

	// Pricing Plans routes
	r.Handle("/v1/pricing_plan", middleware.AdminAuthenticator(http.HandlerFunc(pricing_controller.Create))).
		Methods("POST")
	r.HandleFunc("/v1/pricing_plans", pricing_controller.GetList).
		Methods("GET")
	r.Handle("/v1/pricing_plan/{id}", middleware.AdminAuthenticator(http.HandlerFunc(pricing_controller.Update))).
		Methods("PUT")
	r.Handle("/v1/pricing_plan/{id}", middleware.AdminAuthenticator(http.HandlerFunc(pricing_controller.Delete))).
		Methods("DELETE")

	// Log Acitivity routes
	r.HandleFunc("/v1/log_activity", log_activity_controller.Create).
		Methods("POST")
	r.Handle("/v1/log_activities", middleware.AdminAuthenticator(http.HandlerFunc(log_activity_controller.GetList))).
		Methods("GET")

	r.HandleFunc("/validate/{phone_number}", whatsapp_api.ValidateNumber).Methods("GET")

	log.Fatal(http.ListenAndServe(":9990", r))
}
