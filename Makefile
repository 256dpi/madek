PKGS=$(shell glide nv)

all: fmt vet lint

vet:
	go vet $(PKGS)

fmt:
	go fmt $(PKGS)

lint:
	golint .
	golint cmd/madek

err:
	errcheck -ignoretests -asserts $(PKGS)
