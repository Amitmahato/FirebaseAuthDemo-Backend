package main

import (
	"fmt"
	"log"
	"net/http"

	"firebase-authentication-backend/infrastructures"

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

	// listen and serve on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080",httpRouter))
}
