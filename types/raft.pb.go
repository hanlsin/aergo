// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raft.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// cluster member for raft consensus
type MembershipChangeType int32

const (
	MembershipChangeType_ADD_MEMBER    MembershipChangeType = 0
	MembershipChangeType_REMOVE_MEMBER MembershipChangeType = 1
)

var MembershipChangeType_name = map[int32]string{
	0: "ADD_MEMBER",
	1: "REMOVE_MEMBER",
}
var MembershipChangeType_value = map[string]int32{
	"ADD_MEMBER":    0,
	"REMOVE_MEMBER": 1,
}

func (x MembershipChangeType) String() string {
	return proto.EnumName(MembershipChangeType_name, int32(x))
}
func (MembershipChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{0}
}

type ConfChangeState int32

const (
	ConfChangeState_CONF_CHANGE_STATE_PROPOSED ConfChangeState = 0
	ConfChangeState_CONF_CHANGE_STATE_SAVED    ConfChangeState = 1
	ConfChangeState_CONF_CHANGE_STATE_APPLIED  ConfChangeState = 2
)

var ConfChangeState_name = map[int32]string{
	0: "CONF_CHANGE_STATE_PROPOSED",
	1: "CONF_CHANGE_STATE_SAVED",
	2: "CONF_CHANGE_STATE_APPLIED",
}
var ConfChangeState_value = map[string]int32{
	"CONF_CHANGE_STATE_PROPOSED": 0,
	"CONF_CHANGE_STATE_SAVED":    1,
	"CONF_CHANGE_STATE_APPLIED":  2,
}

func (x ConfChangeState) String() string {
	return proto.EnumName(ConfChangeState_name, int32(x))
}
func (ConfChangeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{1}
}

type MemberAttr struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PeerID               []byte   `protobuf:"bytes,4,opt,name=peerID,proto3" json:"peerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberAttr) Reset()         { *m = MemberAttr{} }
func (m *MemberAttr) String() string { return proto.CompactTextString(m) }
func (*MemberAttr) ProtoMessage()    {}
func (*MemberAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{0}
}
func (m *MemberAttr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberAttr.Unmarshal(m, b)
}
func (m *MemberAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberAttr.Marshal(b, m, deterministic)
}
func (dst *MemberAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberAttr.Merge(dst, src)
}
func (m *MemberAttr) XXX_Size() int {
	return xxx_messageInfo_MemberAttr.Size(m)
}
func (m *MemberAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberAttr.DiscardUnknown(m)
}

var xxx_messageInfo_MemberAttr proto.InternalMessageInfo

func (m *MemberAttr) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *MemberAttr) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MemberAttr) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MemberAttr) GetPeerID() []byte {
	if m != nil {
		return m.PeerID
	}
	return nil
}

type MembershipChange struct {
	Type                 MembershipChangeType `protobuf:"varint,1,opt,name=type,proto3,enum=types.MembershipChangeType" json:"type,omitempty"`
	RequestID            uint64               `protobuf:"varint,2,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Attr                 *MemberAttr          `protobuf:"bytes,3,opt,name=attr,proto3" json:"attr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MembershipChange) Reset()         { *m = MembershipChange{} }
func (m *MembershipChange) String() string { return proto.CompactTextString(m) }
func (*MembershipChange) ProtoMessage()    {}
func (*MembershipChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{1}
}
func (m *MembershipChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MembershipChange.Unmarshal(m, b)
}
func (m *MembershipChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MembershipChange.Marshal(b, m, deterministic)
}
func (dst *MembershipChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MembershipChange.Merge(dst, src)
}
func (m *MembershipChange) XXX_Size() int {
	return xxx_messageInfo_MembershipChange.Size(m)
}
func (m *MembershipChange) XXX_DiscardUnknown() {
	xxx_messageInfo_MembershipChange.DiscardUnknown(m)
}

var xxx_messageInfo_MembershipChange proto.InternalMessageInfo

func (m *MembershipChange) GetType() MembershipChangeType {
	if m != nil {
		return m.Type
	}
	return MembershipChangeType_ADD_MEMBER
}

