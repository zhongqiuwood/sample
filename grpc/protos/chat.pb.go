// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	OkMessage
	PeerAddress
	PeerID
	PeerEndpoint
	PeersMessage
	PeersAddresses
	HelloMessage
	TXOutput
	TXInput
	Transaction
	Block
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type OkMessage_Type int32

const (
	OkMessage_UNDEFINED               OkMessage_Type = 0
	OkMessage_DISC_HELLO              OkMessage_Type = 1
	OkMessage_DISC_DISCONNECT         OkMessage_Type = 2
	OkMessage_DISC_GET_PEERS          OkMessage_Type = 3
	OkMessage_DISC_PEERS              OkMessage_Type = 4
	OkMessage_DISC_NEWMSG             OkMessage_Type = 5
	OkMessage_CHAIN_TRANSACTION       OkMessage_Type = 6
	OkMessage_SYNC_GET_BLOCKS         OkMessage_Type = 11
	OkMessage_SYNC_BLOCKS             OkMessage_Type = 12
	OkMessage_SYNC_BLOCK_ADDED        OkMessage_Type = 13
	OkMessage_SYNC_STATE_GET_SNAPSHOT OkMessage_Type = 14
	OkMessage_SYNC_STATE_SNAPSHOT     OkMessage_Type = 15
	OkMessage_SYNC_STATE_GET_DELTAS   OkMessage_Type = 16
	OkMessage_SYNC_STATE_DELTAS       OkMessage_Type = 17
	OkMessage_RESPONSE                OkMessage_Type = 20
	OkMessage_CONSENSUS               OkMessage_Type = 21
)

var OkMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "DISC_HELLO",
	2:  "DISC_DISCONNECT",
	3:  "DISC_GET_PEERS",
	4:  "DISC_PEERS",
	5:  "DISC_NEWMSG",
	6:  "CHAIN_TRANSACTION",
	11: "SYNC_GET_BLOCKS",
	12: "SYNC_BLOCKS",
	13: "SYNC_BLOCK_ADDED",
	14: "SYNC_STATE_GET_SNAPSHOT",
	15: "SYNC_STATE_SNAPSHOT",
	16: "SYNC_STATE_GET_DELTAS",
	17: "SYNC_STATE_DELTAS",
	20: "RESPONSE",
	21: "CONSENSUS",
}
var OkMessage_Type_value = map[string]int32{
	"UNDEFINED":               0,
	"DISC_HELLO":              1,
	"DISC_DISCONNECT":         2,
	"DISC_GET_PEERS":          3,
	"DISC_PEERS":              4,
	"DISC_NEWMSG":             5,
	"CHAIN_TRANSACTION":       6,
	"SYNC_GET_BLOCKS":         11,
	"SYNC_BLOCKS":             12,
	"SYNC_BLOCK_ADDED":        13,
	"SYNC_STATE_GET_SNAPSHOT": 14,
	"SYNC_STATE_SNAPSHOT":     15,
	"SYNC_STATE_GET_DELTAS":   16,
	"SYNC_STATE_DELTAS":       17,
	"RESPONSE":                20,
	"CONSENSUS":               21,
}

func (x OkMessage_Type) String() string {
	return proto.EnumName(OkMessage_Type_name, int32(x))
}
func (OkMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type PeerEndpoint_Type int32

const (
	PeerEndpoint_UNDEFINED     PeerEndpoint_Type = 0
	PeerEndpoint_VALIDATOR     PeerEndpoint_Type = 1
	PeerEndpoint_NON_VALIDATOR PeerEndpoint_Type = 2
)

var PeerEndpoint_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "VALIDATOR",
	2: "NON_VALIDATOR",
}
var PeerEndpoint_Type_value = map[string]int32{
	"UNDEFINED":     0,
	"VALIDATOR":     1,
	"NON_VALIDATOR": 2,
}

