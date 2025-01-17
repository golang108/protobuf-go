// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/lazy/lazy_extension_test.proto

package lazy

import (
	messagesetpb "google.golang.org/protobuf/internal/testprotos/messageset/messagesetpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type FlyingFoxSpecies int32

const (
	FlyingFoxSpecies_FLYING_FOX_UNDEFINED FlyingFoxSpecies = 0
	FlyingFoxSpecies_GREY_HEADED          FlyingFoxSpecies = 1
	FlyingFoxSpecies_BLACK                FlyingFoxSpecies = 2
	FlyingFoxSpecies_SPECTACLED           FlyingFoxSpecies = 3
	FlyingFoxSpecies_LARGE_EARED          FlyingFoxSpecies = 4
	FlyingFoxSpecies_DUSKY                FlyingFoxSpecies = 5
	FlyingFoxSpecies_TORRESIAN            FlyingFoxSpecies = 6
	FlyingFoxSpecies_BARE_BACKED          FlyingFoxSpecies = 7
)

// Enum value maps for FlyingFoxSpecies.
var (
	FlyingFoxSpecies_name = map[int32]string{
		0: "FLYING_FOX_UNDEFINED",
		1: "GREY_HEADED",
		2: "BLACK",
		3: "SPECTACLED",
		4: "LARGE_EARED",
		5: "DUSKY",
		6: "TORRESIAN",
		7: "BARE_BACKED",
	}
	FlyingFoxSpecies_value = map[string]int32{
		"FLYING_FOX_UNDEFINED": 0,
		"GREY_HEADED":          1,
		"BLACK":                2,
		"SPECTACLED":           3,
		"LARGE_EARED":          4,
		"DUSKY":                5,
		"TORRESIAN":            6,
		"BARE_BACKED":          7,
	}
)

func (x FlyingFoxSpecies) Enum() *FlyingFoxSpecies {
	p := new(FlyingFoxSpecies)
	*p = x
	return p
}

func (x FlyingFoxSpecies) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlyingFoxSpecies) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_enumTypes[0].Descriptor()
}

func (FlyingFoxSpecies) Type() protoreflect.EnumType {
	return &file_internal_testprotos_lazy_lazy_extension_test_proto_enumTypes[0]
}

func (x FlyingFoxSpecies) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *FlyingFoxSpecies) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = FlyingFoxSpecies(num)
	return nil
}

// Deprecated: Use FlyingFoxSpecies.Descriptor instead.
func (FlyingFoxSpecies) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{0}
}

type PipistrelleSpecies int32

const (
	PipistrelleSpecies_PIPISTRELLE_UNDEFINED PipistrelleSpecies = 0
	PipistrelleSpecies_FOREST                PipistrelleSpecies = 1
	PipistrelleSpecies_INDIAN                PipistrelleSpecies = 2
	PipistrelleSpecies_EGYPTIAN              PipistrelleSpecies = 3
	PipistrelleSpecies_RUSTY                 PipistrelleSpecies = 4
	PipistrelleSpecies_LEAST                 PipistrelleSpecies = 5
)

// Enum value maps for PipistrelleSpecies.
var (
	PipistrelleSpecies_name = map[int32]string{
		0: "PIPISTRELLE_UNDEFINED",
		1: "FOREST",
		2: "INDIAN",
		3: "EGYPTIAN",
		4: "RUSTY",
		5: "LEAST",
	}
	PipistrelleSpecies_value = map[string]int32{
		"PIPISTRELLE_UNDEFINED": 0,
		"FOREST":                1,
		"INDIAN":                2,
		"EGYPTIAN":              3,
		"RUSTY":                 4,
		"LEAST":                 5,
	}
)

func (x PipistrelleSpecies) Enum() *PipistrelleSpecies {
	p := new(PipistrelleSpecies)
	*p = x
	return p
}

func (x PipistrelleSpecies) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PipistrelleSpecies) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_enumTypes[1].Descriptor()
}

func (PipistrelleSpecies) Type() protoreflect.EnumType {
	return &file_internal_testprotos_lazy_lazy_extension_test_proto_enumTypes[1]
}

func (x PipistrelleSpecies) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *PipistrelleSpecies) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = PipistrelleSpecies(num)
	return nil
}

// Deprecated: Use PipistrelleSpecies.Descriptor instead.
func (PipistrelleSpecies) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{1}
}

