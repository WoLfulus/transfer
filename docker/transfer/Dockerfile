ARG TRANSFER_VERSION=latest

#
# Builder
#
FROM wolfulus/transfer:${TRANSFER_VERSION}-builder AS builder

#
# Dist
#
FROM alpine

# Transfer executable
COPY --from=builder /transfer/dist/transfer_linux_amd64/docker-transfer /bin/transfer

# Entrypoint
ENTRYPOINT ["transfer"]
CMD []
