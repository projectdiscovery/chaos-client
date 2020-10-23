FROM golang:1.14-alpine AS builder
COPY . /app
WORKDIR /app
RUN go get ./cmd/chaos
RUN go build -o chaos ./cmd/chaos

FROM alpine
RUN adduser --home /app --shell /bin/sh --disabled-password appuser
COPY --from=builder --chown=appuser:appuser /app/chaos /app
USER appuser

WORKDIR /app
ENTRYPOINT ["/app/chaos"]
CMD ["-h"]
