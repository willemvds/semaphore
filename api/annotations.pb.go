// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: api/annotations.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package   string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Host      string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Transport string `protobuf:"bytes,4,opt,name=transport,proto3" json:"transport,omitempty"`
	Codec     string `protobuf:"bytes,5,opt,name=codec,proto3" json:"codec,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_api_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_api_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Service) GetTransport() string {
	if x != nil {
		return x.Transport
	}
	return ""
}

func (x *Service) GetCodec() string {
	if x != nil {
		return x.Codec
	}
	return ""
}

type HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Method   string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *HTTP) Reset() {
	*x = HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTP) ProtoMessage() {}

func (x *HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_api_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTP.ProtoReflect.Descriptor instead.
func (*HTTP) Descriptor() ([]byte, []int) {
	return file_api_annotations_proto_rawDescGZIP(), []int{1}
}

func (x *HTTP) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *HTTP) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

var file_api_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*Service)(nil),
		Field:         50012,
		Name:          "semaphore.api.service",
		Tag:           "bytes,50012,opt,name=service",
		Filename:      "api/annotations.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*HTTP)(nil),
		Field:         50011,
		Name:          "semaphore.api.http",
		Tag:           "bytes,50011,opt,name=http",
		Filename:      "api/annotations.proto",
	},
}

// Extension fields to descriptor.ServiceOptions.
var (
	// optional semaphore.api.Service service = 50012;
	E_Service = &file_api_annotations_proto_extTypes[0]
)

// Extension fields to descriptor.MethodOptions.
var (
	// optional semaphore.api.HTTP http = 50011;
	E_Http = &file_api_annotations_proto_extTypes[1]
)

var File_api_annotations_proto protoreflect.FileDescriptor

var file_api_annotations_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x22, 0x3a, 0x0a, 0x04, 0x48, 0x54, 0x54,
	0x50, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x3a, 0x53, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xdc, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x6d, 0x61,
	0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x49, 0x0a, 0x04, 0x68, 0x74,
	0x74, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xdb, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x6d,
	0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x78, 0x69, 0x61, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68,
	0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_annotations_proto_rawDescOnce sync.Once
	file_api_annotations_proto_rawDescData = file_api_annotations_proto_rawDesc
)

func file_api_annotations_proto_rawDescGZIP() []byte {
	file_api_annotations_proto_rawDescOnce.Do(func() {
		file_api_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_annotations_proto_rawDescData)
	})
	return file_api_annotations_proto_rawDescData
}

var file_api_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_annotations_proto_goTypes = []interface{}{
	(*Service)(nil),                   // 0: semaphore.api.Service
	(*HTTP)(nil),                      // 1: semaphore.api.HTTP
	(*descriptor.ServiceOptions)(nil), // 2: google.protobuf.ServiceOptions
	(*descriptor.MethodOptions)(nil),  // 3: google.protobuf.MethodOptions
}
var file_api_annotations_proto_depIdxs = []int32{
	2, // 0: semaphore.api.service:extendee -> google.protobuf.ServiceOptions
	3, // 1: semaphore.api.http:extendee -> google.protobuf.MethodOptions
	0, // 2: semaphore.api.service:type_name -> semaphore.api.Service
	1, // 3: semaphore.api.http:type_name -> semaphore.api.HTTP
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_annotations_proto_init() }
func file_api_annotations_proto_init() {
	if File_api_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_api_annotations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTP); i {
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
			RawDescriptor: file_api_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_api_annotations_proto_goTypes,
		DependencyIndexes: file_api_annotations_proto_depIdxs,
		MessageInfos:      file_api_annotations_proto_msgTypes,
		ExtensionInfos:    file_api_annotations_proto_extTypes,
	}.Build()
	File_api_annotations_proto = out.File
	file_api_annotations_proto_rawDesc = nil
	file_api_annotations_proto_goTypes = nil
	file_api_annotations_proto_depIdxs = nil
}
