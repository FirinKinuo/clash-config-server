FROM golang:1.25-alpine as builder

RUN apk add --no-cache \
    make\
    git

WORKDIR /build

COPY . .

RUN make build-docker

FROM alpine:3.22

RUN apk add --no-cache \
    tzdata

LABEL org.opencontainers.image.title="clash-config-server" \
      org.opencontainers.image.description="Configuration server for Clash" \
      org.opencontainers.image.authors="Firin Kinuo <me@fkinuo.dev>" \
      org.opencontainers.image.url="https://github.com/FirinKinuo/clash-config-server" \
      org.opencontainers.image.source="https://github.com/FirinKinuo/clash-config-server" \
      org.opencontainers.image.licenses="MIT" \
      maintainer="Firin Kinuo <me@fkinuo.dev>"

ENV CCS_CONFIG_PATH=/data/config.yml \
    CCS_ADDR=0.0.0.0:8000 \
    TZ=UTC

WORKDIR /app
COPY --from=builder /build/clash-config-server ./

EXPOSE 8000

ENTRYPOINT ["./clash-config-server"]