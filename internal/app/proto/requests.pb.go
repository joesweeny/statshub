// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/app/proto/requests.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FixtureRequest struct {
	FixtureId            uint64   `protobuf:"varint,1,opt,name=fixture_id,json=fixtureId,proto3" json:"fixture_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FixtureRequest) Reset()         { *m = FixtureRequest{} }
func (m *FixtureRequest) String() string { return proto.CompactTextString(m) }
func (*FixtureRequest) ProtoMessage()    {}
func (*FixtureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94245552992eb109, []int{0}
}

func (m *FixtureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixtureRequest.Unmarshal(m, b)
}
func (m *FixtureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixtureRequest.Marshal(b, m, deterministic)
}
func (m *FixtureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixtureRequest.Merge(m, src)
}
func (m *FixtureRequest) XXX_Size() int {
	return xxx_messageInfo_FixtureRequest.Size(m)
}
func (m *FixtureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FixtureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FixtureRequest proto.InternalMessageInfo

func (m *FixtureRequest) GetFixtureId() uint64 {
	if m != nil {
		return m.FixtureId
	}
	return 0
}

func init() {
	proto.RegisterType((*FixtureRequest)(nil), "proto.FixtureRequest")
}

func init() { proto.RegisterFile("internal/app/proto/requests.proto", fileDescriptor_94245552992eb109) }

var fileDescriptor_94245552992eb109 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0x2c, 0x28, 0xd0, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x29, 0xd6, 0x03, 0x73, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x3e,
	0x17, 0x9f, 0x5b, 0x66, 0x45, 0x49, 0x69, 0x51, 0x6a, 0x10, 0x44, 0x5e, 0x48, 0x96, 0x8b, 0x2b,
	0x0d, 0x22, 0x12, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0xc4, 0x09, 0x15, 0xf1,
	0x4c, 0x49, 0x62, 0x03, 0xeb, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x52, 0xe5, 0x69,
	0x63, 0x00, 0x00, 0x00,
}