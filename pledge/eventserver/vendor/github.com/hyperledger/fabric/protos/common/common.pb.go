// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto
	common/configtx.proto
	common/configuration.proto
	common/ledger.proto
	common/policies.proto

It has these top-level messages:
	LastConfig
	Metadata
	MetadataSignature
	Header
	ChannelHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigEnvelope
	ConfigGroupSchema
	ConfigValueSchema
	ConfigPolicySchema
	Config
	ConfigUpdateEnvelope
	ConfigUpdate
	ConfigGroup
	ConfigValue
	ConfigPolicy
	ConfigSignature
	HashingAlgorithm
	BlockDataHashingStructure
	OrdererAddresses
	Consortium
	BlockchainInfo
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
	ImplicitMetaPolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN                  Status = 0
	Status_SUCCESS                  Status = 200
	Status_BAD_REQUEST              Status = 400
	Status_FORBIDDEN                Status = 403
	Status_NOT_FOUND                Status = 404
	Status_REQUEST_ENTITY_TOO_LARGE Status = 413
	Status_INTERNAL_SERVER_ERROR    Status = 500
	Status_SERVICE_UNAVAILABLE      Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	413: "REQUEST_ENTITY_TOO_LARGE",
	500: "INTERNAL_SERVER_ERROR",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":                  0,
	"SUCCESS":                  200,
	"BAD_REQUEST":              400,
	"FORBIDDEN":                403,
	"NOT_FOUND":                404,
	"REQUEST_ENTITY_TOO_LARGE": 413,
	"INTERNAL_SERVER_ERROR":    500,
	"SERVICE_UNAVAILABLE":      503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HeaderType int32

