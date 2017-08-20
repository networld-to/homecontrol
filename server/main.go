package main

import (
	"flag"
	"net"
	"os"
	"os/user"

	"github.com/networld-to/homecontrol/hue"
	"github.com/networld-to/homecontrol/version"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var tls = flag.Bool("tls", false, "Activate TLS communication channel encryption.")
var endpoint = flag.String("endpoint", ":50051", "Host and port where the service is listening.s")

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func getServerOptions(tls bool) []grpc.ServerOption {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.ServerOption{}
	if tls {
		cred, err := credentials.NewServerTLSFromFile(
			usr.HomeDir+"/.homecontrol/server.crt",
			usr.HomeDir+"/.homecontrol/server.key",
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
	hue.LoadHueBridgeConfig()
	flag.Parse()

	lis, err := net.Listen("tcp", *endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := getServerOptions(*tls)
	s := grpc.NewServer(opts...)
	hue.RegisterLightsServer(s, hue.Server{})
	version.RegisterVersionServer(s, version.Server{})

	// Register reflection service on gRPC server.
	// reflection.Register(s)

	log.WithField("endpoint", *endpoint).WithField("type", "grpc").
		WithField("version", version.Version).WithField("build", version.Build).Infof("Server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
