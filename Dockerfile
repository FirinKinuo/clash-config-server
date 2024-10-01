FROM golang:1.22-alpine as builder

RUN apk add --no-cache make git

WORKDIR /build

COPY . .
RUN make build-docker

FROM alpine:3.20

WORKDIR /
COPY --from=builder /build/clash-config-server ./

EXPOSE 8000

ENTRYPOINT ["./clash-config-server", "-c", "/data/config.yml"]