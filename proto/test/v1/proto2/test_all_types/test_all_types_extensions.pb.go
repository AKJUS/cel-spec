// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/test/v1/proto2/test_all_types_extensions.proto

package test_all_types

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

type Proto2ExtensionScopedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Proto2ExtensionScopedMessage) Reset() {
	*x = Proto2ExtensionScopedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_test_v1_proto2_test_all_types_extensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proto2ExtensionScopedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto2ExtensionScopedMessage) ProtoMessage() {}

func (x *Proto2ExtensionScopedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_test_v1_proto2_test_all_types_extensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto2ExtensionScopedMessage.ProtoReflect.Descriptor instead.
func (*Proto2ExtensionScopedMessage) Descriptor() ([]byte, []int) {
	return file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDescGZIP(), []int{0}
}

var file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: (*int32)(nil),
		Field:         1000,
		Name:          "google.api.expr.test.v1.proto2.int32_ext",
		Tag:           "varint,1000,opt,name=int32_ext",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: (*TestAllTypes)(nil),
		Field:         1001,
		Name:          "google.api.expr.test.v1.proto2.nested_ext",
		Tag:           "bytes,1001,opt,name=nested_ext",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: (*TestAllTypes)(nil),
		Field:         1002,
		Name:          "google.api.expr.test.v1.proto2.test_all_types_ext",
		Tag:           "bytes,1002,opt,name=test_all_types_ext",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: (*TestAllTypes_NestedEnum)(nil),
		Field:         1003,
		Name:          "google.api.expr.test.v1.proto2.nested_enum_ext",
		Tag:           "varint,1003,opt,name=nested_enum_ext,enum=google.api.expr.test.v1.proto2.TestAllTypes_NestedEnum",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: ([]*TestAllTypes)(nil),
		Field:         1004,
		Name:          "google.api.expr.test.v1.proto2.repeated_test_all_types",
		Tag:           "bytes,1004,rep,name=repeated_test_all_types",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: (*int64)(nil),
		Field:         1005,
		Name:          "google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.int64_ext",
		Tag:           "varint,1005,opt,name=int64_ext",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: (*TestAllTypes)(nil),
		Field:         1006,
		Name:          "google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.message_scoped_nested_ext",
		Tag:           "bytes,1006,opt,name=message_scoped_nested_ext",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: (*TestAllTypes_NestedEnum)(nil),
		Field:         1007,
		Name:          "google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.nested_enum_ext",
		Tag:           "varint,1007,opt,name=nested_enum_ext,enum=google.api.expr.test.v1.proto2.TestAllTypes_NestedEnum",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
	{
		ExtendedType:  (*TestAllTypes)(nil),
		ExtensionType: ([]*TestAllTypes)(nil),
		Field:         1008,
		Name:          "google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.message_scoped_repeated_test_all_types",
		Tag:           "bytes,1008,rep,name=message_scoped_repeated_test_all_types",
		Filename:      "proto/test/v1/proto2/test_all_types_extensions.proto",
	},
}

// Extension fields to TestAllTypes.
var (
	// optional int32 int32_ext = 1000;
	E_Int32Ext = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[0]
	// optional google.api.expr.test.v1.proto2.TestAllTypes nested_ext = 1001;
	E_NestedExt = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[1]
	// optional google.api.expr.test.v1.proto2.TestAllTypes test_all_types_ext = 1002;
	E_TestAllTypesExt = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[2]
	// optional google.api.expr.test.v1.proto2.TestAllTypes.NestedEnum nested_enum_ext = 1003;
	E_NestedEnumExt = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[3]
	// repeated google.api.expr.test.v1.proto2.TestAllTypes repeated_test_all_types = 1004;
	E_RepeatedTestAllTypes = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[4]
	// optional int64 int64_ext = 1005;
	E_Proto2ExtensionScopedMessage_Int64Ext = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[5]
	// optional google.api.expr.test.v1.proto2.TestAllTypes message_scoped_nested_ext = 1006;
	E_Proto2ExtensionScopedMessage_MessageScopedNestedExt = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[6]
	// optional google.api.expr.test.v1.proto2.TestAllTypes.NestedEnum nested_enum_ext = 1007;
	E_Proto2ExtensionScopedMessage_NestedEnumExt = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[7]
	// repeated google.api.expr.test.v1.proto2.TestAllTypes message_scoped_repeated_test_all_types = 1008;
	E_Proto2ExtensionScopedMessage_MessageScopedRepeatedTestAllTypes = &file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes[8]
)

var File_proto_test_v1_proto2_test_all_types_extensions_proto protoreflect.FileDescriptor

