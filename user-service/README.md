# User-Service Service

This is the User-Service service

Generated with

```
micro new github.com/srleyva/turbine/user-service --namespace=go.micro --type=srv --plugin=registry=kubernetes
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.user-service
- Type: srv
- Alias: user-service

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-service-srv
```

Build a docker image
```
make docker
```