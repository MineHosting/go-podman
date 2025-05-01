// Go-Podman/internal/socket/send_request.go
package socket

import (
	"fmt"
	"net/http"

	"github.com/MineHosting/go-podman/internal/network"
)

func SendRequest(method, url string, body any, socketPath string) ([]byte, error) {
	serializer := &network.RealPayloadSerializer{}
	serializedBody, err := serializer.SerializePayload(body)
	if err != nil {
		return nil, fmt.Errorf("[Socket]: failed to serialize payload: %w", err)
	}

	requestBuilder := &network.RealHTTPRequestBuilder{}
	req, err := requestBuilder.NewRequest(method, url, serializedBody)
	if err != nil {
		return nil, fmt.Errorf("[Socket]: failed to create request: %w", err)
	}

	transportCreator := &network.RealTransportCreator{}
	client := &http.Client{
		Transport: transportCreator.NewUnixTransport(socketPath),
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[Socket]: request failed: %w", err)
	}

	defer resp.Body.Close()

	responseReader := &network.RealResponseReader{}
	bodyResp, err := responseReader.ReadBody(resp)
	if err != nil {
		return nil, fmt.Errorf("[Socket]: failed to read response body: %w", err)
	}

	responseValidator := &network.RealResponseValidator{}
	err = responseValidator.ValidateStatus(resp, bodyResp)
	if err != nil {
		return nil, fmt.Errorf("[Network]: invalid response: %w", err)
	}

	return bodyResp, nil
}
