// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_http_access_log.proto

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
	HTTPAccessLogService_CreateHTTPAccessLogs_FullMethodName        = "/pb.HTTPAccessLogService/createHTTPAccessLogs"
	HTTPAccessLogService_ListHTTPAccessLogs_FullMethodName          = "/pb.HTTPAccessLogService/listHTTPAccessLogs"
	HTTPAccessLogService_FindHTTPAccessLog_FullMethodName           = "/pb.HTTPAccessLogService/findHTTPAccessLog"
	HTTPAccessLogService_FindHTTPAccessLogPartitions_FullMethodName = "/pb.HTTPAccessLogService/findHTTPAccessLogPartitions"
)

// HTTPAccessLogServiceClient is the client API for HTTPAccessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPAccessLogServiceClient interface {
	// 创建访问日志
	CreateHTTPAccessLogs(ctx context.Context, in *CreateHTTPAccessLogsRequest, opts ...grpc.CallOption) (*CreateHTTPAccessLogsResponse, error)
	// 列出单页访问日志
	ListHTTPAccessLogs(ctx context.Context, in *ListHTTPAccessLogsRequest, opts ...grpc.CallOption) (*ListHTTPAccessLogsResponse, error)
	// 查找单个日志
	FindHTTPAccessLog(ctx context.Context, in *FindHTTPAccessLogRequest, opts ...grpc.CallOption) (*FindHTTPAccessLogResponse, error)
	// 查找日志分区
	FindHTTPAccessLogPartitions(ctx context.Context, in *FindHTTPAccessLogPartitionsRequest, opts ...grpc.CallOption) (*FindHTTPAccessLogPartitionsResponse, error)
}

type hTTPAccessLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPAccessLogServiceClient(cc grpc.ClientConnInterface) HTTPAccessLogServiceClient {
	return &hTTPAccessLogServiceClient{cc}
}

