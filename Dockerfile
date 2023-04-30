# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

RUN apt-get update && apt-get install -y protobuf-compiler
COPY ./ ./
RUN make generate
RUN CGO_ENABLED=0 GOOS=linux go build -o . ./cmd/app

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage ./app/app app

USER nonroot:nonroot

# Wait for db
COPY --from=ghcr.io/ufoscout/docker-compose-wait:latest /wait /wait

ENV WAIT_COMMAND="./app"

ENTRYPOINT ["/wait"]