// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/newworsearch.proto

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

type SearchItem struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SearchCount          int64    `protobuf:"varint,2,opt,name=search_count,json=searchCount,proto3" json:"search_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchItem) Reset()         { *m = SearchItem{} }
func (m *SearchItem) String() string { return proto.CompactTextString(m) }
func (*SearchItem) ProtoMessage()    {}
func (*SearchItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cf4e25ab144b61, []int{0}
}

func (m *SearchItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchItem.Unmarshal(m, b)
}
func (m *SearchItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchItem.Marshal(b, m, deterministic)
}
func (m *SearchItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchItem.Merge(m, src)
}
func (m *SearchItem) XXX_Size() int {
	return xxx_messageInfo_SearchItem.Size(m)
}
func (m *SearchItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchItem.DiscardUnknown(m)
}

var xxx_messageInfo_SearchItem proto.InternalMessageInfo

func (m *SearchItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchItem) GetSearchCount() int64 {
	if m != nil {
		return m.SearchCount
	}
	return 0
}

type SingleWordSearchRequest struct {
	Word                 string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleWordSearchRequest) Reset()         { *m = SingleWordSearchRequest{} }
func (m *SingleWordSearchRequest) String() string { return proto.CompactTextString(m) }
func (*SingleWordSearchRequest) ProtoMessage()    {}
func (*SingleWordSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cf4e25ab144b61, []int{1}
}

func (m *SingleWordSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleWordSearchRequest.Unmarshal(m, b)
}
func (m *SingleWordSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleWordSearchRequest.Marshal(b, m, deterministic)
}
func (m *SingleWordSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleWordSearchRequest.Merge(m, src)
}
func (m *SingleWordSearchRequest) XXX_Size() int {
	return xxx_messageInfo_SingleWordSearchRequest.Size(m)
}
func (m *SingleWordSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleWordSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SingleWordSearchRequest proto.InternalMessageInfo

func (m *SingleWordSearchRequest) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type SingleWordSearchResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleWordSearchResponse) Reset()         { *m = SingleWordSearchResponse{} }
func (m *SingleWordSearchResponse) String() string { return proto.CompactTextString(m) }
func (*SingleWordSearchResponse) ProtoMessage()    {}
func (*SingleWordSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cf4e25ab144b61, []int{2}
}

func (m *SingleWordSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleWordSearchResponse.Unmarshal(m, b)
}
func (m *SingleWordSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleWordSearchResponse.Marshal(b, m, deterministic)
}
func (m *SingleWordSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleWordSearchResponse.Merge(m, src)
}
func (m *SingleWordSearchResponse) XXX_Size() int {
	return xxx_messageInfo_SingleWordSearchResponse.Size(m)
}
func (m *SingleWordSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleWordSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SingleWordSearchResponse proto.InternalMessageInfo

func (m *SingleWordSearchResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UpdateSearchListRequest struct {
	Word                 string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSearchListRequest) Reset()         { *m = UpdateSearchListRequest{} }
func (m *UpdateSearchListRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSearchListRequest) ProtoMessage()    {}
func (*UpdateSearchListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cf4e25ab144b61, []int{3}
}

func (m *UpdateSearchListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSearchListRequest.Unmarshal(m, b)
}
func (m *UpdateSearchListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSearchListRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSearchListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSearchListRequest.Merge(m, src)
}
func (m *UpdateSearchListRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSearchListRequest.Size(m)
}
func (m *UpdateSearchListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSearchListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSearchListRequest proto.InternalMessageInfo

func (m *UpdateSearchListRequest) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type UpdateSearchListResponse struct {
	Message              string        `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	WordList             []*SearchItem `protobuf:"bytes,2,rep,name=word_list,json=wordList,proto3" json:"word_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateSearchListResponse) Reset()         { *m = UpdateSearchListResponse{} }
func (m *UpdateSearchListResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateSearchListResponse) ProtoMessage()    {}
func (*UpdateSearchListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cf4e25ab144b61, []int{4}
}

func (m *UpdateSearchListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSearchListResponse.Unmarshal(m, b)
}
func (m *UpdateSearchListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSearchListResponse.Marshal(b, m, deterministic)
}
func (m *UpdateSearchListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSearchListResponse.Merge(m, src)
}
func (m *UpdateSearchListResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateSearchListResponse.Size(m)
}
func (m *UpdateSearchListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSearchListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSearchListResponse proto.InternalMessageInfo

func (m *UpdateSearchListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UpdateSearchListResponse) GetWordList() []*SearchItem {
	if m != nil {
		return m.WordList
	}
	return nil
}

type TopFiveRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopFiveRequest) Reset()         { *m = TopFiveRequest{} }
func (m *TopFiveRequest) String() string { return proto.CompactTextString(m) }
func (*TopFiveRequest) ProtoMessage()    {}
func (*TopFiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cf4e25ab144b61, []int{5}
}

func (m *TopFiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopFiveRequest.Unmarshal(m, b)
}
func (m *TopFiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopFiveRequest.Marshal(b, m, deterministic)
}
func (m *TopFiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopFiveRequest.Merge(m, src)
}
func (m *TopFiveRequest) XXX_Size() int {
	return xxx_messageInfo_TopFiveRequest.Size(m)
}
func (m *TopFiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopFiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopFiveRequest proto.InternalMessageInfo

type TopFiveResponse struct {
	TopFive              []*SearchItem `protobuf:"bytes,1,rep,name=top_five,json=topFive,proto3" json:"top_five,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TopFiveResponse) Reset()         { *m = TopFiveResponse{} }