func (c *hTTPAccessLogServiceClient) CreateHTTPAccessLogs(ctx context.Context, in *CreateHTTPAccessLogsRequest, opts ...grpc.CallOption) (*CreateHTTPAccessLogsResponse, error) {
	out := new(CreateHTTPAccessLogsResponse)
	err := c.cc.Invoke(ctx, HTTPAccessLogService_CreateHTTPAccessLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPAccessLogServiceClient) ListHTTPAccessLogs(ctx context.Context, in *ListHTTPAccessLogsRequest, opts ...grpc.CallOption) (*ListHTTPAccessLogsResponse, error) {
	out := new(ListHTTPAccessLogsResponse)
	err := c.cc.Invoke(ctx, HTTPAccessLogService_ListHTTPAccessLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPAccessLogServiceClient) FindHTTPAccessLog(ctx context.Context, in *FindHTTPAccessLogRequest, opts ...grpc.CallOption) (*FindHTTPAccessLogResponse, error) {
	out := new(FindHTTPAccessLogResponse)
	err := c.cc.Invoke(ctx, HTTPAccessLogService_FindHTTPAccessLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPAccessLogServiceClient) FindHTTPAccessLogPartitions(ctx context.Context, in *FindHTTPAccessLogPartitionsRequest, opts ...grpc.CallOption) (*FindHTTPAccessLogPartitionsResponse, error) {
	out := new(FindHTTPAccessLogPartitionsResponse)
	err := c.cc.Invoke(ctx, HTTPAccessLogService_FindHTTPAccessLogPartitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPAccessLogServiceServer is the server API for HTTPAccessLogService service.
// All implementations should embed UnimplementedHTTPAccessLogServiceServer
// for forward compatibility
type HTTPAccessLogServiceServer interface {
	// 创建访问日志
	CreateHTTPAccessLogs(context.Context, *CreateHTTPAccessLogsRequest) (*CreateHTTPAccessLogsResponse, error)
	// 列出单页访问日志
	ListHTTPAccessLogs(context.Context, *ListHTTPAccessLogsRequest) (*ListHTTPAccessLogsResponse, error)
	// 查找单个日志
	FindHTTPAccessLog(context.Context, *FindHTTPAccessLogRequest) (*FindHTTPAccessLogResponse, error)
	// 查找日志分区
	FindHTTPAccessLogPartitions(context.Context, *FindHTTPAccessLogPartitionsRequest) (*FindHTTPAccessLogPartitionsResponse, error)
}

// UnimplementedHTTPAccessLogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHTTPAccessLogServiceServer struct {
}

func (UnimplementedHTTPAccessLogServiceServer) CreateHTTPAccessLogs(context.Context, *CreateHTTPAccessLogsRequest) (*CreateHTTPAccessLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHTTPAccessLogs not implemented")
}
func (UnimplementedHTTPAccessLogServiceServer) ListHTTPAccessLogs(context.Context, *ListHTTPAccessLogsRequest) (*ListHTTPAccessLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHTTPAccessLogs not implemented")
}
func (UnimplementedHTTPAccessLogServiceServer) FindHTTPAccessLog(context.Context, *FindHTTPAccessLogRequest) (*FindHTTPAccessLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindHTTPAccessLog not implemented")
}
func (UnimplementedHTTPAccessLogServiceServer) FindHTTPAccessLogPartitions(context.Context, *FindHTTPAccessLogPartitionsRequest) (*FindHTTPAccessLogPartitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindHTTPAccessLogPartitions not implemented")
}

// UnsafeHTTPAccessLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPAccessLogServiceServer will
// result in compilation errors.
type UnsafeHTTPAccessLogServiceServer interface {
	mustEmbedUnimplementedHTTPAccessLogServiceServer()
}

func RegisterHTTPAccessLogServiceServer(s grpc.ServiceRegistrar, srv HTTPAccessLogServiceServer) {
	s.RegisterService(&HTTPAccessLogService_ServiceDesc, srv)
}

func _HTTPAccessLogService_CreateHTTPAccessLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHTTPAccessLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAccessLogServiceServer).CreateHTTPAccessLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPAccessLogService_CreateHTTPAccessLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAccessLogServiceServer).CreateHTTPAccessLogs(ctx, req.(*CreateHTTPAccessLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPAccessLogService_ListHTTPAccessLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHTTPAccessLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAccessLogServiceServer).ListHTTPAccessLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPAccessLogService_ListHTTPAccessLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAccessLogServiceServer).ListHTTPAccessLogs(ctx, req.(*ListHTTPAccessLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPAccessLogService_FindHTTPAccessLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindHTTPAccessLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAccessLogServiceServer).FindHTTPAccessLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPAccessLogService_FindHTTPAccessLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAccessLogServiceServer).FindHTTPAccessLog(ctx, req.(*FindHTTPAccessLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPAccessLogService_FindHTTPAccessLogPartitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindHTTPAccessLogPartitionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPAccessLogServiceServer).FindHTTPAccessLogPartitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPAccessLogService_FindHTTPAccessLogPartitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPAccessLogServiceServer).FindHTTPAccessLogPartitions(ctx, req.(*FindHTTPAccessLogPartitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPAccessLogService_ServiceDesc is the grpc.ServiceDesc for HTTPAccessLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPAccessLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPAccessLogService",
	HandlerType: (*HTTPAccessLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createHTTPAccessLogs",
			Handler:    _HTTPAccessLogService_CreateHTTPAccessLogs_Handler,
		},
		{
			MethodName: "listHTTPAccessLogs",
			Handler:    _HTTPAccessLogService_ListHTTPAccessLogs_Handler,
		},
		{
			MethodName: "findHTTPAccessLog",
			Handler:    _HTTPAccessLogService_FindHTTPAccessLog_Handler,
		},
		{
			MethodName: "findHTTPAccessLogPartitions",
			Handler:    _HTTPAccessLogService_FindHTTPAccessLogPartitions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_access_log.proto",
}