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
)

var KarmaSourceTarget_name = map[int32]string{
	0: "CALL",
	1: "DEPLOY",
}

var KarmaSourceTarget_value = map[string]int32{
	"CALL":   0,
	"DEPLOY": 1,
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
	Config               *KarmaConfig          `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
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

func (m *KarmaInitRequest) GetConfig() *KarmaConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type KarmaConfig struct {
	MinKarmaToDeploy     int64    `protobuf:"varint,1,opt,name=min_karma_to_deploy,json=minKarmaToDeploy,proto3" json:"min_karma_to_deploy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KarmaConfig) Reset()         { *m = KarmaConfig{} }
func (m *KarmaConfig) String() string { return proto.CompactTextString(m) }
func (*KarmaConfig) ProtoMessage()    {}
func (*KarmaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{1}
}
func (m *KarmaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KarmaConfig.Unmarshal(m, b)
}
func (m *KarmaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KarmaConfig.Marshal(b, m, deterministic)
}
func (m *KarmaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KarmaConfig.Merge(m, src)
}
func (m *KarmaConfig) XXX_Size() int {
	return xxx_messageInfo_KarmaConfig.Size(m)
}
func (m *KarmaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_KarmaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_KarmaConfig proto.InternalMessageInfo

func (m *KarmaConfig) GetMinKarmaToDeploy() int64 {
	if m != nil {
		return m.MinKarmaToDeploy
	}
	return 0
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
	return fileDescriptor_4325a7150e5ab06e, []int{2}
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
	return fileDescriptor_4325a7150e5ab06e, []int{3}
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
	return fileDescriptor_4325a7150e5ab06e, []int{4}
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
	return fileDescriptor_4325a7150e5ab06e, []int{5}
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
	return fileDescriptor_4325a7150e5ab06e, []int{6}
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
	return fileDescriptor_4325a7150e5ab06e, []int{7}
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
	Period               int64    `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KarmaUpkeepParams) Reset()         { *m = KarmaUpkeepParams{} }
func (m *KarmaUpkeepParams) String() string { return proto.CompactTextString(m) }
func (*KarmaUpkeepParams) ProtoMessage()    {}
func (*KarmaUpkeepParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{8}
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
	return fileDescriptor_4325a7150e5ab06e, []int{9}
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
	return fileDescriptor_4325a7150e5ab06e, []int{10}
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
	return fileDescriptor_4325a7150e5ab06e, []int{11}
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
	return fileDescriptor_4325a7150e5ab06e, []int{12}
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
	return fileDescriptor_4325a7150e5ab06e, []int{13}
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
	return fileDescriptor_4325a7150e5ab06e, []int{14}
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

type GetConfigRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigRequest) Reset()         { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()    {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{15}
}
func (m *GetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigRequest.Unmarshal(m, b)
}
func (m *GetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigRequest.Marshal(b, m, deterministic)
}
func (m *GetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigRequest.Merge(m, src)
}
func (m *GetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigRequest.Size(m)
}
func (m *GetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigRequest proto.InternalMessageInfo

type GetSourceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSourceRequest) Reset()         { *m = GetSourceRequest{} }
func (m *GetSourceRequest) String() string { return proto.CompactTextString(m) }
func (*GetSourceRequest) ProtoMessage()    {}
func (*GetSourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4325a7150e5ab06e, []int{16}
}
func (m *GetSourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSourceRequest.Unmarshal(m, b)
}
func (m *GetSourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSourceRequest.Marshal(b, m, deterministic)
}
func (m *GetSourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSourceRequest.Merge(m, src)
}
func (m *GetSourceRequest) XXX_Size() int {
	return xxx_messageInfo_GetSourceRequest.Size(m)
}
func (m *GetSourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSourceRequest proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("karma.KarmaSourceTarget", KarmaSourceTarget_name, KarmaSourceTarget_value)
	proto.RegisterType((*KarmaInitRequest)(nil), "karma.KarmaInitRequest")
	proto.RegisterType((*KarmaConfig)(nil), "karma.KarmaConfig")
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
	proto.RegisterType((*GetConfigRequest)(nil), "karma.GetConfigRequest")
	proto.RegisterType((*GetSourceRequest)(nil), "karma.GetSourceRequest")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto", fileDescriptor_4325a7150e5ab06e)
}

var fileDescriptor_4325a7150e5ab06e = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0xbd, 0xae, 0x93, 0xb4, 0x99, 0xb6, 0xb9, 0x66, 0x41, 0x28, 0x20, 0x40, 0x91, 0x25, 0x44,
	0xa8, 0xda, 0xa4, 0x0a, 0x12, 0x08, 0x09, 0x81, 0xd2, 0x16, 0xa1, 0xaa, 0x15, 0x2d, 0xa6, 0x7d,
	0xe0, 0xc9, 0x6c, 0x9c, 0x25, 0x58, 0xb1, 0x77, 0xd3, 0xdd, 0xb5, 0xa2, 0x7c, 0x06, 0xdf, 0xc7,
	0x87, 0xf0, 0x8a, 0x3c, 0xbb, 0xae, 0xdc, 0xc6, 0x52, 0x2a, 0x5e, 0xa2, 0xec, 0x99, 0x73, 0x66,
	0x66, 0xc7, 0x67, 0x6c, 0x78, 0x3f, 0x89, 0xf5, 0xcf, 0x6c, 0xd4, 0x8b, 0x44, 0xda, 0x4f, 0x84,
	0x48, 0x39, 0xd3, 0x73, 0x21, 0xa7, 0xfd, 0x89, 0xd8, 0xcb, 0x8f, 0xfd, 0x51, 0x16, 0x27, 0x3a,
	0xe6, 0x7d, 0xbd, 0x98, 0x31, 0xd5, 0x9f, 0x52, 0x99, 0x52, 0xf3, 0xdb, 0x9b, 0x49, 0xa1, 0x05,
	0xa9, 0xe3, 0xe1, 0xf1, 0xfe, 0x8a, 0x34, 0x46, 0x8e, 0xbf, 0x46, 0xe8, 0xff, 0x71, 0xc0, 0x3b,
	0xc9, 0xb5, 0xc7, 0x3c, 0xd6, 0x01, 0xbb, 0xca, 0x98, 0xd2, 0xa4, 0x03, 0x8d, 0x33, 0x49, 0xa3,
	0x84, 0xb5, 0x9d, 0x8e, 0xd3, 0xdd, 0x1c, 0x6c, 0xf4, 0x86, 0xe3, 0xb1, 0x64, 0x4a, 0x05, 0x16,
	0x27, 0x03, 0x58, 0x57, 0x22, 0x93, 0x11, 0x53, 0xed, 0xb5, 0x8e, 0xdb, 0xdd, 0x1c, 0xb4, 0x7b,
	0xa6, 0x1d, 0xcc, 0xf5, 0x15, 0x43, 0x01, 0x9b, 0x53, 0x39, 0x0e, 0x0a, 0x22, 0xe9, 0x43, 0x3d,
	0x53, 0x4c, 0xaa, 0xb6, 0x8b, 0x8a, 0x47, 0x65, 0x85, 0xcd, 0x6f, 0x85, 0x86, 0x47, 0xf6, 0xa1,
	0x91, 0xcd, 0xa6, 0x8c, 0xcd, 0xda, 0x35, 0x6c, 0xe3, 0x46, 0x8d, 0x4b, 0x8c, 0x9c, 0x53, 0x49,
	0x53, 0x15, 0x58, 0x1e, 0xd9, 0x81, 0x46, 0x24, 0xf8, 0x8f, 0x78, 0xd2, 0xae, 0xa3, 0x82, 0x94,
	0x15, 0x87, 0x18, 0x09, 0x2c, 0xc3, 0x7f, 0x07, 0x9b, 0x25, 0x98, 0xec, 0xc1, 0xfd, 0x34, 0xe6,
	0x21, 0xf2, 0x43, 0x2d, 0xc2, 0x31, 0x9b, 0x25, 0x62, 0x81, 0x03, 0x70, 0x03, 0x2f, 0x8d, 0x39,
	0x92, 0x2f, 0xc4, 0x11, 0xe2, 0xfe, 0x01, 0x6c, 0x95, 0xae, 0xaa, 0xca, 0x03, 0x71, 0xee, 0x38,
	0x10, 0xff, 0x2d, 0xb4, 0x30, 0xfa, 0x99, 0xcd, 0xed, 0x58, 0x5f, 0x00, 0x70, 0x36, 0x0f, 0x45,
	0xf5, 0xf0, 0x9b, 0xbc, 0x20, 0xfa, 0x14, 0xfe, 0x37, 0x53, 0x50, 0x4c, 0x5e, 0x50, 0x39, 0x61,
	0x9a, 0x3c, 0x81, 0x5a, 0x3e, 0xb6, 0x25, 0x15, 0xa2, 0xf9, 0x2c, 0x35, 0xf2, 0xda, 0x6b, 0x1d,
	0xa7, 0xdb, 0xaa, 0x6a, 0xcf, 0xe4, 0x09, 0x2c, 0xcf, 0xff, 0x52, 0x2a, 0x31, 0x4c, 0x45, 0xc6,
	0x57, 0x95, 0xe8, 0x40, 0x83, 0x22, 0x0f, 0x4b, 0xe4, 0xf1, 0x83, 0x78, 0x72, 0x79, 0xcc, 0x75,
	0x60, 0x71, 0xff, 0x0a, 0xee, 0x2d, 0x8d, 0x83, 0x10, 0xa8, 0x71, 0x9a, 0x9a, 0xdb, 0x36, 0x03,
	0xfc, 0x4f, 0x1e, 0x42, 0x43, 0x62, 0x14, 0x53, 0xb9, 0x81, 0x3d, 0x95, 0x6e, 0xe1, 0xde, 0xf1,
	0x16, 0x43, 0xfb, 0x94, 0x4d, 0xb0, 0xb2, 0xd8, 0x33, 0xa8, 0x47, 0x95, 0x6d, 0x1b, 0xd8, 0xff,
	0x60, 0xbb, 0x2e, 0x3b, 0x2e, 0x4f, 0x14, 0x09, 0xa5, 0xad, 0x3f, 0xf0, 0x7f, 0xde, 0xf5, 0x8c,
	0xc9, 0x58, 0x8c, 0xb1, 0x3b, 0x37, 0xb0, 0x27, 0xff, 0x3b, 0x90, 0x65, 0x93, 0xaf, 0x18, 0xe6,
	0xee, 0xed, 0x05, 0x23, 0x15, 0x7e, 0xba, 0x76, 0xd2, 0x6f, 0x07, 0xc0, 0x04, 0x34, 0xd5, 0x8c,
	0xbc, 0x81, 0x6d, 0x13, 0x09, 0x55, 0x7e, 0x2e, 0x2c, 0x59, 0x95, 0x62, 0xcb, 0x10, 0x51, 0xa7,
	0xc8, 0x6b, 0x20, 0xc6, 0xf7, 0xd7, 0x7b, 0xa0, 0x69, 0xb2, 0x34, 0x17, 0xcf, 0x70, 0xec, 0x42,
	0x68, 0x9a, 0x90, 0x01, 0x78, 0x11, 0x4d, 0x92, 0x1b, 0x2a, 0xf7, 0x96, 0xaa, 0x95, 0x33, 0x4a,
	0x9a, 0x2e, 0x78, 0x09, 0x55, 0x3a, 0xcc, 0x66, 0x63, 0xaa, 0x59, 0xa8, 0xe3, 0x94, 0xe1, 0x9e,
	0xbb, 0x41, 0x2b, 0xc7, 0x2f, 0x11, 0xbe, 0x88, 0x53, 0xe6, 0x4f, 0xec, 0x9e, 0x60, 0x93, 0xb9,
	0x1d, 0xff, 0xfd, 0x82, 0xc5, 0xd0, 0xd7, 0xaa, 0x86, 0xee, 0x9f, 0x17, 0xfe, 0xcc, 0xc9, 0x27,
	0x6c, 0x81, 0xb5, 0x9e, 0x02, 0x60, 0x91, 0x70, 0xca, 0x16, 0xa6, 0x50, 0x33, 0x68, 0x2a, 0xcb,
	0x58, 0x95, 0x71, 0xd7, 0x3e, 0x17, 0x73, 0xe5, 0x6b, 0xa7, 0x39, 0xd5, 0x4e, 0xfb, 0xe5, 0x40,
	0xeb, 0x50, 0x70, 0x2d, 0x69, 0xa4, 0x03, 0x16, 0x09, 0x39, 0xce, 0x25, 0x62, 0xce, 0x2b, 0x6c,
	0x62, 0x60, 0xe2, 0xc3, 0x3a, 0x35, 0xc8, 0x52, 0x07, 0x45, 0x80, 0x3c, 0x87, 0x56, 0x24, 0x19,
	0xd5, 0xb1, 0xe0, 0xe1, 0x28, 0x11, 0xd1, 0xd4, 0xfa, 0x73, 0xbb, 0x40, 0x0f, 0x72, 0x90, 0x3c,
	0x80, 0x3a, 0x17, 0x3c, 0x2a, 0x9e, 0x82, 0x39, 0xf8, 0x04, 0xbc, 0x4f, 0x4c, 0xdb, 0x77, 0xa7,
	0xf9, 0x3e, 0x58, 0xac, 0xd8, 0x62, 0xc4, 0x76, 0x5e, 0xde, 0xd8, 0x6d, 0xfb, 0x4e, 0xda, 0x80,
	0xda, 0xe1, 0xf0, 0xf4, 0xd4, 0xfb, 0x8f, 0x00, 0x34, 0x8e, 0x3e, 0x9e, 0x9f, 0x9e, 0x7d, 0xf3,
	0x9c, 0x51, 0x03, 0x3f, 0x3d, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x42, 0x2f, 0x42,
	0xf5, 0x06, 0x00, 0x00,
}
