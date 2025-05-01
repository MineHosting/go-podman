// go-podman/internal/structures/network/payload_serializer_interface.go
package network

import "io"

type PayloadSerializer interface {
	SerializePayload(payload any) (io.Reader, error)
}
