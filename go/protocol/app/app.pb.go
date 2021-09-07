// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: app.proto

package app

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

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Args          []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Env           []string `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	Stdin         []byte   `protobuf:"bytes,4,opt,name=stdin,proto3" json:"stdin,omitempty"`
	TimeoutMillis int64    `protobuf:"varint,5,opt,name=timeoutMillis,proto3" json:"timeoutMillis,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Command) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Command) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Command) GetStdin() []byte {
	if x != nil {
		return x.Stdin
	}
	return nil
}

func (x *Command) GetTimeoutMillis() int64 {
	if x != nil {
		return x.TimeoutMillis
	}
	return 0
}

type CommandResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId uint64 `protobuf:"varint,1,opt,name=commandId,proto3" json:"commandId,omitempty"`
	ExitCode  int32  `protobuf:"varint,2,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	Stdout    []byte `protobuf:"bytes,3,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr    []byte `protobuf:"bytes,4,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (x *CommandResult) Reset() {
	*x = CommandResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResult) ProtoMessage() {}

func (x *CommandResult) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResult.ProtoReflect.Descriptor instead.
func (*CommandResult) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{1}
}

func (x *CommandResult) GetCommandId() uint64 {
	if x != nil {
		return x.CommandId
	}
	return 0
}

func (x *CommandResult) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *CommandResult) GetStdout() []byte {
	if x != nil {
		return x.Stdout
	}
	return nil
}

func (x *CommandResult) GetStderr() []byte {
	if x != nil {
		return x.Stderr
	}
	return nil
}

type CommandListMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []*Command `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *CommandListMsg) Reset() {
	*x = CommandListMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandListMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandListMsg) ProtoMessage() {}

func (x *CommandListMsg) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandListMsg.ProtoReflect.Descriptor instead.
func (*CommandListMsg) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{2}
}

func (x *CommandListMsg) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

type CommandResultListMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandResults []*CommandResult `protobuf:"bytes,1,rep,name=commandResults,proto3" json:"commandResults,omitempty"`
}

func (x *CommandResultListMsg) Reset() {
	*x = CommandResultListMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResultListMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResultListMsg) ProtoMessage() {}

func (x *CommandResultListMsg) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResultListMsg.ProtoReflect.Descriptor instead.
func (*CommandResultListMsg) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{3}
}

func (x *CommandResultListMsg) GetCommandResults() []*CommandResult {
	if x != nil {
		return x.CommandResults
	}
	return nil
}

type GetCommandListMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCommandListMsg) Reset() {
	*x = GetCommandListMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommandListMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommandListMsg) ProtoMessage() {}

func (x *GetCommandListMsg) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommandListMsg.ProtoReflect.Descriptor instead.
func (*GetCommandListMsg) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{4}
}

type AgentMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MsgType:
	//	*AgentMsg_CommandResultListMsg
	//	*AgentMsg_GetCommandListMsg
	MsgType isAgentMsg_MsgType `protobuf_oneof:"MsgType"`
}

func (x *AgentMsg) Reset() {
	*x = AgentMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentMsg) ProtoMessage() {}

func (x *AgentMsg) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentMsg.ProtoReflect.Descriptor instead.
func (*AgentMsg) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{5}
}

func (m *AgentMsg) GetMsgType() isAgentMsg_MsgType {
	if m != nil {
		return m.MsgType
	}
	return nil
}

func (x *AgentMsg) GetCommandResultListMsg() *CommandResultListMsg {
	if x, ok := x.GetMsgType().(*AgentMsg_CommandResultListMsg); ok {
		return x.CommandResultListMsg
	}
	return nil
}

func (x *AgentMsg) GetGetCommandListMsg() *GetCommandListMsg {
	if x, ok := x.GetMsgType().(*AgentMsg_GetCommandListMsg); ok {
		return x.GetCommandListMsg
	}
	return nil
}

type isAgentMsg_MsgType interface {
	isAgentMsg_MsgType()
}

type AgentMsg_CommandResultListMsg struct {
	CommandResultListMsg *CommandResultListMsg `protobuf:"bytes,1,opt,name=CommandResultListMsg,proto3,oneof"`
}

type AgentMsg_GetCommandListMsg struct {
	GetCommandListMsg *GetCommandListMsg `protobuf:"bytes,2,opt,name=GetCommandListMsg,proto3,oneof"`
}

func (*AgentMsg_CommandResultListMsg) isAgentMsg_MsgType() {}

func (*AgentMsg_GetCommandListMsg) isAgentMsg_MsgType() {}

type SuccessMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SuccessMsg) Reset() {
	*x = SuccessMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessMsg) ProtoMessage() {}

func (x *SuccessMsg) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessMsg.ProtoReflect.Descriptor instead.
func (*SuccessMsg) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{6}
}

type LPMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MsgType:
	//	*LPMsg_SuccessMsg
	//	*LPMsg_CommandListMsg
	MsgType isLPMsg_MsgType `protobuf_oneof:"MsgType"`
}

func (x *LPMsg) Reset() {
	*x = LPMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LPMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LPMsg) ProtoMessage() {}

func (x *LPMsg) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LPMsg.ProtoReflect.Descriptor instead.
func (*LPMsg) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{7}
}

func (m *LPMsg) GetMsgType() isLPMsg_MsgType {
	if m != nil {
		return m.MsgType
	}
	return nil
}

func (x *LPMsg) GetSuccessMsg() *SuccessMsg {
	if x, ok := x.GetMsgType().(*LPMsg_SuccessMsg); ok {
		return x.SuccessMsg
	}
	return nil
}

func (x *LPMsg) GetCommandListMsg() *CommandListMsg {
	if x, ok := x.GetMsgType().(*LPMsg_CommandListMsg); ok {
		return x.CommandListMsg
	}
	return nil
}

type isLPMsg_MsgType interface {
	isLPMsg_MsgType()
}

type LPMsg_SuccessMsg struct {
	SuccessMsg *SuccessMsg `protobuf:"bytes,1,opt,name=SuccessMsg,proto3,oneof"`
}

type LPMsg_CommandListMsg struct {
	CommandListMsg *CommandListMsg `protobuf:"bytes,2,opt,name=CommandListMsg,proto3,oneof"`
}

func (*LPMsg_SuccessMsg) isLPMsg_MsgType() {}

func (*LPMsg_CommandListMsg) isLPMsg_MsgType() {}

var File_app_proto protoreflect.FileDescriptor

var file_app_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e,
	0x76, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x64, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x64,
	0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c,
	0x6c, 0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0x79, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x22, 0x36, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x4e, 0x0a, 0x14, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x12, 0x36, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67,
	0x22, 0xa6, 0x01, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x4b, 0x0a,
	0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x73, 0x67, 0x48, 0x00, 0x52, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x42, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x42, 0x09,
	0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x7c, 0x0a, 0x05, 0x4c, 0x50, 0x4d, 0x73, 0x67,
	0x12, 0x2d, 0x0a, 0x0a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x73,
	0x67, 0x48, 0x00, 0x52, 0x0a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x73, 0x67, 0x12,
	0x39, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x0e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x4d, 0x73,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_proto_rawDescOnce sync.Once
	file_app_proto_rawDescData = file_app_proto_rawDesc
)

func file_app_proto_rawDescGZIP() []byte {
	file_app_proto_rawDescOnce.Do(func() {
		file_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proto_rawDescData)
	})
	return file_app_proto_rawDescData
}

var file_app_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_proto_goTypes = []interface{}{
	(*Command)(nil),              // 0: Command
	(*CommandResult)(nil),        // 1: CommandResult
	(*CommandListMsg)(nil),       // 2: CommandListMsg
	(*CommandResultListMsg)(nil), // 3: CommandResultListMsg
	(*GetCommandListMsg)(nil),    // 4: GetCommandListMsg
	(*AgentMsg)(nil),             // 5: AgentMsg
	(*SuccessMsg)(nil),           // 6: SuccessMsg
	(*LPMsg)(nil),                // 7: LPMsg
}
var file_app_proto_depIdxs = []int32{
	0, // 0: CommandListMsg.commands:type_name -> Command
	1, // 1: CommandResultListMsg.commandResults:type_name -> CommandResult
	3, // 2: AgentMsg.CommandResultListMsg:type_name -> CommandResultListMsg
	4, // 3: AgentMsg.GetCommandListMsg:type_name -> GetCommandListMsg
	6, // 4: LPMsg.SuccessMsg:type_name -> SuccessMsg
	2, // 5: LPMsg.CommandListMsg:type_name -> CommandListMsg
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_app_proto_init() }
func file_app_proto_init() {
	if File_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResult); i {
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
		file_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandListMsg); i {
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
		file_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResultListMsg); i {
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
		file_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommandListMsg); i {
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
		file_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentMsg); i {
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
		file_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessMsg); i {
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
		file_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LPMsg); i {
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
	file_app_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*AgentMsg_CommandResultListMsg)(nil),
		(*AgentMsg_GetCommandListMsg)(nil),
	}
	file_app_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*LPMsg_SuccessMsg)(nil),
		(*LPMsg_CommandListMsg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_proto_goTypes,
		DependencyIndexes: file_app_proto_depIdxs,
		MessageInfos:      file_app_proto_msgTypes,
	}.Build()
	File_app_proto = out.File
	file_app_proto_rawDesc = nil
	file_app_proto_goTypes = nil
	file_app_proto_depIdxs = nil
}