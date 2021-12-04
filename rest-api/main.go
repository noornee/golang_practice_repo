package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	Movies = []*Movie{
		{
			ID:    "1",
			Title: "First Movie",
			Desc:  "This is the first movie",
		},
		{
			ID:    "2",
			Title: "Second Movie",
			Desc:  "This is the second movie",
		},
	}

	handleRequests()

}

type Movie struct {
	ID    string    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var Movies []*Movie

func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World!\nThis is the HomePage")
}

func getAllMovies(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(Movies)
}

func getMovieById(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	for _, movie := range Movies {
		if vars["id"] == movie.ID {
			json.NewEncoder(rw).Encode(movie)
		}

	}
}

func handleRequests() {

	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/movies", getAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")


	http.ListenAndServe(":8000", router)
}
