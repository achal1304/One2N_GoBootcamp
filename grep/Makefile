.PHONY: build
build:
	go build -o grepgo

.PHONY: test
test:
	go test ./...

.PHONY: testcover
testcover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out