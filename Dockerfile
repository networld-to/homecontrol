# Compiling application
FROM golang:1.15-alpine as builder

RUN apk add --no-cache git gcc musl-dev bash

ENV PROJ github.com/networld-to/homecontrol
ENV PROJ_DIR /go/src/${PROJ}

WORKDIR ${PROJ_DIR}

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -buildmode=pie -mod=readonly \
    -ldflags "-extldflags "-static" -linkmode external -s -w -X ${PROJ}/utils.Version=$(git describe --dirty --always) -X '${PROJ}/utils.Build=$(date -u '+%Y-%m-%dT%H:%M:%SZ')'" \
    -a -o /tmp/server ${PROJ}/server

RUN go build -trimpath -buildmode=pie -mod=readonly \
    -ldflags "-extldflags "-static" -linkmode external -s -w  -X ${PROJ}/utils.Version=$(git describe --dirty --always) -X '${PROJ}/utils.Build=$(date -u '+%Y-%m-%dT%H:%M:%SZ')'" \
    -a -o /tmp/client ${PROJ}/client

RUN mv /tmp/server /go/bin && \
    mv /tmp/client /go/bin

#####################################################
# Final, minimized docker image usable in production
#####################################################
FROM alpine:3.12

RUN apk add --no-cache ca-certificates

WORKDIR /
COPY --from=builder /go/bin/server /
COPY --from=builder /go/bin/client /

EXPOSE 50051

