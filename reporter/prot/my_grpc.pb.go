// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: my.proto

package prot

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

// RepServClient is the client API for RepServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepServClient interface {
	// определение метода с использованием rpc.
	GetData(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Point, error)
}

type repServClient struct {
	cc grpc.ClientConnInterface
}

func NewRepServClient(cc grpc.ClientConnInterface) RepServClient {
	return &repServClient{cc}
}

func (c *repServClient) GetData(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Point, error) {
	out := new(Point)
	err := c.cc.Invoke(ctx, "/RepServ/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepServServer is the server API for RepServ service.
// All implementations must embed UnimplementedRepServServer
// for forward compatibility
type RepServServer interface {
	// определение метода с использованием rpc.
	GetData(context.Context, *Point) (*Point, error)
	mustEmbedUnimplementedRepServServer()
}

// UnimplementedRepServServer must be embedded to have forward compatible implementations.
type UnimplementedRepServServer struct {
}

func (UnimplementedRepServServer) GetData(context.Context, *Point) (*Point, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedRepServServer) mustEmbedUnimplementedRepServServer() {}

// UnsafeRepServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepServServer will
// result in compilation errors.
type UnsafeRepServServer interface {
	mustEmbedUnimplementedRepServServer()
}

func RegisterRepServServer(s grpc.ServiceRegistrar, srv RepServServer) {
	s.RegisterService(&RepServ_ServiceDesc, srv)
}

func _RepServ_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepServServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RepServ/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepServServer).GetData(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

// RepServ_ServiceDesc is the grpc.ServiceDesc for RepServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RepServ",
	HandlerType: (*RepServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _RepServ_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "my.proto",
}
