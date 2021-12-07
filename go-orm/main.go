package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noornee/golang_practice_repo/go-orm/handlers"
	"github.com/noornee/golang_practice_repo/go-orm/models"
)

const port string = ":8000"

func main() {
	fmt.Printf("Server running on port %v\n", port)

	models.InitialMigration()

	handleRequests()
	
	
}



func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	
	router.HandleFunc("/users", handlers.AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{name}", handlers.UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(port, router))


}