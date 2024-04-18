// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: observer/crosschain_flags.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type GasPriceIncreaseFlags struct {
	EpochLength             int64         `protobuf:"varint,1,opt,name=epochLength,proto3" json:"epochLength,omitempty"`
	RetryInterval           time.Duration `protobuf:"bytes,2,opt,name=retryInterval,proto3,stdduration" json:"retryInterval"`
	GasPriceIncreasePercent uint32        `protobuf:"varint,3,opt,name=gasPriceIncreasePercent,proto3" json:"gasPriceIncreasePercent,omitempty"`
	// Maximum gas price increase in percent of the median gas price
	// Default is used if 0
	GasPriceIncreaseMax uint32 `protobuf:"varint,4,opt,name=gasPriceIncreaseMax,proto3" json:"gasPriceIncreaseMax,omitempty"`
	// Maximum number of pending crosschain transactions to check for gas price increase
	MaxPendingCctxs uint32 `protobuf:"varint,5,opt,name=maxPendingCctxs,proto3" json:"maxPendingCctxs,omitempty"`
}

func (m *GasPriceIncreaseFlags) Reset()         { *m = GasPriceIncreaseFlags{} }
func (m *GasPriceIncreaseFlags) String() string { return proto.CompactTextString(m) }
func (*GasPriceIncreaseFlags) ProtoMessage()    {}
func (*GasPriceIncreaseFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_b948b59e4d986f49, []int{0}
}
func (m *GasPriceIncreaseFlags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasPriceIncreaseFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasPriceIncreaseFlags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasPriceIncreaseFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasPriceIncreaseFlags.Merge(m, src)
}
func (m *GasPriceIncreaseFlags) XXX_Size() int {
	return m.Size()
}
func (m *GasPriceIncreaseFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_GasPriceIncreaseFlags.DiscardUnknown(m)
}

var xxx_messageInfo_GasPriceIncreaseFlags proto.InternalMessageInfo

func (m *GasPriceIncreaseFlags) GetEpochLength() int64 {
	if m != nil {
		return m.EpochLength
	}
	return 0
}

func (m *GasPriceIncreaseFlags) GetRetryInterval() time.Duration {
	if m != nil {
		return m.RetryInterval
	}
	return 0
}

func (m *GasPriceIncreaseFlags) GetGasPriceIncreasePercent() uint32 {
	if m != nil {
		return m.GasPriceIncreasePercent
	}
	return 0
}

func (m *GasPriceIncreaseFlags) GetGasPriceIncreaseMax() uint32 {
	if m != nil {
		return m.GasPriceIncreaseMax
	}
	return 0
}

func (m *GasPriceIncreaseFlags) GetMaxPendingCctxs() uint32 {
	if m != nil {
		return m.MaxPendingCctxs
	}
	return 0
}

// Deprecated(v16): Use VerificationFlags in the lightclient module instead
type BlockHeaderVerificationFlags struct {
	IsEthTypeChainEnabled bool `protobuf:"varint,1,opt,name=isEthTypeChainEnabled,proto3" json:"isEthTypeChainEnabled,omitempty"`
	IsBtcTypeChainEnabled bool `protobuf:"varint,2,opt,name=isBtcTypeChainEnabled,proto3" json:"isBtcTypeChainEnabled,omitempty"`
}

func (m *BlockHeaderVerificationFlags) Reset()         { *m = BlockHeaderVerificationFlags{} }
func (m *BlockHeaderVerificationFlags) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderVerificationFlags) ProtoMessage()    {}
func (*BlockHeaderVerificationFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_b948b59e4d986f49, []int{1}
}
func (m *BlockHeaderVerificationFlags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeaderVerificationFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeaderVerificationFlags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeaderVerificationFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderVerificationFlags.Merge(m, src)
}
func (m *BlockHeaderVerificationFlags) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeaderVerificationFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderVerificationFlags.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderVerificationFlags proto.InternalMessageInfo

func (m *BlockHeaderVerificationFlags) GetIsEthTypeChainEnabled() bool {
	if m != nil {
		return m.IsEthTypeChainEnabled
	}
	return false
}

func (m *BlockHeaderVerificationFlags) GetIsBtcTypeChainEnabled() bool {
	if m != nil {
		return m.IsBtcTypeChainEnabled
	}
	return false
}

