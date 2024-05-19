FROM golang:1.22-alpine

RUN mkdir /go/src/app

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@v1.52.0