FROM alpine:latest
COPY iacker /usr/local/bin/iacker
ENTRYPOINT ["/usr/local/bin/iacker"]
