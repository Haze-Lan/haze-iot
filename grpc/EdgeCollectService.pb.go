// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.6
// source: EdgeCollectService.proto

package grpc

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

type DataCollections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32          `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Data []*DataRequest `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DataCollections) Reset() {
	*x = DataCollections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataCollections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataCollections) ProtoMessage() {}

func (x *DataCollections) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataCollections.ProtoReflect.Descriptor instead.
func (*DataCollections) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{0}
}

func (x *DataCollections) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *DataCollections) GetData() []*DataRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

//检测数据
type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//设备编号
	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	//环境项目
	Variable string `protobuf:"bytes,2,opt,name=variable,proto3" json:"variable,omitempty"`
	//环境项目值
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	//数据发生时间
	Date int64 `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{1}
}

func (x *DataRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *DataRequest) GetVariable() string {
	if x != nil {
		return x.Variable
	}
	return ""
}

func (x *DataRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *DataRequest) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

type StateCollections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32           `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Data []*StateRequest `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *StateCollections) Reset() {
	*x = StateCollections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateCollections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateCollections) ProtoMessage() {}

func (x *StateCollections) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateCollections.ProtoReflect.Descriptor instead.
func (*StateCollections) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{2}
}

func (x *StateCollections) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StateCollections) GetData() []*StateRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

//设备连接状态
type StateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//设备编号
	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	//设备状态
	State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	//携带的消息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	//时间发生时间
	Date int64 `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *StateRequest) Reset() {
	*x = StateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateRequest) ProtoMessage() {}

func (x *StateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateRequest.ProtoReflect.Descriptor instead.
func (*StateRequest) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{3}
}

func (x *StateRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *StateRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *StateRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StateRequest) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

//指令执行结果上报
type CommandStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//设备编号
	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	//指令Key
	CommandKey string `protobuf:"bytes,2,opt,name=commandKey,proto3" json:"commandKey,omitempty"`
	//是否成功
	IsSuccess bool `protobuf:"varint,3,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	//执行结果描述
	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommandStateRequest) Reset() {
	*x = CommandStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandStateRequest) ProtoMessage() {}

func (x *CommandStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandStateRequest.ProtoReflect.Descriptor instead.
func (*CommandStateRequest) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{4}
}

func (x *CommandStateRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CommandStateRequest) GetCommandKey() string {
	if x != nil {
		return x.CommandKey
	}
	return ""
}

func (x *CommandStateRequest) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *CommandStateRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AttributesCollections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32                `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Data []*AttributesRequest `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AttributesCollections) Reset() {
	*x = AttributesCollections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributesCollections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributesCollections) ProtoMessage() {}

func (x *AttributesCollections) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributesCollections.ProtoReflect.Descriptor instead.
func (*AttributesCollections) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{5}
}

func (x *AttributesCollections) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *AttributesCollections) GetData() []*AttributesRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

//指令执行结果上报
type AttributesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//设备编号
	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	//属性键名
	Attributes string `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	//旧值
	OldValue string `protobuf:"bytes,3,opt,name=oldValue,proto3" json:"oldValue,omitempty"`
	//新值
	NewValue string `protobuf:"bytes,4,opt,name=newValue,proto3" json:"newValue,omitempty"`
	//数据发生时间
	Date int64 `protobuf:"varint,5,opt,name=date,proto3" json:"date,omitempty"`
	//中文 名称
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AttributesRequest) Reset() {
	*x = AttributesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributesRequest) ProtoMessage() {}

func (x *AttributesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributesRequest.ProtoReflect.Descriptor instead.
func (*AttributesRequest) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{6}
}

func (x *AttributesRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *AttributesRequest) GetAttributes() string {
	if x != nil {
		return x.Attributes
	}
	return ""
}

func (x *AttributesRequest) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

func (x *AttributesRequest) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

func (x *AttributesRequest) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *AttributesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MetricsCollections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32             `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Data []*MetricsRequest `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MetricsCollections) Reset() {
	*x = MetricsCollections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsCollections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsCollections) ProtoMessage() {}

func (x *MetricsCollections) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsCollections.ProtoReflect.Descriptor instead.
func (*MetricsCollections) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{7}
}

func (x *MetricsCollections) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *MetricsCollections) GetData() []*MetricsRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

//监测参数上报
type MetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// 名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//*
	// 唯一标识
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	//*
	// 值
	Value float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	//*
	// 描述
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	//*
	// 单位
	Unit string `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *MetricsRequest) Reset() {
	*x = MetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsRequest) ProtoMessage() {}

func (x *MetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsRequest.ProtoReflect.Descriptor instead.
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{8}
}

func (x *MetricsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricsRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MetricsRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *MetricsRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MetricsRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

//这个没什么用
type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{9}
}

type ReportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//这个是否ok绝对数据是否从新发送
	IsOk bool `protobuf:"varint,1,opt,name=isOk,proto3" json:"isOk,omitempty"`
}

func (x *ReportReply) Reset() {
	*x = ReportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportReply) ProtoMessage() {}

func (x *ReportReply) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportReply.ProtoReflect.Descriptor instead.
func (*ReportReply) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{10}
}

func (x *ReportReply) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//设备编号
	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	//指令Key
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	//指令携带参数
	Parameters map[string]string `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EdgeCollectService_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EdgeCollectService_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_EdgeCollectService_proto_rawDescGZIP(), []int{11}
}

