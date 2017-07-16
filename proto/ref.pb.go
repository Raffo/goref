// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ref.proto

/*
Package goref is a generated protocol buffer package.

It is generated from these files:
	ref.proto

It has these top-level messages:
	Ref
	Location
	Position
*/
package goref

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Type int32

const (
	Type_Instantiation  Type = 0
	Type_Call           Type = 1
	Type_Implementation Type = 2
	Type_Extension      Type = 3
	Type_Import         Type = 4
	Type_Reference      Type = 5
)

var Type_name = map[int32]string{
	0: "Instantiation",
	1: "Call",
	2: "Implementation",
	3: "Extension",
	4: "Import",
	5: "Reference",
}
var Type_value = map[string]int32{
	"Instantiation":  0,
	"Call":           1,
	"Implementation": 2,
	"Extension":      3,
	"Import":         4,
	"Reference":      5,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Ref struct {
	Version int64     `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	From    *Location `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	To      *Location `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	Type    Type      `protobuf:"varint,4,opt,name=type,enum=goref.Type" json:"type,omitempty"`
}

func (m *Ref) Reset()                    { *m = Ref{} }
func (m *Ref) String() string            { return proto.CompactTextString(m) }
func (*Ref) ProtoMessage()               {}
func (*Ref) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ref) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ref) GetFrom() *Location {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Ref) GetTo() *Location {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Ref) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_Instantiation
}

type Location struct {
	Position *Position `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	Package  string    `protobuf:"bytes,2,opt,name=package" json:"package,omitempty"`
	Ident    string    `protobuf:"bytes,3,opt,name=ident" json:"ident,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Location) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Location) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Location) GetIdent() string {
	if m != nil {
		return m.Ident
	}
	return ""
}

type Position struct {
	Filename  string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	StartLine int32  `protobuf:"varint,2,opt,name=start_line,json=startLine" json:"start_line,omitempty"`
	StartCol  int32  `protobuf:"varint,3,opt,name=start_col,json=startCol" json:"start_col,omitempty"`
	EndLine   int32  `protobuf:"varint,4,opt,name=end_line,json=endLine" json:"end_line,omitempty"`
	EndCol    int32  `protobuf:"varint,5,opt,name=end_col,json=endCol" json:"end_col,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Position) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Position) GetStartLine() int32 {
	if m != nil {
		return m.StartLine
	}
	return 0
}

func (m *Position) GetStartCol() int32 {
	if m != nil {
		return m.StartCol
	}
	return 0
}

func (m *Position) GetEndLine() int32 {
	if m != nil {
		return m.EndLine
	}
	return 0
}

func (m *Position) GetEndCol() int32 {
	if m != nil {
		return m.EndCol
	}
	return 0
}

func init() {
	proto.RegisterType((*Ref)(nil), "goref.Ref")
	proto.RegisterType((*Location)(nil), "goref.Location")
	proto.RegisterType((*Position)(nil), "goref.Position")
	proto.RegisterEnum("goref.Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("ref.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xab, 0x7f, 0xb6, 0x34, 0xc6, 0xae, 0x3a, 0x14, 0xaa, 0xb6, 0x14, 0x1b, 0xf7, 0x62,
	0x5a, 0xf0, 0xc1, 0x7d, 0x04, 0xd3, 0x83, 0xc1, 0x87, 0xb0, 0xe4, 0x6e, 0x36, 0xd2, 0xc8, 0x2c,
	0x59, 0xed, 0x8a, 0xd5, 0x12, 0xe2, 0x17, 0xc8, 0x23, 0xe4, 0x79, 0x83, 0x46, 0x96, 0x73, 0xc9,
	0x6d, 0xbf, 0xf9, 0xcd, 0x37, 0xfb, 0x31, 0x03, 0x99, 0xa3, 0x7a, 0xdb, 0x3a, 0xeb, 0x2d, 0x26,
	0x67, 0xeb, 0xa8, 0x5e, 0xbf, 0x04, 0x10, 0x09, 0xaa, 0xb1, 0x80, 0xe9, 0x13, 0xb9, 0x4e, 0x59,
	0x53, 0x04, 0xab, 0x60, 0x13, 0x89, 0x51, 0xe2, 0x6f, 0x88, 0x6b, 0x67, 0x9b, 0x22, 0x5c, 0x05,
	0x9b, 0xd9, 0xee, 0xf3, 0x96, 0x7d, 0xdb, 0xa3, 0x2d, 0xa5, 0x57, 0xd6, 0x08, 0x86, 0xb8, 0x84,
	0xd0, 0xdb, 0x22, 0xfa, 0xb8, 0x25, 0xf4, 0x16, 0x97, 0x10, 0xfb, 0x4b, 0x4b, 0x45, 0xbc, 0x0a,
	0x36, 0x8b, 0xdd, 0xec, 0xda, 0x72, 0x7f, 0x69, 0x49, 0x30, 0x58, 0x9f, 0x21, 0x1d, 0x0d, 0xf8,
	0x17, 0xd2, 0xd6, 0x76, 0xca, 0x8f, 0x69, 0xde, 0x67, 0xde, 0x5d, 0xcb, 0xe2, 0xd6, 0xd0, 0x27,
	0x6f, 0x65, 0xf9, 0x28, 0xcf, 0xc4, 0x11, 0x33, 0x31, 0x4a, 0xfc, 0x0a, 0x89, 0xaa, 0xc8, 0x78,
	0xce, 0x95, 0x89, 0x41, 0xac, 0x5f, 0x03, 0x48, 0xc7, 0x31, 0xf8, 0x03, 0xd2, 0x5a, 0x69, 0x32,
	0xb2, 0x21, 0xfe, 0x29, 0x13, 0x37, 0x8d, 0xbf, 0x00, 0x3a, 0x2f, 0x9d, 0x3f, 0x69, 0x65, 0x86,
	0xd9, 0x89, 0xc8, 0xb8, 0x72, 0x54, 0x86, 0xf0, 0x27, 0x0c, 0xe2, 0x54, 0x5a, 0xcd, 0x3f, 0x24,
	0x22, 0xe5, 0xc2, 0xde, 0x6a, 0xfc, 0x0e, 0x29, 0x99, 0x6a, 0x70, 0xc6, 0xcc, 0xa6, 0x64, 0x2a,
	0xf6, 0x7d, 0x83, 0xfe, 0xc9, 0xae, 0x84, 0xc9, 0x84, 0x4c, 0xb5, 0xb7, 0xfa, 0x8f, 0x84, 0xb8,
	0xdf, 0x07, 0x7e, 0x81, 0xf9, 0xc1, 0x74, 0x5e, 0x1a, 0xaf, 0x78, 0x1d, 0xf9, 0x27, 0x4c, 0x21,
	0xde, 0x4b, 0xad, 0xf3, 0x00, 0x11, 0x16, 0x87, 0xa6, 0xd5, 0xd4, 0x90, 0xf1, 0x03, 0x0d, 0x71,
	0x0e, 0xd9, 0xff, 0x67, 0x4f, 0xa6, 0x3f, 0x57, 0x1e, 0x21, 0xc0, 0xe4, 0xd0, 0xb4, 0xd6, 0xf9,
	0x3c, 0xee, 0x91, 0xa0, 0x9a, 0x1c, 0x99, 0x92, 0xf2, 0xe4, 0x61, 0xc2, 0xb7, 0xff, 0xf7, 0x16,
	0x00, 0x00, 0xff, 0xff, 0xca, 0x63, 0x79, 0xe7, 0x08, 0x02, 0x00, 0x00,
}
