package user

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	Svc service
}

func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body"))
		log.Printf("Could not decode json: %s", err)
		return
	}
	// Call the AddUser function
	message, err := h.Svc.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to add user"))
		log.Printf("Failed to add user: %s", err)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
