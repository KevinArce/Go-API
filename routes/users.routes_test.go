package routes

import (
	"encoding/json"     // We need the json package for encoding and decoding the data
	"fmt"               // We need the fmt package for printing the data
	"net/http"          // We need the net/http package for the handler function
	"net/http/httptest" // We need the net/http/httptest package for testing the handler
	"strings"           // We need the strings package for the NewReader function
	"testing"           // We need the testing package for testing our handler

	"github.com/KevinArce/Go-Api/models" // We need the models package for the User struct
	"github.com/gorilla/mux"             // We need the gorilla/mux package for the router
)

func TestCreateUserHandler(t *testing.T) { // This is the function that will test our route
	router := mux.NewRouter()                                             // We create a new router
	router.HandleFunc("/api/v1/users", CreateUserHandler).Methods("POST") // We add our route to the router we created above
	//Print in the console the router
	fmt.Println(router)

	// We create a new request to our route with the method POST and the body we want to send to our route
	req, err := http.NewRequest("POST", "/api/v1/users", strings.NewReader(`{"username": "Test5", "firstname": "Kevinsky", "lastname": "Arcelyk", "email": "email@test.com"}`))

	//Print in the console the request with the header
	fmt.Println("Paso header", req)

	//Print in the console the error
	fmt.Println("Muestra error", err)

	// We check if there is an error
	if err != nil {
		//Print the value of the error
		fmt.Println(err)
		t.Fatal(err) // If there is an error we stop the test
	}

	rr := httptest.NewRecorder() // We create a new recorder
	router.ServeHTTP(rr, req)    // We send the request to our route

	if status := rr.Code; status != http.StatusOK { // We check if the status code is the one we expect (200)
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK) // If the status code is not the one we expect we stop the test and print the error
	}

	var user models.User                   // We create a new user
	json.NewDecoder(rr.Body).Decode(&user) // We decode the body of the response and save it in the user variable we created above

	if user.Username != "Test5" { // We check if the username of the user is the one we expect (Test)
		t.Errorf("handler returned unexpected body: got %v want %v", user.Username, "Test5") // If the username is not the one we expect we stop the test and print the error
	}

	if user.FirstName != "Kevinsky" { // We check if the first name of the user is the one we expect (Luisillo)
		t.Errorf("handler returned unexpected body: got %v want %v", user.FirstName, "Kevinsky") // If the first name is not the one we expect we stop the test and print the error
	}

	if user.LastName != "Arcelyk" { // We check if the last name of the user is the one we expect (Arcelyk)
		t.Errorf("handler returned unexpected body: got %v want %v", user.LastName, "Arcelyk") // If the last name is not the one we expect we stop the test and print the error
	}

	if user.Email != "email@test.com" { // We check if the email of the user is the one we expect (
		t.Errorf("handler returned unexpected body: got %v want %v", user.Email, "email@test.com") // If the email is not the one we expect we stop the test and print the error
	}
}

// func TestGetUsersHandler(t *testing.T) {
// 	req := httptest.NewRequest("GET", "/api/v1/users", nil)
// 	rr := httptest.NewRecorder()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/api/v1/users", GetUsersHandler).Methods("GET")
// 	router.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 	}

// 	var users []models.User
// 	json.NewDecoder(rr.Body).Decode(&users)

// 	if len(users) != 2 {
// 		t.Errorf("handler returned unexpected body: got %v want %v", len(users), 2)
// 	}
// }
