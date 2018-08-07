# Build Server
FROM golang:alpine as server
RUN apk update && apk add make git tree curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN mkdir -p $GOPATH/src/github.com/srleyva/turbine/user-service
WORKDIR $GOPATH/src/github.com/srleyva/turbine/user-service
ADD ./ ./
RUN make && cp user-service /user-service

# Inject Binary into container
FROM scratch
COPY --from=server /user-service /bin/user-service
ENTRYPOINT ["user-service"]
