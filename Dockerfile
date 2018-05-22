FROM alpine:3.7

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

COPY tower-slack /

ENV PORT 8080

ENTRYPOINT ["/tower-slack"]
