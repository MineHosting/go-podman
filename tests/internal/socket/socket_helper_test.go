// go-podman/tests/internal/socket/socket_helper_test.go
package socket_test

import (
	"io"
	"net/http"
)

type MockRequestBuilder struct {
	Err error
}

func (m *MockRequestBuilder) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	req, _ := http.NewRequest(method, url, body)
	return req, nil
}

type MockResponseReader struct {
	Body []byte
	Err  error
}

func (m *MockResponseReader) ReadBody(resp *http.Response) ([]byte, error) {
	return m.Body, m.Err
}

type MockResponseValidator struct {
	Err error
}

func (m *MockResponseValidator) ValidateStatus(resp *http.Response, body []byte) error {
	return m.Err
}

type MockTransport struct {
	Response *http.Response
	Err      error
}

func (m *MockTransport) NewUnixTransport(socketPath string) http.RoundTripper {
	return &mockRoundTripper{
		resp: m.Response,
		err:  m.Err,
	}
}

type mockRoundTripper struct {
	resp *http.Response
	err  error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.resp, m.err
}
