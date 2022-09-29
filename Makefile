.PHONY: build-mycontainer run run-docker integration-test install-dependency integration-test-report

build-mycontainer:
	docker build -t mycontainer:v0.0.0 .

run:
	go run main.go

run-docker:
	docker run mycontainer:v0.0.0

integration-test: build-mycontainer
	go test

install-dependency:
	go install github.com/vakenbolt/go-test-report@v0.9.3

integration-test-report: install-dependency
	go test -json | go-test-report -v