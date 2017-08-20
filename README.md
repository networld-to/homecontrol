# Introduction

Collection of different IoT services implemented in golang with the help of
gRPC and Protobuf.

The first implemented service allows the control of Philips Hue lights if
they are in the same network or if the Philips Hub API is accessible.

# Getting Started

    # Build server and client.
    # Generate TLS certificates
    # Install server and client to ${GOPATH}/bin
    # Build docker image with server and client included
    make all

    # Install server and client to ${GOPATH}/bin
    make install

    # Build and start docker image
    make docker-build docker-run

    # Clean generate protobuf, binaries and TLS certificates
    make clean

    # Start local server with TLS
    go run server/main.go -tls

    # Blink lights that are part of group 2 via homecontrol service hosted
    # locally.
    go run client/main.go -host 127.0.0.1:50051 -tls -cmd blink -group 2

    # Dockerized client example with service hosted under 192.168.1.2
    docker run -it --rm networld/homecontrol /go/bin/client -host 192.168.1.2:50051 -cmd blink

# Features

* TLS encrypted connections
* Lights: Philips Hue support
