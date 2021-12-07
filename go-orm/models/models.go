package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
