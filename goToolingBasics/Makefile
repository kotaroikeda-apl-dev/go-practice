.PHONY: tidy test run install clean-modcache

tidy:
	go mod tidy

test:
	go test ./...

run:
	go run ./cmd/goToolingBasics

install:
	go install ./cmd/goToolingBasics

clean-modcache:
	go clean -modcache


