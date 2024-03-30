// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_traffic_package_price.proto

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

// 流量包价格定义
type TrafficPackagePrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrafficPackageId       int64   `protobuf:"varint,1,opt,name=trafficPackageId,proto3" json:"trafficPackageId,omitempty"`
	NodeRegionId           int64   `protobuf:"varint,2,opt,name=nodeRegionId,proto3" json:"nodeRegionId,omitempty"`
	TrafficPackagePeriodId int64   `protobuf:"varint,3,opt,name=trafficPackagePeriodId,proto3" json:"trafficPackagePeriodId,omitempty"`
	Price                  float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *TrafficPackagePrice) Reset() {
	*x = TrafficPackagePrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_traffic_package_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficPackagePrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficPackagePrice) ProtoMessage() {}

func (x *TrafficPackagePrice) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_traffic_package_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficPackagePrice.ProtoReflect.Descriptor instead.
func (*TrafficPackagePrice) Descriptor() ([]byte, []int) {
	return file_models_model_traffic_package_price_proto_rawDescGZIP(), []int{0}
}

func (x *TrafficPackagePrice) GetTrafficPackageId() int64 {
	if x != nil {
		return x.TrafficPackageId
	}
	return 0
}

func (x *TrafficPackagePrice) GetNodeRegionId() int64 {
	if x != nil {
		return x.NodeRegionId
	}
	return 0
}

func (x *TrafficPackagePrice) GetTrafficPackagePeriodId() int64 {
	if x != nil {
		return x.TrafficPackagePeriodId
	}
	return 0
}

func (x *TrafficPackagePrice) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_models_model_traffic_package_price_proto protoreflect.FileDescriptor

var file_models_model_traffic_package_price_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xb3,
	0x01, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_traffic_package_price_proto_rawDescOnce sync.Once
	file_models_model_traffic_package_price_proto_rawDescData = file_models_model_traffic_package_price_proto_rawDesc
)

func file_models_model_traffic_package_price_proto_rawDescGZIP() []byte {
	file_models_model_traffic_package_price_proto_rawDescOnce.Do(func() {
		file_models_model_traffic_package_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_traffic_package_price_proto_rawDescData)
	})
	return file_models_model_traffic_package_price_proto_rawDescData
}

var file_models_model_traffic_package_price_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_traffic_package_price_proto_goTypes = []interface{}{
	(*TrafficPackagePrice)(nil), // 0: pb.TrafficPackagePrice
}
var file_models_model_traffic_package_price_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_traffic_package_price_proto_init() }
func file_models_model_traffic_package_price_proto_init() {
	if File_models_model_traffic_package_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_traffic_package_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficPackagePrice); i {
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
			RawDescriptor: file_models_model_traffic_package_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_traffic_package_price_proto_goTypes,
		DependencyIndexes: file_models_model_traffic_package_price_proto_depIdxs,
		MessageInfos:      file_models_model_traffic_package_price_proto_msgTypes,
	}.Build()
	File_models_model_traffic_package_price_proto = out.File
	file_models_model_traffic_package_price_proto_rawDesc = nil
	file_models_model_traffic_package_price_proto_goTypes = nil
	file_models_model_traffic_package_price_proto_depIdxs = nil
}