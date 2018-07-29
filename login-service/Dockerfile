# Build Angular app
FROM node:alpine AS client
RUN mkdir /tmp/wkdir
WORKDIR /tmp/wkdir
ADD ./client ./
RUN npm install -g @angular/cli && npm install
RUN ng build --prod --build-optimizer=false --base-href=/static/

# Build Go App
FROM golang:alpine AS server
RUN apk update && apk add make git
RUN mkdir -p $GOPATH/src/github.com/srleyva/turbine/login-service
WORKDIR $GOPATH/src/github.com/srleyva/turbine/login-service
ADD ./ ./
RUN make && cp login /login

# Runtime Image
FROM scratch
WORKDIR /app
COPY --from=client /tmp/wkdir/dist/ /app/wwwroot
COPY --from=server /login /app
ENTRYPOINT [ "/app/login", "--app_dir", "/app/wwwroot" ]

