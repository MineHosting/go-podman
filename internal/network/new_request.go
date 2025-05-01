// go-podman/internal/network/new_request.go
package network

import (
	"fmt"
	"io"
	"net/http"
)

type RealHTTPRequestBuilder struct{}

func (r *RealHTTPRequestBuilder) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("[Network]: error in create a new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