// This message contains a message set.
type Holder struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Data          *messagesetpb.MessageSet `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Holder) Reset() {
	*x = Holder{}
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Holder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Holder) ProtoMessage() {}

func (x *Holder) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Holder.ProtoReflect.Descriptor instead.
func (*Holder) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{0}
}

func (x *Holder) GetData() *messagesetpb.MessageSet {
	if x != nil {
		return x.Data
	}
	return nil
}

// This message may be inserted into a message set.
type Rabbit struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          *string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Rabbit) Reset() {
	*x = Rabbit{}
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Rabbit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rabbit) ProtoMessage() {}

func (x *Rabbit) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rabbit.ProtoReflect.Descriptor instead.
func (*Rabbit) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{1}
}

func (x *Rabbit) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type FlyingFox struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Species       *FlyingFoxSpecies      `protobuf:"varint,1,opt,name=species,enum=lazy_extension_test.FlyingFoxSpecies" json:"species,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FlyingFox) Reset() {
	*x = FlyingFox{}
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlyingFox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlyingFox) ProtoMessage() {}

func (x *FlyingFox) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlyingFox.ProtoReflect.Descriptor instead.
func (*FlyingFox) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{2}
}

func (x *FlyingFox) GetSpecies() FlyingFoxSpecies {
	if x != nil && x.Species != nil {
		return *x.Species
	}
	return FlyingFoxSpecies_FLYING_FOX_UNDEFINED
}

type Tree struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Eucalyptus      *bool                  `protobuf:"varint,1,opt,name=eucalyptus" json:"eucalyptus,omitempty"`
	extensionFields protoimpl.ExtensionFields
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Tree) Reset() {
	*x = Tree{}
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tree) ProtoMessage() {}

func (x *Tree) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tree.ProtoReflect.Descriptor instead.
func (*Tree) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{3}
}

func (x *Tree) GetEucalyptus() bool {
	if x != nil && x.Eucalyptus != nil {
		return *x.Eucalyptus
	}
	return false
}

type Pipistrelle struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Species       *PipistrelleSpecies    `protobuf:"varint,1,opt,name=species,enum=lazy_extension_test.PipistrelleSpecies" json:"species,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pipistrelle) Reset() {
	*x = Pipistrelle{}
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pipistrelle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipistrelle) ProtoMessage() {}

func (x *Pipistrelle) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipistrelle.ProtoReflect.Descriptor instead.
func (*Pipistrelle) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{4}
}

func (x *Pipistrelle) GetSpecies() PipistrelleSpecies {
	if x != nil && x.Species != nil {
		return *x.Species
	}
	return PipistrelleSpecies_PIPISTRELLE_UNDEFINED
}

type Pipistrelles struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Species       *PipistrelleSpecies    `protobuf:"varint,1,opt,name=species,enum=lazy_extension_test.PipistrelleSpecies" json:"species,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pipistrelles) Reset() {
	*x = Pipistrelles{}
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pipistrelles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipistrelles) ProtoMessage() {}

func (x *Pipistrelles) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipistrelles.ProtoReflect.Descriptor instead.
func (*Pipistrelles) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{5}
}

func (x *Pipistrelles) GetSpecies() PipistrelleSpecies {
	if x != nil && x.Species != nil {
		return *x.Species
	}
	return PipistrelleSpecies_PIPISTRELLE_UNDEFINED
}

// And the ugly version that is not encouraged
type BatNest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatNest) Reset() {
	*x = BatNest{}
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatNest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatNest) ProtoMessage() {}

func (x *BatNest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatNest.ProtoReflect.Descriptor instead.
func (*BatNest) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP(), []int{6}
}

