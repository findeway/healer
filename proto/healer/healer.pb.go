// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/healer/healer.proto

package healer

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

type HealerSingleData_MediaType int32

const (
	HealerSingleData_Video HealerSingleData_MediaType = 0
	HealerSingleData_Image HealerSingleData_MediaType = 1
	HealerSingleData_Other HealerSingleData_MediaType = 2
)

// Enum value maps for HealerSingleData_MediaType.
var (
	HealerSingleData_MediaType_name = map[int32]string{
		0: "Video",
		1: "Image",
		2: "Other",
	}
	HealerSingleData_MediaType_value = map[string]int32{
		"Video": 0,
		"Image": 1,
		"Other": 2,
	}
)

func (x HealerSingleData_MediaType) Enum() *HealerSingleData_MediaType {
	p := new(HealerSingleData_MediaType)
	*p = x
	return p
}

func (x HealerSingleData_MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealerSingleData_MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_healer_healer_proto_enumTypes[0].Descriptor()
}

func (HealerSingleData_MediaType) Type() protoreflect.EnumType {
	return &file_proto_healer_healer_proto_enumTypes[0]
}

func (x HealerSingleData_MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealerSingleData_MediaType.Descriptor instead.
func (HealerSingleData_MediaType) EnumDescriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{6, 0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Say string `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSay() string {
	if x != nil {
		return x.Say
	}
	return ""
}

type CategoryInnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category []*CategoryData `protobuf:"bytes,1,rep,name=category,proto3" json:"category,omitempty"`
}

func (x *CategoryInnerResponse) Reset() {
	*x = CategoryInnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryInnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryInnerResponse) ProtoMessage() {}

func (x *CategoryInnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryInnerResponse.ProtoReflect.Descriptor instead.
func (*CategoryInnerResponse) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryInnerResponse) GetCategory() []*CategoryData {
	if x != nil {
		return x.Category
	}
	return nil
}

type CateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *CategoryInnerResponse `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CateResponse) Reset() {
	*x = CateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CateResponse) ProtoMessage() {}

func (x *CateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CateResponse.ProtoReflect.Descriptor instead.
func (*CateResponse) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{2}
}

func (x *CateResponse) GetData() *CategoryInnerResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type CategoryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *CategoryData) Reset() {
	*x = CategoryData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryData) ProtoMessage() {}

func (x *CategoryData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryData.ProtoReflect.Descriptor instead.
func (*CategoryData) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{3}
}

func (x *CategoryData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryData) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type HealerInnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HealerData []*HealerGroupData `protobuf:"bytes,1,rep,name=healer_data,json=healerData,proto3" json:"healer_data,omitempty"`
}

func (x *HealerInnerResponse) Reset() {
	*x = HealerInnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealerInnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealerInnerResponse) ProtoMessage() {}

func (x *HealerInnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealerInnerResponse.ProtoReflect.Descriptor instead.
func (*HealerInnerResponse) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{4}
}

func (x *HealerInnerResponse) GetHealerData() []*HealerGroupData {
	if x != nil {
		return x.HealerData
	}
	return nil
}

type HealerGroupData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag   string              `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Items []*HealerSingleData `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *HealerGroupData) Reset() {
	*x = HealerGroupData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealerGroupData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealerGroupData) ProtoMessage() {}

func (x *HealerGroupData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealerGroupData.ProtoReflect.Descriptor instead.
func (*HealerGroupData) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{5}
}

func (x *HealerGroupData) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *HealerGroupData) GetItems() []*HealerSingleData {
	if x != nil {
		return x.Items
	}
	return nil
}

type HealerSingleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url      string                     `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	CoverUrl string                     `protobuf:"bytes,3,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	Type     HealerSingleData_MediaType `protobuf:"varint,4,opt,name=type,proto3,enum=go.micro.service.healer.HealerSingleData_MediaType" json:"type,omitempty"`
	Desc     string                     `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *HealerSingleData) Reset() {
	*x = HealerSingleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealerSingleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealerSingleData) ProtoMessage() {}

func (x *HealerSingleData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealerSingleData.ProtoReflect.Descriptor instead.
func (*HealerSingleData) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{6}
}

func (x *HealerSingleData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HealerSingleData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HealerSingleData) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *HealerSingleData) GetType() HealerSingleData_MediaType {
	if x != nil {
		return x.Type
	}
	return HealerSingleData_Video
}

func (x *HealerSingleData) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{7}
}

func (x *CallRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HealResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *HealerInnerResponse `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *HealResponse) Reset() {
	*x = HealResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealResponse) ProtoMessage() {}

func (x *HealResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealResponse.ProtoReflect.Descriptor instead.
func (*HealResponse) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{8}
}

func (x *HealResponse) GetData() *HealerInnerResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type StreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StreamingRequest) Reset() {
	*x = StreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingRequest) ProtoMessage() {}

func (x *StreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingRequest.ProtoReflect.Descriptor instead.
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{9}
}

func (x *StreamingRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StreamingResponse) Reset() {
	*x = StreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingResponse) ProtoMessage() {}

func (x *StreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingResponse.ProtoReflect.Descriptor instead.
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{10}
}

func (x *StreamingResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{11}
}

