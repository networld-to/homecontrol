# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.224.3/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.19-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# see https://github.com/protocolbuffers/protobuf/releases
ARG PROTOC_VERSION=21.11
# see https://pkg.go.dev/google.golang.org/protobuf
ARG PROTOBUF_GO_VERSION=v1.28.1
# see https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc
ARG PROTOBUF_GO_GRPC_VERSION=v1.2.0

ENV PROTOBUF_URL https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip
RUN curl -L -o /tmp/protobuf.zip ${PROTOBUF_URL}
WORKDIR /tmp/
RUN unzip protobuf.zip && mv bin/protoc /usr/local/bin
RUN chmod 777 /usr/local/bin/*

RUN sudo ln -sf /usr/share/zoneinfo/America/New_York /etc/localtime

ENV GO111MODULE=on
WORKDIR /workspaces

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOBUF_GO_VERSION} && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOBUF_GO_GRPC_VERSION}

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

RUN chown vscode:vscode -R /go/
