// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: pal.proto

package pal

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
	PalServer_GetPal_FullMethodName      = "/pal.PalServer/GetPal"
	PalServer_ListPal_FullMethodName     = "/pal.PalServer/ListPal"
	PalServer_GetPalByIds_FullMethodName = "/pal.PalServer/GetPalByIds"
)

// PalServerClient is the client API for PalServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PalServerClient interface {
	GetPal(ctx context.Context, in *GetPalReq, opts ...grpc.CallOption) (*GetPalResp, error)
	ListPal(ctx context.Context, in *ListPalReq, opts ...grpc.CallOption) (*ListPalResp, error)
	GetPalByIds(ctx context.Context, in *GetPalByIdsReq, opts ...grpc.CallOption) (*GetPalByIdsResp, error)
}

type palServerClient struct {
	cc grpc.ClientConnInterface
}

func NewPalServerClient(cc grpc.ClientConnInterface) PalServerClient {
	return &palServerClient{cc}
}

func (c *palServerClient) GetPal(ctx context.Context, in *GetPalReq, opts ...grpc.CallOption) (*GetPalResp, error) {
	out := new(GetPalResp)
	err := c.cc.Invoke(ctx, PalServer_GetPal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *palServerClient) ListPal(ctx context.Context, in *ListPalReq, opts ...grpc.CallOption) (*ListPalResp, error) {
	out := new(ListPalResp)
	err := c.cc.Invoke(ctx, PalServer_ListPal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *palServerClient) GetPalByIds(ctx context.Context, in *GetPalByIdsReq, opts ...grpc.CallOption) (*GetPalByIdsResp, error) {
	out := new(GetPalByIdsResp)
	err := c.cc.Invoke(ctx, PalServer_GetPalByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PalServerServer is the server API for PalServer service.
// All implementations must embed UnimplementedPalServerServer
// for forward compatibility
type PalServerServer interface {
	GetPal(context.Context, *GetPalReq) (*GetPalResp, error)
	ListPal(context.Context, *ListPalReq) (*ListPalResp, error)
	GetPalByIds(context.Context, *GetPalByIdsReq) (*GetPalByIdsResp, error)
	mustEmbedUnimplementedPalServerServer()
}

// UnimplementedPalServerServer must be embedded to have forward compatible implementations.
type UnimplementedPalServerServer struct {
}

func (UnimplementedPalServerServer) GetPal(context.Context, *GetPalReq) (*GetPalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPal not implemented")
}
func (UnimplementedPalServerServer) ListPal(context.Context, *ListPalReq) (*ListPalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPal not implemented")
}
func (UnimplementedPalServerServer) GetPalByIds(context.Context, *GetPalByIdsReq) (*GetPalByIdsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPalByIds not implemented")
}
func (UnimplementedPalServerServer) mustEmbedUnimplementedPalServerServer() {}

// UnsafePalServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PalServerServer will
// result in compilation errors.
type UnsafePalServerServer interface {
	mustEmbedUnimplementedPalServerServer()
}

func RegisterPalServerServer(s grpc.ServiceRegistrar, srv PalServerServer) {
	s.RegisterService(&PalServer_ServiceDesc, srv)
}

func _PalServer_GetPal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PalServerServer).GetPal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PalServer_GetPal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PalServerServer).GetPal(ctx, req.(*GetPalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PalServer_ListPal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PalServerServer).ListPal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PalServer_ListPal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PalServerServer).ListPal(ctx, req.(*ListPalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PalServer_GetPalByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPalByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PalServerServer).GetPalByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PalServer_GetPalByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PalServerServer).GetPalByIds(ctx, req.(*GetPalByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PalServer_ServiceDesc is the grpc.ServiceDesc for PalServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PalServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pal.PalServer",
	HandlerType: (*PalServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPal",
			Handler:    _PalServer_GetPal_Handler,
		},
		{
			MethodName: "ListPal",
			Handler:    _PalServer_ListPal_Handler,
		},
		{
			MethodName: "GetPalByIds",
			Handler:    _PalServer_GetPalByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pal.proto",
}
