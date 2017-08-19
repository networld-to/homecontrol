.PHONY: build install clean hue docker

build: hue/hue.proto
	protoc -I hue hue/hue.proto --go_out=plugins=grpc:hue
	./build.sh

install:
	make build
	cp build/homecontrol_server ${GOPATH}/bin
	cp build/homecontrol_client ${GOPATH}/bin

clean:
	@rm -f */*.pb.go \
		build/homecontrol*

docker:
	@glide install
	@docker build -t homecontrol .
	@docker run -v ~/.philips-hue.json:/root/.philips-hue.json -p 50051:50051 -it --rm homecontrol /go/bin/server