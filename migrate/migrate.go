package main

import (
	"firebase-authentication-backend/infrastructures"
	"firebase-authentication-backend/model"
	"fmt"
)

func main() {
	db := infrastructures.SetupDB()
	
	userTableExists := db.Migrator().HasTable(&model.Users{})

	fmt.Println("User table exists ? ", userTableExists)

	if !userTableExists {
		db.AutoMigrate(&model.Users{})
	}
}

