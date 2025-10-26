package handlers

import (
	"log"
	"net/http"

	"github.com/jmechavez/user-management-api/internal/domain"
	"github.com/jmechavez/user-management-api/internal/services"
)

// methodHandler wraps a handler function and only allows specific HTTP methods
func methodHandler(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

func Start() {
	// Initialize the user handler
	uh := UserHandler{services.NewUserService(domain.NewUserRepositoryDb())}
	// uh := UserHandler{services.NewUserService(domain.NewUserRepositoryStub())}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register routes with method checking
	mux.HandleFunc("/users", methodHandler(http.MethodGet, uh.FindAll))

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
