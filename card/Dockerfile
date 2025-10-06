FROM golang:1.25 AS development

RUN apt-get update && apt-get install -y protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 && \
    go install github.com/pressly/goose/v3/cmd/goose@v3.24.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 50051
