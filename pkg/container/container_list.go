// go-podman/pkg/container/container_list.go
package container

import (
	"encoding/json"
	"fmt"

	"github.com/MineHosting/go-podman/pkg/engine"
)

// ListContainers retrieves all containers managed by the Podman engine.
//
// It sends a GET request to the `/libpod/containers/json` endpoint using the provided
// engine client. The response is expected to be a JSON array of containers, which will
// be unmarshaled into a slice of `Container` structs.
//
// Parameters:
//   - pd: An instance of engine.Client responsible for handling HTTP requests.
//
// Returns:
//   - []Container: A slice of containers returned by the Podman API.
//   - error: An error if the request fails or the response cannot be parsed.
//
// Example:
//
//	containers, err := ListContainers(client)
//	if err != nil {
//	    log.Fatalf("failed to list containers: %v", err)
//	}
//	for _, c := range containers {
//	    fmt.Printf("ID: %s, Names: %v\n", c.ID, c.Names)
//	}
func ListContainers(pd engine.Client) ([]Container, error) {
	data, err := pd.Send("GET", "/libpod/containers/json", nil)
	if err != nil {
		return nil, err
	}

	var containers []Container

	err = json.Unmarshal(data, &containers)
	if err != nil {
		return nil, fmt.Errorf("erro ao deserializar resposta: %v", err)
	}

	return containers, nil
}
