.PHONY: all build install tls clean hue docker-build docker-run

all: build install

build:
	protoc -I hue hue/hue.proto --go_out=plugins=grpc:hue
	protoc -I version version/version.proto --go_out=plugins=grpc:version
	./build.sh

install: build
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
		build/homecontrol*
	@echo "Keeping ${HOME}/.homecontrol"

docker-build:
	docker build -t networld/homecontrol .

docker-run:
	docker run -v ~/.philips-hue.json:/root/.philips-hue.json \
		-v ~/.homecontrol:/root/.homecontrol -p 50051:50051 \
		-it --rm networld/homecontrol /server -tls -endpoint=":50051"

push:
	docker build -t $(USER)/homecontrol .
	docker push $(USER)/homecontrol
