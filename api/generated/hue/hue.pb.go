// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: hue.proto

package hue

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{0}
}

type LightsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group int32 `protobuf:"varint,1,opt,name=Group,proto3" json:"Group,omitempty"`
	// Value between 0 and 1
	BrightnessPercent float32 `protobuf:"fixed32,2,opt,name=BrightnessPercent,proto3" json:"BrightnessPercent,omitempty"`
	// Value between 0 and 1
	SaturationPercent float32 `protobuf:"fixed32,3,opt,name=SaturationPercent,proto3" json:"SaturationPercent,omitempty"`
	// Value between 0 and 65535
	// Red: 65535
	// Green: 25500
	// Blue: 46920
	Hue float32 `protobuf:"fixed32,4,opt,name=Hue,proto3" json:"Hue,omitempty"`
}

func (x *LightsRequest) Reset() {
	*x = LightsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightsRequest) ProtoMessage() {}

func (x *LightsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightsRequest.ProtoReflect.Descriptor instead.
func (*LightsRequest) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{1}
}

func (x *LightsRequest) GetGroup() int32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *LightsRequest) GetBrightnessPercent() float32 {
	if x != nil {
		return x.BrightnessPercent
	}
	return 0
}

func (x *LightsRequest) GetSaturationPercent() float32 {
	if x != nil {
		return x.SaturationPercent
	}
	return 0
}

func (x *LightsRequest) GetHue() float32 {
	if x != nil {
		return x.Hue
	}
	return 0
}

type LightsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LightsResponse) Reset() {
	*x = LightsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightsResponse) ProtoMessage() {}

func (x *LightsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightsResponse.ProtoReflect.Descriptor instead.
func (*LightsResponse) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{2}
}

func (x *LightsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Groups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Groups) Reset() {
	*x = Groups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Groups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Groups) ProtoMessage() {}

func (x *Groups) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Groups.ProtoReflect.Descriptor instead.
func (*Groups) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{3}
}

func (x *Groups) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	On         bool   `protobuf:"varint,3,opt,name=On,proto3" json:"On,omitempty"`
	Brightness int32  `protobuf:"varint,4,opt,name=Brightness,proto3" json:"Brightness,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{4}
}

func (x *Group) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetOn() bool {
	if x != nil {
		return x.On
	}
	return false
}

func (x *Group) GetBrightness() int32 {
	if x != nil {
		return x.Brightness
	}
	return 0
}

type SensorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor int32 `protobuf:"varint,1,opt,name=Sensor,proto3" json:"Sensor,omitempty"`
}

func (x *SensorRequest) Reset() {
	*x = SensorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorRequest) ProtoMessage() {}

func (x *SensorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorRequest.ProtoReflect.Descriptor instead.
func (*SensorRequest) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{5}
}

func (x *SensorRequest) GetSensor() int32 {
	if x != nil {
		return x.Sensor
	}
	return 0
}

type Sensors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensors []*Sensor `protobuf:"bytes,1,rep,name=sensors,proto3" json:"sensors,omitempty"`
}

func (x *Sensors) Reset() {
	*x = Sensors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sensors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sensors) ProtoMessage() {}

func (x *Sensors) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sensors.ProtoReflect.Descriptor instead.
func (*Sensors) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{6}
}

func (x *Sensors) GetSensors() []*Sensor {
	if x != nil {
		return x.Sensors
	}
	return nil
}

type Sensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UniqueID         string `protobuf:"bytes,2,opt,name=UniqueID,proto3" json:"UniqueID,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Type             string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	ManufacturerName string `protobuf:"bytes,5,opt,name=ManufacturerName,proto3" json:"ManufacturerName,omitempty"`
	ModelID          string `protobuf:"bytes,6,opt,name=ModelID,proto3" json:"ModelID,omitempty"`
	SWVersion        string `protobuf:"bytes,7,opt,name=SWVersion,proto3" json:"SWVersion,omitempty"`
	State            *State `protobuf:"bytes,8,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *Sensor) Reset() {
	*x = Sensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sensor) ProtoMessage() {}

func (x *Sensor) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sensor.ProtoReflect.Descriptor instead.
func (*Sensor) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{7}
}

func (x *Sensor) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Sensor) GetUniqueID() string {
	if x != nil {
		return x.UniqueID
	}
	return ""
}

func (x *Sensor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sensor) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Sensor) GetManufacturerName() string {
	if x != nil {
		return x.ManufacturerName
	}
	return ""
}

func (x *Sensor) GetModelID() string {
	if x != nil {
		return x.ModelID
	}
	return ""
}

func (x *Sensor) GetSWVersion() string {
	if x != nil {
		return x.SWVersion
	}
	return ""
}

func (x *Sensor) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ButtonEvent int32  `protobuf:"varint,1,opt,name=ButtonEvent,proto3" json:"ButtonEvent,omitempty"`
	Dark        bool   `protobuf:"varint,2,opt,name=Dark,proto3" json:"Dark,omitempty"`
	Daylight    bool   `protobuf:"varint,3,opt,name=Daylight,proto3" json:"Daylight,omitempty"`
	LastUpdated string `protobuf:"bytes,4,opt,name=LastUpdated,proto3" json:"LastUpdated,omitempty"`
	LightLevel  int32  `protobuf:"varint,5,opt,name=LightLevel,proto3" json:"LightLevel,omitempty"`
	Presence    bool   `protobuf:"varint,6,opt,name=Presence,proto3" json:"Presence,omitempty"`
	Status      int32  `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`
	Temperature int32  `protobuf:"varint,8,opt,name=Temperature,proto3" json:"Temperature,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hue_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_hue_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_hue_proto_rawDescGZIP(), []int{8}
}

