all: fmt vet lint test

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golint ./...

test:
	go test ./...

install:
	go install ./cmd/madek
