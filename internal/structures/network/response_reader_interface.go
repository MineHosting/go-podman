// go-podman/internal/structures/network/response_reader_interface.go
package network

import "net/http"

type ResponseReader interface {
	ReadBody(req *http.Response) ([]byte, error)
}
