.PHONY: build fmt run test collect-cover-data test-cover-html test-cover-func
export GO15VENDOREXPERIMENT=1

default: build
build: fmt
	go build -v -o omega-haproxyctl ./
fmt: 
	go fmt ./...

run: build
	./omega-haproxyctl

test:
	go test -v `go list ./... | grep -v /vendor/`

PACKAGES = $(shell go list ./... | grep -v /vendor/)
collect-cover-data:
	echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES),\
		go test -v -coverprofile=coverage.out -covermode=count $(pkg);\
		if [ -f coverage.out ]; then\
			tail -n +2 coverage.out >> coverage-all.out;\
		fi;)

test-cover-html:
	go tool cover -html=coverage-all.out -o coverage.html

test-cover-func:
	go tool cover -func=coverage-all.out