type CrosschainFlags struct {
	IsInboundEnabled      bool                   `protobuf:"varint,1,opt,name=isInboundEnabled,proto3" json:"isInboundEnabled,omitempty"`
	IsOutboundEnabled     bool                   `protobuf:"varint,2,opt,name=isOutboundEnabled,proto3" json:"isOutboundEnabled,omitempty"`
	GasPriceIncreaseFlags *GasPriceIncreaseFlags `protobuf:"bytes,3,opt,name=gasPriceIncreaseFlags,proto3" json:"gasPriceIncreaseFlags,omitempty"`
	// Deprecated(v16): Use VerificationFlags in the lightclient module instead
	BlockHeaderVerificationFlags *BlockHeaderVerificationFlags `protobuf:"bytes,4,opt,name=blockHeaderVerificationFlags,proto3" json:"blockHeaderVerificationFlags,omitempty"`
}

func (m *CrosschainFlags) Reset()         { *m = CrosschainFlags{} }
func (m *CrosschainFlags) String() string { return proto.CompactTextString(m) }
func (*CrosschainFlags) ProtoMessage()    {}
func (*CrosschainFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_b948b59e4d986f49, []int{2}
}
func (m *CrosschainFlags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrosschainFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrosschainFlags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrosschainFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrosschainFlags.Merge(m, src)
}
func (m *CrosschainFlags) XXX_Size() int {
	return m.Size()
}
func (m *CrosschainFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_CrosschainFlags.DiscardUnknown(m)
}

var xxx_messageInfo_CrosschainFlags proto.InternalMessageInfo

func (m *CrosschainFlags) GetIsInboundEnabled() bool {
	if m != nil {
		return m.IsInboundEnabled
	}
	return false
}

func (m *CrosschainFlags) GetIsOutboundEnabled() bool {
	if m != nil {
		return m.IsOutboundEnabled
	}
	return false
}

func (m *CrosschainFlags) GetGasPriceIncreaseFlags() *GasPriceIncreaseFlags {
	if m != nil {
		return m.GasPriceIncreaseFlags
	}
	return nil
}

func (m *CrosschainFlags) GetBlockHeaderVerificationFlags() *BlockHeaderVerificationFlags {
	if m != nil {
		return m.BlockHeaderVerificationFlags
	}
	return nil
}

type LegacyCrosschainFlags struct {
	IsInboundEnabled      bool                   `protobuf:"varint,1,opt,name=isInboundEnabled,proto3" json:"isInboundEnabled,omitempty"`
	IsOutboundEnabled     bool                   `protobuf:"varint,2,opt,name=isOutboundEnabled,proto3" json:"isOutboundEnabled,omitempty"`
	GasPriceIncreaseFlags *GasPriceIncreaseFlags `protobuf:"bytes,3,opt,name=gasPriceIncreaseFlags,proto3" json:"gasPriceIncreaseFlags,omitempty"`
}

