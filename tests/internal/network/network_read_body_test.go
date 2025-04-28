package network_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/MineHosting/go-podman/internal/network"
)

func TestReadBody(t *testing.T) {
	tests := []struct {
		name       string
		body       io.ReadCloser
		want       string
		shouldFail bool
	}{
		{
			name:       "Valid body",
			body:       io.NopCloser(strings.NewReader("response")),
			want:       "response",
			shouldFail: false,
		},
		{
			name:       "Invalid read",
			body:       io.NopCloser(brokenReader{}),
			want:       "",
			shouldFail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{Body: tt.body}
			data, err := network.ReadBody(resp)
			if (err != nil) != tt.shouldFail {
				t.Errorf("ReadBody() error = %v, shouldFail = %v", err, tt.shouldFail)
			}
			if !tt.shouldFail && string(data) != tt.want {
				t.Errorf("Expected body = %q, got %q", tt.want, data)
			}
		})
	}
}
