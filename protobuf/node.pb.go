// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: protobuf/node.proto

package avsproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageOp int32

const (
	MessageOp_Unset              MessageOp = 0
	MessageOp_MonitorTaskTrigger MessageOp = 1
	MessageOp_CancelTask         MessageOp = 2
	MessageOp_DeleteTask         MessageOp = 3
	MessageOp_CompletedTask      MessageOp = 4
)

// Enum value maps for MessageOp.
var (
	MessageOp_name = map[int32]string{
		0: "Unset",
		1: "MonitorTaskTrigger",
		2: "CancelTask",
		3: "DeleteTask",
		4: "CompletedTask",
	}
	MessageOp_value = map[string]int32{
		"Unset":              0,
		"MonitorTaskTrigger": 1,
		"CancelTask":         2,
		"DeleteTask":         3,
		"CompletedTask":      4,
	}
)

func (x MessageOp) Enum() *MessageOp {
	p := new(MessageOp)
	*p = x
	return p
}

func (x MessageOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageOp) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_node_proto_enumTypes[0].Descriptor()
}

func (MessageOp) Type() protoreflect.EnumType {
	return &file_protobuf_node_proto_enumTypes[0]
}

func (x MessageOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageOp.Descriptor instead.
func (MessageOp) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{0}
}

type Checkin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address     string          `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Signature   string          `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Status      *Checkin_Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Version     string          `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	MetricsPort int32           `protobuf:"varint,6,opt,name=metricsPort,proto3" json:"metricsPort,omitempty"`
	RemoteIP    string          `protobuf:"bytes,7,opt,name=remoteIP,proto3" json:"remoteIP,omitempty"`
}

func (x *Checkin) Reset() {
	*x = Checkin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkin) ProtoMessage() {}

func (x *Checkin) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkin.ProtoReflect.Descriptor instead.
func (*Checkin) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{0}
}

func (x *Checkin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Checkin) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Checkin) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Checkin) GetStatus() *Checkin_Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Checkin) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Checkin) GetMetricsPort() int32 {
	if x != nil {
		return x.MetricsPort
	}
	return 0
}

func (x *Checkin) GetRemoteIP() string {
	if x != nil {
		return x.RemoteIP
	}
	return ""
}

type CheckinResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CheckinResp) Reset() {
	*x = CheckinResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckinResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckinResp) ProtoMessage() {}

func (x *CheckinResp) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckinResp.ProtoReflect.Descriptor instead.
func (*CheckinResp) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{1}
}

func (x *CheckinResp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type SyncMessagesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address        string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Signature      []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	MonotonicClock int64  `protobuf:"varint,4,opt,name=monotonic_clock,json=monotonicClock,proto3" json:"monotonic_clock,omitempty"`
}

func (x *SyncMessagesReq) Reset() {
	*x = SyncMessagesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMessagesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMessagesReq) ProtoMessage() {}

func (x *SyncMessagesReq) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMessagesReq.ProtoReflect.Descriptor instead.
func (*SyncMessagesReq) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{2}
}

func (x *SyncMessagesReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SyncMessagesReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SyncMessagesReq) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SyncMessagesReq) GetMonotonicClock() int64 {
	if x != nil {
		return x.MonotonicClock
	}
	return 0
}

type SyncMessagesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// message id is used to support dedup
	Id           string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Op           MessageOp                      `protobuf:"varint,2,opt,name=op,proto3,enum=aggregator.MessageOp" json:"op,omitempty"`
	TaskMetadata *SyncMessagesResp_TaskMetadata `protobuf:"bytes,3,opt,name=task_metadata,json=taskMetadata,proto3" json:"task_metadata,omitempty"`
}

func (x *SyncMessagesResp) Reset() {
	*x = SyncMessagesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMessagesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMessagesResp) ProtoMessage() {}

func (x *SyncMessagesResp) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMessagesResp.ProtoReflect.Descriptor instead.
func (*SyncMessagesResp) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{3}
}

func (x *SyncMessagesResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SyncMessagesResp) GetOp() MessageOp {
	if x != nil {
		return x.Op
	}
	return MessageOp_Unset
}

func (x *SyncMessagesResp) GetTaskMetadata() *SyncMessagesResp_TaskMetadata {
	if x != nil {
		return x.TaskMetadata
	}
	return nil
}

type AckMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AckMessageReq) Reset() {
	*x = AckMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessageReq) ProtoMessage() {}

func (x *AckMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessageReq.ProtoReflect.Descriptor instead.
func (*AckMessageReq) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{4}
}

func (x *AckMessageReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NotifyTriggersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address         string           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Signature       string           `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	TaskId          string           `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TriggerMetadata *TriggerMetadata `protobuf:"bytes,4,opt,name=trigger_metadata,json=triggerMetadata,proto3" json:"trigger_metadata,omitempty"`
}

func (x *NotifyTriggersReq) Reset() {
	*x = NotifyTriggersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyTriggersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyTriggersReq) ProtoMessage() {}

func (x *NotifyTriggersReq) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyTriggersReq.ProtoReflect.Descriptor instead.
func (*NotifyTriggersReq) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{5}
}

func (x *NotifyTriggersReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *NotifyTriggersReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *NotifyTriggersReq) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *NotifyTriggersReq) GetTriggerMetadata() *TriggerMetadata {
	if x != nil {
		return x.TriggerMetadata
	}
	return nil
}

type NotifyTriggersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *NotifyTriggersResp) Reset() {
	*x = NotifyTriggersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyTriggersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyTriggersResp) ProtoMessage() {}

func (x *NotifyTriggersResp) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyTriggersResp.ProtoReflect.Descriptor instead.
func (*NotifyTriggersResp) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{6}
}

func (x *NotifyTriggersResp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Checkin_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uptime        int64                  `protobuf:"varint,1,opt,name=uptime,proto3" json:"uptime,omitempty"`
	QueueDepth    int64                  `protobuf:"varint,2,opt,name=queueDepth,proto3" json:"queueDepth,omitempty"`
	LastHeartbeat *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_heartbeat,json=lastHeartbeat,proto3" json:"last_heartbeat,omitempty"`
}

