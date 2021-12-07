package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"
	"github.com/noornee/golang_practice_repo/go-orm/models"
)

func AllUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Get Users endpoint")
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var users []models.User

	db.Find(&users)
	fmt.Println(users)
	json.NewEncoder(rw).Encode(users)

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Create User endpoint")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	json.Unmarshal(reqBody, &user)

	defer db.Close()

	// vars := mux.Vars(r)
	// name := vars["name"]
	// email := vars["email"]

	db.Create(&models.User{
		Name:  user.Name,
		Email: user.Email,
	})

	fmt.Fprintf(rw, "User created successfully")

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Update User endpoint")

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Delete User endpoint")

}
