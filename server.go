package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func main() {
	httpRouter := mux.NewRouter()

	// testing if req to http://localhost:8080 is reachable or not
	httpRouter.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw,"Server is up and running")
	})

	// listen and serve on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080",httpRouter))
}
