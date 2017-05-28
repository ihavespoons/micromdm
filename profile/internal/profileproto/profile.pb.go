// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

/*
Package profileproto is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
*/
package profileproto

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

type Profile struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Mobileconfig []byte `protobuf:"bytes,2,opt,name=mobileconfig,proto3" json:"mobileconfig,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetMobileconfig() []byte {
	if m != nil {
		return m.Mobileconfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "profileproto.Profile")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0xcb, 0xcc, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0x72, 0xc1, 0x3c, 0x25,
	0x5b, 0x2e, 0xf6, 0x00, 0x08, 0x5f, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x89, 0x8b, 0x27, 0x37, 0x3f, 0x29, 0x33, 0x27, 0x35,
	0x39, 0x3f, 0x2f, 0x2d, 0x33, 0x5d, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08, 0x45, 0x2c, 0x89,
	0x0d, 0x6c, 0x8a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xea, 0x33, 0x2f, 0xfa, 0x64, 0x00, 0x00,
	0x00,
}