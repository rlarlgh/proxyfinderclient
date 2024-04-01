// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: grpc/proto/proxyfinder.proto

package proxyfinder

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proxyfinder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proxyfinder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proxyfinder_proto_rawDescGZIP(), []int{0}
}

type FindProxyServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUrl string `protobuf:"bytes,1,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
}

func (x *FindProxyServerRequest) Reset() {
	*x = FindProxyServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proxyfinder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindProxyServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindProxyServerRequest) ProtoMessage() {}

func (x *FindProxyServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proxyfinder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindProxyServerRequest.ProtoReflect.Descriptor instead.
func (*FindProxyServerRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proxyfinder_proto_rawDescGZIP(), []int{1}
}

func (x *FindProxyServerRequest) GetTargetUrl() string {
	if x != nil {
		return x.TargetUrl
	}
	return ""
}

type FindProxyServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *FindProxyServerResponse) Reset() {
	*x = FindProxyServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proxyfinder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindProxyServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindProxyServerResponse) ProtoMessage() {}

func (x *FindProxyServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proxyfinder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindProxyServerResponse.ProtoReflect.Descriptor instead.
func (*FindProxyServerResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proxyfinder_proto_rawDescGZIP(), []int{2}
}

func (x *FindProxyServerResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *FindProxyServerResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_grpc_proto_proxyfinder_proto protoreflect.FileDescriptor

var file_grpc_proto_proxyfinder_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x37, 0x0a, 0x16, 0x46, 0x69, 0x6e,
	0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x41, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x32, 0xd2, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x46,
	0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x66, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x15,
	0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x66, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_proxyfinder_proto_rawDescOnce sync.Once
	file_grpc_proto_proxyfinder_proto_rawDescData = file_grpc_proto_proxyfinder_proto_rawDesc
)

func file_grpc_proto_proxyfinder_proto_rawDescGZIP() []byte {
	file_grpc_proto_proxyfinder_proto_rawDescOnce.Do(func() {
		file_grpc_proto_proxyfinder_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_proxyfinder_proto_rawDescData)
	})
	return file_grpc_proto_proxyfinder_proto_rawDescData
}

var file_grpc_proto_proxyfinder_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_proto_proxyfinder_proto_goTypes = []interface{}{
	(*Empty)(nil),                   // 0: grpc.proxyfinder.Empty
	(*FindProxyServerRequest)(nil),  // 1: grpc.proxyfinder.FindProxyServerRequest
	(*FindProxyServerResponse)(nil), // 2: grpc.proxyfinder.FindProxyServerResponse
}
var file_grpc_proto_proxyfinder_proto_depIdxs = []int32{
	0, // 0: grpc.proxyfinder.ProxyFinder.FindProxyServer:input_type -> grpc.proxyfinder.Empty
	1, // 1: grpc.proxyfinder.ProxyFinder.FindProxyServerStrict:input_type -> grpc.proxyfinder.FindProxyServerRequest
	2, // 2: grpc.proxyfinder.ProxyFinder.FindProxyServer:output_type -> grpc.proxyfinder.FindProxyServerResponse
	2, // 3: grpc.proxyfinder.ProxyFinder.FindProxyServerStrict:output_type -> grpc.proxyfinder.FindProxyServerResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_proxyfinder_proto_init() }
func file_grpc_proto_proxyfinder_proto_init() {
	if File_grpc_proto_proxyfinder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_proxyfinder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_grpc_proto_proxyfinder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindProxyServerRequest); i {
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
		file_grpc_proto_proxyfinder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindProxyServerResponse); i {
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
			RawDescriptor: file_grpc_proto_proxyfinder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_proxyfinder_proto_goTypes,
		DependencyIndexes: file_grpc_proto_proxyfinder_proto_depIdxs,
		MessageInfos:      file_grpc_proto_proxyfinder_proto_msgTypes,
	}.Build()
	File_grpc_proto_proxyfinder_proto = out.File
	file_grpc_proto_proxyfinder_proto_rawDesc = nil
	file_grpc_proto_proxyfinder_proto_goTypes = nil
	file_grpc_proto_proxyfinder_proto_depIdxs = nil
}
