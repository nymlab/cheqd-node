// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/resource/v2/fee.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// FeeParams defines the parameters for the `resource` module fixed fee.
type FeeParams struct {
	// Media types define the fixed fee each for the `resource` module.
	Image      types.Coin                             `protobuf:"bytes,1,opt,name=image,proto3" json:"image"`
	Json       types.Coin                             `protobuf:"bytes,2,opt,name=json,proto3" json:"json"`
	Default    types.Coin                             `protobuf:"bytes,3,opt,name=default,proto3" json:"default"`
	BurnFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=burn_factor,json=burnFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"burn_factor"`
}

func (m *FeeParams) Reset()         { *m = FeeParams{} }
func (m *FeeParams) String() string { return proto.CompactTextString(m) }
func (*FeeParams) ProtoMessage()    {}
func (*FeeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_133abe56c2e24f1e, []int{0}
}
func (m *FeeParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeParams.Merge(m, src)
}
func (m *FeeParams) XXX_Size() int {
	return m.Size()
}
func (m *FeeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeParams.DiscardUnknown(m)
}

var xxx_messageInfo_FeeParams proto.InternalMessageInfo

func (m *FeeParams) GetImage() types.Coin {
	if m != nil {
		return m.Image
	}
	return types.Coin{}
}

func (m *FeeParams) GetJson() types.Coin {
	if m != nil {
		return m.Json
	}
	return types.Coin{}
}

func (m *FeeParams) GetDefault() types.Coin {
	if m != nil {
		return m.Default
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*FeeParams)(nil), "cheqd.resource.v2.FeeParams")
}

func init() { proto.RegisterFile("cheqd/resource/v2/fee.proto", fileDescriptor_133abe56c2e24f1e) }

var fileDescriptor_133abe56c2e24f1e = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0x48, 0x2d,
	0x4c, 0xd1, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2f, 0x33, 0xd2, 0x4f, 0x4b,
	0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x04, 0x4b, 0xea, 0xc1, 0x24, 0xf5, 0xca,
	0x8c, 0xa4, 0xe4, 0x92, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xf5, 0x93, 0x12, 0x8b, 0x53, 0xf5, 0xcb,
	0x0c, 0x93, 0x52, 0x4b, 0x12, 0x0d, 0xf5, 0x93, 0xf3, 0x33, 0xf3, 0x20, 0x5a, 0xa4, 0x24, 0x21,
	0xf2, 0xf1, 0x60, 0x9e, 0x3e, 0x84, 0x03, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x87, 0x88, 0x83,
	0x58, 0x10, 0x51, 0xa5, 0x89, 0x4c, 0x5c, 0x9c, 0x6e, 0xa9, 0xa9, 0x01, 0x89, 0x45, 0x89, 0xb9,
	0xc5, 0x42, 0xa6, 0x5c, 0xac, 0x99, 0xb9, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0x92, 0x7a, 0x50, 0x13, 0x40, 0xd6, 0xe9, 0x41, 0xad, 0xd3, 0x73, 0xce, 0xcf, 0xcc, 0x73,
	0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xa2, 0x5a, 0xc8, 0x98, 0x8b, 0x25, 0xab, 0x38, 0x3f,
	0x4f, 0x82, 0x89, 0x38, 0x5d, 0x60, 0xc5, 0x42, 0x96, 0x5c, 0xec, 0x29, 0xa9, 0x69, 0x89, 0xa5,
	0x39, 0x25, 0x12, 0xcc, 0xc4, 0xe9, 0x83, 0xa9, 0x17, 0x8a, 0xe5, 0xe2, 0x4e, 0x2a, 0x2d, 0xca,
	0x8b, 0x4f, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0x92, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x74, 0xb2, 0x01,
	0xa9, 0xb9, 0x75, 0x4f, 0x5e, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x17,
	0x1a, 0x00, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x58, 0xcf, 0x25,
	0x35, 0xf9, 0xd2, 0x16, 0x5d, 0x2e, 0xa8, 0x7d, 0x2e, 0xa9, 0xc9, 0x41, 0x5c, 0x20, 0x03, 0xdd,
	0xc0, 0xe6, 0x39, 0x79, 0xad, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7,
	0x72, 0x0c, 0x51, 0x3a, 0xc8, 0xe6, 0x83, 0x63, 0x0f, 0x4c, 0xea, 0xe6, 0xe5, 0xa7, 0xa4, 0xea,
	0x57, 0x20, 0xa2, 0x12, 0x6c, 0x53, 0x12, 0x1b, 0x38, 0x98, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x7e, 0x2c, 0x51, 0xb6, 0xe9, 0x01, 0x00, 0x00,
}

func (this *FeeParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FeeParams)
	if !ok {
		that2, ok := that.(FeeParams)
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
	if !this.Image.Equal(&that1.Image) {
		return false
	}
	if !this.Json.Equal(&that1.Json) {
		return false
	}
	if !this.Default.Equal(&that1.Default) {
		return false
	}
	if !this.BurnFactor.Equal(that1.BurnFactor) {
		return false
	}
	return true
}
func (m *FeeParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BurnFactor.Size()
		i -= size
		if _, err := m.BurnFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Default.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Json.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Image.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Image.Size()
	n += 1 + l + sovFee(uint64(l))
	l = m.Json.Size()
	n += 1 + l + sovFee(uint64(l))
	l = m.Default.Size()
	n += 1 + l + sovFee(uint64(l))
	l = m.BurnFactor.Size()
	n += 1 + l + sovFee(uint64(l))
	return n
}

func sovFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFee(x uint64) (n int) {
	return sovFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFee
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
			return fmt.Errorf("proto: FeeParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Image.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Json", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Json.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Default", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Default.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFee
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
				return ErrInvalidLengthFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFee
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
func skipFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
					return 0, ErrIntOverflowFee
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
				return 0, ErrInvalidLengthFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFee = fmt.Errorf("proto: unexpected end of group")
)