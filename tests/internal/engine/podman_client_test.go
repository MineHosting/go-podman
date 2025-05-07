// go-podman/tests/internal/engine/podman_client_test.go
package engine_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPodmanClient_Connect_Success(t *testing.T) {
	mockSocket := new(MockSocketClient)

	mockSocket.On("Connect").Return(nil)

	err := mockSocket.Connect()

	assert.NoError(t, err, "Expected Connect to execute without errors")
	mockSocket.AssertExpectations(t)
}

func TestPodmanClient_SendRequest(t *testing.T) {
	mockSocket := new(MockSocketClient)

	mockSocket.On("SendRequest", "testRequest").Return("response", nil)

	response, err := mockSocket.SendRequest("testRequest")

	assert.NoError(t, err, "Expected SendRequest to execute without errors")
	assert.Equal(t, "response", response, "The response did not match the expected result")
	mockSocket.AssertExpectations(t)
}
