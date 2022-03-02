// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/relationships/v1/models.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Relationship is the struct of a relationship.
// It represent the concept of "follow" of traditional social networks.
type Relationship struct {
	// Creator represents the creator of the relationship
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	// Counterparty represents the other user involved in the relationship
	Counterparty string `protobuf:"bytes,2,opt,name=counterparty,proto3" json:"counterparty,omitempty" yaml:"recipient"`
	// SubspaceID represents the id of the subspace for which the relationship is
	// valid
	SubspaceID uint64 `protobuf:"varint,3,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty" yaml:"subspace"`
}

func (m *Relationship) Reset()         { *m = Relationship{} }
func (m *Relationship) String() string { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()    {}
func (*Relationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_9096e6e86fd356f7, []int{0}
}
func (m *Relationship) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Relationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Relationship.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Relationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relationship.Merge(m, src)
}
func (m *Relationship) XXX_Size() int {
	return m.Size()
}
func (m *Relationship) XXX_DiscardUnknown() {
	xxx_messageInfo_Relationship.DiscardUnknown(m)
}

var xxx_messageInfo_Relationship proto.InternalMessageInfo

func (m *Relationship) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Relationship) GetCounterparty() string {
	if m != nil {
		return m.Counterparty
	}
	return ""
}

func (m *Relationship) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

// UserBlock represents the fact that the Blocker has blocked the given Blocked
// user.
type UserBlock struct {
	// Blocker represents the address of the user blocking another one
	Blocker string `protobuf:"bytes,1,opt,name=blocker,proto3" json:"blocker,omitempty" yaml:"blocker"`
	// Blocked represents the address of the blocked user
	Blocked string `protobuf:"bytes,2,opt,name=blocked,proto3" json:"blocked,omitempty" yaml:"blocked"`
	// Reason represents the optional reason the user has been blocked for.
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty" yaml:"reason"`
	// SubspaceID represents the ID of the subspace inside which the user should
	// be blocked
	SubspaceID uint64 `protobuf:"varint,4,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty" yaml:"subspace"`
}

func (m *UserBlock) Reset()         { *m = UserBlock{} }
func (m *UserBlock) String() string { return proto.CompactTextString(m) }
func (*UserBlock) ProtoMessage()    {}
func (*UserBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_9096e6e86fd356f7, []int{1}
}
func (m *UserBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlock.Merge(m, src)
}
func (m *UserBlock) XXX_Size() int {
	return m.Size()
}
func (m *UserBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlock.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlock proto.InternalMessageInfo

func (m *UserBlock) GetBlocker() string {
	if m != nil {
		return m.Blocker
	}
	return ""
}

func (m *UserBlock) GetBlocked() string {
	if m != nil {
		return m.Blocked
	}
	return ""
}

func (m *UserBlock) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *UserBlock) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func init() {
	proto.RegisterType((*Relationship)(nil), "desmos.relationships.v1.Relationship")
	proto.RegisterType((*UserBlock)(nil), "desmos.relationships.v1.UserBlock")
}

func init() {
	proto.RegisterFile("desmos/relationships/v1/models.proto", fileDescriptor_9096e6e86fd356f7)
}

