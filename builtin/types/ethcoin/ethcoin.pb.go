// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/ethcoin/ethcoin.proto

package ethcoin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ETHCoinInitRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ETHCoinInitRequest) Reset()         { *m = ETHCoinInitRequest{} }
func (m *ETHCoinInitRequest) String() string { return proto.CompactTextString(m) }
func (*ETHCoinInitRequest) ProtoMessage()    {}
func (*ETHCoinInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethcoin_15cf8379a9e6a48d, []int{0}
}
func (m *ETHCoinInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETHCoinInitRequest.Unmarshal(m, b)
}
func (m *ETHCoinInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETHCoinInitRequest.Marshal(b, m, deterministic)
}
func (dst *ETHCoinInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETHCoinInitRequest.Merge(dst, src)
}
func (m *ETHCoinInitRequest) XXX_Size() int {
	return xxx_messageInfo_ETHCoinInitRequest.Size(m)
}
func (m *ETHCoinInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ETHCoinInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ETHCoinInitRequest proto.InternalMessageInfo

type ETHCoinMintToGatewayRequest struct {
	Amount               *types.BigUInt `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ETHCoinMintToGatewayRequest) Reset()         { *m = ETHCoinMintToGatewayRequest{} }
func (m *ETHCoinMintToGatewayRequest) String() string { return proto.CompactTextString(m) }
func (*ETHCoinMintToGatewayRequest) ProtoMessage()    {}
func (*ETHCoinMintToGatewayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethcoin_15cf8379a9e6a48d, []int{1}
}
func (m *ETHCoinMintToGatewayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETHCoinMintToGatewayRequest.Unmarshal(m, b)
}
func (m *ETHCoinMintToGatewayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETHCoinMintToGatewayRequest.Marshal(b, m, deterministic)
}
func (dst *ETHCoinMintToGatewayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETHCoinMintToGatewayRequest.Merge(dst, src)
}
func (m *ETHCoinMintToGatewayRequest) XXX_Size() int {
	return xxx_messageInfo_ETHCoinMintToGatewayRequest.Size(m)
}
func (m *ETHCoinMintToGatewayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ETHCoinMintToGatewayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ETHCoinMintToGatewayRequest proto.InternalMessageInfo

func (m *ETHCoinMintToGatewayRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*ETHCoinInitRequest)(nil), "ETHCoinInitRequest")
	proto.RegisterType((*ETHCoinMintToGatewayRequest)(nil), "ETHCoinMintToGatewayRequest")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/ethcoin/ethcoin.proto", fileDescriptor_ethcoin_15cf8379a9e6a48d)
}

var fileDescriptor_ethcoin_15cf8379a9e6a48d = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xcf, 0xcf, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x4f, 0xcf, 0xd7, 0x05, 0x71, 0xf5, 0x93, 0x4a, 0x33, 0x73, 0x4a, 0x32, 0xf3,
	0xf4, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x53, 0x4b, 0x32, 0x92, 0xf3, 0x33, 0xf3, 0x60, 0xb4,
	0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x94, 0x01, 0x01, 0x33, 0x20, 0x7a, 0xc1, 0x24, 0x44, 0x87,
	0x92, 0x08, 0x97, 0x90, 0x6b, 0x88, 0x87, 0x73, 0x7e, 0x66, 0x9e, 0x67, 0x5e, 0x66, 0x49, 0x50,
	0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x92, 0x3d, 0x97, 0x34, 0x54, 0xd4, 0x37, 0x33, 0xaf, 0x24,
	0x24, 0xdf, 0x3d, 0xb1, 0x24, 0xb5, 0x3c, 0xb1, 0x12, 0x2a, 0x2d, 0xa4, 0xc0, 0xc5, 0x96, 0x98,
	0x9b, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xa1, 0xe7, 0x94, 0x99,
	0x1e, 0xea, 0x99, 0x57, 0x12, 0x04, 0x15, 0x4f, 0x62, 0x03, 0x9b, 0x6e, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x56, 0xda, 0xd7, 0x3c, 0xd5, 0x00, 0x00, 0x00,
}
