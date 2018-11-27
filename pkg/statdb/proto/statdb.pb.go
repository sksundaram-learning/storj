// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statdb.proto

package statdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Node is info for a updating a single storagenode, used in the Update rpc calls
type Node struct {
	NodeId               []byte   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	LatencyList          []int64  `protobuf:"varint,2,rep,packed,name=latency_list,json=latencyList,proto3" json:"latency_list,omitempty"`
	AuditSuccess         bool     `protobuf:"varint,3,opt,name=audit_success,json=auditSuccess,proto3" json:"audit_success,omitempty"`
	IsUp                 bool     `protobuf:"varint,4,opt,name=is_up,json=isUp,proto3" json:"is_up,omitempty"`
	UpdateLatency        bool     `protobuf:"varint,5,opt,name=update_latency,json=updateLatency,proto3" json:"update_latency,omitempty"`
	UpdateAuditSuccess   bool     `protobuf:"varint,6,opt,name=update_audit_success,json=updateAuditSuccess,proto3" json:"update_audit_success,omitempty"`
	UpdateUptime         bool     `protobuf:"varint,7,opt,name=update_uptime,json=updateUptime,proto3" json:"update_uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Node) GetLatencyList() []int64 {
	if m != nil {
		return m.LatencyList
	}
	return nil
}

func (m *Node) GetAuditSuccess() bool {
	if m != nil {
		return m.AuditSuccess
	}
	return false
}

func (m *Node) GetIsUp() bool {
	if m != nil {
		return m.IsUp
	}
	return false
}

func (m *Node) GetUpdateLatency() bool {
	if m != nil {
		return m.UpdateLatency
	}
	return false
}

func (m *Node) GetUpdateAuditSuccess() bool {
	if m != nil {
		return m.UpdateAuditSuccess
	}
	return false
}

func (m *Node) GetUpdateUptime() bool {
	if m != nil {
		return m.UpdateUptime
	}
	return false
}

// NodeStats is info about a single storagenode stored in the stats db
type NodeStats struct {
	NodeId               []byte   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Latency_90           int64    `protobuf:"varint,2,opt,name=latency_90,json=latency90,proto3" json:"latency_90,omitempty"`
	AuditSuccessRatio    float64  `protobuf:"fixed64,3,opt,name=audit_success_ratio,json=auditSuccessRatio,proto3" json:"audit_success_ratio,omitempty"`
	UptimeRatio          float64  `protobuf:"fixed64,4,opt,name=uptime_ratio,json=uptimeRatio,proto3" json:"uptime_ratio,omitempty"`
	AuditCount           int64    `protobuf:"varint,5,opt,name=audit_count,json=auditCount,proto3" json:"audit_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStats) Reset()         { *m = NodeStats{} }
func (m *NodeStats) String() string { return proto.CompactTextString(m) }
func (*NodeStats) ProtoMessage()    {}
func (*NodeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{1}
}
func (m *NodeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStats.Unmarshal(m, b)
}
func (m *NodeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStats.Marshal(b, m, deterministic)
}
func (dst *NodeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStats.Merge(dst, src)
}
func (m *NodeStats) XXX_Size() int {
	return xxx_messageInfo_NodeStats.Size(m)
}
func (m *NodeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStats proto.InternalMessageInfo

func (m *NodeStats) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *NodeStats) GetLatency_90() int64 {
	if m != nil {
		return m.Latency_90
	}
	return 0
}

func (m *NodeStats) GetAuditSuccessRatio() float64 {
	if m != nil {
		return m.AuditSuccessRatio
	}
	return 0
}

func (m *NodeStats) GetUptimeRatio() float64 {
	if m != nil {
		return m.UptimeRatio
	}
	return 0
}

func (m *NodeStats) GetAuditCount() int64 {
	if m != nil {
		return m.AuditCount
	}
	return 0
}

// CreateRequest is a request message for the Create rpc call
type CreateRequest struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{2}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *CreateRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// CreateResponse is a response message for the Create rpc call
type CreateResponse struct {
	Stats                *NodeStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{3}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStats() *NodeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// GetRequest is a request message for the Get rpc call
type GetRequest struct {
	NodeId               []byte   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{4}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *GetRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// GetResponse is a response message for the Get rpc call
type GetResponse struct {
	Stats                *NodeStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{5}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetStats() *NodeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// FindValidNodesRequest is a request message for the FindValidNodes rpc call
type FindValidNodesRequest struct {
	NodeIds              [][]byte   `protobuf:"bytes,1,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
	MinStats             *NodeStats `protobuf:"bytes,2,opt,name=min_stats,json=minStats,proto3" json:"min_stats,omitempty"`
	APIKey               []byte     `protobuf:"bytes,3,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindValidNodesRequest) Reset()         { *m = FindValidNodesRequest{} }
func (m *FindValidNodesRequest) String() string { return proto.CompactTextString(m) }
func (*FindValidNodesRequest) ProtoMessage()    {}
func (*FindValidNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{6}
}
func (m *FindValidNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindValidNodesRequest.Unmarshal(m, b)
}
func (m *FindValidNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindValidNodesRequest.Marshal(b, m, deterministic)
}
func (dst *FindValidNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindValidNodesRequest.Merge(dst, src)
}
func (m *FindValidNodesRequest) XXX_Size() int {
	return xxx_messageInfo_FindValidNodesRequest.Size(m)
}
func (m *FindValidNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindValidNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindValidNodesRequest proto.InternalMessageInfo

func (m *FindValidNodesRequest) GetNodeIds() [][]byte {
	if m != nil {
		return m.NodeIds
	}
	return nil
}

func (m *FindValidNodesRequest) GetMinStats() *NodeStats {
	if m != nil {
		return m.MinStats
	}
	return nil
}

func (m *FindValidNodesRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// FindValidNodesResponse is a response message for the FindValidNodes rpc call
type FindValidNodesResponse struct {
	PassedIds            [][]byte `protobuf:"bytes,1,rep,name=passed_ids,json=passedIds,proto3" json:"passed_ids,omitempty"`
	FailedIds            [][]byte `protobuf:"bytes,2,rep,name=failed_ids,json=failedIds,proto3" json:"failed_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValidNodesResponse) Reset()         { *m = FindValidNodesResponse{} }
func (m *FindValidNodesResponse) String() string { return proto.CompactTextString(m) }
func (*FindValidNodesResponse) ProtoMessage()    {}
func (*FindValidNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{7}
}
func (m *FindValidNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindValidNodesResponse.Unmarshal(m, b)
}
func (m *FindValidNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindValidNodesResponse.Marshal(b, m, deterministic)
}
func (dst *FindValidNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindValidNodesResponse.Merge(dst, src)
}
func (m *FindValidNodesResponse) XXX_Size() int {
	return xxx_messageInfo_FindValidNodesResponse.Size(m)
}
func (m *FindValidNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindValidNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindValidNodesResponse proto.InternalMessageInfo

func (m *FindValidNodesResponse) GetPassedIds() [][]byte {
	if m != nil {
		return m.PassedIds
	}
	return nil
}

func (m *FindValidNodesResponse) GetFailedIds() [][]byte {
	if m != nil {
		return m.FailedIds
	}
	return nil
}

// UpdateRequest is a request message for the Update rpc call
type UpdateRequest struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{8}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *UpdateRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// UpdateRequest is a response message for the Update rpc call
type UpdateResponse struct {
	Stats                *NodeStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{9}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetStats() *NodeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// UpdateBatchRequest is a request message for the UpdateBatch rpc call
type UpdateBatchRequest struct {
	NodeList             []*Node  `protobuf:"bytes,1,rep,name=node_list,json=nodeList,proto3" json:"node_list,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBatchRequest) Reset()         { *m = UpdateBatchRequest{} }
func (m *UpdateBatchRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBatchRequest) ProtoMessage()    {}
func (*UpdateBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{10}
}
func (m *UpdateBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBatchRequest.Unmarshal(m, b)
}
func (m *UpdateBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBatchRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBatchRequest.Merge(dst, src)
}
func (m *UpdateBatchRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBatchRequest.Size(m)
}
func (m *UpdateBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBatchRequest proto.InternalMessageInfo

func (m *UpdateBatchRequest) GetNodeList() []*Node {
	if m != nil {
		return m.NodeList
	}
	return nil
}

func (m *UpdateBatchRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// UpdateBatchResponse is a response message for the UpdateBatch rpc call
type UpdateBatchResponse struct {
	StatsList            []*NodeStats `protobuf:"bytes,1,rep,name=stats_list,json=statsList,proto3" json:"stats_list,omitempty"`
	FailedNodes          []*Node      `protobuf:"bytes,2,rep,name=failed_nodes,json=failedNodes,proto3" json:"failed_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateBatchResponse) Reset()         { *m = UpdateBatchResponse{} }
func (m *UpdateBatchResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBatchResponse) ProtoMessage()    {}
func (*UpdateBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{11}
}
func (m *UpdateBatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBatchResponse.Unmarshal(m, b)
}
func (m *UpdateBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBatchResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBatchResponse.Merge(dst, src)
}
func (m *UpdateBatchResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBatchResponse.Size(m)
}
func (m *UpdateBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBatchResponse proto.InternalMessageInfo

func (m *UpdateBatchResponse) GetStatsList() []*NodeStats {
	if m != nil {
		return m.StatsList
	}
	return nil
}

func (m *UpdateBatchResponse) GetFailedNodes() []*Node {
	if m != nil {
		return m.FailedNodes
	}
	return nil
}

// CreateEntryIfNotExistsRequest is a request message for the CreateEntryIfNotExists rpc call
type CreateEntryIfNotExistsRequest struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntryIfNotExistsRequest) Reset()         { *m = CreateEntryIfNotExistsRequest{} }
func (m *CreateEntryIfNotExistsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEntryIfNotExistsRequest) ProtoMessage()    {}
func (*CreateEntryIfNotExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{12}
}
func (m *CreateEntryIfNotExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntryIfNotExistsRequest.Unmarshal(m, b)
}
func (m *CreateEntryIfNotExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntryIfNotExistsRequest.Marshal(b, m, deterministic)
}
func (dst *CreateEntryIfNotExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntryIfNotExistsRequest.Merge(dst, src)
}
func (m *CreateEntryIfNotExistsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEntryIfNotExistsRequest.Size(m)
}
func (m *CreateEntryIfNotExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntryIfNotExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntryIfNotExistsRequest proto.InternalMessageInfo

func (m *CreateEntryIfNotExistsRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *CreateEntryIfNotExistsRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// CreateEntryIfNotExistsResponse is a response message for the CreateEntryIfNotExists rpc call
type CreateEntryIfNotExistsResponse struct {
	Stats                *NodeStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateEntryIfNotExistsResponse) Reset()         { *m = CreateEntryIfNotExistsResponse{} }
func (m *CreateEntryIfNotExistsResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEntryIfNotExistsResponse) ProtoMessage()    {}
func (*CreateEntryIfNotExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_5a3c71d6f7d0333f, []int{13}
}
func (m *CreateEntryIfNotExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntryIfNotExistsResponse.Unmarshal(m, b)
}
func (m *CreateEntryIfNotExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntryIfNotExistsResponse.Marshal(b, m, deterministic)
}
func (dst *CreateEntryIfNotExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntryIfNotExistsResponse.Merge(dst, src)
}
func (m *CreateEntryIfNotExistsResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEntryIfNotExistsResponse.Size(m)
}
func (m *CreateEntryIfNotExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntryIfNotExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntryIfNotExistsResponse proto.InternalMessageInfo

func (m *CreateEntryIfNotExistsResponse) GetStats() *NodeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "statdb.Node")
	proto.RegisterType((*NodeStats)(nil), "statdb.NodeStats")
	proto.RegisterType((*CreateRequest)(nil), "statdb.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "statdb.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "statdb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "statdb.GetResponse")
	proto.RegisterType((*FindValidNodesRequest)(nil), "statdb.FindValidNodesRequest")
	proto.RegisterType((*FindValidNodesResponse)(nil), "statdb.FindValidNodesResponse")
	proto.RegisterType((*UpdateRequest)(nil), "statdb.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "statdb.UpdateResponse")
	proto.RegisterType((*UpdateBatchRequest)(nil), "statdb.UpdateBatchRequest")
	proto.RegisterType((*UpdateBatchResponse)(nil), "statdb.UpdateBatchResponse")
	proto.RegisterType((*CreateEntryIfNotExistsRequest)(nil), "statdb.CreateEntryIfNotExistsRequest")
	proto.RegisterType((*CreateEntryIfNotExistsResponse)(nil), "statdb.CreateEntryIfNotExistsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatDBClient is the client API for StatDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatDBClient interface {
	// Create a db entry for the provided storagenode ID
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get uses a storagenode ID to get that storagenode's stats
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// FindValidNodes gets a subset of storagenodes that fit minimum reputation args
	FindValidNodes(ctx context.Context, in *FindValidNodesRequest, opts ...grpc.CallOption) (*FindValidNodesResponse, error)
	// Update updates storagenode stats for a single storagenode
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// UpdateBatch updates storagenode stats for multiple farmers at a time
	UpdateBatch(ctx context.Context, in *UpdateBatchRequest, opts ...grpc.CallOption) (*UpdateBatchResponse, error)
	// CreateEntryIfNotExists creates a db entry if it didn't exist
	CreateEntryIfNotExists(ctx context.Context, in *CreateEntryIfNotExistsRequest, opts ...grpc.CallOption) (*CreateEntryIfNotExistsResponse, error)
}

type statDBClient struct {
	cc *grpc.ClientConn
}

func NewStatDBClient(cc *grpc.ClientConn) StatDBClient {
	return &statDBClient{cc}
}

func (c *statDBClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statDBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statDBClient) FindValidNodes(ctx context.Context, in *FindValidNodesRequest, opts ...grpc.CallOption) (*FindValidNodesResponse, error) {
	out := new(FindValidNodesResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/FindValidNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statDBClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statDBClient) UpdateBatch(ctx context.Context, in *UpdateBatchRequest, opts ...grpc.CallOption) (*UpdateBatchResponse, error) {
	out := new(UpdateBatchResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/UpdateBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statDBClient) CreateEntryIfNotExists(ctx context.Context, in *CreateEntryIfNotExistsRequest, opts ...grpc.CallOption) (*CreateEntryIfNotExistsResponse, error) {
	out := new(CreateEntryIfNotExistsResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/CreateEntryIfNotExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatDBServer is the server API for StatDB service.
type StatDBServer interface {
	// Create a db entry for the provided storagenode ID
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get uses a storagenode ID to get that storagenode's stats
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// FindValidNodes gets a subset of storagenodes that fit minimum reputation args
	FindValidNodes(context.Context, *FindValidNodesRequest) (*FindValidNodesResponse, error)
	// Update updates storagenode stats for a single storagenode
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// UpdateBatch updates storagenode stats for multiple farmers at a time
	UpdateBatch(context.Context, *UpdateBatchRequest) (*UpdateBatchResponse, error)
	// CreateEntryIfNotExists creates a db entry if it didn't exist
	CreateEntryIfNotExists(context.Context, *CreateEntryIfNotExistsRequest) (*CreateEntryIfNotExistsResponse, error)
}

func RegisterStatDBServer(s *grpc.Server, srv StatDBServer) {
	s.RegisterService(&_StatDB_serviceDesc, srv)
}

func _StatDB_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatDB_FindValidNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindValidNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).FindValidNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/FindValidNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).FindValidNodes(ctx, req.(*FindValidNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatDB_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatDB_UpdateBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).UpdateBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/UpdateBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).UpdateBatch(ctx, req.(*UpdateBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatDB_CreateEntryIfNotExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryIfNotExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).CreateEntryIfNotExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/CreateEntryIfNotExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).CreateEntryIfNotExists(ctx, req.(*CreateEntryIfNotExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "statdb.StatDB",
	HandlerType: (*StatDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StatDB_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StatDB_Get_Handler,
		},
		{
			MethodName: "FindValidNodes",
			Handler:    _StatDB_FindValidNodes_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StatDB_Update_Handler,
		},
		{
			MethodName: "UpdateBatch",
			Handler:    _StatDB_UpdateBatch_Handler,
		},
		{
			MethodName: "CreateEntryIfNotExists",
			Handler:    _StatDB_CreateEntryIfNotExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "statdb.proto",
}

func init() { proto.RegisterFile("statdb.proto", fileDescriptor_statdb_5a3c71d6f7d0333f) }

var fileDescriptor_statdb_5a3c71d6f7d0333f = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0xd4, 0x6d, 0xc6, 0x6e, 0xa5, 0x6e, 0x68, 0x30, 0x41, 0x29, 0x21, 0x55, 0x21,
	0x5c, 0x42, 0x54, 0x24, 0xaa, 0x1e, 0x38, 0xb4, 0xa5, 0xad, 0x22, 0xaa, 0x82, 0xb6, 0x6a, 0x11,
	0x27, 0xcb, 0x8d, 0xb7, 0xb0, 0x52, 0x6a, 0x9b, 0xec, 0x5a, 0x6a, 0xf9, 0x26, 0xee, 0xfc, 0x1b,
	0x27, 0xb4, 0x3b, 0x6b, 0xc5, 0x8e, 0x62, 0xa1, 0x08, 0x8e, 0x7e, 0x33, 0xfb, 0xde, 0xdb, 0x37,
	0xb3, 0x09, 0x78, 0x42, 0x86, 0x32, 0xba, 0x19, 0xa4, 0xd3, 0x44, 0x26, 0xc4, 0xc1, 0xaf, 0xde,
	0x6f, 0x0b, 0xea, 0x17, 0x49, 0xc4, 0xc8, 0x63, 0x58, 0x8d, 0x93, 0x88, 0x05, 0x3c, 0xf2, 0xad,
	0xae, 0xd5, 0xf7, 0xa8, 0xa3, 0x3e, 0x47, 0x11, 0x79, 0x0e, 0xde, 0x24, 0x94, 0x2c, 0x1e, 0x3f,
	0x04, 0x13, 0x2e, 0xa4, 0x5f, 0xeb, 0xda, 0x7d, 0x9b, 0xba, 0x06, 0x3b, 0xe7, 0x42, 0x92, 0x1d,
	0x58, 0x0f, 0xb3, 0x88, 0xcb, 0x40, 0x64, 0xe3, 0x31, 0x13, 0xc2, 0xb7, 0xbb, 0x56, 0x7f, 0x8d,
	0x7a, 0x1a, 0xbc, 0x44, 0x8c, 0x34, 0x61, 0x85, 0x8b, 0x20, 0x4b, 0xfd, 0xba, 0x2e, 0xd6, 0xb9,
	0xb8, 0x4a, 0xc9, 0x2e, 0x6c, 0x64, 0x69, 0x14, 0x4a, 0x16, 0x18, 0x3e, 0x7f, 0x45, 0x57, 0xd7,
	0x11, 0x3d, 0x47, 0x90, 0x0c, 0xe1, 0x91, 0x69, 0x2b, 0xeb, 0x38, 0xba, 0x99, 0x60, 0xed, 0xb0,
	0xa8, 0xb6, 0x03, 0x86, 0x22, 0xc8, 0x52, 0xc9, 0xef, 0x98, 0xbf, 0x8a, 0x96, 0x10, 0xbc, 0xd2,
	0x58, 0xef, 0x97, 0x05, 0x0d, 0x75, 0xf9, 0x4b, 0x19, 0x4a, 0x51, 0x9d, 0x40, 0x07, 0x20, 0x4f,
	0xe0, 0x60, 0xe8, 0xd7, 0xba, 0x56, 0xdf, 0xa6, 0x0d, 0x83, 0x1c, 0x0c, 0xc9, 0x00, 0x9a, 0x25,
	0x57, 0xc1, 0x34, 0x94, 0x3c, 0xd1, 0x19, 0x58, 0x74, 0xb3, 0x98, 0x01, 0x55, 0x05, 0x15, 0x28,
	0x7a, 0x32, 0x8d, 0x75, 0xdd, 0xe8, 0x22, 0x86, 0x2d, 0xcf, 0xc0, 0x45, 0xca, 0x71, 0x92, 0xc5,
	0x52, 0x67, 0x62, 0x53, 0xd0, 0xd0, 0xb1, 0x42, 0x7a, 0x23, 0x58, 0x3f, 0x9e, 0xb2, 0x50, 0x32,
	0xca, 0xbe, 0x67, 0x4c, 0x48, 0xd2, 0x85, 0xba, 0x72, 0xab, 0x9d, 0xbb, 0x7b, 0xde, 0xc0, 0x0c,
	0x5b, 0xdd, 0x8e, 0xea, 0x0a, 0x69, 0x81, 0x73, 0xf8, 0x69, 0xf4, 0x81, 0x3d, 0xe8, 0x1b, 0x78,
	0xd4, 0x7c, 0xf5, 0x0e, 0x60, 0x23, 0xa7, 0x12, 0x69, 0x12, 0x0b, 0x46, 0x5e, 0xc2, 0x8a, 0x3a,
	0x2e, 0x0c, 0xd9, 0x66, 0x91, 0x4c, 0x47, 0x45, 0xb1, 0xde, 0x7b, 0x07, 0x70, 0xc6, 0x64, 0x6e,
	0xa1, 0x32, 0xbf, 0x2a, 0xe5, 0xb7, 0xe0, 0xea, 0xe3, 0xcb, 0xca, 0xfe, 0x80, 0xad, 0x53, 0x1e,
	0x47, 0xd7, 0xe1, 0x84, 0x47, 0xaa, 0x28, 0x72, 0x07, 0x4f, 0x60, 0xcd, 0x38, 0x50, 0x24, 0x76,
	0xdf, 0xa3, 0xab, 0x68, 0x41, 0x90, 0x01, 0x34, 0xee, 0x78, 0x1c, 0xa0, 0x40, 0xad, 0x4a, 0x60,
	0xed, 0x8e, 0xc7, 0xb8, 0x0c, 0x33, 0xcf, 0x76, 0xc9, 0xf3, 0x35, 0xb4, 0xe6, 0xb5, 0x8d, 0xfd,
	0x0e, 0x40, 0x1a, 0x0a, 0xc1, 0xa2, 0x82, 0x7c, 0x03, 0x11, 0x65, 0xa0, 0x03, 0x70, 0x1b, 0xf2,
	0x89, 0x29, 0xd7, 0xb0, 0x8c, 0xc8, 0x28, 0x12, 0x6a, 0xa0, 0x57, 0x7a, 0x35, 0xff, 0xcb, 0x40,
	0x73, 0xaa, 0x65, 0x93, 0xfd, 0x0c, 0x04, 0x8f, 0x1e, 0x85, 0x72, 0xfc, 0x2d, 0xb7, 0xf2, 0x0a,
	0x1a, 0x3a, 0x56, 0xfd, 0xfc, 0xd5, 0xc5, 0xe6, 0xfd, 0xe8, 0xd4, 0xf5, 0x2f, 0x41, 0x95, 0xa7,
	0x7b, 0x68, 0x96, 0x88, 0x8d, 0xb1, 0x21, 0x80, 0x16, 0x2e, 0x52, 0x2f, 0x70, 0xd7, 0xd0, 0x4d,
	0x5a, 0xe0, 0x35, 0x78, 0x26, 0x46, 0xa5, 0x89, 0x41, 0xce, 0xdb, 0x71, 0xb1, 0x43, 0x8f, 0xa7,
	0xf7, 0x05, 0x3a, 0xb8, 0xde, 0x27, 0xb1, 0x9c, 0x3e, 0x8c, 0x6e, 0x2f, 0x12, 0x79, 0x72, 0xcf,
	0x85, 0x14, 0xff, 0x1e, 0xf4, 0x08, 0xb6, 0xab, 0xa8, 0x97, 0x0c, 0x7e, 0xef, 0xa7, 0x0d, 0x8e,
	0x02, 0xde, 0x1f, 0x91, 0x7d, 0x70, 0x90, 0x95, 0x6c, 0xe5, 0xed, 0xa5, 0xa7, 0xde, 0x6e, 0xcd,
	0xc3, 0x46, 0x6c, 0x00, 0xf6, 0x19, 0x93, 0x84, 0xe4, 0xe5, 0xd9, 0xd3, 0x6c, 0x37, 0x4b, 0x98,
	0xe9, 0xff, 0x08, 0x1b, 0xe5, 0x55, 0x26, 0x9d, 0xbc, 0x6d, 0xe1, 0xf3, 0x6a, 0x6f, 0x57, 0x95,
	0x0d, 0xe1, 0x3e, 0x38, 0x38, 0xe4, 0x99, 0xf3, 0xd2, 0x4e, 0xcf, 0x9c, 0xcf, 0xed, 0xe7, 0x29,
	0xb8, 0x85, 0xed, 0x20, 0xed, 0x72, 0x5b, 0x71, 0x17, 0xdb, 0x4f, 0x17, 0xd6, 0x0c, 0xcf, 0x57,
	0x68, 0x2d, 0x1e, 0x08, 0xd9, 0x2d, 0x67, 0x56, 0xb1, 0x0b, 0xed, 0x17, 0x7f, 0x6b, 0x43, 0xa1,
	0x1b, 0x47, 0xff, 0x89, 0xbe, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x63, 0x10, 0xde, 0x1a, 0x54,
	0x07, 0x00, 0x00,
}
