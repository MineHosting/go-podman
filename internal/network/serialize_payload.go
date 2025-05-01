// go-podman/internal/network/serialize_payload.go
package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type RealPayloadSerializer struct{}

func (r *RealPayloadSerializer) SerializePayload(payload any) (io.Reader, error) {
	if payload == nil {
		return nil, nil
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("[Network]: failed in marshalling payload: %v", err)
	}
	return bytes.NewReader(raw), nil
}