func (x *CommandResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CommandResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CommandResponse) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

var File_EdgeCollectService_proto protoreflect.FileDescriptor

var file_EdgeCollectService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x22, 0x4c, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6b,
	0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6a, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x58, 0x0a, 0x15, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x82, 0x01, 0x0a, 0x0e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x69, 0x73, 0x4f, 0x6b, 0x22, 0xc1, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xdb, 0x02, 0x0a, 0x12, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0c, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x13, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x66, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x11, 0x50, 0x01, 0x5a, 0x0d, 0x68, 0x61, 0x7a, 0x65, 0x2d, 0x69, 0x6f, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EdgeCollectService_proto_rawDescOnce sync.Once
	file_EdgeCollectService_proto_rawDescData = file_EdgeCollectService_proto_rawDesc
)

func file_EdgeCollectService_proto_rawDescGZIP() []byte {
	file_EdgeCollectService_proto_rawDescOnce.Do(func() {
		file_EdgeCollectService_proto_rawDescData = protoimpl.X.CompressGZIP(file_EdgeCollectService_proto_rawDescData)
	})
	return file_EdgeCollectService_proto_rawDescData
}

var file_EdgeCollectService_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_EdgeCollectService_proto_goTypes = []interface{}{
	(*DataCollections)(nil),       // 0: grpc.DataCollections
	(*DataRequest)(nil),           // 1: grpc.DataRequest
	(*StateCollections)(nil),      // 2: grpc.StateCollections
	(*StateRequest)(nil),          // 3: grpc.StateRequest
	(*CommandStateRequest)(nil),   // 4: grpc.CommandStateRequest
	(*AttributesCollections)(nil), // 5: grpc.AttributesCollections
	(*AttributesRequest)(nil),     // 6: grpc.AttributesRequest
	(*MetricsCollections)(nil),    // 7: grpc.MetricsCollections
	(*MetricsRequest)(nil),        // 8: grpc.MetricsRequest
	(*CommandRequest)(nil),        // 9: grpc.CommandRequest
	(*ReportReply)(nil),           // 10: grpc.ReportReply
	(*CommandResponse)(nil),       // 11: grpc.CommandResponse
	nil,                           // 12: grpc.CommandResponse.ParametersEntry
}
var file_EdgeCollectService_proto_depIdxs = []int32{
	1,  // 0: grpc.DataCollections.data:type_name -> grpc.DataRequest
	3,  // 1: grpc.StateCollections.data:type_name -> grpc.StateRequest
	6,  // 2: grpc.AttributesCollections.data:type_name -> grpc.AttributesRequest
	8,  // 3: grpc.MetricsCollections.data:type_name -> grpc.MetricsRequest
	12, // 4: grpc.CommandResponse.parameters:type_name -> grpc.CommandResponse.ParametersEntry
	0,  // 5: grpc.DataCollectService.collectData:input_type -> grpc.DataCollections
	2,  // 6: grpc.DataCollectService.collectState:input_type -> grpc.StateCollections
	4,  // 7: grpc.DataCollectService.collectCommandState:input_type -> grpc.CommandStateRequest
	5,  // 8: grpc.DataCollectService.collectAttributes:input_type -> grpc.AttributesCollections
	7,  // 9: grpc.DataCollectService.collectMetrics:input_type -> grpc.MetricsCollections
	9,  // 10: grpc.CommandDistributionService.commandDistribution:input_type -> grpc.CommandRequest
	10, // 11: grpc.DataCollectService.collectData:output_type -> grpc.ReportReply
	10, // 12: grpc.DataCollectService.collectState:output_type -> grpc.ReportReply
	10, // 13: grpc.DataCollectService.collectCommandState:output_type -> grpc.ReportReply
	10, // 14: grpc.DataCollectService.collectAttributes:output_type -> grpc.ReportReply
	10, // 15: grpc.DataCollectService.collectMetrics:output_type -> grpc.ReportReply
	11, // 16: grpc.CommandDistributionService.commandDistribution:output_type -> grpc.CommandResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_EdgeCollectService_proto_init() }
func file_EdgeCollectService_proto_init() {
	if File_EdgeCollectService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EdgeCollectService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataCollections); i {
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
		file_EdgeCollectService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_EdgeCollectService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateCollections); i {
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
		file_EdgeCollectService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateRequest); i {
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
		file_EdgeCollectService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandStateRequest); i {
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
		file_EdgeCollectService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributesCollections); i {
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
		file_EdgeCollectService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributesRequest); i {
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
		file_EdgeCollectService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsCollections); i {
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
		file_EdgeCollectService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsRequest); i {
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
		file_EdgeCollectService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_EdgeCollectService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportReply); i {
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
		file_EdgeCollectService_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
			RawDescriptor: file_EdgeCollectService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_EdgeCollectService_proto_goTypes,
		DependencyIndexes: file_EdgeCollectService_proto_depIdxs,
		MessageInfos:      file_EdgeCollectService_proto_msgTypes,
	}.Build()
	File_EdgeCollectService_proto = out.File
	file_EdgeCollectService_proto_rawDesc = nil
	file_EdgeCollectService_proto_goTypes = nil
	file_EdgeCollectService_proto_depIdxs = nil
}
