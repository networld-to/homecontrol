package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/user"
	"time"

	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/lights"

	"github.com/networld-to/homecontrol/hue"
	"github.com/networld-to/homecontrol/version"
	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

func main() {
	loadHueBridgeConfig()
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hue.RegisterLightsServer(s, &server{})
	// Register reflection service on gRPC server.
	//reflection.Register(s)
	log.WithField("endpoint", endpoint).WithField("type", "grpc").
		WithField("version", version.Version).WithField("build", version.Build).Infof("Server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}