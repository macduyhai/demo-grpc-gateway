// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gatewayproto/gatewaypb.proto

package gatewaypb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

type PingRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e562cc4ef2f7e84, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e562cc4ef2f7e84, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e562cc4ef2f7e84, []int{2}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type CreateUserResponse struct {
	Code                 int32                    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Messsage             string                   `protobuf:"bytes,2,opt,name=messsage,proto3" json:"messsage,omitempty"`
	Data                 *CreateUserResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e562cc4ef2f7e84, []int{3}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateUserResponse) GetMesssage() string {
	if m != nil {
		return m.Messsage
	}
	return ""
}

func (m *CreateUserResponse) GetData() *CreateUserResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateUserResponse_Data struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse_Data) Reset()         { *m = CreateUserResponse_Data{} }
func (m *CreateUserResponse_Data) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse_Data) ProtoMessage()    {}
func (*CreateUserResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e562cc4ef2f7e84, []int{3, 0}
}

func (m *CreateUserResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse_Data.Unmarshal(m, b)
}
func (m *CreateUserResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse_Data.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse_Data.Merge(m, src)
}
func (m *CreateUserResponse_Data) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse_Data.Size(m)
}
func (m *CreateUserResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse_Data proto.InternalMessageInfo

func (m *CreateUserResponse_Data) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
	proto.RegisterType((*PingResponse)(nil), "PingResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "CreateUserResponse")
	proto.RegisterType((*CreateUserResponse_Data)(nil), "CreateUserResponse.Data")
}

func init() { proto.RegisterFile("gatewayproto/gatewaypb.proto", fileDescriptor_8e562cc4ef2f7e84) }

var fileDescriptor_8e562cc4ef2f7e84 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x99, 0x36, 0x6d, 0x6f, 0x4e, 0x7a, 0xa1, 0x3d, 0x17, 0xee, 0x0d, 0xb9, 0x05, 0x4b,
	0x36, 0x86, 0x22, 0x09, 0xd6, 0x95, 0x5d, 0x5a, 0xa1, 0xb8, 0x93, 0x88, 0x1b, 0x37, 0x32, 0x6d,
	0x86, 0x38, 0xd8, 0x66, 0x62, 0x66, 0xaa, 0x74, 0xeb, 0x2b, 0x74, 0xef, 0x4b, 0xf9, 0x0a, 0x3e,
	0x88, 0xcc, 0xa4, 0xa9, 0x81, 0x8a, 0xbb, 0xf3, 0x71, 0xfe, 0xf3, 0x9f, 0xc3, 0x3f, 0x03, 0x83,
	0x94, 0x2a, 0xf6, 0x42, 0x37, 0x79, 0x21, 0x94, 0x88, 0x2a, 0x98, 0x87, 0x86, 0xbd, 0x41, 0x2a,
	0x44, 0xba, 0x64, 0x11, 0xcd, 0x79, 0x44, 0xb3, 0x4c, 0x28, 0xaa, 0xb8, 0xc8, 0xe4, 0xae, 0xfb,
	0xbf, 0xd6, 0x7d, 0xe6, 0x92, 0xcf, 0xf9, 0x92, 0xab, 0x4d, 0xd9, 0xf4, 0x8f, 0xc1, 0xb9, 0xe6,
	0x59, 0x1a, 0xb3, 0xa7, 0x35, 0x93, 0x0a, 0x5d, 0xe8, 0xac, 0x98, 0x94, 0x34, 0x65, 0x2e, 0x19,
	0x92, 0xc0, 0x8e, 0x2b, 0xf4, 0x03, 0xe8, 0x96, 0x42, 0x99, 0x8b, 0x4c, 0xb2, 0x1f, 0x94, 0xe7,
	0xd0, 0x9f, 0x16, 0x8c, 0x2a, 0x76, 0x2b, 0x59, 0x51, 0x19, 0x23, 0x58, 0x19, 0x5d, 0x55, 0x5a,
	0x53, 0x63, 0x0f, 0x9a, 0x7a, 0xbc, 0x31, 0x24, 0x41, 0x2b, 0xd6, 0xa5, 0xbf, 0x25, 0x80, 0xf5,
	0xd9, 0xdd, 0x2e, 0x04, 0x6b, 0x21, 0x92, 0x72, 0xb8, 0x15, 0x9b, 0x1a, 0x3d, 0xf8, 0xa5, 0x17,
	0xca, 0xca, 0xc1, 0x8e, 0xf7, 0x8c, 0x27, 0x60, 0x25, 0x54, 0x51, 0xb7, 0x39, 0x24, 0x81, 0x33,
	0x76, 0xc3, 0x43, 0xcb, 0xf0, 0x92, 0x2a, 0x1a, 0x1b, 0x95, 0x77, 0x04, 0x96, 0x26, 0xfc, 0x07,
	0x9d, 0xb5, 0x64, 0xc5, 0x3d, 0x4f, 0xcc, 0xa2, 0x66, 0xdc, 0xd6, 0x78, 0x95, 0x8c, 0xdf, 0x08,
	0x38, 0xb3, 0x22, 0x5f, 0xcc, 0xca, 0xd8, 0x71, 0x02, 0xf6, 0xf4, 0x81, 0x2d, 0x1e, 0x75, 0x1e,
	0xd8, 0x0d, 0x6b, 0xf9, 0x79, 0xbf, 0xc3, 0x7a, 0x48, 0x7e, 0xef, 0xf5, 0xfd, 0x63, 0xdb, 0x00,
	0x6c, 0x45, 0x39, 0xcf, 0xd2, 0x09, 0x19, 0xe1, 0x0d, 0xc0, 0xd7, 0x35, 0x88, 0xe1, 0x41, 0x52,
	0xde, 0x9f, 0x6f, 0xce, 0xf5, 0x07, 0xc6, 0xe8, 0xaf, 0xdf, 0x2f, 0x5f, 0xf1, 0x34, 0xd2, 0xa7,
	0xc9, 0x88, 0x26, 0xc9, 0x84, 0x8c, 0x2e, 0x9c, 0x3b, 0x7b, 0xff, 0x25, 0xe6, 0x6d, 0xf3, 0xb0,
	0x67, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x93, 0x77, 0x4c, 0x33, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GrpcGatewayClient is the client API for GrpcGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcGatewayClient interface {
	CheckPing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type grpcGatewayClient struct {
	cc *grpc.ClientConn
}

func NewGrpcGatewayClient(cc *grpc.ClientConn) GrpcGatewayClient {
	return &grpcGatewayClient{cc}
}

func (c *grpcGatewayClient) CheckPing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/GrpcGateway/CheckPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/GrpcGateway/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcGatewayServer is the server API for GrpcGateway service.
type GrpcGatewayServer interface {
	CheckPing(context.Context, *PingRequest) (*PingResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedGrpcGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcGatewayServer struct {
}

func (*UnimplementedGrpcGatewayServer) CheckPing(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPing not implemented")
}
func (*UnimplementedGrpcGatewayServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterGrpcGatewayServer(s *grpc.Server, srv GrpcGatewayServer) {
	s.RegisterService(&_GrpcGateway_serviceDesc, srv)
}

func _GrpcGateway_CheckPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).CheckPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GrpcGateway/CheckPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).CheckPing(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcGateway_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GrpcGateway/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GrpcGateway",
	HandlerType: (*GrpcGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPing",
			Handler:    _GrpcGateway_CheckPing_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _GrpcGateway_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gatewayproto/gatewaypb.proto",
}
