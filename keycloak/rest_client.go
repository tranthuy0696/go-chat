package keycloak

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

var (
	DefaultBaseURL  = os.Getenv("KEYCLOAK_BASE_URL")
	DefaultClientId = os.Getenv("KEYCLOAK_CLIENT_ID")
	DefaultRealm    = os.Getenv("KEYCLOAK_REALMS")
)

type RestClient struct {
	baseURL   string
	userAgent string
	config    clientcredentials.Config
}

type RestError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
}

//A RestClientOption sets an option on a RestClient
type RestClientOption func(*RestClient)

func PasswordAuth(username, password string) RestClientOption {
	return func(c *RestClient) {
		c.config.EndpointParams = url.Values{
			"grant_type": {"password"},
			"username":   {username},
			"password":   {password},
		}
	}
}

// getTokenURLPathForRealm returns the path component of the token URL of realm.
func getTokenURLPathForRealm(realm string) string {
	return "/realms/" + string(realm) + "/protocol/openid-connect/token"
}

// newRequest creats a new http.Request.
func newRequest(method, urlStr string, values url.Values, body io.Reader) (*http.Request, error) {
	if values != nil {
		urlStr += "?" + values.Encode()
	}
	return http.NewRequest(method, urlStr, body)
}

// newRequestJSON creates a new http.Request with a JSON body.
func newRequestJSON(method, urlStr string, body interface{}) (*http.Request, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := newRequest(method, urlStr, nil, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
