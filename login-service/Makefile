BINARY = login
VERSION=0.1.1
# COMMIT=$(shell git rev-parse HEAD)
# BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_USERNAME=srleyva
GIT_REPO=${GIT_USERNAME}/${BINARY}
DOCKER_REPO=sleyva97/$(BINARY)

LDFLAGS = -ldflags "-X main.VERSION=${VERSION}" #TODO -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"
target = cmd/server/server.go cmd/server/plugin.go
pkgs = $(shell go list ./... | grep -v /vendor/ | grep -v /test/)
gobuild = go build ${LDFLAGS} -o ${BINARY} ${target}



all: format dep build-linux

dep:
	dep ensure -v

format:
	go fmt $(pkgs)

test:
	go test `go list ./... | grep -v apis | grep -v client`

build:
	$(gobuild)

build-linux:
	CGO_ENABLED=0 GOOS=linux $(gobuild)

docker-build:
	docker build . -t ${DOCKER_REPO}:${VERSION}
