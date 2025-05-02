// Go-podman/tests/internal/engine/engine_helper_test.go
package engine_test

import (
	"io"
	"net/http"
)

type MockSerializer struct {
	Serialized io.Reader
	Err        error
}

func (m *MockSerializer) SerializePayload(payload any) (io.Reader, error) {
	return m.Serialized, m.Err
}

type MockRequestBuilder struct {
	Request *http.Request
	Err     error
}

func (m *MockRequestBuilder) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return m.Request, m.Err
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

func (m *MockTransport) NewUnixTransport(_ string) http.RoundTripper {
	return roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		return m.Response, m.Err
	})
}

type roundTripperFunc func(req *http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
