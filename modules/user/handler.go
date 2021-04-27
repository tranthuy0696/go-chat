package user

import (
	"encoding/json"
	"go-chat/keycloak"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body User
	json.NewDecoder(r.Body).Decode(&body)
	if body.Password != body.PasswordConfirm {
		http.Error(w, "Passsword did not match!", http.StatusForbidden)
		return
	}
	user := keycloak.User{
		Username:  body.Username,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	response, err := keycloak.CreateUser(user)
	if err != nil {
		http.Error(w, "Bad Request!", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(response)
}