func (x PeerEndpoint_Type) String() string {
	return proto.EnumName(PeerEndpoint_Type_name, int32(x))
}
func (PeerEndpoint_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type OkMessage struct {
	Type      OkMessage_Type             `protobuf:"varint,1,opt,name=type,enum=protos.OkMessage_Type" json:"type,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte                     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *OkMessage) Reset()                    { *m = OkMessage{} }
func (m *OkMessage) String() string            { return proto.CompactTextString(m) }
func (*OkMessage) ProtoMessage()               {}
func (*OkMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OkMessage) GetType() OkMessage_Type {
	if m != nil {
		return m.Type
	}
	return OkMessage_UNDEFINED
}

func (m *OkMessage) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *OkMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *OkMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type PeerAddress struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (m *PeerAddress) Reset()                    { *m = PeerAddress{} }
func (m *PeerAddress) String() string            { return proto.CompactTextString(m) }
func (*PeerAddress) ProtoMessage()               {}
func (*PeerAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PeerAddress) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PeerAddress) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type PeerID struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PeerID) Reset()                    { *m = PeerID{} }
func (m *PeerID) String() string            { return proto.CompactTextString(m) }
func (*PeerID) ProtoMessage()               {}
func (*PeerID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PeerID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PeerEndpoint struct {
	ID      string            `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Address string            `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Type    PeerEndpoint_Type `protobuf:"varint,3,opt,name=type,enum=protos.PeerEndpoint_Type" json:"type,omitempty"`
}

func (m *PeerEndpoint) Reset()                    { *m = PeerEndpoint{} }
func (m *PeerEndpoint) String() string            { return proto.CompactTextString(m) }
func (*PeerEndpoint) ProtoMessage()               {}
func (*PeerEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PeerEndpoint) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *PeerEndpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PeerEndpoint) GetType() PeerEndpoint_Type {
	if m != nil {
		return m.Type
	}
	return PeerEndpoint_UNDEFINED
}

type PeersMessage struct {
	Peers []*PeerEndpoint `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
}

func (m *PeersMessage) Reset()                    { *m = PeersMessage{} }
func (m *PeersMessage) String() string            { return proto.CompactTextString(m) }
func (*PeersMessage) ProtoMessage()               {}
func (*PeersMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PeersMessage) GetPeers() []*PeerEndpoint {
	if m != nil {
		return m.Peers
	}
	return nil
}

type PeersAddresses struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *PeersAddresses) Reset()                    { *m = PeersAddresses{} }
func (m *PeersAddresses) String() string            { return proto.CompactTextString(m) }
func (*PeersAddresses) ProtoMessage()               {}
func (*PeersAddresses) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PeersAddresses) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type HelloMessage struct {
	Endpoint *PeerEndpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *HelloMessage) Reset()                    { *m = HelloMessage{} }
func (m *HelloMessage) String() string            { return proto.CompactTextString(m) }
func (*HelloMessage) ProtoMessage()               {}
func (*HelloMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *HelloMessage) GetEndpoint() *PeerEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type TXOutput struct {
	Value      uint64 `protobuf:"varint,1,opt,name=Value" json:"Value,omitempty"`
	PubKeyHash []byte `protobuf:"bytes,2,opt,name=PubKeyHash,proto3" json:"PubKeyHash,omitempty"`
}

func (m *TXOutput) Reset()                    { *m = TXOutput{} }
func (m *TXOutput) String() string            { return proto.CompactTextString(m) }
func (*TXOutput) ProtoMessage()               {}
func (*TXOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TXOutput) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TXOutput) GetPubKeyHash() []byte {
	if m != nil {
		return m.PubKeyHash
	}
	return nil
}

