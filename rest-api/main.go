package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	Movies = []*Movie{
		{
			ID: 1,
			Title: "First Movie",
			Desc: "This is the first movie",
		},
		{
			ID: 2,
			Title: "Second Movie",
			Desc: "This is the second movie",
		},
	}

	handleRequests()

}

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var Movies []*Movie

func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World!\nThis is the HomePage")
}

func getAllMovies(rw http.ResponseWriter, r *http.Request){
	json.NewEncoder(rw).Encode(Movies)
}



func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/movies", getAllMovies)
	http.ListenAndServe(":8000", nil)
}
