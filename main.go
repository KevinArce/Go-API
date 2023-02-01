package main

import (
	"net/http"

	"github.com/KevinArce/Go-Api/db"
	"github.com/KevinArce/Go-Api/models"
	"github.com/KevinArce/Go-Api/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.Connect()
	db.DB.AutoMigrate(&models.User{}, &models.Task{})

	r := mux.NewRouter()
	
	r.HandleFunc("/api/v1/", routes.HomeHandler)
	r.HandleFunc("/api/v1/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/api/v1/users", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", routes.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	
	http.ListenAndServe(":3000", r)
}