type TXInput struct {
	Txid      []byte `protobuf:"bytes,1,opt,name=Txid,proto3" json:"Txid,omitempty"`
	Vout      uint64 `protobuf:"varint,2,opt,name=Vout" json:"Vout,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=Signature,proto3" json:"Signature,omitempty"`
	PubKey    []byte `protobuf:"bytes,4,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
}

func (m *TXInput) Reset()                    { *m = TXInput{} }
func (m *TXInput) String() string            { return proto.CompactTextString(m) }
func (*TXInput) ProtoMessage()               {}
func (*TXInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TXInput) GetTxid() []byte {
	if m != nil {
		return m.Txid
	}
	return nil
}

func (m *TXInput) GetVout() uint64 {
	if m != nil {
		return m.Vout
	}
	return 0
}

func (m *TXInput) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *TXInput) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type Transaction struct {
	ID   []byte      `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Vin  []*TXInput  `protobuf:"bytes,2,rep,name=Vin" json:"Vin,omitempty"`
	Vout []*TXOutput `protobuf:"bytes,3,rep,name=Vout" json:"Vout,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Transaction) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Transaction) GetVin() []*TXInput {
	if m != nil {
		return m.Vin
	}
	return nil
}

func (m *Transaction) GetVout() []*TXOutput {
	if m != nil {
		return m.Vout
	}
	return nil
}

type Block struct {
	Version       uint32                     `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Timestamp     *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=Timestamp" json:"Timestamp,omitempty"`
	Transactions  []*Transaction             `protobuf:"bytes,3,rep,name=Transactions" json:"Transactions,omitempty"`
	PrevBlockHash []byte                     `protobuf:"bytes,4,opt,name=PrevBlockHash,proto3" json:"PrevBlockHash,omitempty"`
	Hash          []byte                     `protobuf:"bytes,5,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Nonce         uint64                     `protobuf:"varint,6,opt,name=Nonce" json:"Nonce,omitempty"`
	Height        uint64                     `protobuf:"varint,7,opt,name=Height" json:"Height,omitempty"`
	StateHash     []byte                     `protobuf:"bytes,8,opt,name=StateHash,proto3" json:"StateHash,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Block) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Block) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func init() {
	proto.RegisterType((*OkMessage)(nil), "protos.OkMessage")
	proto.RegisterType((*PeerAddress)(nil), "protos.PeerAddress")
	proto.RegisterType((*PeerID)(nil), "protos.PeerID")
	proto.RegisterType((*PeerEndpoint)(nil), "protos.PeerEndpoint")
	proto.RegisterType((*PeersMessage)(nil), "protos.PeersMessage")
	proto.RegisterType((*PeersAddresses)(nil), "protos.PeersAddresses")
	proto.RegisterType((*HelloMessage)(nil), "protos.HelloMessage")
	proto.RegisterType((*TXOutput)(nil), "protos.TXOutput")
	proto.RegisterType((*TXInput)(nil), "protos.TXInput")
	proto.RegisterType((*Transaction)(nil), "protos.Transaction")
	proto.RegisterType((*Block)(nil), "protos.Block")
	proto.RegisterEnum("protos.OkMessage_Type", OkMessage_Type_name, OkMessage_Type_value)
	proto.RegisterEnum("protos.PeerEndpoint_Type", PeerEndpoint_Type_name, PeerEndpoint_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Peer service

type PeerClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (Peer_ChatClient, error)
}

type peerClient struct {
	cc *grpc.ClientConn
}

func NewPeerClient(cc *grpc.ClientConn) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Peer_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Peer_serviceDesc.Streams[0], c.cc, "/protos.Peer/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerChatClient{stream}
	return x, nil
}

type Peer_ChatClient interface {
	Send(*OkMessage) error
	Recv() (*OkMessage, error)
	grpc.ClientStream
}

type peerChatClient struct {
	grpc.ClientStream
}

func (x *peerChatClient) Send(m *OkMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peerChatClient) Recv() (*OkMessage, error) {
	m := new(OkMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Peer service

type PeerServer interface {
	Chat(Peer_ChatServer) error
}

func RegisterPeerServer(s *grpc.Server, srv PeerServer) {
	s.RegisterService(&_Peer_serviceDesc, srv)
}

func _Peer_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeerServer).Chat(&peerChatServer{stream})
}

