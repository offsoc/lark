// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: pb_gw/gw.proto

package pb_gw

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

const (
	MessageGateway_SendTopicMessage_FullMethodName = "/pb_gw.MessageGateway/SendTopicMessage"
	MessageGateway_HealthCheck_FullMethodName      = "/pb_gw.MessageGateway/HealthCheck"
)

// MessageGatewayClient is the client API for MessageGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageGatewayClient interface {
	SendTopicMessage(ctx context.Context, in *SendTopicMessageReq, opts ...grpc.CallOption) (*SendTopicMessageResp, error)
	HealthCheck(ctx context.Context, in *HealthCheckReq, opts ...grpc.CallOption) (*HealthCheckResp, error)
}

type messageGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageGatewayClient(cc grpc.ClientConnInterface) MessageGatewayClient {
	return &messageGatewayClient{cc}
}

func (c *messageGatewayClient) SendTopicMessage(ctx context.Context, in *SendTopicMessageReq, opts ...grpc.CallOption) (*SendTopicMessageResp, error) {
	out := new(SendTopicMessageResp)
	err := c.cc.Invoke(ctx, MessageGateway_SendTopicMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageGatewayClient) HealthCheck(ctx context.Context, in *HealthCheckReq, opts ...grpc.CallOption) (*HealthCheckResp, error) {
	out := new(HealthCheckResp)
	err := c.cc.Invoke(ctx, MessageGateway_HealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageGatewayServer is the server API for MessageGateway service.
// All implementations must embed UnimplementedMessageGatewayServer
// for forward compatibility
type MessageGatewayServer interface {
	SendTopicMessage(context.Context, *SendTopicMessageReq) (*SendTopicMessageResp, error)
	HealthCheck(context.Context, *HealthCheckReq) (*HealthCheckResp, error)
	mustEmbedUnimplementedMessageGatewayServer()
}

// UnimplementedMessageGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedMessageGatewayServer struct {
}

func (UnimplementedMessageGatewayServer) SendTopicMessage(context.Context, *SendTopicMessageReq) (*SendTopicMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTopicMessage not implemented")
}
func (UnimplementedMessageGatewayServer) HealthCheck(context.Context, *HealthCheckReq) (*HealthCheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedMessageGatewayServer) mustEmbedUnimplementedMessageGatewayServer() {}

// UnsafeMessageGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageGatewayServer will
// result in compilation errors.
type UnsafeMessageGatewayServer interface {
	mustEmbedUnimplementedMessageGatewayServer()
}

func RegisterMessageGatewayServer(s grpc.ServiceRegistrar, srv MessageGatewayServer) {
	s.RegisterService(&MessageGateway_ServiceDesc, srv)
}

func _MessageGateway_SendTopicMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTopicMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageGatewayServer).SendTopicMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageGateway_SendTopicMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageGatewayServer).SendTopicMessage(ctx, req.(*SendTopicMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageGateway_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageGatewayServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageGateway_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageGatewayServer).HealthCheck(ctx, req.(*HealthCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageGateway_ServiceDesc is the grpc.ServiceDesc for MessageGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_gw.MessageGateway",
	HandlerType: (*MessageGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTopicMessage",
			Handler:    _MessageGateway_SendTopicMessage_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _MessageGateway_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_gw/gw.proto",
}
