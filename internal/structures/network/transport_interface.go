// go-podman/internal/structures/network/transport_interface.go
package network

import "net/http"

type TransportCreator interface {
	NewUnixTransport(socketPath string) http.RoundTripper
}
