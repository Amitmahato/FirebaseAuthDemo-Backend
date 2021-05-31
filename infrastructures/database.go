package infrastructures

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func SetupDB() *gorm.DB {
	log.Println("Setting up database")

	// credentials for making connection to the database
	DB_USER := "<user>"
	DB_PASS := "<password>"
	DB_HOST := "localhost"
	DB_PORT := "3306"
	DB_NAME := "FireBaseAuthDemo"

	// construct connection url through which database connection is made and use gorm to establish a connection
	connectionUrl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(connectionUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Failure establishing database connection", err)
	}
	return db
}
