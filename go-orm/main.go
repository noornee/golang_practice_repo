package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noornee/golang_practice_repo/go-orm/handlers"
)

const port string = ":8000"

func main() {
	fmt.Printf("Server running on port %v\n", port)

	handlers.InitialMigration()

	handleRequests()

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user", handlers.CreateUser).Methods("POST")

	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(port, router))

}
