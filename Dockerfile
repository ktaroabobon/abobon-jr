FROM golang:1.22-alpine

RUN mkdir /go/src/app

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@v1.52.0
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.2
RUN go install github.com/daixiang0/gci@latest
RUN go install mvdan.cc/gofumpt@latest
RUN go install golang.org/x/tools/cmd/goimports@latest