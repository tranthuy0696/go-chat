package authentication

import (
	"encoding/json"
	"go-chat/keycloak"
	"net/http"
)

type Authen struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body Authen
	json.NewDecoder(r.Body).Decode(&body)
	token, err := keycloak.GetToken(body.Username, body.Password)
	if err != nil {
		json.NewEncoder(w).Encode(keycloak.TokenResponse{})
		return
	}
	json.NewEncoder(w).Encode(token)
}
