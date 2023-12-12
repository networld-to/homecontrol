// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: hue.proto

package hue

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LightsClient is the client API for Lights service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
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
// All implementations must embed UnimplementedLightsServer
// for forward compatibility
type LightsServer interface {
	GetGroups(context.Context, *Empty) (*Groups, error)
	GetSensors(context.Context, *Empty) (*Sensors, error)
	SwitchOn(context.Context, *LightsRequest) (*LightsResponse, error)
	SwitchOff(context.Context, *LightsRequest) (*LightsResponse, error)
	Blink(context.Context, *LightsRequest) (*LightsResponse, error)
	mustEmbedUnimplementedLightsServer()
}

// UnimplementedLightsServer must be embedded to have forward compatible implementations.
type UnimplementedLightsServer struct {
}

func (UnimplementedLightsServer) GetGroups(context.Context, *Empty) (*Groups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedLightsServer) GetSensors(context.Context, *Empty) (*Sensors, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensors not implemented")
}
func (UnimplementedLightsServer) SwitchOn(context.Context, *LightsRequest) (*LightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchOn not implemented")
}
func (UnimplementedLightsServer) SwitchOff(context.Context, *LightsRequest) (*LightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchOff not implemented")
}
func (UnimplementedLightsServer) Blink(context.Context, *LightsRequest) (*LightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Blink not implemented")
}
func (UnimplementedLightsServer) mustEmbedUnimplementedLightsServer() {}

// UnsafeLightsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LightsServer will
// result in compilation errors.
type UnsafeLightsServer interface {
	mustEmbedUnimplementedLightsServer()
}

func RegisterLightsServer(s grpc.ServiceRegistrar, srv LightsServer) {
	s.RegisterService(&Lights_ServiceDesc, srv)
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

// Lights_ServiceDesc is the grpc.ServiceDesc for Lights service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lights_ServiceDesc = grpc.ServiceDesc{
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
