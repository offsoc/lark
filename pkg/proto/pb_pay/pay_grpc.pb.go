// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: pb_pay/pay.proto

package pb_pay

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
	Pay_AlipayReturn_FullMethodName = "/pb_pay.Pay/AlipayReturn"
	Pay_AlipayNotify_FullMethodName = "/pb_pay.Pay/AlipayNotify"
	Pay_PaypalReturn_FullMethodName = "/pb_pay.Pay/PaypalReturn"
	Pay_PaypalCancel_FullMethodName = "/pb_pay.Pay/PaypalCancel"
	Pay_PaypalNotify_FullMethodName = "/pb_pay.Pay/PaypalNotify"
)

// PayClient is the client API for Pay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayClient interface {
	AlipayReturn(ctx context.Context, in *AlipayReturnReq, opts ...grpc.CallOption) (*AlipayReturnResp, error)
	AlipayNotify(ctx context.Context, in *AlipayNotifyReq, opts ...grpc.CallOption) (*AlipayNotifyResp, error)
	PaypalReturn(ctx context.Context, in *PaypalReturnReq, opts ...grpc.CallOption) (*PaypalReturnResp, error)
	PaypalCancel(ctx context.Context, in *PaypalCancelReq, opts ...grpc.CallOption) (*PaypalCancelResp, error)
	PaypalNotify(ctx context.Context, in *PaypalNotifyReq, opts ...grpc.CallOption) (*PaypalNotifyResp, error)
}

type payClient struct {
	cc grpc.ClientConnInterface
}

func NewPayClient(cc grpc.ClientConnInterface) PayClient {
	return &payClient{cc}
}

func (c *payClient) AlipayReturn(ctx context.Context, in *AlipayReturnReq, opts ...grpc.CallOption) (*AlipayReturnResp, error) {
	out := new(AlipayReturnResp)
	err := c.cc.Invoke(ctx, Pay_AlipayReturn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) AlipayNotify(ctx context.Context, in *AlipayNotifyReq, opts ...grpc.CallOption) (*AlipayNotifyResp, error) {
	out := new(AlipayNotifyResp)
	err := c.cc.Invoke(ctx, Pay_AlipayNotify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) PaypalReturn(ctx context.Context, in *PaypalReturnReq, opts ...grpc.CallOption) (*PaypalReturnResp, error) {
	out := new(PaypalReturnResp)
	err := c.cc.Invoke(ctx, Pay_PaypalReturn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) PaypalCancel(ctx context.Context, in *PaypalCancelReq, opts ...grpc.CallOption) (*PaypalCancelResp, error) {
	out := new(PaypalCancelResp)
	err := c.cc.Invoke(ctx, Pay_PaypalCancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) PaypalNotify(ctx context.Context, in *PaypalNotifyReq, opts ...grpc.CallOption) (*PaypalNotifyResp, error) {
	out := new(PaypalNotifyResp)
	err := c.cc.Invoke(ctx, Pay_PaypalNotify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServer is the server API for Pay service.
// All implementations must embed UnimplementedPayServer
// for forward compatibility
type PayServer interface {
	AlipayReturn(context.Context, *AlipayReturnReq) (*AlipayReturnResp, error)
	AlipayNotify(context.Context, *AlipayNotifyReq) (*AlipayNotifyResp, error)
	PaypalReturn(context.Context, *PaypalReturnReq) (*PaypalReturnResp, error)
	PaypalCancel(context.Context, *PaypalCancelReq) (*PaypalCancelResp, error)
	PaypalNotify(context.Context, *PaypalNotifyReq) (*PaypalNotifyResp, error)
	mustEmbedUnimplementedPayServer()
}

// UnimplementedPayServer must be embedded to have forward compatible implementations.
type UnimplementedPayServer struct {
}

func (UnimplementedPayServer) AlipayReturn(context.Context, *AlipayReturnReq) (*AlipayReturnResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlipayReturn not implemented")
}
func (UnimplementedPayServer) AlipayNotify(context.Context, *AlipayNotifyReq) (*AlipayNotifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlipayNotify not implemented")
}
func (UnimplementedPayServer) PaypalReturn(context.Context, *PaypalReturnReq) (*PaypalReturnResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaypalReturn not implemented")
}
func (UnimplementedPayServer) PaypalCancel(context.Context, *PaypalCancelReq) (*PaypalCancelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaypalCancel not implemented")
}
func (UnimplementedPayServer) PaypalNotify(context.Context, *PaypalNotifyReq) (*PaypalNotifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaypalNotify not implemented")
}
func (UnimplementedPayServer) mustEmbedUnimplementedPayServer() {}

// UnsafePayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayServer will
// result in compilation errors.
type UnsafePayServer interface {
	mustEmbedUnimplementedPayServer()
}

func RegisterPayServer(s grpc.ServiceRegistrar, srv PayServer) {
	s.RegisterService(&Pay_ServiceDesc, srv)
}

func _Pay_AlipayReturn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlipayReturnReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).AlipayReturn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pay_AlipayReturn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).AlipayReturn(ctx, req.(*AlipayReturnReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_AlipayNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlipayNotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).AlipayNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pay_AlipayNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).AlipayNotify(ctx, req.(*AlipayNotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_PaypalReturn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaypalReturnReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).PaypalReturn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pay_PaypalReturn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).PaypalReturn(ctx, req.(*PaypalReturnReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_PaypalCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaypalCancelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).PaypalCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pay_PaypalCancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).PaypalCancel(ctx, req.(*PaypalCancelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_PaypalNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaypalNotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).PaypalNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pay_PaypalNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).PaypalNotify(ctx, req.(*PaypalNotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pay_ServiceDesc is the grpc.ServiceDesc for Pay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_pay.Pay",
	HandlerType: (*PayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AlipayReturn",
			Handler:    _Pay_AlipayReturn_Handler,
		},
		{
			MethodName: "AlipayNotify",
			Handler:    _Pay_AlipayNotify_Handler,
		},
		{
			MethodName: "PaypalReturn",
			Handler:    _Pay_PaypalReturn_Handler,
		},
		{
			MethodName: "PaypalCancel",
			Handler:    _Pay_PaypalCancel_Handler,
		},
		{
			MethodName: "PaypalNotify",
			Handler:    _Pay_PaypalNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_pay/pay.proto",
}
