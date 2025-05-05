FROM alpine:latest

LABEL org.opencontainers.image.source = "https://github.com/melkor217/bitmagnet"
RUN ["apk", "--no-cache", "add", "ca-certificates","curl","iproute2-ss"]

COPY bitmagnet /usr/local/bin/bitmagnet
ENTRYPOINT ["/usr/local/bin/bitmagnet"]
