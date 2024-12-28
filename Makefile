.PHONY:	testwc
testwc:
	go test ./wordcount/...

.PHONY: testwccover
testwccover:
	go test -coverprofile=coverage.out ./wordcount/...
	go tool cover -html=coverage.out