// Code generated by protoc-gen-go. DO NOT EDIT.
// source: portsgrpc.proto

package portsgrpc

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

type Ports struct {
	PortId               string    `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string    `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string    `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string  `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string  `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []float32 `protobuf:"fixed32,7,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string    `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string    `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string  `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 string    `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Ports) Reset()         { *m = Ports{} }
func (m *Ports) String() string { return proto.CompactTextString(m) }
func (*Ports) ProtoMessage()    {}
func (*Ports) Descriptor() ([]byte, []int) {
	return fileDescriptor_09e7f44acbd5ff64, []int{0}
}

func (m *Ports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ports.Unmarshal(m, b)
}
func (m *Ports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ports.Marshal(b, m, deterministic)
}
func (m *Ports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ports.Merge(m, src)
}
func (m *Ports) XXX_Size() int {
	return xxx_messageInfo_Ports.Size(m)
}
func (m *Ports) XXX_DiscardUnknown() {
	xxx_messageInfo_Ports.DiscardUnknown(m)
}

var xxx_messageInfo_Ports proto.InternalMessageInfo

func (m *Ports) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *Ports) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ports) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Ports) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Ports) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Ports) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Ports) GetCoordinates() []float32 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Ports) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Ports) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Ports) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *Ports) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Response struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_09e7f44acbd5ff64, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Ports)(nil), "portsgrpc.Ports")
	proto.RegisterType((*Response)(nil), "portsgrpc.Response")
}

func init() { proto.RegisterFile("portsgrpc.proto", fileDescriptor_09e7f44acbd5ff64) }

var fileDescriptor_09e7f44acbd5ff64 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x3f, 0x6b, 0xf3, 0x30,
	0x10, 0x87, 0xdf, 0xfc, 0xb1, 0x1d, 0x5f, 0x86, 0xb7, 0xa8, 0xa5, 0x3d, 0x32, 0x99, 0x4c, 0x9e,
	0x42, 0x69, 0x97, 0x42, 0xd7, 0x2e, 0xdd, 0x8a, 0xa1, 0x73, 0x71, 0xe4, 0xc3, 0x08, 0x62, 0x9d,
	0x90, 0x94, 0x42, 0xfa, 0x49, 0xfa, 0x71, 0x8b, 0xe4, 0xc8, 0xc9, 0x76, 0xcf, 0xf3, 0x13, 0x9c,
	0xee, 0x0e, 0xfe, 0x1b, 0xb6, 0xde, 0xf5, 0xd6, 0xc8, 0x9d, 0xb1, 0xec, 0x59, 0x94, 0x93, 0xd8,
	0xfe, 0xce, 0x21, 0xfb, 0x08, 0x24, 0x1e, 0xa0, 0x08, 0xfa, 0x4b, 0x75, 0x38, 0xab, 0x66, 0x75,
	0xd9, 0xe4, 0x01, 0xdf, 0x3b, 0x21, 0x60, 0xa9, 0xdb, 0x81, 0x70, 0x1e, 0x6d, 0xac, 0x83, 0x93,
	0xca, 0x9f, 0x70, 0x31, 0xba, 0x50, 0x0b, 0x84, 0x42, 0xf2, 0x51, 0x7b, 0x7b, 0xc2, 0x65, 0xd4,
	0x09, 0xc5, 0x1d, 0x64, 0xed, 0x41, 0xb5, 0x0e, 0xb3, 0x6a, 0x51, 0x97, 0xcd, 0x08, 0xe1, 0xbd,
	0xa5, 0x5e, 0xb1, 0x76, 0x98, 0x47, 0x9f, 0x50, 0x54, 0xb0, 0x96, 0xcc, 0xb6, 0x53, 0xba, 0xf5,
	0xe4, 0xb0, 0xa8, 0x16, 0xf5, 0xbc, 0xb9, 0x56, 0x62, 0x03, 0x2b, 0x63, 0xf9, 0x5b, 0x69, 0x49,
	0xb8, 0x8a, 0xcd, 0x26, 0x0e, 0x99, 0x57, 0x03, 0xfd, 0xb0, 0x26, 0x2c, 0xc7, 0x2c, 0xb1, 0xb8,
	0x87, 0xfc, 0xa8, 0x0f, 0x2c, 0x1d, 0x42, 0x6c, 0x79, 0xa6, 0x38, 0x0f, 0x77, 0x84, 0xeb, 0xf3,
	0x3c, 0xdc, 0xd1, 0xf6, 0x05, 0x56, 0x0d, 0x39, 0xc3, 0xda, 0xd1, 0x94, 0xcf, 0x2e, 0x79, 0xf8,
	0xff, 0x40, 0xce, 0xb5, 0x7d, 0x5a, 0x4d, 0xc2, 0xa7, 0x57, 0x28, 0xe2, 0x4e, 0xdf, 0xf6, 0xe2,
	0x11, 0xb2, 0x4f, 0xe3, 0xc8, 0x8b, 0x9b, 0xdd, 0xe5, 0x0a, 0x31, 0xdc, 0xdc, 0x5e, 0x99, 0xd4,
	0x68, 0xfb, 0x6f, 0x9f, 0xc7, 0x1b, 0x3d, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xea, 0x0c, 0x14,
	0x76, 0xb6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PortsDbClient is the client API for PortsDb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortsDbClient interface {
	Upset(ctx context.Context, in *Ports, opts ...grpc.CallOption) (*Response, error)
}

type portsDbClient struct {
	cc *grpc.ClientConn
}

func NewPortsDbClient(cc *grpc.ClientConn) PortsDbClient {
	return &portsDbClient{cc}
}

func (c *portsDbClient) Upset(ctx context.Context, in *Ports, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/portsgrpc.PortsDb/Upset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortsDbServer is the server API for PortsDb service.
type PortsDbServer interface {
	Upset(context.Context, *Ports) (*Response, error)
}

// UnimplementedPortsDbServer can be embedded to have forward compatible implementations.
type UnimplementedPortsDbServer struct {
}

func (*UnimplementedPortsDbServer) Upset(ctx context.Context, req *Ports) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upset not implemented")
}

func RegisterPortsDbServer(s *grpc.Server, srv PortsDbServer) {
	s.RegisterService(&_PortsDb_serviceDesc, srv)
}

func _PortsDb_Upset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ports)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortsDbServer).Upset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portsgrpc.PortsDb/Upset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortsDbServer).Upset(ctx, req.(*Ports))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortsDb_serviceDesc = grpc.ServiceDesc{
	ServiceName: "portsgrpc.PortsDb",
	HandlerType: (*PortsDbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upset",
			Handler:    _PortsDb_Upset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portsgrpc.proto",
}