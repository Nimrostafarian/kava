// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/issuance/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the issuance module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params   Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Supplies []AssetSupply `protobuf:"bytes,2,rep,name=supplies,proto3" json:"supplies" yaml:"supplies,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e567e34e5c078b96, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSupplies() []AssetSupply {
	if m != nil {
		return m.Supplies
	}
	return nil
}

// Params defines the parameters for the issuance module.
type Params struct {
	Assets []Asset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets" yaml:"assets,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e567e34e5c078b96, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAssets() []Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

// Asset type for assets in the issuance module
type Asset struct {
	Owner            string    `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner,omitempty"`
	Denom            string    `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom,omitempty"`
	BlockedAddresses []string  `protobuf:"bytes,3,rep,name=blocked_addresses,json=blockedAddresses,proto3" json:"blocked_addresses,omitempty" yaml:"blocked_addresses,omitempty"`
	Paused           bool      `protobuf:"varint,4,opt,name=paused,proto3" json:"paused,omitempty" yaml:"paused,omitempty"`
	Blockable        bool      `protobuf:"varint,5,opt,name=blockable,proto3" json:"blockable,omitempty" yaml:"blockable,omitempty"`
	RateLimit        RateLimit `protobuf:"bytes,6,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit" yaml:"rate_limit,omitempty"`
}

func (m *Asset) Reset()      { *m = Asset{} }
func (*Asset) ProtoMessage() {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_e567e34e5c078b96, []int{2}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Asset) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Asset) GetBlockedAddresses() []string {
	if m != nil {
		return m.BlockedAddresses
	}
	return nil
}

func (m *Asset) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func (m *Asset) GetBlockable() bool {
	if m != nil {
		return m.Blockable
	}
	return false
}

func (m *Asset) GetRateLimit() RateLimit {
	if m != nil {
		return m.RateLimit
	}
	return RateLimit{}
}

// RateLimit parameters for rate-limiting the supply of an issued asset
type RateLimit struct {
	Active     bool                                   `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty" yaml:"active,omitempty"`
	Limit      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=limit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"limit,omitempty" yaml:"limit"`
	TimePeriod time.Duration                          `protobuf:"bytes,3,opt,name=time_period,json=timePeriod,proto3,stdduration" json:"time_period" yaml:"time_period,omitempty"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e567e34e5c078b96, []int{3}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return m.Size()
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *RateLimit) GetTimePeriod() time.Duration {
	if m != nil {
		return m.TimePeriod
	}
	return 0
}

// AssetSupply contains information about an asset's rate-limited supply (the
// total supply of the asset is tracked in the top-level supply module)
type AssetSupply struct {
	CurrentSupply types.Coin    `protobuf:"bytes,1,opt,name=current_supply,json=currentSupply,proto3" json:"current_supply"`
	TimeElapsed   time.Duration `protobuf:"bytes,3,opt,name=time_elapsed,json=timeElapsed,proto3,stdduration" json:"time_elapsed" yaml:"time_elapsed,omitempty"`
}

func (m *AssetSupply) Reset()      { *m = AssetSupply{} }
func (*AssetSupply) ProtoMessage() {}
func (*AssetSupply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e567e34e5c078b96, []int{4}
}
func (m *AssetSupply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetSupply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetSupply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetSupply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetSupply.Merge(m, src)
}
func (m *AssetSupply) XXX_Size() int {
	return m.Size()
}
func (m *AssetSupply) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetSupply.DiscardUnknown(m)
}

var xxx_messageInfo_AssetSupply proto.InternalMessageInfo

func (m *AssetSupply) GetCurrentSupply() types.Coin {
	if m != nil {
		return m.CurrentSupply
	}
	return types.Coin{}
}

func (m *AssetSupply) GetTimeElapsed() time.Duration {
	if m != nil {
		return m.TimeElapsed
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kava.issuance.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "kava.issuance.v1beta1.Params")
	proto.RegisterType((*Asset)(nil), "kava.issuance.v1beta1.Asset")
	proto.RegisterType((*RateLimit)(nil), "kava.issuance.v1beta1.RateLimit")
	proto.RegisterType((*AssetSupply)(nil), "kava.issuance.v1beta1.AssetSupply")
}

func init() {
	proto.RegisterFile("kava/issuance/v1beta1/genesis.proto", fileDescriptor_e567e34e5c078b96)
}

