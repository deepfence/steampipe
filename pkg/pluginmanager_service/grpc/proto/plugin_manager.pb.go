// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.3
// source: plugin_manager.proto

package proto

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connections []string `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetConnections() []string {
	if x != nil {
		return x.Connections
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReattachMap map[string]*ReattachConfig `protobuf:"bytes,1,rep,name=reattach_map,json=reattachMap,proto3" json:"reattach_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FailureMap  map[string]string          `protobuf:"bytes,2,rep,name=failure_map,json=failureMap,proto3" json:"failure_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetReattachMap() map[string]*ReattachConfig {
	if x != nil {
		return x.ReattachMap
	}
	return nil
}

func (x *GetResponse) GetFailureMap() map[string]string {
	if x != nil {
		return x.FailureMap
	}
	return nil
}

type RefreshConnectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshConnectionsRequest) Reset() {
	*x = RefreshConnectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshConnectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshConnectionsRequest) ProtoMessage() {}

func (x *RefreshConnectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshConnectionsRequest.ProtoReflect.Descriptor instead.
func (*RefreshConnectionsRequest) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{2}
}

type RefreshConnectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshConnectionsResponse) Reset() {
	*x = RefreshConnectionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshConnectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshConnectionsResponse) ProtoMessage() {}

func (x *RefreshConnectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshConnectionsResponse.ProtoReflect.Descriptor instead.
func (*RefreshConnectionsResponse) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{3}
}

type ShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShutdownRequest) Reset() {
	*x = ShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownRequest) ProtoMessage() {}

func (x *ShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownRequest.ProtoReflect.Descriptor instead.
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{4}
}

type ShutdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShutdownResponse) Reset() {
	*x = ShutdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownResponse) ProtoMessage() {}

func (x *ShutdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownResponse.ProtoReflect.Descriptor instead.
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{5}
}

type ReattachConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol            string               `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ProtocolVersion     int64                `protobuf:"varint,2,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	Addr                *NetAddr             `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Pid                 int64                `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
	SupportedOperations *SupportedOperations `protobuf:"bytes,5,opt,name=supported_operations,json=supportedOperations,proto3" json:"supported_operations,omitempty"`
	Connections         []string             `protobuf:"bytes,6,rep,name=connections,proto3" json:"connections,omitempty"`
	Plugin              string               `protobuf:"bytes,7,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *ReattachConfig) Reset() {
	*x = ReattachConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReattachConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReattachConfig) ProtoMessage() {}

func (x *ReattachConfig) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReattachConfig.ProtoReflect.Descriptor instead.
func (*ReattachConfig) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{6}
}

func (x *ReattachConfig) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ReattachConfig) GetProtocolVersion() int64 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *ReattachConfig) GetAddr() *NetAddr {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *ReattachConfig) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ReattachConfig) GetSupportedOperations() *SupportedOperations {
	if x != nil {
		return x.SupportedOperations
	}
	return nil
}

func (x *ReattachConfig) GetConnections() []string {
	if x != nil {
		return x.Connections
	}
	return nil
}

func (x *ReattachConfig) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

// NOTE: this must be consistent with GetSupportedOperationsResponse in steampipe-plugin-sdk/grpc/proto/plugin.proto
type SupportedOperations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryCache          bool `protobuf:"varint,1,opt,name=query_cache,json=queryCache,proto3" json:"query_cache,omitempty"`
	MultipleConnections bool `protobuf:"varint,2,opt,name=multiple_connections,json=multipleConnections,proto3" json:"multiple_connections,omitempty"`
	MessageStream       bool `protobuf:"varint,3,opt,name=message_stream,json=messageStream,proto3" json:"message_stream,omitempty"`
}

func (x *SupportedOperations) Reset() {
	*x = SupportedOperations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupportedOperations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportedOperations) ProtoMessage() {}

func (x *SupportedOperations) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportedOperations.ProtoReflect.Descriptor instead.
func (*SupportedOperations) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{7}
}

func (x *SupportedOperations) GetQueryCache() bool {
	if x != nil {
		return x.QueryCache
	}
	return false
}

func (x *SupportedOperations) GetMultipleConnections() bool {
	if x != nil {
		return x.MultipleConnections
	}
	return false
}

func (x *SupportedOperations) GetMessageStream() bool {
	if x != nil {
		return x.MessageStream
	}
	return false
}

type NetAddr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=Network,proto3" json:"Network,omitempty"` // name of the network (for example, "tcp", "udp")
	Address string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"` // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

