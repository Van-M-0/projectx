// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Helloworld
	Person
*/
package protocol

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

type Helloworld struct {
	Id  int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Str string `protobuf:"bytes,2,opt,name=str" json:"str,omitempty"`
	Opt int32  `protobuf:"varint,3,opt,name=opt" json:"opt,omitempty"`
}

func (m *Helloworld) Reset()                    { *m = Helloworld{} }
func (m *Helloworld) String() string            { return proto.CompactTextString(m) }
func (*Helloworld) ProtoMessage()               {}
func (*Helloworld) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Helloworld) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Helloworld) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Helloworld) GetOpt() int32 {
	if m != nil {
		return m.Opt
	}
	return 0
}

type Person struct {
	Len  int32  `protobuf:"varint,1,opt,name=len" json:"len,omitempty"`
	Info string `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Person) GetLen() int32 {
	if m != nil {
		return m.Len
	}
	return 0
}

func (m *Person) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*Helloworld)(nil), "protocol.helloworld")
	proto.RegisterType((*Person)(nil), "protocol.person")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x0e, 0x5c,
	0x5c, 0x19, 0xa9, 0x39, 0x39, 0xf9, 0xe5, 0xf9, 0x45, 0x39, 0x29, 0x42, 0x7c, 0x5c, 0x4c, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x40, 0x96, 0x90, 0x00, 0x17, 0x73, 0x71, 0x49,
	0x91, 0x04, 0x13, 0x50, 0x80, 0x33, 0x08, 0xc4, 0x04, 0x89, 0xe4, 0x17, 0x94, 0x48, 0x30, 0x83,
	0x95, 0x80, 0x98, 0x4a, 0x7a, 0x5c, 0x6c, 0x05, 0xa9, 0x45, 0xc5, 0xf9, 0x79, 0x20, 0xb9, 0x9c,
	0xd4, 0x3c, 0xa8, 0x76, 0x10, 0x53, 0x48, 0x88, 0x8b, 0x25, 0x33, 0x2f, 0x2d, 0x1f, 0x6a, 0x00,
	0x98, 0x9d, 0xc4, 0x06, 0xb6, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x60, 0xfe, 0xba,
	0x90, 0x00, 0x00, 0x00,
}
