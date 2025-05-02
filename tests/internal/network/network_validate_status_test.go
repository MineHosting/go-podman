package network_test

import (
	"net/http"
	"testing"

	"github.com/MineHosting/go-podman/internal/network"
)

func TestValidateStatus(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		body    []byte
		wantErr bool
	}{
		{"200 OK", 200, []byte("ok"), false},
		{"201 Created", 201, []byte("created"), false},
		{"400 Bad Request", 400, []byte("bad"), true},
		{"500 Server Error", 500, []byte("boom"), true},
	}

	statusValidator := network.RealResponseValidator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{StatusCode: tt.status}
			err := statusValidator.ValidateStatus(resp, tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
