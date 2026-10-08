FROM alpine:latest

LABEL org.opencontainers.image.authors="ProjectDiscovery"
LABEL org.opencontainers.image.description="Go client to communicate with Chaos dataset API."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.title="chaos-client"
LABEL org.opencontainers.image.url="https://github.com/projectdiscovery/chaos-client"

RUN apk -U upgrade --no-cache \
    && apk add --no-cache bind-tools ca-certificates

ARG TARGETPLATFORM
COPY $TARGETPLATFORM/chaos-client /usr/local/bin/chaos

ENTRYPOINT ["chaos"]