var file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDesc = []byte{
	0x0a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x4a, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x65, 0x78, 0x74, 0x12,
	0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70,
	0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0xed, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x78, 0x74, 0x32, 0x96,
	0x01, 0x0a, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x64, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0xee, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52,
	0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x45, 0x78, 0x74, 0x32, 0x8e, 0x01, 0x0a, 0x0f, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0xef, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78,
	0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x74, 0x32, 0xae, 0x01, 0x0a, 0x26, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x18, 0xf0, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x21, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x4a, 0x0a, 0x09, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x45, 0x78, 0x74, 0x3a, 0x7a, 0x0a, 0x0a, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x65, 0x78, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x09, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x78,
	0x74, 0x3a, 0x88, 0x01, 0x0a, 0x12, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0f, 0x74, 0x65, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x45, 0x78, 0x74, 0x3a, 0x8e, 0x01, 0x0a,
	0x0f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x78, 0x74,
	0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78,
	0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0xeb,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0d,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x74, 0x3a, 0x92, 0x01,
	0x0a, 0x17, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0xec, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x14, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x42, 0x78, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x42, 0x16, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x65, 0x6c, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
}

var (
	file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDescOnce sync.Once
	file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDescData = file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDesc
)

func file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDescGZIP() []byte {
	file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDescOnce.Do(func() {
		file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDescData)
	})
	return file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDescData
}

var file_proto_test_v1_proto2_test_all_types_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_test_v1_proto2_test_all_types_extensions_proto_goTypes = []interface{}{
	(*Proto2ExtensionScopedMessage)(nil), // 0: google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage
	(*TestAllTypes)(nil),                 // 1: google.api.expr.test.v1.proto2.TestAllTypes
	(TestAllTypes_NestedEnum)(0),         // 2: google.api.expr.test.v1.proto2.TestAllTypes.NestedEnum
}
var file_proto_test_v1_proto2_test_all_types_extensions_proto_depIdxs = []int32{
	1,  // 0: google.api.expr.test.v1.proto2.int32_ext:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 1: google.api.expr.test.v1.proto2.nested_ext:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 2: google.api.expr.test.v1.proto2.test_all_types_ext:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 3: google.api.expr.test.v1.proto2.nested_enum_ext:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 4: google.api.expr.test.v1.proto2.repeated_test_all_types:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 5: google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.int64_ext:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 6: google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.message_scoped_nested_ext:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 7: google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.nested_enum_ext:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 8: google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.message_scoped_repeated_test_all_types:extendee -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 9: google.api.expr.test.v1.proto2.nested_ext:type_name -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 10: google.api.expr.test.v1.proto2.test_all_types_ext:type_name -> google.api.expr.test.v1.proto2.TestAllTypes
	2,  // 11: google.api.expr.test.v1.proto2.nested_enum_ext:type_name -> google.api.expr.test.v1.proto2.TestAllTypes.NestedEnum
	1,  // 12: google.api.expr.test.v1.proto2.repeated_test_all_types:type_name -> google.api.expr.test.v1.proto2.TestAllTypes
	1,  // 13: google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.message_scoped_nested_ext:type_name -> google.api.expr.test.v1.proto2.TestAllTypes
	2,  // 14: google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.nested_enum_ext:type_name -> google.api.expr.test.v1.proto2.TestAllTypes.NestedEnum
	1,  // 15: google.api.expr.test.v1.proto2.Proto2ExtensionScopedMessage.message_scoped_repeated_test_all_types:type_name -> google.api.expr.test.v1.proto2.TestAllTypes
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	9,  // [9:16] is the sub-list for extension type_name
	0,  // [0:9] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_test_v1_proto2_test_all_types_extensions_proto_init() }
func file_proto_test_v1_proto2_test_all_types_extensions_proto_init() {
	if File_proto_test_v1_proto2_test_all_types_extensions_proto != nil {
		return
	}
	file_proto_test_v1_proto2_test_all_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_test_v1_proto2_test_all_types_extensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proto2ExtensionScopedMessage); i {
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
			RawDescriptor: file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 9,
			NumServices:   0,
		},
		GoTypes:           file_proto_test_v1_proto2_test_all_types_extensions_proto_goTypes,
		DependencyIndexes: file_proto_test_v1_proto2_test_all_types_extensions_proto_depIdxs,
		MessageInfos:      file_proto_test_v1_proto2_test_all_types_extensions_proto_msgTypes,
		ExtensionInfos:    file_proto_test_v1_proto2_test_all_types_extensions_proto_extTypes,
	}.Build()
	File_proto_test_v1_proto2_test_all_types_extensions_proto = out.File
	file_proto_test_v1_proto2_test_all_types_extensions_proto_rawDesc = nil
	file_proto_test_v1_proto2_test_all_types_extensions_proto_goTypes = nil
	file_proto_test_v1_proto2_test_all_types_extensions_proto_depIdxs = nil
}