var file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*Tree)(nil),
		ExtensionType: (*FlyingFox)(nil),
		Field:         345570595,
		Name:          "lazy_extension_test.bat",
		Tag:           "bytes,345570595,opt,name=bat",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
	{
		ExtendedType:  (*Tree)(nil),
		ExtensionType: (*FlyingFox)(nil),
		Field:         345570596,
		Name:          "lazy_extension_test.bat_pup",
		Tag:           "bytes,345570596,opt,name=bat_pup",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
	{
		ExtendedType:  (*Tree)(nil),
		ExtensionType: ([]*FlyingFox)(nil),
		Field:         345570597,
		Name:          "lazy_extension_test.bat_posse",
		Tag:           "bytes,345570597,rep,name=bat_posse",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
	{
		ExtendedType:  (*Tree)(nil),
		ExtensionType: ([]byte)(nil),
		Field:         345570598,
		Name:          "lazy_extension_test.binary_bat",
		Tag:           "bytes,345570598,opt,name=binary_bat",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
	{
		ExtendedType:  (*Tree)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         345570599,
		Name:          "lazy_extension_test.integer_bat",
		Tag:           "varint,345570599,opt,name=integer_bat",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
	{
		ExtendedType:  (*Tree)(nil),
		ExtensionType: (*Pipistrelle)(nil),
		Field:         345570600,
		Name:          "lazy_extension_test.pipistrelle",
		Tag:           "group,345570600,opt,name=Pipistrelle",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
	{
		ExtendedType:  (*Tree)(nil),
		ExtensionType: ([]*Pipistrelles)(nil),
		Field:         345570601,
		Name:          "lazy_extension_test.pipistrelles",
		Tag:           "group,345570601,rep,name=Pipistrelles",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
	{
		ExtendedType:  (*messagesetpb.MessageSet)(nil),
		ExtensionType: (*Rabbit)(nil),
		Field:         345570595,
		Name:          "lazy_extension_test.Rabbit.message_set_extension",
		Tag:           "bytes,345570595,opt,name=message_set_extension",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
	{
		ExtendedType:  (*Tree)(nil),
		ExtensionType: (*FlyingFox)(nil),
		Field:         345570602,
		Name:          "lazy_extension_test.BatNest.bat",
		Tag:           "bytes,345570602,opt,name=bat",
		Filename:      "internal/testprotos/lazy/lazy_extension_test.proto",
	},
}

// Extension fields to Tree.
var (
	// optional lazy_extension_test.FlyingFox bat = 345570595;
	E_Bat = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[0]
	// optional lazy_extension_test.FlyingFox bat_pup = 345570596;
	E_BatPup = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[1]
	// repeated lazy_extension_test.FlyingFox bat_posse = 345570597;
	E_BatPosse = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[2]
	// optional bytes binary_bat = 345570598;
	E_BinaryBat = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[3]
	// optional uint32 integer_bat = 345570599;
	E_IntegerBat = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[4]
	// optional lazy_extension_test.Pipistrelle pipistrelle = 345570600;
	E_Pipistrelle = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[5]
	// repeated lazy_extension_test.Pipistrelles pipistrelles = 345570601;
	E_Pipistrelles = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[6]
	// optional lazy_extension_test.FlyingFox bat = 345570602;
	E_BatNest_Bat = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[8]
)

// Extension fields to messagesetpb.MessageSet.
var (
	// optional lazy_extension_test.Rabbit message_set_extension = 345570595;
	E_Rabbit_MessageSetExtension = &file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes[7]
)

var File_internal_testprotos_lazy_lazy_extension_test_proto protoreflect.FileDescriptor

var file_internal_testprotos_lazy_lazy_extension_test_proto_rawDesc = string([]byte{
	0x0a, 0x32, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x65, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x06, 0x48, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x97, 0x01, 0x0a,
	0x06, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x79, 0x0a, 0x15, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x18, 0xa3, 0xfa, 0xe3, 0xa4, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x61, 0x62, 0x62, 0x69,
	0x74, 0x52, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x09, 0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67,
	0x46, 0x6f, 0x78, 0x12, 0x3f, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x6c, 0x79, 0x69, 0x6e,
	0x67, 0x46, 0x6f, 0x78, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x07, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x75, 0x63, 0x61, 0x6c, 0x79, 0x70, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x65, 0x75, 0x63, 0x61, 0x6c, 0x79, 0x70, 0x74, 0x75, 0x73, 0x2a, 0x09, 0x08, 0x90,
	0x4e, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x22, 0x50, 0x0a, 0x0b, 0x50, 0x69, 0x70, 0x69, 0x73,
	0x74, 0x72, 0x65, 0x6c, 0x6c, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x69,
	0x70, 0x69, 0x73, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x22, 0x51, 0x0a, 0x0c, 0x50, 0x69, 0x70,
	0x69, 0x73, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6c, 0x61, 0x7a,
	0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x69, 0x70, 0x69, 0x73, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x65, 0x73, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x07,
	0x42, 0x61, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x32, 0x4f, 0x0a, 0x03, 0x62, 0x61, 0x74, 0x12, 0x19,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x18, 0xaa, 0xfa, 0xe3, 0xa4, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67,
	0x46, 0x6f, 0x78, 0x52, 0x03, 0x62, 0x61, 0x74, 0x2a, 0x94, 0x01, 0x0a, 0x10, 0x46, 0x6c, 0x79,
	0x69, 0x6e, 0x67, 0x46, 0x6f, 0x78, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x14, 0x46, 0x4c, 0x59, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x58, 0x5f, 0x55, 0x4e, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x52, 0x45, 0x59, 0x5f,
	0x48, 0x45, 0x41, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x41, 0x43,
	0x4b, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x50, 0x45, 0x43, 0x54, 0x41, 0x43, 0x4c, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x45, 0x41, 0x52,
	0x45, 0x44, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x55, 0x53, 0x4b, 0x59, 0x10, 0x05, 0x12,
	0x0d, 0x0a, 0x09, 0x54, 0x4f, 0x52, 0x52, 0x45, 0x53, 0x49, 0x41, 0x4e, 0x10, 0x06, 0x12, 0x0f,
	0x0a, 0x0b, 0x42, 0x41, 0x52, 0x45, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x07, 0x2a,
	0x6b, 0x0a, 0x12, 0x50, 0x69, 0x70, 0x69, 0x73, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x49, 0x50, 0x49, 0x53, 0x54, 0x52,
	0x45, 0x4c, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4f, 0x52, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x49, 0x4e, 0x44, 0x49, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x47, 0x59, 0x50,
	0x54, 0x49, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x55, 0x53, 0x54, 0x59, 0x10,
	0x04, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x45, 0x41, 0x53, 0x54, 0x10, 0x05, 0x3a, 0x4f, 0x0a, 0x03,
	0x62, 0x61, 0x74, 0x12, 0x19, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x18, 0xa3,
	0xfa, 0xe3, 0xa4, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46,
	0x6c, 0x79, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x78, 0x52, 0x03, 0x62, 0x61, 0x74, 0x3a, 0x56, 0x0a,
	0x07, 0x62, 0x61, 0x74, 0x5f, 0x70, 0x75, 0x70, 0x12, 0x19, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x72, 0x65, 0x65, 0x18, 0xa4, 0xfa, 0xe3, 0xa4, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x46, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x78, 0x52, 0x06, 0x62,
	0x61, 0x74, 0x50, 0x75, 0x70, 0x3a, 0x5a, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x5f, 0x70, 0x6f, 0x73,
	0x73, 0x65, 0x12, 0x19, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x18, 0xa5, 0xfa,
	0xe3, 0xa4, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x6c,
	0x79, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x78, 0x52, 0x08, 0x62, 0x61, 0x74, 0x50, 0x6f, 0x73, 0x73,
	0x65, 0x3a, 0x3c, 0x0a, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x62, 0x61, 0x74, 0x12,
	0x19, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x18, 0xa6, 0xfa, 0xe3, 0xa4, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x42, 0x61, 0x74, 0x3a,
	0x3e, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x74, 0x12, 0x19,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x18, 0xa7, 0xfa, 0xe3, 0xa4, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x42, 0x61, 0x74, 0x3a,
	0x61, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x69, 0x73, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x65, 0x12, 0x19,
	0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x18, 0xa8, 0xfa, 0xe3, 0xa4, 0x01, 0x20,
	0x01, 0x28, 0x0a, 0x32, 0x20, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x69, 0x70, 0x69, 0x73, 0x74,
	0x72, 0x65, 0x6c, 0x6c, 0x65, 0x52, 0x0b, 0x70, 0x69, 0x70, 0x69, 0x73, 0x74, 0x72, 0x65, 0x6c,
	0x6c, 0x65, 0x3a, 0x64, 0x0a, 0x0c, 0x70, 0x69, 0x70, 0x69, 0x73, 0x74, 0x72, 0x65, 0x6c, 0x6c,
	0x65, 0x73, 0x12, 0x19, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x18, 0xa9, 0xfa,
	0xe3, 0xa4, 0x01, 0x20, 0x03, 0x28, 0x0a, 0x32, 0x21, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x69,
	0x70, 0x69, 0x73, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x65, 0x73, 0x52, 0x0c, 0x70, 0x69, 0x70, 0x69,
	0x73, 0x74, 0x72, 0x65, 0x6c, 0x6c, 0x65, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a, 0x79,
})

var (
	file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescData []byte
)

func file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_testprotos_lazy_lazy_extension_test_proto_rawDesc), len(file_internal_testprotos_lazy_lazy_extension_test_proto_rawDesc)))
	})
	return file_internal_testprotos_lazy_lazy_extension_test_proto_rawDescData
}