func (m *MembershipChange) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *MembershipChange) GetAttr() *MemberAttr {
	if m != nil {
		return m.Attr
	}
	return nil
}

type MembershipChangeReply struct {
	Attr                 *MemberAttr `protobuf:"bytes,1,opt,name=attr,proto3" json:"attr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MembershipChangeReply) Reset()         { *m = MembershipChangeReply{} }
func (m *MembershipChangeReply) String() string { return proto.CompactTextString(m) }
func (*MembershipChangeReply) ProtoMessage()    {}
func (*MembershipChangeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{2}
}
func (m *MembershipChangeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MembershipChangeReply.Unmarshal(m, b)
}
func (m *MembershipChangeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MembershipChangeReply.Marshal(b, m, deterministic)
}
func (dst *MembershipChangeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MembershipChangeReply.Merge(dst, src)
}
func (m *MembershipChangeReply) XXX_Size() int {
	return xxx_messageInfo_MembershipChangeReply.Size(m)
}
func (m *MembershipChangeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MembershipChangeReply.DiscardUnknown(m)
}

var xxx_messageInfo_MembershipChangeReply proto.InternalMessageInfo

func (m *MembershipChangeReply) GetAttr() *MemberAttr {
	if m != nil {
		return m.Attr
	}
	return nil
}

type HardStateInfo struct {
	Term                 uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Commit               uint64   `protobuf:"varint,2,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HardStateInfo) Reset()         { *m = HardStateInfo{} }
func (m *HardStateInfo) String() string { return proto.CompactTextString(m) }
func (*HardStateInfo) ProtoMessage()    {}
func (*HardStateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{3}
}
func (m *HardStateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HardStateInfo.Unmarshal(m, b)
}
func (m *HardStateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HardStateInfo.Marshal(b, m, deterministic)
}
func (dst *HardStateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HardStateInfo.Merge(dst, src)
}
func (m *HardStateInfo) XXX_Size() int {
	return xxx_messageInfo_HardStateInfo.Size(m)
}
func (m *HardStateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HardStateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HardStateInfo proto.InternalMessageInfo

func (m *HardStateInfo) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *HardStateInfo) GetCommit() uint64 {
	if m != nil {
		return m.Commit
	}
	return 0
}

// data types for raft support
// GetClusterInfoRequest
type GetClusterInfoRequest struct {
	BestBlockHash        []byte   `protobuf:"bytes,1,opt,name=bestBlockHash,proto3" json:"bestBlockHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClusterInfoRequest) Reset()         { *m = GetClusterInfoRequest{} }
func (m *GetClusterInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterInfoRequest) ProtoMessage()    {}
func (*GetClusterInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{4}
}
func (m *GetClusterInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClusterInfoRequest.Unmarshal(m, b)
}
func (m *GetClusterInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClusterInfoRequest.Marshal(b, m, deterministic)
}
func (dst *GetClusterInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClusterInfoRequest.Merge(dst, src)
}
func (m *GetClusterInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetClusterInfoRequest.Size(m)
}
func (m *GetClusterInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClusterInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClusterInfoRequest proto.InternalMessageInfo

func (m *GetClusterInfoRequest) GetBestBlockHash() []byte {
	if m != nil {
		return m.BestBlockHash
	}
	return nil
}

type GetClusterInfoResponse struct {
	ChainID              []byte         `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	ClusterID            uint64         `protobuf:"varint,2,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	Error                string         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	MbrAttrs             []*MemberAttr  `protobuf:"bytes,4,rep,name=mbrAttrs,proto3" json:"mbrAttrs,omitempty"`
	BestBlockNo          uint64         `protobuf:"varint,5,opt,name=bestBlockNo,proto3" json:"bestBlockNo,omitempty"`
	HardStateInfo        *HardStateInfo `protobuf:"bytes,6,opt,name=hardStateInfo,proto3" json:"hardStateInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetClusterInfoResponse) Reset()         { *m = GetClusterInfoResponse{} }
func (m *GetClusterInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetClusterInfoResponse) ProtoMessage()    {}
func (*GetClusterInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{5}
}
func (m *GetClusterInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClusterInfoResponse.Unmarshal(m, b)
}
func (m *GetClusterInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClusterInfoResponse.Marshal(b, m, deterministic)
}
func (dst *GetClusterInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClusterInfoResponse.Merge(dst, src)
}
func (m *GetClusterInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetClusterInfoResponse.Size(m)
}
func (m *GetClusterInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClusterInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetClusterInfoResponse proto.InternalMessageInfo

func (m *GetClusterInfoResponse) GetChainID() []byte {
	if m != nil {
		return m.ChainID
	}
	return nil
}

func (m *GetClusterInfoResponse) GetClusterID() uint64 {
	if m != nil {
		return m.ClusterID
	}
	return 0
}

func (m *GetClusterInfoResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetClusterInfoResponse) GetMbrAttrs() []*MemberAttr {
	if m != nil {
		return m.MbrAttrs
	}
	return nil
}

func (m *GetClusterInfoResponse) GetBestBlockNo() uint64 {
	if m != nil {
		return m.BestBlockNo
	}
	return 0
}

func (m *GetClusterInfoResponse) GetHardStateInfo() *HardStateInfo {
	if m != nil {
		return m.HardStateInfo
	}
	return nil
}

type ConfChangeProgress struct {
	State                ConfChangeState `protobuf:"varint,1,opt,name=State,proto3,enum=types.ConfChangeState" json:"State,omitempty"`
	Err                  string          `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	Members              []*MemberAttr   `protobuf:"bytes,3,rep,name=Members,proto3" json:"Members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ConfChangeProgress) Reset()         { *m = ConfChangeProgress{} }
func (m *ConfChangeProgress) String() string { return proto.CompactTextString(m) }
func (*ConfChangeProgress) ProtoMessage()    {}
func (*ConfChangeProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{6}
}
func (m *ConfChangeProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfChangeProgress.Unmarshal(m, b)
}
func (m *ConfChangeProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfChangeProgress.Marshal(b, m, deterministic)
}
func (dst *ConfChangeProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfChangeProgress.Merge(dst, src)
}
func (m *ConfChangeProgress) XXX_Size() int {
	return xxx_messageInfo_ConfChangeProgress.Size(m)
}
func (m *ConfChangeProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfChangeProgress.DiscardUnknown(m)
}

var xxx_messageInfo_ConfChangeProgress proto.InternalMessageInfo

func (m *ConfChangeProgress) GetState() ConfChangeState {
	if m != nil {
		return m.State
	}
	return ConfChangeState_CONF_CHANGE_STATE_PROPOSED
}

func (m *ConfChangeProgress) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *ConfChangeProgress) GetMembers() []*MemberAttr {
	if m != nil {
		return m.Members
	}
	return nil
}

// SnapshotResponse is response message of receiving peer
type SnapshotResponse struct {
	Status               ResultStatus `protobuf:"varint,1,opt,name=status,proto3,enum=types.ResultStatus" json:"status,omitempty"`
	Message              string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SnapshotResponse) Reset()         { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string { return proto.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()    {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_58cb27e06f04b826, []int{7}
}
func (m *SnapshotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotResponse.Unmarshal(m, b)
}
func (m *SnapshotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotResponse.Marshal(b, m, deterministic)
}
func (dst *SnapshotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotResponse.Merge(dst, src)
}
func (m *SnapshotResponse) XXX_Size() int {
	return xxx_messageInfo_SnapshotResponse.Size(m)
}
func (m *SnapshotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotResponse proto.InternalMessageInfo

func (m *SnapshotResponse) GetStatus() ResultStatus {
	if m != nil {
		return m.Status
	}
	return ResultStatus_OK
}

func (m *SnapshotResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MemberAttr)(nil), "types.MemberAttr")
	proto.RegisterType((*MembershipChange)(nil), "types.MembershipChange")
	proto.RegisterType((*MembershipChangeReply)(nil), "types.MembershipChangeReply")
	proto.RegisterType((*HardStateInfo)(nil), "types.HardStateInfo")
	proto.RegisterType((*GetClusterInfoRequest)(nil), "types.GetClusterInfoRequest")
	proto.RegisterType((*GetClusterInfoResponse)(nil), "types.GetClusterInfoResponse")
	proto.RegisterType((*ConfChangeProgress)(nil), "types.ConfChangeProgress")
	proto.RegisterType((*SnapshotResponse)(nil), "types.SnapshotResponse")
	proto.RegisterEnum("types.MembershipChangeType", MembershipChangeType_name, MembershipChangeType_value)
	proto.RegisterEnum("types.ConfChangeState", ConfChangeState_name, ConfChangeState_value)
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor_raft_58cb27e06f04b826) }

var fileDescriptor_raft_58cb27e06f04b826 = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x13, 0x27, 0xa5, 0xd3, 0xa6, 0xb8, 0x4b, 0x5b, 0x4c, 0x0b, 0x28, 0xb2, 0x40, 0x8a,
	0x5a, 0x08, 0x52, 0x38, 0x01, 0x02, 0xc9, 0x8d, 0x4d, 0x63, 0x89, 0x7c, 0x68, 0x1d, 0x55, 0xe2,
	0x14, 0x6d, 0x92, 0x4d, 0x1c, 0x11, 0x7f, 0xb0, 0xbb, 0x39, 0xe4, 0xc8, 0x8d, 0xff, 0xcb, 0x1f,
	0x40, 0x5e, 0xaf, 0xdd, 0x26, 0x2d, 0xdc, 0x76, 0x66, 0xde, 0xcc, 0xbc, 0x79, 0x33, 0x36, 0x00,
	0x23, 0x33, 0xd1, 0x4c, 0x58, 0x2c, 0x62, 0x54, 0x11, 0xeb, 0x84, 0xf2, 0xb3, 0xbd, 0xa4, 0x95,
	0x64, 0x1e, 0x6b, 0x0c, 0xd0, 0xa5, 0xe1, 0x98, 0x32, 0x5b, 0x08, 0x86, 0x0e, 0xa1, 0xe4, 0x39,
	0xa6, 0x56, 0xd7, 0x1a, 0x3a, 0x2e, 0x79, 0x0e, 0x42, 0xa0, 0x47, 0x24, 0xa4, 0x66, 0xa9, 0xae,
	0x35, 0xf6, 0xb0, 0x7c, 0x23, 0x13, 0x76, 0xc9, 0x74, 0xca, 0x28, 0xe7, 0x66, 0x59, 0xba, 0x73,
	0x13, 0x9d, 0x42, 0x35, 0xa1, 0x94, 0x79, 0x8e, 0xa9, 0xd7, 0xb5, 0xc6, 0x01, 0x56, 0x96, 0xf5,
	0x5b, 0x03, 0x23, 0x6b, 0xc2, 0x83, 0x45, 0xd2, 0x0e, 0x48, 0x34, 0xa7, 0xe8, 0x1d, 0xe8, 0x29,
	0x19, 0xd9, 0xec, 0xb0, 0x75, 0xde, 0x94, 0xcc, 0x9a, 0xdb, 0xb0, 0xe1, 0x3a, 0xa1, 0x58, 0x02,
	0xd1, 0x73, 0xd8, 0x63, 0xf4, 0xe7, 0x8a, 0x72, 0xe1, 0x39, 0x92, 0x90, 0x8e, 0x6f, 0x1d, 0xe8,
	0x35, 0xe8, 0x44, 0x08, 0x26, 0x29, 0xed, 0xb7, 0x8e, 0x36, 0xca, 0xa5, 0xa3, 0x61, 0x19, 0xb6,
	0xbe, 0xc0, 0xc9, 0x76, 0x0b, 0x4c, 0x93, 0xe5, 0xba, 0xc8, 0xd7, 0xfe, 0x9f, 0xff, 0x09, 0x6a,
	0x1d, 0xc2, 0xa6, 0xbe, 0x20, 0x82, 0x7a, 0xd1, 0x2c, 0x4e, 0x15, 0x12, 0x94, 0x85, 0x4a, 0x33,
	0xf9, 0x4e, 0x75, 0x98, 0xc4, 0x61, 0xb8, 0x10, 0x8a, 0xa6, 0xb2, 0xac, 0xcf, 0x70, 0x72, 0x4d,
	0x45, 0x7b, 0xb9, 0xe2, 0x82, 0xb2, 0x34, 0x1b, 0x67, 0xf4, 0xd1, 0x2b, 0xa8, 0x8d, 0x29, 0x17,
	0x57, 0xcb, 0x78, 0xf2, 0xa3, 0x43, 0x78, 0x20, 0xab, 0x1d, 0xe0, 0x4d, 0xa7, 0xf5, 0x47, 0x83,
	0xd3, 0xed, 0x7c, 0x9e, 0xc4, 0x11, 0x97, 0x3b, 0x99, 0x04, 0x64, 0x11, 0xa9, 0xe5, 0x1d, 0xe0,
	0xdc, 0x4c, 0x55, 0x9b, 0xa8, 0x84, 0x42, 0xb5, 0xc2, 0x81, 0x8e, 0xa1, 0x42, 0x19, 0x8b, 0x99,
	0xda, 0x64, 0x66, 0xa0, 0xb7, 0xf0, 0x28, 0x1c, 0xcb, 0xa9, 0xb9, 0xa9, 0xd7, 0xcb, 0x0f, 0xeb,
	0x51, 0x40, 0x50, 0x1d, 0xf6, 0x0b, 0xa2, 0xbd, 0xd8, 0xac, 0xc8, 0x26, 0x77, 0x5d, 0xe8, 0x23,
	0xd4, 0x82, 0xbb, 0xaa, 0x99, 0x55, 0xa9, 0xf2, 0xb1, 0xaa, 0xba, 0xa1, 0x28, 0xde, 0x84, 0x5a,
	0xbf, 0x34, 0x40, 0xed, 0x38, 0x9a, 0x65, 0xcb, 0x1a, 0xb0, 0x78, 0x2e, 0x6f, 0xed, 0x0d, 0x54,
	0x24, 0x46, 0xdd, 0xcf, 0xa9, 0x2a, 0x75, 0x8b, 0x94, 0x51, 0x9c, 0x81, 0x90, 0x01, 0x65, 0x97,
	0x31, 0x75, 0xc6, 0xe9, 0x13, 0x5d, 0xc2, 0xae, 0x3a, 0x04, 0xb3, 0xfc, 0xaf, 0x11, 0x73, 0x84,
	0xf5, 0x1d, 0x0c, 0x3f, 0x22, 0x09, 0x0f, 0x62, 0x51, 0x48, 0x7e, 0x09, 0x55, 0x2e, 0x88, 0x58,
	0x71, 0xc5, 0xe0, 0x89, 0xca, 0xc7, 0x94, 0xaf, 0x96, 0xc2, 0x97, 0x21, 0xac, 0x20, 0xe9, 0x7e,
	0x42, 0xca, 0x39, 0x99, 0xe7, 0x9f, 0x52, 0x6e, 0x5e, 0x7c, 0x80, 0xe3, 0x87, 0x6e, 0x1e, 0x1d,
	0x02, 0xd8, 0x8e, 0x33, 0xea, 0xba, 0xdd, 0x2b, 0x17, 0x1b, 0x3b, 0xe8, 0x08, 0x6a, 0xd8, 0xed,
	0xf6, 0x6f, 0xdc, 0xdc, 0xa5, 0x5d, 0x84, 0xf0, 0x78, 0x6b, 0x5c, 0xf4, 0x12, 0xce, 0xda, 0xfd,
	0xde, 0xd7, 0x51, 0xbb, 0x63, 0xf7, 0xae, 0xdd, 0x91, 0x3f, 0xb4, 0x87, 0xee, 0x68, 0x80, 0xfb,
	0x83, 0xbe, 0xef, 0x3a, 0xc6, 0x0e, 0x3a, 0x87, 0xa7, 0xf7, 0xe3, 0xbe, 0x7d, 0xe3, 0x3a, 0x86,
	0x86, 0x5e, 0xc0, 0xb3, 0xfb, 0x41, 0x7b, 0x30, 0xf8, 0xe6, 0xb9, 0x8e, 0x51, 0x1a, 0x57, 0xe5,
	0x0f, 0xe3, 0xfd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x67, 0x2b, 0xa7, 0x50, 0x04, 0x00,
	0x00,
}
