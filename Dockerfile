FROM golang:1.19.1-alpine as build-env
RUN go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest

FROM alpine:latest
RUN apk add --no-cache bind-tools ca-certificates
COPY --from=build-env /go/bin/chaos /usr/local/bin/chaos
ENTRYPOINT ["chaos"]