var file_internal_testprotos_lazy_lazy_extension_test_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_testprotos_lazy_lazy_extension_test_proto_goTypes = []any{
	(FlyingFoxSpecies)(0),           // 0: lazy_extension_test.FlyingFoxSpecies
	(PipistrelleSpecies)(0),         // 1: lazy_extension_test.PipistrelleSpecies
	(*Holder)(nil),                  // 2: lazy_extension_test.Holder
	(*Rabbit)(nil),                  // 3: lazy_extension_test.Rabbit
	(*FlyingFox)(nil),               // 4: lazy_extension_test.FlyingFox
	(*Tree)(nil),                    // 5: lazy_extension_test.Tree
	(*Pipistrelle)(nil),             // 6: lazy_extension_test.Pipistrelle
	(*Pipistrelles)(nil),            // 7: lazy_extension_test.Pipistrelles
	(*BatNest)(nil),                 // 8: lazy_extension_test.BatNest
	(*messagesetpb.MessageSet)(nil), // 9: goproto.proto.messageset.MessageSet
}
var file_internal_testprotos_lazy_lazy_extension_test_proto_depIdxs = []int32{
	9,  // 0: lazy_extension_test.Holder.data:type_name -> goproto.proto.messageset.MessageSet
	0,  // 1: lazy_extension_test.FlyingFox.species:type_name -> lazy_extension_test.FlyingFoxSpecies
	1,  // 2: lazy_extension_test.Pipistrelle.species:type_name -> lazy_extension_test.PipistrelleSpecies
	1,  // 3: lazy_extension_test.Pipistrelles.species:type_name -> lazy_extension_test.PipistrelleSpecies
	5,  // 4: lazy_extension_test.bat:extendee -> lazy_extension_test.Tree
	5,  // 5: lazy_extension_test.bat_pup:extendee -> lazy_extension_test.Tree
	5,  // 6: lazy_extension_test.bat_posse:extendee -> lazy_extension_test.Tree
	5,  // 7: lazy_extension_test.binary_bat:extendee -> lazy_extension_test.Tree
	5,  // 8: lazy_extension_test.integer_bat:extendee -> lazy_extension_test.Tree
	5,  // 9: lazy_extension_test.pipistrelle:extendee -> lazy_extension_test.Tree
	5,  // 10: lazy_extension_test.pipistrelles:extendee -> lazy_extension_test.Tree
	9,  // 11: lazy_extension_test.Rabbit.message_set_extension:extendee -> goproto.proto.messageset.MessageSet
	5,  // 12: lazy_extension_test.BatNest.bat:extendee -> lazy_extension_test.Tree
	4,  // 13: lazy_extension_test.bat:type_name -> lazy_extension_test.FlyingFox
	4,  // 14: lazy_extension_test.bat_pup:type_name -> lazy_extension_test.FlyingFox
	4,  // 15: lazy_extension_test.bat_posse:type_name -> lazy_extension_test.FlyingFox
	6,  // 16: lazy_extension_test.pipistrelle:type_name -> lazy_extension_test.Pipistrelle
	7,  // 17: lazy_extension_test.pipistrelles:type_name -> lazy_extension_test.Pipistrelles
	3,  // 18: lazy_extension_test.Rabbit.message_set_extension:type_name -> lazy_extension_test.Rabbit
	4,  // 19: lazy_extension_test.BatNest.bat:type_name -> lazy_extension_test.FlyingFox
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	13, // [13:20] is the sub-list for extension type_name
	4,  // [4:13] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_internal_testprotos_lazy_lazy_extension_test_proto_init() }
func file_internal_testprotos_lazy_lazy_extension_test_proto_init() {
	if File_internal_testprotos_lazy_lazy_extension_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_lazy_lazy_extension_test_proto_rawDesc), len(file_internal_testprotos_lazy_lazy_extension_test_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 9,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_lazy_lazy_extension_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_lazy_lazy_extension_test_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_lazy_lazy_extension_test_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_lazy_lazy_extension_test_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_lazy_lazy_extension_test_proto_extTypes,
	}.Build()
	File_internal_testprotos_lazy_lazy_extension_test_proto = out.File
	file_internal_testprotos_lazy_lazy_extension_test_proto_goTypes = nil
	file_internal_testprotos_lazy_lazy_extension_test_proto_depIdxs = nil
}
