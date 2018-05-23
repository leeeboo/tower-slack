APP?=tower-slack
PORT?=8080
VERSION?=$(shell git describe --tags --always)
DOCKER_USER?=leeeboo
CONTAINER_IMAGE?=${DOCKER_USER}/${APP}

GOOS?=linux
GOARCH?=amd64

GO_VERSION?=1.10.2

default: build

.PHONY: clean
clean:
	rm -fr ${APP}

.PHONY: build-binary
build-binary: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${APP}

.PHONY: build
build:
	docker run --rm -v "$$PWD":/go/src/github.com/leeeboo/${APP} -w /go/src/github.com/leeeboo/${APP} golang:${GO_VERSION} make build-binary

.PHONY: docker-image
docker-image:
	docker build -t $(CONTAINER_IMAGE):$(VERSION) -f ./Dockerfile .

.PHONY: push
push:
	docker push $(CONTAINER_IMAGE):$(VERSION)

.PHONY: test
test:
	rm -f ./.testCoverage.txt
	go test $(shell go list ./... | grep -v /vendor/) -coverprofile ./.testCoverage.txt
	go tool cover -func=./.testCoverage.txt | grep 'total:' | grep 'statements' | awk '{print $$3}' | xargs echo "FullCoverage:"
