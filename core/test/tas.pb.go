// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tas.proto

package tas_middleware_test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BlockHeader struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=Hash" json:"Hash,omitempty"`
	Height               *string  `protobuf:"bytes,2,opt,name=Height" json:"Height,omitempty"`
	PreHash              []byte   `protobuf:"bytes,3,opt,name=PreHash" json:"PreHash,omitempty"`
	Elapsed              *int32   `protobuf:"varint,4,opt,name=Elapsed" json:"Elapsed,omitempty"`
	ProveValue           []byte   `protobuf:"bytes,5,opt,name=ProveValue" json:"ProveValue,omitempty"`
	TotalQN              *uint64  `protobuf:"varint,6,opt,name=TotalQN" json:"TotalQN,omitempty"`
	CurTime              *int64   `protobuf:"varint,7,opt,name=CurTime" json:"CurTime,omitempty"`
	Castor               []byte   `protobuf:"bytes,8,opt,name=Castor" json:"Castor,omitempty"`
	GroupId              []byte   `protobuf:"bytes,9,opt,name=GroupId" json:"GroupId,omitempty"`
	Signature            []byte   `protobuf:"bytes,10,opt,name=Signature" json:"Signature,omitempty"`
	Nonce                *int32   `protobuf:"varint,11,opt,name=Nonce" json:"Nonce,omitempty"`
	TxTree               []byte   `protobuf:"bytes,12,opt,name=TxTree" json:"TxTree,omitempty"`
	ReceiptTree          []byte   `protobuf:"bytes,13,opt,name=ReceiptTree" json:"ReceiptTree,omitempty"`
	StateTree            []byte   `protobuf:"bytes,14,opt,name=StateTree" json:"StateTree,omitempty"`
	ExtraData            []byte   `protobuf:"bytes,15,opt,name=ExtraData" json:"ExtraData,omitempty"`
	Random               []byte   `protobuf:"bytes,16,opt,name=Random" json:"Random,omitempty"`
	GasFee               *uint64  `protobuf:"varint,17,opt,name=GasFee" json:"GasFee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_045c4e31cfb1792c, []int{0}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetHeight() string {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return ""
}

func (m *BlockHeader) GetPreHash() []byte {
	if m != nil {
		return m.PreHash
	}
	return nil
}

func (m *BlockHeader) GetElapsed() int32 {
	if m != nil && m.Elapsed != nil {
		return *m.Elapsed
	}
	return 0
}

func (m *BlockHeader) GetProveValue() []byte {
	if m != nil {
		return m.ProveValue
	}
	return nil
}

func (m *BlockHeader) GetTotalQN() uint64 {
	if m != nil && m.TotalQN != nil {
		return *m.TotalQN
	}
	return 0
}

func (m *BlockHeader) GetCurTime() int64 {
	if m != nil && m.CurTime != nil {
		return *m.CurTime
	}
	return 0
}

func (m *BlockHeader) GetCastor() []byte {
	if m != nil {
		return m.Castor
	}
	return nil
}

func (m *BlockHeader) GetGroupId() []byte {
	if m != nil {
		return m.GroupId
	}
	return nil
}

func (m *BlockHeader) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *BlockHeader) GetNonce() int32 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetTxTree() []byte {
	if m != nil {
		return m.TxTree
	}
	return nil
}

func (m *BlockHeader) GetReceiptTree() []byte {
	if m != nil {
		return m.ReceiptTree
	}
	return nil
}

func (m *BlockHeader) GetStateTree() []byte {
	if m != nil {
		return m.StateTree
	}
	return nil
}

func (m *BlockHeader) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

func (m *BlockHeader) GetRandom() []byte {
	if m != nil {
		return m.Random
	}
	return nil
}

func (m *BlockHeader) GetGasFee() uint64 {
	if m != nil && m.GasFee != nil {
		return *m.GasFee
	}
	return 0
}

type Block struct {
	Header               *BlockHeader   `protobuf:"bytes,1,req,name=Header" json:"Header,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_045c4e31cfb1792c, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Transaction struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
	Nonce                *uint64  `protobuf:"varint,3,req,name=Nonce" json:"Nonce,omitempty"`
	Target               []byte   `protobuf:"bytes,4,opt,name=Target" json:"Target,omitempty"`
	GasLimit             []byte   `protobuf:"bytes,5,req,name=GasLimit" json:"GasLimit,omitempty"`
	GasPrice             []byte   `protobuf:"bytes,6,req,name=GasPrice" json:"GasPrice,omitempty"`
	Hash                 []byte   `protobuf:"bytes,7,req,name=Hash" json:"Hash,omitempty"`
	ExtraData            []byte   `protobuf:"bytes,8,opt,name=ExtraData" json:"ExtraData,omitempty"`
	Type                 *int32   `protobuf:"varint,9,req,name=Type" json:"Type,omitempty"`
	Sign                 []byte   `protobuf:"bytes,10,opt,name=Sign" json:"Sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_045c4e31cfb1792c, []int{2}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

func (m *Transaction) GetTarget() []byte {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Transaction) GetGasLimit() []byte {
	if m != nil {
		return m.GasLimit
	}
	return nil
}

func (m *Transaction) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

func (m *Transaction) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type BlockResponseMsg struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=Blocks" json:"Blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockResponseMsg) Reset()         { *m = BlockResponseMsg{} }
func (m *BlockResponseMsg) String() string { return proto.CompactTextString(m) }
func (*BlockResponseMsg) ProtoMessage()    {}
func (*BlockResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_045c4e31cfb1792c, []int{3}
}

func (m *BlockResponseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockResponseMsg.Unmarshal(m, b)
}
func (m *BlockResponseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockResponseMsg.Marshal(b, m, deterministic)
}
func (m *BlockResponseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockResponseMsg.Merge(m, src)
}
func (m *BlockResponseMsg) XXX_Size() int {
	return xxx_messageInfo_BlockResponseMsg.Size(m)
}
func (m *BlockResponseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockResponseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BlockResponseMsg proto.InternalMessageInfo

func (m *BlockResponseMsg) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "tas.middleware.test.BlockHeader")
	proto.RegisterType((*Block)(nil), "tas.middleware.test.Block")
	proto.RegisterType((*Transaction)(nil), "tas.middleware.test.Transaction")
	proto.RegisterType((*BlockResponseMsg)(nil), "tas.middleware.test.BlockResponseMsg")
}

func init() { proto.RegisterFile("tas.proto", fileDescriptor_045c4e31cfb1792c) }

var fileDescriptor_045c4e31cfb1792c = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x95, 0xd3, 0xa6, 0x1f, 0x6e, 0x81, 0xc5, 0xac, 0xd0, 0x68, 0x85, 0x90, 0xd5, 0x53, 0x4e,
	0x3d, 0xec, 0x89, 0x33, 0xfb, 0xd1, 0x22, 0xc1, 0xaa, 0x98, 0x88, 0xbb, 0xd5, 0x8c, 0xba, 0x11,
	0x6d, 0x1c, 0xd9, 0x2e, 0x2c, 0xbf, 0x80, 0x5f, 0xc7, 0xef, 0xe1, 0x8a, 0x3c, 0x76, 0xdb, 0x2c,
	0x62, 0x6f, 0xf3, 0xde, 0x9b, 0xc9, 0x38, 0xf3, 0x1e, 0x1f, 0x7b, 0xed, 0xe6, 0xad, 0x35, 0xde,
	0x88, 0x57, 0xa1, 0xdc, 0xd5, 0x55, 0xb5, 0xc5, 0x1f, 0xda, 0xe2, 0xdc, 0xa3, 0xf3, 0xb3, 0xdf,
	0x3d, 0x3e, 0x79, 0xbf, 0x35, 0xeb, 0x6f, 0x4b, 0xd4, 0x15, 0x5a, 0x21, 0x78, 0x7f, 0xa9, 0xdd,
	0x3d, 0x30, 0xc9, 0x8a, 0xa9, 0xa2, 0x5a, 0xbc, 0xe6, 0x83, 0x25, 0xd6, 0x9b, 0x7b, 0x0f, 0x99,
	0x64, 0xc5, 0x58, 0x25, 0x24, 0x80, 0x0f, 0x57, 0x16, 0xa9, 0xbd, 0x47, 0xed, 0x07, 0x18, 0x94,
	0x9b, 0xad, 0x6e, 0x1d, 0x56, 0xd0, 0x97, 0xac, 0xc8, 0xd5, 0x01, 0x8a, 0xb7, 0x9c, 0xaf, 0xac,
	0xf9, 0x8e, 0x5f, 0xf5, 0x76, 0x8f, 0x90, 0xd3, 0x58, 0x87, 0x09, 0x93, 0xa5, 0xf1, 0x7a, 0xfb,
	0xf9, 0x0e, 0x06, 0x92, 0x15, 0x7d, 0x75, 0x80, 0x41, 0xb9, 0xda, 0xdb, 0xb2, 0xde, 0x21, 0x0c,
	0x25, 0x2b, 0x7a, 0xea, 0x00, 0xc3, 0xfb, 0xae, 0xb4, 0xf3, 0xc6, 0xc2, 0x88, 0xbe, 0x97, 0x50,
	0x98, 0x58, 0x58, 0xb3, 0x6f, 0x3f, 0x54, 0x30, 0x8e, 0xef, 0x4b, 0x50, 0xbc, 0xe1, 0xe3, 0x2f,
	0xf5, 0xa6, 0xd1, 0x7e, 0x6f, 0x11, 0x38, 0x69, 0x27, 0x42, 0x9c, 0xf3, 0xfc, 0xce, 0x34, 0x6b,
	0x84, 0x09, 0xbd, 0x3d, 0x82, 0xb0, 0xa5, 0x7c, 0x28, 0x2d, 0x22, 0x4c, 0xe3, 0x96, 0x88, 0x84,
	0xe4, 0x13, 0x85, 0x6b, 0xac, 0x5b, 0x4f, 0xe2, 0x33, 0x12, 0xbb, 0x14, 0x6d, 0xf3, 0xda, 0x23,
	0xe9, 0xcf, 0xd3, 0xb6, 0x03, 0x11, 0xd4, 0x9b, 0x07, 0x6f, 0xf5, 0xb5, 0xf6, 0x1a, 0x5e, 0x44,
	0xf5, 0x48, 0x84, 0xad, 0x4a, 0x37, 0x95, 0xd9, 0xc1, 0x59, 0xdc, 0x1a, 0x51, 0xe0, 0x17, 0xda,
	0xdd, 0x22, 0xc2, 0x4b, 0x3a, 0x53, 0x42, 0xb3, 0x5f, 0x8c, 0xe7, 0xe4, 0xa7, 0x78, 0x17, 0x5c,
	0x0b, 0x9e, 0x02, 0x93, 0x59, 0x31, 0xb9, 0x94, 0xf3, 0xff, 0xf8, 0x3f, 0xef, 0x78, 0xaf, 0x52,
	0xbf, 0xb8, 0xe6, 0x53, 0x6f, 0x75, 0xe3, 0xf4, 0xda, 0xd7, 0xa6, 0x71, 0x90, 0xc9, 0xde, 0x93,
	0xf3, 0xe5, 0xa9, 0x51, 0x3d, 0x9a, 0x9a, 0xfd, 0x61, 0x7c, 0xd2, 0x51, 0x43, 0xb2, 0xe8, 0x17,
	0x53, 0xb2, 0xe8, 0xef, 0xce, 0x79, 0x1e, 0x83, 0x90, 0x11, 0x19, 0xc1, 0xe9, 0xfe, 0x3d, 0x99,
	0x15, 0xfd, 0xee, 0xfd, 0xb5, 0xdd, 0xa0, 0xa7, 0x48, 0x85, 0xfb, 0x13, 0x12, 0x17, 0x7c, 0xb4,
	0xd0, 0xee, 0x63, 0xbd, 0xab, 0x3d, 0xe4, 0x32, 0x2b, 0xa6, 0xea, 0x88, 0x93, 0xb6, 0xb2, 0xf5,
	0x1a, 0x61, 0x70, 0xd4, 0x08, 0x1f, 0x93, 0x3e, 0x24, 0x3e, 0x26, 0xfd, 0x91, 0x17, 0xa3, 0x7f,
	0xbd, 0x10, 0xbc, 0x5f, 0xfe, 0x6c, 0x11, 0xc6, 0x32, 0x2b, 0x72, 0x45, 0x75, 0xe0, 0x42, 0x70,
	0x52, 0x88, 0xa8, 0x9e, 0xdd, 0xf2, 0x33, 0x3a, 0xab, 0x42, 0xd7, 0x9a, 0xc6, 0xe1, 0x27, 0xb7,
	0x11, 0x97, 0x7c, 0x40, 0x9c, 0x03, 0x46, 0xd7, 0xbc, 0x78, 0xda, 0x0d, 0x95, 0x3a, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0xc8, 0xb9, 0x56, 0xbc, 0x03, 0x00, 0x00,
}
