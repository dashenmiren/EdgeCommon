// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_ip_library_file.proto

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
	IPLibraryFileService_FindAllFinishedIPLibraryFiles_FullMethodName     = "/pb.IPLibraryFileService/findAllFinishedIPLibraryFiles"
	IPLibraryFileService_FindAllUnfinishedIPLibraryFiles_FullMethodName   = "/pb.IPLibraryFileService/findAllUnfinishedIPLibraryFiles"
	IPLibraryFileService_FindIPLibraryFile_FullMethodName                 = "/pb.IPLibraryFileService/findIPLibraryFile"
	IPLibraryFileService_CreateIPLibraryFile_FullMethodName               = "/pb.IPLibraryFileService/createIPLibraryFile"
	IPLibraryFileService_CheckCountriesWithIPLibraryFileId_FullMethodName = "/pb.IPLibraryFileService/checkCountriesWithIPLibraryFileId"
	IPLibraryFileService_CheckProvincesWithIPLibraryFileId_FullMethodName = "/pb.IPLibraryFileService/checkProvincesWithIPLibraryFileId"
	IPLibraryFileService_CheckCitiesWithIPLibraryFileId_FullMethodName    = "/pb.IPLibraryFileService/checkCitiesWithIPLibraryFileId"
	IPLibraryFileService_CheckTownsWithIPLibraryFileId_FullMethodName     = "/pb.IPLibraryFileService/checkTownsWithIPLibraryFileId"
	IPLibraryFileService_CheckProvidersWithIPLibraryFileId_FullMethodName = "/pb.IPLibraryFileService/checkProvidersWithIPLibraryFileId"
	IPLibraryFileService_GenerateIPLibraryFile_FullMethodName             = "/pb.IPLibraryFileService/generateIPLibraryFile"
	IPLibraryFileService_UpdateIPLibraryFileFinished_FullMethodName       = "/pb.IPLibraryFileService/updateIPLibraryFileFinished"
	IPLibraryFileService_DeleteIPLibraryFile_FullMethodName               = "/pb.IPLibraryFileService/deleteIPLibraryFile"
)

