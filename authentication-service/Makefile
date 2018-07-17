BINARY = user-service
VERSION=0.1.0
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GITHUB_USERNAME=srleyva
GITHUB_REPO=${GITHUB_USERNAME}/${BINARY}

LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"
pkgs = $(shell go list ./... | grep -v /vendor/ | grep -v /test/)
gobuild = go build ${LDFLAGS} -o ${BINARY}

server = cmd/server/server.go


all: format build-linux

format:
	go fmt $(pkgs)

test:
	go test `go list ./... | grep -v apis | grep -v client`

gen-proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/authentication/authentication.proto

build:
	$(gobuild) $(server)

build-linux:
	CGO_ENABLED=0 GOOS=linux $(gobuild) $(server)

docker-build:
	docker build . -t ${GITHUB_REPO}:${COMMIT}

# For testing purposes
consul-start:
	mkdir /tmp/consul
	consul agent -server -bootstrap-expect=1 -data-dir=/tmp/consul -advertise=127.0.0.1 > /tmp/consul/consul.log 2>&1 &
