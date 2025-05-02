package network_test

import (
	"bytes"
	"testing"

	"github.com/MineHosting/go-podman/internal/network"
)

func TestSerializePayload(t *testing.T) {
	tests := []struct {
		name     string
		payload  any
		expected string
		wantErr  bool
	}{
		{"Nil payload", nil, "", false},
		{"Valid map", map[string]int{"a": 1}, `{"a":1}`, false},
		{"Unsupported payload", make(chan int), "", true},
	}

	serializePayloader := network.RealPayloadSerializer{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := serializePayloader.SerializePayload(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("SerializePayload() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && reader != nil {
				buf := new(bytes.Buffer)
				buf.ReadFrom(reader)
				got := buf.String()
				if got != tt.expected {
					t.Errorf("Expected payload = %q, got %q", tt.expected, got)
				}
			}
		})
	}
}
