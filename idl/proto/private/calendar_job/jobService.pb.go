// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: idl/proto/private/calendar_job/jobService.proto

package private_pd

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

// 创建一个job
type CreateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     uint64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	RecordID   uint64  `protobuf:"varint,2,opt,name=RecordID,proto3" json:"RecordID,omitempty"`
	Title      string  `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	Content    string  `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	NoticeTime string  `protobuf:"bytes,5,opt,name=NoticeTime,proto3" json:"NoticeTime,omitempty"`
	Phone      *string `protobuf:"bytes,6,opt,name=Phone,proto3,oneof" json:"Phone,omitempty"`
	Email      string  `protobuf:"bytes,7,opt,name=Email,proto3" json:"Email,omitempty"`
	WeChat     *string `protobuf:"bytes,8,opt,name=WeChat,proto3,oneof" json:"WeChat,omitempty"`
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest.ProtoReflect.Descriptor instead.
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJobRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateJobRequest) GetRecordID() uint64 {
	if x != nil {
		return x.RecordID
	}
	return 0
}

func (x *CreateJobRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateJobRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateJobRequest) GetNoticeTime() string {
	if x != nil {
		return x.NoticeTime
	}
	return ""
}

func (x *CreateJobRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *CreateJobRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateJobRequest) GetWeChat() string {
	if x != nil && x.WeChat != nil {
		return *x.WeChat
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID  uint64 `protobuf:"varint,1,opt,name=JobID,proto3" json:"JobID,omitempty"`
	Status uint32 `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[1]
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
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetJobID() uint64 {
	if x != nil {
		return x.JobID
	}
	return 0
}

func (x *Message) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data *Message `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *CreateJobResponse) Reset() {
	*x = CreateJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobResponse) ProtoMessage() {}

func (x *CreateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobResponse.ProtoReflect.Descriptor instead.
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{2}
}

func (x *CreateJobResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateJobResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateJobResponse) GetData() *Message {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetJobListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Status   uint32 `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	RecordID uint64 `protobuf:"varint,3,opt,name=RecordID,proto3" json:"RecordID,omitempty"`
}

func (x *GetJobListRequest) Reset() {
	*x = GetJobListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobListRequest) ProtoMessage() {}

func (x *GetJobListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobListRequest.ProtoReflect.Descriptor instead.
func (*GetJobListRequest) Descriptor() ([]byte, []int) {
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{3}
}

func (x *GetJobListRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetJobListRequest) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetJobListRequest) GetRecordID() uint64 {
	if x != nil {
		return x.RecordID
	}
	return 0
}

type GetJobListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32      `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string      `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data []*JobModel `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetJobListResponse) Reset() {
	*x = GetJobListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobListResponse) ProtoMessage() {}

func (x *GetJobListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobListResponse.ProtoReflect.Descriptor instead.
func (*GetJobListResponse) Descriptor() ([]byte, []int) {
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{4}
}

func (x *GetJobListResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetJobListResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetJobListResponse) GetData() []*JobModel {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetJobDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobID uint64 `protobuf:"varint,1,opt,name=JobID,proto3" json:"JobID,omitempty"`
}

func (x *GetJobDetailRequest) Reset() {
	*x = GetJobDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobDetailRequest) ProtoMessage() {}

func (x *GetJobDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobDetailRequest.ProtoReflect.Descriptor instead.
func (*GetJobDetailRequest) Descriptor() ([]byte, []int) {
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{5}
}

func (x *GetJobDetailRequest) GetJobID() uint64 {
	if x != nil {
		return x.JobID
	}
	return 0
}

type GetJobDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string    `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data *JobModel `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetJobDetailResponse) Reset() {
	*x = GetJobDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobDetailResponse) ProtoMessage() {}

func (x *GetJobDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobDetailResponse.ProtoReflect.Descriptor instead.
func (*GetJobDetailResponse) Descriptor() ([]byte, []int) {
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{6}
}

func (x *GetJobDetailResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetJobDetailResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetJobDetailResponse) GetData() *JobModel {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordID   uint64  `protobuf:"varint,1,opt,name=RecordID,proto3" json:"RecordID,omitempty"`
	NoticeTime *string `protobuf:"bytes,5,opt,name=NoticeTime,proto3,oneof" json:"NoticeTime,omitempty"`
	UserID     *uint64 `protobuf:"varint,2,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
}

func (x *DeleteJobRequest) Reset() {
	*x = DeleteJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobRequest) Descriptor() ([]byte, []int) {
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteJobRequest) GetRecordID() uint64 {
	if x != nil {
		return x.RecordID
	}
	return 0
}

func (x *DeleteJobRequest) GetNoticeTime() string {
	if x != nil && x.NoticeTime != nil {
		return *x.NoticeTime
	}
	return ""
}

func (x *DeleteJobRequest) GetUserID() uint64 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_private_calendar_job_jobService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP(), []int{8}
}

func (x *CommonResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CommonResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_idl_proto_private_calendar_job_jobService_proto protoreflect.FileDescriptor

var file_idl_proto_private_calendar_job_jobService_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x6a, 0x6f, 0x62,
	0x2f, 0x6a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x06, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x57, 0x65, 0x43, 0x68, 0x61, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x57, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x22, 0x37, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x4a, 0x6f, 0x62, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5d, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5f, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x22, 0x5f, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4a, 0x6f, 0x62, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2b,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4a, 0x6f, 0x62, 0x49, 0x44, 0x22, 0x61, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x23, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4a, 0x6f, 0x62, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x8a,
	0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x12,
	0x23, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4a, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0x95, 0x02, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x23, 0x5a, 0x21, 0x2e, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x3b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_proto_private_calendar_job_jobService_proto_rawDescOnce sync.Once
	file_idl_proto_private_calendar_job_jobService_proto_rawDescData = file_idl_proto_private_calendar_job_jobService_proto_rawDesc
)

func file_idl_proto_private_calendar_job_jobService_proto_rawDescGZIP() []byte {
	file_idl_proto_private_calendar_job_jobService_proto_rawDescOnce.Do(func() {
		file_idl_proto_private_calendar_job_jobService_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_proto_private_calendar_job_jobService_proto_rawDescData)
	})
	return file_idl_proto_private_calendar_job_jobService_proto_rawDescData
}

var file_idl_proto_private_calendar_job_jobService_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_idl_proto_private_calendar_job_jobService_proto_goTypes = []interface{}{
	(*CreateJobRequest)(nil),     // 0: proto.CreateJobRequest
	(*Message)(nil),              // 1: proto.Message
	(*CreateJobResponse)(nil),    // 2: proto.CreateJobResponse
	(*GetJobListRequest)(nil),    // 3: proto.GetJobListRequest
	(*GetJobListResponse)(nil),   // 4: proto.GetJobListResponse
	(*GetJobDetailRequest)(nil),  // 5: proto.GetJobDetailRequest
	(*GetJobDetailResponse)(nil), // 6: proto.GetJobDetailResponse
	(*DeleteJobRequest)(nil),     // 7: proto.DeleteJobRequest
	(*CommonResponse)(nil),       // 8: proto.CommonResponse
	(*JobModel)(nil),             // 9: proto.JobModel
}
var file_idl_proto_private_calendar_job_jobService_proto_depIdxs = []int32{
	1, // 0: proto.CreateJobResponse.Data:type_name -> proto.Message
	9, // 1: proto.GetJobListResponse.Data:type_name -> proto.JobModel
	9, // 2: proto.GetJobDetailResponse.Data:type_name -> proto.JobModel
	0, // 3: proto.JobService.JobCreate:input_type -> proto.CreateJobRequest
	3, // 4: proto.JobService.GetJobList:input_type -> proto.GetJobListRequest
	5, // 5: proto.JobService.GetJobDetail:input_type -> proto.GetJobDetailRequest
	7, // 6: proto.JobService.DeleteJob:input_type -> proto.DeleteJobRequest
	2, // 7: proto.JobService.JobCreate:output_type -> proto.CreateJobResponse
	4, // 8: proto.JobService.GetJobList:output_type -> proto.GetJobListResponse
	6, // 9: proto.JobService.GetJobDetail:output_type -> proto.GetJobDetailResponse
	8, // 10: proto.JobService.DeleteJob:output_type -> proto.CommonResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_idl_proto_private_calendar_job_jobService_proto_init() }
func file_idl_proto_private_calendar_job_jobService_proto_init() {
	if File_idl_proto_private_calendar_job_jobService_proto != nil {
		return
	}
	file_idl_proto_private_calendar_job_jobModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobRequest); i {
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
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobResponse); i {
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
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobListRequest); i {
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
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobListResponse); i {
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
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobDetailRequest); i {
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
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobDetailResponse); i {
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
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobRequest); i {
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
		file_idl_proto_private_calendar_job_jobService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
	file_idl_proto_private_calendar_job_jobService_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_idl_proto_private_calendar_job_jobService_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_proto_private_calendar_job_jobService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_proto_private_calendar_job_jobService_proto_goTypes,
		DependencyIndexes: file_idl_proto_private_calendar_job_jobService_proto_depIdxs,
		MessageInfos:      file_idl_proto_private_calendar_job_jobService_proto_msgTypes,
	}.Build()
	File_idl_proto_private_calendar_job_jobService_proto = out.File
	file_idl_proto_private_calendar_job_jobService_proto_rawDesc = nil
	file_idl_proto_private_calendar_job_jobService_proto_goTypes = nil
	file_idl_proto_private_calendar_job_jobService_proto_depIdxs = nil
}