type Peer_ChatServer interface {
	Send(*OkMessage) error
	Recv() (*OkMessage, error)
	grpc.ServerStream
}

type peerChatServer struct {
	grpc.ServerStream
}

func (x *peerChatServer) Send(m *OkMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peerChatServer) Recv() (*OkMessage, error) {
	m := new(OkMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Peer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _Peer_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 831 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x0e, 0x25, 0x4a, 0xb6, 0x46, 0x94, 0x4c, 0xaf, 0xe5, 0x84, 0x71, 0x83, 0xd6, 0x25, 0x72,
	0x10, 0x0a, 0x54, 0x09, 0x54, 0x14, 0x29, 0x72, 0x0a, 0x23, 0xb2, 0x11, 0x11, 0x85, 0x14, 0x96,
	0x1b, 0x37, 0x3d, 0x09, 0xb4, 0xb5, 0x91, 0x84, 0xc8, 0x24, 0x21, 0xae, 0x82, 0xfa, 0x55, 0x7a,
	0xea, 0xdb, 0xf4, 0x81, 0xfa, 0x02, 0xc5, 0xce, 0xf2, 0x47, 0x69, 0xd2, 0x43, 0x2e, 0xf6, 0x7c,
	0xdf, 0x7c, 0xf3, 0xc3, 0x99, 0x21, 0x05, 0x70, 0xb3, 0x8e, 0xc5, 0x28, 0xdb, 0xa5, 0x22, 0x25,
	0x6d, 0xfc, 0x97, 0x5f, 0x7c, 0xb7, 0x4a, 0xd3, 0xd5, 0x96, 0x3f, 0x41, 0x78, 0xbd, 0x7f, 0xff,
	0x44, 0x6c, 0x6e, 0x79, 0x2e, 0xe2, 0xdb, 0x4c, 0x09, 0xed, 0x7f, 0x9a, 0xd0, 0x09, 0x3f, 0xbc,
	0xe1, 0x79, 0x1e, 0xaf, 0x38, 0xf9, 0x01, 0x74, 0x71, 0x97, 0x71, 0x4b, 0xbb, 0xd4, 0x86, 0xfd,
	0xf1, 0x7d, 0xa5, 0xc9, 0x47, 0x95, 0x60, 0xc4, 0xee, 0x32, 0x4e, 0x51, 0x43, 0x7e, 0x81, 0x4e,
	0x95, 0xcc, 0x6a, 0x5c, 0x6a, 0xc3, 0xee, 0xf8, 0x62, 0xa4, 0xca, 0x8d, 0xca, 0x72, 0x23, 0x56,
	0x2a, 0x68, 0x2d, 0x26, 0x16, 0x1c, 0x65, 0xf1, 0xdd, 0x36, 0x8d, 0x97, 0x56, 0xf3, 0x52, 0x1b,
	0x1a, 0xb4, 0x84, 0xe4, 0x11, 0x74, 0xf2, 0xcd, 0x2a, 0x89, 0xc5, 0x7e, 0xc7, 0x2d, 0x1d, 0x7d,
	0x35, 0x61, 0xff, 0xdd, 0x00, 0x5d, 0x36, 0x40, 0x7a, 0xd0, 0x79, 0x1b, 0xb8, 0xde, 0xaf, 0x7e,
	0xe0, 0xb9, 0xe6, 0x3d, 0xd2, 0x07, 0x70, 0xfd, 0x68, 0xb2, 0x98, 0x7a, 0xb3, 0x59, 0x68, 0x6a,
	0xe4, 0x0c, 0x4e, 0x10, 0xcb, 0x3f, 0x61, 0x10, 0x78, 0x13, 0x66, 0x36, 0x08, 0x81, 0x3e, 0x92,
	0xaf, 0x3c, 0xb6, 0x98, 0x7b, 0x1e, 0x8d, 0xcc, 0x66, 0x15, 0xa8, 0xb0, 0x4e, 0x4e, 0xa0, 0x8b,
	0x38, 0xf0, 0x7e, 0x7b, 0x13, 0xbd, 0x32, 0x5b, 0xe4, 0x1c, 0x4e, 0x27, 0x53, 0xc7, 0x0f, 0x16,
	0x8c, 0x3a, 0x41, 0xe4, 0x4c, 0x98, 0x1f, 0x06, 0x66, 0x5b, 0x16, 0x88, 0x7e, 0x0f, 0x54, 0xae,
	0x97, 0xb3, 0x70, 0xf2, 0x3a, 0x32, 0xbb, 0x32, 0x18, 0xc9, 0x82, 0x30, 0xc8, 0x00, 0xcc, 0x9a,
	0x58, 0x38, 0xae, 0xeb, 0xb9, 0x66, 0x8f, 0x7c, 0x03, 0x0f, 0x90, 0x8d, 0x98, 0xc3, 0x3c, 0xcc,
	0x10, 0x05, 0xce, 0x3c, 0x9a, 0x86, 0xcc, 0xec, 0x93, 0x07, 0x70, 0x76, 0xe0, 0xac, 0x1c, 0x27,
	0xe4, 0x21, 0x9c, 0xff, 0x27, 0xca, 0xf5, 0x66, 0xcc, 0x89, 0x4c, 0x53, 0xf6, 0x78, 0xe0, 0x2a,
	0xe8, 0x53, 0x62, 0xc0, 0x31, 0xf5, 0xa2, 0x79, 0x18, 0x44, 0x9e, 0x39, 0x90, 0x13, 0x9b, 0x48,
	0x33, 0x88, 0xde, 0x46, 0xe6, 0xb9, 0xfd, 0x33, 0x74, 0xe7, 0x9c, 0xef, 0x9c, 0xe5, 0x72, 0xc7,
	0xf3, 0x9c, 0x10, 0xd0, 0xd7, 0x69, 0x2e, 0x70, 0xed, 0x1d, 0x8a, 0xb6, 0xe4, 0xb2, 0x74, 0x27,
	0x70, 0xb3, 0x2d, 0x8a, 0xb6, 0xfd, 0x08, 0xda, 0x32, 0xcc, 0x77, 0xa5, 0x37, 0x89, 0x6f, 0x79,
	0x19, 0x21, 0x6d, 0xfb, 0x2f, 0x0d, 0x0c, 0xe9, 0xf6, 0x92, 0x65, 0x96, 0x6e, 0x12, 0x41, 0xfa,
	0xd0, 0xf0, 0xdd, 0x42, 0xd2, 0xf0, 0x5d, 0xb9, 0xf7, 0x58, 0x55, 0xc4, 0xac, 0x1d, 0x5a, 0x42,
	0xf2, 0x63, 0x71, 0x77, 0x4d, 0xbc, 0xbb, 0x87, 0xe5, 0xdd, 0x1d, 0x66, 0x3b, 0x38, 0x3d, 0xfb,
	0xd9, 0x97, 0xef, 0xa0, 0x07, 0x9d, 0x2b, 0x67, 0xe6, 0xbb, 0x0e, 0x0b, 0xa9, 0xa9, 0x91, 0x53,
	0xe8, 0x05, 0x61, 0xb0, 0xa8, 0xa9, 0x86, 0xfd, 0x5c, 0x75, 0x98, 0xd7, 0xf7, 0xde, 0xca, 0x24,
	0xb6, 0xb4, 0xcb, 0xe6, 0xb0, 0x3b, 0x1e, 0x7c, 0xa9, 0x30, 0x55, 0x12, 0x7b, 0x04, 0x7d, 0x8c,
	0x2d, 0x86, 0xc6, 0x73, 0x79, 0xad, 0x71, 0x09, 0x30, 0x43, 0x87, 0xd6, 0x84, 0xfd, 0x02, 0x8c,
	0x29, 0xdf, 0x6e, 0xd3, 0xb2, 0xd6, 0x53, 0x38, 0xe6, 0x45, 0x4a, 0x9c, 0xc9, 0xff, 0x95, 0xab,
	0x54, 0xf6, 0x0b, 0x38, 0x66, 0xef, 0xc2, 0xbd, 0xc8, 0xf6, 0x82, 0x0c, 0xa0, 0x75, 0x15, 0x6f,
	0xf7, 0x6a, 0xe2, 0x3a, 0x55, 0x80, 0x7c, 0x0b, 0x30, 0xdf, 0x5f, 0xbf, 0xe6, 0x77, 0xd3, 0x38,
	0x5f, 0xe3, 0x50, 0x0d, 0x7a, 0xc0, 0xd8, 0x2b, 0x38, 0x62, 0xef, 0xfc, 0x44, 0x26, 0x20, 0xa0,
	0xb3, 0x3f, 0x36, 0x4b, 0x8c, 0x37, 0x28, 0xda, 0x92, 0xbb, 0x4a, 0xf7, 0x6a, 0xc7, 0x3a, 0x45,
	0x5b, 0x3e, 0x54, 0x54, 0xbd, 0x82, 0xea, 0xf5, 0xac, 0x09, 0x72, 0x1f, 0xda, 0x2a, 0x7d, 0xf1,
	0x76, 0x16, 0xc8, 0x7e, 0x0f, 0x5d, 0xb6, 0x8b, 0x93, 0x3c, 0xbe, 0x11, 0x9b, 0x34, 0x39, 0xd8,
	0xbc, 0x81, 0x9b, 0xff, 0x1e, 0x9a, 0x57, 0x9b, 0xc4, 0x6a, 0xe0, 0x94, 0x4f, 0xca, 0xc7, 0x2e,
	0x5a, 0xa3, 0xd2, 0x47, 0x1e, 0x17, 0xbd, 0x34, 0x51, 0x63, 0xd6, 0x1a, 0x35, 0x00, 0xd5, 0x9d,
	0xfd, 0x67, 0x03, 0x5a, 0x2f, 0xb7, 0xe9, 0xcd, 0x07, 0x79, 0x4c, 0x1f, 0xf9, 0x2e, 0xdf, 0xa4,
	0x09, 0xd6, 0xe9, 0xd1, 0x12, 0xca, 0x0f, 0x13, 0xfb, 0x9a, 0x0f, 0x53, 0x65, 0x92, 0x67, 0x60,
	0x1c, 0x3c, 0x45, 0x5e, 0xf4, 0x72, 0x56, 0xf5, 0x52, 0xfb, 0xe8, 0x27, 0x42, 0xf2, 0x18, 0x7a,
	0xf3, 0x1d, 0xff, 0x88, 0x9d, 0xe1, 0x2a, 0xd4, 0x74, 0x3e, 0x25, 0xe5, 0xb8, 0xd1, 0xd9, 0x52,
	0x2b, 0x40, 0x6e, 0x00, 0xad, 0x20, 0x4d, 0x6e, 0xb8, 0xd5, 0x56, 0x7b, 0x45, 0x20, 0xc7, 0x3c,
	0xe5, 0x9b, 0xd5, 0x5a, 0x58, 0x47, 0x48, 0x17, 0x08, 0x97, 0x23, 0x62, 0xc1, 0x31, 0xcd, 0x71,
	0xb1, 0x9c, 0x92, 0x18, 0x3f, 0x07, 0x5d, 0x5e, 0x12, 0x19, 0x83, 0x3e, 0x59, 0xc7, 0x82, 0x9c,
	0x7e, 0xf6, 0xfd, 0xbe, 0xf8, 0x9c, 0xb2, 0xef, 0x0d, 0xb5, 0xa7, 0xda, 0xb5, 0xfa, 0xc1, 0xf8,
	0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xf0, 0x52, 0x54, 0x45, 0x06, 0x00, 0x00,
}
