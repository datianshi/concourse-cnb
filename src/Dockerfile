FROM golang:1.16
WORKDIR /go/src/github.com/datianshi/concourse-cnb/config
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o buildpack cmd/config/*.go

FROM alpine:latest  
RUN apk --no-cache add --no-cache ca-certificates wget jq bash unzip
WORKDIR /root/
COPY --from=0 /go/src/github.com/datianshi/concourse-cnb/config/buildpack /usr/bin/
RUN chmod +x /usr/bin/buildpack