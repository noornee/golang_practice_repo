package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Printf("failed to connect to database, %v", err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var users []User

	db.Find(&users)
	json.NewEncoder(rw).Encode(users)
}


func GetUserByID(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var users []User

	db.Find(&users, id)
	json.NewEncoder(rw).Encode(users)

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	json.Unmarshal(reqBody, &user)

	db.Create(&User{
		Name:  user.Name,
		Email: user.Email,
	})

	fmt.Fprintf(rw, "User created successfully")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(r)
	id := vars["id"]


	var user User
	json.Unmarshal(reqBody, &user)

	// db.Find(&user, id).Updates(map[string]interface{}{
	// 	"Name": user.Name,
	// 	"Email": user.Email,
	// })


	db.Model(&user).Where("ID", id).Updates(map[string]interface{}{
		"Name": user.Name,
		"Email": user.Email,
	})

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	db.Delete(&user, id)

	fmt.Fprintf(rw, "Deleted")


}