func (x *NetAddr) Reset() {
	*x = NetAddr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetAddr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetAddr) ProtoMessage() {}

func (x *NetAddr) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetAddr.ProtoReflect.Descriptor instead.
func (*NetAddr) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{8}
}

func (x *NetAddr) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *NetAddr) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_plugin_manager_proto protoreflect.FileDescriptor

var file_plugin_manager_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb0, 0x02,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0c, 0x72, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x4d, 0x61, 0x70, 0x12, 0x43, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x55, 0x0a, 0x10, 0x52, 0x65,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3d, 0x0a, 0x0f, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x1b, 0x0a, 0x19, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a,
	0x1a, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x53,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x96, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x4d, 0x0a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x13,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x3d,
	0x0a, 0x07, 0x4e, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xdb, 0x01,
	0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5b, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08,
	0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_manager_proto_rawDescOnce sync.Once
	file_plugin_manager_proto_rawDescData = file_plugin_manager_proto_rawDesc
)

func file_plugin_manager_proto_rawDescGZIP() []byte {
	file_plugin_manager_proto_rawDescOnce.Do(func() {
		file_plugin_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_manager_proto_rawDescData)
	})
	return file_plugin_manager_proto_rawDescData
}

var file_plugin_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_plugin_manager_proto_goTypes = []interface{}{
	(*GetRequest)(nil),                 // 0: proto.GetRequest
	(*GetResponse)(nil),                // 1: proto.GetResponse
	(*RefreshConnectionsRequest)(nil),  // 2: proto.RefreshConnectionsRequest
	(*RefreshConnectionsResponse)(nil), // 3: proto.RefreshConnectionsResponse
	(*ShutdownRequest)(nil),            // 4: proto.ShutdownRequest
	(*ShutdownResponse)(nil),           // 5: proto.ShutdownResponse
	(*ReattachConfig)(nil),             // 6: proto.ReattachConfig
	(*SupportedOperations)(nil),        // 7: proto.SupportedOperations
	(*NetAddr)(nil),                    // 8: proto.NetAddr
	nil,                                // 9: proto.GetResponse.ReattachMapEntry
	nil,                                // 10: proto.GetResponse.FailureMapEntry
}
var file_plugin_manager_proto_depIdxs = []int32{
	9,  // 0: proto.GetResponse.reattach_map:type_name -> proto.GetResponse.ReattachMapEntry
	10, // 1: proto.GetResponse.failure_map:type_name -> proto.GetResponse.FailureMapEntry
	8,  // 2: proto.ReattachConfig.addr:type_name -> proto.NetAddr
	7,  // 3: proto.ReattachConfig.supported_operations:type_name -> proto.SupportedOperations
	6,  // 4: proto.GetResponse.ReattachMapEntry.value:type_name -> proto.ReattachConfig
	0,  // 5: proto.PluginManager.Get:input_type -> proto.GetRequest
	2,  // 6: proto.PluginManager.RefreshConnections:input_type -> proto.RefreshConnectionsRequest
	4,  // 7: proto.PluginManager.Shutdown:input_type -> proto.ShutdownRequest
	1,  // 8: proto.PluginManager.Get:output_type -> proto.GetResponse
	3,  // 9: proto.PluginManager.RefreshConnections:output_type -> proto.RefreshConnectionsResponse
	5,  // 10: proto.PluginManager.Shutdown:output_type -> proto.ShutdownResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_plugin_manager_proto_init() }
func file_plugin_manager_proto_init() {
	if File_plugin_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_plugin_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_plugin_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshConnectionsRequest); i {
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
		file_plugin_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshConnectionsResponse); i {
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
		file_plugin_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownRequest); i {
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
		file_plugin_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownResponse); i {
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
		file_plugin_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReattachConfig); i {
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
		file_plugin_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupportedOperations); i {
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
		file_plugin_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetAddr); i {
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
			RawDescriptor: file_plugin_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_manager_proto_goTypes,
		DependencyIndexes: file_plugin_manager_proto_depIdxs,
		MessageInfos:      file_plugin_manager_proto_msgTypes,
	}.Build()
	File_plugin_manager_proto = out.File
	file_plugin_manager_proto_rawDesc = nil
	file_plugin_manager_proto_goTypes = nil
	file_plugin_manager_proto_depIdxs = nil
}
