// go-podman/internal/engine/ping.go
package engine

import "fmt"

func Ping(pd Client) error {
	resp, err := pd.Send("GET", "libpod/_ping", nil)
	if err != nil {
		return fmt.Errorf("failed to ping Podman: %w", err)
	}

	if string(resp) != "OK" {
		return fmt.Errorf("unexpected ping response: %s", string(resp))
	}

	return nil
}
