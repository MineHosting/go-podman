// go-podman/internal/socket/socket_types.go
package socket

import (
	"fmt"
	"os/user"
)

type SocketPath string

var (
	Rootfull SocketPath = "/run/podman/io.podman/podman.sock"
	Rootless SocketPath
)

func GetRootlessPath() error {
	CollectUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("[GetRootlessPath]: Can't get info about user: %w", err)
	}

	Rootless = SocketPath(fmt.Sprintf("/run/user/%s/podman/podman.sock", CollectUser.Uid))

	return nil
}
