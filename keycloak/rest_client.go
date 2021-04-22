package keycloak

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var (
	DefaultBaseURL       = ""
	DefaultClientId      = ""
	DefaultRealm         = ""
	DefaultAdminUsername = ""
	DefaultAdminPassword = ""
)

func init() {
	godotenv.Load(".env")
	DefaultBaseURL = os.Getenv("KEYCLOAK_BASE_URL")
	DefaultClientId = os.Getenv("KEYCLOAK_CLIENT_ID")
	DefaultRealm = os.Getenv("KEYCLOAK_REALMS")
	DefaultAdminUsername = os.Getenv("KEYCLOAK_ADMIN_USERNAME")
	DefaultAdminPassword = os.Getenv("KEYCLOAK_ADMIN_PASSWORD")
}

//A RestClientOption sets an option on a RestClient
type RestClientOption func(*RestClient)

func PasswordAuth(username, password string) RestClientOption {
	return func(c *RestClient) {
		c.Config.EndpointParams = url.Values{
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

func GetToken(username, password string) (TokenResponse, error) {

	body := url.Values{
		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
		"client_id":  {DefaultClientId},
	}

	res, err := http.PostForm(DefaultBaseURL+getTokenURLPathForRealm(DefaultRealm), body)
	var response TokenResponse
	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	errParseJson := json.NewDecoder(res.Body).Decode(&response)

	if errParseJson != nil {
		return response, errParseJson
	}
	return response, nil

}

func GetTokenFromAdminUser() (TokenResponse, error) {
	body := url.Values{
		"grant_type": {"password"},
		"username":   {DefaultAdminUsername},
		"password":   {DefaultAdminPassword},
		"client_id":  {DefaultClientId},
	}

	res, err := http.PostForm(DefaultBaseURL+getTokenURLPathForRealm(DefaultRealm), body)
	var response TokenResponse
	if err != nil {
		return response, err
	}

	defer res.Body.Close()
	// bodyRes, _ := io.ReadAll(res.Body)
	// sb := string(bodyRes)

	errParseJson := json.NewDecoder(res.Body).Decode(&response)

	if errParseJson != nil {
		return response, errParseJson
	}
	return response, nil
}

func CreateUser(data User) (*User, error) {
	token, err := GetTokenFromAdminUser()
	if err != nil {
		return nil, err
	}
	var response User
	urlCreateUser := DefaultBaseURL + "/admin/realms/" + DefaultRealm + "/users"
	client := &http.Client{}
	// var body = []byte(`{
	// 	"username": "thuthuytthuy1",
	// 	"lastName": "tran",
	// 	"enabled": true
	// }`)
	// req, err := http.NewRequest("POST", urlCreateUser, bytes.NewBuffer(body))
	data.Enabled = true
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", urlCreateUser, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	errParseJson := json.NewDecoder(res.Body).Decode(&response)
	if errParseJson != nil {
		return nil, errParseJson
	}
	return &response, nil

}
