SHELL = /bin/bash

.PHONY: all build install tls clean hue docker-build docker-run

help: ## Shows this help
	@IFS=$$'\n' ; \
		help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
		printf "%-30s %s\n" "command" "description" ; \
		printf "%-30s %s\n" "-------" "-----------" ; \
		for help_line in $${help_lines[@]}; do \
			IFS=$$':' ; \
			help_split=($$help_line) ; \
			help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
			printf '\033[36m'; \
			printf "%-30s %s" $$help_command ; \
			printf '\033[0m'; \
			printf "%s\n" $$help_info; \
		done


build-server:		## Build the server
	go build -trimpath -buildmode=pie -mod=readonly -ldflags "-buildid= -s -w \
			-X github.com/networld-to/homecontrol/utils.Version=$(shell git describe --dirty --always) \
			-X 'github.com/networld-to/homecontrol/utils.Build=REPRODUCIBLE'" \
		-a -o ./server/server github.com/networld-to/homecontrol/server

build-client:   ## Build the client
	go build -trimpath -buildmode=pie -mod=readonly -ldflags "-buildid= -s -w \
			-X github.com/networld-to/homecontrol/utils.Version=$(shell git describe --dirty --always) \
			-X 'github.com/networld-to/homecontrol/utils.Build=REPRODUCIBLE'" \
		-a -o ./client/client github.com/networld-to/homecontrol/client

protoc:						## Generates protobuf code
	protoc -I api/proto hue.proto --go_out=plugins=grpc:api/generated/hue
	protoc -I api/proto version.proto --go_out=plugins=grpc:api/generated/version

build: protoc build-server build-client # Generates the protobuf code and buils the server and client

build-only: build-server build-client 	# Compiles the server and the client without generating new version of the protobuf code

install: build		## Builds and installs a new version
	cp server/server ${GOPATH}/bin/homecontrol_server
	cp client/client ${GOPATH}/bin/homecontrol_client

run: build      ## Executes first make build and then ./server/server
	./server/server -endpoint=":50051"

run-client: 	## Executes ./client/client with ./config.yaml (needed for TLS)
	@./client/client

tls:              ## Generates TLS certificates for the server, under ~/.homecontrol
	mkdir -p ${HOME}/.homecontrol
	openssl genrsa -out ${HOME}/.homecontrol/server.key 4096
	openssl req -new -x509 -sha256 -key ${HOME}/.homecontrol/server.key \
		-out ${HOME}/.homecontrol/server.crt -days 3650 \
		-subj "/C=US/ST=MA/L=Cambridge/O=Networld/CN=Homecontrol/emailAddress=foo@bar.com" \
		-addext "subjectAltName=IP:0.0.0.0,IP:127.0.0.1,IP:192.168.1.2"

clean:          ## Removes generated protobuffer code and binaries. Keeps ~/.homecontrol
	@rm -f api/generated/*/*.pb.go
	rm -f server/server client/client
	@echo "Keeping ${HOME}/.homecontrol"

docker-build:    ## Builds a docker image with the client and server inside
	docker build --compress --rm -t $(USER)/homecontrol .

docker-run:      ## Runs the docker image and mounts all necessary config files from the host
	docker run --name homecontrol-inst \
		-v ~/.philips-hue.json:/root/.philips-hue.json \
		-v ~/.homecontrol:/root/.homecontrol -p 50051:50051 \
		-it --rm $(USER)/homecontrol /server -tls -endpoint=":50051"

push: docker-build ## Builds a docker image and pushes it to hub.docker.com/$(USER)/homecontrol
	docker push $(USER)/homecontrol

