// go-podman/internal/network/validate_status.go
package network

import (
	"fmt"
	"net/http"
)

func ValidateStatus(resp *http.Response, body []byte) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("[Network]: unexpected status code %d – response body: %s", resp.StatusCode, string(body))
	}
	return nil
}
