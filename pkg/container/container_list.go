// go-podman/pkg/container/container_list.go
package container

import (
	"encoding/json"
	"fmt"

	"github.com/MineHosting/go-podman/pkg/engine"
)

func List_Containers(pd engine.Client) ([]Container, error) {
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
