// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonClient interface {
	SocketFeature(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleReply, error)
	ListFeature(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (Common_ListFeatureClient, error)
	JSONFeature(ctx context.Context, in *JsonRequest, opts ...grpc.CallOption) (*JsonReply, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) SocketFeature(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleReply, error) {
	out := new(SimpleReply)
	err := c.cc.Invoke(ctx, "/pb.Common/SocketFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) ListFeature(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (Common_ListFeatureClient, error) {
	stream, err := c.cc.NewStream(ctx, &Common_ServiceDesc.Streams[0], "/pb.Common/ListFeature", opts...)
	if err != nil {
		return nil, err
	}
	x := &commonListFeatureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Common_ListFeatureClient interface {
	Recv() (*SimpleReply, error)
	grpc.ClientStream
}

type commonListFeatureClient struct {
	grpc.ClientStream
}

func (x *commonListFeatureClient) Recv() (*SimpleReply, error) {
	m := new(SimpleReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commonClient) JSONFeature(ctx context.Context, in *JsonRequest, opts ...grpc.CallOption) (*JsonReply, error) {
	out := new(JsonReply)
	err := c.cc.Invoke(ctx, "/pb.Common/JSONFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
// All implementations must embed UnimplementedCommonServer
// for forward compatibility
type CommonServer interface {
	SocketFeature(context.Context, *SimpleRequest) (*SimpleReply, error)
	ListFeature(*SimpleRequest, Common_ListFeatureServer) error
	JSONFeature(context.Context, *JsonRequest) (*JsonReply, error)
	mustEmbedUnimplementedCommonServer()
}

// UnimplementedCommonServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (UnimplementedCommonServer) SocketFeature(context.Context, *SimpleRequest) (*SimpleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SocketFeature not implemented")
}
func (UnimplementedCommonServer) ListFeature(*SimpleRequest, Common_ListFeatureServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeature not implemented")
}
func (UnimplementedCommonServer) JSONFeature(context.Context, *JsonRequest) (*JsonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JSONFeature not implemented")
}
func (UnimplementedCommonServer) mustEmbedUnimplementedCommonServer() {}

// UnsafeCommonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServer will
// result in compilation errors.
type UnsafeCommonServer interface {
	mustEmbedUnimplementedCommonServer()
}

func RegisterCommonServer(s grpc.ServiceRegistrar, srv CommonServer) {
	s.RegisterService(&Common_ServiceDesc, srv)
}

func _Common_SocketFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SocketFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Common/SocketFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SocketFeature(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_ListFeature_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommonServer).ListFeature(m, &commonListFeatureServer{stream})
}

type Common_ListFeatureServer interface {
	Send(*SimpleReply) error
	grpc.ServerStream
}

type commonListFeatureServer struct {
	grpc.ServerStream
}

func (x *commonListFeatureServer) Send(m *SimpleReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Common_JSONFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).JSONFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Common/JSONFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).JSONFeature(ctx, req.(*JsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Common_ServiceDesc is the grpc.ServiceDesc for Common service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Common_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SocketFeature",
			Handler:    _Common_SocketFeature_Handler,
		},
		{
			MethodName: "JSONFeature",
			Handler:    _Common_JSONFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeature",
			Handler:       _Common_ListFeature_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}
