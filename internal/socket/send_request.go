// Go-Podman/internal/socket/send_request.go
package socket

import (
	"fmt"
	"net/http"

	"github.com/MineHosting/go-podman/internal/network"
)

func SendRequest(method, url string, body any, socketPath string) ([]byte, error) {
	serializedBody, err := network.SerializePayload(body)
	if err != nil {
		return nil, fmt.Errorf("[Socket]: failed to serialize payload: %w", err)
	}

	req, err := network.NewRequest(method, url, serializedBody)
	if err != nil {
		return nil, fmt.Errorf("[Socket]: failed to create request: %w", err)
	}

	client := &http.Client{
		Transport: network.NewUnixTransport(socketPath),
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[Socket]: request failed: %w", err)
	}

	defer resp.Body.Close()

	bodyResp, err := network.ReadBody(resp)
	if err != nil {
		return nil, fmt.Errorf("[Socket]: failed to read response body: %w", err)
	}

	err = network.ValidateStatus(resp, bodyResp)
	if err != nil {
		return nil, fmt.Errorf("[Network]: invalid response: %w", err)
	}

	return bodyResp, nil
}
