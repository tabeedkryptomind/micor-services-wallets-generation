FROM golang:1.17.5-alpine3.15 as dev

RUN apk update && apk upgrade && \
    apk add --no-cache git \
        dpkg \
        gcc \
        git \
        musl-dev



WORKDIR /wallets

EXPOSE 8080