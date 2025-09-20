package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	// Create explicit ServeMux for better control (Go 1.22+)
	mux := http.NewServeMux()

	// register /greet endpoint → handled by greet function
	mux.HandleFunc("GET /greet", greet)
	// register /users endpoint → handled by GetAllUser function
	mux.HandleFunc("GET /users", GetAllUser)
	// register /users/{user_id} endpoint with path parameters (Go 1.22+)
	mux.HandleFunc("GET /users/{user_id}", getUser)

	// start server on port 8080, log.Fatal stops program if server fails
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// Use Go 1.22+ path parameter extraction (replaces mux.Vars)
	userID := r.PathValue("user_id")
	fmt.Fprint(w, userID)
}
