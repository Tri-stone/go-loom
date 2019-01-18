// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto

package karma

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/loomnetwork/go-loom/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type KarmaSourceTarget int32

const (
	KarmaSourceTarget_CALL   KarmaSourceTarget = 0
	KarmaSourceTarget_DEPLOY KarmaSourceTarget = 1
	KarmaSourceTarget_ALL    KarmaSourceTarget = 2
)

var KarmaSourceTarget_name = map[int32]string{
	0: "CALL",
	1: "DEPLOY",
	2: "ALL",
}

var KarmaSourceTarget_value = map[string]int32{
	"CALL":   0,
	"DEPLOY": 1,
	"ALL":    2,
}

func (x KarmaSourceTarget) String() string {
	return proto.EnumName(KarmaSourceTarget_name, int32(x))
}

func (KarmaSourceTarget) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{0}
}

type KarmaInitRequest struct {
	Oracle               *types.Address        `protobuf:"bytes,1,opt,name=Oracle,proto3" json:"Oracle,omitempty"`
	Sources              []*KarmaSourceReward  `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty"`
	Users                []*KarmaAddressSource `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Upkeep               *KarmaUpkeepParams    `protobuf:"bytes,4,opt,name=upkeep,proto3" json:"upkeep,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KarmaInitRequest) Reset()         { *m = KarmaInitRequest{} }
func (m *KarmaInitRequest) String() string { return proto.CompactTextString(m) }
func (*KarmaInitRequest) ProtoMessage()    {}
func (*KarmaInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{0}
}
func (m *KarmaInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaInitRequest.Unmarshal(m, b)
}
func (m *KarmaInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaInitRequest.Marshal(b, m, deterministic)
}
func (m *KarmaInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaInitRequest.Merge(m, src)
}
func (m *KarmaInitRequest) XXX_Size() int {
	return xxx_messageInfo_KarmaInitRequest.Size(m)
}
func (m *KarmaInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaInitRequest proto.InternalMessageInfo

func (m *KarmaInitRequest) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func (m *KarmaInitRequest) GetSources() []*KarmaSourceReward {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *KarmaInitRequest) GetUsers() []*KarmaAddressSource {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *KarmaInitRequest) GetUpkeep() *KarmaUpkeepParams {
	if m != nil {
		return m.Upkeep
	}
	return nil
}

type KarmaSources struct {
	Sources              []*KarmaSourceReward `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KarmaSources) Reset()         { *m = KarmaSources{} }
func (m *KarmaSources) String() string { return proto.CompactTextString(m) }
func (*KarmaSources) ProtoMessage()    {}
func (*KarmaSources) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{1}
}
func (m *KarmaSources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaSources.Unmarshal(m, b)
}
func (m *KarmaSources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaSources.Marshal(b, m, deterministic)
}
func (m *KarmaSources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaSources.Merge(m, src)
}
func (m *KarmaSources) XXX_Size() int {
	return xxx_messageInfo_KarmaSources.Size(m)
}
func (m *KarmaSources) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaSources.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaSources proto.InternalMessageInfo

func (m *KarmaSources) GetSources() []*KarmaSourceReward {
	if m != nil {
		return m.Sources
	}
	return nil
}

type KarmaNewOracle struct {
	NewOracle            *types.Address `protobuf:"bytes,1,opt,name=new_oracle,json=newOracle,proto3" json:"new_oracle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaNewOracle) Reset()         { *m = KarmaNewOracle{} }
func (m *KarmaNewOracle) String() string { return proto.CompactTextString(m) }
func (*KarmaNewOracle) ProtoMessage()    {}
func (*KarmaNewOracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{2}
}
func (m *KarmaNewOracle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaNewOracle.Unmarshal(m, b)
}
func (m *KarmaNewOracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaNewOracle.Marshal(b, m, deterministic)
}
func (m *KarmaNewOracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaNewOracle.Merge(m, src)
}
func (m *KarmaNewOracle) XXX_Size() int {
	return xxx_messageInfo_KarmaNewOracle.Size(m)
}
func (m *KarmaNewOracle) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaNewOracle.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaNewOracle proto.InternalMessageInfo

func (m *KarmaNewOracle) GetNewOracle() *types.Address {
	if m != nil {
		return m.NewOracle
	}
	return nil
}

type KarmaUserTarget struct {
	User                 *types.Address    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Target               KarmaSourceTarget `protobuf:"varint,2,opt,name=target,proto3,enum=karma.KarmaSourceTarget" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *KarmaUserTarget) Reset()         { *m = KarmaUserTarget{} }
func (m *KarmaUserTarget) String() string { return proto.CompactTextString(m) }
func (*KarmaUserTarget) ProtoMessage()    {}
func (*KarmaUserTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{3}
}
func (m *KarmaUserTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaUserTarget.Unmarshal(m, b)
}
func (m *KarmaUserTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaUserTarget.Marshal(b, m, deterministic)
}
func (m *KarmaUserTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaUserTarget.Merge(m, src)
}
func (m *KarmaUserTarget) XXX_Size() int {
	return xxx_messageInfo_KarmaUserTarget.Size(m)
}
func (m *KarmaUserTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaUserTarget.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaUserTarget proto.InternalMessageInfo

func (m *KarmaUserTarget) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaUserTarget) GetTarget() KarmaSourceTarget {
	if m != nil {
		return m.Target
	}
	return KarmaSourceTarget_CALL
}

type KarmaUserAmount struct {
	User                 *types.Address `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Amount               *types.BigUInt `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaUserAmount) Reset()         { *m = KarmaUserAmount{} }
func (m *KarmaUserAmount) String() string { return proto.CompactTextString(m) }
func (*KarmaUserAmount) ProtoMessage()    {}
func (*KarmaUserAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{4}
}
func (m *KarmaUserAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaUserAmount.Unmarshal(m, b)
}
func (m *KarmaUserAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaUserAmount.Marshal(b, m, deterministic)
}
func (m *KarmaUserAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaUserAmount.Merge(m, src)
}
func (m *KarmaUserAmount) XXX_Size() int {
	return xxx_messageInfo_KarmaUserAmount.Size(m)
}
func (m *KarmaUserAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaUserAmount.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaUserAmount proto.InternalMessageInfo

func (m *KarmaUserAmount) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaUserAmount) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type KarmaSourceReward struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reward               int64             `protobuf:"varint,2,opt,name=reward,proto3" json:"reward,omitempty"`
	Target               KarmaSourceTarget `protobuf:"varint,3,opt,name=target,proto3,enum=karma.KarmaSourceTarget" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *KarmaSourceReward) Reset()         { *m = KarmaSourceReward{} }
func (m *KarmaSourceReward) String() string { return proto.CompactTextString(m) }
func (*KarmaSourceReward) ProtoMessage()    {}
func (*KarmaSourceReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{5}
}
func (m *KarmaSourceReward) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaSourceReward.Unmarshal(m, b)
}
func (m *KarmaSourceReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaSourceReward.Marshal(b, m, deterministic)
}
func (m *KarmaSourceReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaSourceReward.Merge(m, src)
}
func (m *KarmaSourceReward) XXX_Size() int {
	return xxx_messageInfo_KarmaSourceReward.Size(m)
}
func (m *KarmaSourceReward) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaSourceReward.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaSourceReward proto.InternalMessageInfo

func (m *KarmaSourceReward) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KarmaSourceReward) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

func (m *KarmaSourceReward) GetTarget() KarmaSourceTarget {
	if m != nil {
		return m.Target
	}
	return KarmaSourceTarget_CALL
}

type KarmaSource struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count                *types.BigUInt `protobuf:"bytes,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaSource) Reset()         { *m = KarmaSource{} }
func (m *KarmaSource) String() string { return proto.CompactTextString(m) }
func (*KarmaSource) ProtoMessage()    {}
func (*KarmaSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{6}
}
func (m *KarmaSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaSource.Unmarshal(m, b)
}
func (m *KarmaSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaSource.Marshal(b, m, deterministic)
}
func (m *KarmaSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaSource.Merge(m, src)
}
func (m *KarmaSource) XXX_Size() int {
	return xxx_messageInfo_KarmaSource.Size(m)
}
func (m *KarmaSource) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaSource.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaSource proto.InternalMessageInfo

func (m *KarmaSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KarmaSource) GetCount() *types.BigUInt {
	if m != nil {
		return m.Count
	}
	return nil
}

type KarmaUpkeepParams struct {
	Cost                 int64    `protobuf:"varint,1,opt,name=cost,proto3" json:"cost,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Period               int64    `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KarmaUpkeepParams) Reset()         { *m = KarmaUpkeepParams{} }
func (m *KarmaUpkeepParams) String() string { return proto.CompactTextString(m) }
func (*KarmaUpkeepParams) ProtoMessage()    {}
func (*KarmaUpkeepParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{7}
}
func (m *KarmaUpkeepParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaUpkeepParams.Unmarshal(m, b)
}
func (m *KarmaUpkeepParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaUpkeepParams.Marshal(b, m, deterministic)
}
func (m *KarmaUpkeepParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaUpkeepParams.Merge(m, src)
}
func (m *KarmaUpkeepParams) XXX_Size() int {
	return xxx_messageInfo_KarmaUpkeepParams.Size(m)
}
func (m *KarmaUpkeepParams) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaUpkeepParams.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaUpkeepParams proto.InternalMessageInfo

func (m *KarmaUpkeepParams) GetCost() int64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *KarmaUpkeepParams) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *KarmaUpkeepParams) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

