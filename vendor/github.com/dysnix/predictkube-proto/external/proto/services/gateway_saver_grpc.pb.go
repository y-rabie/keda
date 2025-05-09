// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: proto/services/gateway_saver.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GatewaySaverService_GetMetricsOffset_FullMethodName = "/services.GatewaySaverService/GetMetricsOffset"
	GatewaySaverService_SendMetrics_FullMethodName      = "/services.GatewaySaverService/SendMetrics"
)

// GatewaySaverServiceClient is the client API for GatewaySaverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewaySaverServiceClient interface {
	GetMetricsOffset(ctx context.Context, in *ReqGetMetricsOffset, opts ...grpc.CallOption) (*ResGetMetricsOffset, error)
	SendMetrics(ctx context.Context, in *ReqSendMetrics, opts ...grpc.CallOption) (*ResSendMetrics, error)
}

type gatewaySaverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewaySaverServiceClient(cc grpc.ClientConnInterface) GatewaySaverServiceClient {
	return &gatewaySaverServiceClient{cc}
}

func (c *gatewaySaverServiceClient) GetMetricsOffset(ctx context.Context, in *ReqGetMetricsOffset, opts ...grpc.CallOption) (*ResGetMetricsOffset, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResGetMetricsOffset)
	err := c.cc.Invoke(ctx, GatewaySaverService_GetMetricsOffset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewaySaverServiceClient) SendMetrics(ctx context.Context, in *ReqSendMetrics, opts ...grpc.CallOption) (*ResSendMetrics, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResSendMetrics)
	err := c.cc.Invoke(ctx, GatewaySaverService_SendMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewaySaverServiceServer is the server API for GatewaySaverService service.
// All implementations must embed UnimplementedGatewaySaverServiceServer
// for forward compatibility.
type GatewaySaverServiceServer interface {
	GetMetricsOffset(context.Context, *ReqGetMetricsOffset) (*ResGetMetricsOffset, error)
	SendMetrics(context.Context, *ReqSendMetrics) (*ResSendMetrics, error)
	mustEmbedUnimplementedGatewaySaverServiceServer()
}

// UnimplementedGatewaySaverServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGatewaySaverServiceServer struct{}

func (UnimplementedGatewaySaverServiceServer) GetMetricsOffset(context.Context, *ReqGetMetricsOffset) (*ResGetMetricsOffset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricsOffset not implemented")
}
func (UnimplementedGatewaySaverServiceServer) SendMetrics(context.Context, *ReqSendMetrics) (*ResSendMetrics, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMetrics not implemented")
}
func (UnimplementedGatewaySaverServiceServer) mustEmbedUnimplementedGatewaySaverServiceServer() {}
func (UnimplementedGatewaySaverServiceServer) testEmbeddedByValue()                             {}

// UnsafeGatewaySaverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewaySaverServiceServer will
// result in compilation errors.
type UnsafeGatewaySaverServiceServer interface {
	mustEmbedUnimplementedGatewaySaverServiceServer()
}

func RegisterGatewaySaverServiceServer(s grpc.ServiceRegistrar, srv GatewaySaverServiceServer) {
	// If the following call pancis, it indicates UnimplementedGatewaySaverServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GatewaySaverService_ServiceDesc, srv)
}

func _GatewaySaverService_GetMetricsOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetMetricsOffset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaySaverServiceServer).GetMetricsOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewaySaverService_GetMetricsOffset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaySaverServiceServer).GetMetricsOffset(ctx, req.(*ReqGetMetricsOffset))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewaySaverService_SendMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSendMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewaySaverServiceServer).SendMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewaySaverService_SendMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewaySaverServiceServer).SendMetrics(ctx, req.(*ReqSendMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewaySaverService_ServiceDesc is the grpc.ServiceDesc for GatewaySaverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewaySaverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.GatewaySaverService",
	HandlerType: (*GatewaySaverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetricsOffset",
			Handler:    _GatewaySaverService_GetMetricsOffset_Handler,
		},
		{
			MethodName: "SendMetrics",
			Handler:    _GatewaySaverService_SendMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/services/gateway_saver.proto",
}