// IPLibraryFileServiceClient is the client API for IPLibraryFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IPLibraryFileServiceClient interface {
	// 查找所有已完成的IP库文件
	FindAllFinishedIPLibraryFiles(ctx context.Context, in *FindAllFinishedIPLibraryFilesRequest, opts ...grpc.CallOption) (*FindAllFinishedIPLibraryFilesResponse, error)
	// 查找所有未完成的IP库文件
	FindAllUnfinishedIPLibraryFiles(ctx context.Context, in *FindAllUnfinishedIPLibraryFilesRequest, opts ...grpc.CallOption) (*FindAllUnfinishedIPLibraryFilesResponse, error)
	// 查找单个IP库文件
	FindIPLibraryFile(ctx context.Context, in *FindIPLibraryFileRequest, opts ...grpc.CallOption) (*FindIPLibraryFileResponse, error)
	// 创建IP库文件
	CreateIPLibraryFile(ctx context.Context, in *CreateIPLibraryFileRequest, opts ...grpc.CallOption) (*CreateIPLibraryFileResponse, error)
	// 检查国家/地区
	CheckCountriesWithIPLibraryFileId(ctx context.Context, in *CheckCountriesWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckCountriesWithIPLibraryFileIdResponse, error)
	// 检查省份/州
	CheckProvincesWithIPLibraryFileId(ctx context.Context, in *CheckProvincesWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckProvincesWithIPLibraryFileIdResponse, error)
	// 检查城市/市
	CheckCitiesWithIPLibraryFileId(ctx context.Context, in *CheckCitiesWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckCitiesWithIPLibraryFileIdResponse, error)
	// 检查区县
	CheckTownsWithIPLibraryFileId(ctx context.Context, in *CheckTownsWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckTownsWithIPLibraryFileIdResponse, error)
	// 检查ISP运营商
	CheckProvidersWithIPLibraryFileId(ctx context.Context, in *CheckProvidersWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckProvidersWithIPLibraryFileIdResponse, error)
	// 生成IP库文件
	GenerateIPLibraryFile(ctx context.Context, in *GenerateIPLibraryFileRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 设置某个IP库为已完成
	UpdateIPLibraryFileFinished(ctx context.Context, in *UpdateIPLibraryFileFinishedRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 删除IP库文件
	DeleteIPLibraryFile(ctx context.Context, in *DeleteIPLibraryFileRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type iPLibraryFileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIPLibraryFileServiceClient(cc grpc.ClientConnInterface) IPLibraryFileServiceClient {
	return &iPLibraryFileServiceClient{cc}
}

func (c *iPLibraryFileServiceClient) FindAllFinishedIPLibraryFiles(ctx context.Context, in *FindAllFinishedIPLibraryFilesRequest, opts ...grpc.CallOption) (*FindAllFinishedIPLibraryFilesResponse, error) {
	out := new(FindAllFinishedIPLibraryFilesResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_FindAllFinishedIPLibraryFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) FindAllUnfinishedIPLibraryFiles(ctx context.Context, in *FindAllUnfinishedIPLibraryFilesRequest, opts ...grpc.CallOption) (*FindAllUnfinishedIPLibraryFilesResponse, error) {
	out := new(FindAllUnfinishedIPLibraryFilesResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_FindAllUnfinishedIPLibraryFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) FindIPLibraryFile(ctx context.Context, in *FindIPLibraryFileRequest, opts ...grpc.CallOption) (*FindIPLibraryFileResponse, error) {
	out := new(FindIPLibraryFileResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_FindIPLibraryFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) CreateIPLibraryFile(ctx context.Context, in *CreateIPLibraryFileRequest, opts ...grpc.CallOption) (*CreateIPLibraryFileResponse, error) {
	out := new(CreateIPLibraryFileResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_CreateIPLibraryFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) CheckCountriesWithIPLibraryFileId(ctx context.Context, in *CheckCountriesWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckCountriesWithIPLibraryFileIdResponse, error) {
	out := new(CheckCountriesWithIPLibraryFileIdResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_CheckCountriesWithIPLibraryFileId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) CheckProvincesWithIPLibraryFileId(ctx context.Context, in *CheckProvincesWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckProvincesWithIPLibraryFileIdResponse, error) {
	out := new(CheckProvincesWithIPLibraryFileIdResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_CheckProvincesWithIPLibraryFileId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) CheckCitiesWithIPLibraryFileId(ctx context.Context, in *CheckCitiesWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckCitiesWithIPLibraryFileIdResponse, error) {
	out := new(CheckCitiesWithIPLibraryFileIdResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_CheckCitiesWithIPLibraryFileId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) CheckTownsWithIPLibraryFileId(ctx context.Context, in *CheckTownsWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckTownsWithIPLibraryFileIdResponse, error) {
	out := new(CheckTownsWithIPLibraryFileIdResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_CheckTownsWithIPLibraryFileId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) CheckProvidersWithIPLibraryFileId(ctx context.Context, in *CheckProvidersWithIPLibraryFileIdRequest, opts ...grpc.CallOption) (*CheckProvidersWithIPLibraryFileIdResponse, error) {
	out := new(CheckProvidersWithIPLibraryFileIdResponse)
	err := c.cc.Invoke(ctx, IPLibraryFileService_CheckProvidersWithIPLibraryFileId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) GenerateIPLibraryFile(ctx context.Context, in *GenerateIPLibraryFileRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, IPLibraryFileService_GenerateIPLibraryFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) UpdateIPLibraryFileFinished(ctx context.Context, in *UpdateIPLibraryFileFinishedRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, IPLibraryFileService_UpdateIPLibraryFileFinished_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPLibraryFileServiceClient) DeleteIPLibraryFile(ctx context.Context, in *DeleteIPLibraryFileRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, IPLibraryFileService_DeleteIPLibraryFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPLibraryFileServiceServer is the server API for IPLibraryFileService service.
// All implementations should embed UnimplementedIPLibraryFileServiceServer
// for forward compatibility
type IPLibraryFileServiceServer interface {
	// 查找所有已完成的IP库文件
	FindAllFinishedIPLibraryFiles(context.Context, *FindAllFinishedIPLibraryFilesRequest) (*FindAllFinishedIPLibraryFilesResponse, error)
	// 查找所有未完成的IP库文件
	FindAllUnfinishedIPLibraryFiles(context.Context, *FindAllUnfinishedIPLibraryFilesRequest) (*FindAllUnfinishedIPLibraryFilesResponse, error)
	// 查找单个IP库文件
	FindIPLibraryFile(context.Context, *FindIPLibraryFileRequest) (*FindIPLibraryFileResponse, error)
	// 创建IP库文件
	CreateIPLibraryFile(context.Context, *CreateIPLibraryFileRequest) (*CreateIPLibraryFileResponse, error)
	// 检查国家/地区
	CheckCountriesWithIPLibraryFileId(context.Context, *CheckCountriesWithIPLibraryFileIdRequest) (*CheckCountriesWithIPLibraryFileIdResponse, error)
	// 检查省份/州
	CheckProvincesWithIPLibraryFileId(context.Context, *CheckProvincesWithIPLibraryFileIdRequest) (*CheckProvincesWithIPLibraryFileIdResponse, error)
	// 检查城市/市
	CheckCitiesWithIPLibraryFileId(context.Context, *CheckCitiesWithIPLibraryFileIdRequest) (*CheckCitiesWithIPLibraryFileIdResponse, error)
	// 检查区县
	CheckTownsWithIPLibraryFileId(context.Context, *CheckTownsWithIPLibraryFileIdRequest) (*CheckTownsWithIPLibraryFileIdResponse, error)
	// 检查ISP运营商
	CheckProvidersWithIPLibraryFileId(context.Context, *CheckProvidersWithIPLibraryFileIdRequest) (*CheckProvidersWithIPLibraryFileIdResponse, error)
	// 生成IP库文件
	GenerateIPLibraryFile(context.Context, *GenerateIPLibraryFileRequest) (*RPCSuccess, error)
	// 设置某个IP库为已完成
	UpdateIPLibraryFileFinished(context.Context, *UpdateIPLibraryFileFinishedRequest) (*RPCSuccess, error)
	// 删除IP库文件
	DeleteIPLibraryFile(context.Context, *DeleteIPLibraryFileRequest) (*RPCSuccess, error)
}

// UnimplementedIPLibraryFileServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIPLibraryFileServiceServer struct {
}

func (UnimplementedIPLibraryFileServiceServer) FindAllFinishedIPLibraryFiles(context.Context, *FindAllFinishedIPLibraryFilesRequest) (*FindAllFinishedIPLibraryFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllFinishedIPLibraryFiles not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) FindAllUnfinishedIPLibraryFiles(context.Context, *FindAllUnfinishedIPLibraryFilesRequest) (*FindAllUnfinishedIPLibraryFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllUnfinishedIPLibraryFiles not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) FindIPLibraryFile(context.Context, *FindIPLibraryFileRequest) (*FindIPLibraryFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindIPLibraryFile not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) CreateIPLibraryFile(context.Context, *CreateIPLibraryFileRequest) (*CreateIPLibraryFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIPLibraryFile not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) CheckCountriesWithIPLibraryFileId(context.Context, *CheckCountriesWithIPLibraryFileIdRequest) (*CheckCountriesWithIPLibraryFileIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCountriesWithIPLibraryFileId not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) CheckProvincesWithIPLibraryFileId(context.Context, *CheckProvincesWithIPLibraryFileIdRequest) (*CheckProvincesWithIPLibraryFileIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProvincesWithIPLibraryFileId not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) CheckCitiesWithIPLibraryFileId(context.Context, *CheckCitiesWithIPLibraryFileIdRequest) (*CheckCitiesWithIPLibraryFileIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCitiesWithIPLibraryFileId not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) CheckTownsWithIPLibraryFileId(context.Context, *CheckTownsWithIPLibraryFileIdRequest) (*CheckTownsWithIPLibraryFileIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTownsWithIPLibraryFileId not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) CheckProvidersWithIPLibraryFileId(context.Context, *CheckProvidersWithIPLibraryFileIdRequest) (*CheckProvidersWithIPLibraryFileIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProvidersWithIPLibraryFileId not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) GenerateIPLibraryFile(context.Context, *GenerateIPLibraryFileRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateIPLibraryFile not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) UpdateIPLibraryFileFinished(context.Context, *UpdateIPLibraryFileFinishedRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIPLibraryFileFinished not implemented")
}
func (UnimplementedIPLibraryFileServiceServer) DeleteIPLibraryFile(context.Context, *DeleteIPLibraryFileRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIPLibraryFile not implemented")
}

// UnsafeIPLibraryFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IPLibraryFileServiceServer will
// result in compilation errors.
type UnsafeIPLibraryFileServiceServer interface {
	mustEmbedUnimplementedIPLibraryFileServiceServer()
}

func RegisterIPLibraryFileServiceServer(s grpc.ServiceRegistrar, srv IPLibraryFileServiceServer) {
	s.RegisterService(&IPLibraryFileService_ServiceDesc, srv)
}

func _IPLibraryFileService_FindAllFinishedIPLibraryFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllFinishedIPLibraryFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).FindAllFinishedIPLibraryFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_FindAllFinishedIPLibraryFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).FindAllFinishedIPLibraryFiles(ctx, req.(*FindAllFinishedIPLibraryFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_FindAllUnfinishedIPLibraryFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllUnfinishedIPLibraryFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).FindAllUnfinishedIPLibraryFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_FindAllUnfinishedIPLibraryFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).FindAllUnfinishedIPLibraryFiles(ctx, req.(*FindAllUnfinishedIPLibraryFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_FindIPLibraryFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindIPLibraryFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).FindIPLibraryFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_FindIPLibraryFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).FindIPLibraryFile(ctx, req.(*FindIPLibraryFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_CreateIPLibraryFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIPLibraryFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).CreateIPLibraryFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_CreateIPLibraryFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).CreateIPLibraryFile(ctx, req.(*CreateIPLibraryFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_CheckCountriesWithIPLibraryFileId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCountriesWithIPLibraryFileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).CheckCountriesWithIPLibraryFileId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_CheckCountriesWithIPLibraryFileId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).CheckCountriesWithIPLibraryFileId(ctx, req.(*CheckCountriesWithIPLibraryFileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_CheckProvincesWithIPLibraryFileId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckProvincesWithIPLibraryFileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).CheckProvincesWithIPLibraryFileId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_CheckProvincesWithIPLibraryFileId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).CheckProvincesWithIPLibraryFileId(ctx, req.(*CheckProvincesWithIPLibraryFileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_CheckCitiesWithIPLibraryFileId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCitiesWithIPLibraryFileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).CheckCitiesWithIPLibraryFileId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_CheckCitiesWithIPLibraryFileId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).CheckCitiesWithIPLibraryFileId(ctx, req.(*CheckCitiesWithIPLibraryFileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_CheckTownsWithIPLibraryFileId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTownsWithIPLibraryFileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).CheckTownsWithIPLibraryFileId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_CheckTownsWithIPLibraryFileId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).CheckTownsWithIPLibraryFileId(ctx, req.(*CheckTownsWithIPLibraryFileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_CheckProvidersWithIPLibraryFileId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckProvidersWithIPLibraryFileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).CheckProvidersWithIPLibraryFileId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_CheckProvidersWithIPLibraryFileId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).CheckProvidersWithIPLibraryFileId(ctx, req.(*CheckProvidersWithIPLibraryFileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_GenerateIPLibraryFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateIPLibraryFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).GenerateIPLibraryFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_GenerateIPLibraryFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).GenerateIPLibraryFile(ctx, req.(*GenerateIPLibraryFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_UpdateIPLibraryFileFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIPLibraryFileFinishedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).UpdateIPLibraryFileFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_UpdateIPLibraryFileFinished_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).UpdateIPLibraryFileFinished(ctx, req.(*UpdateIPLibraryFileFinishedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPLibraryFileService_DeleteIPLibraryFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIPLibraryFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPLibraryFileServiceServer).DeleteIPLibraryFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IPLibraryFileService_DeleteIPLibraryFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPLibraryFileServiceServer).DeleteIPLibraryFile(ctx, req.(*DeleteIPLibraryFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IPLibraryFileService_ServiceDesc is the grpc.ServiceDesc for IPLibraryFileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IPLibraryFileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.IPLibraryFileService",
	HandlerType: (*IPLibraryFileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findAllFinishedIPLibraryFiles",
			Handler:    _IPLibraryFileService_FindAllFinishedIPLibraryFiles_Handler,
		},
		{
			MethodName: "findAllUnfinishedIPLibraryFiles",
			Handler:    _IPLibraryFileService_FindAllUnfinishedIPLibraryFiles_Handler,
		},
		{
			MethodName: "findIPLibraryFile",
			Handler:    _IPLibraryFileService_FindIPLibraryFile_Handler,
		},
		{
			MethodName: "createIPLibraryFile",
			Handler:    _IPLibraryFileService_CreateIPLibraryFile_Handler,
		},
		{
			MethodName: "checkCountriesWithIPLibraryFileId",
			Handler:    _IPLibraryFileService_CheckCountriesWithIPLibraryFileId_Handler,
		},
		{
			MethodName: "checkProvincesWithIPLibraryFileId",
			Handler:    _IPLibraryFileService_CheckProvincesWithIPLibraryFileId_Handler,
		},
		{
			MethodName: "checkCitiesWithIPLibraryFileId",
			Handler:    _IPLibraryFileService_CheckCitiesWithIPLibraryFileId_Handler,
		},
		{
			MethodName: "checkTownsWithIPLibraryFileId",
			Handler:    _IPLibraryFileService_CheckTownsWithIPLibraryFileId_Handler,
		},
		{
			MethodName: "checkProvidersWithIPLibraryFileId",
			Handler:    _IPLibraryFileService_CheckProvidersWithIPLibraryFileId_Handler,
		},
		{
			MethodName: "generateIPLibraryFile",
			Handler:    _IPLibraryFileService_GenerateIPLibraryFile_Handler,
		},
		{
			MethodName: "updateIPLibraryFileFinished",
			Handler:    _IPLibraryFileService_UpdateIPLibraryFileFinished_Handler,
		},
		{
			MethodName: "deleteIPLibraryFile",
			Handler:    _IPLibraryFileService_DeleteIPLibraryFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ip_library_file.proto",
}
