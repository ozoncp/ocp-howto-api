// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: ocp-howto-api.proto

package ocp_howto_api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateHowtoV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId uint64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Question string `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	Answer   string `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *CreateHowtoV1Request) Reset() {
	*x = CreateHowtoV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHowtoV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHowtoV1Request) ProtoMessage() {}

func (x *CreateHowtoV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHowtoV1Request.ProtoReflect.Descriptor instead.
func (*CreateHowtoV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHowtoV1Request) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *CreateHowtoV1Request) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *CreateHowtoV1Request) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type CreateHowtoV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateHowtoV1Response) Reset() {
	*x = CreateHowtoV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHowtoV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHowtoV1Response) ProtoMessage() {}

func (x *CreateHowtoV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHowtoV1Response.ProtoReflect.Descriptor instead.
func (*CreateHowtoV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHowtoV1Response) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeHowtoV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DescribeHowtoV1Request) Reset() {
	*x = DescribeHowtoV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeHowtoV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeHowtoV1Request) ProtoMessage() {}

func (x *DescribeHowtoV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeHowtoV1Request.ProtoReflect.Descriptor instead.
func (*DescribeHowtoV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeHowtoV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeHowtoV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Howto *Howto `protobuf:"bytes,1,opt,name=howto,proto3" json:"howto,omitempty"`
}

func (x *DescribeHowtoV1Response) Reset() {
	*x = DescribeHowtoV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeHowtoV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeHowtoV1Response) ProtoMessage() {}

func (x *DescribeHowtoV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeHowtoV1Response.ProtoReflect.Descriptor instead.
func (*DescribeHowtoV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeHowtoV1Response) GetHowto() *Howto {
	if x != nil {
		return x.Howto
	}
	return nil
}

type ListHowtosV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartWithId uint64 `protobuf:"varint,1,opt,name=start_with_id,json=startWithId,proto3" json:"start_with_id,omitempty"`
	Length      uint64 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *ListHowtosV1Request) Reset() {
	*x = ListHowtosV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHowtosV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHowtosV1Request) ProtoMessage() {}

func (x *ListHowtosV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHowtosV1Request.ProtoReflect.Descriptor instead.
func (*ListHowtosV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListHowtosV1Request) GetStartWithId() uint64 {
	if x != nil {
		return x.StartWithId
	}
	return 0
}

func (x *ListHowtosV1Request) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ListHowtosV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Howtos []*Howto `protobuf:"bytes,1,rep,name=howtos,proto3" json:"howtos,omitempty"`
}

func (x *ListHowtosV1Response) Reset() {
	*x = ListHowtosV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHowtosV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHowtosV1Response) ProtoMessage() {}

func (x *ListHowtosV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHowtosV1Response.ProtoReflect.Descriptor instead.
func (*ListHowtosV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListHowtosV1Response) GetHowtos() []*Howto {
	if x != nil {
		return x.Howtos
	}
	return nil
}

type RemoveHowtoV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveHowtoV1Request) Reset() {
	*x = RemoveHowtoV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveHowtoV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveHowtoV1Request) ProtoMessage() {}

func (x *RemoveHowtoV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveHowtoV1Request.ProtoReflect.Descriptor instead.
func (*RemoveHowtoV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveHowtoV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveHowtoV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveHowtoV1Response) Reset() {
	*x = RemoveHowtoV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveHowtoV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveHowtoV1Response) ProtoMessage() {}

func (x *RemoveHowtoV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveHowtoV1Response.ProtoReflect.Descriptor instead.
func (*RemoveHowtoV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{7}
}

// Описание сущности howto
type Howto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CourseId uint64 `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	Answer   string `protobuf:"bytes,4,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *Howto) Reset() {
	*x = Howto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_howto_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Howto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Howto) ProtoMessage() {}

func (x *Howto) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_howto_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Howto.ProtoReflect.Descriptor instead.
func (*Howto) Descriptor() ([]byte, []int) {
	return file_ocp_howto_api_proto_rawDescGZIP(), []int{8}
}

func (x *Howto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Howto) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *Howto) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *Howto) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

var File_ocp_howto_api_proto protoreflect.FileDescriptor

