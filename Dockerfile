ARG GOLANG_VERSION=1.23.5
ARG ALPINE_VERSION=3.21

# https://hub.docker.com/_/golang
FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} as golang

ENV GO111MODULE=on
ENV GOROOT=/usr/local/go
ENV GOOS="linux"
ARG UID=1000
ARG GID=1000

RUN addgroup -S tmt --gid $GID &&\
    adduser --uid $UID -S tmt -G tmt &&\
    apk add --no-cache git gcc musl-dev make bash &&\
    go install github.com/air-verse/air@v1.61.7 &&\
    go install github.com/go-delve/delve/cmd/dlv@v1.24.0 &&\
    go install github.com/pressly/goose/v3/cmd/goose@v3.24.1&&\
    mkdir -p /go/src/app &&\
    chown tmt:tmt -R /go

WORKDIR /go/src/app

USER tmt

FROM golang AS air

EXPOSE 8080 40000

ENTRYPOINT ["air"]
