.PHONY: test

test:
	@go test -v ./tests/...
	@go clean -testcache