func (m *LegacyCrosschainFlags) Reset()         { *m = LegacyCrosschainFlags{} }
func (m *LegacyCrosschainFlags) String() string { return proto.CompactTextString(m) }
func (*LegacyCrosschainFlags) ProtoMessage()    {}
func (*LegacyCrosschainFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_b948b59e4d986f49, []int{3}
}
func (m *LegacyCrosschainFlags) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyCrosschainFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyCrosschainFlags.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyCrosschainFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyCrosschainFlags.Merge(m, src)
}
func (m *LegacyCrosschainFlags) XXX_Size() int {
	return m.Size()
}
func (m *LegacyCrosschainFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyCrosschainFlags.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyCrosschainFlags proto.InternalMessageInfo

func (m *LegacyCrosschainFlags) GetIsInboundEnabled() bool {
	if m != nil {
		return m.IsInboundEnabled
	}
	return false
}

func (m *LegacyCrosschainFlags) GetIsOutboundEnabled() bool {
	if m != nil {
		return m.IsOutboundEnabled
	}
	return false
}

func (m *LegacyCrosschainFlags) GetGasPriceIncreaseFlags() *GasPriceIncreaseFlags {
	if m != nil {
		return m.GasPriceIncreaseFlags
	}
	return nil
}

func init() {
	proto.RegisterType((*GasPriceIncreaseFlags)(nil), "zetachain.zetacore.observer.GasPriceIncreaseFlags")
	proto.RegisterType((*BlockHeaderVerificationFlags)(nil), "zetachain.zetacore.observer.BlockHeaderVerificationFlags")
	proto.RegisterType((*CrosschainFlags)(nil), "zetachain.zetacore.observer.CrosschainFlags")
	proto.RegisterType((*LegacyCrosschainFlags)(nil), "zetachain.zetacore.observer.LegacyCrosschainFlags")
}

func init() { proto.RegisterFile("observer/crosschain_flags.proto", fileDescriptor_b948b59e4d986f49) }

var fileDescriptor_b948b59e4d986f49 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0x5d, 0x95, 0x65, 0xca, 0xb2, 0x1a, 0x2d, 0xd6, 0x75, 0x49, 0x4b, 0x4f, 0x45,
	0x34, 0x23, 0xd5, 0x83, 0x5e, 0x5b, 0x57, 0x0d, 0xac, 0x58, 0x82, 0x78, 0xf0, 0x22, 0x93, 0xc9,
	0xeb, 0x64, 0x30, 0x3b, 0x53, 0x66, 0x26, 0x4b, 0x2b, 0xf8, 0x05, 0x3c, 0x79, 0x14, 0x3f, 0xd1,
	0x1e, 0xf7, 0xe0, 0x41, 0x10, 0x54, 0xda, 0x2f, 0x22, 0x9d, 0xd8, 0xc5, 0xb6, 0x31, 0x1f, 0xc0,
	0xdb, 0xe4, 0xfd, 0xdf, 0xff, 0xfd, 0x92, 0xf7, 0x5e, 0x06, 0xb7, 0x55, 0x6c, 0x40, 0x9f, 0x82,
	0x26, 0x4c, 0x2b, 0x63, 0x58, 0x4a, 0x85, 0x7c, 0x3b, 0xce, 0x28, 0x37, 0xc1, 0x44, 0x2b, 0xab,
	0xbc, 0xdb, 0xef, 0xc1, 0x52, 0x17, 0x0e, 0xdc, 0x49, 0x69, 0x08, 0x56, 0x9e, 0x83, 0x1b, 0x5c,
	0x71, 0xe5, 0xf2, 0xc8, 0xf2, 0x54, 0x58, 0x0e, 0x7c, 0xae, 0x14, 0xcf, 0x80, 0xb8, 0xa7, 0x38,
	0x1f, 0x93, 0x24, 0xd7, 0xd4, 0x0a, 0x25, 0x0b, 0xbd, 0xfb, 0xa5, 0x8e, 0x9b, 0xcf, 0xa8, 0x19,
	0x69, 0xc1, 0x20, 0x94, 0x4c, 0x03, 0x35, 0xf0, 0x74, 0x89, 0xf4, 0x3a, 0xb8, 0x01, 0x13, 0xc5,
	0xd2, 0x63, 0x90, 0xdc, 0xa6, 0x2d, 0xd4, 0x41, 0xbd, 0x9d, 0xe8, 0xef, 0x90, 0x17, 0xe2, 0x3d,
	0x0d, 0x56, 0xcf, 0x42, 0x69, 0x41, 0x9f, 0xd2, 0xac, 0x55, 0xef, 0xa0, 0x5e, 0xa3, 0x7f, 0x2b,
	0x28, 0x98, 0xc1, 0x8a, 0x19, 0x3c, 0xf9, 0xc3, 0x1c, 0xec, 0x9e, 0xfd, 0x68, 0xd7, 0x3e, 0xff,
	0x6c, 0xa3, 0x68, 0xdd, 0xe9, 0x3d, 0xc2, 0x37, 0xf9, 0xc6, 0x5b, 0x8c, 0x40, 0x33, 0x90, 0xb6,
	0xb5, 0xd3, 0x41, 0xbd, 0xbd, 0xe8, 0x5f, 0xb2, 0x77, 0x1f, 0x5f, 0xdf, 0x94, 0x5e, 0xd0, 0x69,
	0xeb, 0x92, 0x73, 0x95, 0x49, 0x5e, 0x0f, 0xef, 0x9f, 0xd0, 0xe9, 0x08, 0x64, 0x22, 0x24, 0x1f,
	0x32, 0x3b, 0x35, 0xad, 0xcb, 0x2e, 0x7b, 0x33, 0xdc, 0xfd, 0x88, 0xf0, 0xe1, 0x20, 0x53, 0xec,
	0xdd, 0x73, 0xa0, 0x09, 0xe8, 0xd7, 0xa0, 0xc5, 0x58, 0x30, 0xf7, 0x29, 0x45, 0x8f, 0x1e, 0xe2,
	0xa6, 0x30, 0x47, 0x36, 0x7d, 0x35, 0x9b, 0xc0, 0x70, 0x39, 0x97, 0x23, 0x49, 0xe3, 0x0c, 0x12,
	0xd7, 0xad, 0xdd, 0xa8, 0x5c, 0x2c, 0x5c, 0x03, 0xcb, 0xb6, 0x5c, 0xf5, 0x95, 0xab, 0x44, 0xec,
	0x7e, 0xad, 0xe3, 0xfd, 0xe1, 0xc5, 0x5e, 0x14, 0xfc, 0x3b, 0xf8, 0xaa, 0x30, 0xa1, 0x8c, 0x55,
	0x2e, 0x93, 0x75, 0xf4, 0x56, 0xdc, 0xbb, 0x8b, 0xaf, 0x09, 0xf3, 0x32, 0xb7, 0x6b, 0xc9, 0x05,
	0x71, 0x5b, 0xf0, 0x52, 0xdc, 0xe4, 0x65, 0x6b, 0xe1, 0xc6, 0xd1, 0xe8, 0xf7, 0x83, 0x8a, 0x55,
	0x0c, 0x4a, 0x17, 0x2a, 0x2a, 0x2f, 0xe8, 0x7d, 0xc0, 0x87, 0x71, 0x45, 0x8f, 0xdd, 0x24, 0x1b,
	0xfd, 0xc7, 0x95, 0xc0, 0xaa, 0x21, 0x45, 0x95, 0xe5, 0xbb, 0xdf, 0x11, 0x6e, 0x1e, 0x03, 0xa7,
	0x6c, 0xf6, 0x1f, 0x36, 0x77, 0x10, 0x9e, 0xcd, 0x7d, 0x74, 0x3e, 0xf7, 0xd1, 0xaf, 0xb9, 0x8f,
	0x3e, 0x2d, 0xfc, 0xda, 0xf9, 0xc2, 0xaf, 0x7d, 0x5b, 0xf8, 0xb5, 0x37, 0x84, 0x0b, 0x9b, 0xe6,
	0x71, 0xc0, 0xd4, 0x09, 0x59, 0x42, 0xee, 0x39, 0x1e, 0x59, 0xf1, 0xc8, 0x94, 0x5c, 0xdc, 0x46,
	0x76, 0x36, 0x01, 0x13, 0x5f, 0x71, 0xbf, 0xf3, 0x83, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde,
	0xd8, 0x41, 0xc9, 0xa6, 0x04, 0x00, 0x00,
}

