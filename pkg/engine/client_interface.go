// Go-Podman/pkg/engine/client_interface
package engine

type Client interface {
	Send(method, endpoint string, body any) ([]byte, error)
	ChangeApiVersion() error
}
