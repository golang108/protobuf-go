// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/annotations/annotations.proto

package annotations

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type AnnotationsTestEnum int32

const (
	AnnotationsTestEnum_ANNOTATIONS_TEST_ENUM_VALUE AnnotationsTestEnum = 0
)

// Enum value maps for AnnotationsTestEnum.
var (
	AnnotationsTestEnum_name = map[int32]string{
		0: "ANNOTATIONS_TEST_ENUM_VALUE",
	}
	AnnotationsTestEnum_value = map[string]int32{
		"ANNOTATIONS_TEST_ENUM_VALUE": 0,
	}
)

func (x AnnotationsTestEnum) Enum() *AnnotationsTestEnum {
	p := new(AnnotationsTestEnum)
	*p = x
	return p
}

func (x AnnotationsTestEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnnotationsTestEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_enumTypes[0].Descriptor()
}

func (AnnotationsTestEnum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_enumTypes[0]
}

func (x AnnotationsTestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AnnotationsTestEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AnnotationsTestEnum(num)
	return nil
}

// Deprecated: Use AnnotationsTestEnum.Descriptor instead.
func (AnnotationsTestEnum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDescGZIP(), []int{0}
}

type AnnotationsTestMessage struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	AnnotationsTestField *string                `protobuf:"bytes,1,opt,name=AnnotationsTestField" json:"AnnotationsTestField,omitempty"`
	// Deprecated: Do not use. This will be deleted in the near future.
	XXX_weak_M    struct{} `protobuf:"bytes,2,opt,name=m,weak=fmt.M" json:"m,omitempty"`
	weakFields    protoimpl.WeakFields
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnnotationsTestMessage) Reset() {
	*x = AnnotationsTestMessage{}
	mi := &file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnnotationsTestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationsTestMessage) ProtoMessage() {}

func (x *AnnotationsTestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationsTestMessage.ProtoReflect.Descriptor instead.
func (*AnnotationsTestMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotationsTestMessage) GetAnnotationsTestField() string {
	if x != nil && x.AnnotationsTestField != nil {
		return *x.AnnotationsTestField
	}
	return ""
}

func (x *AnnotationsTestMessage) GetM() proto.Message {
	var w protoimpl.WeakFields
	if x != nil {
		w = x.weakFields
	}
	return protoimpl.X.GetWeak(w, 2, "fmt.M")
}

func (x *AnnotationsTestMessage) SetM(v proto.Message) {
	var w *protoimpl.WeakFields
	if x != nil {
		w = &x.weakFields
	}
	protoimpl.X.SetWeak(w, 2, "fmt.M", v)
}

var File_cmd_protoc_gen_go_testdata_annotations_annotations_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDesc = string([]byte{
	0x0a, 0x38, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x2e, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x66, 0x6d, 0x74, 0x2f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x16, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x32, 0x0a, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54,
	0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x01, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x66, 0x6d, 0x74, 0x2e, 0x4d, 0x42, 0x02, 0x50, 0x01, 0x52, 0x01, 0x6d, 0x2a, 0x36,
	0x0a, 0x13, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73,
	0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x58, 0x00,
})

var (
	file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDescData []byte
)

func file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDesc)))
	})
	return file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_goTypes = []any{
	(AnnotationsTestEnum)(0),       // 0: goproto.protoc.annotations.AnnotationsTestEnum
	(*AnnotationsTestMessage)(nil), // 1: goproto.protoc.annotations.AnnotationsTestMessage
}
var file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_init() }
func file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_init() {
	if File_cmd_protoc_gen_go_testdata_annotations_annotations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_annotations_annotations_proto = out.File
	file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_annotations_annotations_proto_depIdxs = nil
}
