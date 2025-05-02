// Go-Podman/internal/engine/podman_interface
package engine

type Client interface {
	Send(method, endpoint string, body any) ([]byte, error)
	ChangeApiVersion() error
}
