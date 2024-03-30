// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_ns_access_log.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 创建访问日志
type CreateNSAccessLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsAccessLogs []*NSAccessLog `protobuf:"bytes,1,rep,name=nsAccessLogs,proto3" json:"nsAccessLogs,omitempty"`
}

func (x *CreateNSAccessLogsRequest) Reset() {
	*x = CreateNSAccessLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_access_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNSAccessLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNSAccessLogsRequest) ProtoMessage() {}

func (x *CreateNSAccessLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_access_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNSAccessLogsRequest.ProtoReflect.Descriptor instead.
func (*CreateNSAccessLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_access_log_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNSAccessLogsRequest) GetNsAccessLogs() []*NSAccessLog {
	if x != nil {
		return x.NsAccessLogs
	}
	return nil
}

type CreateNSAccessLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateNSAccessLogsResponse) Reset() {
	*x = CreateNSAccessLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_access_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNSAccessLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNSAccessLogsResponse) ProtoMessage() {}

func (x *CreateNSAccessLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_access_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNSAccessLogsResponse.ProtoReflect.Descriptor instead.
func (*CreateNSAccessLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_access_log_proto_rawDescGZIP(), []int{1}
}

// 列出往前的单页访问日志
type ListNSAccessLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId   string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`      // 上一页请求ID，可选
	NsClusterId int64  `protobuf:"varint,9,opt,name=nsClusterId,proto3" json:"nsClusterId,omitempty"` // 集群
	NsNodeId    int64  `protobuf:"varint,2,opt,name=nsNodeId,proto3" json:"nsNodeId,omitempty"`       // 节点ID
	NsDomainId  int64  `protobuf:"varint,3,opt,name=nsDomainId,proto3" json:"nsDomainId,omitempty"`   // 域名ID
	NsRecordId  int64  `protobuf:"varint,4,opt,name=nsRecordId,proto3" json:"nsRecordId,omitempty"`   // 记录ID
	Size        int64  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`               // 单页条数
	Day         string `protobuf:"bytes,6,opt,name=day,proto3" json:"day,omitempty"`                  // 日期，格式YYYYMMDD
	Reverse     bool   `protobuf:"varint,7,opt,name=reverse,proto3" json:"reverse,omitempty"`         // 是否反向查找，可选
	Keyword     string `protobuf:"bytes,8,opt,name=keyword,proto3" json:"keyword,omitempty"`          // 关键词
	RecordType  string `protobuf:"bytes,10,opt,name=recordType,proto3" json:"recordType,omitempty"`   // 记录类型
}

func (x *ListNSAccessLogsRequest) Reset() {
	*x = ListNSAccessLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_access_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNSAccessLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNSAccessLogsRequest) ProtoMessage() {}

