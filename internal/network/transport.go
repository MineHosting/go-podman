// go-podman/internal/network/transport.go
package network

import (
	"context"
	"net"
	"net/http"
	"time"
)

func NewUnixTransport(socketPath string) http.RoundTripper {
	return &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial("unix", socketPath)
		},
		MaxIdleConns:          10,
		IdleConnTimeout:       30 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
	}
}
