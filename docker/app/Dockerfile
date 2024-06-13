FROM golang:1.22-alpine

RUN mkdir /go/src/app

WORKDIR /go/src/app

# Install golangci-lint, goimports, and air
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.2 && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/air-verse/air@latest

# Copy air configuration file
COPY ./app/.air.toml ./
