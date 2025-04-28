package network_test

import "errors"

type brokenReader struct{}

func (brokenReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("broken read")
}
