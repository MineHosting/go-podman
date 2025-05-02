// go-podman/internal/socket/socket_client.go
package socket

import (
	"fmt"
	"io"
	"net/http"

	"github.com/MineHosting/go-podman/internal/structures/network"
)

var _ SocketClientInterface = (*SocketClient)(nil)

type SocketClient struct {
	Serializer        network.PayloadSerializer
	RequestBuilder    network.HTTPRequestBuilder
	ResponseReader    network.ResponseReader
	ResponseValidator network.ResponseValidator
	NetworkTransport  network.TransportCreator
}

func NewSocketClient(Serializer network.PayloadSerializer, Request network.HTTPRequestBuilder, Response network.ResponseReader, Validator network.ResponseValidator, Transport network.TransportCreator) *SocketClient {
	return &SocketClient{
		Serializer:        Serializer,
		RequestBuilder:    Request,
		ResponseReader:    Response,
		ResponseValidator: Validator,
		NetworkTransport:  Transport,
	}
}

func (SC *SocketClient) Send(method, url string, body io.Reader, socket SocketPath) ([]byte, error) {
	injectedUrl := fmt.Sprintf("http://d/%s", url)
	req, err := SC.RequestBuilder.NewRequest(method, injectedUrl, body)
	if err != nil {
		return nil, fmt.Errorf("[SocketClient]: failed to build request: %w", err)
	}

	client := &http.Client{
		Transport: SC.NetworkTransport.NewUnixTransport(string(socket)),
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[SocketClient]: request failed: %w", err)
	}

	defer resp.Body.Close()

	bodyResp, err := SC.ResponseReader.ReadBody(resp)
	if err != nil {
		return nil, fmt.Errorf("[SocketClient]: failed to read response body: %w", err)
	}

	if err := SC.ResponseValidator.ValidateStatus(resp, bodyResp); err != nil {
		return nil, fmt.Errorf("[SocketClient]: invalid response status: %w", err)
	}

	return bodyResp, nil
}