type KarmaAddressSource struct {
	User                 *types.Address `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Sources              []*KarmaSource `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaAddressSource) Reset()         { *m = KarmaAddressSource{} }
func (m *KarmaAddressSource) String() string { return proto.CompactTextString(m) }
func (*KarmaAddressSource) ProtoMessage()    {}
func (*KarmaAddressSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{8}
}
func (m *KarmaAddressSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaAddressSource.Unmarshal(m, b)
}
func (m *KarmaAddressSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaAddressSource.Marshal(b, m, deterministic)
}
func (m *KarmaAddressSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaAddressSource.Merge(m, src)
}
func (m *KarmaAddressSource) XXX_Size() int {
	return xxx_messageInfo_KarmaAddressSource.Size(m)
}
func (m *KarmaAddressSource) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaAddressSource.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaAddressSource proto.InternalMessageInfo

func (m *KarmaAddressSource) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaAddressSource) GetSources() []*KarmaSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

type KarmaState struct {
	SourceStates         []*KarmaSource `protobuf:"bytes,1,rep,name=source_states,json=sourceStates,proto3" json:"source_states,omitempty"`
	DeployKarmaTotal     *types.BigUInt `protobuf:"bytes,2,opt,name=deploy_karma_total,json=deployKarmaTotal,proto3" json:"deploy_karma_total,omitempty"`
	CallKarmaTotal       *types.BigUInt `protobuf:"bytes,3,opt,name=call_karma_total,json=callKarmaTotal,proto3" json:"call_karma_total,omitempty"`
	LastUpdateTime       int64          `protobuf:"varint,4,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaState) Reset()         { *m = KarmaState{} }
func (m *KarmaState) String() string { return proto.CompactTextString(m) }
func (*KarmaState) ProtoMessage()    {}
func (*KarmaState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{9}
}
func (m *KarmaState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaState.Unmarshal(m, b)
}
func (m *KarmaState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaState.Marshal(b, m, deterministic)
}
func (m *KarmaState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaState.Merge(m, src)
}
func (m *KarmaState) XXX_Size() int {
	return xxx_messageInfo_KarmaState.Size(m)
}
func (m *KarmaState) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaState.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaState proto.InternalMessageInfo

func (m *KarmaState) GetSourceStates() []*KarmaSource {
	if m != nil {
		return m.SourceStates
	}
	return nil
}

func (m *KarmaState) GetDeployKarmaTotal() *types.BigUInt {
	if m != nil {
		return m.DeployKarmaTotal
	}
	return nil
}

func (m *KarmaState) GetCallKarmaTotal() *types.BigUInt {
	if m != nil {
		return m.CallKarmaTotal
	}
	return nil
}

func (m *KarmaState) GetLastUpdateTime() int64 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

type KarmaStateUser struct {
	SourceStates         []*KarmaSource `protobuf:"bytes,1,rep,name=source_states,json=sourceStates,proto3" json:"source_states,omitempty"`
	User                 *types.Address `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaStateUser) Reset()         { *m = KarmaStateUser{} }
