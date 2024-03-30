// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_user_email_verification.proto

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
	UserEmailVerificationService_VerifyUserEmail_FullMethodName                 = "/pb.UserEmailVerificationService/verifyUserEmail"
	UserEmailVerificationService_SendUserEmailVerification_FullMethodName       = "/pb.UserEmailVerificationService/sendUserEmailVerification"
	UserEmailVerificationService_FindLatestUserEmailVerification_FullMethodName = "/pb.UserEmailVerificationService/findLatestUserEmailVerification"
)

// UserEmailVerificationServiceClient is the client API for UserEmailVerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserEmailVerificationServiceClient interface {
	// 认证邮箱
	VerifyUserEmail(ctx context.Context, in *VerifyUserEmailRequest, opts ...grpc.CallOption) (*VerifyUserEmailResponse, error)
	// 发送邮箱认证
	SendUserEmailVerification(ctx context.Context, in *SendUserEmailVerificationRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找用户正在等待激活的认证
	FindLatestUserEmailVerification(ctx context.Context, in *FindLatestUserEmailVerificationRequest, opts ...grpc.CallOption) (*FindLatestUserEmailVerificationResponse, error)
}

type userEmailVerificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserEmailVerificationServiceClient(cc grpc.ClientConnInterface) UserEmailVerificationServiceClient {
	return &userEmailVerificationServiceClient{cc}
}

func (c *userEmailVerificationServiceClient) VerifyUserEmail(ctx context.Context, in *VerifyUserEmailRequest, opts ...grpc.CallOption) (*VerifyUserEmailResponse, error) {
	out := new(VerifyUserEmailResponse)
	err := c.cc.Invoke(ctx, UserEmailVerificationService_VerifyUserEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEmailVerificationServiceClient) SendUserEmailVerification(ctx context.Context, in *SendUserEmailVerificationRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, UserEmailVerificationService_SendUserEmailVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userEmailVerificationServiceClient) FindLatestUserEmailVerification(ctx context.Context, in *FindLatestUserEmailVerificationRequest, opts ...grpc.CallOption) (*FindLatestUserEmailVerificationResponse, error) {
	out := new(FindLatestUserEmailVerificationResponse)
	err := c.cc.Invoke(ctx, UserEmailVerificationService_FindLatestUserEmailVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserEmailVerificationServiceServer is the server API for UserEmailVerificationService service.
// All implementations should embed UnimplementedUserEmailVerificationServiceServer
// for forward compatibility
type UserEmailVerificationServiceServer interface {
	// 认证邮箱
	VerifyUserEmail(context.Context, *VerifyUserEmailRequest) (*VerifyUserEmailResponse, error)
	// 发送邮箱认证
	SendUserEmailVerification(context.Context, *SendUserEmailVerificationRequest) (*RPCSuccess, error)
	// 查找用户正在等待激活的认证
	FindLatestUserEmailVerification(context.Context, *FindLatestUserEmailVerificationRequest) (*FindLatestUserEmailVerificationResponse, error)
}

// UnimplementedUserEmailVerificationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserEmailVerificationServiceServer struct {
}

func (UnimplementedUserEmailVerificationServiceServer) VerifyUserEmail(context.Context, *VerifyUserEmailRequest) (*VerifyUserEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserEmail not implemented")
}
func (UnimplementedUserEmailVerificationServiceServer) SendUserEmailVerification(context.Context, *SendUserEmailVerificationRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserEmailVerification not implemented")
}
func (UnimplementedUserEmailVerificationServiceServer) FindLatestUserEmailVerification(context.Context, *FindLatestUserEmailVerificationRequest) (*FindLatestUserEmailVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLatestUserEmailVerification not implemented")
}

// UnsafeUserEmailVerificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserEmailVerificationServiceServer will
// result in compilation errors.
type UnsafeUserEmailVerificationServiceServer interface {
	mustEmbedUnimplementedUserEmailVerificationServiceServer()
}

func RegisterUserEmailVerificationServiceServer(s grpc.ServiceRegistrar, srv UserEmailVerificationServiceServer) {
	s.RegisterService(&UserEmailVerificationService_ServiceDesc, srv)
}

func _UserEmailVerificationService_VerifyUserEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEmailVerificationServiceServer).VerifyUserEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserEmailVerificationService_VerifyUserEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEmailVerificationServiceServer).VerifyUserEmail(ctx, req.(*VerifyUserEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEmailVerificationService_SendUserEmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserEmailVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEmailVerificationServiceServer).SendUserEmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserEmailVerificationService_SendUserEmailVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEmailVerificationServiceServer).SendUserEmailVerification(ctx, req.(*SendUserEmailVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserEmailVerificationService_FindLatestUserEmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindLatestUserEmailVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserEmailVerificationServiceServer).FindLatestUserEmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserEmailVerificationService_FindLatestUserEmailVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserEmailVerificationServiceServer).FindLatestUserEmailVerification(ctx, req.(*FindLatestUserEmailVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserEmailVerificationService_ServiceDesc is the grpc.ServiceDesc for UserEmailVerificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserEmailVerificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserEmailVerificationService",
	HandlerType: (*UserEmailVerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "verifyUserEmail",
			Handler:    _UserEmailVerificationService_VerifyUserEmail_Handler,
		},
		{
			MethodName: "sendUserEmailVerification",
			Handler:    _UserEmailVerificationService_SendUserEmailVerification_Handler,
		},
		{
			MethodName: "findLatestUserEmailVerification",
			Handler:    _UserEmailVerificationService_FindLatestUserEmailVerification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_email_verification.proto",
}