package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/KevinArce/Go-Api/db"
	"github.com/KevinArce/Go-Api/routes"
)

func main() {

	db.Connect()

	r := mux.NewRouter()
	
	r.HandleFunc("/api/v1/home", routes.HomeHandler)
	
	http.ListenAndServe(":3000", r)
}