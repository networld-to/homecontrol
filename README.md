# Introduction

![Go](https://github.com/networld-to/homecontrol/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/networld-to/homecontrol)](https://goreportcard.com/report/github.com/networld-to/homecontrol)

Collection of different IoT services implemented in golang with the help of
gRPC and Protobuf.

The first implemented service allows the control of Philips Hue lights if
they are in the same network or if the Philips Hub API is accessible.

# Getting Started

    # Build server and client.
    # Install server and client to ${GOPATH}/bin
    make all

    # Install server and client to ${GOPATH}/bin
    make install

    # Build and start docker image
    make docker-build docker-run

    # Clean generate protobuf and binaries
    make clean

    # Start local server with TLS
    go run server/main.go -tls

    # Blink lights that are part of group 2 via homecontrol service hosted
    # locally.
    go run client/main.go -host 127.0.0.1:50051 -tls -cmd blink -group 2

    # Dockerized client example with service hosted under 192.168.1.2
    docker run -it --rm networld/homecontrol /client -host 192.168.1.2:50051 -cmd blink

    # Dockerized client example with service hosted under 192.168.1.2. TLS support
    make tls
    docker run -it --rm --name homecontrol -v ~/.homecontrol:/root/.homecontrol networld/homecontrol /client -host 192.168.1.2:50051 -tls -cmd blink
    docker exec -it homecontrol /client -tls

# Examples: Client

    # Change the lights in group 3 to GREEN
    $ ./client/client -group 3 -cmd on -brightness 1 -sat 1 -hue 25500

    # Change the lights in group 3 to RED
    $ ./client/client -group 3 -cmd on -brightness 1 -sat 1 -hue 65535

# Features

* TLS encrypted connections
* Lights: Philips Hue support

# Appendix A: Directory Structure

    ├─ api
    │  ├─ generated                 # Generated Code from protobuf files
    │  │  ├─ hue                        # Service 1 package
    │  │  │  └─ hue.pb.go
    │  │  └─ version                    # Service 2 package
    │  │     └─ version.pb.go
    │  └─ proto                     # All the protobuf service definitions
    │     ├─ hue.proto
    │     └─ version.proto
    ├─ client                       # Client implementation, using api/generated code
    │  ├─ client
    │  └─ main.go
    ├─ go.mod
    ├─ go.sum
    ├─ server                       # Server implementation, using api/generated code
    │  ├─ hue.go
    │  ├─ main.go
    │  ├─ server
    │  └─ version.go
    └─ utils                        # Helper code used by client and server
        └─ version.go