var file_ocp_howto_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x63, 0x70, 0x2d, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x63, 0x70, 0x2e, 0x68, 0x6f, 0x77, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x67, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x77, 0x74,
	0x6f, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45,
	0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x68, 0x6f, 0x77,
	0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x68,
	0x6f, 0x77, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x52, 0x05,
	0x68, 0x6f, 0x77, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x77,
	0x74, 0x6f, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x6f, 0x77, 0x74, 0x6f, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x52, 0x06, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x73, 0x22, 0x26,
	0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x68, 0x0a, 0x05, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x32, 0xda, 0x03, 0x0a, 0x0b, 0x4f, 0x63,
	0x70, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x41, 0x70, 0x69, 0x12, 0x6e, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x0a, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x73, 0x12, 0x79, 0x0a, 0x0f, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x12, 0x25, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x48, 0x6f, 0x77, 0x74,
	0x6f, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x77, 0x74,
	0x6f, 0x73, 0x56, 0x31, 0x12, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x68, 0x6f, 0x77, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x68,
	0x6f, 0x77, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x77,
	0x74, 0x6f, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x77, 0x74, 0x6f,
	0x73, 0x12, 0x73, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f,
	0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f, 0x77, 0x74, 0x6f, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x68, 0x6f,
	0x77, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x6f,
	0x77, 0x74, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x77, 0x74, 0x6f,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63, 0x70, 0x2d,
	0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x63,
	0x70, 0x2d, 0x68, 0x6f, 0x77, 0x74, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f,
	0x68, 0x6f, 0x77, 0x74, 0x6f, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ocp_howto_api_proto_rawDescOnce sync.Once
	file_ocp_howto_api_proto_rawDescData = file_ocp_howto_api_proto_rawDesc
)

func file_ocp_howto_api_proto_rawDescGZIP() []byte {
	file_ocp_howto_api_proto_rawDescOnce.Do(func() {
		file_ocp_howto_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocp_howto_api_proto_rawDescData)
	})
	return file_ocp_howto_api_proto_rawDescData
}

var file_ocp_howto_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ocp_howto_api_proto_goTypes = []interface{}{
	(*CreateHowtoV1Request)(nil),    // 0: ocp.howto.api.CreateHowtoV1Request
	(*CreateHowtoV1Response)(nil),   // 1: ocp.howto.api.CreateHowtoV1Response
	(*DescribeHowtoV1Request)(nil),  // 2: ocp.howto.api.DescribeHowtoV1Request
	(*DescribeHowtoV1Response)(nil), // 3: ocp.howto.api.DescribeHowtoV1Response
	(*ListHowtosV1Request)(nil),     // 4: ocp.howto.api.ListHowtosV1Request
	(*ListHowtosV1Response)(nil),    // 5: ocp.howto.api.ListHowtosV1Response
	(*RemoveHowtoV1Request)(nil),    // 6: ocp.howto.api.RemoveHowtoV1Request
	(*RemoveHowtoV1Response)(nil),   // 7: ocp.howto.api.RemoveHowtoV1Response
	(*Howto)(nil),                   // 8: ocp.howto.api.Howto
}
var file_ocp_howto_api_proto_depIdxs = []int32{
	8, // 0: ocp.howto.api.DescribeHowtoV1Response.howto:type_name -> ocp.howto.api.Howto
	8, // 1: ocp.howto.api.ListHowtosV1Response.howtos:type_name -> ocp.howto.api.Howto
	0, // 2: ocp.howto.api.OcpHowtoApi.CreateHowtoV1:input_type -> ocp.howto.api.CreateHowtoV1Request
	2, // 3: ocp.howto.api.OcpHowtoApi.DescribeHowtoV1:input_type -> ocp.howto.api.DescribeHowtoV1Request
	4, // 4: ocp.howto.api.OcpHowtoApi.ListHowtosV1:input_type -> ocp.howto.api.ListHowtosV1Request
	6, // 5: ocp.howto.api.OcpHowtoApi.RemoveHowtoV1:input_type -> ocp.howto.api.RemoveHowtoV1Request
	1, // 6: ocp.howto.api.OcpHowtoApi.CreateHowtoV1:output_type -> ocp.howto.api.CreateHowtoV1Response
	3, // 7: ocp.howto.api.OcpHowtoApi.DescribeHowtoV1:output_type -> ocp.howto.api.DescribeHowtoV1Response
	5, // 8: ocp.howto.api.OcpHowtoApi.ListHowtosV1:output_type -> ocp.howto.api.ListHowtosV1Response
	7, // 9: ocp.howto.api.OcpHowtoApi.RemoveHowtoV1:output_type -> ocp.howto.api.RemoveHowtoV1Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ocp_howto_api_proto_init() }
func file_ocp_howto_api_proto_init() {
	if File_ocp_howto_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocp_howto_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHowtoV1Request); i {
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
		file_ocp_howto_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHowtoV1Response); i {
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
		file_ocp_howto_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeHowtoV1Request); i {
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
		file_ocp_howto_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeHowtoV1Response); i {
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
		file_ocp_howto_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHowtosV1Request); i {
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
		file_ocp_howto_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHowtosV1Response); i {
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
		file_ocp_howto_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveHowtoV1Request); i {
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
		file_ocp_howto_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveHowtoV1Response); i {
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
		file_ocp_howto_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Howto); i {
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
			RawDescriptor: file_ocp_howto_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ocp_howto_api_proto_goTypes,
		DependencyIndexes: file_ocp_howto_api_proto_depIdxs,
		MessageInfos:      file_ocp_howto_api_proto_msgTypes,
	}.Build()
	File_ocp_howto_api_proto = out.File
	file_ocp_howto_api_proto_rawDesc = nil
	file_ocp_howto_api_proto_goTypes = nil
	file_ocp_howto_api_proto_depIdxs = nil
}