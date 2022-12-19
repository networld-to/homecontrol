package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"time"

	"context"

	log "github.com/sirupsen/logrus"

	hue "github.com/networld-to/homecontrol/api/generated/hue"

	"github.com/networld-to/go-hue/groups"
	"github.com/networld-to/go-hue/lights"
	"github.com/networld-to/go-hue/sensors"
)

// HueServer : The implementation of the Philips Hue gRPC service
type HueServer struct{
	hue.LightsServer
}

// Bridge : Connection and authentication information related to the Philips Hue Bridge
type Bridge struct {
	Bridge     string `json:"bridge"`
	Username   string `json:"username"`
	Devicetype string `json:"devicetype"`
}

var hueBridge Bridge

// LoadHueBridgeConfig : Loads the Philips Hue configuration file from file.
func LoadHueBridgeConfig() Bridge {
	usr, err := user.Current()
	var prefix string
	if err == nil {
		prefix = usr.HomeDir
	}
	content, err := ioutil.ReadFile(prefix + "/.philips-hue.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &hueBridge)
	if err != nil {
		fmt.Print("Error:", err)
	}
	return hueBridge
}

/************************************************************************************
 * Start gRPC Service implementation
 */

// GetSensors : Returns all known sensors
func (s HueServer) GetSensors(ctx context.Context, in *hue.Empty) (*hue.Sensors, error) {
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

// GetGroups implements hue.Lights : Returns group of lights.
func (s HueServer) GetGroups(ctx context.Context, in *hue.Empty) (*hue.Groups, error) {
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
func (s HueServer) SwitchOn(ctx context.Context, in *hue.LightsRequest) (*hue.LightsResponse, error) {
	log.WithField("message", in).
		WithField("call", "SwitchOn").Print("Incoming gRPC call")
	gg := groups.New(hueBridge.Bridge, hueBridge.Username)
	g, err := gg.GetGroup(int(in.GetGroup()))
	if err != nil {
		return nil, err
	}
	brightness := uint8(255 * in.GetBrightnessPercent())

	saturation := uint8(254 * in.GetSaturationPercent()) // 254 is the most saturated color, 0 is the least saturated color (white)
	hueInt := uint16(in.GetHue())                        // Value between 0 and 65535 with Red=5535 and Green=25500 and Blue=46920

	gg.SetGroupState(g.ID, lights.State{
		On:  true,
		Bri: brightness,
		Sat: saturation,
		Hue: hueInt,
	})
	return &hue.LightsResponse{Success: true}, nil
}

// SwitchOff implements hue.Lights
func (s HueServer) SwitchOff(ctx context.Context, in *hue.LightsRequest) (*hue.LightsResponse, error) {
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

// Blink implement hue.Lights : It switches on and off the lights in a short time period.
// After done it resets to the previous state.
func (s HueServer) Blink(ctx context.Context, in *hue.LightsRequest) (*hue.LightsResponse, error) {
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
