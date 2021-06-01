package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

	firebaseApp := infrastructures.InitializeFirebase()
	firebaseAuth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		panic("Error Creating Firebase Client")
	}
	
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

	// authentication middlerware
	userAuthMiddleware := func (next http.Handler) http.Handler {	
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[INFO] %v\n", r.URL.Path)
		
			authorizationToken := r.Header.Get("Authorization")
			// idToken is different from customToken and customToken should be used in firebase client app in signInWithCustomToken
			idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer ", "", 1))

			ctx := context.Background()
			fbUser, err := firebaseAuth.VerifyIDToken(ctx, idToken)

			if err != nil {
				log.Println("User authentication failure, ", err)
				return
			}

			var user model.Users
			err = db.Find(&user, "UID = ?", fbUser.UID).Error
			if err != nil {
				log.Println("User with given UID not found, ", err)
				return
			}

			ctx = r.Context()
			ctx = context.WithValue(ctx, "user_id", user.Id)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}

	// create a authenticated route group for user
	authenticatedUserRoute := httpRouter.PathPrefix("/user").Subrouter()
	authenticatedUserRoute.Use(userAuthMiddleware)

	authenticatedUserRoute.HandleFunc("", func(rw http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("user_id").(uint64)
		
		var user model.Users
		err := db.First(&user, id).Error

		if err != nil {
			log.Println("User with given credential doesn't exist")

			msg := "User doesn't exist"
			payload := map[string]string{"error": msg}
			response, _ := json.Marshal(payload)
			
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write(response)
			return
		}

		var _user struct {
			Id        uint64 `json:"user_id"`
			UID       string `json:"user_uid"`
			FirstName string `json:"user_firstName"`
			LastName  string `json:"user_lastName"`
			Email     string `json:"user_email"`
			Phone     string `json:"user_phone"`
		}

		_user.Id = user.Id
		_user.UID = user.UID
		_user.FirstName = user.FirstName
		_user.LastName = user.LastName
		_user.Email = user.Email
		_user.Phone = user.Phone

		response, _ := json.Marshal(_user)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(response)
	})


	// listen and serve on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080",httpRouter))
}
