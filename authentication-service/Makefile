BINARY = authentication-service
VERSION=0.1.0
GITHUB_USERNAME=srleyva
GITHUB_REPO=${GITHUB_USERNAME}/${BINARY}
DOCKER_REPO=sleyva97/$(BINARY)

LDFLAGS = -ldflags "-X main.VERSION=${VERSION}"
pkgs = $(shell go list ./... | grep -v /vendor/ | grep -v /test/)
gobuild = go build ${LDFLAGS} -o ${BINARY}

plugin = cmd/server/plugin.go
server = cmd/server/server.go


all: format build-linux

format:
	go fmt $(pkgs)

test:
	go test `go list ./... | grep -v apis | grep -v client`

gen-proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/authentication/authentication.proto

build:
	$(gobuild) $(server) $(plugin)

build-linux:
	CGO_ENABLED=0 GOOS=linux $(gobuild) $(server) $(plugin)

docker-build:
	docker build . -t ${DOCKER_REPO}:${VERSION}

# For testing purposes
consul-start:
	mkdir /tmp/consul
	consul agent -server -bootstrap-expect=1 -data-dir=/tmp/consul -advertise=127.0.0.1 > /tmp/consul/consul.log 2>&1 &
