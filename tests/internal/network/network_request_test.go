// go-podman/tests/internal/network/network_request_test.go
package network_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/MineHosting/go-podman/internal/network"
)

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name    string
		method  string
		url     string
		body    io.Reader
		wantErr bool
	}{
		{"Valid GET", http.MethodGet, "http://localhost", nil, false},
		{"Valid POST with body", http.MethodPost, "http://localhost", strings.NewReader(`{"a":1}`), false},
		{"Invalid URL", http.MethodGet, "://", nil, true},
	}

	requestBuilder := network.RealHTTPRequestBuilder{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := requestBuilder.NewRequest(tt.method, tt.url, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && req.Header.Get("Content-Type") != "application/json" {
				t.Errorf("Expected Content-Type to be application/json")
			}
		})
	}
}
