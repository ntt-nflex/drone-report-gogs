.PHONY: go-build
go-build: 
	@echo "Build project binary..."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o report-gogs
