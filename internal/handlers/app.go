package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmechavez/user-management-api/internal/domain"
	"github.com/jmechavez/user-management-api/internal/services"
)

func Start() {
	router := mux.NewRouter()

	uh := UserHandler{services.NewUserService(domain.NewUserRepositoryStub())}
	router.HandleFunc("/users", uh.FindAll).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
