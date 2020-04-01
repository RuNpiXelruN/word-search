// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protowordsearch.proto

package wordsearch

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ListWordsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWordsRequest) Reset()         { *m = ListWordsRequest{} }
func (m *ListWordsRequest) String() string { return proto.CompactTextString(m) }
func (*ListWordsRequest) ProtoMessage()    {}
func (*ListWordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c99cceefb9f3eb54, []int{0}
}

func (m *ListWordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWordsRequest.Unmarshal(m, b)
}
func (m *ListWordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWordsRequest.Marshal(b, m, deterministic)
}
func (m *ListWordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWordsRequest.Merge(m, src)
}
func (m *ListWordsRequest) XXX_Size() int {
	return xxx_messageInfo_ListWordsRequest.Size(m)
}
func (m *ListWordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWordsRequest proto.InternalMessageInfo

type ListWordsResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWordsResponse) Reset()         { *m = ListWordsResponse{} }
func (m *ListWordsResponse) String() string { return proto.CompactTextString(m) }
func (*ListWordsResponse) ProtoMessage()    {}
func (*ListWordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c99cceefb9f3eb54, []int{1}
}

func (m *ListWordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWordsResponse.Unmarshal(m, b)
}
func (m *ListWordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWordsResponse.Marshal(b, m, deterministic)
}
func (m *ListWordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWordsResponse.Merge(m, src)
}
func (m *ListWordsResponse) XXX_Size() int {
	return xxx_messageInfo_ListWordsResponse.Size(m)
}
func (m *ListWordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWordsResponse proto.InternalMessageInfo

func (m *ListWordsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ListWordsRequest)(nil), "wordsearch.ListWordsRequest")
	proto.RegisterType((*ListWordsResponse)(nil), "wordsearch.ListWordsResponse")
}

func init() {
	proto.RegisterFile("protowordsearch.proto", fileDescriptor_c99cceefb9f3eb54)
}

var fileDescriptor_c99cceefb9f3eb54 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xcf, 0x2f, 0x4a, 0x29, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0xd0, 0x03, 0xf3, 0x85, 0xb8,
	0x10, 0x22, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89,
	0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x95, 0x4a, 0x42, 0x5c, 0x02,
	0x3e, 0x99, 0xc5, 0x25, 0xe1, 0x20, 0xf5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0xba,
	0x5c, 0x82, 0x48, 0x62, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9,
	0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x51, 0x01,
	0x17, 0x17, 0x48, 0x69, 0x30, 0xd8, 0x3a, 0xa1, 0x24, 0x2e, 0x7e, 0x90, 0x66, 0x08, 0x2f, 0x24,
	0xb5, 0x28, 0xb7, 0x58, 0x48, 0x46, 0x0f, 0xc9, 0x81, 0xe8, 0xb6, 0x49, 0xc9, 0xe2, 0x90, 0x85,
	0xd8, 0xab, 0x24, 0xd4, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x1e, 0x21, 0x2e, 0xb0, 0x07, 0xc0, 0x4a,
	0x93, 0xd8, 0xc0, 0x6e, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xa0, 0x30, 0xe9, 0xfe,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WordSearchClient is the client API for WordSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WordSearchClient interface {
	ListSearchTerms(ctx context.Context, in *ListWordsRequest, opts ...grpc.CallOption) (*ListWordsResponse, error)
}

type wordSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewWordSearchClient(cc grpc.ClientConnInterface) WordSearchClient {
	return &wordSearchClient{cc}
}

func (c *wordSearchClient) ListSearchTerms(ctx context.Context, in *ListWordsRequest, opts ...grpc.CallOption) (*ListWordsResponse, error) {
	out := new(ListWordsResponse)
	err := c.cc.Invoke(ctx, "/wordsearch.WordSearch/ListSearchTerms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordSearchServer is the server API for WordSearch service.
type WordSearchServer interface {
	ListSearchTerms(context.Context, *ListWordsRequest) (*ListWordsResponse, error)
}

// UnimplementedWordSearchServer can be embedded to have forward compatible implementations.
type UnimplementedWordSearchServer struct {
}

func (*UnimplementedWordSearchServer) ListSearchTerms(ctx context.Context, req *ListWordsRequest) (*ListWordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSearchTerms not implemented")
}

func RegisterWordSearchServer(s *grpc.Server, srv WordSearchServer) {
	s.RegisterService(&_WordSearch_serviceDesc, srv)
}

func _WordSearch_ListSearchTerms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordSearchServer).ListSearchTerms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wordsearch.WordSearch/ListSearchTerms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordSearchServer).ListSearchTerms(ctx, req.(*ListWordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordSearch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wordsearch.WordSearch",
	HandlerType: (*WordSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSearchTerms",
			Handler:    _WordSearch_ListSearchTerms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protowordsearch.proto",
}
