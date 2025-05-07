// go-podman/internal/socket/socket_types.go
package socket

import (
	"fmt"
	"os"
	"os/user"
)

type SocketPath string

var (
	Rootfull SocketPath = "/run/podman/io.podman/podman.sock"
	Rootless SocketPath = GetRootlessPath()
)

func GetRootlessPath() SocketPath {
	CollectUser, err := user.Current()
	if err != nil {
		fmt.Printf("[GetRootlessPath]: Can't get info about user: %v", err)
		os.Exit(1)
	}

	return SocketPath(fmt.Sprintf("/run/user/%s/podman/podman.sock", CollectUser.Uid))
}
