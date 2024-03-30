// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_user_traffic_bill.proto

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
	UserTrafficBillService_FindUserTrafficBills_FullMethodName = "/pb.UserTrafficBillService/findUserTrafficBills"
)

// UserTrafficBillServiceClient is the client API for UserTrafficBillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserTrafficBillServiceClient interface {
	// 列出某个账单下的流量带宽子账单
	FindUserTrafficBills(ctx context.Context, in *FindUserTrafficBillsRequest, opts ...grpc.CallOption) (*FindUserTrafficBillsResponse, error)
}

type userTrafficBillServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserTrafficBillServiceClient(cc grpc.ClientConnInterface) UserTrafficBillServiceClient {
	return &userTrafficBillServiceClient{cc}
}

func (c *userTrafficBillServiceClient) FindUserTrafficBills(ctx context.Context, in *FindUserTrafficBillsRequest, opts ...grpc.CallOption) (*FindUserTrafficBillsResponse, error) {
	out := new(FindUserTrafficBillsResponse)
	err := c.cc.Invoke(ctx, UserTrafficBillService_FindUserTrafficBills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserTrafficBillServiceServer is the server API for UserTrafficBillService service.
// All implementations should embed UnimplementedUserTrafficBillServiceServer
// for forward compatibility
type UserTrafficBillServiceServer interface {
	// 列出某个账单下的流量带宽子账单
	FindUserTrafficBills(context.Context, *FindUserTrafficBillsRequest) (*FindUserTrafficBillsResponse, error)
}

// UnimplementedUserTrafficBillServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserTrafficBillServiceServer struct {
}

func (UnimplementedUserTrafficBillServiceServer) FindUserTrafficBills(context.Context, *FindUserTrafficBillsRequest) (*FindUserTrafficBillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserTrafficBills not implemented")
}

// UnsafeUserTrafficBillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserTrafficBillServiceServer will
// result in compilation errors.
type UnsafeUserTrafficBillServiceServer interface {
	mustEmbedUnimplementedUserTrafficBillServiceServer()
}

func RegisterUserTrafficBillServiceServer(s grpc.ServiceRegistrar, srv UserTrafficBillServiceServer) {
	s.RegisterService(&UserTrafficBillService_ServiceDesc, srv)
}

func _UserTrafficBillService_FindUserTrafficBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserTrafficBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTrafficBillServiceServer).FindUserTrafficBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserTrafficBillService_FindUserTrafficBills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTrafficBillServiceServer).FindUserTrafficBills(ctx, req.(*FindUserTrafficBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserTrafficBillService_ServiceDesc is the grpc.ServiceDesc for UserTrafficBillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserTrafficBillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserTrafficBillService",
	HandlerType: (*UserTrafficBillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findUserTrafficBills",
			Handler:    _UserTrafficBillService_FindUserTrafficBills_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_traffic_bill.proto",
}
