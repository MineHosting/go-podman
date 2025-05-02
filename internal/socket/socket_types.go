// go-podman/internal/socket/socket_types.go
package socket

type SocketPath string

const (
	Rootfull SocketPath = "/run/podman/io.podman/podman.sock"
	Rootless SocketPath = "/run/user/1000/podman/podman.sock"
)