var fileDescriptor_9096e6e86fd356f7 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0x6b, 0xa8, 0x0a, 0x35, 0xe5, 0x5f, 0x54, 0x89, 0xd2, 0x21, 0x2e, 0x16, 0x43, 0x91,
	0x20, 0x56, 0x41, 0x42, 0xa8, 0x63, 0xc4, 0xd2, 0x0d, 0x05, 0xb1, 0xb0, 0x54, 0x4e, 0x62, 0xd2,
	0x88, 0x24, 0x8e, 0x6c, 0xa7, 0x22, 0x6f, 0xc1, 0xc8, 0xd8, 0x87, 0x61, 0x60, 0xec, 0xc8, 0x14,
	0x55, 0xe9, 0xc2, 0x9c, 0x27, 0x40, 0x89, 0x93, 0x7b, 0x6f, 0x7b, 0xef, 0x76, 0xb7, 0xcf, 0xe7,
	0xfb, 0x9d, 0xe3, 0xf3, 0x49, 0x07, 0xbe, 0xf4, 0x99, 0x8c, 0xb9, 0x24, 0x82, 0x45, 0x54, 0x85,
	0x3c, 0x91, 0x9b, 0x30, 0x95, 0x64, 0xbb, 0x20, 0x31, 0xf7, 0x59, 0x24, 0xad, 0x54, 0x70, 0xc5,
	0x8d, 0x67, 0x9a, 0xb2, 0x4e, 0x28, 0x6b, 0xbb, 0x98, 0x8e, 0x03, 0x1e, 0xf0, 0x86, 0x21, 0xb5,
	0xd2, 0xf8, 0xf4, 0x79, 0xc0, 0x79, 0x10, 0x31, 0xd2, 0xbc, 0xdc, 0xec, 0x1b, 0xa1, 0x49, 0xde,
	0x5a, 0xe8, 0xdc, 0x52, 0x61, 0xcc, 0xa4, 0xa2, 0x71, 0xda, 0xf5, 0x7a, 0xbc, 0xfe, 0x6a, 0xad,
	0x87, 0xea, 0x87, 0xb6, 0xf0, 0x6f, 0x00, 0x47, 0xce, 0x95, 0x0d, 0x8c, 0xd7, 0xf0, 0x9e, 0x27,
	0x18, 0x55, 0x5c, 0x4c, 0xc0, 0x0c, 0xcc, 0x87, 0xb6, 0x51, 0x15, 0xe8, 0x51, 0x4e, 0xe3, 0x68,
	0x89, 0x5b, 0x03, 0x3b, 0x1d, 0x62, 0x7c, 0x80, 0x23, 0x8f, 0x67, 0x89, 0x62, 0x22, 0xa5, 0x42,
	0xe5, 0x93, 0x3b, 0x4d, 0xcb, 0xb8, 0x2a, 0xd0, 0x13, 0xdd, 0x22, 0x98, 0x17, 0xa6, 0x21, 0x4b,
	0x14, 0x76, 0x4e, 0x48, 0xc3, 0x86, 0x0f, 0x64, 0xe6, 0xca, 0x94, 0x7a, 0x6c, 0x1d, 0xfa, 0x93,
	0xbb, 0x33, 0x30, 0xef, 0xdb, 0x2f, 0xca, 0x02, 0xc1, 0xcf, 0x6d, 0x79, 0xf5, 0xb1, 0x2a, 0xd0,
	0x63, 0x3d, 0xa6, 0x43, 0xb1, 0x03, 0x3b, 0xb9, 0xf2, 0x97, 0xf7, 0x7f, 0xed, 0x10, 0xf8, 0xb7,
	0x43, 0x00, 0x1f, 0x00, 0x1c, 0x7e, 0x91, 0x4c, 0xd8, 0x11, 0xf7, 0xbe, 0xd7, 0x19, 0xdc, 0x5a,
	0xb0, 0x1b, 0x32, 0xb4, 0x06, 0x76, 0x3a, 0xe4, 0x92, 0xf6, 0xdb, 0xf5, 0xaf, 0xd1, 0xfe, 0x05,
	0xed, 0x1b, 0xaf, 0xe0, 0x40, 0x30, 0x2a, 0x79, 0xd2, 0xac, 0x3c, 0xb4, 0x9f, 0x56, 0x05, 0x7a,
	0xd8, 0x65, 0xad, 0xeb, 0xd8, 0x69, 0x81, 0xf3, 0x88, 0xfd, 0x5b, 0x45, 0xb4, 0x3f, 0xfd, 0x29,
	0x4d, 0xb0, 0x2f, 0x4d, 0x70, 0x28, 0x4d, 0xf0, 0xf3, 0x68, 0xf6, 0xf6, 0x47, 0xb3, 0xf7, 0xf7,
	0x68, 0xf6, 0xbe, 0xbe, 0x0f, 0x42, 0xb5, 0xc9, 0x5c, 0xcb, 0xe3, 0x31, 0xd1, 0x47, 0xf5, 0x26,
	0xa2, 0xae, 0x6c, 0x35, 0xd9, 0xbe, 0x25, 0x3f, 0xce, 0x6e, 0x51, 0xe5, 0x29, 0x93, 0xee, 0xa0,
	0x39, 0x81, 0x77, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe1, 0xdf, 0x14, 0xb0, 0x02, 0x00,
	0x00,
}

func (this *Relationship) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Relationship)
	if !ok {
		that2, ok := that.(Relationship)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if this.Counterparty != that1.Counterparty {
		return false
	}
	if this.SubspaceID != that1.SubspaceID {
		return false
	}
	return true
}
func (this *UserBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserBlock)
	if !ok {
		that2, ok := that.(UserBlock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Blocker != that1.Blocker {
		return false
	}
	if this.Blocked != that1.Blocked {
		return false
	}
	if this.Reason != that1.Reason {
		return false
	}
	if this.SubspaceID != that1.SubspaceID {
		return false
	}
	return true
}
func (m *Relationship) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Relationship) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Relationship) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubspaceID != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Counterparty) > 0 {
		i -= len(m.Counterparty)
		copy(dAtA[i:], m.Counterparty)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Counterparty)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubspaceID != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Blocked) > 0 {
		i -= len(m.Blocked)
		copy(dAtA[i:], m.Blocked)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Blocked)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Blocker) > 0 {
		i -= len(m.Blocker)
		copy(dAtA[i:], m.Blocker)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Blocker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModels(dAtA []byte, offset int, v uint64) int {
	offset -= sovModels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Relationship) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Counterparty)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.SubspaceID != 0 {
		n += 1 + sovModels(uint64(m.SubspaceID))
	}
	return n
}

func (m *UserBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Blocker)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Blocked)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.SubspaceID != 0 {
		n += 1 + sovModels(uint64(m.SubspaceID))
	}
	return n
}

func sovModels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModels(x uint64) (n int) {
	return sovModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Relationship) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: Relationship: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Relationship: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counterparty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Counterparty = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModels
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
func (m *UserBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: UserBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocked = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModels
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
func skipModels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
				return 0, ErrInvalidLengthModels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModels = fmt.Errorf("proto: unexpected end of group")
)