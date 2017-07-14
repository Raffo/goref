// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serve.proto

/*
Package serve is a generated protocol buffer package.

It is generated from these files:
	serve.proto

It has these top-level messages:
	GetFileRequest
	GetFileResponse
	GetAnnotationsRequest
	GetAnnotationsResponse
	GetFilesRequest
	GetFilesResponse
*/
package serve

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import goref_proto "github.com/korfuri/goref/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetFileRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *GetFileRequest) Reset()                    { *m = GetFileRequest{} }
func (m *GetFileRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFileRequest) ProtoMessage()               {}
func (*GetFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetFileRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type GetFileResponse struct {
	Path     string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Contents string `protobuf:"bytes,2,opt,name=contents" json:"contents,omitempty"`
}

func (m *GetFileResponse) Reset()                    { *m = GetFileResponse{} }
func (m *GetFileResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFileResponse) ProtoMessage()               {}
func (*GetFileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetFileResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GetFileResponse) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

type GetAnnotationsRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *GetAnnotationsRequest) Reset()                    { *m = GetAnnotationsRequest{} }
func (m *GetAnnotationsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAnnotationsRequest) ProtoMessage()               {}
func (*GetAnnotationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetAnnotationsRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type GetAnnotationsResponse struct {
	Path       string             `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Annotation []*goref_proto.Ref `protobuf:"bytes,2,rep,name=annotation" json:"annotation,omitempty"`
}

func (m *GetAnnotationsResponse) Reset()                    { *m = GetAnnotationsResponse{} }
func (m *GetAnnotationsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAnnotationsResponse) ProtoMessage()               {}
func (*GetAnnotationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetAnnotationsResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GetAnnotationsResponse) GetAnnotation() []*goref_proto.Ref {
	if m != nil {
		return m.Annotation
	}
	return nil
}

type GetFilesRequest struct {
	Package string `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
}

func (m *GetFilesRequest) Reset()                    { *m = GetFilesRequest{} }
func (m *GetFilesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFilesRequest) ProtoMessage()               {}
func (*GetFilesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetFilesRequest) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

type GetFilesResponse struct {
	Package  string   `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
	Filename []string `protobuf:"bytes,2,rep,name=filename" json:"filename,omitempty"`
}

func (m *GetFilesResponse) Reset()                    { *m = GetFilesResponse{} }
func (m *GetFilesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFilesResponse) ProtoMessage()               {}
func (*GetFilesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetFilesResponse) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *GetFilesResponse) GetFilename() []string {
	if m != nil {
		return m.Filename
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFileRequest)(nil), "serve.GetFileRequest")
	proto.RegisterType((*GetFileResponse)(nil), "serve.GetFileResponse")
	proto.RegisterType((*GetAnnotationsRequest)(nil), "serve.GetAnnotationsRequest")
	proto.RegisterType((*GetAnnotationsResponse)(nil), "serve.GetAnnotationsResponse")
	proto.RegisterType((*GetFilesRequest)(nil), "serve.GetFilesRequest")
	proto.RegisterType((*GetFilesResponse)(nil), "serve.GetFilesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Goref service

type GorefClient interface {
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
	GetAnnotations(ctx context.Context, in *GetAnnotationsRequest, opts ...grpc.CallOption) (*GetAnnotationsResponse, error)
	GetFiles(ctx context.Context, in *GetFilesRequest, opts ...grpc.CallOption) (*GetFilesResponse, error)
}

type gorefClient struct {
	cc *grpc.ClientConn
}

func NewGorefClient(cc *grpc.ClientConn) GorefClient {
	return &gorefClient{cc}
}

func (c *gorefClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := grpc.Invoke(ctx, "/serve.Goref/GetFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gorefClient) GetAnnotations(ctx context.Context, in *GetAnnotationsRequest, opts ...grpc.CallOption) (*GetAnnotationsResponse, error) {
	out := new(GetAnnotationsResponse)
	err := grpc.Invoke(ctx, "/serve.Goref/GetAnnotations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gorefClient) GetFiles(ctx context.Context, in *GetFilesRequest, opts ...grpc.CallOption) (*GetFilesResponse, error) {
	out := new(GetFilesResponse)
	err := grpc.Invoke(ctx, "/serve.Goref/GetFiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goref service

type GorefServer interface {
	GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error)
	GetAnnotations(context.Context, *GetAnnotationsRequest) (*GetAnnotationsResponse, error)
	GetFiles(context.Context, *GetFilesRequest) (*GetFilesResponse, error)
}

func RegisterGorefServer(s *grpc.Server, srv GorefServer) {
	s.RegisterService(&_Goref_serviceDesc, srv)
}

func _Goref_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GorefServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serve.Goref/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GorefServer).GetFile(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goref_GetAnnotations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnnotationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GorefServer).GetAnnotations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serve.Goref/GetAnnotations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GorefServer).GetAnnotations(ctx, req.(*GetAnnotationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goref_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GorefServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serve.Goref/GetFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GorefServer).GetFiles(ctx, req.(*GetFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goref_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serve.Goref",
	HandlerType: (*GorefServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFile",
			Handler:    _Goref_GetFile_Handler,
		},
		{
			MethodName: "GetAnnotations",
			Handler:    _Goref_GetAnnotations_Handler,
		},
		{
			MethodName: "GetFiles",
			Handler:    _Goref_GetFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serve.proto",
}

func init() { proto.RegisterFile("serve.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0x43, 0xff, 0x7f, 0x04, 0x86, 0x44, 0xc9, 0x26, 0x60, 0xd3, 0xa0, 0x62, 0x43, 0x22,
	0x81, 0xa4, 0xab, 0x78, 0xf6, 0xc0, 0x45, 0x3c, 0xf7, 0xe8, 0xc1, 0x64, 0x21, 0xd3, 0xd2, 0x00,
	0xbb, 0xb5, 0x5d, 0xb8, 0x18, 0x2f, 0xbe, 0x82, 0xef, 0xe3, 0x4b, 0xf8, 0x0a, 0x3e, 0x88, 0xe9,
	0xb6, 0x6c, 0x0b, 0x08, 0xb7, 0xce, 0xec, 0x97, 0xf9, 0xcd, 0x7c, 0x5f, 0xa1, 0x1e, 0x63, 0xb4,
	0x46, 0x27, 0x8c, 0x84, 0x14, 0xa4, 0xac, 0x0a, 0xab, 0xed, 0x0b, 0xe1, 0x2f, 0x90, 0xb2, 0x30,
	0xa0, 0x8c, 0x73, 0x21, 0x99, 0x0c, 0x04, 0x8f, 0x53, 0x91, 0xd5, 0xf3, 0x03, 0x39, 0x5b, 0x4d,
	0x9c, 0xa9, 0x58, 0xd2, 0xb9, 0x88, 0xbc, 0x55, 0x14, 0x50, 0x5f, 0x44, 0xe8, 0x51, 0xf5, 0x4e,
	0x23, 0xf4, 0x52, 0xa5, 0xdd, 0x85, 0xd3, 0x31, 0xca, 0xc7, 0x60, 0x81, 0x2e, 0xbe, 0xae, 0x30,
	0x96, 0x84, 0xc0, 0xff, 0x90, 0xc9, 0x99, 0x59, 0xea, 0x94, 0x7a, 0x35, 0x57, 0x7d, 0xdb, 0x23,
	0x38, 0xd3, 0xaa, 0x38, 0x14, 0x3c, 0xc6, 0xbf, 0x64, 0xc4, 0x82, 0xea, 0x54, 0x70, 0x89, 0x5c,
	0xc6, 0xa6, 0xa1, 0xfa, 0xba, 0xb6, 0x07, 0xd0, 0x1c, 0xa3, 0x1c, 0xe5, 0xab, 0x1e, 0xe3, 0xbd,
	0x40, 0x6b, 0x57, 0x7c, 0x04, 0x7b, 0x0b, 0x90, 0x5b, 0x60, 0x1a, 0x9d, 0x7f, 0xbd, 0xfa, 0xb0,
	0xe1, 0xa8, 0x7b, 0xd3, 0x2b, 0x1d, 0x17, 0x3d, 0xb7, 0xa0, 0xb1, 0x07, 0xfa, 0x1e, 0xbd, 0x86,
	0x09, 0x95, 0x90, 0x4d, 0xe7, 0xcc, 0xc7, 0x6c, 0xf6, 0xa6, 0xb4, 0x9f, 0xa0, 0x91, 0x8b, 0xb3,
	0x35, 0x0e, 0xaa, 0x13, 0x0f, 0xbc, 0x60, 0x81, 0x9c, 0x2d, 0x51, 0xad, 0x52, 0x73, 0x75, 0x3d,
	0xfc, 0x32, 0xa0, 0x3c, 0x4e, 0xd6, 0x22, 0xcf, 0x50, 0xc9, 0x66, 0x92, 0xa6, 0x93, 0xc6, 0xbb,
	0x1d, 0x83, 0xd5, 0xda, 0x6d, 0xa7, 0x64, 0xbb, 0xf3, 0xf1, 0xfd, 0xf3, 0x69, 0x58, 0xc4, 0xa4,
	0xeb, 0xbb, 0x2c, 0xd3, 0x64, 0x3e, 0x7d, 0x4b, 0xac, 0x78, 0xe8, 0xf7, 0xdf, 0xc9, 0x5a, 0x45,
	0x5a, 0x30, 0x8f, 0xb4, 0xf3, 0x59, 0xfb, 0x01, 0x58, 0x17, 0x07, 0x5e, 0x33, 0xe0, 0x8d, 0x02,
	0x5e, 0x93, 0xab, 0x1c, 0x58, 0xf8, 0xe1, 0x0a, 0x5c, 0x06, 0xd5, 0x8d, 0x4f, 0x64, 0x67, 0x7b,
	0xcd, 0x3a, 0xdf, 0xeb, 0x67, 0x94, 0xae, 0xa2, 0x5c, 0x92, 0xf6, 0xf6, 0x59, 0x6a, 0xbe, 0x32,
	0x36, 0x41, 0x4c, 0x4e, 0x54, 0x9c, 0xf7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x06, 0xfe, 0x88,
	0x7a, 0x12, 0x03, 0x00, 0x00,
}
