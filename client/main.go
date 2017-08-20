package main

import (
	"flag"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"

	"github.com/networld-to/homecontrol/hue"
	"github.com/networld-to/homecontrol/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	address = flag.String("host", "127.0.0.1:50051", "The gRPC service endpoint that will be contacted.")
	cmd     = flag.String("cmd", "groups", "The command that will be executed.")
	group   = flag.Int("group", 2, "The light group ID.")
	tls     = flag.Bool("tls", false, "Activate TLS communication channel encryption.")
)

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getGroups(client hue.LightsClient) *hue.Groups {
	opts := getCallOptions()
	resp, err := client.GetGroups(context.Background(), &hue.LightsRequest{}, opts...)
	must(err)
	return resp
}

func getSensors(client hue.LightsClient) *hue.Sensors {
	opts := getCallOptions()
	resp, err := client.GetSensors(context.Background(), &hue.SensorRequest{}, opts...)
	must(err)
	return resp
}

func switchOn(client hue.LightsClient, group int) {
	opts := getCallOptions()
	_, err := client.SwitchOn(context.Background(), &hue.LightsRequest{Group: int32(group), BrightnessPercent: 0.33}, opts...)
	must(err)
	// log.Printf("Lights switched on: %v", r.Success)
}

func switchOff(client hue.LightsClient, group int) {
	opts := getCallOptions()
	_, err := client.SwitchOff(context.Background(), &hue.LightsRequest{Group: int32(group)}, opts...)
	must(err)
	// log.Printf("Lights switched off: %v", r.Success)
}

func blink(client hue.LightsClient, group int) {
	opts := getCallOptions()
	_, err := client.Blink(context.Background(), &hue.LightsRequest{Group: int32(group), BrightnessPercent: 0.33}, opts...)
	must(err)
	// log.Printf("Lights switched on: %v", r.Success)
}

func getCallOptions() []grpc.CallOption {
	opts := []grpc.CallOption{
		grpc.FailFast(true),
		grpc.MaxCallSendMsgSize(1024),
		grpc.MaxCallRecvMsgSize(5120),
	}
	return opts
}

func getDialOptions() []grpc.DialOption {
	opts := []grpc.DialOption{
		// grpc.WithInsecure(),
		grpc.WithTimeout(10 * time.Second),
		grpc.WithBlock(),
		grpc.WithBackoffMaxDelay(1 * time.Second),
	}
	return opts
}

func main() {
	flag.Parse()

	start := time.Now()
	opts := getDialOptions()
	if *tls {
		cred, err := credentials.NewClientTLSFromFile(os.Getenv("HOME")+"/.homecontrol/server.crt", "Homecontrol")
		info := cred.Info()
		log.WithField("tls", true).
			WithField("tls_version", info.ProtocolVersion).
			WithField("tls_server_name", info.ServerName).
			Info("TLS Enabled")
		if err != nil {
			log.Fatalf("failed TLS: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		log.Warn("Insecure connection. No TLS")
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*address, opts...)
	must(err)
	defer conn.Close()

	hue := hue.NewLightsClient(conn)
	ver := version.NewVersionClient(conn)

	log.WithField("address", *address).WithField("cmd", *cmd).WithField("group", *group).Info("Executing command")
	switch *cmd {
	case "on":
		switchOn(hue, *group)
	case "off":
		switchOff(hue, *group)
	case "blink":
		blink(hue, *group)
	case "sensors":
		resp := getSensors(hue)
		for _, sensor := range resp.Sensors {
			log.Printf("Hue Light Sensor: %v", sensor.String())
		}
	case "version":
		log.WithField("version", version.Version).WithField("build", version.Build).Info("Client Version")
		resp, err := ver.Version(context.Background(), &version.VersionMessage{})
		must(err)
		log.WithField("version", resp.Version).WithField("build", resp.Build).Info("Server Version")
	default:
		resp := getGroups(hue)
		log.Printf("Hue Light Groups: %v", resp.Groups)
	}

	end := time.Now()
	log.WithField("exec_time", end.Sub(start)).Printf("Execution finished")
}
