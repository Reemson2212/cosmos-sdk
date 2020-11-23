// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/client/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the ibc client submodule's genesis state.
type GenesisState struct {
	// client states with their corresponding identifiers
	Clients IdentifiedClientStates `protobuf:"bytes,1,rep,name=clients,proto3,castrepeated=IdentifiedClientStates" json:"clients"`
	// consensus states from each client
	ClientsConsensus ClientsConsensusStates `protobuf:"bytes,2,rep,name=clients_consensus,json=clientsConsensus,proto3,castrepeated=ClientsConsensusStates" json:"clients_consensus" yaml:"clients_consensus"`
	// metadata from each client
	ClientsMetadata []IdentifiedGenesisMetadata `protobuf:"bytes,3,rep,name=clients_metadata,json=clientsMetadata,proto3" json:"clients_metadata" yaml:"clients_metadata"`
	Params          Params                      `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
	// create localhost on initialization
	CreateLocalhost bool `protobuf:"varint,5,opt,name=create_localhost,json=createLocalhost,proto3" json:"create_localhost,omitempty" yaml:"create_localhost"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcd0c0f1f2e6a91a, []int{0}
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

func (m *GenesisState) GetClients() IdentifiedClientStates {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (m *GenesisState) GetClientsConsensus() ClientsConsensusStates {
	if m != nil {
		return m.ClientsConsensus
	}
	return nil
}

func (m *GenesisState) GetClientsMetadata() []IdentifiedGenesisMetadata {
	if m != nil {
		return m.ClientsMetadata
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetCreateLocalhost() bool {
	if m != nil {
		return m.CreateLocalhost
	}
	return false
}

// GenesisMetadata defines the genesis type for metadata that clients may return
// with ExportMetadata
type GenesisMetadata struct {
	// store key of metadata without clientID-prefix
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// metadata value
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GenesisMetadata) Reset()         { *m = GenesisMetadata{} }
func (m *GenesisMetadata) String() string { return proto.CompactTextString(m) }
func (*GenesisMetadata) ProtoMessage()    {}
func (*GenesisMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcd0c0f1f2e6a91a, []int{1}
}
func (m *GenesisMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisMetadata.Merge(m, src)
}
func (m *GenesisMetadata) XXX_Size() int {
	return m.Size()
}
func (m *GenesisMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisMetadata proto.InternalMessageInfo

// IdentifiedGenesisMetadata has the client metadata with the corresponding client id.
type IdentifiedGenesisMetadata struct {
	ClientId       string            `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" yaml:"client_id"`
	ClientMetadata []GenesisMetadata `protobuf:"bytes,2,rep,name=client_metadata,json=clientMetadata,proto3" json:"client_metadata" yaml:"client_metadata"`
}

func (m *IdentifiedGenesisMetadata) Reset()         { *m = IdentifiedGenesisMetadata{} }
func (m *IdentifiedGenesisMetadata) String() string { return proto.CompactTextString(m) }
func (*IdentifiedGenesisMetadata) ProtoMessage()    {}
func (*IdentifiedGenesisMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcd0c0f1f2e6a91a, []int{2}
}
func (m *IdentifiedGenesisMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentifiedGenesisMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentifiedGenesisMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentifiedGenesisMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifiedGenesisMetadata.Merge(m, src)
}
func (m *IdentifiedGenesisMetadata) XXX_Size() int {
	return m.Size()
}
func (m *IdentifiedGenesisMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifiedGenesisMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifiedGenesisMetadata proto.InternalMessageInfo

func (m *IdentifiedGenesisMetadata) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *IdentifiedGenesisMetadata) GetClientMetadata() []GenesisMetadata {
	if m != nil {
		return m.ClientMetadata
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.client.v1.GenesisState")
	proto.RegisterType((*GenesisMetadata)(nil), "ibc.core.client.v1.GenesisMetadata")
	proto.RegisterType((*IdentifiedGenesisMetadata)(nil), "ibc.core.client.v1.IdentifiedGenesisMetadata")
}

func init() { proto.RegisterFile("ibc/core/client/v1/genesis.proto", fileDescriptor_bcd0c0f1f2e6a91a) }

var fileDescriptor_bcd0c0f1f2e6a91a = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x8e, 0xd7, 0x6e, 0x6c, 0xde, 0xc4, 0x8a, 0x55, 0x8d, 0x50, 0xa4, 0x24, 0x0a, 0x97, 0x72,
	0x68, 0x42, 0xcb, 0x01, 0xd4, 0x0b, 0x52, 0x26, 0x81, 0x26, 0x81, 0x04, 0xe1, 0xc6, 0x25, 0x72,
	0x1d, 0xd3, 0x45, 0x4b, 0xe2, 0xaa, 0x76, 0x2b, 0xfa, 0x00, 0x48, 0x1c, 0x11, 0x4f, 0xc0, 0x99,
	0x67, 0xe0, 0x01, 0x76, 0xdc, 0x91, 0x53, 0x41, 0xed, 0x1b, 0xec, 0x09, 0x50, 0x6c, 0xa7, 0xdb,
	0xb2, 0x76, 0xa7, 0xfe, 0xfd, 0xfc, 0x7d, 0xdf, 0xff, 0xfb, 0xfb, 0x63, 0xe8, 0x24, 0x03, 0xe2,
	0x13, 0x36, 0xa6, 0x3e, 0x49, 0x13, 0x9a, 0x0b, 0x7f, 0xda, 0xf5, 0x87, 0x34, 0xa7, 0x3c, 0xe1,
	0xde, 0x68, 0xcc, 0x04, 0x43, 0x28, 0x19, 0x10, 0xaf, 0x60, 0x78, 0x8a, 0xe1, 0x4d, 0xbb, 0x2d,
	0x7b, 0x8d, 0x4a, 0x9f, 0x4a, 0x51, 0xab, 0x39, 0x64, 0x43, 0x26, 0x4b, 0xbf, 0xa8, 0x14, 0xea,
	0x7e, 0xad, 0xc3, 0x83, 0x37, 0xca, 0xfc, 0xa3, 0xc0, 0x82, 0x22, 0x02, 0xef, 0x29, 0x19, 0x37,
	0x81, 0x53, 0x6b, 0xef, 0xf7, 0x9e, 0x7a, 0xb7, 0xbb, 0x79, 0x27, 0x31, 0xcd, 0x45, 0xf2, 0x39,
	0xa1, 0xf1, 0xb1, 0xc4, 0xa4, 0x36, 0xb0, 0xce, 0xe7, 0xb6, 0xf1, 0xeb, 0xaf, 0x7d, 0xb4, 0xf6,
	0x98, 0x87, 0xa5, 0x33, 0xfa, 0x01, 0xe0, 0x03, 0x5d, 0x47, 0x84, 0xe5, 0x9c, 0xe6, 0x7c, 0xc2,
	0xcd, 0xad, 0xcd, 0xfd, 0x94, 0xcd, 0x71, 0x49, 0x55, 0x7e, 0x41, 0xbf, 0xe8, 0x77, 0x39, 0xb7,
	0xcd, 0x19, 0xce, 0xd2, 0xbe, 0x7b, 0xcb, 0xd1, 0x2d, 0x66, 0x51, 0x52, 0x5e, 0xd1, 0x86, 0x0d,
	0x52, 0xc1, 0xd1, 0x0c, 0x96, 0x58, 0x94, 0x51, 0x81, 0x63, 0x2c, 0xb0, 0x59, 0x93, 0x23, 0x75,
	0xee, 0x8e, 0x40, 0xe7, 0xf7, 0x4e, 0x8b, 0x02, 0x5b, 0x8f, 0xf5, 0xf0, 0xe6, 0x58, 0xa5, 0xa9,
	0x1b, 0x1e, 0x6a, 0xa8, 0x54, 0xa0, 0x97, 0x70, 0x67, 0x84, 0xc7, 0x38, 0xe3, 0x66, 0xdd, 0x01,
	0xed, 0xfd, 0x5e, 0x6b, 0x5d, 0xc3, 0xf7, 0x92, 0x11, 0xd4, 0x0b, 0xf7, 0x50, 0xf3, 0xd1, 0x6b,
	0xd8, 0x20, 0x63, 0x8a, 0x05, 0x8d, 0x52, 0x46, 0x70, 0x7a, 0xca, 0xb8, 0x30, 0xb7, 0x1d, 0xd0,
	0xde, 0x0d, 0x1e, 0x5f, 0x9b, 0xa0, 0xc2, 0x28, 0x26, 0x90, 0xd0, 0xdb, 0x15, 0xf2, 0x0a, 0x1e,
	0x56, 0xae, 0x81, 0x1a, 0xb0, 0x76, 0x46, 0x67, 0x26, 0x70, 0x40, 0xfb, 0x20, 0x2c, 0x4a, 0xd4,
	0x84, 0xdb, 0x53, 0x9c, 0x4e, 0xa8, 0xb9, 0x25, 0x31, 0xf5, 0xa7, 0x5f, 0xff, 0xf6, 0xd3, 0x36,
	0xdc, 0xdf, 0x00, 0x3e, 0xda, 0x18, 0x09, 0xea, 0xc2, 0x3d, 0x75, 0x91, 0x28, 0x89, 0xa5, 0xe3,
	0x5e, 0xd0, 0xbc, 0x9c, 0xdb, 0x8d, 0xeb, 0x09, 0x45, 0x49, 0xec, 0x86, 0xbb, 0xaa, 0x3e, 0x89,
	0x51, 0x0a, 0x75, 0x4c, 0x57, 0xdb, 0x50, 0x1f, 0xc8, 0x93, 0x75, 0xe1, 0x54, 0x77, 0x60, 0xe9,
	0x1d, 0x1c, 0xdd, 0xe8, 0x70, 0xb5, 0x82, 0xfb, 0x0a, 0x59, 0xf1, 0x3f, 0x9c, 0x2f, 0x2c, 0x70,
	0xb1, 0xb0, 0xc0, 0xbf, 0x85, 0x05, 0xbe, 0x2f, 0x2d, 0xe3, 0x62, 0x69, 0x19, 0x7f, 0x96, 0x96,
	0xf1, 0xe9, 0xc5, 0x30, 0x11, 0xa7, 0x93, 0x81, 0x47, 0x58, 0xe6, 0x13, 0xc6, 0x33, 0xc6, 0xf5,
	0x4f, 0x87, 0xc7, 0x67, 0xfe, 0x17, 0x7f, 0xf5, 0xee, 0x9e, 0xf5, 0x3a, 0xfa, 0xe9, 0x89, 0xd9,
	0x88, 0xf2, 0xc1, 0x8e, 0x7c, 0x61, 0xcf, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xd0, 0xc9,
	0x20, 0xd0, 0x03, 0x00, 0x00,
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
	if m.CreateLocalhost {
		i--
		if m.CreateLocalhost {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
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
	dAtA[i] = 0x22
	if len(m.ClientsMetadata) > 0 {
		for iNdEx := len(m.ClientsMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientsMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClientsConsensus) > 0 {
		for iNdEx := len(m.ClientsConsensus) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientsConsensus[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Clients) > 0 {
		for iNdEx := len(m.Clients) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Clients[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IdentifiedGenesisMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentifiedGenesisMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentifiedGenesisMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientMetadata) > 0 {
		for iNdEx := len(m.ClientMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
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
	if len(m.Clients) > 0 {
		for _, e := range m.Clients {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientsConsensus) > 0 {
		for _, e := range m.ClientsConsensus {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientsMetadata) > 0 {
		for _, e := range m.ClientsMetadata {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.CreateLocalhost {
		n += 2
	}
	return n
}

func (m *GenesisMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *IdentifiedGenesisMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ClientMetadata) > 0 {
		for _, e := range m.ClientMetadata {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Clients", wireType)
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
			m.Clients = append(m.Clients, IdentifiedClientState{})
			if err := m.Clients[len(m.Clients)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientsConsensus", wireType)
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
			m.ClientsConsensus = append(m.ClientsConsensus, ClientConsensusStates{})
			if err := m.ClientsConsensus[len(m.ClientsConsensus)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientsMetadata", wireType)
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
			m.ClientsMetadata = append(m.ClientsMetadata, IdentifiedGenesisMetadata{})
			if err := m.ClientsMetadata[len(m.ClientsMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateLocalhost", wireType)
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
			m.CreateLocalhost = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
func (m *GenesisMetadata) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
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
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
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
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
func (m *IdentifiedGenesisMetadata) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IdentifiedGenesisMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentifiedGenesisMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
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
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientMetadata", wireType)
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
			m.ClientMetadata = append(m.ClientMetadata, GenesisMetadata{})
			if err := m.ClientMetadata[len(m.ClientMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
