// user-management-api: simple HTTP server with /greet and /users endpoints.
package main

import (
	"encoding/json" // for encoding Go structs to JSON
	"fmt"           // for formatted I/O
	"log"           // for logging errors
	"net/http"      // for HTTP server and request handling
)

// User represents a basic user model.
type User struct {
	Name  string `json:"name"`  // user's name, mapped to JSON key "name"
	Email string `json:"email"` // user's email, mapped to JSON key "email"
	ID    int    `json:"id"`    // user's ID, mapped to JSON key "id"
}

// main registers routes and starts the server.
func main() {
	// register /greet endpoint → handled by greet function
	http.HandleFunc("/greet", greet)

	// register /users endpoint → handled by GetUser function
	http.HandleFunc("/users", GetUser)

	// start server on port 8080, log.Fatal stops program if server fails
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// greet handles /greet requests.
func greet(w http.ResponseWriter, r *http.Request) {
	// write plain text response to client
	fmt.Fprintf(w, "Greetings!")
}

// GetUser handles /users requests and returns mock data.
func GetUser(w http.ResponseWriter, r *http.Request) {
	// create a slice (list) of User structs with sample data
	users := []User{
		{"John Doe", "johndoe@test.com", 1234},
		{"test", "test01@test.com", 2345},
	}

	// set response header so client knows it’s JSON
	w.Header().Set("Content-Type", "application/json")

	// encode users slice as JSON and write to response
	json.NewEncoder(w).Encode(users)
}
