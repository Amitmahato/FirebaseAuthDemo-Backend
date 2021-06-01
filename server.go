package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"firebase-authentication-backend/infrastructures"
	"firebase-authentication-backend/model"

	"github.com/golang/gddo/httputil/header"

	"github.com/gorilla/mux"
)



func main() {
	// setup a databse connection using gormdb
	db := infrastructures.SetupDB()

	// jsut to use db instance so `go` don't throw error
	log.Println("Database server is ",db.Name())

	httpRouter := mux.NewRouter()

	// testing if req to http://localhost:8080 is reachable or not
	httpRouter.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw,"Server is up and running")
	})

	httpRouter.HandleFunc("/users/signup", func(rw http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "" {
			value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
			if value != "application/json" {
				msg := "Content-Type header is not application/json"

				payload := map[string]string{"error": msg}
				response, _ := json.Marshal(payload)
				
				rw.Header().Set("Content-Type", "application/json")
				rw.WriteHeader(http.StatusUnsupportedMediaType)
				rw.Write(response)
				return
			}
		}

		var user model.Users

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&user)		
		if err != nil {
			log.Println("Failed to decode json object to users struct")

			msg := "Bad Request"
			payload := map[string]string{"error": msg}
			response, _ := json.Marshal(payload)
			
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(response)
			return
		}

		err = db.Create(&user).Error
		if err != nil {
			log.Println("Failed to insert user into the users table")

			msg := "Bad Request"
			payload := map[string]string{"error": msg}
			response, _ := json.Marshal(payload)
			
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(response)
			return
		}

		payload := map[string]string{"msg": "User created successfully" }
		response, _ := json.Marshal(payload)
		
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(response)
	})

	// listen and serve on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080",httpRouter))
}
