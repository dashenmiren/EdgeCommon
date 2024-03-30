// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_api_token.proto

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

const (
	APITokenService_FindAllEnabledAPITokens_FullMethodName = "/pb.APITokenService/findAllEnabledAPITokens"
)

// APITokenServiceClient is the client API for APITokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APITokenServiceClient interface {
	// 获取API令牌
	FindAllEnabledAPITokens(ctx context.Context, in *FindAllEnabledAPITokensRequest, opts ...grpc.CallOption) (*FindAllEnabledAPITokensResponse, error)
}

type aPITokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPITokenServiceClient(cc grpc.ClientConnInterface) APITokenServiceClient {
	return &aPITokenServiceClient{cc}
}

func (c *aPITokenServiceClient) FindAllEnabledAPITokens(ctx context.Context, in *FindAllEnabledAPITokensRequest, opts ...grpc.CallOption) (*FindAllEnabledAPITokensResponse, error) {
	out := new(FindAllEnabledAPITokensResponse)
	err := c.cc.Invoke(ctx, APITokenService_FindAllEnabledAPITokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APITokenServiceServer is the server API for APITokenService service.
// All implementations should embed UnimplementedAPITokenServiceServer
// for forward compatibility
type APITokenServiceServer interface {
	// 获取API令牌
	FindAllEnabledAPITokens(context.Context, *FindAllEnabledAPITokensRequest) (*FindAllEnabledAPITokensResponse, error)
}

// UnimplementedAPITokenServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAPITokenServiceServer struct {
}

func (UnimplementedAPITokenServiceServer) FindAllEnabledAPITokens(context.Context, *FindAllEnabledAPITokensRequest) (*FindAllEnabledAPITokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledAPITokens not implemented")
}

// UnsafeAPITokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APITokenServiceServer will
// result in compilation errors.
type UnsafeAPITokenServiceServer interface {
	mustEmbedUnimplementedAPITokenServiceServer()
}

func RegisterAPITokenServiceServer(s grpc.ServiceRegistrar, srv APITokenServiceServer) {
	s.RegisterService(&APITokenService_ServiceDesc, srv)
}

func _APITokenService_FindAllEnabledAPITokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledAPITokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APITokenServiceServer).FindAllEnabledAPITokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APITokenService_FindAllEnabledAPITokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APITokenServiceServer).FindAllEnabledAPITokens(ctx, req.(*FindAllEnabledAPITokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// APITokenService_ServiceDesc is the grpc.ServiceDesc for APITokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APITokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.APITokenService",
	HandlerType: (*APITokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllEnabledAPITokens",
			Handler:    _APITokenService_FindAllEnabledAPITokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_api_token.proto",
}