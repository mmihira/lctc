FROM golang:alpine as build
RUN apk update \
    && apk add --virtual build-dependencies \
        build-base \
        gcc \
        wget \
        git \
    && apk add \
        bash

RUN mkdir /src
COPY go.mod /src
COPY go.sum /src
COPY *.go /src/
COPY models/* /src/models/

WORKDIR /src
RUN go build -o main .

CMD ["./main"]
