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
	go build -v -o server/server server/main.go

build-client:   ## Build the client
	go build -v -o client/client client/main.go

protoc:						## Generates protobuf code
	protoc -I hue hue/hue.proto --go_out=plugins=grpc:hue
	protoc -I version version/version.proto --go_out=plugins=grpc:version

build: protoc build-server build-client # Generates the protobuf code and buils the server and client

build-only: build-server build-client 	# Compiles the server and the client without generating new version of the protobuf code

install: build		## Builds and installs a new version
	cp build/homecontrol_server ${GOPATH}/bin
	cp build/homecontrol_client ${GOPATH}/bin

run: build      ## Executes first make build and then ./server/server
	./server/server -endpoint="127.0.0.1:50051"

run-client: 	## Executes ./client/client with ./config.yaml (needed for TLS)
	@./client/client

tls:              ## Generates TLS certificates for the server, under ~/.homecontrol
	mkdir -p ${HOME}/.homecontrol
	openssl genrsa -out ${HOME}/.homecontrol/server.key 4096
	openssl req -new -x509 -sha256 -key ${HOME}/.homecontrol/server.key \
		-out ${HOME}/.homecontrol/server.crt -days 3650 \
		-subj "/C=US/ST=MA/L=Cambridge/O=Networld/CN=Homecontrol/emailAddress=foo@bar.com"

clean:          ## Removes generated protobuffer code and binaries. Keeps ~/.homecontrol
	rm -f */*.pb.go
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