func (m *TopFiveResponse) String() string { return proto.CompactTextString(m) }
func (*TopFiveResponse) ProtoMessage()    {}
func (*TopFiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cf4e25ab144b61, []int{6}
}

func (m *TopFiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopFiveResponse.Unmarshal(m, b)
}
func (m *TopFiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopFiveResponse.Marshal(b, m, deterministic)
}
func (m *TopFiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopFiveResponse.Merge(m, src)
}
func (m *TopFiveResponse) XXX_Size() int {
	return xxx_messageInfo_TopFiveResponse.Size(m)
}
func (m *TopFiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TopFiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TopFiveResponse proto.InternalMessageInfo

func (m *TopFiveResponse) GetTopFive() []*SearchItem {
	if m != nil {
		return m.TopFive
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchItem)(nil), "wordsearch.SearchItem")
	proto.RegisterType((*SingleWordSearchRequest)(nil), "wordsearch.SingleWordSearchRequest")
	proto.RegisterType((*SingleWordSearchResponse)(nil), "wordsearch.SingleWordSearchResponse")
	proto.RegisterType((*UpdateSearchListRequest)(nil), "wordsearch.UpdateSearchListRequest")
	proto.RegisterType((*UpdateSearchListResponse)(nil), "wordsearch.UpdateSearchListResponse")
	proto.RegisterType((*TopFiveRequest)(nil), "wordsearch.TopFiveRequest")
	proto.RegisterType((*TopFiveResponse)(nil), "wordsearch.TopFiveResponse")
}

func init() {
	proto.RegisterFile("proto/newworsearch.proto", fileDescriptor_d8cf4e25ab144b61)
}

