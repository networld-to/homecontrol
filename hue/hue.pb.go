// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hue.proto

/*
Package hue is a generated protocol buffer package.

It is generated from these files:
	hue.proto

It has these top-level messages:
	EchoMessage
	LightsRequest
	LightsResponse
	Groups
	Group
	SensorRequest
	Sensors
	Sensor
	State
*/
package hue

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoMessage struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *EchoMessage) Reset()                    { *m = EchoMessage{} }
func (m *EchoMessage) String() string            { return proto.CompactTextString(m) }
func (*EchoMessage) ProtoMessage()               {}
func (*EchoMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LightsRequest struct {
	Group int32 `protobuf:"varint,1,opt,name=Group" json:"Group,omitempty"`
	// Value between 0 and 1
	BrightnessPercent float32 `protobuf:"fixed32,2,opt,name=BrightnessPercent" json:"BrightnessPercent,omitempty"`
}

func (m *LightsRequest) Reset()                    { *m = LightsRequest{} }
func (m *LightsRequest) String() string            { return proto.CompactTextString(m) }
func (*LightsRequest) ProtoMessage()               {}
func (*LightsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LightsRequest) GetGroup() int32 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *LightsRequest) GetBrightnessPercent() float32 {
	if m != nil {
		return m.BrightnessPercent
	}
	return 0
}

type LightsResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *LightsResponse) Reset()                    { *m = LightsResponse{} }
func (m *LightsResponse) String() string            { return proto.CompactTextString(m) }
func (*LightsResponse) ProtoMessage()               {}
func (*LightsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LightsResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Groups struct {
	Groups []*Group `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
}

func (m *Groups) Reset()                    { *m = Groups{} }
func (m *Groups) String() string            { return proto.CompactTextString(m) }
func (*Groups) ProtoMessage()               {}
func (*Groups) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Groups) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type Group struct {
	ID         int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	On         bool   `protobuf:"varint,3,opt,name=On" json:"On,omitempty"`
	Brightness int32  `protobuf:"varint,4,opt,name=Brightness" json:"Brightness,omitempty"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Group) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *Group) GetBrightness() int32 {
	if m != nil {
		return m.Brightness
	}
	return 0
}

type SensorRequest struct {
	Sensor int32 `protobuf:"varint,1,opt,name=Sensor" json:"Sensor,omitempty"`
}

func (m *SensorRequest) Reset()                    { *m = SensorRequest{} }
func (m *SensorRequest) String() string            { return proto.CompactTextString(m) }
func (*SensorRequest) ProtoMessage()               {}
func (*SensorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SensorRequest) GetSensor() int32 {
	if m != nil {
		return m.Sensor
	}
	return 0
}

type Sensors struct {
	Sensors []*Sensor `protobuf:"bytes,1,rep,name=sensors" json:"sensors,omitempty"`
}

func (m *Sensors) Reset()                    { *m = Sensors{} }
func (m *Sensors) String() string            { return proto.CompactTextString(m) }
func (*Sensors) ProtoMessage()               {}
func (*Sensors) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Sensors) GetSensors() []*Sensor {
	if m != nil {
		return m.Sensors
	}
	return nil
}

type Sensor struct {
	ID               int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	UniqueID         string `protobuf:"bytes,2,opt,name=UniqueID" json:"UniqueID,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Type             string `protobuf:"bytes,4,opt,name=Type" json:"Type,omitempty"`
	ManufacturerName string `protobuf:"bytes,5,opt,name=ManufacturerName" json:"ManufacturerName,omitempty"`
	ModelID          string `protobuf:"bytes,6,opt,name=ModelID" json:"ModelID,omitempty"`
	SWVersion        string `protobuf:"bytes,7,opt,name=SWVersion" json:"SWVersion,omitempty"`
	State            *State `protobuf:"bytes,8,opt,name=State" json:"State,omitempty"`
}

func (m *Sensor) Reset()                    { *m = Sensor{} }
func (m *Sensor) String() string            { return proto.CompactTextString(m) }
func (*Sensor) ProtoMessage()               {}
func (*Sensor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Sensor) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Sensor) GetUniqueID() string {
	if m != nil {
		return m.UniqueID
	}
	return ""
}

func (m *Sensor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sensor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Sensor) GetManufacturerName() string {
	if m != nil {
		return m.ManufacturerName
	}
	return ""
}

func (m *Sensor) GetModelID() string {
	if m != nil {
		return m.ModelID
	}
	return ""
}

func (m *Sensor) GetSWVersion() string {
	if m != nil {
		return m.SWVersion
	}
	return ""
}

func (m *Sensor) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

type State struct {
	ButtonEvent int32  `protobuf:"varint,1,opt,name=ButtonEvent" json:"ButtonEvent,omitempty"`
	Dark        bool   `protobuf:"varint,2,opt,name=Dark" json:"Dark,omitempty"`
	Daylight    bool   `protobuf:"varint,3,opt,name=Daylight" json:"Daylight,omitempty"`
	LastUpdated string `protobuf:"bytes,4,opt,name=LastUpdated" json:"LastUpdated,omitempty"`
	LightLevel  int32  `protobuf:"varint,5,opt,name=LightLevel" json:"LightLevel,omitempty"`
	Presence    bool   `protobuf:"varint,6,opt,name=Presence" json:"Presence,omitempty"`
	Status      int32  `protobuf:"varint,7,opt,name=Status" json:"Status,omitempty"`
	Temperature int32  `protobuf:"varint,8,opt,name=Temperature" json:"Temperature,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *State) GetButtonEvent() int32 {
	if m != nil {
		return m.ButtonEvent
	}
	return 0
}

