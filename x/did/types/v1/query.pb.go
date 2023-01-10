// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/did/v1/query.proto

package v1

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryGetDidRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetDidRequest) Reset()         { *m = QueryGetDidRequest{} }
func (m *QueryGetDidRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDidRequest) ProtoMessage()    {}
func (*QueryGetDidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3bd6aa714c99150, []int{0}
}
func (m *QueryGetDidRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDidRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDidRequest.Merge(m, src)
}
func (m *QueryGetDidRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDidRequest proto.InternalMessageInfo

func (m *QueryGetDidRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryGetDidResponse struct {
	Did      *Did      `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *QueryGetDidResponse) Reset()         { *m = QueryGetDidResponse{} }
func (m *QueryGetDidResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDidResponse) ProtoMessage()    {}
func (*QueryGetDidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3bd6aa714c99150, []int{1}
}
func (m *QueryGetDidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDidResponse.Merge(m, src)
}
func (m *QueryGetDidResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDidResponse proto.InternalMessageInfo

func (m *QueryGetDidResponse) GetDid() *Did {
	if m != nil {
		return m.Did
	}
	return nil
}

func (m *QueryGetDidResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetDidRequest)(nil), "cheqdid.cheqdnode.cheqd.v1.QueryGetDidRequest")
	proto.RegisterType((*QueryGetDidResponse)(nil), "cheqdid.cheqdnode.cheqd.v1.QueryGetDidResponse")
}

func init() { proto.RegisterFile("cheqd/did/v1/query.proto", fileDescriptor_a3bd6aa714c99150) }

var fileDescriptor_a3bd6aa714c99150 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4b, 0x33, 0x31,
	0x18, 0xc0, 0x9b, 0x2b, 0xef, 0x8b, 0x46, 0x70, 0x88, 0x22, 0xe5, 0xd0, 0x28, 0xa5, 0x83, 0x08,
	0x26, 0x5c, 0xfd, 0x02, 0x22, 0x05, 0x27, 0x07, 0x3b, 0x38, 0xb8, 0xa5, 0x7d, 0x1e, 0xda, 0x40,
	0x7b, 0xb9, 0x36, 0xb9, 0x62, 0x11, 0x97, 0x8e, 0xe2, 0x20, 0xf8, 0xa5, 0x1c, 0x0b, 0x2e, 0x8e,
	0xd2, 0xf3, 0x83, 0xc8, 0xe5, 0x0e, 0xe1, 0x10, 0x8b, 0xcb, 0x5d, 0xc8, 0xf3, 0xfb, 0x3d, 0xff,
	0x42, 0x1b, 0xfd, 0x21, 0x4e, 0x40, 0x82, 0x06, 0x39, 0x8b, 0xe4, 0x24, 0xc5, 0xe9, 0x5c, 0x24,
	0x53, 0xe3, 0x0c, 0x0b, 0x7d, 0x44, 0x83, 0xf0, 0xff, 0xd8, 0x00, 0x16, 0x27, 0x31, 0x8b, 0xc2,
	0xbd, 0x8a, 0x95, 0x43, 0xde, 0x09, 0x0f, 0x2a, 0xf7, 0xd6, 0x29, 0x87, 0x37, 0x6a, 0x94, 0x62,
	0x19, 0xde, 0x1f, 0x18, 0x33, 0x18, 0xa1, 0x54, 0x89, 0x96, 0x2a, 0x8e, 0x8d, 0x53, 0x4e, 0x9b,
	0xd8, 0x16, 0xd1, 0x66, 0x8b, 0xb2, 0xeb, 0xbc, 0xfe, 0x25, 0xba, 0x8e, 0x86, 0x2e, 0x4e, 0x52,
	0xb4, 0x8e, 0x6d, 0xd3, 0x40, 0x43, 0x83, 0x1c, 0x91, 0xe3, 0xcd, 0x6e, 0xa0, 0xa1, 0xf9, 0x48,
	0xe8, 0x4e, 0x05, 0xb3, 0x89, 0x89, 0x2d, 0xb2, 0x88, 0xd6, 0xa1, 0x04, 0xb7, 0xda, 0x87, 0xe2,
	0xf7, 0xe6, 0x45, 0x6e, 0xe5, 0x2c, 0x3b, 0xa7, 0x1b, 0x63, 0x74, 0x0a, 0x94, 0x53, 0x8d, 0xc0,
	0x7b, 0xad, 0x75, 0xde, 0x55, 0xc9, 0x76, 0xbf, 0xad, 0xf6, 0x13, 0xa1, 0xff, 0x7c, 0x33, 0x6c,
	0x41, 0x68, 0xbd, 0xa3, 0x81, 0x89, 0x75, 0x19, 0x7e, 0x8e, 0x17, 0xca, 0x3f, 0xf3, 0xc5, 0x9c,
	0xcd, 0x70, 0xf1, 0xf6, 0xf9, 0x12, 0xec, 0x32, 0x26, 0x8b, 0x5d, 0x17, 0xfb, 0x97, 0xf7, 0x1a,
	0x1e, 0x2e, 0x3a, 0xaf, 0x2b, 0x4e, 0x96, 0x2b, 0x4e, 0x3e, 0x56, 0x9c, 0x3c, 0x67, 0xbc, 0xb6,
	0xcc, 0x78, 0xed, 0x3d, 0xe3, 0xb5, 0xdb, 0x93, 0x81, 0x76, 0xc3, 0xb4, 0x27, 0xfa, 0x66, 0x5c,
	0x7a, 0xfe, 0x7b, 0x9a, 0xd7, 0x93, 0x77, 0x3e, 0x83, 0x9b, 0x27, 0x68, 0xe5, 0x2c, 0xea, 0xfd,
	0xf7, 0xcf, 0x71, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x7c, 0xd6, 0xcd, 0x1b, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Did(ctx context.Context, in *QueryGetDidRequest, opts ...grpc.CallOption) (*QueryGetDidResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Did(ctx context.Context, in *QueryGetDidRequest, opts ...grpc.CallOption) (*QueryGetDidResponse, error) {
	out := new(QueryGetDidResponse)
	err := c.cc.Invoke(ctx, "/cheqdid.cheqdnode.cheqd.v1.Query/Did", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Did(context.Context, *QueryGetDidRequest) (*QueryGetDidResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Did(ctx context.Context, req *QueryGetDidRequest) (*QueryGetDidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Did not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Did_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Did(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqdid.cheqdnode.cheqd.v1.Query/Did",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Did(ctx, req.(*QueryGetDidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cheqdid.cheqdnode.cheqd.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Did",
			Handler:    _Query_Did_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheqd/did/v1/query.proto",
}

func (m *QueryGetDidRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDidRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDidRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Did != nil {
		{
			size, err := m.Did.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetDidRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Did != nil {
		l = m.Did.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetDidRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDidRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDidRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetDidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetDidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Did == nil {
				m.Did = &Did{}
			}
			if err := m.Did.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)