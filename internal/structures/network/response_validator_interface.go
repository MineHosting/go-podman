// go-podman/internal/structures/network/response_validator_interface.go
package network

import "net/http"

type ResponseValidator interface {
	ValidateStatus(resp *http.Response, body []byte) error
}
