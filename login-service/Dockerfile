# Build Angular app
FROM node:alpine AS client
RUN mkdir /tmp/wkdir
WORKDIR /tmp/wkdir
ADD ./client ./
RUN npm install -g @angular/cli && npm install
RUN ng build -- --prod --environment prod --progress=false --base-href /congregation-publishers/ --base-href=/static/

# Build Go App
FROM golang:alpine AS server
RUN apk update && apk add ca-certificates && apk add make git
RUN mkdir -p $GOPATH/src/github.com/srleyva/Windmill/authentication
WORKDIR $GOPATH/src/github.com/srleyva/Windmill/authentication
ADD ./ ./
RUN make && cp login /login

# Runtime Image
FROM scratch
WORKDIR /app
COPY --from=client /tmp/wkdir/dist/ /app/wwwroot
COPY --from=server /login /app
ENTRYPOINT [ "/app/login" ]

