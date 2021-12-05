package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	Movies = []Movie{
		{
			ID:    1,
			Title: "First Movie",
			Desc:  "This is the first movie",
		},
		{
			ID:    2,
			Title: "Second Movie",
			Desc:  "This is the second movie",
		},
	}

	handleRequests()

}

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var Movies []Movie

func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World!\nThis is the HomePage")
}

func getAllMovies(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(Movies)
}

func getMovieById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, movie := range Movies {
		if id == movie.ID {
			json.NewEncoder(rw).Encode(movie)
		}

	}
}

func addMovie(rw http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("there was an error", err)
	}

	var movie Movie

	json.Unmarshal(reqBody, &movie)

	Movies = append(Movies, movie)
	json.NewEncoder(rw).Encode(movie)

}

func updateMovie(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var updateMovie Movie

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("there was an error", err)
	}

	json.Unmarshal(reqBody, &updateMovie)


	for index, movie := range Movies {
		if movie.ID == id {
			movie.Title = updateMovie.Title
			movie.Desc = updateMovie.Desc
			Movies = append(Movies[:index], movie)
			json.NewEncoder(rw).Encode(movie)


		}	

	}



}

func deleteMovie(rw http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, movie := range Movies {
		if movie.ID == id {
			Movies = append(Movies[:index], Movies[index+1:]...)
		}
	}
}

func handleRequests() {

	router := mux.NewRouter()

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/movies", getAllMovies).Methods("GET")
	router.HandleFunc("/movies", addMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	

	http.ListenAndServe(":8000", router)
}

/*

{
    "id": 5,
    "title": "fifth movie"
    "desc": "this is the fifth movie"
}

*/