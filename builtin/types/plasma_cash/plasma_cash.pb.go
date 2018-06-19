// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

/*
Package plasma_cash is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

It has these top-level messages:
	PlasmaBlock
	PlasmaTx
	PlasmaBookKeeping
	GetCurrentBlockRequest
	GetCurrentBlockResponse
	GetBlockRequest
	GetBlockResponse
	SubmitBlockToMainnetRequest
	SubmitBlockToMainnetResponse
	PlasmaTxRequest
	PlasmaTxResponse
	DepositRequest
	DepositResponse
	GetProofRequest
	GetProofResponse
	PlasmaCashParams
	PlasmaCashInitRequest
*/
package plasma_cash

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
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

type PlasmaBlock struct {
	Uid          *types.BigUInt `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Transactions []*PlasmaTx    `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Signature    []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	MerkleHash   []byte         `protobuf:"bytes,4,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Hash         []byte         `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *PlasmaBlock) Reset()                    { *m = PlasmaBlock{} }
func (m *PlasmaBlock) String() string            { return proto.CompactTextString(m) }
func (*PlasmaBlock) ProtoMessage()               {}
func (*PlasmaBlock) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{0} }

func (m *PlasmaBlock) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *PlasmaBlock) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *PlasmaBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaBlock) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type PlasmaTx struct {
	Uid           uint64         `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PreviousBlock *types.BigUInt `protobuf:"bytes,2,opt,name=previous_block,json=previousBlock" json:"previous_block,omitempty"`
	Denomination  *types.BigUInt `protobuf:"bytes,3,opt,name=denomination" json:"denomination,omitempty"`
	NewOwner      *types.Address `protobuf:"bytes,4,opt,name=new_owner,json=newOwner" json:"new_owner,omitempty"`
	Signature     []byte         `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Hash          []byte         `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	MerkleHash    []byte         `protobuf:"bytes,7,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Sender        *types.Address `protobuf:"bytes,8,opt,name=sender" json:"sender,omitempty"`
}

func (m *PlasmaTx) Reset()                    { *m = PlasmaTx{} }
func (m *PlasmaTx) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTx) ProtoMessage()               {}
func (*PlasmaTx) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{1} }

func (m *PlasmaTx) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PlasmaTx) GetPreviousBlock() *types.BigUInt {
	if m != nil {
		return m.PreviousBlock
	}
	return nil
}

func (m *PlasmaTx) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *PlasmaTx) GetNewOwner() *types.Address {
	if m != nil {
		return m.NewOwner
	}
	return nil
}

func (m *PlasmaTx) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaTx) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PlasmaTx) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaTx) GetSender() *types.Address {
	if m != nil {
		return m.Sender
	}
	return nil
}

type PlasmaBookKeeping struct {
	CurrentHeight *types.BigUInt `protobuf:"bytes,1,opt,name=current_height,json=currentHeight" json:"current_height,omitempty"`
}

func (m *PlasmaBookKeeping) Reset()                    { *m = PlasmaBookKeeping{} }
func (m *PlasmaBookKeeping) String() string            { return proto.CompactTextString(m) }
func (*PlasmaBookKeeping) ProtoMessage()               {}
func (*PlasmaBookKeeping) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{2} }

func (m *PlasmaBookKeeping) GetCurrentHeight() *types.BigUInt {
	if m != nil {
		return m.CurrentHeight
	}
	return nil
}

type GetCurrentBlockRequest struct {
}

func (m *GetCurrentBlockRequest) Reset()                    { *m = GetCurrentBlockRequest{} }
func (m *GetCurrentBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCurrentBlockRequest) ProtoMessage()               {}
func (*GetCurrentBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{3} }

type GetCurrentBlockResponse struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *GetCurrentBlockResponse) Reset()         { *m = GetCurrentBlockResponse{} }
func (m *GetCurrentBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentBlockResponse) ProtoMessage()    {}
func (*GetCurrentBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{4}
}

