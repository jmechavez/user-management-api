package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/jmechavez/user-management-api/internal/services"
)

type UserHandler struct {
	service services.UserServices
}

func (uh *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, _ := uh.service.FindAllUser()
	if r.Header.Get("Content-Type") != "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	} else {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(users)
	}
}