func (m *GasPriceIncreaseFlags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasPriceIncreaseFlags) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasPriceIncreaseFlags) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxPendingCctxs != 0 {
		i = encodeVarintCrosschainFlags(dAtA, i, uint64(m.MaxPendingCctxs))
		i--
		dAtA[i] = 0x28
	}
	if m.GasPriceIncreaseMax != 0 {
		i = encodeVarintCrosschainFlags(dAtA, i, uint64(m.GasPriceIncreaseMax))
		i--
		dAtA[i] = 0x20
	}
	if m.GasPriceIncreasePercent != 0 {
		i = encodeVarintCrosschainFlags(dAtA, i, uint64(m.GasPriceIncreasePercent))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.RetryInterval, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.RetryInterval):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCrosschainFlags(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.EpochLength != 0 {
		i = encodeVarintCrosschainFlags(dAtA, i, uint64(m.EpochLength))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeaderVerificationFlags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeaderVerificationFlags) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeaderVerificationFlags) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsBtcTypeChainEnabled {
		i--
		if m.IsBtcTypeChainEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.IsEthTypeChainEnabled {
		i--
		if m.IsEthTypeChainEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CrosschainFlags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrosschainFlags) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrosschainFlags) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeaderVerificationFlags != nil {
		{
			size, err := m.BlockHeaderVerificationFlags.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCrosschainFlags(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.GasPriceIncreaseFlags != nil {
		{
			size, err := m.GasPriceIncreaseFlags.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCrosschainFlags(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.IsOutboundEnabled {
		i--
		if m.IsOutboundEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.IsInboundEnabled {
		i--
		if m.IsInboundEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LegacyCrosschainFlags) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyCrosschainFlags) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyCrosschainFlags) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasPriceIncreaseFlags != nil {
		{
			size, err := m.GasPriceIncreaseFlags.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCrosschainFlags(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.IsOutboundEnabled {
		i--
		if m.IsOutboundEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.IsInboundEnabled {
		i--
		if m.IsInboundEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCrosschainFlags(dAtA []byte, offset int, v uint64) int {
	offset -= sovCrosschainFlags(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GasPriceIncreaseFlags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochLength != 0 {
		n += 1 + sovCrosschainFlags(uint64(m.EpochLength))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.RetryInterval)
	n += 1 + l + sovCrosschainFlags(uint64(l))
	if m.GasPriceIncreasePercent != 0 {
		n += 1 + sovCrosschainFlags(uint64(m.GasPriceIncreasePercent))
	}
	if m.GasPriceIncreaseMax != 0 {
		n += 1 + sovCrosschainFlags(uint64(m.GasPriceIncreaseMax))
	}
	if m.MaxPendingCctxs != 0 {
		n += 1 + sovCrosschainFlags(uint64(m.MaxPendingCctxs))
	}
	return n
}

func (m *BlockHeaderVerificationFlags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsEthTypeChainEnabled {
		n += 2
	}
	if m.IsBtcTypeChainEnabled {
		n += 2
	}
	return n
}

func (m *CrosschainFlags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsInboundEnabled {
		n += 2
	}
	if m.IsOutboundEnabled {
		n += 2
	}
	if m.GasPriceIncreaseFlags != nil {
		l = m.GasPriceIncreaseFlags.Size()
		n += 1 + l + sovCrosschainFlags(uint64(l))
	}
	if m.BlockHeaderVerificationFlags != nil {
		l = m.BlockHeaderVerificationFlags.Size()
		n += 1 + l + sovCrosschainFlags(uint64(l))
	}
	return n
}

func (m *LegacyCrosschainFlags) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsInboundEnabled {
		n += 2
	}
	if m.IsOutboundEnabled {
		n += 2
	}
	if m.GasPriceIncreaseFlags != nil {
		l = m.GasPriceIncreaseFlags.Size()
		n += 1 + l + sovCrosschainFlags(uint64(l))
	}
	return n
}

func sovCrosschainFlags(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCrosschainFlags(x uint64) (n int) {
	return sovCrosschainFlags(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GasPriceIncreaseFlags) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrosschainFlags
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
			return fmt.Errorf("proto: GasPriceIncreaseFlags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasPriceIncreaseFlags: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochLength", wireType)
			}
			m.EpochLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochLength |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetryInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
				return ErrInvalidLengthCrosschainFlags
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.RetryInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceIncreasePercent", wireType)
			}
			m.GasPriceIncreasePercent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPriceIncreasePercent |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceIncreaseMax", wireType)
			}
			m.GasPriceIncreaseMax = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPriceIncreaseMax |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPendingCctxs", wireType)
			}
			m.MaxPendingCctxs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPendingCctxs |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCrosschainFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrosschainFlags
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
func (m *BlockHeaderVerificationFlags) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrosschainFlags
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
			return fmt.Errorf("proto: BlockHeaderVerificationFlags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeaderVerificationFlags: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsEthTypeChainEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
			m.IsEthTypeChainEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsBtcTypeChainEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
			m.IsBtcTypeChainEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCrosschainFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrosschainFlags
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
func (m *CrosschainFlags) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrosschainFlags
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
			return fmt.Errorf("proto: CrosschainFlags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrosschainFlags: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsInboundEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
			m.IsInboundEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsOutboundEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
			m.IsOutboundEnabled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceIncreaseFlags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
				return ErrInvalidLengthCrosschainFlags
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GasPriceIncreaseFlags == nil {
				m.GasPriceIncreaseFlags = &GasPriceIncreaseFlags{}
			}
			if err := m.GasPriceIncreaseFlags.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeaderVerificationFlags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
				return ErrInvalidLengthCrosschainFlags
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockHeaderVerificationFlags == nil {
				m.BlockHeaderVerificationFlags = &BlockHeaderVerificationFlags{}
			}
			if err := m.BlockHeaderVerificationFlags.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrosschainFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrosschainFlags
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
func (m *LegacyCrosschainFlags) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrosschainFlags
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
			return fmt.Errorf("proto: LegacyCrosschainFlags: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyCrosschainFlags: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsInboundEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
			m.IsInboundEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsOutboundEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
			m.IsOutboundEnabled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceIncreaseFlags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrosschainFlags
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
				return ErrInvalidLengthCrosschainFlags
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCrosschainFlags
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GasPriceIncreaseFlags == nil {
				m.GasPriceIncreaseFlags = &GasPriceIncreaseFlags{}
			}
			if err := m.GasPriceIncreaseFlags.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrosschainFlags(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrosschainFlags
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
func skipCrosschainFlags(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCrosschainFlags
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
					return 0, ErrIntOverflowCrosschainFlags
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
					return 0, ErrIntOverflowCrosschainFlags
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
				return 0, ErrInvalidLengthCrosschainFlags
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCrosschainFlags
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCrosschainFlags
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCrosschainFlags        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCrosschainFlags          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCrosschainFlags = fmt.Errorf("proto: unexpected end of group")
)
