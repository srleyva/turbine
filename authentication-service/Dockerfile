# Build Server
FROM golang:alpine as server
RUN apk update && apk add make git tree
RUN mkdir -p $GOPATH/src/github.com/srleyva/turbine/authentication-service
WORKDIR $GOPATH/src/github.com/srleyva/turbine/authentication-service
ADD ./ ./
RUN make && cp authentication-service /authentication-service

# Inject Binary into container
FROM scratch
COPY --from=server /authentication-service /bin/authentication-service
ENTRYPOINT ["authentication-service"]
