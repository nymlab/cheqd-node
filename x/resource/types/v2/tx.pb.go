// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/resource/v2/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cheqd/cheqd-node/x/did/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// MsgCreateResource defines the Msg/CreateResource request type.
// It describes the parameters of a request for creating a resource.
type MsgCreateResource struct {
	// Payload containing the resource to be created.
	Payload *MsgCreateResourcePayload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// Signatures of the corresponding DID Document's controller(s).
	Signatures []*types.SignInfo `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (m *MsgCreateResource) Reset()         { *m = MsgCreateResource{} }
func (m *MsgCreateResource) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResource) ProtoMessage()    {}
func (*MsgCreateResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d13b428c5ed4ca4, []int{0}
}
func (m *MsgCreateResource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResource.Merge(m, src)
}
func (m *MsgCreateResource) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResource.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResource proto.InternalMessageInfo

func (m *MsgCreateResource) GetPayload() *MsgCreateResourcePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *MsgCreateResource) GetSignatures() []*types.SignInfo {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// MsgCreateResourcePayload defines the structure of the payload for creating a resource.
//
// If a resource with the given id does not exist already,
// it will be created. The resource will be created in the resource collection.
//
// If a resource with the given id, collection_id already exists, an error code 2200 will be returned.
//
// A new version of the resource in an existing collection will be created,
// if a resource in that collection with the same name, resource_type and empty next_version_id exists.
//
// An update operation is not possible, because the resource is immutable by design.
type MsgCreateResourcePayload struct {
	// data is a byte-representation of the actual Data the user wants to store.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// collection_id is an identifier of the DidDocument the resource belongs to.
	// Format: <unique-identifier>
	//
	// Examples:
	// - c82f2b02-bdab-4dd7-b833-3e143745d612
	// - wGHEXrZvJxR8vw5P3UWH1j
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"resourceCollectionId"`
	// id is a unique id of the resource.
	// Format: <uuid>
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"resourceId"`
	// name is a human-readable name of the resource.
	// Format: <string>
	//
	// Does not change between different versions.
	// Example: PassportSchema, EducationTrustRegistry
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"resourceName"`
	// version is a version of the resource.
	// Format: <string>
	// Stored as a string. OPTIONAL.
	//
	// Example: 1.0.0, v2.1.0
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"resourceVersion"`
	// resource_type is a type of the resource.
	// Format: <string>
	//
	// This is NOT the same as the resource's media type.
	// Example: AnonCredsSchema, StatusList2021
	ResourceType string `protobuf:"bytes,6,opt,name=resource_type,json=resourceType,proto3" json:"resourceType"`
	// also_known_as is a list of URIs that can be used to get the resource.
	AlsoKnownAs []*AlternativeUri `protobuf:"bytes,7,rep,name=also_known_as,json=alsoKnownAs,proto3" json:"resourceAlternativeUri"`
}

func (m *MsgCreateResourcePayload) Reset()         { *m = MsgCreateResourcePayload{} }
func (m *MsgCreateResourcePayload) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResourcePayload) ProtoMessage()    {}
func (*MsgCreateResourcePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d13b428c5ed4ca4, []int{1}
}
func (m *MsgCreateResourcePayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResourcePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResourcePayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResourcePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResourcePayload.Merge(m, src)
}
func (m *MsgCreateResourcePayload) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResourcePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResourcePayload.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResourcePayload proto.InternalMessageInfo

func (m *MsgCreateResourcePayload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MsgCreateResourcePayload) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *MsgCreateResourcePayload) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgCreateResourcePayload) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreateResourcePayload) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *MsgCreateResourcePayload) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *MsgCreateResourcePayload) GetAlsoKnownAs() []*AlternativeUri {
	if m != nil {
		return m.AlsoKnownAs
	}
	return nil
}

type MsgCreateResourceResponse struct {
	// Return the created resource metadata.
	Resource *Metadata `protobuf:"bytes,1,opt,name=resource,proto3" json:"linkedResourceMetadata"`
}

func (m *MsgCreateResourceResponse) Reset()         { *m = MsgCreateResourceResponse{} }
func (m *MsgCreateResourceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateResourceResponse) ProtoMessage()    {}
func (*MsgCreateResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d13b428c5ed4ca4, []int{2}
}
func (m *MsgCreateResourceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateResourceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateResourceResponse.Merge(m, src)
}
func (m *MsgCreateResourceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateResourceResponse proto.InternalMessageInfo

func (m *MsgCreateResourceResponse) GetResource() *Metadata {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreateResource)(nil), "cheqd.resource.v2.MsgCreateResource")
	proto.RegisterType((*MsgCreateResourcePayload)(nil), "cheqd.resource.v2.MsgCreateResourcePayload")
	proto.RegisterType((*MsgCreateResourceResponse)(nil), "cheqd.resource.v2.MsgCreateResourceResponse")
}

func init() { proto.RegisterFile("cheqd/resource/v2/tx.proto", fileDescriptor_1d13b428c5ed4ca4) }

var fileDescriptor_1d13b428c5ed4ca4 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0x9b, 0xb6, 0xac, 0xe0, 0xb6, 0x83, 0x59, 0x63, 0x98, 0x22, 0xa5, 0xa5, 0xda, 0xa1,
	0x12, 0x23, 0x11, 0x45, 0x70, 0xe3, 0xb0, 0x4e, 0x20, 0x55, 0x68, 0x08, 0x79, 0xc0, 0x81, 0x4b,
	0xe5, 0xc5, 0x26, 0xb3, 0x9a, 0xda, 0x21, 0x76, 0xcb, 0x7a, 0xe7, 0x07, 0x20, 0x7e, 0xd5, 0x8e,
	0x3b, 0x72, 0xaa, 0x50, 0x7b, 0xeb, 0xaf, 0x40, 0x71, 0xea, 0xd2, 0x91, 0x21, 0xed, 0x12, 0xd9,
	0x79, 0x9f, 0xf7, 0xf5, 0x97, 0x7c, 0x9f, 0x41, 0x23, 0x38, 0x63, 0x5f, 0xa9, 0x9f, 0x30, 0x25,
	0xc7, 0x49, 0xc0, 0xfc, 0x49, 0xd7, 0xd7, 0xe7, 0x5e, 0x9c, 0x48, 0x2d, 0xe1, 0x8e, 0xd1, 0x3c,
	0xab, 0x79, 0x93, 0x6e, 0xe3, 0x7e, 0x86, 0x53, 0x4e, 0x37, 0xc9, 0x46, 0x2b, 0x9f, 0xb2, 0x76,
	0x65, 0xc4, 0x6e, 0x28, 0x43, 0x69, 0x96, 0x7e, 0xba, 0xca, 0xde, 0xb6, 0x7f, 0x3a, 0x60, 0xe7,
	0x58, 0x85, 0x47, 0x09, 0x23, 0x9a, 0xe1, 0x95, 0x03, 0xbe, 0x06, 0x95, 0x98, 0x4c, 0x23, 0x49,
	0x28, 0x72, 0x5a, 0x4e, 0xa7, 0xda, 0x7d, 0xe2, 0xe5, 0x2a, 0xf1, 0x72, 0xb6, 0xf7, 0x99, 0x05,
	0x5b, 0x2f, 0x7c, 0x09, 0x80, 0xe2, 0xa1, 0x20, 0x7a, 0x9c, 0x30, 0x85, 0x8a, 0xad, 0x52, 0xa7,
	0xda, 0xdd, 0x5b, 0x25, 0x51, 0x4e, 0xd3, 0x90, 0x13, 0x1e, 0x8a, 0xbe, 0xf8, 0x22, 0xf1, 0x06,
	0xd9, 0xfe, 0x5e, 0x02, 0xe8, 0x7f, 0xe9, 0x10, 0x82, 0x32, 0x25, 0x9a, 0x98, 0xc2, 0x6a, 0xd8,
	0xac, 0xe1, 0x2b, 0x50, 0x0f, 0x64, 0x14, 0xb1, 0x40, 0x73, 0x29, 0x06, 0x9c, 0xa2, 0x62, 0xcb,
	0xe9, 0xdc, 0xe9, 0xa1, 0xe5, 0xac, 0xb9, 0x6b, 0x4b, 0x3e, 0x5a, 0x03, 0x7d, 0x8a, 0x6b, 0xc1,
	0xc6, 0x0e, 0xba, 0xa0, 0xc8, 0x29, 0x2a, 0x19, 0xcf, 0xf6, 0x72, 0xd6, 0x04, 0xd6, 0xd3, 0xa7,
	0xb8, 0xc8, 0x29, 0xdc, 0x07, 0x65, 0x41, 0x46, 0x0c, 0x95, 0x0d, 0x71, 0x6f, 0x39, 0x6b, 0xd6,
	0x2c, 0xf1, 0x8e, 0x8c, 0x18, 0x36, 0x2a, 0x7c, 0x06, 0x2a, 0x13, 0x96, 0x28, 0x2e, 0x05, 0xba,
	0x65, 0xc0, 0x07, 0x17, 0xb3, 0xa6, 0xb3, 0x9c, 0x35, 0xef, 0x5a, 0xf8, 0x53, 0x26, 0x63, 0xcb,
	0xc1, 0x17, 0xa0, 0x6e, 0xb5, 0x81, 0x9e, 0xc6, 0x0c, 0x6d, 0xe5, 0x4f, 0xf8, 0x30, 0x8d, 0x19,
	0xbe, 0xb2, 0x83, 0x0c, 0xd4, 0x49, 0xa4, 0xe4, 0x60, 0x28, 0xe4, 0x37, 0x31, 0x20, 0x0a, 0x55,
	0xcc, 0xaf, 0x7d, 0x7c, 0x4d, 0x93, 0x0e, 0x23, 0xcd, 0x12, 0x41, 0x34, 0x9f, 0xb0, 0x8f, 0x09,
	0xef, 0xb9, 0xab, 0x92, 0xf6, 0x2c, 0x73, 0x55, 0xc7, 0xd5, 0x34, 0xf7, 0x6d, 0x1a, 0x7b, 0xa8,
	0xda, 0x31, 0x78, 0x98, 0xeb, 0x02, 0x66, 0x2a, 0x96, 0x42, 0x31, 0x78, 0x02, 0x6e, 0xdb, 0x8c,
	0xd5, 0x8c, 0x3c, 0xba, 0x6e, 0x46, 0x98, 0x26, 0x69, 0x87, 0x7a, 0x8d, 0xf4, 0xd0, 0x88, 0x8b,
	0x21, 0xa3, 0x36, 0xca, 0x6a, 0x78, 0x1d, 0xd4, 0x1d, 0x82, 0xd2, 0xb1, 0x0a, 0x21, 0x05, 0xdb,
	0xff, 0x0c, 0xe4, 0xfe, 0x4d, 0xe6, 0xaf, 0x71, 0x70, 0x13, 0xca, 0x7e, 0x41, 0xef, 0xcd, 0xc5,
	0xdc, 0x75, 0x2e, 0xe7, 0xae, 0xf3, 0x7b, 0xee, 0x3a, 0x3f, 0x16, 0x6e, 0xe1, 0x72, 0xe1, 0x16,
	0x7e, 0x2d, 0xdc, 0xc2, 0xe7, 0x83, 0x90, 0xeb, 0xb3, 0xf1, 0xa9, 0x17, 0xc8, 0x91, 0x9f, 0xdd,
	0x2b, 0xf3, 0x7c, 0x2a, 0x24, 0x65, 0xfe, 0xf9, 0xdf, 0x4b, 0x96, 0xb6, 0x4c, 0x9d, 0x6e, 0x99,
	0x9b, 0xf4, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x6d, 0x3e, 0x2f, 0xc9, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateResource defines a method for creating a resource.
	CreateResource(ctx context.Context, in *MsgCreateResource, opts ...grpc.CallOption) (*MsgCreateResourceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateResource(ctx context.Context, in *MsgCreateResource, opts ...grpc.CallOption) (*MsgCreateResourceResponse, error) {
	out := new(MsgCreateResourceResponse)
	err := c.cc.Invoke(ctx, "/cheqd.resource.v2.Msg/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateResource defines a method for creating a resource.
	CreateResource(context.Context, *MsgCreateResource) (*MsgCreateResourceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateResource(ctx context.Context, req *MsgCreateResource) (*MsgCreateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateResource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqd.resource.v2.Msg/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateResource(ctx, req.(*MsgCreateResource))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cheqd.resource.v2.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResource",
			Handler:    _Msg_CreateResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheqd/resource/v2/tx.proto",
}

func (m *MsgCreateResource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for iNdEx := len(m.Signatures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signatures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Payload != nil {
		{
			size, err := m.Payload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateResourcePayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResourcePayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResourcePayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AlsoKnownAs) > 0 {
		for iNdEx := len(m.AlsoKnownAs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AlsoKnownAs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ResourceType) > 0 {
		i -= len(m.ResourceType)
		copy(dAtA[i:], m.ResourceType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ResourceType)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CollectionId) > 0 {
		i -= len(m.CollectionId)
		copy(dAtA[i:], m.CollectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CollectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateResourceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateResourceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateResourceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Resource != nil {
		{
			size, err := m.Resource.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateResource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, e := range m.Signatures {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateResourcePayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CollectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ResourceType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.AlsoKnownAs) > 0 {
		for _, e := range m.AlsoKnownAs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateResourceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Resource != nil {
		l = m.Resource.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateResource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &MsgCreateResourcePayload{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, &types.SignInfo{})
			if err := m.Signatures[len(m.Signatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateResourcePayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateResourcePayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResourcePayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlsoKnownAs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AlsoKnownAs = append(m.AlsoKnownAs, &AlternativeUri{})
			if err := m.AlsoKnownAs[len(m.AlsoKnownAs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateResourceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateResourceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateResourceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resource == nil {
				m.Resource = &Metadata{}
			}
			if err := m.Resource.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)