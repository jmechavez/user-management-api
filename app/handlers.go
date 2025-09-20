package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// User represents a basic user model.
type User struct {
	Name  string `json:"name" xml:"name"`   // user's name, mapped to JSON key "name"
	Email string `json:"email" xml:"email"` // user's email, mapped to JSON key "email"
	ID    int    `json:"id" xml:"id"`       // user's ID, mapped to JSON key "id"
}

// greet handles /greet requests.
func greet(w http.ResponseWriter, r *http.Request) {
	// write plain text response to client
	fmt.Fprintf(w, "Greetings!")
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	// create a slice (list) of User structs with sample data
	users := []User{
		{"John Doe", "johndoe@test.com", 1234},
		{"test", "test01@test.com", 2345},
	}
	// set response header so client knows it’s XML if not JSON
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(users)
	} else {
		// set response header so client knows it’s JSON
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}
