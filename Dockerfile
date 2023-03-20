# Base
FROM golang:1.20.1-alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build ./cmd/chaos

# Release
FROM alpine:latest
RUN apk add --no-cache bind-tools ca-certificates
COPY --from=build-env /app/chaos /usr/local/bin/
ENTRYPOINT ["chaos"]