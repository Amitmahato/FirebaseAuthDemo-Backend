---------------------------- Standalone Server -----------------------------------

1. Setup a router and start a server on port 8080 and handle / path - test http://localhost:8080 is reachable
2. establish a database connection	- test if database connection is established?
3. define a model	& migrate the model to database to create a table - structure of table to store our user information and how the fields should appear in json	& add migration logic, utilize & explain gorm automigrate
4. add controller for signup and create a entry in database - signup route should be reachable, handle json from req.body and make an entry in users table

------------------- Requires FrontEnd To Be Taken Parallely ----------------------

5. add firebase app initializaiton and on signup should create user in firebase - instruct how to get serviceAccountKey.json, use it to initalize firebase app, and user firebase app to create user 
6. add some authenticated route - initially show it can be assessed even without logging into user account
7. add authentication middleware for firebase idToken verification & make authenticated route use it - demo this middleware is hit before the controll is sent to respective controller & demo authenticated route is no longer availabe to not signed in user
