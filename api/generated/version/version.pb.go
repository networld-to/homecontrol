// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: version.proto

package version

import (
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

type VersionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Build   string `protobuf:"bytes,2,opt,name=build,proto3" json:"build,omitempty"`
}

func (x *VersionMessage) Reset() {
	*x = VersionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionMessage) ProtoMessage() {}

func (x *VersionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionMessage.ProtoReflect.Descriptor instead.
func (*VersionMessage) Descriptor() ([]byte, []int) {
	return file_version_proto_rawDescGZIP(), []int{0}
}

func (x *VersionMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionMessage) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

var File_version_proto protoreflect.FileDescriptor

var file_version_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x32, 0x48, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x17, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_version_proto_rawDescOnce sync.Once
	file_version_proto_rawDescData = file_version_proto_rawDesc
)

func file_version_proto_rawDescGZIP() []byte {
	file_version_proto_rawDescOnce.Do(func() {
		file_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_version_proto_rawDescData)
	})
	return file_version_proto_rawDescData
}

var file_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_version_proto_goTypes = []interface{}{
	(*VersionMessage)(nil), // 0: version.VersionMessage
}
var file_version_proto_depIdxs = []int32{
	0, // 0: version.Version.Version:input_type -> version.VersionMessage
	0, // 1: version.Version.Version:output_type -> version.VersionMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_version_proto_init() }
func file_version_proto_init() {
	if File_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionMessage); i {
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
			RawDescriptor: file_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_version_proto_goTypes,
		DependencyIndexes: file_version_proto_depIdxs,
		MessageInfos:      file_version_proto_msgTypes,
	}.Build()
	File_version_proto = out.File
	file_version_proto_rawDesc = nil
	file_version_proto_goTypes = nil
	file_version_proto_depIdxs = nil
}
