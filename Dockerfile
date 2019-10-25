FROM golang:1.11-alpine AS registry

# Flags
ARG REGISTRY_VERSION=v2.7.1
ARG GOOS=linux
ARG GOARCH=amd64
ARG GOARM=6

ENV BUILDTAGS include_oss include_gcs

# Dependencies
RUN set -ex && \
    apk add --no-cache make git file

WORKDIR /go/src/github.com/docker/distribution

# Clone and build
RUN cd .. && \
    mkdir -p /go/src/github.com/wolfulus/transfer/ && \
    git clone --branch ${REGISTRY_VERSION} --depth 1 https://github.com/docker/distribution.git && \
    cd distribution && \
    sed -i '/import (/a _ "github.com/docker/distribution/transfer/hooks"' ./cmd/registry/main.go && \
    cat ./cmd/registry/main.go

# Additional dependencies
RUN go get github.com/foomo/htpasswd && \
    go get github.com/gin-gonic/gin

# Copy our files
COPY . .
COPY . /go/src/github.com/wolfulus/transfer/

# Build
RUN CGO_ENABLED=0 make PREFIX=/go clean binaries && file ./bin/registry | grep "statically linked"

#
# Dist
#
# FROM alpine
#
# # Additional files
# COPY ./configs /etc/docker/registry
#
# # Registry executable
# COPY --from=registry /go/src/github.com/docker/distribution/bin/registry /bin/registry
#
# # Default port
# EXPOSE 5000
#
# # Entrypoint
# ENTRYPOINT ["registry"]
# CMD ["serve", "/etc/docker/registry/config.yml"]
#
