// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: comments.proto

package comments

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
	CommentsServer_AddComment_FullMethodName    = "/comments.CommentsServer/AddComment"
	CommentsServer_ListComment_FullMethodName   = "/comments.CommentsServer/ListComment"
	CommentsServer_DeleteComment_FullMethodName = "/comments.CommentsServer/DeleteComment"
)

// CommentsServerClient is the client API for CommentsServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentsServerClient interface {
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
	ListComment(ctx context.Context, in *ListCommentReq, opts ...grpc.CallOption) (*CommentListResp, error)
	DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error)
}

type commentsServerClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentsServerClient(cc grpc.ClientConnInterface) CommentsServerClient {
	return &commentsServerClient{cc}
}

func (c *commentsServerClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, CommentsServer_AddComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsServerClient) ListComment(ctx context.Context, in *ListCommentReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	out := new(CommentListResp)
	err := c.cc.Invoke(ctx, CommentsServer_ListComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsServerClient) DeleteComment(ctx context.Context, in *DeleteCommentReq, opts ...grpc.CallOption) (*DeleteCommentResp, error) {
	out := new(DeleteCommentResp)
	err := c.cc.Invoke(ctx, CommentsServer_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentsServerServer is the server API for CommentsServer service.
// All implementations must embed UnimplementedCommentsServerServer
// for forward compatibility
type CommentsServerServer interface {
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
	ListComment(context.Context, *ListCommentReq) (*CommentListResp, error)
	DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentResp, error)
	mustEmbedUnimplementedCommentsServerServer()
}

// UnimplementedCommentsServerServer must be embedded to have forward compatible implementations.
type UnimplementedCommentsServerServer struct {
}

func (UnimplementedCommentsServerServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedCommentsServerServer) ListComment(context.Context, *ListCommentReq) (*CommentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComment not implemented")
}
func (UnimplementedCommentsServerServer) DeleteComment(context.Context, *DeleteCommentReq) (*DeleteCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentsServerServer) mustEmbedUnimplementedCommentsServerServer() {}

// UnsafeCommentsServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentsServerServer will
// result in compilation errors.
type UnsafeCommentsServerServer interface {
	mustEmbedUnimplementedCommentsServerServer()
}

func RegisterCommentsServerServer(s grpc.ServiceRegistrar, srv CommentsServerServer) {
	s.RegisterService(&CommentsServer_ServiceDesc, srv)
}

func _CommentsServer_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServerServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsServer_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServerServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsServer_ListComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServerServer).ListComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsServer_ListComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServerServer).ListComment(ctx, req.(*ListCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentsServer_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentsServerServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentsServer_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentsServerServer).DeleteComment(ctx, req.(*DeleteCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentsServer_ServiceDesc is the grpc.ServiceDesc for CommentsServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentsServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comments.CommentsServer",
	HandlerType: (*CommentsServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddComment",
			Handler:    _CommentsServer_AddComment_Handler,
		},
		{
			MethodName: "ListComment",
			Handler:    _CommentsServer_ListComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _CommentsServer_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comments.proto",
}
