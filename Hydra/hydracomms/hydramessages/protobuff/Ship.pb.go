// Code generated by protoc-gen-go.
// source: Ship.proto
// DO NOT EDIT!

/*
Package protobuff is a generated protocol buffer package.

It is generated from these files:
	Ship.proto

It has these top-level messages:
	Ship
*/
package protobuff

//import proto "github.com/golang/protobuf/proto"
import "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Ship struct {
	Shipname    string             `protobuf:"bytes,1,opt,name=shipname" json:"shipname,omitempty"`
	CaptainName string             `protobuf:"bytes,2,opt,name=CaptainName,json=captainName" json:"CaptainName,omitempty"`
	Crew        []*Ship_CrewMember `protobuf:"bytes,3,rep,name=Crew,json=crew" json:"Crew,omitempty"`
}

func (m *Ship) Reset()                    { *m = Ship{} }
func (m *Ship) String() string            { return proto.CompactTextString(m) }
func (*Ship) ProtoMessage()               {}
func (*Ship) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ship) GetCrew() []*Ship_CrewMember {
	if m != nil {
		return m.Crew
	}
	return nil
}

type Ship_CrewMember struct {
	Id           int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SecClearance int32  `protobuf:"varint,3,opt,name=secClearance" json:"secClearance,omitempty"`
	Position     string `protobuf:"bytes,4,opt,name=position" json:"position,omitempty"`
}

func (m *Ship_CrewMember) Reset()                    { *m = Ship_CrewMember{} }
func (m *Ship_CrewMember) String() string            { return proto.CompactTextString(m) }
func (*Ship_CrewMember) ProtoMessage()               {}
func (*Ship_CrewMember) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func init() {
	proto.RegisterType((*Ship)(nil), "protobuff.Ship")
	proto.RegisterType((*Ship_CrewMember)(nil), "protobuff.Ship.CrewMember")
}

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x8e, 0x83, 0x30,
	0x10, 0x45, 0x05, 0x78, 0x57, 0xcb, 0xb0, 0x4a, 0x31, 0x15, 0xa2, 0x42, 0x54, 0x54, 0x2e, 0x92,
	0x23, 0x50, 0x27, 0x05, 0x39, 0x81, 0x71, 0x06, 0xc5, 0x52, 0xb0, 0x2d, 0x43, 0x94, 0x1b, 0xe7,
	0x1c, 0xb1, 0x1d, 0x05, 0x94, 0x6a, 0xfe, 0xfc, 0x67, 0xe9, 0x79, 0x00, 0xce, 0x57, 0x65, 0xb9,
	0x75, 0x66, 0x31, 0x98, 0xc7, 0x31, 0xdc, 0xc7, 0xb1, 0x79, 0x26, 0xc0, 0x02, 0xc1, 0x0a, 0xfe,
	0x66, 0x3f, 0xb5, 0x98, 0xa8, 0x4c, 0xea, 0xa4, 0xcd, 0xfb, 0x75, 0xc7, 0x1a, 0x8a, 0x4e, 0xd8,
	0x45, 0x28, 0x7d, 0x0a, 0x38, 0x8d, 0xb8, 0x90, 0x5b, 0x85, 0x1c, 0x58, 0xe7, 0xe8, 0x51, 0x66,
	0x75, 0xd6, 0x16, 0xfb, 0x8a, 0xaf, 0x02, 0x1e, 0xb5, 0x81, 0x1d, 0x69, 0x1a, 0xc8, 0xf5, 0x4c,
	0xfa, 0x5c, 0x59, 0x80, 0xad, 0xc3, 0x1d, 0xa4, 0xea, 0x12, 0xad, 0x3f, 0xbd, 0x4f, 0x88, 0xc0,
	0xf4, 0x26, 0x8a, 0x19, 0x1b, 0xf8, 0x9f, 0x49, 0x76, 0x37, 0x12, 0x4e, 0x68, 0x49, 0xde, 0x14,
	0x5e, 0x7f, 0x75, 0xe1, 0x06, 0x6b, 0x66, 0xb5, 0x28, 0xa3, 0x4b, 0xf6, 0xbe, 0xe1, 0xb3, 0x0f,
	0xbf, 0xf1, 0x4b, 0x87, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xf7, 0x92, 0x55, 0x08, 0x01,
	0x00, 0x00,
}