var fileDescriptor_e567e34e5c078b96 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0x9b, 0x26, 0x6a, 0x26, 0xf9, 0x3e, 0xc0, 0xe2, 0xc7, 0x4d, 0x5b, 0x3b, 0xb8, 0x52,
	0x15, 0x10, 0xb5, 0x69, 0xbb, 0x2b, 0x6c, 0x6a, 0x0a, 0x08, 0x89, 0x45, 0x49, 0xc5, 0x86, 0x4d,
	0x34, 0xb6, 0x07, 0x33, 0xaa, 0xed, 0xb1, 0x3c, 0xe3, 0x42, 0x9e, 0x02, 0x96, 0xdd, 0xc1, 0x13,
	0xf0, 0x14, 0x2c, 0xba, 0xec, 0x12, 0xb1, 0x30, 0xa8, 0xdd, 0xb1, 0xcc, 0x13, 0xa0, 0xf9, 0x49,
	0xe3, 0x14, 0x0a, 0xab, 0x4c, 0xee, 0xbd, 0xe7, 0xdc, 0x73, 0xe7, 0xcc, 0x35, 0x58, 0x3d, 0x80,
	0x87, 0xd0, 0xc5, 0x94, 0x16, 0x30, 0x0d, 0x90, 0x7b, 0xb8, 0xe1, 0x23, 0x06, 0x37, 0xdc, 0x08,
	0xa5, 0x88, 0x62, 0xea, 0x64, 0x39, 0x61, 0x44, 0xbf, 0xc1, 0x8b, 0x9c, 0x49, 0x91, 0xa3, 0x8a,
	0xba, 0x66, 0x40, 0x68, 0x42, 0xa8, 0xeb, 0x43, 0x3a, 0x45, 0x06, 0x04, 0xa7, 0x12, 0xd6, 0xbd,
	0x1e, 0x91, 0x88, 0x88, 0xa3, 0xcb, 0x4f, 0x2a, 0x6a, 0x46, 0x84, 0x44, 0x31, 0x72, 0xc5, 0x3f,
	0xbf, 0x78, 0xed, 0x86, 0x45, 0x0e, 0x19, 0x26, 0x0a, 0x65, 0x7f, 0xd6, 0x40, 0xe7, 0xa9, 0x6c,
	0xbf, 0xcf, 0x20, 0x43, 0xfa, 0x03, 0xd0, 0xcc, 0x60, 0x0e, 0x13, 0x6a, 0x68, 0x3d, 0xad, 0xdf,
	0xde, 0x5c, 0x71, 0xfe, 0x28, 0xc7, 0xd9, 0x13, 0x45, 0xde, 0xfc, 0x71, 0x69, 0xd5, 0x06, 0x0a,
	0xa2, 0x43, 0xb0, 0x40, 0x8b, 0x2c, 0x8b, 0x31, 0xa2, 0xc6, 0x5c, 0xaf, 0xde, 0x6f, 0x6f, 0xda,
	0x97, 0xc0, 0x77, 0x28, 0x45, 0x6c, 0x9f, 0xd7, 0x8e, 0xbc, 0xdb, 0x9c, 0x63, 0x5c, 0x5a, 0x8b,
	0x23, 0x98, 0xc4, 0xdb, 0xf6, 0x84, 0xe1, 0x1e, 0x49, 0x30, 0x43, 0x49, 0xc6, 0x46, 0xf6, 0xe0,
	0x9c, 0xd6, 0x46, 0xa0, 0x29, 0x5b, 0xeb, 0x2f, 0x41, 0x13, 0x72, 0x16, 0xae, 0x94, 0xb7, 0x5a,
	0xfe, 0x5b, 0x2b, 0xcf, 0x52, 0x4d, 0x6e, 0xc9, 0x26, 0x12, 0x59, 0x6d, 0xa1, 0xc8, 0xb6, 0xe7,
	0x8f, 0x3e, 0x59, 0x35, 0xfb, 0x63, 0x1d, 0x34, 0x04, 0x50, 0xbf, 0x0f, 0x1a, 0xe4, 0x6d, 0x8a,
	0x72, 0x71, 0x1f, 0x2d, 0xaf, 0x3b, 0x2e, 0xad, 0x9b, 0x92, 0x43, 0x84, 0xab, 0x14, 0xb2, 0x90,
	0x23, 0x42, 0x94, 0x92, 0xc4, 0x98, 0xbb, 0x88, 0x10, 0xe1, 0x19, 0x84, 0x88, 0xe8, 0xfb, 0xe0,
	0x9a, 0x1f, 0x93, 0xe0, 0x00, 0x85, 0x43, 0x18, 0x86, 0x39, 0xa2, 0x14, 0x51, 0xa3, 0xde, 0xab,
	0xf7, 0x5b, 0xde, 0xda, 0xb8, 0xb4, 0x6c, 0x89, 0xfe, 0xad, 0xa4, 0xca, 0x74, 0x55, 0x65, 0x77,
	0x26, 0x49, 0x7d, 0x8b, 0x3b, 0x59, 0x50, 0x14, 0x1a, 0xf3, 0x3d, 0xad, 0xbf, 0xe0, 0x2d, 0x4d,
	0xa7, 0x97, 0xf1, 0x99, 0xe9, 0x65, 0x48, 0x7f, 0x08, 0x5a, 0x82, 0x08, 0xfa, 0x31, 0x32, 0x1a,
	0x02, 0x67, 0x8e, 0x4b, 0xab, 0x5b, 0x51, 0xc0, 0x53, 0x55, 0xe8, 0x14, 0xa0, 0x87, 0x00, 0xe4,
	0x90, 0xa1, 0x61, 0x8c, 0x13, 0xcc, 0x8c, 0xa6, 0x78, 0x40, 0xbd, 0x4b, 0x6c, 0x19, 0x40, 0x86,
	0x9e, 0xf3, 0x3a, 0x6f, 0x55, 0x59, 0xb3, 0x24, 0x9b, 0x4c, 0x19, 0x66, 0xba, 0xe4, 0x93, 0x7a,
	0xe5, 0xd0, 0xfb, 0x39, 0xd0, 0x3a, 0xe7, 0xe0, 0xc3, 0xc2, 0x80, 0xe1, 0x43, 0x24, 0x6c, 0x9a,
	0x19, 0x56, 0xc6, 0x67, 0xad, 0x16, 0x21, 0x3d, 0x02, 0x0d, 0xa9, 0x94, 0x1b, 0xd5, 0xf1, 0x5e,
	0x70, 0x1d, 0xdf, 0x4a, 0x6b, 0x2d, 0xc2, 0xec, 0x4d, 0xe1, 0x3b, 0x01, 0x49, 0x5c, 0xb5, 0x74,
	0xf2, 0x67, 0x9d, 0x86, 0x07, 0x2e, 0x1b, 0x65, 0x88, 0x3a, 0xcf, 0x52, 0xf6, 0xb3, 0xb4, 0xae,
	0x5c, 0x90, 0x39, 0x2e, 0xad, 0x8e, 0x6c, 0x2a, 0x12, 0xf6, 0x40, 0xf2, 0xeb, 0x21, 0x68, 0x33,
	0x9c, 0xa0, 0x61, 0x86, 0x72, 0x4c, 0x42, 0xa3, 0x2e, 0x2e, 0x66, 0xd1, 0x91, 0xbb, 0xe9, 0x4c,
	0x76, 0xd3, 0xd9, 0x55, 0xbb, 0xe9, 0xf5, 0xd5, 0x8d, 0x2c, 0x4b, 0xb2, 0x0a, 0xb6, 0x32, 0xc6,
	0xd1, 0x77, 0x4b, 0x1b, 0x00, 0x9e, 0xdb, 0x13, 0x29, 0xfb, 0x8b, 0x06, 0xda, 0x95, 0xbd, 0xd2,
	0x9f, 0x80, 0xff, 0x83, 0x22, 0xcf, 0x51, 0xca, 0x86, 0x62, 0x7d, 0x46, 0x6a, 0xa5, 0x17, 0x1d,
	0x39, 0x8e, 0xc3, 0x3f, 0x25, 0xe7, 0x7e, 0x3c, 0x22, 0x38, 0x55, 0xeb, 0xfc, 0x9f, 0x82, 0x29,
	0x9e, 0x08, 0x74, 0x84, 0x02, 0x14, 0xc3, 0x8c, 0x3f, 0xa7, 0x7f, 0xca, 0xbf, 0xa3, 0xe4, 0xaf,
	0x54, 0xe4, 0x2b, 0xf0, 0x45, 0xfd, 0xe2, 0x5e, 0x1e, 0xcb, 0x9c, 0x34, 0xd6, 0xdb, 0x3d, 0x3e,
	0x35, 0xb5, 0x93, 0x53, 0x53, 0xfb, 0x71, 0x6a, 0x6a, 0x1f, 0xce, 0xcc, 0xda, 0xc9, 0x99, 0x59,
	0xfb, 0x7a, 0x66, 0xd6, 0x5e, 0xdd, 0xad, 0x18, 0xc3, 0x1f, 0xd5, 0x7a, 0x0c, 0x7d, 0x2a, 0x4e,
	0xee, 0xbb, 0xe9, 0x57, 0x55, 0x18, 0xe4, 0x37, 0x85, 0xac, 0xad, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xb4, 0x61, 0x43, 0x86, 0x73, 0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Supplies) > 0 {
		for iNdEx := len(m.Supplies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Supplies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RateLimit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.Blockable {
		i--
		if m.Blockable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.BlockedAddresses) > 0 {
		for iNdEx := len(m.BlockedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.BlockedAddresses[iNdEx])
			copy(dAtA[i:], m.BlockedAddresses[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.BlockedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.TimePeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.TimePeriod):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	{
		size := m.Limit.Size()
		i -= size
		if _, err := m.Limit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AssetSupply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetSupply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetSupply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.TimeElapsed, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.TimeElapsed):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintGenesis(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.CurrentSupply.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Supplies) > 0 {
		for _, e := range m.Supplies {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.BlockedAddresses) > 0 {
		for _, s := range m.BlockedAddresses {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Paused {
		n += 2
	}
	if m.Blockable {
		n += 2
	}
	l = m.RateLimit.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *RateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Active {
		n += 2
	}
	l = m.Limit.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.TimePeriod)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *AssetSupply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CurrentSupply.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.TimeElapsed)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supplies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Supplies = append(m.Supplies, AssetSupply{})
			if err := m.Supplies[len(m.Supplies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, Asset{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockedAddresses = append(m.BlockedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blockable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Blockable = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Limit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimePeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.TimePeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AssetSupply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AssetSupply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetSupply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSupply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeElapsed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.TimeElapsed, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
