package main

import (
	"flag"
	"net"
	"os"
	"os/user"

	hue "github.com/networld-to/homecontrol/api/generated/hue"
	version "github.com/networld-to/homecontrol/api/generated/version"
	"github.com/networld-to/homecontrol/utils"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var tls = flag.Bool("tls", false, "Activate TLS communication channel encryption.")
var endpoint = flag.String("endpoint", ":50051", "Host and port where the service is listening.")

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func getServerOptions(tls bool) []grpc.ServerOption {
	usr, err := user.Current()
	var prefix string
	if err == nil {
		prefix = usr.HomeDir
	}
	opts := []grpc.ServerOption{}
	if tls {
		cred, err := credentials.NewServerTLSFromFile(
			prefix+"/.homecontrol/server.crt",
			prefix+"/.homecontrol/server.key",
		)
		if err != nil {
			log.Fatalf("failed TLS: %v", err)
		}
		log.WithField("tls", true).Info("TLS Enabled")
		opts = append(opts, grpc.Creds(cred))
	} else {
		log.Warn("insecure connection. no TLS")
	}
	return opts
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := getServerOptions(*tls)
	s := grpc.NewServer(opts...)
	LoadHueBridgeConfig()
	hue.RegisterLightsServer(s, HueServer{})
	version.RegisterVersionServer(s, VersionServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.WithField("endpoint", *endpoint).WithField("type", "grpc").
		WithField("version", utils.Version).WithField("build", utils.Build).Infof("Server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
