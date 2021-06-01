
# Getting the project running on your machine locally
## Backend
1. Create a my sql database with name `FireBaseAuthDemo` or with any name and replace the values in the `infrastructures/database` for database username, password & database name.

2. Install the packages with ```go install```
3. Run ```go run ./migrate/migrate.go``` to make the migration of users table.
4. Build the app with ```go build```
5. Run the executable with ```./firebase-authentication-backend```

## Frontend repository - [FirebaseAuthDemo-Frontend](https://github.com/Amitmahato/FirebaseAuthDemo-Frontend)



# Topics Covered
## Standalone Server
- Note that each step has its separate branch available on github
1. Setup a router and start a server on port 8080 and handle / path - test http://localhost:8080 is reachable
2. establish a database connection	- test if database connection is established?
3. define a model	& migrate the model to database to create a table - structure of table to store our user information and how the fields should appear in json	& add migration logic, utilize & explain gorm automigrate
4. add controller for signup and create a entry in database - signup route should be reachable, handle json from req.body and make an entry in users table

## Requires FrontEnd To Be Taken Parallely
- All the topics covere below is in a single branch [5.FirebaseAppInitialization](https://github.com/Amitmahato/FirebaseAuthDemo-Backend/tree/5.FirebaseAppInitialization)
5. add firebase app initializaiton and on signup should create user in firebase - instruct how to get serviceAccountKey.json, use it to initalize firebase app, and user firebase app to create user 
6. add some authenticated route - initially show it can be assessed even without logging into user account
7. add authentication middleware for firebase idToken verification & make authenticated route use it - demo this middleware is hit before the controll is sent to respective controller & demo authenticated route is no longer availabe to not signed in user

