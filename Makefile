.PHONY: all build install tls clean hue docker-build docker-run

all: build tls install docker-build

build:
	protoc -I hue hue/hue.proto --go_out=plugins=grpc:hue
	./build.sh

install:
	make build
	cp build/homecontrol_server ${GOPATH}/bin
	cp build/homecontrol_client ${GOPATH}/bin

tls:
	mkdir -p ${HOME}/.homecontrol
	openssl genrsa -out ${HOME}/.homecontrol/server.key 4096
	openssl req -new -x509 -sha256 -key ${HOME}/.homecontrol/server.key \
		-out ${HOME}/.homecontrol/server.crt -days 3650 \
		-subj "/C=US/ST=MA/L=Cambridge/O=Networld/CN=Homecontrol/emailAddress=foo@bar.com"

clean:
	rm -rf */*.pb.go \
		build/homecontrol* ${HOME}/.homecontrol

docker-build:
	docker build -t networld/homecontrol .

docker-run:
	docker run -v ~/.philips-hue.json:/root/.philips-hue.json -p 50051:50051 -it --rm networld/homecontrol /go/bin/server

push:
	docker build -t networld/homecontrol .
	docker push networld/homecontrol
