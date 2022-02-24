# Compiling application
FROM golang:1.17-alpine3.15 as builder

RUN apk add --no-cache git gcc musl-dev bash

ENV PROJ github.com/networld-to/homecontrol
ENV PROJ_DIR /go/src/${PROJ}

WORKDIR ${PROJ_DIR}

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -buildmode=pie -mod=readonly \
    -ldflags "-buildid= -extldflags "-static" -linkmode external -s -w -X ${PROJ}/utils.Version=$(git describe --dirty --always) -X '${PROJ}/utils.Build=REPRODUCIBLE'" \
    -a -o /tmp/server ${PROJ}/server

RUN go build -trimpath -buildmode=pie -mod=readonly \
    -ldflags "-buildid= -extldflags "-static" -linkmode external -s -w  -X ${PROJ}/utils.Version=$(git describe --dirty --always) -X '${PROJ}/utils.Build=REPRODUCIBLE'" \
    -a -o /tmp/client ${PROJ}/client

RUN mv /tmp/server /go/bin && \
    mv /tmp/client /go/bin

RUN sha256sum /go/bin/server /go/bin/client

#####################################################
# Final, minimized docker image usable in production
#####################################################
FROM scratch

COPY --from=alpine:3.15 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /
COPY --from=builder /go/bin/server /
COPY --from=builder /go/bin/client /

EXPOSE 50051

