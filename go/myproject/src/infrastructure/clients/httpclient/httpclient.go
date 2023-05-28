package httpclient

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

type Config struct {
	ContentType   string
	Accept        string
	UserAgent     string
	Username      string
	Password      string
	RequestedWith string
	Authorization string
	BaseURL       string
	Timeout       time.Duration
}

type httpClient struct {
	config Config
}

func New(config Config) HttpClient {
	if config.ContentType == "" {
		config.ContentType = "application/json"
	}
	if config.Accept == "" {
		config.Accept = "application/json"
	}
	if config.UserAgent == "" {
		config.UserAgent = "{{ .ProjectName }}"
	}
	return &httpClient{config}
}

func (h *httpClient) GET(url string) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	return h.doRequest(url, req)
}

func (h *httpClient) POST(url string, data []byte) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.SetBody(data)
	return h.doRequest(url, req)
}

func (h *httpClient) PUT(url string, data []byte) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("PUT")
	req.SetBody(data)
	return h.doRequest(url, req)
}

func (h *httpClient) DELETE(url string) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("DELETE")
	return h.doRequest(url, req)
}

func (h *httpClient) PATCH(url string, data []byte) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("PATCH")
	req.SetBody(data)
	return h.doRequest(url, req)
}

func (h *httpClient) Failed(err error, statusCode int) bool {
	return err != nil || statusCode >= 400
}

func SaveRequestInFile(req *http.Request, filename string) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	if filename != "" {
		return os.WriteFile(filename, body, 0644)
	}
	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (h *httpClient) doRequest(url string, req *fasthttp.Request) ([]byte, int, error) {
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if h.config.Username != "" || h.config.Password != "" {
		req.Header.Set(fasthttp.HeaderAuthorization, "Basic "+basicAuth(h.config.Username, h.config.Password))
	}
	if h.config.RequestedWith != "" {
		req.Header.Set("X-Requested-With", h.config.RequestedWith)
	}
	if h.config.Authorization != "" {
		req.Header.Set("Authorization", "Bearer "+h.config.Authorization)
	}

	if h.config.BaseURL != "" {
		url = fmt.Sprintf("%s%s", h.config.BaseURL, url)
	}
	req.SetRequestURI(url)
	req.Header.SetContentType(h.config.ContentType)
	req.Header.Set(fasthttp.HeaderAccept, h.config.Accept)
	req.Header.SetUserAgent(h.config.UserAgent)
	timeout := h.config.Timeout * time.Second
	err := fasthttp.DoTimeout(req, resp, timeout)
	if err != nil {
		return []byte(""), 0, err
	}

	statusCode := resp.StatusCode()
	bodyBytes := resp.Body()
	return bodyBytes, statusCode, nil
}
