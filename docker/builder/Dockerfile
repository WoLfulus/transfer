#
# Builder
#
FROM golang:1.13-alpine

ENV GO111MODULE=on

ARG TRANSFER_VERSION=latest
ENV TRANSFER_VERSION=${TRANSFER_VERSION}

# Prerequisites
WORKDIR /transfer
RUN apk add --no-cache bash curl git gcc make upx && \
    curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

# Dependencies
COPY go.* ./
RUN go mod download

# Build
COPY . .
RUN ./bin/goreleaser --rm-dist --skip-publish --skip-sign --skip-validate
