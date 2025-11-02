package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strconv"
	"strings"

	apperrors "github.com/jmechavez/user-management-api/internal/appErrors"
	"github.com/jmechavez/user-management-api/internal/services"
)

type UserHandler struct {
	service services.UserServices
}

// func (uh *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
// 	users, _ := uh.service.FindAllUser()
// 	if r.Header.Get("Content-Type") != "application/json" {
// 		writeJSONResponse(w, http.StatusOK, users)
// 		// w.Header().Add("Content-Type", "application/json")
// 		// json.NewEncoder(w).Encode(users)
// 	} else {
// 		w.Header().Set("Content-Type", "application/xml")
// 		xml.NewEncoder(w).Encode(users)
// 	}
// }

func (uh *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, appErr := uh.service.FindAllUser()
	if appErr != nil {
		writeJSONResponse(w, appErr.Code, appErr)
		return
	}

	accept := r.Header.Get("Accept")

	if strings.Contains(accept, "application/xml") {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(users)
	} else {
		// Default to JSON
		writeJSONResponse(w, http.StatusOK, users)
	}
}

func (uh *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path, e.g. /users/123
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		idError := apperrors.NewInvalidUserIDError()
		writeJSONResponse(w, http.StatusBadRequest, idError)
		// http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, appErr := uh.service.FindById(id)
	if appErr != nil {
		writeJSONResponse(w, appErr.Code, appErr)
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(appErr.Code)
		// json.NewEncoder(w).Encode(appErr)
	} else {
		// Return the user as JSON
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(user)
		writeJSONResponse(w, http.StatusOK, user)
	}
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
