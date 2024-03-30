// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: service_acme_authentication.proto

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

// 获取Key
type FindACMEAuthenticationKeyWithTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *FindACMEAuthenticationKeyWithTokenRequest) Reset() {
	*x = FindACMEAuthenticationKeyWithTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_acme_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindACMEAuthenticationKeyWithTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindACMEAuthenticationKeyWithTokenRequest) ProtoMessage() {}

func (x *FindACMEAuthenticationKeyWithTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_acme_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindACMEAuthenticationKeyWithTokenRequest.ProtoReflect.Descriptor instead.
func (*FindACMEAuthenticationKeyWithTokenRequest) Descriptor() ([]byte, []int) {
	return file_service_acme_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *FindACMEAuthenticationKeyWithTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type FindACMEAuthenticationKeyWithTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *FindACMEAuthenticationKeyWithTokenResponse) Reset() {
	*x = FindACMEAuthenticationKeyWithTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_acme_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindACMEAuthenticationKeyWithTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindACMEAuthenticationKeyWithTokenResponse) ProtoMessage() {}

func (x *FindACMEAuthenticationKeyWithTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_acme_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindACMEAuthenticationKeyWithTokenResponse.ProtoReflect.Descriptor instead.
func (*FindACMEAuthenticationKeyWithTokenResponse) Descriptor() ([]byte, []int) {
	return file_service_acme_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *FindACMEAuthenticationKeyWithTokenResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_service_acme_authentication_proto protoreflect.FileDescriptor

var file_service_acme_authentication_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x6d, 0x65, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x41, 0x0a, 0x29, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x43, 0x4d, 0x45, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x2a, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x43, 0x4d, 0x45, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32, 0xa1, 0x01, 0x0a, 0x19, 0x41,
	0x43, 0x4d, 0x45, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x22, 0x66, 0x69, 0x6e,
	0x64, 0x41, 0x43, 0x4d, 0x45, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x43, 0x4d, 0x45, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x57, 0x69,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x43, 0x4d, 0x45, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x57, 0x69, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_acme_authentication_proto_rawDescOnce sync.Once
	file_service_acme_authentication_proto_rawDescData = file_service_acme_authentication_proto_rawDesc
)

func file_service_acme_authentication_proto_rawDescGZIP() []byte {
	file_service_acme_authentication_proto_rawDescOnce.Do(func() {
		file_service_acme_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_acme_authentication_proto_rawDescData)
	})
	return file_service_acme_authentication_proto_rawDescData
}

var file_service_acme_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_acme_authentication_proto_goTypes = []interface{}{
	(*FindACMEAuthenticationKeyWithTokenRequest)(nil),  // 0: pb.FindACMEAuthenticationKeyWithTokenRequest
	(*FindACMEAuthenticationKeyWithTokenResponse)(nil), // 1: pb.FindACMEAuthenticationKeyWithTokenResponse
}
var file_service_acme_authentication_proto_depIdxs = []int32{
	0, // 0: pb.ACMEAuthenticationService.findACMEAuthenticationKeyWithToken:input_type -> pb.FindACMEAuthenticationKeyWithTokenRequest
	1, // 1: pb.ACMEAuthenticationService.findACMEAuthenticationKeyWithToken:output_type -> pb.FindACMEAuthenticationKeyWithTokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_acme_authentication_proto_init() }
func file_service_acme_authentication_proto_init() {
	if File_service_acme_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_acme_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindACMEAuthenticationKeyWithTokenRequest); i {
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
		file_service_acme_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindACMEAuthenticationKeyWithTokenResponse); i {
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
			RawDescriptor: file_service_acme_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_acme_authentication_proto_goTypes,
		DependencyIndexes: file_service_acme_authentication_proto_depIdxs,
		MessageInfos:      file_service_acme_authentication_proto_msgTypes,
	}.Build()
	File_service_acme_authentication_proto = out.File
	file_service_acme_authentication_proto_rawDesc = nil
	file_service_acme_authentication_proto_goTypes = nil
	file_service_acme_authentication_proto_depIdxs = nil
}
