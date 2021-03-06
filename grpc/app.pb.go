// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

package app

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

type App struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cover                string   `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	Download             string   `protobuf:"bytes,3,opt,name=download,proto3" json:"download,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{0}
}

func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *App) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *App) GetDownload() string {
	if m != nil {
		return m.Download
	}
	return ""
}

type GetAppReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppReq) Reset()         { *m = GetAppReq{} }
func (m *GetAppReq) String() string { return proto.CompactTextString(m) }
func (*GetAppReq) ProtoMessage()    {}
func (*GetAppReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{1}
}

func (m *GetAppReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppReq.Unmarshal(m, b)
}
func (m *GetAppReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppReq.Marshal(b, m, deterministic)
}
func (m *GetAppReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppReq.Merge(m, src)
}
func (m *GetAppReq) XXX_Size() int {
	return xxx_messageInfo_GetAppReq.Size(m)
}
func (m *GetAppReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppReq proto.InternalMessageInfo

func (m *GetAppReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAppRes struct {
	App                  *App     `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppRes) Reset()         { *m = GetAppRes{} }
func (m *GetAppRes) String() string { return proto.CompactTextString(m) }
func (*GetAppRes) ProtoMessage()    {}
func (*GetAppRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{2}
}

func (m *GetAppRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppRes.Unmarshal(m, b)
}
func (m *GetAppRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppRes.Marshal(b, m, deterministic)
}
func (m *GetAppRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppRes.Merge(m, src)
}
func (m *GetAppRes) XXX_Size() int {
	return xxx_messageInfo_GetAppRes.Size(m)
}
func (m *GetAppRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppRes proto.InternalMessageInfo

func (m *GetAppRes) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

func init() {
	proto.RegisterType((*App)(nil), "app.App")
	proto.RegisterType((*GetAppReq)(nil), "app.GetAppReq")
	proto.RegisterType((*GetAppRes)(nil), "app.GetAppRes")
}

func init() { proto.RegisterFile("app.proto", fileDescriptor_e0f9056a14b86d47) }

var fileDescriptor_e0f9056a14b86d47 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0x28, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0x28, 0x50, 0x72, 0xe7, 0x62, 0x76, 0x2c,
	0x28, 0x10, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xca,
	0x4c, 0x11, 0x12, 0xe1, 0x62, 0x4d, 0xce, 0x2f, 0x4b, 0x2d, 0x92, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x0c, 0x82, 0x70, 0x84, 0xa4, 0xb8, 0x38, 0x52, 0xf2, 0xcb, 0xf3, 0x72, 0xf2, 0x13, 0x53, 0x24,
	0x98, 0xc1, 0x12, 0x70, 0xbe, 0x92, 0x34, 0x17, 0xa7, 0x7b, 0x6a, 0x89, 0x63, 0x41, 0x41, 0x50,
	0x6a, 0x21, 0xba, 0x71, 0x4a, 0xea, 0x08, 0xc9, 0x62, 0x21, 0x29, 0x2e, 0x90, 0xcd, 0x60, 0x59,
	0x6e, 0x23, 0x0e, 0x3d, 0x90, 0x83, 0x40, 0x32, 0x20, 0x41, 0x23, 0x0b, 0x2e, 0x2e, 0xc7, 0x82,
	0x82, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x2d, 0x2e, 0x36, 0x88, 0x36, 0x21, 0x3e,
	0xb0, 0x32, 0xb8, 0x05, 0x52, 0xa8, 0xfc, 0x62, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xa7, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0x7c, 0x22, 0xf5, 0xe1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppServiceClient interface {
	GetApp(ctx context.Context, in *GetAppReq, opts ...grpc.CallOption) (*GetAppRes, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) GetApp(ctx context.Context, in *GetAppReq, opts ...grpc.CallOption) (*GetAppRes, error) {
	out := new(GetAppRes)
	err := c.cc.Invoke(ctx, "/app.AppService/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
type AppServiceServer interface {
	GetApp(context.Context, *GetAppReq) (*GetAppRes, error)
}

// UnimplementedAppServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (*UnimplementedAppServiceServer) GetApp(ctx context.Context, req *GetAppReq) (*GetAppRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}

func RegisterAppServiceServer(s *grpc.Server, srv AppServiceServer) {
	s.RegisterService(&_AppService_serviceDesc, srv)
}

func _AppService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.AppService/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetApp(ctx, req.(*GetAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApp",
			Handler:    _AppService_GetApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
