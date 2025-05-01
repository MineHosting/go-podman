// go-podman/internal/network/read_body.go
package network

import (
	"fmt"
	"io"
	"net/http"
)

type RealResponseReader struct{}

func (r *RealResponseReader) ReadBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[Network]: failed in read response: %w", err)
	}
	return body, nil
}
