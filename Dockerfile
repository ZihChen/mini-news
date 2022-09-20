FROM golang:1.16.7-alpine

RUN apk add git \
    && go get github.com/pilu/fresh \