var fileDescriptor_d8cf4e25ab144b61 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x53, 0xae, 0xf1, 0xde, 0x7b, 0xee, 0x8d, 0xe2, 0x84, 0x60, 0x53, 0x59, 0x60, 0x75,
	0x41, 0x4c, 0x68, 0x23, 0xb8, 0xd2, 0x9d, 0x18, 0x13, 0x13, 0xc3, 0xa2, 0x68, 0x34, 0x6e, 0xc8,
	0x00, 0x63, 0x99, 0xd8, 0xce, 0x0c, 0x9d, 0x29, 0x2c, 0x88, 0x1b, 0x5f, 0xc1, 0x47, 0xf3, 0x15,
	0xdc, 0xf8, 0x16, 0x66, 0x66, 0x0a, 0xa5, 0xfc, 0xbb, 0xab, 0x4e, 0x4f, 0xbf, 0xf3, 0xfd, 0x4e,
	0xbf, 0xd3, 0x82, 0x2b, 0x32, 0xae, 0x78, 0xc8, 0xc8, 0x6a, 0xc5, 0x33, 0x49, 0x70, 0x36, 0x9d,
	0x07, 0xa6, 0x84, 0x60, 0xc5, 0xb3, 0x99, 0xad, 0x78, 0xad, 0x98, 0xf3, 0x38, 0x21, 0x21, 0x16,
	0x34, 0xc4, 0x8c, 0x71, 0x85, 0x15, 0xe5, 0x4c, 0x5a, 0xa5, 0x3f, 0x00, 0x18, 0x19, 0xdd, 0x07,
	0x45, 0x52, 0x84, 0xe0, 0x1e, 0xc3, 0x29, 0x71, 0x9d, 0xb6, 0xd3, 0xb9, 0x8e, 0xcc, 0x19, 0x3d,
	0x85, 0x5b, 0xeb, 0x34, 0x9e, 0xf2, 0x9c, 0x29, 0xb7, 0xd6, 0x76, 0x3a, 0x17, 0xd1, 0x8d, 0xad,
	0x0d, 0x74, 0xc9, 0xef, 0xc2, 0xe3, 0x11, 0x65, 0x71, 0x42, 0xbe, 0xf0, 0x6c, 0x66, 0xed, 0x22,
	0xb2, 0xc8, 0x89, 0x54, 0xda, 0x51, 0xcf, 0xb2, 0x71, 0xd4, 0x67, 0xff, 0x15, 0xb8, 0x87, 0x72,
	0x29, 0x38, 0x93, 0x04, 0xb9, 0x70, 0x99, 0x12, 0x29, 0x71, 0xbc, 0x19, 0x62, 0x73, 0xab, 0x21,
	0x9f, 0xc5, 0x0c, 0x2b, 0x62, 0x3b, 0x3e, 0x52, 0xa9, 0xce, 0x41, 0x28, 0xb8, 0x87, 0xf2, 0xbb,
	0x20, 0xa8, 0x0f, 0xd7, 0xba, 0x7b, 0x9c, 0x50, 0xa9, 0xdf, 0xf4, 0xa2, 0x73, 0xd3, 0x6b, 0x06,
	0x65, 0x98, 0x41, 0x99, 0x55, 0x74, 0xa5, 0xcb, 0xda, 0xd6, 0xaf, 0xc3, 0x83, 0x4f, 0x5c, 0xbc,
	0xa7, 0x4b, 0x52, 0x0c, 0xe4, 0xbf, 0x83, 0x87, 0xdb, 0x4a, 0xc1, 0x7c, 0x09, 0x57, 0x8a, 0x8b,
	0xf1, 0x77, 0xba, 0xd4, 0xd0, 0x73, 0xc6, 0x97, 0xca, 0xb6, 0xf6, 0xfe, 0xd5, 0x00, 0xca, 0x88,
	0xd0, 0x02, 0xea, 0xfb, 0xb1, 0xa1, 0x67, 0x15, 0x8f, 0xe3, 0x3b, 0xf0, 0x9e, 0x9f, 0x17, 0xd9,
	0x01, 0x7d, 0xf4, 0xeb, 0xcf, 0xdf, 0xdf, 0xb5, 0x5b, 0x04, 0xe6, 0x4b, 0x31, 0x1d, 0x68, 0x0d,
	0xf5, 0xfd, 0x10, 0xab, 0xc8, 0x13, 0x1b, 0xa9, 0x22, 0x4f, 0xed, 0xc1, 0x6f, 0x19, 0x64, 0xd3,
	0x7b, 0x54, 0x22, 0xc3, 0xb5, 0xbe, 0xfc, 0x7c, 0xed, 0xbc, 0x40, 0x3f, 0xa0, 0x51, 0x84, 0xb8,
	0x9d, 0x34, 0x4f, 0x94, 0x44, 0xde, 0xae, 0x77, 0x35, 0x78, 0xef, 0xc9, 0xd1, 0x67, 0x05, 0xce,
	0x33, 0xb8, 0x06, 0x42, 0x3b, 0x38, 0xc1, 0x45, 0x9e, 0xe0, 0xec, 0x6d, 0xf8, 0xad, 0x1b, 0x53,
	0x35, 0xcf, 0x27, 0xc1, 0x94, 0xa7, 0x61, 0x94, 0x0f, 0x05, 0xfd, 0x4a, 0x92, 0x2c, 0x1f, 0x1a,
	0x5d, 0xd7, 0x3a, 0xbe, 0x29, 0xcd, 0x27, 0xf7, 0xcd, 0xff, 0xd3, 0xff, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x4d, 0x96, 0x86, 0xb1, 0x85, 0x03, 0x00, 0x00,
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
	SingleWordSearch(ctx context.Context, in *SingleWordSearchRequest, opts ...grpc.CallOption) (*SingleWordSearchResponse, error)
	UpdateSearchList(ctx context.Context, in *UpdateSearchListRequest, opts ...grpc.CallOption) (*UpdateSearchListResponse, error)
	TopFiveSearchResults(ctx context.Context, in *TopFiveRequest, opts ...grpc.CallOption) (*TopFiveResponse, error)
}

type wordSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewWordSearchClient(cc grpc.ClientConnInterface) WordSearchClient {
	return &wordSearchClient{cc}
}

func (c *wordSearchClient) SingleWordSearch(ctx context.Context, in *SingleWordSearchRequest, opts ...grpc.CallOption) (*SingleWordSearchResponse, error) {
	out := new(SingleWordSearchResponse)
	err := c.cc.Invoke(ctx, "/wordsearch.WordSearch/SingleWordSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordSearchClient) UpdateSearchList(ctx context.Context, in *UpdateSearchListRequest, opts ...grpc.CallOption) (*UpdateSearchListResponse, error) {
	out := new(UpdateSearchListResponse)
	err := c.cc.Invoke(ctx, "/wordsearch.WordSearch/UpdateSearchList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordSearchClient) TopFiveSearchResults(ctx context.Context, in *TopFiveRequest, opts ...grpc.CallOption) (*TopFiveResponse, error) {
	out := new(TopFiveResponse)
	err := c.cc.Invoke(ctx, "/wordsearch.WordSearch/TopFiveSearchResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordSearchServer is the server API for WordSearch service.
type WordSearchServer interface {
	SingleWordSearch(context.Context, *SingleWordSearchRequest) (*SingleWordSearchResponse, error)
	UpdateSearchList(context.Context, *UpdateSearchListRequest) (*UpdateSearchListResponse, error)
	TopFiveSearchResults(context.Context, *TopFiveRequest) (*TopFiveResponse, error)
}

// UnimplementedWordSearchServer can be embedded to have forward compatible implementations.
type UnimplementedWordSearchServer struct {
}

func (*UnimplementedWordSearchServer) SingleWordSearch(ctx context.Context, req *SingleWordSearchRequest) (*SingleWordSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SingleWordSearch not implemented")
}
func (*UnimplementedWordSearchServer) UpdateSearchList(ctx context.Context, req *UpdateSearchListRequest) (*UpdateSearchListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSearchList not implemented")
}
func (*UnimplementedWordSearchServer) TopFiveSearchResults(ctx context.Context, req *TopFiveRequest) (*TopFiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopFiveSearchResults not implemented")
}

func RegisterWordSearchServer(s *grpc.Server, srv WordSearchServer) {
	s.RegisterService(&_WordSearch_serviceDesc, srv)
}

func _WordSearch_SingleWordSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleWordSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordSearchServer).SingleWordSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wordsearch.WordSearch/SingleWordSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordSearchServer).SingleWordSearch(ctx, req.(*SingleWordSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordSearch_UpdateSearchList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSearchListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordSearchServer).UpdateSearchList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wordsearch.WordSearch/UpdateSearchList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordSearchServer).UpdateSearchList(ctx, req.(*UpdateSearchListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordSearch_TopFiveSearchResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopFiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordSearchServer).TopFiveSearchResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wordsearch.WordSearch/TopFiveSearchResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordSearchServer).TopFiveSearchResults(ctx, req.(*TopFiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordSearch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wordsearch.WordSearch",
	HandlerType: (*WordSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SingleWordSearch",
			Handler:    _WordSearch_SingleWordSearch_Handler,
		},
		{
			MethodName: "UpdateSearchList",
			Handler:    _WordSearch_UpdateSearchList_Handler,
		},
		{
			MethodName: "TopFiveSearchResults",
			Handler:    _WordSearch_TopFiveSearchResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/newworsearch.proto",
}
