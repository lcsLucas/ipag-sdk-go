FROM golang:1.20.3-alpine

ARG user=app
ARG uid=1000
ARG dir=ipag-sdk-go

WORKDIR /go/src/github.com/lcslucas/${dir}

RUN apk add --no-cache \
    bash

RUN adduser -D -u 1000 -s /bin/bash ${user} && \
    mkdir -p /var/log/${dir} && \
    chown -R ${user} . && \
    chown -R ${user} /var/log/${dir} && \
    find /var/log/${dir} -type f | xargs -I{} chmod -v 644 {} && \
    find /var/log/${dir} -type d | xargs -I{} chmod -v 755 {} && \
    find . -type f | xargs -I{} chmod -v 644 {} && \
    find . -type d | xargs -I{} chmod -v 755 {};

USER ${user}

ENV HOME /home/${user}