func (x *Ping) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stroke int64 `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_healer_healer_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_proto_healer_healer_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_proto_healer_healer_proto_rawDescGZIP(), []int{12}
}

func (x *Pong) GetStroke() int64 {
	if x != nil {
		return x.Stroke
	}
	return 0
}

var File_proto_healer_healer_proto protoreflect.FileDescriptor

var file_proto_healer_healer_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x65, 0x72, 0x22, 0x1b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x61,
	0x79, 0x22, 0x5a, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x52, 0x0a,
	0x0c, 0x43, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x36, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x60, 0x0a, 0x13, 0x48, 0x65, 0x61,
	0x6c, 0x65, 0x72, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0a, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x0f, 0x48,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x12, 0x3f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xe0, 0x01, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x2c, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x10, 0x02, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1e,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x22, 0x1e,
	0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x72, 0x6f, 0x6b, 0x65, 0x32, 0xf5,
	0x02, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x08, 0x48, 0x65, 0x61,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x63, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x29, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x6e, 0x67,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x3b, 0x68, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_healer_healer_proto_rawDescOnce sync.Once
	file_proto_healer_healer_proto_rawDescData = file_proto_healer_healer_proto_rawDesc
)

func file_proto_healer_healer_proto_rawDescGZIP() []byte {
	file_proto_healer_healer_proto_rawDescOnce.Do(func() {
		file_proto_healer_healer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_healer_healer_proto_rawDescData)
	})
	return file_proto_healer_healer_proto_rawDescData
}

var file_proto_healer_healer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_healer_healer_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_healer_healer_proto_goTypes = []interface{}{
	(HealerSingleData_MediaType)(0), // 0: go.micro.service.healer.HealerSingleData.MediaType
	(*Message)(nil),                 // 1: go.micro.service.healer.Message
	(*CategoryInnerResponse)(nil),   // 2: go.micro.service.healer.CategoryInnerResponse
	(*CateResponse)(nil),            // 3: go.micro.service.healer.CateResponse
	(*CategoryData)(nil),            // 4: go.micro.service.healer.CategoryData
	(*HealerInnerResponse)(nil),     // 5: go.micro.service.healer.HealerInnerResponse
	(*HealerGroupData)(nil),         // 6: go.micro.service.healer.HealerGroupData
	(*HealerSingleData)(nil),        // 7: go.micro.service.healer.HealerSingleData
	(*CallRequest)(nil),             // 8: go.micro.service.healer.CallRequest
	(*HealResponse)(nil),            // 9: go.micro.service.healer.HealResponse
	(*StreamingRequest)(nil),        // 10: go.micro.service.healer.StreamingRequest
	(*StreamingResponse)(nil),       // 11: go.micro.service.healer.StreamingResponse
	(*Ping)(nil),                    // 12: go.micro.service.healer.Ping
	(*Pong)(nil),                    // 13: go.micro.service.healer.Pong
}
var file_proto_healer_healer_proto_depIdxs = []int32{
	4,  // 0: go.micro.service.healer.CategoryInnerResponse.category:type_name -> go.micro.service.healer.CategoryData
	2,  // 1: go.micro.service.healer.CateResponse.data:type_name -> go.micro.service.healer.CategoryInnerResponse
	6,  // 2: go.micro.service.healer.HealerInnerResponse.healer_data:type_name -> go.micro.service.healer.HealerGroupData
	7,  // 3: go.micro.service.healer.HealerGroupData.items:type_name -> go.micro.service.healer.HealerSingleData
	0,  // 4: go.micro.service.healer.HealerSingleData.type:type_name -> go.micro.service.healer.HealerSingleData.MediaType
	5,  // 5: go.micro.service.healer.HealResponse.data:type_name -> go.micro.service.healer.HealerInnerResponse
	8,  // 6: go.micro.service.healer.Healer.HealList:input_type -> go.micro.service.healer.CallRequest
	8,  // 7: go.micro.service.healer.Healer.Categories:input_type -> go.micro.service.healer.CallRequest
	10, // 8: go.micro.service.healer.Healer.Stream:input_type -> go.micro.service.healer.StreamingRequest
	12, // 9: go.micro.service.healer.Healer.PingPong:input_type -> go.micro.service.healer.Ping
	9,  // 10: go.micro.service.healer.Healer.HealList:output_type -> go.micro.service.healer.HealResponse
	3,  // 11: go.micro.service.healer.Healer.Categories:output_type -> go.micro.service.healer.CateResponse
	11, // 12: go.micro.service.healer.Healer.Stream:output_type -> go.micro.service.healer.StreamingResponse
	13, // 13: go.micro.service.healer.Healer.PingPong:output_type -> go.micro.service.healer.Pong
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_healer_healer_proto_init() }
func file_proto_healer_healer_proto_init() {
	if File_proto_healer_healer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_healer_healer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_healer_healer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryInnerResponse); i {
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
		file_proto_healer_healer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CateResponse); i {
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
		file_proto_healer_healer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryData); i {
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
		file_proto_healer_healer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealerInnerResponse); i {
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
		file_proto_healer_healer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealerGroupData); i {
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
		file_proto_healer_healer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealerSingleData); i {
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
		file_proto_healer_healer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_proto_healer_healer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealResponse); i {
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
		file_proto_healer_healer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingRequest); i {
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
		file_proto_healer_healer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingResponse); i {
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
		file_proto_healer_healer_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_proto_healer_healer_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_proto_healer_healer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_healer_healer_proto_goTypes,
		DependencyIndexes: file_proto_healer_healer_proto_depIdxs,
		EnumInfos:         file_proto_healer_healer_proto_enumTypes,
		MessageInfos:      file_proto_healer_healer_proto_msgTypes,
	}.Build()
	File_proto_healer_healer_proto = out.File
	file_proto_healer_healer_proto_rawDesc = nil
	file_proto_healer_healer_proto_goTypes = nil
	file_proto_healer_healer_proto_depIdxs = nil
}
