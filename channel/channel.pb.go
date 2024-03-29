// Code generated by protoc-gen-go. DO NOT EDIT.
// source: channel.proto

package infomation

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Param struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Param) Reset()         { *m = Param{} }
func (m *Param) String() string { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()    {}
func (*Param) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8f385724121f37b, []int{0}
}

func (m *Param) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Param.Unmarshal(m, b)
}
func (m *Param) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Param.Marshal(b, m, deterministic)
}
func (m *Param) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Param.Merge(m, src)
}
func (m *Param) XXX_Size() int {
	return xxx_messageInfo_Param.Size(m)
}
func (m *Param) XXX_DiscardUnknown() {
	xxx_messageInfo_Param.DiscardUnknown(m)
}

var xxx_messageInfo_Param proto.InternalMessageInfo

func (m *Param) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ChannelInfo struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BmsCode              string   `protobuf:"bytes,3,opt,name=bms_code,json=bmsCode,proto3" json:"bms_code,omitempty"`
	DefinitionCode       string   `protobuf:"bytes,4,opt,name=definition_code,json=definitionCode,proto3" json:"definition_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelInfo) Reset()         { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()    {}
func (*ChannelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8f385724121f37b, []int{1}
}

func (m *ChannelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelInfo.Unmarshal(m, b)
}
func (m *ChannelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelInfo.Marshal(b, m, deterministic)
}
func (m *ChannelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelInfo.Merge(m, src)
}
func (m *ChannelInfo) XXX_Size() int {
	return xxx_messageInfo_ChannelInfo.Size(m)
}
func (m *ChannelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelInfo proto.InternalMessageInfo

func (m *ChannelInfo) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ChannelInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChannelInfo) GetBmsCode() string {
	if m != nil {
		return m.BmsCode
	}
	return ""
}

func (m *ChannelInfo) GetDefinitionCode() string {
	if m != nil {
		return m.DefinitionCode
	}
	return ""
}

func init() {
	proto.RegisterType((*Param)(nil), "infomation.Param")
	proto.RegisterType((*ChannelInfo)(nil), "infomation.ChannelInfo")
}

func init() { proto.RegisterFile("channel.proto", fileDescriptor_c8f385724121f37b) }

var fileDescriptor_c8f385724121f37b = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0xcc,
	0xcb, 0x4b, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xca, 0xcc, 0x4b, 0xcb, 0xcf,
	0x4d, 0x2c, 0xc9, 0xcc, 0xcf, 0x53, 0x92, 0xe6, 0x62, 0x0d, 0x48, 0x2c, 0x4a, 0xcc, 0x15, 0x12,
	0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95,
	0x2a, 0xb9, 0xb8, 0x9d, 0x21, 0x3a, 0x3d, 0xf3, 0xd2, 0xf2, 0xb1, 0x29, 0x01, 0x89, 0xe5, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0x41, 0xc4, 0x40, 0x6c, 0x21, 0x49, 0x2e, 0x8e, 0xa4, 0xdc, 0xe2, 0x78,
	0xb0, 0x5a, 0x66, 0xb0, 0x38, 0x7b, 0x52, 0x6e, 0xb1, 0x33, 0x48, 0xb9, 0x3a, 0x17, 0x7f, 0x4a,
	0x6a, 0x5a, 0x66, 0x5e, 0x26, 0xc8, 0x72, 0x88, 0x0a, 0x16, 0xb0, 0x0a, 0x3e, 0x84, 0x30, 0x48,
	0xa1, 0x91, 0x37, 0x17, 0x1f, 0xd4, 0xea, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x4b,
	0x2e, 0xce, 0xf4, 0xd4, 0x12, 0xa7, 0x4a, 0xb0, 0x39, 0x82, 0x7a, 0x08, 0x3f, 0xe8, 0x81, 0x3d,
	0x20, 0x25, 0x8e, 0x2c, 0x84, 0xe4, 0x6c, 0x25, 0x06, 0x27, 0x4e, 0x2e, 0xf6, 0xd4, 0x0a, 0xbd,
	0xf4, 0xa2, 0x82, 0xe4, 0x24, 0x36, 0x70, 0x10, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0x75, 0xc1, 0x84, 0x13, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChannelServiceClient is the client API for ChannelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChannelServiceClient interface {
	GetByCode(ctx context.Context, in *Param, opts ...grpc.CallOption) (*ChannelInfo, error)
}

type channelServiceClient struct {
	cc *grpc.ClientConn
}

func NewChannelServiceClient(cc *grpc.ClientConn) ChannelServiceClient {
	return &channelServiceClient{cc}
}

func (c *channelServiceClient) GetByCode(ctx context.Context, in *Param, opts ...grpc.CallOption) (*ChannelInfo, error) {
	out := new(ChannelInfo)
	err := c.cc.Invoke(ctx, "/infomation.ChannelService/getByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelServiceServer is the server API for ChannelService service.
type ChannelServiceServer interface {
	GetByCode(context.Context, *Param) (*ChannelInfo, error)
}

// UnimplementedChannelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChannelServiceServer struct {
}

func (*UnimplementedChannelServiceServer) GetByCode(ctx context.Context, req *Param) (*ChannelInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCode not implemented")
}

func RegisterChannelServiceServer(s *grpc.Server, srv ChannelServiceServer) {
	s.RegisterService(&_ChannelService_serviceDesc, srv)
}

func _ChannelService_GetByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Param)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelServiceServer).GetByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infomation.ChannelService/GetByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).GetByCode(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChannelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infomation.ChannelService",
	HandlerType: (*ChannelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getByCode",
			Handler:    _ChannelService_GetByCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channel.proto",
}