const (
	HeaderType_MESSAGE              HeaderType = 0
	HeaderType_CONFIG               HeaderType = 1
	HeaderType_CONFIG_UPDATE        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION HeaderType = 3
	HeaderType_ORDERER_TRANSACTION  HeaderType = 4
	HeaderType_DELIVER_SEEK_INFO    HeaderType = 5
	HeaderType_CHAINCODE_PACKAGE    HeaderType = 6
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIG",
	2: "CONFIG_UPDATE",
	3: "ENDORSER_TRANSACTION",
	4: "ORDERER_TRANSACTION",
	5: "DELIVER_SEEK_INFO",
	6: "CHAINCODE_PACKAGE",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":              0,
	"CONFIG":               1,
	"CONFIG_UPDATE":        2,
	"ENDORSER_TRANSACTION": 3,
	"ORDERER_TRANSACTION":  4,
	"DELIVER_SEEK_INFO":    5,
	"CHAINCODE_PACKAGE":    6,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// This enum enlists indexes of the block metadata array
type BlockMetadataIndex int32

const (
	BlockMetadataIndex_SIGNATURES          BlockMetadataIndex = 0
	BlockMetadataIndex_LAST_CONFIG         BlockMetadataIndex = 1
	BlockMetadataIndex_TRANSACTIONS_FILTER BlockMetadataIndex = 2
	BlockMetadataIndex_ORDERER             BlockMetadataIndex = 3
)

var BlockMetadataIndex_name = map[int32]string{
	0: "SIGNATURES",
	1: "LAST_CONFIG",
	2: "TRANSACTIONS_FILTER",
	3: "ORDERER",
}
var BlockMetadataIndex_value = map[string]int32{
	"SIGNATURES":          0,
	"LAST_CONFIG":         1,
	"TRANSACTIONS_FILTER": 2,
	"ORDERER":             3,
}

func (x BlockMetadataIndex) String() string {
	return proto.EnumName(BlockMetadataIndex_name, int32(x))
}
func (BlockMetadataIndex) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// LastConfig is the encoded value for the Metadata message which is encoded in the LAST_CONFIGURATION block metadata index
type LastConfig struct {
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (m *LastConfig) Reset()                    { *m = LastConfig{} }
func (m *LastConfig) String() string            { return proto.CompactTextString(m) }
func (*LastConfig) ProtoMessage()               {}
func (*LastConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LastConfig) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// Metadata is a common structure to be used to encode block metadata
type Metadata struct {
	Value      []byte               `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Signatures []*MetadataSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Metadata) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Metadata) GetSignatures() []*MetadataSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type MetadataSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MetadataSignature) Reset()                    { *m = MetadataSignature{} }
func (m *MetadataSignature) String() string            { return proto.CompactTextString(m) }
func (*MetadataSignature) ProtoMessage()               {}
func (*MetadataSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetadataSignature) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

func (m *MetadataSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Header struct {
	ChannelHeader   []byte `protobuf:"bytes,1,opt,name=channel_header,json=channelHeader,proto3" json:"channel_header,omitempty"`
	SignatureHeader []byte `protobuf:"bytes,2,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Header) GetChannelHeader() []byte {
	if m != nil {
		return m.ChannelHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() []byte {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChannelHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Identifier of the channel this message is bound for
	ChannelId string `protobuf:"bytes,4,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxId string `protobuf:"bytes,5,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *ChannelHeader) Reset()                    { *m = ChannelHeader{} }
func (m *ChannelHeader) String() string            { return proto.CompactTextString(m) }
func (*ChannelHeader) ProtoMessage()               {}
func (*ChannelHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ChannelHeader) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ChannelHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ChannelHeader) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChannelHeader) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChannelHeader) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *ChannelHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ChannelHeader) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, specified as a certificate chain
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SignatureHeader) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *SignatureHeader) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Envelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// BlockHeader is the element of the block which forms the block chain
// The block header is hashed using the configured chain hashing algorithm
// over the ASN.1 encoding of the BlockHeader
type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BlockHeader) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockHeader) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *BlockHeader) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BlockData) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *BlockMetadata) Reset()                    { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string            { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()               {}
func (*BlockMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *BlockMetadata) GetMetadata() [][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*LastConfig)(nil), "common.LastConfig")
	proto.RegisterType((*Metadata)(nil), "common.Metadata")
	proto.RegisterType((*MetadataSignature)(nil), "common.MetadataSignature")
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChannelHeader)(nil), "common.ChannelHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
	proto.RegisterEnum("common.BlockMetadataIndex", BlockMetadataIndex_name, BlockMetadataIndex_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xdf, 0x6e, 0xe3, 0xc4,
	0x1b, 0xad, 0xe3, 0xfc, 0x69, 0xbe, 0x34, 0xad, 0x3b, 0xd9, 0xfe, 0xd6, 0xbf, 0xc2, 0x6a, 0x23,
	0xc3, 0xa2, 0xd2, 0x4a, 0x89, 0x28, 0x37, 0x70, 0xe9, 0xd8, 0x93, 0xd6, 0x6a, 0xd6, 0x2e, 0x63,
	0x67, 0x11, 0xbb, 0x48, 0x96, 0x93, 0x4c, 0x13, 0x8b, 0xc4, 0x8e, 0x6c, 0xa7, 0x6a, 0x5f, 0x02,
	0x21, 0xc1, 0x0d, 0x17, 0xbc, 0x00, 0x4f, 0xc2, 0x5b, 0xf0, 0x12, 0x48, 0xdc, 0xa2, 0xf1, 0x8c,
	0xbd, 0x49, 0x59, 0x89, 0xab, 0xcc, 0x39, 0x73, 0x3c, 0xdf, 0x99, 0xf3, 0x7d, 0xb1, 0xa1, 0x33,
	0x8d, 0x57, 0xab, 0x38, 0xea, 0xf3, 0x9f, 0xde, 0x3a, 0x89, 0xb3, 0x18, 0xd5, 0x39, 0x3a, 0x7d,
	0x39, 0x8f, 0xe3, 0xf9, 0x92, 0xf6, 0x73, 0x76, 0xb2, 0xb9, 0xeb, 0x67, 0xe1, 0x8a, 0xa6, 0x59,
	0xb0, 0x5a, 0x73, 0xa1, 0xa6, 0x01, 0x8c, 0x82, 0x34, 0x33, 0xe2, 0xe8, 0x2e, 0x9c, 0xa3, 0x67,
	0x50, 0x0b, 0xa3, 0x19, 0x7d, 0x50, 0xa5, 0xae, 0x74, 0x56, 0x25, 0x1c, 0x68, 0xef, 0x60, 0xff,
	0x35, 0xcd, 0x82, 0x59, 0x90, 0x05, 0x4c, 0x71, 0x1f, 0x2c, 0x37, 0x34, 0x57, 0x1c, 0x10, 0x0e,
	0xd0, 0xd7, 0x00, 0x69, 0x38, 0x8f, 0x82, 0x6c, 0x93, 0xd0, 0x54, 0xad, 0x74, 0xe5, 0xb3, 0xd6,
	0xe5, 0xff, 0x7b, 0xc2, 0x51, 0xf1, 0xac, 0x5b, 0x28, 0xc8, 0x96, 0x58, 0xfb, 0x1e, 0x8e, 0xff,
	0x25, 0x40, 0x9f, 0x83, 0x52, 0x4a, 0xfc, 0x05, 0x0d, 0x66, 0x34, 0x11, 0x05, 0x8f, 0x4a, 0xfe,
	0x3a, 0xa7, 0xd1, 0xc7, 0xd0, 0x2c, 0x29, 0xb5, 0x92, 0x6b, 0xde, 0x13, 0xda, 0x5b, 0xa8, 0x0b,
	0xdd, 0x2b, 0x38, 0x9c, 0x2e, 0x82, 0x28, 0xa2, 0xcb, 0xdd, 0x03, 0xdb, 0x82, 0x15, 0xb2, 0x0f,
	0x55, 0xae, 0x7c, 0xb0, 0xb2, 0xf6, 0xa7, 0x04, 0x6d, 0x63, 0xe7, 0x61, 0x04, 0xd5, 0xec, 0x71,
	0xcd, 0xb3, 0xa9, 0x91, 0x7c, 0x8d, 0x54, 0x68, 0xdc, 0xd3, 0x24, 0x0d, 0xe3, 0x28, 0x3f, 0xa7,
	0x46, 0x0a, 0x88, 0xbe, 0x82, 0x66, 0xd9, 0x0d, 0x55, 0xee, 0x4a, 0x67, 0xad, 0xcb, 0xd3, 0x1e,
	0xef, 0x57, 0xaf, 0xe8, 0x57, 0xcf, 0x2b, 0x14, 0xe4, 0xbd, 0x18, 0xbd, 0x00, 0x28, 0xee, 0x12,
	0xce, 0xd4, 0x6a, 0x57, 0x3a, 0x6b, 0x92, 0xa6, 0x60, 0xac, 0x19, 0xea, 0x40, 0x2d, 0x7b, 0x60,
	0x3b, 0xb5, 0x7c, 0xa7, 0x9a, 0x3d, 0x58, 0x33, 0xd6, 0x38, 0xba, 0x8e, 0xa7, 0x0b, 0xb5, 0xce,
	0x5b, 0x9b, 0x03, 0x96, 0x1e, 0x7d, 0xc8, 0x68, 0x94, 0xfb, 0x6b, 0xf0, 0xf4, 0x4a, 0x42, 0xd3,
	0xe1, 0xc8, 0x7d, 0x12, 0xb7, 0x0a, 0x8d, 0x69, 0x42, 0x83, 0x2c, 0x2e, 0xf2, 0x2b, 0x20, 0x2b,
	0x10, 0xc5, 0xd1, 0xb4, 0x68, 0x02, 0x07, 0x1a, 0x86, 0xc6, 0x6d, 0xf0, 0xb8, 0x8c, 0x83, 0x19,
	0xfa, 0x0c, 0xea, 0x5b, 0xc9, 0xb7, 0x2e, 0x0f, 0x8b, 0x01, 0xe1, 0x47, 0x13, 0xb1, 0xcb, 0x52,
	0x64, 0xd3, 0x20, 0xce, 0xc9, 0xd7, 0xda, 0x00, 0xf6, 0x71, 0x74, 0x4f, 0x97, 0x31, 0x4f, 0x74,
	0xcd, 0x8f, 0x2c, 0x2c, 0x08, 0xf8, 0x1f, 0xb3, 0xf0, 0xa3, 0x04, 0xb5, 0xc1, 0x32, 0x9e, 0xfe,
	0x80, 0x2e, 0x9e, 0x38, 0xe9, 0x14, 0x4e, 0xf2, 0xed, 0x27, 0x76, 0x5e, 0x6d, 0xd9, 0x69, 0x5d,
	0x1e, 0xef, 0x48, 0xcd, 0x20, 0x0b, 0xb8, 0x43, 0xf4, 0x05, 0xec, 0xaf, 0xc4, 0x1c, 0x8b, 0x66,
	0x9e, 0xec, 0x48, 0x8b, 0x21, 0x27, 0xa5, 0x4c, 0x9b, 0x43, 0x6b, 0xab, 0x20, 0xfa, 0x1f, 0xd4,
	0xa3, 0xcd, 0x6a, 0x22, 0x5c, 0x55, 0x89, 0x40, 0xe8, 0x13, 0x68, 0xaf, 0x13, 0x7a, 0x1f, 0xc6,
	0x9b, 0xd4, 0x5f, 0x04, 0xe9, 0x42, 0xdc, 0xec, 0xa0, 0x20, 0xaf, 0x83, 0x74, 0x81, 0x3e, 0x82,
	0x26, 0x3b, 0x93, 0x0b, 0xe4, 0x5c, 0xb0, 0xcf, 0x08, 0xb6, 0xa9, 0xbd, 0x84, 0x66, 0x69, 0xb7,
	0x8c, 0x57, 0xea, 0xca, 0x65, 0xbc, 0x17, 0xd0, 0xde, 0x31, 0x89, 0x4e, 0xb7, 0x6e, 0xc3, 0x85,
	0x25, 0x3e, 0xff, 0x5d, 0x82, 0xba, 0x9b, 0x05, 0xd9, 0x26, 0x45, 0x2d, 0x68, 0x8c, 0xed, 0x1b,
	0xdb, 0xf9, 0xd6, 0x56, 0xf6, 0xd0, 0x01, 0x34, 0xdc, 0xb1, 0x61, 0x60, 0xd7, 0x55, 0xfe, 0x90,
	0x90, 0x02, 0xad, 0x81, 0x6e, 0xfa, 0x04, 0x7f, 0x33, 0xc6, 0xae, 0xa7, 0xfc, 0x24, 0xa3, 0x43,
	0x68, 0x0e, 0x1d, 0x32, 0xb0, 0x4c, 0x13, 0xdb, 0xca, 0xcf, 0x39, 0xb6, 0x1d, 0xcf, 0x1f, 0x3a,
	0x63, 0xdb, 0x54, 0x7e, 0x91, 0xd1, 0x0b, 0x50, 0x85, 0xda, 0xc7, 0xb6, 0x67, 0x79, 0xdf, 0xf9,
	0x9e, 0xe3, 0xf8, 0x23, 0x9d, 0x5c, 0x61, 0xe5, 0x37, 0x19, 0x9d, 0xc2, 0x89, 0x65, 0x7b, 0x98,
	0xd8, 0xfa, 0xc8, 0x77, 0x31, 0x79, 0x83, 0x89, 0x8f, 0x09, 0x71, 0x88, 0xf2, 0x97, 0x8c, 0x54,
	0xe8, 0x30, 0xca, 0x32, 0xb0, 0x3f, 0xb6, 0xf5, 0x37, 0xba, 0x35, 0xd2, 0x07, 0x23, 0xac, 0xfc,
	0x2d, 0x9f, 0xff, 0x2a, 0x01, 0xf0, 0x7c, 0x3d, 0xf6, 0x6f, 0x6c, 0x41, 0xe3, 0x35, 0x76, 0x5d,
	0xfd, 0x0a, 0x2b, 0x7b, 0x08, 0xa0, 0x6e, 0x38, 0xf6, 0xd0, 0xba, 0x52, 0x24, 0x74, 0x0c, 0x6d,
	0xbe, 0xf6, 0xc7, 0xb7, 0xa6, 0xee, 0x61, 0xa5, 0x82, 0x54, 0x78, 0x86, 0x6d, 0xd3, 0x21, 0x2e,
	0x26, 0xbe, 0x47, 0x74, 0xdb, 0xd5, 0x0d, 0xcf, 0x72, 0x6c, 0x45, 0x46, 0xcf, 0xa1, 0xe3, 0x10,
	0x13, 0x93, 0x27, 0x1b, 0x55, 0x74, 0x02, 0xc7, 0x26, 0x1e, 0x59, 0xcc, 0x9b, 0x8b, 0xf1, 0x8d,
	0x6f, 0xd9, 0x43, 0x47, 0xa9, 0x31, 0xda, 0xb8, 0xd6, 0x2d, 0xdb, 0x70, 0x4c, 0xec, 0xdf, 0xea,
	0xc6, 0x0d, 0xab, 0x5f, 0x3f, 0x7f, 0x07, 0x68, 0x27, 0x75, 0x8b, 0xbd, 0x6d, 0xd1, 0x21, 0x80,
	0x6b, 0x5d, 0xd9, 0xba, 0x37, 0x26, 0xd8, 0x55, 0xf6, 0xd0, 0x11, 0xb4, 0x46, 0xba, 0xeb, 0xf9,
	0xa5, 0xd5, 0xe7, 0xd0, 0xd9, 0xaa, 0xea, 0xfa, 0x43, 0x6b, 0xe4, 0x61, 0xa2, 0x54, 0xd8, 0xe5,
	0x84, 0x2d, 0x45, 0x1e, 0xb8, 0xf0, 0x69, 0x9c, 0xcc, 0x7b, 0x8b, 0xc7, 0x35, 0x4d, 0x96, 0x74,
	0x36, 0xa7, 0x49, 0xef, 0x2e, 0x98, 0x24, 0xe1, 0x94, 0xbf, 0x5b, 0x52, 0x31, 0x9c, 0x6f, 0x2f,
	0xe6, 0x61, 0xb6, 0xd8, 0x4c, 0x18, 0xec, 0x6f, 0x89, 0xfb, 0x5c, 0xcc, 0x3f, 0x1c, 0xa9, 0xf8,
	0xb8, 0x4c, 0xea, 0x39, 0xfc, 0xf2, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xcb, 0xe1, 0xb4,
	0x74, 0x06, 0x00, 0x00,
}
