// Go-podman/tests/internal/engine/engine_helper_test.go
package engine_test

import "github.com/stretchr/testify/mock"

type MockPodmanClient struct {
	mock.Mock
}

func (m *MockPodmanClient) StartPodman() error {
	args := m.Called()
	return args.Error(0)
}

type MockSocketClient struct {
	mock.Mock
}

func (m *MockSocketClient) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSocketClient) SendRequest(request string) (string, error) {
	args := m.Called(request)
	return args.String(0), args.Error(1)
}
