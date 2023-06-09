// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: author_api.proto

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

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorServiceClient interface {
	GetAuthorsByIds(ctx context.Context, in *GetAuthorsByIdsRequest, opts ...grpc.CallOption) (*GetAuthorsByIdsResponse, error)
	GetAuthorIdsByEmails(ctx context.Context, in *GetAuthorIdsByEmailsRequest, opts ...grpc.CallOption) (*GetAuthorIdsByEmailResponse, error)
}

type authorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorServiceClient(cc grpc.ClientConnInterface) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) GetAuthorsByIds(ctx context.Context, in *GetAuthorsByIdsRequest, opts ...grpc.CallOption) (*GetAuthorsByIdsResponse, error) {
	out := new(GetAuthorsByIdsResponse)
	err := c.cc.Invoke(ctx, "/author_api.AuthorService/GetAuthorsByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAuthorIdsByEmails(ctx context.Context, in *GetAuthorIdsByEmailsRequest, opts ...grpc.CallOption) (*GetAuthorIdsByEmailResponse, error) {
	out := new(GetAuthorIdsByEmailResponse)
	err := c.cc.Invoke(ctx, "/author_api.AuthorService/GetAuthorIdsByEmails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
// All implementations must embed UnimplementedAuthorServiceServer
// for forward compatibility
type AuthorServiceServer interface {
	GetAuthorsByIds(context.Context, *GetAuthorsByIdsRequest) (*GetAuthorsByIdsResponse, error)
	GetAuthorIdsByEmails(context.Context, *GetAuthorIdsByEmailsRequest) (*GetAuthorIdsByEmailResponse, error)
	mustEmbedUnimplementedAuthorServiceServer()
}

// UnimplementedAuthorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (UnimplementedAuthorServiceServer) GetAuthorsByIds(context.Context, *GetAuthorsByIdsRequest) (*GetAuthorsByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorsByIds not implemented")
}
func (UnimplementedAuthorServiceServer) GetAuthorIdsByEmails(context.Context, *GetAuthorIdsByEmailsRequest) (*GetAuthorIdsByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorIdsByEmails not implemented")
}
func (UnimplementedAuthorServiceServer) mustEmbedUnimplementedAuthorServiceServer() {}

// UnsafeAuthorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServiceServer will
// result in compilation errors.
type UnsafeAuthorServiceServer interface {
	mustEmbedUnimplementedAuthorServiceServer()
}

func RegisterAuthorServiceServer(s grpc.ServiceRegistrar, srv AuthorServiceServer) {
	s.RegisterService(&AuthorService_ServiceDesc, srv)
}

func _AuthorService_GetAuthorsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorsByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthorsByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/author_api.AuthorService/GetAuthorsByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthorsByIds(ctx, req.(*GetAuthorsByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetAuthorIdsByEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorIdsByEmailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthorIdsByEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/author_api.AuthorService/GetAuthorIdsByEmails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthorIdsByEmails(ctx, req.(*GetAuthorIdsByEmailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorService_ServiceDesc is the grpc.ServiceDesc for AuthorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "author_api.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthorsByIds",
			Handler:    _AuthorService_GetAuthorsByIds_Handler,
		},
		{
			MethodName: "GetAuthorIdsByEmails",
			Handler:    _AuthorService_GetAuthorIdsByEmails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author_api.proto",
}
