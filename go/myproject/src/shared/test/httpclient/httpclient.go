package httpclient

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

type HTTPClient struct {
	Cookie   *http.Cookie
	Token    string
	Username string
	Password string
}

func DoGET(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

func DoDELETE(url string) *http.Request {
	req, _ := http.NewRequest("DELETE", url, nil)
	return req
}

func DoPOST(url, jsonReq string) *http.Request {
	return makeRequest("POST", url, jsonReq)
}

func DoPUT(url, jsonReq string) *http.Request {
	return makeRequest("PUT", url, jsonReq)
}

func makeRequest(method, url, jsonReq string) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(jsonReq)))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func WithCookie(name, value string) HTTPClient {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	return HTTPClient{Cookie: cookie}
}

func WithTokenInHeader(token string) HTTPClient {
	return HTTPClient{
		Token: token,
	}
}

func WithBasicAuth(username, password string) HTTPClient {
	return HTTPClient{
		Username: username,
		Password: password,
	}
}

func (h HTTPClient) DoPOST(url, jsonReq string) *http.Request {
	req := DoPOST(url, jsonReq)
	return h.addTokenIn(req)
}

func (h HTTPClient) DoGET(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return h.addTokenIn(req)
}

func (h HTTPClient) DoPUT(url, jsonReq string) *http.Request {
	req := DoPUT(url, jsonReq)
	return h.addTokenIn(req)
}

func (h HTTPClient) addTokenIn(req *http.Request) *http.Request {
	if h.Cookie != nil {
		req.AddCookie(h.Cookie)
	} else if h.Token != "" {
		log.Println("Authorization:", fmt.Sprintf("Bearer %s", h.Token))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Token))
	} else {
		credentials := h.Username + ":" + h.Password
		auth := base64.StdEncoding.EncodeToString([]byte(credentials))
		req.Header.Set("Authorization", "basic "+auth)
	}
	return req
}
