FROM golang:1-alpine AS builder

RUN apk update && apk add --no-cache git

RUN go get -v github.com/fullstorydev/grpcui/... && \
    go install -v github.com/fullstorydev/grpcui/cmd/grpcui

# Final image
FROM golang:1-alpine
COPY --from=builder /go/bin/* /

RUN apk update && apk add --no-cache bash

ARG GRPCUI_PORT
ENV GRPCUI_PORT ${GRPCUI_PORT:-8080}

ARG GRPCUI_SERVICE
ENV GRPCUI_SERVICE ${GRPCUI_SERVICE}

ENTRYPOINT /grpcui -bind 0.0.0.0 -plaintext -port ${GRPCUI_PORT} ${GRPCUI_SERVICE}