func (x *State) GetButtonEvent() int32 {
	if x != nil {
		return x.ButtonEvent
	}
	return 0
}

func (x *State) GetDark() bool {
	if x != nil {
		return x.Dark
	}
	return false
}

func (x *State) GetDaylight() bool {
	if x != nil {
		return x.Daylight
	}
	return false
}

func (x *State) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *State) GetLightLevel() int32 {
	if x != nil {
		return x.LightLevel
	}
	return 0
}

func (x *State) GetPresence() bool {
	if x != nil {
		return x.Presence
	}
	return false
}

func (x *State) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *State) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

var File_hue_proto protoreflect.FileDescriptor

var file_hue_proto_rawDesc = []byte{
	0x0a, 0x09, 0x68, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x68, 0x75, 0x65,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x42, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x53, 0x61, 0x74, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x53, 0x61, 0x74, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x48, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x48, 0x75, 0x65, 0x22,
	0x2a, 0x0a, 0x0e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x5b, 0x0a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x4f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x42, 0x72, 0x69, 0x67,
	0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22,
	0x30, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x68, 0x75,
	0x65, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x73, 0x22, 0xe2, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x57, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x57, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x44, 0x61, 0x72, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x44, 0x61, 0x79, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0xfd, 0x01, 0x0a, 0x06, 0x4c,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x0a, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b,
	0x2e, 0x68, 0x75, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x00, 0x12, 0x28, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x0a, 0x2e, 0x68, 0x75,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x4f, 0x6e, 0x12, 0x12, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36,
	0x0a, 0x09, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x4f, 0x66, 0x66, 0x12, 0x12, 0x2e, 0x68, 0x75,
	0x65, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x05, 0x42, 0x6c, 0x69, 0x6e, 0x6b, 0x12,
	0x12, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x68, 0x75, 0x65, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b,
	0x68, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hue_proto_rawDescOnce sync.Once
	file_hue_proto_rawDescData = file_hue_proto_rawDesc
)

func file_hue_proto_rawDescGZIP() []byte {
	file_hue_proto_rawDescOnce.Do(func() {
		file_hue_proto_rawDescData = protoimpl.X.CompressGZIP(file_hue_proto_rawDescData)
	})
	return file_hue_proto_rawDescData
}

