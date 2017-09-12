package goKeycloakAdapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// UserInfo is used to expose user's info
type UserInfo struct {
	name  string
	email string
}

// GetUserInfo get user info
func GetUserInfo(request *http.Request) {
	url := "http://localhost:8282/auth/realms/authentication-server/protocol/openid-connect/userinfo"
	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJQRjNfaURDY1pRcGpzTGlvc3U5UVJCMlhKRHYtT3hZdS16dUZmYmJ1elpvIn0.eyJqdGkiOiI0ZDI4MjRlZS04YThiLTRjNTgtOGIyNy1mY2JhM2ExOTYyYjkiLCJleHAiOjE1MDUyMTM4MDMsIm5iZiI6MCwiaWF0IjoxNTA1MjEzMjAzLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgyODIvYXV0aC9yZWFsbXMvYXV0aGVudGljYXRpb24tc2VydmVyIiwiYXVkIjoiZ28ta2V5Y2xvYWstYWRhcHRlciIsInN1YiI6ImY5NjFjNWYzLTRkNzQtNDc2MS05MTBjLTRkZTM1ZTc0ZDJkNyIsInR5cCI6IkJlYXJlciIsImF6cCI6ImdvLWtleWNsb2FrLWFkYXB0ZXIiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiJlODk2ZmI0Ni01MDYwLTRkYjYtYTJkNy1jY2RiODc0MGFhNzUiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly9nb29nbGUuY29tIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJ1bWFfYXV0aG9yaXphdGlvbiIsImFzc2lzdGFuY2UiXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJuYW1lIjoiRGluaCBQaGFuIiwicHJlZmVycmVkX3VzZXJuYW1lIjoicGhhbnZ1ZGluaCIsImdpdmVuX25hbWUiOiJEaW5oIiwiZmFtaWx5X25hbWUiOiJQaGFuIiwiZW1haWwiOiJpdHBoYW52dWRpbmhAZ21haWwuY29tIn0.c2u9UzDa69TWYFu_L8bSKKAHqLe4o5qvqImNB935_gp9jnnqLufnxOtCFloDlNxfyGWzxklHf3ufgfwRGAdNY9xvTFLX2s8dEX-z7uRn5PmsotxxpAit4PEyK1HIBD-SNyvJHO7Whu4xOiadaxYl2mxBHOULCyxw3DS5S6iKLo1k06YD9iU1h3grgQnMf7PmJETiJqQUEANec7p5dTYXaxBj8pcGCNpNDYSc3hmMpphvelTuPoFCxKXCkx1vNuUBFI97xDKX5Jvt_gd8SlMBdY0mr6lM5-pdoN8mU9PR6AaRy62Zc_fRuB8h3ALesAjo2pbHlWDup09PmS8133UXSg"
	client := http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "bearer "+token)

	res, _ := client.Do(req)

	fmt.Println("status code " + strconv.Itoa(res.StatusCode))
	if res.StatusCode == 200 { // OK
		body, _ := ioutil.ReadAll(res.Body)
		var data map[string]interface{}
		json.Unmarshal(body, &data)
		fmt.Println(data["name"])
		fmt.Println(data["email"])
		fmt.Println(data["given_name"])
	}
}

// IsAuthorized validate user'authorization
func IsAuthorized(request *http.Request) bool {
	requestURL := "http://localhost:8282/auth/realms/authentication-server/protocol/openid-connect/token/introspect"

	form := url.Values{
		"client_id":     {"go-keycloak-adapter"},
		"client_secret": {"d73627f6-87a1-45db-bb7a-314fd4d89dd4"},
		"token":         {"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJQRjNfaURDY1pRcGpzTGlvc3U5UVJCMlhKRHYtT3hZdS16dUZmYmJ1elpvIn0.eyJqdGkiOiI0ZDI4MjRlZS04YThiLTRjNTgtOGIyNy1mY2JhM2ExOTYyYjkiLCJleHAiOjE1MDUyMTM4MDMsIm5iZiI6MCwiaWF0IjoxNTA1MjEzMjAzLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgyODIvYXV0aC9yZWFsbXMvYXV0aGVudGljYXRpb24tc2VydmVyIiwiYXVkIjoiZ28ta2V5Y2xvYWstYWRhcHRlciIsInN1YiI6ImY5NjFjNWYzLTRkNzQtNDc2MS05MTBjLTRkZTM1ZTc0ZDJkNyIsInR5cCI6IkJlYXJlciIsImF6cCI6ImdvLWtleWNsb2FrLWFkYXB0ZXIiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiJlODk2ZmI0Ni01MDYwLTRkYjYtYTJkNy1jY2RiODc0MGFhNzUiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly9nb29nbGUuY29tIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJ1bWFfYXV0aG9yaXphdGlvbiIsImFzc2lzdGFuY2UiXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJuYW1lIjoiRGluaCBQaGFuIiwicHJlZmVycmVkX3VzZXJuYW1lIjoicGhhbnZ1ZGluaCIsImdpdmVuX25hbWUiOiJEaW5oIiwiZmFtaWx5X25hbWUiOiJQaGFuIiwiZW1haWwiOiJpdHBoYW52dWRpbmhAZ21haWwuY29tIn0.c2u9UzDa69TWYFu_L8bSKKAHqLe4o5qvqImNB935_gp9jnnqLufnxOtCFloDlNxfyGWzxklHf3ufgfwRGAdNY9xvTFLX2s8dEX-z7uRn5PmsotxxpAit4PEyK1HIBD-SNyvJHO7Whu4xOiadaxYl2mxBHOULCyxw3DS5S6iKLo1k06YD9iU1h3grgQnMf7PmJETiJqQUEANec7p5dTYXaxBj8pcGCNpNDYSc3hmMpphvelTuPoFCxKXCkx1vNuUBFI97xDKX5Jvt_gd8SlMBdY0mr6lM5-pdoN8mU9PR6AaRy62Zc_fRuB8h3ALesAjo2pbHlWDup09PmS8133UXSg"},
	}

	// func Post(url string, bodyType string, body io.Reader) (resp *Response, err error) {
	body := bytes.NewBufferString(form.Encode())
	res, _ := http.Post(requestURL, "application/x-www-form-urlencoded", body)

	fmt.Println("status code " + strconv.Itoa(res.StatusCode))

	if res.StatusCode == 200 { // OK
		body, _ := ioutil.ReadAll(res.Body)
		var data map[string]interface{}
		json.Unmarshal(body, &data)
		fmt.Println(data["active"])
	}

	return true
}
