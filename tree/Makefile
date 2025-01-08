.PHONY: build
build:
	go build -o treego

.PHONY: test
test:
	go test ./...

.PHONY: testcover
testcover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out