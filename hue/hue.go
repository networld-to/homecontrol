package hue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"time"

	"context"

	log "github.com/sirupsen/logrus"

	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/lights"
	"github.com/heatxsink/go-hue/sensors"
)

// Server : The implementation of the Philips Hue gRPC service
type Server struct{}

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

/************************************************************************************
 * Start gRPC Service implementation
 */

// GetSensors : Returns all known sensors
func (s Server) GetSensors(ctx context.Context, in *SensorRequest) (*Sensors, error) {
	start := time.Now()
	ss := sensors.New(hueBridge.Bridge, hueBridge.Username)
	allSensors, err := ss.GetAllSensors()
	if err != nil {
		return nil, err
	}
	log.WithField("bridge_exec_time", time.Since(start)).
		WithField("call", "GetSensors").Print("Incoming gRPC call")

	sensorsResp := &Sensors{}
	sensorsResp.Sensors = make([]*Sensor, len(allSensors))
	for i, s := range allSensors {
		sensorsResp.Sensors[i] = &Sensor{
			ID:               int32(s.ID),
			UniqueID:         s.UniqueID,
			Name:             s.Name,
			Type:             s.Type,
			ModelID:          s.ModelID,
			SWVersion:        s.SWVersion,
			ManufacturerName: s.ManufacturerName,
			State: &State{
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
func (s Server) GetGroups(ctx context.Context, in *LightsRequest) (*Groups, error) {
	start := time.Now()
	gg := groups.New(hueBridge.Bridge, hueBridge.Username)
	allGroups, err := gg.GetAllGroups()
	if err != nil {
		return nil, err
	}
	log.WithField("bridge_exec_time", time.Since(start)).
		WithField("call", "GetGroups").Print("Incoming gRPC call")

	groups := &Groups{}
	groups.Groups = make([]*Group, len(allGroups))
	for i, g := range allGroups {
		groups.Groups[i] = &Group{ID: int32(g.ID), Name: g.Name, On: g.Action.On, Brightness: int32(g.Action.Bri)}
	}
	return groups, nil
}

// SwitchOn implements hue.Lights
func (s Server) SwitchOn(ctx context.Context, in *LightsRequest) (*LightsResponse, error) {
	log.WithField("message", in).
		WithField("call", "SwitchOn").Print("Incoming gRPC call")
	gg := groups.New(hueBridge.Bridge, hueBridge.Username)
	g, err := gg.GetGroup(int(in.GetGroup()))
	if err != nil {
		return nil, err
	}
	brightness := uint8(255 * in.GetBrightnessPercent())
	gg.SetGroupState(g.ID, lights.State{
		On:  true,
		Bri: brightness,
		Sat: 254,
		Hue: 10000,
	})
	return &LightsResponse{Success: true}, nil
}

// SwitchOff implements hue.Lights
func (s Server) SwitchOff(ctx context.Context, in *LightsRequest) (*LightsResponse, error) {
	log.WithField("message", in).
		WithField("call", "SwitchOff").Print("Incoming gRPC call")
	gg := groups.New(hueBridge.Bridge, hueBridge.Username)
	g, err := gg.GetGroup(int(in.GetGroup()))
	if err != nil {
		return nil, err
	}
	gg.SetGroupState(g.ID, lights.State{On: false})
	return &LightsResponse{Success: true}, nil
}

// Blink implement hue.Lights : It switches on and off the lights in a short time period.
// After done it resets to the previous state.
func (s Server) Blink(ctx context.Context, in *LightsRequest) (*LightsResponse, error) {
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
	return &LightsResponse{Success: true}, nil
}