var file_hue_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_hue_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: hue.Empty
	(*LightsRequest)(nil),  // 1: hue.LightsRequest
	(*LightsResponse)(nil), // 2: hue.LightsResponse
	(*Groups)(nil),         // 3: hue.Groups
	(*Group)(nil),          // 4: hue.Group
	(*SensorRequest)(nil),  // 5: hue.SensorRequest
	(*Sensors)(nil),        // 6: hue.Sensors
	(*Sensor)(nil),         // 7: hue.Sensor
	(*State)(nil),          // 8: hue.State
}
var file_hue_proto_depIdxs = []int32{
	4, // 0: hue.Groups.groups:type_name -> hue.Group
	7, // 1: hue.Sensors.sensors:type_name -> hue.Sensor
	8, // 2: hue.Sensor.State:type_name -> hue.State
	0, // 3: hue.Lights.GetGroups:input_type -> hue.Empty
	0, // 4: hue.Lights.GetSensors:input_type -> hue.Empty
	1, // 5: hue.Lights.SwitchOn:input_type -> hue.LightsRequest
	1, // 6: hue.Lights.SwitchOff:input_type -> hue.LightsRequest
	1, // 7: hue.Lights.Blink:input_type -> hue.LightsRequest
	3, // 8: hue.Lights.GetGroups:output_type -> hue.Groups
	6, // 9: hue.Lights.GetSensors:output_type -> hue.Sensors
	2, // 10: hue.Lights.SwitchOn:output_type -> hue.LightsResponse
	2, // 11: hue.Lights.SwitchOff:output_type -> hue.LightsResponse
	2, // 12: hue.Lights.Blink:output_type -> hue.LightsResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hue_proto_init() }
func file_hue_proto_init() {
	if File_hue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hue_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Groups); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hue_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hue_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hue_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sensors); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hue_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sensor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hue_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hue_proto_goTypes,
		DependencyIndexes: file_hue_proto_depIdxs,
		MessageInfos:      file_hue_proto_msgTypes,
	}.Build()
	File_hue_proto = out.File
	file_hue_proto_rawDesc = nil
	file_hue_proto_goTypes = nil
	file_hue_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LightsClient is the client API for Lights service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LightsClient interface {
	GetGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Groups, error)
	GetSensors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Sensors, error)
	SwitchOn(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error)
	SwitchOff(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error)
	Blink(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error)
}

type lightsClient struct {
	cc grpc.ClientConnInterface
}

func NewLightsClient(cc grpc.ClientConnInterface) LightsClient {
	return &lightsClient{cc}
}

func (c *lightsClient) GetGroups(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Groups, error) {
	out := new(Groups)
	err := c.cc.Invoke(ctx, "/hue.Lights/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) GetSensors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Sensors, error) {
	out := new(Sensors)
	err := c.cc.Invoke(ctx, "/hue.Lights/GetSensors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) SwitchOn(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error) {
	out := new(LightsResponse)
	err := c.cc.Invoke(ctx, "/hue.Lights/SwitchOn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) SwitchOff(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error) {
	out := new(LightsResponse)
	err := c.cc.Invoke(ctx, "/hue.Lights/SwitchOff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightsClient) Blink(ctx context.Context, in *LightsRequest, opts ...grpc.CallOption) (*LightsResponse, error) {
	out := new(LightsResponse)
	err := c.cc.Invoke(ctx, "/hue.Lights/Blink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LightsServer is the server API for Lights service.
type LightsServer interface {
	GetGroups(context.Context, *Empty) (*Groups, error)
	GetSensors(context.Context, *Empty) (*Sensors, error)
	SwitchOn(context.Context, *LightsRequest) (*LightsResponse, error)
	SwitchOff(context.Context, *LightsRequest) (*LightsResponse, error)
	Blink(context.Context, *LightsRequest) (*LightsResponse, error)
}

// UnimplementedLightsServer can be embedded to have forward compatible implementations.
type UnimplementedLightsServer struct {
}

func (*UnimplementedLightsServer) GetGroups(context.Context, *Empty) (*Groups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (*UnimplementedLightsServer) GetSensors(context.Context, *Empty) (*Sensors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensors not implemented")
}
func (*UnimplementedLightsServer) SwitchOn(context.Context, *LightsRequest) (*LightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchOn not implemented")
}
func (*UnimplementedLightsServer) SwitchOff(context.Context, *LightsRequest) (*LightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchOff not implemented")
}
func (*UnimplementedLightsServer) Blink(context.Context, *LightsRequest) (*LightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Blink not implemented")
}

func RegisterLightsServer(s *grpc.Server, srv LightsServer) {
	s.RegisterService(&_Lights_serviceDesc, srv)
}

func _Lights_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
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
		return srv.(LightsServer).GetGroups(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lights_GetSensors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
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
		return srv.(LightsServer).GetSensors(ctx, req.(*Empty))
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