# Introduction

Collection of different IoT services implemented in golang with the help of
gRPC and Protobuf.

The first implemented service allows the control of Philips Hue lights if
they are in the same network or if the Philips Hub API is accessible.

# Getting Started

    # Build server and client, see ./build/ directory
    make

    # Install server and client to ${GOPATH}/bin
    make install

    # Build and start docker image
    make docker

    # Clean generate protobuf and binaries
    make clean

    # Start local server
    go run server/main.go

    # Blink lights that are part of group 2 via homecontrol service hosted
    # locally.
    go run client/main.go -host 127.0.0.1:50051 -cmd blink -group 2

