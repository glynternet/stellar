// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/application/v1/application.proto

/*
Package application is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/application/v1/application.proto

It has these top-level messages:
	InfoRequest
	InfoResponse
	CreateRequest
	DeleteRequest
	ListRequest
	App
	ListResponse
	GetRequest
	GetResponse
	RestartRequest
*/
package application

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"
import stellar_services_node_v1 "github.com/ehazlett/stellar/api/services/node/v1"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type InfoRequest struct {
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{0} }

type InfoResponse struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{1} }

func (m *InfoResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type CreateRequest struct {
	Name     string                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels   []string                            `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	Services []*stellar_services_node_v1.Service `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{2} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreateRequest) GetServices() []*stellar_services_node_v1.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type DeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{3} }

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{4} }

type App struct {
	Name     string                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Services []*stellar_services_node_v1.Service `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{5} }

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetServices() []*stellar_services_node_v1.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type ListResponse struct {
	Applications []*App `protobuf:"bytes,1,rep,name=applications" json:"applications,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{6} }

func (m *ListResponse) GetApplications() []*App {
	if m != nil {
		return m.Applications
	}
	return nil
}

type GetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{7} }

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResponse struct {
	Application *App `protobuf:"bytes,1,opt,name=application" json:"application,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{8} }

func (m *GetResponse) GetApplication() *App {
	if m != nil {
		return m.Application
	}
	return nil
}

type RestartRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *RestartRequest) Reset()                    { *m = RestartRequest{} }
func (m *RestartRequest) String() string            { return proto.CompactTextString(m) }
func (*RestartRequest) ProtoMessage()               {}
func (*RestartRequest) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{9} }

func (m *RestartRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*InfoRequest)(nil), "stellar.services.application.v1.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "stellar.services.application.v1.InfoResponse")
	proto.RegisterType((*CreateRequest)(nil), "stellar.services.application.v1.CreateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "stellar.services.application.v1.DeleteRequest")
	proto.RegisterType((*ListRequest)(nil), "stellar.services.application.v1.ListRequest")
	proto.RegisterType((*App)(nil), "stellar.services.application.v1.App")
	proto.RegisterType((*ListResponse)(nil), "stellar.services.application.v1.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "stellar.services.application.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "stellar.services.application.v1.GetResponse")
	proto.RegisterType((*RestartRequest)(nil), "stellar.services.application.v1.RestartRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Application service

type ApplicationClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type applicationClient struct {
	cc *grpc.ClientConn
}

func NewApplicationClient(cc *grpc.ClientConn) ApplicationClient {
	return &applicationClient{cc}
}

func (c *applicationClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationClient) Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.application.v1.Application/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Application service

type ApplicationServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	Create(context.Context, *CreateRequest) (*google_protobuf1.Empty, error)
	Delete(context.Context, *DeleteRequest) (*google_protobuf1.Empty, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Restart(context.Context, *RestartRequest) (*google_protobuf1.Empty, error)
}

func RegisterApplicationServer(s *grpc.Server, srv ApplicationServer) {
	s.RegisterService(&_Application_serviceDesc, srv)
}

func _Application_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Application_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.application.v1.Application/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServer).Restart(ctx, req.(*RestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Application_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.application.v1.Application",
	HandlerType: (*ApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Application_Info_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Application_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Application_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Application_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Application_Get_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Application_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/application/v1/application.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/application/v1/application.proto", fileDescriptorApplication)
}

var fileDescriptorApplication = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0xec, 0x28, 0xd0, 0x71, 0xc2, 0x61, 0x85, 0xa2, 0xc8, 0x1c, 0x1a, 0x4c, 0x85, 0x22,
	0x51, 0x76, 0x95, 0x72, 0xac, 0x38, 0xa4, 0x1f, 0x94, 0x48, 0x1c, 0x90, 0x11, 0x52, 0xc5, 0x09,
	0xc7, 0x99, 0xba, 0x96, 0x1c, 0xef, 0xe2, 0xdd, 0x44, 0xa2, 0xff, 0x8a, 0x3f, 0xc4, 0x81, 0x5f,
	0x82, 0xec, 0xdd, 0xd0, 0xb5, 0x20, 0xb1, 0xe9, 0x29, 0x3b, 0xc9, 0x7b, 0xf3, 0xde, 0xcc, 0xbe,
	0x0d, 0xcc, 0x93, 0x54, 0xdd, 0xae, 0x17, 0x34, 0xe6, 0x2b, 0x86, 0xb7, 0xd1, 0x5d, 0x86, 0x4a,
	0x31, 0xa9, 0x30, 0xcb, 0xa2, 0x82, 0x45, 0x22, 0x65, 0x12, 0x8b, 0x4d, 0x1a, 0xa3, 0x64, 0x91,
	0x10, 0x59, 0x1a, 0x47, 0x2a, 0xe5, 0x39, 0xdb, 0x4c, 0xed, 0x92, 0x8a, 0x82, 0x2b, 0x4e, 0x0e,
	0x0d, 0x8d, 0x6e, 0x29, 0xd4, 0xc6, 0x6c, 0xa6, 0xfe, 0xd3, 0x84, 0x27, 0xbc, 0xc2, 0xb2, 0xf2,
	0xa4, 0x69, 0xfe, 0xb3, 0x84, 0xf3, 0x24, 0x43, 0x56, 0x55, 0x8b, 0xf5, 0x0d, 0xc3, 0x95, 0x50,
	0xdf, 0xcd, 0x8f, 0xa7, 0xad, 0xed, 0xe5, 0x7c, 0x89, 0xa5, 0xaf, 0xf2, 0x53, 0x93, 0x83, 0x01,
	0x78, 0xf3, 0xfc, 0x86, 0x87, 0xf8, 0x6d, 0x8d, 0x52, 0x05, 0x2f, 0xa1, 0xaf, 0x4b, 0x29, 0x78,
	0x2e, 0x91, 0x0c, 0xc1, 0x49, 0x97, 0xa3, 0xce, 0xb8, 0x33, 0x39, 0x38, 0xeb, 0xfd, 0xfa, 0x79,
	0xe8, 0xcc, 0x2f, 0x42, 0x27, 0x5d, 0x06, 0x77, 0x30, 0x38, 0x2f, 0x30, 0x52, 0x68, 0x88, 0x84,
	0x40, 0x37, 0x8f, 0x56, 0xa8, 0xa1, 0x61, 0x75, 0x26, 0x43, 0xe8, 0x65, 0xd1, 0x02, 0x33, 0x39,
	0x72, 0xc6, 0xee, 0xe4, 0x20, 0x34, 0x15, 0x79, 0x0b, 0x8f, 0xb7, 0x96, 0x46, 0xee, 0xd8, 0x9d,
	0x78, 0x27, 0xcf, 0xe9, 0x5f, 0x7b, 0xa9, 0x3c, 0x6e, 0xa6, 0xf4, 0x93, 0xfe, 0x22, 0xfc, 0x43,
	0x09, 0x5e, 0xc0, 0xe0, 0x02, 0x33, 0xdc, 0xab, 0x5d, 0xce, 0xf5, 0x21, 0x95, 0x6a, 0x3b, 0xd7,
	0x35, 0xb8, 0x33, 0x21, 0xfe, 0xe9, 0xd2, 0x76, 0xe3, 0xfc, 0xbf, 0x9b, 0x6b, 0xe8, 0x6b, 0x21,
	0xb3, 0xb1, 0xf7, 0xd0, 0xb7, 0xae, 0x54, 0x8e, 0x3a, 0x55, 0xcb, 0x23, 0xda, 0x70, 0xf1, 0x74,
	0x26, 0x44, 0x58, 0x63, 0x06, 0x63, 0x80, 0x2b, 0x54, 0xfb, 0x86, 0xfc, 0x0c, 0x5e, 0x85, 0x30,
	0xd2, 0xef, 0xc0, 0xb3, 0x1a, 0x54, 0xc8, 0xb6, 0xca, 0x36, 0x31, 0x38, 0x82, 0x27, 0x21, 0x4a,
	0x15, 0x15, 0xfb, 0xc4, 0x4f, 0x7e, 0x74, 0xc1, 0x9b, 0xdd, 0xb3, 0x48, 0x0c, 0xdd, 0x32, 0x3a,
	0xe4, 0xb8, 0x51, 0xd0, 0x0a, 0x9c, 0xff, 0xba, 0x25, 0xda, 0x8c, 0xf8, 0x11, 0x7a, 0x3a, 0x77,
	0x84, 0x36, 0x12, 0x6b, 0x01, 0xf5, 0x87, 0x54, 0xbf, 0x21, 0xba, 0x7d, 0x43, 0xf4, 0xb2, 0x7c,
	0x43, 0x65, 0x47, 0x9d, 0xa6, 0x16, 0x1d, 0x6b, 0xb1, 0xdb, 0xd9, 0x31, 0x86, 0x6e, 0x99, 0x88,
	0x16, 0x8b, 0xb0, 0x12, 0xda, 0x62, 0x11, 0xb5, 0x98, 0x7d, 0x05, 0xf7, 0x0a, 0x15, 0x79, 0xd5,
	0xc8, 0xba, 0x8f, 0x90, 0x7f, 0xdc, 0x0e, 0x6c, 0x14, 0x42, 0x78, 0x64, 0x52, 0x40, 0x58, 0x23,
	0xb1, 0x9e, 0x97, 0x5d, 0xab, 0x39, 0xbb, 0xfc, 0x72, 0xfe, 0xc0, 0xff, 0xd2, 0x53, 0xab, 0x5c,
	0xf4, 0xaa, 0xb6, 0x6f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x23, 0xdb, 0xc6, 0x99, 0x05,
	0x00, 0x00,
}
