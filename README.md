Building a Go-Powered Event Booking API
1. GET /events → Get a list of available events
2. GET /events/<ID> → Get a specific event by ID
3. POST /events → Create a new bookable event [Authentication Required]
4. PUT /events/<ID> → Update an event [Authentication Required]
5. DELETE /events/<ID> → Delete an event [Authentication Required]
6. POST /signup → Create a new user
7. POST /login → Authenticate user & obtain JWT token [JWT]
8. POST /events/<ID>/register → Register user for an event [Authentication Required]
9. DELETE /events/<ID>/register → Cancel registration [Authentication Required]


Gin library is used instead of the http package because of its more advanced features

1. install the go mod and create main file

2. install the gin package
go get -u github.com/gin-gonic/gin

3. set us a sample routing and events handling to test your server 

4. set up an event model


