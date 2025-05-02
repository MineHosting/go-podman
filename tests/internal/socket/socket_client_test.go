// Go-podman/tests/internal/socket/socket_client_test.go
package socket_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/MineHosting/go-podman/internal/socket"
)

func TestSocketClient_Send_Success(t *testing.T) {
	expectedBody := []byte(`{"ok": true}`)

	client := socket.NewSocketClient(
		nil,
		&MockRequestBuilder{},
		&MockResponseReader{Body: expectedBody},
		&MockResponseValidator{},
		&MockTransport{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(string(expectedBody))),
			},
		},
	)

	body, err := client.Send("GET", "/containers/json", nil, socket.Rootless)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !bytes.Equal(body, expectedBody) {
		t.Errorf("unexpected body: got %s, want %s", body, expectedBody)
	}
}

func TestSocketClient_Send_RequestBuildFails(t *testing.T) {
	client := socket.NewSocketClient(
		nil,
		&MockRequestBuilder{Err: errors.New("boom")},
		nil, nil, nil,
	)

	_, err := client.Send("GET", "/fail", nil, socket.Rootless)
	if err == nil || err.Error() == "" {
		t.Fatal("expected error, got nil")
	}
}

func TestSocketClient_Send_RequestFails(t *testing.T) {
	client := socket.NewSocketClient(
		nil,
		&MockRequestBuilder{},
		nil,
		nil,
		&MockTransport{Err: errors.New("connection refused")},
	)

	_, err := client.Send("GET", "/fail", nil, socket.Rootless)
	if err == nil || err.Error() == "" {
		t.Fatal("expected transport error, got nil")
	}
}

func TestSocketClient_Send_ReadFails(t *testing.T) {
	client := socket.NewSocketClient(
		nil,
		&MockRequestBuilder{},
		&MockResponseReader{Err: errors.New("read failed")},
		&MockResponseValidator{},
		&MockTransport{
			Response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString("ok")),
			},
		},
	)

	_, err := client.Send("GET", "/", nil, socket.Rootless)
	if err == nil || err.Error() == "" {
		t.Fatal("expected read error, got nil")
	}
}

func TestSocketClient_Send_InvalidStatus(t *testing.T) {
	client := socket.NewSocketClient(
		nil,
		&MockRequestBuilder{},
		&MockResponseReader{Body: []byte("unauthorized")},
		&MockResponseValidator{Err: errors.New("invalid status")},
		&MockTransport{
			Response: &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(bytes.NewBufferString("unauthorized")),
			},
		},
	)

	_, err := client.Send("GET", "/", nil, socket.Rootless)
	if err == nil || err.Error() == "" {
		t.Fatal("expected status validation error, got nil")
	}
}
