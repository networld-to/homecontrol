package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/user"
	"time"

	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/lights"
	"github.com/heatxsink/go-hue/sensors"

	"github.com/networld-to/homecontrol/hue"
	"github.com/networld-to/homecontrol/version"
	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	endpoint = ":50051"
)

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

type server struct{}

func (s *server) Echo(ctx context.Context, in *hue.EchoMessage) (*hue.EchoMessage, error) {
	log.WithField("message", in.Msg).
		WithField("call", "Echo").Print("Incoming gRPC call")
	return in, nil
}

func (s *server) Version(ctx context.Context, in *hue.VersionMessage) (*hue.VersionMessage, error) {
	log.WithField("message", in).
		WithField("call", "Version").Print("Incoming gRPC call")
	return &hue.VersionMessage{Version: version.Version, Build: version.Build}, nil
}

func (s *server) GetSensors(ctx context.Context, in *hue.SensorRequest) (*hue.Sensors, error) {
	start := time.Now()
	ss := sensors.New(hueBridge.Bridge, hueBridge.Username)
	allSensors, err := ss.GetAllSensors()
	if err != nil {
		return nil, err
	}
	log.WithField("bridge_exec_time", time.Since(start)).
		WithField("call", "GetSensors").Print("Incoming gRPC call")

	sensorsResp := &hue.Sensors{}
	sensorsResp.Sensors = make([]*hue.Sensor, len(allSensors))
	for i, s := range allSensors {
		sensorsResp.Sensors[i] = &hue.Sensor{
			ID:               int32(s.ID),
			UniqueID:         s.UniqueID,
			Name:             s.Name,
			Type:             s.Type,
			ModelID:          s.ModelID,
			SWVersion:        s.SWVersion,
			ManufacturerName: s.ManufacturerName,
			State: &hue.State{
				ButtonEvent: int32(s.State.ButtonEvent),
				Dark:        s.State.Dark,
				Daylight:    s.State.Daylight,
				LastUpdated: s.State.LastUpdated,
				LightLevel:  int32(s.State.LightLevel),
				Presence:    s.State.Presence,
				Status:      int32(s.State.Status),
				Temperature: int32(s.State.Temperature),
			},
		}
	}
	return sensorsResp, nil
}

func (s *server) GetGroups(ctx context.Context, in *hue.LightsRequest) (*hue.Groups, error) {
	start := time.Now()
	gg := groups.New(hueBridge.Bridge, hueBridge.Username)
	allGroups, err := gg.GetAllGroups()
	if err != nil {
		return nil, err
	}
	log.WithField("bridge_exec_time", time.Since(start)).
		WithField("call", "GetGroups").Print("Incoming gRPC call")

	groups := &hue.Groups{}
	groups.Groups = make([]*hue.Group, len(allGroups))
	for i, g := range allGroups {
		groups.Groups[i] = &hue.Group{ID: int32(g.ID), Name: g.Name, On: g.Action.On, Brightness: int32(g.Action.Bri)}
	}
	return groups, nil
}

// SwitchOn implements hue.Lights
func (s *server) SwitchOn(ctx context.Context, in *hue.LightsRequest) (*hue.LightsResponse, error) {
	log.WithField("message", in).
		WithField("call", "SwitchOn").Print("Incoming gRPC call")
	gg := groups.New(hueBridge.Bridge, hueBridge.Username)
	g, err := gg.GetGroup(int(in.GetGroup()))
	if err != nil {
		return nil, err
	}
	brightness := uint8(255 * in.GetBrightnessPercent())
	gg.SetGroupState(g.ID, lights.State{On: true, Bri: brightness})
	return &hue.LightsResponse{Success: true}, nil
}

// SwitchOff implements hue.Lights
func (s *server) SwitchOff(ctx context.Context, in *hue.LightsRequest) (*hue.LightsResponse, error) {
	log.WithField("message", in).
		WithField("call", "SwitchOff").Print("Incoming gRPC call")
	gg := groups.New(hueBridge.Bridge, hueBridge.Username)
	g, err := gg.GetGroup(int(in.GetGroup()))
	if err != nil {
		return nil, err
	}
	gg.SetGroupState(g.ID, lights.State{On: false})
	return &hue.LightsResponse{Success: true}, nil
}

func (s *server) Blink(ctx context.Context, in *hue.LightsRequest) (*hue.LightsResponse, error) {
	log.WithField("message", in).
		WithField("call", "Blink").Print("Incoming gRPC call")
	gg := groups.New(hueBridge.Bridge, hueBridge.Username)
	g, err := gg.GetGroup(int(in.GetGroup()))
	if err != nil {
		return nil, err
	}

	oldState := g.Action

	brightness := uint8(255 * in.GetBrightnessPercent())
	gg.SetGroupState(g.ID, lights.State{On: true, Bri: brightness, Alert: "lselect"})
	gg.SetGroupState(g.ID, oldState)
	return &hue.LightsResponse{Success: true}, nil
}

// HueBridge : Connection and authentication information related to the Philips Hue Bridge
type HueBridge struct {
	Bridge     string `json:"bridge"`
	Username   string `json:"username"`
	Devicetype string `json:"devicetype"`
}

var hueBridge HueBridge

func loadHueBridgeConfig() HueBridge {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadFile(usr.HomeDir + "/.philips-hue.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &hueBridge)
	if err != nil {
		fmt.Print("Error:", err)
	}
	return hueBridge
}

func getServerOptions() []grpc.ServerOption {
	opts := []grpc.ServerOption{}
	return opts
}

var tls = flag.Bool("tls", false, "Activate TLS communication channel encryption.")

func main() {
	flag.Parse()
	loadHueBridgeConfig()
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := getServerOptions()
	if *tls {
		cred, err := credentials.NewServerTLSFromFile(os.Getenv("HOME")+"/.homecontrol/server.crt", os.Getenv("HOME")+"/.homecontrol/server.key")
		log.WithField("tls", true).Info("TLS Enabled")
		if err != nil {
			log.Fatalf("failed TLS: %v", err)
		}
		opts = append(opts, grpc.Creds(cred))
	} else {
		log.Warn("Insecure connection. No TLS")
	}
	s := grpc.NewServer(opts...)
	hue.RegisterLightsServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.WithField("endpoint", endpoint).WithField("type", "grpc").
		WithField("version", version.Version).WithField("build", version.Build).Infof("Server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
