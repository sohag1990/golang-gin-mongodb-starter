# golang-gin-mongodb-starter

Working example of most popular golang gin framework's mongoDB connection and REST API post method.
 To run this file change the hostname/hostip of your mongodb server
 session, err := mgo.Dial("172.17.0.2")
 
 Then simple run
 ```
 go run main.go
 ```
 Then curl Request to check.
 ```
curl -X POST -H "Content-Type: application/json" -d '{"task":"i never lose i win or i learn!!!"}' localhost:8000/api/v1/task
 ```
 
