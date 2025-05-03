// go-podman/engine/podman_client.go
package engine

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/MineHosting/go-podman/internal/socket"
)

type PodmanClient struct {
	ApiVersion string

	SocketType   socket.SocketPath
	SocketClient socket.SocketClientInterface
}

func NewPodmanClient(st socket.SocketPath, sc socket.SocketClientInterface) *PodmanClient {
	return &PodmanClient{
		ApiVersion:   "v1.0.0",
		SocketType:   st,
		SocketClient: sc,
	}
}

func (pd *PodmanClient) ChangeApiVersion() error {
	const versionEndpoint = "libpod/version"

	resp, err := pd.Send("GET", versionEndpoint, nil)
	if err != nil {
		return fmt.Errorf("[Engine]: failed to fetch version info: %w", err)
	}

	var parsed struct {
		Components []struct {
			Name    string `json:"Name"`
			Details struct {
				APIVersion string `json:"APIVersion"`
			} `json:"Details"`
		} `json:"Components"`
	}

	if err := json.Unmarshal(resp, &parsed); err != nil {
		return fmt.Errorf("[Engine]: failed to parse version components: %w", err)
	}

	for _, comp := range parsed.Components {
		if comp.Name == "Podman Engine" {
			if comp.Details.APIVersion == "" {
				return fmt.Errorf("[Engine]: Podman Engine found but missing APIVersion")
			}
			pd.ApiVersion = fmt.Sprintf("v%s", comp.Details.APIVersion)
			return nil
		}
	}

	return fmt.Errorf("[Engine]: Podman Engine component not found in version response")
}

func (pd *PodmanClient) Send(method, endpoint string, body any) ([]byte, error) {
	type serializer interface {
		SerializePayload(payload any) (io.Reader, error)
	}

	var serializedBody io.Reader
	var err error
	if body != nil {
		if s, ok := pd.SocketClient.(serializer); ok {
			serializedBody, err = s.SerializePayload(body)
			if err != nil {
				return nil, err
			}
		}
	}

	url := fmt.Sprintf("/%s%s", pd.ApiVersion, endpoint)
	return pd.SocketClient.Send(method, url, serializedBody, pd.SocketType)
}
