// go-podman/pkg/engine/Start.go
package engine

import "github.com/MineHosting/go-podman/internal/engine"

// CreateNewClient creates and initializes a new Podman client, configuring whether rootless mode is enabled or not.
//
// ## Parameters:
// - isRootlessMode: A boolean indicating whether the client should be initialized in rootless mode (true) or not (false).
//
// ## Returns:
// - engine.Client: The initialized Podman client, implementing the Client interface.
// - error: Returns an error if client initialization or the ping test fails.
//
// ## Example usage:
// client, err := CreateNewClient(true)
//
//	if err != nil {
//	    log.Fatalf("Error creating client: %v", err)
//	}
//
// ## Notes:
// This function also performs a ping test on the client after initialization to ensure it is functional before being returned.
func CreateNewClient(isRootlessMode bool) (engine.Client, error) {
	client, err := engine.Start(isRootlessMode)
	if err != nil {
		return nil, err
	}

	err = engine.Ping(client)
	if err != nil {
		return nil, err
	}

	return client, err
}
