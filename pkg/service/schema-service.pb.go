// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: internal/proto-files/service/schema-service.proto

package service

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	domain "github.com/luigizuccarelli/golang-eventbus-grpc/pkg/domain"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetDataSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataschema *domain.DataSchema `protobuf:"bytes,1,opt,name=dataschema,proto3" json:"dataschema,omitempty"`
	Error      *Error             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetDataSchemaResponse) Reset() {
	*x = GetDataSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_files_service_schema_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataSchemaResponse) ProtoMessage() {}

func (x *GetDataSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_files_service_schema_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataSchemaResponse.ProtoReflect.Descriptor instead.
func (*GetDataSchemaResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_files_service_schema_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDataSchemaResponse) GetDataschema() *domain.DataSchema {
	if x != nil {
		return x.Dataschema
	}
	return nil
}

func (x *GetDataSchemaResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_files_service_schema_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_files_service_schema_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_internal_proto_files_service_schema_service_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_proto_files_service_schema_service_proto protoreflect.FileDescriptor

var file_internal_proto_files_service_schema_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x28, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x4e, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x17, 0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x52, 0x50,
	0x43, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_proto_files_service_schema_service_proto_rawDescOnce sync.Once
	file_internal_proto_files_service_schema_service_proto_rawDescData = file_internal_proto_files_service_schema_service_proto_rawDesc
)

func file_internal_proto_files_service_schema_service_proto_rawDescGZIP() []byte {
	file_internal_proto_files_service_schema_service_proto_rawDescOnce.Do(func() {
		file_internal_proto_files_service_schema_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_files_service_schema_service_proto_rawDescData)
	})
	return file_internal_proto_files_service_schema_service_proto_rawDescData
}

var file_internal_proto_files_service_schema_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_files_service_schema_service_proto_goTypes = []interface{}{
	(*GetDataSchemaResponse)(nil), // 0: service.GetDataSchemaResponse
	(*Error)(nil),                 // 1: service.Error
	(*domain.DataSchema)(nil),     // 2: domain.DataSchema
}
var file_internal_proto_files_service_schema_service_proto_depIdxs = []int32{
	2, // 0: service.GetDataSchemaResponse.dataschema:type_name -> domain.DataSchema
	1, // 1: service.GetDataSchemaResponse.error:type_name -> service.Error
	2, // 2: service.DataSchemaService.get:input_type -> domain.DataSchema
	0, // 3: service.DataSchemaService.get:output_type -> service.GetDataSchemaResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_files_service_schema_service_proto_init() }
func file_internal_proto_files_service_schema_service_proto_init() {
	if File_internal_proto_files_service_schema_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_files_service_schema_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataSchemaResponse); i {
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
		file_internal_proto_files_service_schema_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_internal_proto_files_service_schema_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_files_service_schema_service_proto_goTypes,
		DependencyIndexes: file_internal_proto_files_service_schema_service_proto_depIdxs,
		MessageInfos:      file_internal_proto_files_service_schema_service_proto_msgTypes,
	}.Build()
	File_internal_proto_files_service_schema_service_proto = out.File
	file_internal_proto_files_service_schema_service_proto_rawDesc = nil
	file_internal_proto_files_service_schema_service_proto_goTypes = nil
	file_internal_proto_files_service_schema_service_proto_depIdxs = nil
}
