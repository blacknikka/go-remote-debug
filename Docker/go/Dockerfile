FROM golang:1.13.4-buster

WORKDIR /go/app
ENV GOPATH="/go"

COPY ../../src .

RUN set -ex && \
    apk update && \
    apk add --no-cache git && \
    go build -o src && \
    go get gopkg.in/urfave/cli.v2@master && \
    go get github.com/oxequa/realize && \
    go get -u github.com/go-delve/delve/cmd/dlv && \
    go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

FROM alpine:3.10

WORKDIR /app

COPY --from=build ../../src .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/src

CMD ["./src"]
