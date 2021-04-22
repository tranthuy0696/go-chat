package keycloak

import (
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type RestClient struct {
	BaseURL   string
	UserAgent string
	Config    clientcredentials.Config
}

type RestError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
}

type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type User struct {
	ID                         string               `json:"id,omitempty"`
	CreatedTimestamp           int64                `json:"createdTimestamp,omitempty"`
	Username                   string               `json:"username,omitempty"`
	Enabled                    bool                 `json:"enabled,omitempty"`
	Totp                       bool                 `json:"totp,omitempty"`
	EmailVerified              bool                 `json:"emailVerified,omitempty"`
	FirstName                  string               `json:"firstName,omitempty"`
	LastName                   string               `json:"lastName,omitempty"`
	Email                      string               `json:"email,omitempty"`
	FederationLink             string               `json:"federationLink,omitempty"`
	Attributes                 *map[string][]string `json:"attributes,omitempty"`
	DisableableCredentialTypes *[]interface{}       `json:"disableableCredentialTypes,omitempty"`
	RequiredActions            *[]string            `json:"requiredActions,omitempty"`
	Access                     *map[string]bool     `json:"access,omitempty"`
	ClientRoles                *map[string][]string `json:"clientRoles,omitempty"`
	RealmRoles                 *[]string            `json:"realmRoles,omitempty"`
	Groups                     *[]string            `json:"groups,omitempty"`
	ServiceAccountClientID     string               `json:"serviceAccountClientId,omitempty"`
}
