package calculating

import (
	"net/http"
	"time"
)

// APIClient HTTP client wrapper
type APIClient struct {
	doHTTPReq func(*http.Request) (*http.Response, error)
}

// NewClient initialises a new API client
func NewClient() *APIClient {
	client := &http.Client{Timeout: 60 * time.Second}
	doReq := func(req *http.Request) (*http.Response, error) {
		return client.Do(req)
	}

	return &APIClient{
		doHTTPReq: doReq,
	}
}
