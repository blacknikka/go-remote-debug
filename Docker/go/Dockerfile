FROM golang:1.13.4-buster

WORKDIR /go

RUN go get -u github.com/go-redis/redis && \
    go get -u github.com/go-delve/delve/cmd/dlv

ENV PATH $PATH:/go/bin

RUN which dlv

# CMD ["which", "dlv"]
# CMD ["go", "run", "./app/main.go"]
CMD ["dlv","debug", "./app", "--listen=:2345", "--headless=true", "--api-version=2"]]
