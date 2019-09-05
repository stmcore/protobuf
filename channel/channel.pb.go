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
	proto.RegisterType((*Param)(nil), "channel.Param")
	proto.RegisterType((*ChannelInfo)(nil), "channel.ChannelInfo")
}

func init() { proto.RegisterFile("channel.proto", fileDescriptor_c8f385724121f37b) }

var fileDescriptor_c8f385724121f37b = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0xcc,
	0xcb, 0x4b, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xa4, 0xb9,
	0x58, 0x03, 0x12, 0x8b, 0x12, 0x73, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa5, 0x4a, 0x2e, 0x6e, 0x67, 0x88, 0x3a, 0xcf, 0xbc,
	0xb4, 0x7c, 0x6c, 0x4a, 0x40, 0x62, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x10, 0x31, 0x10, 0x5b,
	0x48, 0x92, 0x8b, 0x23, 0x29, 0xb7, 0x38, 0x1e, 0xac, 0x96, 0x19, 0x2c, 0xce, 0x9e, 0x94, 0x5b,
	0xec, 0x0c, 0x52, 0xae, 0xce, 0xc5, 0x9f, 0x92, 0x9a, 0x96, 0x99, 0x97, 0x59, 0x92, 0x99, 0x9f,
	0x07, 0x51, 0xc1, 0x02, 0x56, 0xc1, 0x87, 0x10, 0x06, 0x29, 0x34, 0x72, 0xe5, 0xe2, 0x83, 0x5a,
	0x1d, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xcc, 0xc5, 0x99, 0x9e, 0x5a, 0xe2, 0x54,
	0x09, 0x36, 0x87, 0x4f, 0x0f, 0xe6, 0x1f, 0xb0, 0xeb, 0xa5, 0x44, 0xe0, 0x7c, 0x24, 0x07, 0x2b,
	0x31, 0x38, 0x71, 0x72, 0xb1, 0xa7, 0x56, 0xe8, 0xa5, 0x17, 0x15, 0x24, 0x27, 0xb1, 0x81, 0x7d,
	0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x90, 0xf2, 0x3d, 0x0a, 0x01, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/channel.ChannelService/getByCode", in, out, opts...)
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
		FullMethod: "/channel.ChannelService/GetByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelServiceServer).GetByCode(ctx, req.(*Param))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChannelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "channel.ChannelService",
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
