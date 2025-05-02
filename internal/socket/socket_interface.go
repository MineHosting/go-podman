// go-podman/internal/socket/socket_interface.go
package socket

import "io"

type SocketClientInterface interface {
	Send(method, url string, body io.Reader, socket SocketPath) ([]byte, error)
}