func (x *Checkin_Status) Reset() {
	*x = Checkin_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkin_Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkin_Status) ProtoMessage() {}

func (x *Checkin_Status) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkin_Status.ProtoReflect.Descriptor instead.
func (*Checkin_Status) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Checkin_Status) GetUptime() int64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *Checkin_Status) GetQueueDepth() int64 {
	if x != nil {
		return x.QueueDepth
	}
	return 0
}

func (x *Checkin_Status) GetLastHeartbeat() *timestamppb.Timestamp {
	if x != nil {
		return x.LastHeartbeat
	}
	return nil
}

type SyncMessagesResp_TaskMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// how many time this task can run
	Remain    int64        `protobuf:"varint,2,opt,name=remain,proto3" json:"remain,omitempty"`
	ExpiredAt int64        `protobuf:"varint,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	Trigger   *TaskTrigger `protobuf:"bytes,4,opt,name=trigger,proto3" json:"trigger,omitempty"`
}

func (x *SyncMessagesResp_TaskMetadata) Reset() {
	*x = SyncMessagesResp_TaskMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMessagesResp_TaskMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMessagesResp_TaskMetadata) ProtoMessage() {}

func (x *SyncMessagesResp_TaskMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMessagesResp_TaskMetadata.ProtoReflect.Descriptor instead.
func (*SyncMessagesResp_TaskMetadata) Descriptor() ([]byte, []int) {
	return file_protobuf_node_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SyncMessagesResp_TaskMetadata) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *SyncMessagesResp_TaskMetadata) GetRemain() int64 {
	if x != nil {
		return x.Remain
	}
	return 0
}

func (x *SyncMessagesResp_TaskMetadata) GetExpiredAt() int64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *SyncMessagesResp_TaskMetadata) GetTrigger() *TaskTrigger {
	if x != nil {
		return x.Trigger
	}
	return nil
}

var File_protobuf_node_proto protoreflect.FileDescriptor

var file_protobuf_node_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f,
	0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x76, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x02, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x1a, 0x83, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x44, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x41, 0x0a, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c,
	0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x22, 0x48, 0x0a, 0x0b,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x6f, 0x6e, 0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x63, 0x5f,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x6f, 0x6e,
	0x6f, 0x74, 0x6f, 0x6e, 0x69, 0x63, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xad, 0x02, 0x0a, 0x10,
	0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x25, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x4e, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x91, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x52, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x41,
	0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xac, 0x01, 0x0a,
	0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x10, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x12, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x61, 0x0a, 0x09,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x6e, 0x73,
	0x65, 0x74, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54,
	0x61, 0x73, 0x6b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x04, 0x32,
	0xa0, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x13, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x1a, 0x17, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x1b, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x3e, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x12, 0x1d, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x1e, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x61, 0x76, 0x73, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_node_proto_rawDescOnce sync.Once
	file_protobuf_node_proto_rawDescData = file_protobuf_node_proto_rawDesc
)

func file_protobuf_node_proto_rawDescGZIP() []byte {
	file_protobuf_node_proto_rawDescOnce.Do(func() {
		file_protobuf_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_node_proto_rawDescData)
	})
	return file_protobuf_node_proto_rawDescData
}

var file_protobuf_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_node_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protobuf_node_proto_goTypes = []interface{}{
	(MessageOp)(0),                        // 0: aggregator.MessageOp
	(*Checkin)(nil),                       // 1: aggregator.Checkin
	(*CheckinResp)(nil),                   // 2: aggregator.CheckinResp
	(*SyncMessagesReq)(nil),               // 3: aggregator.SyncMessagesReq
	(*SyncMessagesResp)(nil),              // 4: aggregator.SyncMessagesResp
	(*AckMessageReq)(nil),                 // 5: aggregator.AckMessageReq
	(*NotifyTriggersReq)(nil),             // 6: aggregator.NotifyTriggersReq
	(*NotifyTriggersResp)(nil),            // 7: aggregator.NotifyTriggersResp
	(*Checkin_Status)(nil),                // 8: aggregator.Checkin.Status
	(*SyncMessagesResp_TaskMetadata)(nil), // 9: aggregator.SyncMessagesResp.TaskMetadata
	(*timestamppb.Timestamp)(nil),         // 10: google.protobuf.Timestamp
	(*TriggerMetadata)(nil),               // 11: aggregator.TriggerMetadata
	(*TaskTrigger)(nil),                   // 12: aggregator.TaskTrigger
	(*wrapperspb.BoolValue)(nil),          // 13: google.protobuf.BoolValue
}
var file_protobuf_node_proto_depIdxs = []int32{
	8,  // 0: aggregator.Checkin.status:type_name -> aggregator.Checkin.Status
	10, // 1: aggregator.CheckinResp.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: aggregator.SyncMessagesResp.op:type_name -> aggregator.MessageOp
	9,  // 3: aggregator.SyncMessagesResp.task_metadata:type_name -> aggregator.SyncMessagesResp.TaskMetadata
	11, // 4: aggregator.NotifyTriggersReq.trigger_metadata:type_name -> aggregator.TriggerMetadata
	10, // 5: aggregator.NotifyTriggersResp.updated_at:type_name -> google.protobuf.Timestamp
	10, // 6: aggregator.Checkin.Status.last_heartbeat:type_name -> google.protobuf.Timestamp
	12, // 7: aggregator.SyncMessagesResp.TaskMetadata.trigger:type_name -> aggregator.TaskTrigger
	1,  // 8: aggregator.Node.Ping:input_type -> aggregator.Checkin
	3,  // 9: aggregator.Node.SyncMessages:input_type -> aggregator.SyncMessagesReq
	5,  // 10: aggregator.Node.Ack:input_type -> aggregator.AckMessageReq
	6,  // 11: aggregator.Node.NotifyTriggers:input_type -> aggregator.NotifyTriggersReq
	2,  // 12: aggregator.Node.Ping:output_type -> aggregator.CheckinResp
	4,  // 13: aggregator.Node.SyncMessages:output_type -> aggregator.SyncMessagesResp
	13, // 14: aggregator.Node.Ack:output_type -> google.protobuf.BoolValue
	7,  // 15: aggregator.Node.NotifyTriggers:output_type -> aggregator.NotifyTriggersResp
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_protobuf_node_proto_init() }
func file_protobuf_node_proto_init() {
	if File_protobuf_node_proto != nil {
		return
	}
	file_protobuf_avs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protobuf_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkin); i {
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
		file_protobuf_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckinResp); i {
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
		file_protobuf_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMessagesReq); i {
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
		file_protobuf_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMessagesResp); i {
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
		file_protobuf_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckMessageReq); i {
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
		file_protobuf_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyTriggersReq); i {
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
		file_protobuf_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyTriggersResp); i {
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
		file_protobuf_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkin_Status); i {
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
		file_protobuf_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMessagesResp_TaskMetadata); i {
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
			RawDescriptor: file_protobuf_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_node_proto_goTypes,
		DependencyIndexes: file_protobuf_node_proto_depIdxs,
		EnumInfos:         file_protobuf_node_proto_enumTypes,
		MessageInfos:      file_protobuf_node_proto_msgTypes,
	}.Build()
	File_protobuf_node_proto = out.File
	file_protobuf_node_proto_rawDesc = nil
	file_protobuf_node_proto_goTypes = nil
	file_protobuf_node_proto_depIdxs = nil
}
