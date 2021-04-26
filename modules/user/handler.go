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

	}
	token, err := keycloak.GetToken(body.Username, body.Password)
	if err != nil {
		json.NewEncoder(w).Encode(keycloak.TokenResponse{})
		return
	}
	json.NewEncoder(w).Encode(token)
}
