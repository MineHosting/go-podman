// Go-podman/tests/internal/engine/podman_start_test.go
package engine_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartPodman_Success(t *testing.T) {
	mockClient := new(MockPodmanClient)

	mockClient.On("StartPodman").Return(nil)

	err := mockClient.StartPodman()

	assert.NoError(t, err, "Expected StartPodman to execute without errors")
	mockClient.AssertExpectations(t)
}

func TestStartPodman_Failure(t *testing.T) {
	mockClient := new(MockPodmanClient)

	mockClient.On("StartPodman").Return(fmt.Errorf("failed to start"))

	err := mockClient.StartPodman()

	assert.Error(t, err, "Expected StartPodman to return an error")
	assert.EqualError(t, err, "failed to start", "The error did not match the expected message")
	mockClient.AssertExpectations(t)
}
