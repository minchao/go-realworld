FROM golang:1.13-alpine3.11 AS builder

RUN apk add -U --no-cache \
        ca-certificates

WORKDIR /go/src/github.com/minchao/go-realworld
COPY go.mod go.sum ./
RUN go mod download

# build application
ARG BUILD_VERSION
ARG BUILD_DATE
ARG BUILD_COMMIT
ARG CMD_PACKAGE

COPY . .

RUN CGO_ENABLED=0 go build \
        -ldflags "-s -X ${CMD_PACKAGE}.Version=${BUILD_VERSION} -X ${CMD_PACKAGE}.Commit=${BUILD_COMMIT} -X ${CMD_PACKAGE}.Date=${BUILD_DATE}" \
        ./cmd/realworld

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/minchao/go-realworld/realworld /usr/bin/realworld

USER 1000

ENTRYPOINT ["realworld"]
CMD ["serve"]
