FROM golang:1.17.0-alpine as build-env
RUN GO111MODULE=on go get -v github.com/projectdiscovery/chaos-client/cmd/chaos

FROM alpine:latest
RUN apk add --no-cache bind-tools ca-certificates
COPY --from=build-env /go/bin/chaos /usr/local/bin/chaos
ENTRYPOINT ["chaos"]