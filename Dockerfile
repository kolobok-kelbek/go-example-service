ARG GOLANG_VERSION=1.23.4
ARG ALPINE_VERSION=3.17

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} as golang

ENV GO111MODULE=on
ENV GOROOT=/usr/local/go
ENV GOOS="linux"
ARG UID=1000
ARG GID=1000

RUN addgroup -S gos --gid $GID                              &&\
    adduser --uid $UID -S gos -G gos                        &&\
    apk add --no-cache git gcc musl-dev make bash           &&\
    # latest не есть хорошо, заменить на точную версию как будет возможность
    go install github.com/cosmtrek/air@latest               &&\
    go install github.com/go-delve/delve/cmd/dlv@latest     &&\
    # завязка на конкретную ветку не есть хорошо, заменить на точную версию как будет возможность
    go install github.com/google/wire/cmd/wire@main         &&\
    mkdir -p /go/src/app                                    &&\
    chown gos:gos -R /go

WORKDIR /go/src/app

USER gos

FROM golang AS air

EXPOSE 8080 40000

ENTRYPOINT ["air"]
