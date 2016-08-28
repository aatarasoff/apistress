FROM golang:1.7-alpine
MAINTAINER aatarasoff@gmail.com

VOLUME /data
WORKDIR /data

RUN apk update && \
    apk upgrade && \
    apk add git bash

RUN go get github.com/aatarasoff/apistress && \
    go install github.com/aatarasoff/apistress

CMD [ "apistress" ]