func (m *State) GetDark() bool {
	if m != nil {
		return m.Dark
	}
	return false
}

func (m *State) GetDaylight() bool {
	if m != nil {
		return m.Daylight
	}
	return false
}

func (m *State) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

func (m *State) GetLightLevel() int32 {
	if m != nil {
		return m.LightLevel
	}
	return 0
}

func (m *State) GetPresence() bool {
	if m != nil {
		return m.Presence
	}
	return false
}

func (m *State) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *State) GetTemperature() int32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func init() {
	proto.RegisterType((*EchoMessage)(nil), "hue.EchoMessage")
	proto.RegisterType((*LightsRequest)(nil), "hue.LightsRequest")
	proto.RegisterType((*LightsResponse)(nil), "hue.LightsResponse")
	proto.RegisterType((*Groups)(nil), "hue.Groups")
	proto.RegisterType((*Group)(nil), "hue.Group")
	proto.RegisterType((*SensorRequest)(nil), "hue.SensorRequest")
	proto.RegisterType((*Sensors)(nil), "hue.Sensors")
	proto.RegisterType((*Sensor)(nil), "hue.Sensor")
	proto.RegisterType((*State)(nil), "hue.State")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Lights service

type LightsClient interface {
	Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
	GetGroups(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*Groups, error)
	GetSensors(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (*Sensors, error)
	SwitchOn(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error)
	SwitchOff(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error)
	Blink(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error)
}

type lightsClient struct {
	cc *grpc.ClientConn
}

func NewLightsClient(cc *grpc.ClientConn) LightsClient {
	return &lightsClient{cc}
}

func (c *lightsClient) Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := grpc.Invoke(ctx, "/hue.Lights/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) GetGroups(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*Groups, error) {
	out := new(Groups)
	err := grpc.Invoke(ctx, "/hue.Lights/GetGroups", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) GetSensors(ctx context.Context, in *SensorRequest, opts ...grpc.CallOption) (*Sensors, error) {
	out := new(Sensors)
	err := grpc.Invoke(ctx, "/hue.Lights/GetSensors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) SwitchOn(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error) {
	out := new(LightsResponse)
	err := grpc.Invoke(ctx, "/hue.Lights/SwitchOn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) SwitchOff(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error) {
	out := new(LightsResponse)
	err := grpc.Invoke(ctx, "/hue.Lights/SwitchOff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) Blink(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error) {
	out := new(LightsResponse)
	err := grpc.Invoke(ctx, "/hue.Lights/Blink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lights service

type LightsServer interface {
	Echo(context.Context, *EchoMessage) (*EchoMessage, error)
	GetGroups(context.Context, *LightsRequest) (*Groups, error)
	GetSensors(context.Context, *SensorRequest) (*Sensors, error)
	SwitchOn(context.Context, *LightsRequest) (*LightsResponse, error)
	SwitchOff(context.Context, *LightsRequest) (*LightsResponse, error)
	Blink(context.Context, *LightsRequest) (*LightsResponse, error)
}

func RegisterLightsServer(s *grpc.Server, srv LightsServer) {
	s.RegisterService(&_Lights_serviceDesc, srv)
}

func _Lights_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightsServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hue.Lights/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightsServer).Echo(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lights_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightsServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hue.Lights/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightsServer).GetGroups(ctx, req.(*LightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lights_GetSensors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightsServer).GetSensors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hue.Lights/GetSensors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightsServer).GetSensors(ctx, req.(*SensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lights_SwitchOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightsServer).SwitchOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hue.Lights/SwitchOn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightsServer).SwitchOn(ctx, req.(*LightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lights_SwitchOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightsServer).SwitchOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hue.Lights/SwitchOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightsServer).SwitchOff(ctx, req.(*LightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lights_Blink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightsServer).Blink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hue.Lights/Blink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightsServer).Blink(ctx, req.(*LightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lights_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hue.Lights",
	HandlerType: (*LightsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Lights_Echo_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _Lights_GetGroups_Handler,
		},
		{
			MethodName: "GetSensors",
			Handler:    _Lights_GetSensors_Handler,
		},
		{
			MethodName: "SwitchOn",
			Handler:    _Lights_SwitchOn_Handler,
		},
		{
			MethodName: "SwitchOff",
			Handler:    _Lights_SwitchOff_Handler,
		},
		{
			MethodName: "Blink",
			Handler:    _Lights_Blink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hue.proto",
}

func init() { proto.RegisterFile("hue.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x4e, 0x1b, 0x3d,
	0x10, 0x65, 0x93, 0x6c, 0xb2, 0x3b, 0xf9, 0x40, 0x7c, 0x6e, 0x55, 0xad, 0x10, 0x6a, 0x23, 0x4b,
	0x55, 0x11, 0xa2, 0x11, 0xa2, 0x6a, 0x1f, 0x20, 0x0a, 0xa2, 0x91, 0xa0, 0x44, 0x0e, 0xb4, 0x17,
	0xbd, 0xda, 0x6e, 0x86, 0x64, 0x45, 0x62, 0x07, 0xdb, 0x4b, 0xc5, 0xb3, 0xf5, 0x6d, 0xfa, 0x06,
	0x7d, 0x83, 0xca, 0xb3, 0xde, 0xb0, 0x28, 0xbd, 0xe1, 0x6e, 0xe6, 0xcc, 0x8f, 0xcf, 0xf1, 0x8c,
	0x0d, 0xf1, 0xbc, 0xc0, 0xfe, 0x4a, 0x2b, 0xab, 0x58, 0x73, 0x5e, 0x20, 0x7f, 0x03, 0xdd, 0xd3,
	0x6c, 0xae, 0x2e, 0xd0, 0x98, 0x74, 0x86, 0x6c, 0x17, 0x9a, 0x4b, 0x33, 0x4b, 0x82, 0x5e, 0x70,
	0x10, 0x0b, 0x67, 0xf2, 0x09, 0x6c, 0x9f, 0xe7, 0xb3, 0xb9, 0x35, 0x02, 0xef, 0x0a, 0x34, 0x96,
	0xbd, 0x84, 0xf0, 0x4c, 0xab, 0x62, 0x45, 0x49, 0xa1, 0x28, 0x1d, 0x76, 0x04, 0xff, 0x0f, 0xb4,
	0xcb, 0x93, 0x68, 0xcc, 0x18, 0x75, 0x86, 0xd2, 0x26, 0x8d, 0x5e, 0x70, 0xd0, 0x10, 0x9b, 0x01,
	0x7e, 0x08, 0x3b, 0x55, 0x53, 0xb3, 0x52, 0xd2, 0x20, 0x4b, 0xa0, 0x63, 0x8a, 0x2c, 0x43, 0x63,
	0xa8, 0x6f, 0x24, 0x2a, 0x97, 0x1f, 0x41, 0x9b, 0x8e, 0x30, 0x8c, 0x43, 0x7b, 0x46, 0x56, 0x12,
	0xf4, 0x9a, 0x07, 0xdd, 0x13, 0xe8, 0x3b, 0x31, 0x14, 0x14, 0x3e, 0xc2, 0xbf, 0x7b, 0x76, 0x6c,
	0x07, 0x1a, 0xa3, 0xa1, 0xe7, 0xd8, 0x18, 0x0d, 0x19, 0x83, 0xd6, 0x97, 0x74, 0x89, 0xc4, 0x29,
	0x16, 0x64, 0xbb, 0x9c, 0x4b, 0x99, 0x34, 0xe9, 0xbc, 0xc6, 0xa5, 0x64, 0xaf, 0x01, 0x1e, 0xb9,
	0x26, 0x2d, 0xaa, 0xad, 0x21, 0xfc, 0x1d, 0x6c, 0x4f, 0x50, 0x1a, 0xa5, 0xab, 0xbb, 0x78, 0x05,
	0xed, 0x12, 0xf0, 0x07, 0x79, 0x8f, 0x1f, 0x43, 0xa7, 0xb4, 0x0c, 0x7b, 0x0b, 0x1d, 0x53, 0x9a,
	0x9e, 0x75, 0x97, 0x58, 0xfb, 0x3e, 0x55, 0x8c, 0xff, 0x0e, 0xaa, 0x56, 0x1b, 0xcc, 0xf7, 0x20,
	0xba, 0x96, 0xf9, 0x5d, 0x81, 0xa3, 0xa1, 0x67, 0xbf, 0xf6, 0xd7, 0xaa, 0x9a, 0x35, 0x55, 0x0c,
	0x5a, 0x57, 0x0f, 0x2b, 0x24, 0xfe, 0xb1, 0x20, 0x9b, 0x1d, 0xc2, 0xee, 0x45, 0x2a, 0x8b, 0x9b,
	0x34, 0xb3, 0x85, 0x46, 0x4d, 0x35, 0x21, 0xc5, 0x37, 0x70, 0x37, 0x8a, 0x0b, 0x35, 0xc5, 0xc5,
	0x68, 0x98, 0xb4, 0x29, 0xa5, 0x72, 0xd9, 0x3e, 0xc4, 0x93, 0x6f, 0x5f, 0x51, 0x9b, 0x5c, 0xc9,
	0xa4, 0x43, 0xb1, 0x47, 0x80, 0xf5, 0x20, 0x9c, 0xd8, 0xd4, 0x62, 0x12, 0xf5, 0x82, 0xf5, 0x74,
	0x08, 0x11, 0x65, 0x80, 0xff, 0x09, 0x7c, 0x0a, 0xeb, 0x41, 0x77, 0x50, 0x58, 0xab, 0xe4, 0xe9,
	0xbd, 0x5b, 0x94, 0x52, 0x6c, 0x1d, 0x72, 0x2a, 0x86, 0xa9, 0xbe, 0x25, 0xc5, 0x91, 0x20, 0xdb,
	0xdd, 0xc4, 0x30, 0x7d, 0x58, 0xb8, 0x79, 0xf8, 0xa9, 0xad, 0x7d, 0xd7, 0xf1, 0x3c, 0x35, 0xf6,
	0x7a, 0x35, 0x4d, 0x2d, 0x4e, 0xbd, 0xf8, 0x3a, 0xe4, 0xa6, 0x4b, 0x4b, 0x77, 0x8e, 0xf7, 0xb8,
	0x20, 0xf5, 0xa1, 0xa8, 0x21, 0xae, 0xfb, 0x58, 0xa3, 0x41, 0x99, 0x21, 0x09, 0x8f, 0xc4, 0xda,
	0xa7, 0x41, 0xdb, 0xd4, 0x16, 0x86, 0x64, 0xbb, 0x41, 0x93, 0xe7, 0x4e, 0xbd, 0xc2, 0xe5, 0x0a,
	0x75, 0xea, 0xee, 0x8f, 0x94, 0x87, 0xa2, 0x0e, 0x9d, 0xfc, 0x6a, 0x40, 0xbb, 0xdc, 0x75, 0x76,
	0x04, 0x2d, 0xf7, 0xd6, 0xd8, 0x2e, 0xdd, 0x4c, 0xed, 0xd9, 0xed, 0x6d, 0x20, 0x7c, 0x8b, 0xf5,
	0x21, 0x3e, 0x43, 0xeb, 0x57, 0x9f, 0x51, 0xc2, 0x93, 0x87, 0xb8, 0xd7, 0x7d, 0x5c, 0x7f, 0xc3,
	0xb7, 0xd8, 0x31, 0xc0, 0x19, 0xda, 0x6a, 0xed, 0x58, 0x7d, 0xcb, 0x7c, 0xc1, 0x7f, 0x35, 0xcc,
	0x55, 0x7c, 0x84, 0x68, 0xf2, 0x33, 0xb7, 0xd9, 0xfc, 0x52, 0xfe, 0xf3, 0x80, 0x17, 0x4f, 0xb0,
	0xf2, 0xa1, 0xf2, 0x2d, 0xf6, 0x09, 0x62, 0x5f, 0x76, 0x73, 0xf3, 0x9c, 0xba, 0x13, 0x08, 0x07,
	0x8b, 0x5c, 0xde, 0x3e, 0xa3, 0x66, 0x70, 0x08, 0xfb, 0x99, 0x5a, 0xf6, 0xf3, 0xa9, 0x7a, 0x6f,
	0x50, 0xdf, 0xe7, 0x19, 0x9a, 0x3e, 0x01, 0xca, 0xba, 0x82, 0x41, 0xf4, 0xb9, 0xc0, 0xb1, 0xfb,
	0xcd, 0xc6, 0xc1, 0x8f, 0x36, 0x7d, 0x6b, 0x1f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xf1,
	0x1e, 0xe0, 0xe3, 0x04, 0x00, 0x00,
}
