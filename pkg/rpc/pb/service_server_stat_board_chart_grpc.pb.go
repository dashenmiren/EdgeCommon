// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_server_stat_board_chart.proto

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
	ServerStatBoardChartService_EnableServerStatBoardChart_FullMethodName          = "/pb.ServerStatBoardChartService/enableServerStatBoardChart"
	ServerStatBoardChartService_DisableServerStatBoardChart_FullMethodName         = "/pb.ServerStatBoardChartService/disableServerStatBoardChart"
	ServerStatBoardChartService_FindAllEnabledServerStatBoardCharts_FullMethodName = "/pb.ServerStatBoardChartService/findAllEnabledServerStatBoardCharts"
)

// ServerStatBoardChartServiceClient is the client API for ServerStatBoardChartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerStatBoardChartServiceClient interface {
	// 添加图表
	EnableServerStatBoardChart(ctx context.Context, in *EnableServerStatBoardChartRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 取消图表
	DisableServerStatBoardChart(ctx context.Context, in *DisableServerStatBoardChartRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 读取看板中的图表
	FindAllEnabledServerStatBoardCharts(ctx context.Context, in *FindAllEnabledServerStatBoardChartsRequest, opts ...grpc.CallOption) (*FindAllEnabledServerStatBoardChartsResponse, error)
}

type serverStatBoardChartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStatBoardChartServiceClient(cc grpc.ClientConnInterface) ServerStatBoardChartServiceClient {
	return &serverStatBoardChartServiceClient{cc}
}

func (c *serverStatBoardChartServiceClient) EnableServerStatBoardChart(ctx context.Context, in *EnableServerStatBoardChartRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, ServerStatBoardChartService_EnableServerStatBoardChart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverStatBoardChartServiceClient) DisableServerStatBoardChart(ctx context.Context, in *DisableServerStatBoardChartRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, ServerStatBoardChartService_DisableServerStatBoardChart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverStatBoardChartServiceClient) FindAllEnabledServerStatBoardCharts(ctx context.Context, in *FindAllEnabledServerStatBoardChartsRequest, opts ...grpc.CallOption) (*FindAllEnabledServerStatBoardChartsResponse, error) {
	out := new(FindAllEnabledServerStatBoardChartsResponse)
	err := c.cc.Invoke(ctx, ServerStatBoardChartService_FindAllEnabledServerStatBoardCharts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerStatBoardChartServiceServer is the server API for ServerStatBoardChartService service.
// All implementations should embed UnimplementedServerStatBoardChartServiceServer
// for forward compatibility
type ServerStatBoardChartServiceServer interface {
	// 添加图表
	EnableServerStatBoardChart(context.Context, *EnableServerStatBoardChartRequest) (*RPCSuccess, error)
	// 取消图表
	DisableServerStatBoardChart(context.Context, *DisableServerStatBoardChartRequest) (*RPCSuccess, error)
	// 读取看板中的图表
	FindAllEnabledServerStatBoardCharts(context.Context, *FindAllEnabledServerStatBoardChartsRequest) (*FindAllEnabledServerStatBoardChartsResponse, error)
}

// UnimplementedServerStatBoardChartServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServerStatBoardChartServiceServer struct {
}

func (UnimplementedServerStatBoardChartServiceServer) EnableServerStatBoardChart(context.Context, *EnableServerStatBoardChartRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableServerStatBoardChart not implemented")
}
func (UnimplementedServerStatBoardChartServiceServer) DisableServerStatBoardChart(context.Context, *DisableServerStatBoardChartRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableServerStatBoardChart not implemented")
}
func (UnimplementedServerStatBoardChartServiceServer) FindAllEnabledServerStatBoardCharts(context.Context, *FindAllEnabledServerStatBoardChartsRequest) (*FindAllEnabledServerStatBoardChartsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledServerStatBoardCharts not implemented")
}

// UnsafeServerStatBoardChartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerStatBoardChartServiceServer will
// result in compilation errors.
type UnsafeServerStatBoardChartServiceServer interface {
	mustEmbedUnimplementedServerStatBoardChartServiceServer()
}

func RegisterServerStatBoardChartServiceServer(s grpc.ServiceRegistrar, srv ServerStatBoardChartServiceServer) {
	s.RegisterService(&ServerStatBoardChartService_ServiceDesc, srv)
}

func _ServerStatBoardChartService_EnableServerStatBoardChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableServerStatBoardChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatBoardChartServiceServer).EnableServerStatBoardChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerStatBoardChartService_EnableServerStatBoardChart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatBoardChartServiceServer).EnableServerStatBoardChart(ctx, req.(*EnableServerStatBoardChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerStatBoardChartService_DisableServerStatBoardChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableServerStatBoardChartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatBoardChartServiceServer).DisableServerStatBoardChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerStatBoardChartService_DisableServerStatBoardChart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatBoardChartServiceServer).DisableServerStatBoardChart(ctx, req.(*DisableServerStatBoardChartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerStatBoardChartService_FindAllEnabledServerStatBoardCharts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledServerStatBoardChartsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatBoardChartServiceServer).FindAllEnabledServerStatBoardCharts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerStatBoardChartService_FindAllEnabledServerStatBoardCharts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatBoardChartServiceServer).FindAllEnabledServerStatBoardCharts(ctx, req.(*FindAllEnabledServerStatBoardChartsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerStatBoardChartService_ServiceDesc is the grpc.ServiceDesc for ServerStatBoardChartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerStatBoardChartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerStatBoardChartService",
	HandlerType: (*ServerStatBoardChartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "enableServerStatBoardChart",
			Handler:    _ServerStatBoardChartService_EnableServerStatBoardChart_Handler,
		},
		{
			MethodName: "disableServerStatBoardChart",
			Handler:    _ServerStatBoardChartService_DisableServerStatBoardChart_Handler,
		},
		{
			MethodName: "findAllEnabledServerStatBoardCharts",
			Handler:    _ServerStatBoardChartService_FindAllEnabledServerStatBoardCharts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_server_stat_board_chart.proto",
}
