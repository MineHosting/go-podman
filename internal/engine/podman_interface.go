// Go-Podman/internal/engine/podman_interface
package engine

import "net/http"

type Client interface {
	RawSend(method, endpoint string, body any) (*http.Request, error)
	Send(method, endpoint string, body any) ([]byte, error)
	ChangeApiVersion() error
}
