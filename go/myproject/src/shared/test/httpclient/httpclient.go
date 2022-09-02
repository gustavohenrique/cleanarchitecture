package httpclient

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"{{ .ProjectName }}/src/shared/conf"
)

type HTTPClient struct {
	Cookie      *http.Cookie
	Token       string
	TokenHeader string
	Username    string
	Password    string
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

func WithTokenInHeader(header, token string) HTTPClient {
	return HTTPClient{
		TokenHeader: header,
		Token:       token,
	}
}

func WithOauth2Mock() HTTPClient {
	config := conf.Get()
	header := config.Auth.Jwt.Header
	return HTTPClient{
		TokenHeader: header,
		Token:       "any token because authService is mocked",
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
		return req
	}
	if h.Token != "" {
		return h.getReqWithTokenHeader(req)
	}
	credentials := h.Username + ":" + h.Password
	auth := base64.StdEncoding.EncodeToString([]byte(credentials))
	req.Header.Set("Authorization", "basic "+auth)
	return req
}

func (h HTTPClient) getReqWithTokenHeader(req *http.Request) *http.Request {
	if h.TokenHeader != "" {
		req.Header.Set(h.TokenHeader, h.Token)
		return req
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.Token))
	return req
}