func (m *KarmaStateUser) String() string { return proto.CompactTextString(m) }
func (*KarmaStateUser) ProtoMessage()    {}
func (*KarmaStateUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{10}
}
func (m *KarmaStateUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaStateUser.Unmarshal(m, b)
}
func (m *KarmaStateUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaStateUser.Marshal(b, m, deterministic)
}
func (m *KarmaStateUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaStateUser.Merge(m, src)
}
func (m *KarmaStateUser) XXX_Size() int {
	return xxx_messageInfo_KarmaStateUser.Size(m)
}
func (m *KarmaStateUser) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaStateUser.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaStateUser proto.InternalMessageInfo

func (m *KarmaStateUser) GetSourceStates() []*KarmaSource {
	if m != nil {
		return m.SourceStates
	}
	return nil
}

func (m *KarmaStateUser) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

type KarmaStateKeyUser struct {
	StateKeys            []string       `protobuf:"bytes,1,rep,name=state_keys,json=stateKeys,proto3" json:"state_keys,omitempty"`
	User                 *types.Address `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaStateKeyUser) Reset()         { *m = KarmaStateKeyUser{} }
func (m *KarmaStateKeyUser) String() string { return proto.CompactTextString(m) }
func (*KarmaStateKeyUser) ProtoMessage()    {}
func (*KarmaStateKeyUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{11}
}
func (m *KarmaStateKeyUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaStateKeyUser.Unmarshal(m, b)
}
func (m *KarmaStateKeyUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaStateKeyUser.Marshal(b, m, deterministic)
}
func (m *KarmaStateKeyUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaStateKeyUser.Merge(m, src)
}
func (m *KarmaStateKeyUser) XXX_Size() int {
	return xxx_messageInfo_KarmaStateKeyUser.Size(m)
}
func (m *KarmaStateKeyUser) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaStateKeyUser.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaStateKeyUser proto.InternalMessageInfo

func (m *KarmaStateKeyUser) GetStateKeys() []string {
	if m != nil {
		return m.StateKeys
	}
	return nil
}

func (m *KarmaStateKeyUser) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

type KarmaTotal struct {
	Count                *types.BigUInt `protobuf:"bytes,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KarmaTotal) Reset()         { *m = KarmaTotal{} }
func (m *KarmaTotal) String() string { return proto.CompactTextString(m) }
func (*KarmaTotal) ProtoMessage()    {}
func (*KarmaTotal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{12}
}
func (m *KarmaTotal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaTotal.Unmarshal(m, b)
}
func (m *KarmaTotal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaTotal.Marshal(b, m, deterministic)
}
func (m *KarmaTotal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaTotal.Merge(m, src)
}
func (m *KarmaTotal) XXX_Size() int {
	return xxx_messageInfo_KarmaTotal.Size(m)
}
func (m *KarmaTotal) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaTotal.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaTotal proto.InternalMessageInfo

func (m *KarmaTotal) GetCount() *types.BigUInt {
	if m != nil {
		return m.Count
	}
	return nil
}

type ContractRecord struct {
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Address              *types.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	CreationBlock        int64          `protobuf:"varint,3,opt,name=creation_block,json=creationBlock,proto3" json:"creation_block,omitempty"`
	Nonce                int64          `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ContractRecord) Reset()         { *m = ContractRecord{} }
func (m *ContractRecord) String() string { return proto.CompactTextString(m) }
func (*ContractRecord) ProtoMessage()    {}
func (*ContractRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{13}
}
func (m *ContractRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractRecord.Unmarshal(m, b)
}
func (m *ContractRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractRecord.Marshal(b, m, deterministic)
}
func (m *ContractRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRecord.Merge(m, src)
}
func (m *ContractRecord) XXX_Size() int {
	return xxx_messageInfo_ContractRecord.Size(m)
}
func (m *ContractRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRecord proto.InternalMessageInfo

func (m *ContractRecord) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ContractRecord) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ContractRecord) GetCreationBlock() int64 {
	if m != nil {
		return m.CreationBlock
	}
	return 0
}

func (m *ContractRecord) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterEnum("karma.KarmaSourceTarget", KarmaSourceTarget_name, KarmaSourceTarget_value)
	proto.RegisterType((*KarmaInitRequest)(nil), "karma.KarmaInitRequest")
	proto.RegisterType((*KarmaSources)(nil), "karma.KarmaSources")
	proto.RegisterType((*KarmaNewOracle)(nil), "karma.KarmaNewOracle")
	proto.RegisterType((*KarmaUserTarget)(nil), "karma.KarmaUserTarget")
	proto.RegisterType((*KarmaUserAmount)(nil), "karma.KarmaUserAmount")
	proto.RegisterType((*KarmaSourceReward)(nil), "karma.KarmaSourceReward")
	proto.RegisterType((*KarmaSource)(nil), "karma.KarmaSource")
	proto.RegisterType((*KarmaUpkeepParams)(nil), "karma.KarmaUpkeepParams")
	proto.RegisterType((*KarmaAddressSource)(nil), "karma.KarmaAddressSource")
	proto.RegisterType((*KarmaState)(nil), "karma.KarmaState")
	proto.RegisterType((*KarmaStateUser)(nil), "karma.KarmaStateUser")
	proto.RegisterType((*KarmaStateKeyUser)(nil), "karma.KarmaStateKeyUser")
	proto.RegisterType((*KarmaTotal)(nil), "karma.KarmaTotal")
	proto.RegisterType((*ContractRecord)(nil), "karma.ContractRecord")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto", fileDescriptor_4325a7150e5ab06e)
}

var fileDescriptor_4325a7150e5ab06e = []byte{
	// 676 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xd1, 0x6a, 0xdb, 0x4a,
	0x10, 0xbd, 0xb2, 0x6c, 0x27, 0x9e, 0x24, 0xbe, 0xbe, 0xcb, 0xe5, 0xe2, 0x5b, 0xda, 0x62, 0x04,
	0xa5, 0xa6, 0xa4, 0x76, 0x70, 0xa1, 0xa5, 0x2f, 0x05, 0x27, 0xed, 0x43, 0x48, 0x68, 0xd2, 0x6d,
	0x42, 0xe9, 0x93, 0xba, 0x96, 0x17, 0x57, 0x58, 0xd2, 0x2a, 0xbb, 0x2b, 0x8c, 0x3f, 0xa3, 0xdf,
	0xd5, 0x4f, 0xe8, 0xcf, 0x94, 0x9d, 0x5d, 0x05, 0x25, 0x16, 0x38, 0xf4, 0x45, 0x68, 0xce, 0x9c,
	0x33, 0x73, 0x66, 0x35, 0x6b, 0xc3, 0xbb, 0x45, 0xac, 0xbf, 0x17, 0xb3, 0x51, 0x24, 0xd2, 0x71,
	0x22, 0x44, 0x9a, 0x71, 0xbd, 0x12, 0x72, 0x39, 0x5e, 0x88, 0x97, 0x26, 0x1c, 0xcf, 0x8a, 0x38,
	0xd1, 0x71, 0x36, 0xd6, 0xeb, 0x9c, 0xab, 0xf1, 0x92, 0xc9, 0x94, 0xd9, 0xe7, 0x28, 0x97, 0x42,
	0x0b, 0xd2, 0xc2, 0xe0, 0xd1, 0xd1, 0x96, 0x32, 0x56, 0x8e, 0x4f, 0x2b, 0x0c, 0x7e, 0x7a, 0xd0,
	0x3b, 0x33, 0xda, 0xd3, 0x2c, 0xd6, 0x94, 0xdf, 0x14, 0x5c, 0x69, 0x32, 0x80, 0xf6, 0x85, 0x64,
	0x51, 0xc2, 0xfb, 0xde, 0xc0, 0x1b, 0xee, 0x4d, 0x76, 0x47, 0xd3, 0xf9, 0x5c, 0x72, 0xa5, 0xa8,
	0xc3, 0xc9, 0x04, 0x76, 0x94, 0x28, 0x64, 0xc4, 0x55, 0xbf, 0x31, 0xf0, 0x87, 0x7b, 0x93, 0xfe,
	0xc8, 0xda, 0xc1, 0x5a, 0x9f, 0x31, 0x45, 0xf9, 0x8a, 0xc9, 0x39, 0x2d, 0x89, 0x64, 0x0c, 0xad,
	0x42, 0x71, 0xa9, 0xfa, 0x3e, 0x2a, 0xfe, 0xaf, 0x2a, 0x5c, 0x7d, 0x27, 0xb4, 0x3c, 0x72, 0x04,
	0xed, 0x22, 0x5f, 0x72, 0x9e, 0xf7, 0x9b, 0x68, 0xe3, 0x4e, 0x8f, 0x6b, 0xcc, 0x5c, 0x32, 0xc9,
	0x52, 0x45, 0x1d, 0x2f, 0x38, 0x86, 0xfd, 0x8a, 0x01, 0x55, 0xb5, 0xe9, 0x3d, 0xd0, 0x66, 0xf0,
	0x16, 0xba, 0x98, 0xfd, 0xc8, 0x57, 0x6e, 0xd8, 0xe7, 0x00, 0x19, 0x5f, 0x85, 0xa2, 0xfe, 0x48,
	0x3a, 0x59, 0x49, 0x0c, 0x18, 0xfc, 0x6d, 0xbd, 0x29, 0x2e, 0xaf, 0x98, 0x5c, 0x70, 0x4d, 0x1e,
	0x43, 0xd3, 0x0c, 0xb3, 0xa1, 0x42, 0xd4, 0x4c, 0xa8, 0x91, 0xd7, 0x6f, 0x0c, 0xbc, 0x61, 0xb7,
	0xce, 0x9e, 0xad, 0x43, 0x1d, 0x2f, 0xf8, 0x54, 0x69, 0x31, 0x4d, 0x45, 0x91, 0x6d, 0x6b, 0x31,
	0x80, 0x36, 0x43, 0x1e, 0xb6, 0x30, 0xf9, 0xe3, 0x78, 0x71, 0x7d, 0x9a, 0x69, 0xea, 0xf0, 0xe0,
	0x06, 0xfe, 0xd9, 0x38, 0x0e, 0x42, 0xa0, 0x99, 0xb1, 0xd4, 0x4e, 0xdb, 0xa1, 0xf8, 0x4e, 0xfe,
	0x83, 0xb6, 0xc4, 0x2c, 0x96, 0xf2, 0xa9, 0x8b, 0x2a, 0x53, 0xf8, 0x0f, 0x9c, 0x62, 0x0a, 0x7b,
	0x95, 0x64, 0x6d, 0xb3, 0xa7, 0xd0, 0x8a, 0x6a, 0x6d, 0x5b, 0x38, 0xf8, 0xe2, 0x5c, 0x57, 0xf7,
	0xc0, 0x14, 0x8a, 0x84, 0xd2, 0x58, 0xc8, 0xa7, 0xf8, 0x6e, 0x5c, 0xdb, 0x4f, 0x8b, 0x95, 0x3a,
	0xd4, 0x45, 0x06, 0xcf, 0xb9, 0x8c, 0xc5, 0x1c, 0x5d, 0xfb, 0xd4, 0x45, 0xc1, 0x37, 0x20, 0x9b,
	0x2b, 0xb9, 0xe5, 0x90, 0x0f, 0xef, 0x5f, 0x07, 0x52, 0xb3, 0x67, 0xb7, 0x1b, 0xf6, 0xcb, 0x03,
	0xb0, 0x09, 0xcd, 0x34, 0x27, 0x6f, 0xe0, 0xc0, 0x66, 0x42, 0x65, 0xe2, 0x72, 0x55, 0xeb, 0x4a,
	0xec, 0x5b, 0x22, 0xea, 0x14, 0x79, 0x0d, 0x64, 0xce, 0xf3, 0x44, 0xac, 0x43, 0x64, 0x86, 0x5a,
	0x68, 0x96, 0x6c, 0x9c, 0x57, 0xcf, 0x72, 0xb0, 0xcc, 0x95, 0x61, 0x90, 0x09, 0xf4, 0x22, 0x96,
	0x24, 0x77, 0x54, 0xfe, 0x3d, 0x55, 0xd7, 0x30, 0x2a, 0x9a, 0x21, 0xf4, 0x12, 0xa6, 0x74, 0x58,
	0xe4, 0x73, 0xa6, 0x79, 0xa8, 0xe3, 0x94, 0xe3, 0xad, 0xf4, 0x69, 0xd7, 0xe0, 0xd7, 0x08, 0x5f,
	0xc5, 0x29, 0x0f, 0x16, 0xee, 0xfe, 0xa0, 0x49, 0xb3, 0xa6, 0x7f, 0x3e, 0x60, 0x79, 0xe8, 0x8d,
	0xba, 0x43, 0x0f, 0x2e, 0xcb, 0xbd, 0x35, 0xe4, 0x33, 0xbe, 0xc6, 0x5e, 0x4f, 0x00, 0xb0, 0x49,
	0xb8, 0xe4, 0x6b, 0xdb, 0xa8, 0x43, 0x3b, 0xca, 0x31, 0xb6, 0x55, 0x3c, 0x74, 0xdf, 0xc5, 0x8e,
	0x7c, 0xbb, 0x81, 0x5e, 0xfd, 0x06, 0xfe, 0xf0, 0xa0, 0x7b, 0x22, 0x32, 0x2d, 0x59, 0xa4, 0x29,
	0x8f, 0x84, 0x9c, 0x1b, 0x89, 0x58, 0x65, 0x35, 0x6b, 0x62, 0x61, 0x12, 0xc0, 0x0e, 0xb3, 0xc8,
	0x86, 0x83, 0x32, 0x41, 0x9e, 0x41, 0x37, 0x92, 0x9c, 0xe9, 0x58, 0x64, 0xe1, 0x2c, 0x11, 0xd1,
	0xd2, 0xed, 0xe7, 0x41, 0x89, 0x1e, 0x1b, 0x90, 0xfc, 0x0b, 0xad, 0x4c, 0x64, 0x51, 0xf9, 0x15,
	0x6c, 0xf0, 0x62, 0x72, 0xe7, 0x2e, 0xbb, 0xdf, 0xa0, 0x5d, 0x68, 0x9e, 0x4c, 0xcf, 0xcf, 0x7b,
	0x7f, 0x11, 0x80, 0xf6, 0xfb, 0x0f, 0x97, 0xe7, 0x17, 0x5f, 0x7b, 0x1e, 0xd9, 0x01, 0xdf, 0x80,
	0x8d, 0x59, 0x1b, 0xff, 0x09, 0x5e, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x00, 0x06, 0x54, 0x73,
	0x84, 0x06, 0x00, 0x00,
}
