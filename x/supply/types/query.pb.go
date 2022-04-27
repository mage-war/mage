// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/supply/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
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

// QueryTotalRequest is the request type for Query/Total RPC method
type QueryTotalRequest struct {
	// coin denom to query the circulating supply for
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// divider_exponent is a factor used to power the divider used to convert the
	// supply to the desired representation
	DividerExponent uint64 `protobuf:"varint,2,opt,name=divider_exponent,json=dividerExponent,proto3" json:"divider_exponent,omitempty"`
}

func (m *QueryTotalRequest) Reset()         { *m = QueryTotalRequest{} }
func (m *QueryTotalRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalRequest) ProtoMessage()    {}
func (*QueryTotalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{0}
}
func (m *QueryTotalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalRequest.Merge(m, src)
}
func (m *QueryTotalRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalRequest proto.InternalMessageInfo

// QueryTotalResponse is the response type for the Query/Total RPC method
type QueryTotalResponse struct {
	TotalSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_supply"`
}

func (m *QueryTotalResponse) Reset()         { *m = QueryTotalResponse{} }
func (m *QueryTotalResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTotalResponse) ProtoMessage()    {}
func (*QueryTotalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{1}
}
func (m *QueryTotalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalResponse.Merge(m, src)
}
func (m *QueryTotalResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalResponse proto.InternalMessageInfo

// QueryCirculatingRequest is the request type for the Query/Circulating RPC
// method
type QueryCirculatingRequest struct {
	// coin denom to query the circulating supply for
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// divider_exponent is a factor used to power the divider used to convert the
	// supply to the desired representation
	DividerExponent uint64 `protobuf:"varint,2,opt,name=divider_exponent,json=dividerExponent,proto3" json:"divider_exponent,omitempty"`
}

func (m *QueryCirculatingRequest) Reset()         { *m = QueryCirculatingRequest{} }
func (m *QueryCirculatingRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCirculatingRequest) ProtoMessage()    {}
func (*QueryCirculatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{2}
}
func (m *QueryCirculatingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCirculatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCirculatingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCirculatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCirculatingRequest.Merge(m, src)
}
func (m *QueryCirculatingRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCirculatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCirculatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCirculatingRequest proto.InternalMessageInfo

// QueryCirculatingResponse is the response type for the Query/Circulating RPC
// method
type QueryCirculatingResponse struct {
	CirculatingSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=circulating_supply,json=circulatingSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"circulating_supply"`
}

func (m *QueryCirculatingResponse) Reset()         { *m = QueryCirculatingResponse{} }
func (m *QueryCirculatingResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCirculatingResponse) ProtoMessage()    {}
func (*QueryCirculatingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{3}
}
func (m *QueryCirculatingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCirculatingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCirculatingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCirculatingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCirculatingResponse.Merge(m, src)
}
func (m *QueryCirculatingResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCirculatingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCirculatingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCirculatingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryTotalRequest)(nil), "desmos.supply.v1.QueryTotalRequest")
	proto.RegisterType((*QueryTotalResponse)(nil), "desmos.supply.v1.QueryTotalResponse")
	proto.RegisterType((*QueryCirculatingRequest)(nil), "desmos.supply.v1.QueryCirculatingRequest")
	proto.RegisterType((*QueryCirculatingResponse)(nil), "desmos.supply.v1.QueryCirculatingResponse")
}

func init() { proto.RegisterFile("desmos/supply/v1/query.proto", fileDescriptor_821941a4adac8710) }

var fileDescriptor_821941a4adac8710 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xde, 0x0d, 0x46, 0x74, 0x2a, 0xd8, 0x0e, 0x05, 0xe3, 0x52, 0x76, 0x6b, 0xfc, 0xd1, 0x56,
	0xc8, 0x0e, 0xb1, 0x37, 0x8f, 0x15, 0x0f, 0x3d, 0x36, 0x7a, 0x12, 0x34, 0xcc, 0xee, 0x8e, 0xeb,
	0xe2, 0x66, 0xde, 0x64, 0x67, 0x76, 0x6d, 0x28, 0x5e, 0xf4, 0xe2, 0x51, 0xe8, 0xcd, 0x53, 0xff,
	0x9c, 0x1e, 0x0b, 0x5e, 0xc4, 0x43, 0x91, 0xc4, 0x83, 0x7f, 0x86, 0x64, 0x66, 0xa2, 0x4b, 0x16,
	0x45, 0x90, 0x9e, 0x92, 0xf7, 0xbe, 0x99, 0xef, 0x7d, 0xdf, 0xf7, 0x66, 0xd1, 0x46, 0xc2, 0xe4,
	0x08, 0x24, 0x91, 0xa5, 0x10, 0xf9, 0x84, 0x54, 0x7d, 0x32, 0x2e, 0x59, 0x31, 0x09, 0x45, 0x01,
	0x0a, 0xf0, 0xaa, 0x41, 0x43, 0x83, 0x86, 0x55, 0xdf, 0x5b, 0x4f, 0x21, 0x05, 0x0d, 0x92, 0xf9,
	0x3f, 0x73, 0xce, 0xdb, 0x48, 0x01, 0xd2, 0x9c, 0x11, 0x2a, 0x32, 0x42, 0x39, 0x07, 0x45, 0x55,
	0x06, 0x5c, 0x5a, 0xf4, 0xa6, 0x45, 0x75, 0x15, 0x95, 0x2f, 0x09, 0xe5, 0x76, 0x80, 0xe7, 0x2f,
	0x43, 0x6f, 0x0a, 0x2a, 0x04, 0x2b, 0x7e, 0x5d, 0x8d, 0x61, 0x2e, 0x60, 0x68, 0x26, 0x9a, 0x62,
	0x71, 0xd5, 0x54, 0x24, 0xa2, 0x92, 0x91, 0xaa, 0x1f, 0x31, 0x45, 0xfb, 0x24, 0x86, 0x8c, 0x1b,
	0xbc, 0xfb, 0x02, 0xad, 0x1d, 0xcc, 0xad, 0x3c, 0x05, 0x45, 0xf3, 0x01, 0x1b, 0x97, 0x4c, 0x2a,
	0xbc, 0x8e, 0xda, 0x09, 0xe3, 0x30, 0xea, 0xb8, 0x9b, 0xee, 0xf6, 0xd5, 0x81, 0x29, 0xf0, 0x0e,
	0x5a, 0x4d, 0xb2, 0x2a, 0x4b, 0x58, 0x31, 0x64, 0x87, 0x02, 0x38, 0xe3, 0xaa, 0xd3, 0xda, 0x74,
	0xb7, 0x2f, 0x0d, 0xae, 0xdb, 0xfe, 0x63, 0xdb, 0x7e, 0x78, 0xe5, 0xc3, 0x49, 0xe0, 0xfc, 0x38,
	0x09, 0x9c, 0xee, 0x18, 0xe1, 0x3a, 0xbf, 0x14, 0xc0, 0x25, 0xc3, 0x07, 0xe8, 0x9a, 0x9a, 0x37,
	0x86, 0x26, 0x32, 0x33, 0x67, 0x2f, 0x3c, 0x3d, 0x0f, 0x9c, 0xaf, 0xe7, 0xc1, 0xbd, 0x34, 0x53,
	0xaf, 0xca, 0x28, 0x8c, 0x61, 0x64, 0xcd, 0xd8, 0x9f, 0x9e, 0x4c, 0x5e, 0x13, 0x35, 0x11, 0x4c,
	0x86, 0xfb, 0x5c, 0x0d, 0x56, 0x34, 0xc7, 0x13, 0x4d, 0x51, 0x1b, 0x99, 0xa0, 0x1b, 0x7a, 0xe4,
	0xa3, 0xac, 0x88, 0xcb, 0x9c, 0xaa, 0x8c, 0xa7, 0x17, 0x60, 0xec, 0xbd, 0x8b, 0x3a, 0xcd, 0x31,
	0xd6, 0xdf, 0x73, 0x84, 0xe3, 0xdf, 0xed, 0xff, 0x73, 0xb9, 0x56, 0x63, 0x5a, 0xf6, 0xfa, 0xe0,
	0x53, 0x0b, 0xb5, 0xb5, 0x0a, 0x7c, 0x84, 0xda, 0x3a, 0x63, 0x7c, 0x3b, 0x5c, 0x7e, 0x8e, 0x61,
	0x63, 0xc3, 0xde, 0x9d, 0xbf, 0x1f, 0x32, 0x36, 0xba, 0x5b, 0xef, 0x3e, 0x7f, 0x3f, 0x6e, 0xdd,
	0xc2, 0x01, 0x69, 0xbc, 0x7f, 0x1d, 0x3d, 0x39, 0xd2, 0x01, 0xbe, 0xc5, 0xc7, 0x2e, 0x5a, 0xa9,
	0xe5, 0x80, 0x77, 0xfe, 0x40, 0xdf, 0x5c, 0x89, 0x77, 0xff, 0x5f, 0x8e, 0x5a, 0x3d, 0x3d, 0xad,
	0x67, 0x0b, 0xdf, 0x6d, 0xea, 0xa9, 0x85, 0xb4, 0x50, 0xb5, 0xb7, 0x7f, 0x3a, 0xf5, 0xdd, 0xb3,
	0xa9, 0xef, 0x7e, 0x9b, 0xfa, 0xee, 0xc7, 0x99, 0xef, 0x9c, 0xcd, 0x7c, 0xe7, 0xcb, 0xcc, 0x77,
	0x9e, 0x91, 0x5a, 0xf6, 0x86, 0xaa, 0x97, 0xd3, 0x48, 0x2e, 0x68, 0xab, 0x5d, 0x72, 0xb8, 0xe0,
	0xd6, 0x8b, 0x88, 0x2e, 0xeb, 0xaf, 0x65, 0xf7, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xb5,
	0x0f, 0x03, 0x09, 0x04, 0x00, 0x00,
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
	// Total queries the total supply of the given denom
	Total(ctx context.Context, in *QueryTotalRequest, opts ...grpc.CallOption) (*QueryTotalResponse, error)
	// Circulating queries the amount of tokens circulating in the market of the
	// given denom
	Circulating(ctx context.Context, in *QueryCirculatingRequest, opts ...grpc.CallOption) (*QueryCirculatingResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Total(ctx context.Context, in *QueryTotalRequest, opts ...grpc.CallOption) (*QueryTotalResponse, error) {
	out := new(QueryTotalResponse)
	err := c.cc.Invoke(ctx, "/desmos.supply.v1.Query/Total", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Circulating(ctx context.Context, in *QueryCirculatingRequest, opts ...grpc.CallOption) (*QueryCirculatingResponse, error) {
	out := new(QueryCirculatingResponse)
	err := c.cc.Invoke(ctx, "/desmos.supply.v1.Query/Circulating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Total queries the total supply of the given denom
	Total(context.Context, *QueryTotalRequest) (*QueryTotalResponse, error)
	// Circulating queries the amount of tokens circulating in the market of the
	// given denom
	Circulating(context.Context, *QueryCirculatingRequest) (*QueryCirculatingResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Total(ctx context.Context, req *QueryTotalRequest) (*QueryTotalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Total not implemented")
}
func (*UnimplementedQueryServer) Circulating(ctx context.Context, req *QueryCirculatingRequest) (*QueryCirculatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Circulating not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Total_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Total(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.supply.v1.Query/Total",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Total(ctx, req.(*QueryTotalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Circulating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCirculatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Circulating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.supply.v1.Query/Circulating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Circulating(ctx, req.(*QueryCirculatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.supply.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Total",
			Handler:    _Query_Total_Handler,
		},
		{
			MethodName: "Circulating",
			Handler:    _Query_Circulating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/supply/v1/query.proto",
}

func (m *QueryTotalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DividerExponent != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.DividerExponent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTotalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalSupply.Size()
		i -= size
		if _, err := m.TotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCirculatingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCirculatingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCirculatingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DividerExponent != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.DividerExponent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCirculatingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCirculatingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCirculatingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CirculatingSupply.Size()
		i -= size
		if _, err := m.CirculatingSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *QueryTotalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.DividerExponent != 0 {
		n += 1 + sovQuery(uint64(m.DividerExponent))
	}
	return n
}

func (m *QueryTotalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalSupply.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryCirculatingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.DividerExponent != 0 {
		n += 1 + sovQuery(uint64(m.DividerExponent))
	}
	return n
}

func (m *QueryCirculatingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CirculatingSupply.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryTotalRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DividerExponent", wireType)
			}
			m.DividerExponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DividerExponent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryTotalResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
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
			if err := m.TotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCirculatingRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCirculatingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCirculatingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DividerExponent", wireType)
			}
			m.DividerExponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DividerExponent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *QueryCirculatingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCirculatingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCirculatingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CirculatingSupply", wireType)
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
			if err := m.CirculatingSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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