func (m *GetCurrentBlockResponse) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type GetBlockRequest struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *GetBlockRequest) Reset()                    { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()               {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{5} }

func (m *GetBlockRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type GetBlockResponse struct {
	Block *PlasmaBlock `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *GetBlockResponse) Reset()                    { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()               {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{6} }

func (m *GetBlockResponse) GetBlock() *PlasmaBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

// This only originates from the validator
type SubmitBlockToMainnetRequest struct {
}

func (m *SubmitBlockToMainnetRequest) Reset()         { *m = SubmitBlockToMainnetRequest{} }
func (m *SubmitBlockToMainnetRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetRequest) ProtoMessage()    {}
func (*SubmitBlockToMainnetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{7}
}

type SubmitBlockToMainnetResponse struct {
}

func (m *SubmitBlockToMainnetResponse) Reset()         { *m = SubmitBlockToMainnetResponse{} }
func (m *SubmitBlockToMainnetResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetResponse) ProtoMessage()    {}
func (*SubmitBlockToMainnetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{8}
}

type PlasmaTxRequest struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *PlasmaTxRequest) Reset()                    { *m = PlasmaTxRequest{} }
func (m *PlasmaTxRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxRequest) ProtoMessage()               {}
func (*PlasmaTxRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{9} }

func (m *PlasmaTxRequest) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type PlasmaTxResponse struct {
}

func (m *PlasmaTxResponse) Reset()                    { *m = PlasmaTxResponse{} }
func (m *PlasmaTxResponse) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxResponse) ProtoMessage()               {}
func (*PlasmaTxResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{10} }

// This only originates from the validator
type DepositRequest struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *DepositRequest) Reset()                    { *m = DepositRequest{} }
func (m *DepositRequest) String() string            { return proto.CompactTextString(m) }
func (*DepositRequest) ProtoMessage()               {}
func (*DepositRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{11} }

func (m *DepositRequest) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type DepositResponse struct {
}

func (m *DepositResponse) Reset()                    { *m = DepositResponse{} }
func (m *DepositResponse) String() string            { return proto.CompactTextString(m) }
func (*DepositResponse) ProtoMessage()               {}
func (*DepositResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{12} }

type GetProofRequest struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
	Uid         *types.BigUInt `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *GetProofRequest) Reset()                    { *m = GetProofRequest{} }
func (m *GetProofRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProofRequest) ProtoMessage()               {}
func (*GetProofRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{13} }

func (m *GetProofRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

func (m *GetProofRequest) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

type GetProofResponse struct {
	Proof []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *GetProofResponse) Reset()                    { *m = GetProofResponse{} }
func (m *GetProofResponse) String() string            { return proto.CompactTextString(m) }
func (*GetProofResponse) ProtoMessage()               {}
func (*GetProofResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{14} }

func (m *GetProofResponse) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

// Initialization of state from Genesis.json
type PlasmaCashParams struct {
	BlockInterval uint64 `protobuf:"varint,1,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
}

func (m *PlasmaCashParams) Reset()                    { *m = PlasmaCashParams{} }
func (m *PlasmaCashParams) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashParams) ProtoMessage()               {}
func (*PlasmaCashParams) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{15} }

func (m *PlasmaCashParams) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

type PlasmaCashInitRequest struct {
	Params *PlasmaCashParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
}

func (m *PlasmaCashInitRequest) Reset()                    { *m = PlasmaCashInitRequest{} }
func (m *PlasmaCashInitRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashInitRequest) ProtoMessage()               {}
func (*PlasmaCashInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{16} }

func (m *PlasmaCashInitRequest) GetParams() *PlasmaCashParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*PlasmaBlock)(nil), "PlasmaBlock")
	proto.RegisterType((*PlasmaTx)(nil), "PlasmaTx")
	proto.RegisterType((*PlasmaBookKeeping)(nil), "PlasmaBookKeeping")
	proto.RegisterType((*GetCurrentBlockRequest)(nil), "GetCurrentBlockRequest")
	proto.RegisterType((*GetCurrentBlockResponse)(nil), "GetCurrentBlockResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "GetBlockRequest")
	proto.RegisterType((*GetBlockResponse)(nil), "GetBlockResponse")
	proto.RegisterType((*SubmitBlockToMainnetRequest)(nil), "SubmitBlockToMainnetRequest")
	proto.RegisterType((*SubmitBlockToMainnetResponse)(nil), "SubmitBlockToMainnetResponse")
	proto.RegisterType((*PlasmaTxRequest)(nil), "PlasmaTxRequest")
	proto.RegisterType((*PlasmaTxResponse)(nil), "PlasmaTxResponse")
	proto.RegisterType((*DepositRequest)(nil), "DepositRequest")
	proto.RegisterType((*DepositResponse)(nil), "DepositResponse")
	proto.RegisterType((*GetProofRequest)(nil), "GetProofRequest")
	proto.RegisterType((*GetProofResponse)(nil), "GetProofResponse")
	proto.RegisterType((*PlasmaCashParams)(nil), "PlasmaCashParams")
	proto.RegisterType((*PlasmaCashInitRequest)(nil), "PlasmaCashInitRequest")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto", fileDescriptorPlasmaCash)
}

var fileDescriptorPlasmaCash = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x55, 0xbb, 0xb5, 0x74, 0x5f, 0xd3, 0x75, 0x8d, 0xf8, 0x89, 0xc6, 0xa0, 0x55, 0xa4, 0x49,
	0x45, 0xb0, 0x16, 0x6d, 0x12, 0x3f, 0x12, 0x37, 0x74, 0x48, 0xdb, 0x40, 0xb0, 0x2a, 0x0c, 0x4d,
	0xe2, 0xa6, 0x72, 0x5a, 0x2f, 0xb1, 0xda, 0xd8, 0xc1, 0x76, 0xd6, 0xf1, 0x4c, 0xbc, 0x00, 0x4f,
	0xb3, 0x8b, 0x3d, 0x09, 0x8a, 0xed, 0x2e, 0x49, 0x37, 0x40, 0x88, 0x9b, 0xca, 0xdf, 0xdf, 0x39,
	0xdf, 0x39, 0xa9, 0x0d, 0xef, 0x03, 0x22, 0xc3, 0xc4, 0xef, 0x8d, 0x59, 0xd4, 0x9f, 0x31, 0x16,
	0x51, 0x2c, 0xe7, 0x8c, 0x4f, 0xfb, 0x01, 0xdb, 0x49, 0xc3, 0xbe, 0x9f, 0x90, 0x99, 0x24, 0xb4,
	0x2f, 0xbf, 0xc7, 0x58, 0xf4, 0xe3, 0x19, 0x12, 0x11, 0x1a, 0x8d, 0x91, 0x08, 0xf3, 0xe7, 0x5e,
	0xcc, 0x99, 0x64, 0x9b, 0x3b, 0x39, 0xac, 0x80, 0x05, 0xac, 0xaf, 0xd2, 0x7e, 0x72, 0xa6, 0x22,
	0x15, 0xa8, 0x93, 0x69, 0x7f, 0xfe, 0x17, 0x6a, 0x4d, 0xa9, 0x7e, 0xf5, 0x84, 0xfb, 0xa3, 0x04,
	0xf5, 0xa1, 0xa2, 0x1d, 0xcc, 0xd8, 0x78, 0x6a, 0x6f, 0xc2, 0x4a, 0x42, 0x26, 0x4e, 0xa9, 0x53,
	0xea, 0xd6, 0x77, 0x6b, 0xbd, 0x01, 0x09, 0xbe, 0x1c, 0x51, 0xe9, 0xa5, 0x49, 0x7b, 0x07, 0x2c,
	0xc9, 0x11, 0x15, 0x68, 0x2c, 0x09, 0xa3, 0xc2, 0x29, 0x77, 0x56, 0xba, 0xf5, 0xdd, 0xb5, 0x9e,
	0x9e, 0x3f, 0xb9, 0xf0, 0x0a, 0x65, 0x7b, 0x0b, 0xd6, 0x04, 0x09, 0x28, 0x92, 0x09, 0xc7, 0xce,
	0x4a, 0xa7, 0xd4, 0xb5, 0xbc, 0x2c, 0x61, 0xb7, 0xa1, 0x1e, 0x61, 0x3e, 0x9d, 0xe1, 0x51, 0x88,
	0x44, 0xe8, 0xac, 0xaa, 0x3a, 0xe8, 0xd4, 0x21, 0x12, 0xa1, 0x6d, 0xc3, 0xaa, 0xaa, 0x54, 0x54,
	0x45, 0x9d, 0xdd, 0x9f, 0x65, 0xa8, 0x2d, 0xd8, 0xec, 0x8d, 0x6c, 0xd5, 0x55, 0xbd, 0xe0, 0x00,
	0xd6, 0x63, 0x8e, 0xcf, 0x09, 0x4b, 0xc4, 0xc8, 0x4f, 0xe5, 0x38, 0xe5, 0xa2, 0x8e, 0x41, 0xeb,
	0xea, 0xb2, 0xdd, 0x18, 0x9a, 0x1e, 0xa5, 0xd8, 0x6b, 0xc4, 0xf9, 0xd0, 0x7e, 0x06, 0xd6, 0x04,
	0x53, 0x16, 0x11, 0x8a, 0x52, 0x19, 0x6a, 0xf1, 0xbc, 0x13, 0x85, 0xaa, 0xbd, 0x07, 0x6b, 0x14,
	0xcf, 0x47, 0x6c, 0x4e, 0x31, 0x57, 0x1a, 0xd2, 0xd6, 0xb7, 0x93, 0x09, 0xc7, 0x42, 0x0c, 0xac,
	0xab, 0xcb, 0x76, 0xed, 0x13, 0x9e, 0x1f, 0xa7, 0x55, 0xaf, 0x46, 0xcd, 0xa9, 0x68, 0x4c, 0x65,
	0xd9, 0x98, 0x85, 0xee, 0x6a, 0xa6, 0x7b, 0xd9, 0xac, 0x3b, 0x37, 0xcc, 0xea, 0x40, 0x55, 0x60,
	0x3a, 0xc1, 0xdc, 0xa9, 0x15, 0x97, 0xf0, 0x4c, 0xde, 0x3d, 0x85, 0x96, 0xf9, 0xce, 0x8c, 0x4d,
	0x3f, 0x60, 0x1c, 0x13, 0x1a, 0xa4, 0x86, 0x8d, 0x13, 0xce, 0x31, 0x95, 0xa3, 0x10, 0x93, 0x20,
	0x94, 0xcb, 0x1f, 0x5e, 0x1b, 0xb6, 0xaf, 0x7b, 0x0e, 0x55, 0x8b, 0xd7, 0x18, 0xe7, 0x43, 0xd7,
	0x81, 0xfb, 0x07, 0x58, 0x9a, 0x16, 0x6d, 0x29, 0xfe, 0x96, 0x60, 0x21, 0xdd, 0x53, 0x78, 0x70,
	0xa3, 0x22, 0x62, 0x46, 0x05, 0xb6, 0xdf, 0x80, 0xa5, 0x3e, 0xd0, 0xef, 0x68, 0x9b, 0x57, 0x97,
	0xed, 0xba, 0x1a, 0x31, 0xa4, 0x75, 0x3f, 0x0b, 0xdc, 0x63, 0x68, 0x1e, 0xe0, 0x02, 0xd7, 0x7f,
	0x02, 0xbe, 0x80, 0x8d, 0x0c, 0xd0, 0xac, 0xe8, 0x42, 0x45, 0xff, 0x87, 0x34, 0x94, 0xd5, 0xcb,
	0x5d, 0x13, 0x4f, 0x97, 0xdc, 0x47, 0xf0, 0xf0, 0x73, 0xe2, 0x47, 0x44, 0x8f, 0x9e, 0xb0, 0x8f,
	0x88, 0x50, 0x8a, 0xe5, 0xc2, 0x80, 0xc7, 0xb0, 0x75, 0x7b, 0x59, 0x53, 0xb8, 0xaf, 0xa0, 0x79,
	0x7d, 0x77, 0x8c, 0x8e, 0x6d, 0xa8, 0xe9, 0x57, 0x40, 0x5e, 0x18, 0xe2, 0xdc, 0xfd, 0xba, 0x2e,
	0xb9, 0x36, 0x6c, 0x64, 0x93, 0x06, 0xed, 0x25, 0xac, 0xbf, 0xc3, 0x31, 0x13, 0x44, 0xfe, 0x23,
	0x58, 0x0b, 0x9a, 0xd7, 0x83, 0x06, 0xeb, 0xab, 0x72, 0x78, 0xc8, 0x19, 0x3b, 0x5b, 0x80, 0x3d,
	0xfd, 0xb3, 0xc3, 0x05, 0x43, 0x17, 0xcf, 0x48, 0xf9, 0x96, 0x67, 0xc4, 0xed, 0x2a, 0xb3, 0x0d,
	0xb6, 0x31, 0xfb, 0x2e, 0x54, 0xe2, 0x34, 0xa1, 0x50, 0x2d, 0x4f, 0x07, 0xee, 0xeb, 0x85, 0xca,
	0x7d, 0x24, 0xc2, 0x21, 0xe2, 0x28, 0x12, 0xf6, 0x36, 0xac, 0xeb, 0x35, 0x08, 0x95, 0x98, 0x9f,
	0xa3, 0x99, 0x79, 0x00, 0x1a, 0x2a, 0x7b, 0x64, 0x92, 0xee, 0x00, 0xee, 0x65, 0xa3, 0x47, 0x34,
	0xf3, 0xe4, 0x09, 0x54, 0x63, 0x85, 0x64, 0x04, 0xb4, 0x7a, 0xcb, 0x14, 0x9e, 0x69, 0xf0, 0xab,
	0xea, 0x89, 0xdc, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xca, 0xdb, 0x6f, 0xd1, 0x05, 0x00,
	0x00,
}