func (x *ListNSAccessLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_access_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNSAccessLogsRequest.ProtoReflect.Descriptor instead.
func (*ListNSAccessLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_access_log_proto_rawDescGZIP(), []int{2}
}

func (x *ListNSAccessLogsRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ListNSAccessLogsRequest) GetNsClusterId() int64 {
	if x != nil {
		return x.NsClusterId
	}
	return 0
}

func (x *ListNSAccessLogsRequest) GetNsNodeId() int64 {
	if x != nil {
		return x.NsNodeId
	}
	return 0
}

func (x *ListNSAccessLogsRequest) GetNsDomainId() int64 {
	if x != nil {
		return x.NsDomainId
	}
	return 0
}

func (x *ListNSAccessLogsRequest) GetNsRecordId() int64 {
	if x != nil {
		return x.NsRecordId
	}
	return 0
}

func (x *ListNSAccessLogsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListNSAccessLogsRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ListNSAccessLogsRequest) GetReverse() bool {
	if x != nil {
		return x.Reverse
	}
	return false
}

func (x *ListNSAccessLogsRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListNSAccessLogsRequest) GetRecordType() string {
	if x != nil {
		return x.RecordType
	}
	return ""
}

type ListNSAccessLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsAccessLogs []*NSAccessLog `protobuf:"bytes,1,rep,name=nsAccessLogs,proto3" json:"nsAccessLogs,omitempty"`
	RequestId    string         `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
	HasMore      bool           `protobuf:"varint,3,opt,name=hasMore,proto3" json:"hasMore,omitempty"`
}

func (x *ListNSAccessLogsResponse) Reset() {
	*x = ListNSAccessLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_access_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNSAccessLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNSAccessLogsResponse) ProtoMessage() {}

func (x *ListNSAccessLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_access_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNSAccessLogsResponse.ProtoReflect.Descriptor instead.
func (*ListNSAccessLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_access_log_proto_rawDescGZIP(), []int{3}
}

func (x *ListNSAccessLogsResponse) GetNsAccessLogs() []*NSAccessLog {
	if x != nil {
		return x.NsAccessLogs
	}
	return nil
}

func (x *ListNSAccessLogsResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ListNSAccessLogsResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

// 查找单个日志
type FindNSAccessLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *FindNSAccessLogRequest) Reset() {
	*x = FindNSAccessLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_access_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNSAccessLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNSAccessLogRequest) ProtoMessage() {}

func (x *FindNSAccessLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_access_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNSAccessLogRequest.ProtoReflect.Descriptor instead.
func (*FindNSAccessLogRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_access_log_proto_rawDescGZIP(), []int{4}
}

func (x *FindNSAccessLogRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type FindNSAccessLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsAccessLog *NSAccessLog `protobuf:"bytes,1,opt,name=nsAccessLog,proto3" json:"nsAccessLog,omitempty"`
}

func (x *FindNSAccessLogResponse) Reset() {
	*x = FindNSAccessLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_access_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNSAccessLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNSAccessLogResponse) ProtoMessage() {}

func (x *FindNSAccessLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_access_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNSAccessLogResponse.ProtoReflect.Descriptor instead.
func (*FindNSAccessLogResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_access_log_proto_rawDescGZIP(), []int{5}
}

func (x *FindNSAccessLogResponse) GetNsAccessLog() *NSAccessLog {
	if x != nil {
		return x.NsAccessLog
	}
	return nil
}

var File_service_ns_access_log_proto protoreflect.FileDescriptor

var file_service_ns_access_log_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x73, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6e, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x33, 0x0a, 0x0c, 0x6e, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x0c, 0x6e, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x53, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xaf, 0x02, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x53, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6e, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6e,
	0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6e,
	0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x53,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x6e, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x0c, 0x6e, 0x73, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x22,
	0x36, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x4e,
	0x53, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x0b, 0x6e, 0x73, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x6f, 0x67, 0x32, 0x84, 0x02, 0x0a, 0x12, 0x4e, 0x53, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x12,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x4e, 0x53, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x53, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x53, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_ns_access_log_proto_rawDescOnce sync.Once
	file_service_ns_access_log_proto_rawDescData = file_service_ns_access_log_proto_rawDesc
)

func file_service_ns_access_log_proto_rawDescGZIP() []byte {
	file_service_ns_access_log_proto_rawDescOnce.Do(func() {
		file_service_ns_access_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ns_access_log_proto_rawDescData)
	})
	return file_service_ns_access_log_proto_rawDescData
}

var file_service_ns_access_log_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_ns_access_log_proto_goTypes = []interface{}{
	(*CreateNSAccessLogsRequest)(nil),  // 0: pb.CreateNSAccessLogsRequest
	(*CreateNSAccessLogsResponse)(nil), // 1: pb.CreateNSAccessLogsResponse
	(*ListNSAccessLogsRequest)(nil),    // 2: pb.ListNSAccessLogsRequest
	(*ListNSAccessLogsResponse)(nil),   // 3: pb.ListNSAccessLogsResponse
	(*FindNSAccessLogRequest)(nil),     // 4: pb.FindNSAccessLogRequest
	(*FindNSAccessLogResponse)(nil),    // 5: pb.FindNSAccessLogResponse
	(*NSAccessLog)(nil),                // 6: pb.NSAccessLog
}
var file_service_ns_access_log_proto_depIdxs = []int32{
	6, // 0: pb.CreateNSAccessLogsRequest.nsAccessLogs:type_name -> pb.NSAccessLog
	6, // 1: pb.ListNSAccessLogsResponse.nsAccessLogs:type_name -> pb.NSAccessLog
	6, // 2: pb.FindNSAccessLogResponse.nsAccessLog:type_name -> pb.NSAccessLog
	0, // 3: pb.NSAccessLogService.createNSAccessLogs:input_type -> pb.CreateNSAccessLogsRequest
	2, // 4: pb.NSAccessLogService.listNSAccessLogs:input_type -> pb.ListNSAccessLogsRequest
	4, // 5: pb.NSAccessLogService.findNSAccessLog:input_type -> pb.FindNSAccessLogRequest
	1, // 6: pb.NSAccessLogService.createNSAccessLogs:output_type -> pb.CreateNSAccessLogsResponse
	3, // 7: pb.NSAccessLogService.listNSAccessLogs:output_type -> pb.ListNSAccessLogsResponse
	5, // 8: pb.NSAccessLogService.findNSAccessLog:output_type -> pb.FindNSAccessLogResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_ns_access_log_proto_init() }
func file_service_ns_access_log_proto_init() {
	if File_service_ns_access_log_proto != nil {
		return
	}
	file_models_model_ns_access_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ns_access_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNSAccessLogsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_ns_access_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNSAccessLogsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_ns_access_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNSAccessLogsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_ns_access_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNSAccessLogsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_ns_access_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNSAccessLogRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_ns_access_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNSAccessLogResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_ns_access_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ns_access_log_proto_goTypes,
		DependencyIndexes: file_service_ns_access_log_proto_depIdxs,
		MessageInfos:      file_service_ns_access_log_proto_msgTypes,
	}.Build()
	File_service_ns_access_log_proto = out.File
	file_service_ns_access_log_proto_rawDesc = nil
	file_service_ns_access_log_proto_goTypes = nil
	file_service_ns_access_log_proto_depIdxs = nil
}