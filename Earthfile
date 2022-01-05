VERSION 0.6
FROM golang:1.17
WORKDIR /rtmp-wav-server

RUN apt update && apt -y install golang-go ca-certificates openssl

fdkaac:
  FROM ghcr.io/izumisy/fdkaac:latest
  SAVE ARTIFACT /fdkaac-include
  SAVE ARTIFACT /fdkaac-lib

deps:
  COPY go.mod go.sum .
  RUN go mod download

build:
  FROM +deps
  COPY +fdkaac/fdkaac-include /usr/include/fdk-aac
  COPY +fdkaac/fdkaac-lib /usr/lib/fdk-aac
  COPY main.go .
  RUN go build -o build/rtmp-wav-server main.go
  SAVE ARTIFACT build/rtmp-wav-server

run:
  LOCALLY
  COPY +build/rtmp-wav-server .
  RUN ./rtmp-wav-server
