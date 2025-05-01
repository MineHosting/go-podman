// go-podman/internal/structures/network/request_builder_interface.go
package network

import (
	"io"
	"net/http"
)

type HTTPRequestBuilder interface {
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
}
