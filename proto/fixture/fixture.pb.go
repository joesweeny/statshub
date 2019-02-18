// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/fixture/fixture.proto

package statshub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

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

type Request struct {
	DateFrom             string   `protobuf:"bytes,1,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo               string   `protobuf:"bytes,2,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fixture_e87420d3e76139d7, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetDateFrom() string {
	if m != nil {
		return m.DateFrom
	}
	return ""
}

func (m *Request) GetDateTo() string {
	if m != nil {
		return m.DateTo
	}
	return ""
}

type Competition struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Competition) Reset()         { *m = Competition{} }
func (m *Competition) String() string { return proto.CompactTextString(m) }
func (*Competition) ProtoMessage()    {}
func (*Competition) Descriptor() ([]byte, []int) {
	return fileDescriptor_fixture_e87420d3e76139d7, []int{1}
}
func (m *Competition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Competition.Unmarshal(m, b)
}
func (m *Competition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Competition.Marshal(b, m, deterministic)
}
func (dst *Competition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Competition.Merge(dst, src)
}
func (m *Competition) XXX_Size() int {
	return xxx_messageInfo_Competition.Size(m)
}
func (m *Competition) XXX_DiscardUnknown() {
	xxx_messageInfo_Competition.DiscardUnknown(m)
}

var xxx_messageInfo_Competition proto.InternalMessageInfo

func (m *Competition) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Competition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Season struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Season) Reset()         { *m = Season{} }
func (m *Season) String() string { return proto.CompactTextString(m) }
func (*Season) ProtoMessage()    {}
func (*Season) Descriptor() ([]byte, []int) {
	return fileDescriptor_fixture_e87420d3e76139d7, []int{2}
}
func (m *Season) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Season.Unmarshal(m, b)
}
func (m *Season) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Season.Marshal(b, m, deterministic)
}
func (dst *Season) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Season.Merge(dst, src)
}
func (m *Season) XXX_Size() int {
	return xxx_messageInfo_Season.Size(m)
}
func (m *Season) XXX_DiscardUnknown() {
	xxx_messageInfo_Season.DiscardUnknown(m)
}

var xxx_messageInfo_Season proto.InternalMessageInfo

func (m *Season) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Season) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Team struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_fixture_e87420d3e76139d7, []int{3}
}
func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (dst *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(dst, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Venue struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Venue) Reset()         { *m = Venue{} }
func (m *Venue) String() string { return proto.CompactTextString(m) }
func (*Venue) ProtoMessage()    {}
func (*Venue) Descriptor() ([]byte, []int) {
	return fileDescriptor_fixture_e87420d3e76139d7, []int{4}
}
func (m *Venue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Venue.Unmarshal(m, b)
}
func (m *Venue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Venue.Marshal(b, m, deterministic)
}
func (dst *Venue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Venue.Merge(dst, src)
}
func (m *Venue) XXX_Size() int {
	return xxx_messageInfo_Venue.Size(m)
}
func (m *Venue) XXX_DiscardUnknown() {
	xxx_messageInfo_Venue.DiscardUnknown(m)
}

var xxx_messageInfo_Venue proto.InternalMessageInfo

func (m *Venue) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Venue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Fixture struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Competition          *Competition         `protobuf:"bytes,2,opt,name=competition,proto3" json:"competition,omitempty"`
	Season               *Season              `protobuf:"bytes,3,opt,name=season,proto3" json:"season,omitempty"`
	HomeTeam             *Team                `protobuf:"bytes,4,opt,name=home_team,json=homeTeam,proto3" json:"home_team,omitempty"`
	AwayTeam             *Team                `protobuf:"bytes,5,opt,name=away_team,json=awayTeam,proto3" json:"away_team,omitempty"`
	Venue                *Venue               `protobuf:"bytes,6,opt,name=venue,proto3" json:"venue,omitempty"`
	RefereeId            *wrappers.Int64Value `protobuf:"bytes,7,opt,name=referee_id,json=refereeId,proto3" json:"referee_id,omitempty"`
	DateTime             int64                `protobuf:"varint,8,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Fixture) Reset()         { *m = Fixture{} }
func (m *Fixture) String() string { return proto.CompactTextString(m) }
func (*Fixture) ProtoMessage()    {}
func (*Fixture) Descriptor() ([]byte, []int) {
	return fileDescriptor_fixture_e87420d3e76139d7, []int{5}
}
func (m *Fixture) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fixture.Unmarshal(m, b)
}
func (m *Fixture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fixture.Marshal(b, m, deterministic)
}
func (dst *Fixture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fixture.Merge(dst, src)
}
func (m *Fixture) XXX_Size() int {
	return xxx_messageInfo_Fixture.Size(m)
}
func (m *Fixture) XXX_DiscardUnknown() {
	xxx_messageInfo_Fixture.DiscardUnknown(m)
}

var xxx_messageInfo_Fixture proto.InternalMessageInfo

func (m *Fixture) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Fixture) GetCompetition() *Competition {
	if m != nil {
		return m.Competition
	}
	return nil
}

func (m *Fixture) GetSeason() *Season {
	if m != nil {
		return m.Season
	}
	return nil
}

func (m *Fixture) GetHomeTeam() *Team {
	if m != nil {
		return m.HomeTeam
	}
	return nil
}

func (m *Fixture) GetAwayTeam() *Team {
	if m != nil {
		return m.AwayTeam
	}
	return nil
}

func (m *Fixture) GetVenue() *Venue {
	if m != nil {
		return m.Venue
	}
	return nil
}

func (m *Fixture) GetRefereeId() *wrappers.Int64Value {
	if m != nil {
		return m.RefereeId
	}
	return nil
}

func (m *Fixture) GetDateTime() int64 {
	if m != nil {
		return m.DateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "statshub.Request")
	proto.RegisterType((*Competition)(nil), "statshub.Competition")
	proto.RegisterType((*Season)(nil), "statshub.Season")
	proto.RegisterType((*Team)(nil), "statshub.Team")
	proto.RegisterType((*Venue)(nil), "statshub.Venue")
	proto.RegisterType((*Fixture)(nil), "statshub.Fixture")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FixtureServiceClient is the client API for FixtureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FixtureServiceClient interface {
	ListFixtures(ctx context.Context, in *Request, opts ...grpc.CallOption) (FixtureService_ListFixturesClient, error)
}

type fixtureServiceClient struct {
	cc *grpc.ClientConn
}

func NewFixtureServiceClient(cc *grpc.ClientConn) FixtureServiceClient {
	return &fixtureServiceClient{cc}
}

func (c *fixtureServiceClient) ListFixtures(ctx context.Context, in *Request, opts ...grpc.CallOption) (FixtureService_ListFixturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FixtureService_serviceDesc.Streams[0], "/statshub.FixtureService/ListFixtures", opts...)
	if err != nil {
		return nil, err
	}
	x := &fixtureServiceListFixturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FixtureService_ListFixturesClient interface {
	Recv() (*Fixture, error)
	grpc.ClientStream
}

type fixtureServiceListFixturesClient struct {
	grpc.ClientStream
}

func (x *fixtureServiceListFixturesClient) Recv() (*Fixture, error) {
	m := new(Fixture)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FixtureServiceServer is the server API for FixtureService service.
type FixtureServiceServer interface {
	ListFixtures(*Request, FixtureService_ListFixturesServer) error
}

func RegisterFixtureServiceServer(s *grpc.Server, srv FixtureServiceServer) {
	s.RegisterService(&_FixtureService_serviceDesc, srv)
}

func _FixtureService_ListFixtures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FixtureServiceServer).ListFixtures(m, &fixtureServiceListFixturesServer{stream})
}

type FixtureService_ListFixturesServer interface {
	Send(*Fixture) error
	grpc.ServerStream
}

type fixtureServiceListFixturesServer struct {
	grpc.ServerStream
}

func (x *fixtureServiceListFixturesServer) Send(m *Fixture) error {
	return x.ServerStream.SendMsg(m)
}

var _FixtureService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "statshub.FixtureService",
	HandlerType: (*FixtureServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFixtures",
			Handler:       _FixtureService_ListFixtures_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/fixture/fixture.proto",
}

func init() {
	proto.RegisterFile("proto/fixture/fixture.proto", fileDescriptor_fixture_e87420d3e76139d7)
}

var fileDescriptor_fixture_e87420d3e76139d7 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0xab, 0xd3, 0x40,
	0x14, 0xc5, 0x6d, 0x5e, 0x9b, 0x3f, 0xb7, 0x52, 0x75, 0x40, 0x0c, 0xaf, 0x20, 0x12, 0x10, 0x1e,
	0x56, 0x52, 0x7d, 0x8a, 0x8a, 0x1b, 0x17, 0x42, 0xa1, 0xe2, 0x6a, 0x5a, 0xba, 0x2d, 0xd3, 0xe6,
	0xa6, 0x1d, 0xe8, 0x64, 0xe2, 0xcc, 0xa4, 0xd5, 0x2f, 0xea, 0xe7, 0x91, 0xdc, 0xa4, 0xa6, 0xa0,
	0x8b, 0xae, 0xda, 0x9c, 0xfb, 0x3b, 0x30, 0xe7, 0x70, 0x60, 0x5c, 0x1a, 0xed, 0xf4, 0x34, 0x97,
	0x3f, 0x5d, 0x65, 0xf0, 0xfc, 0x9b, 0x92, 0xca, 0x42, 0xeb, 0x84, 0xb3, 0xfb, 0x6a, 0x73, 0xfb,
	0x7c, 0xa7, 0xf5, 0xee, 0x80, 0x53, 0xd2, 0x37, 0x55, 0x3e, 0x3d, 0x19, 0x51, 0x96, 0x68, 0x6c,
	0x43, 0x26, 0x5f, 0x20, 0xe0, 0xf8, 0xa3, 0x42, 0xeb, 0xd8, 0x18, 0xa2, 0x4c, 0x38, 0x5c, 0xe7,
	0x46, 0xab, 0xb8, 0xf7, 0xa2, 0x77, 0x17, 0xf1, 0xb0, 0x16, 0x66, 0x46, 0x2b, 0xf6, 0x0c, 0x02,
	0x3a, 0x3a, 0x1d, 0x7b, 0x74, 0xf2, 0xeb, 0xcf, 0xa5, 0x4e, 0xde, 0xc2, 0xf0, 0xab, 0x56, 0x25,
	0x3a, 0xe9, 0xa4, 0x2e, 0xd8, 0x08, 0x3c, 0x99, 0x91, 0xfb, 0x86, 0x7b, 0x32, 0x63, 0x0c, 0xfa,
	0x85, 0x50, 0xd8, 0x9a, 0xe8, 0x7f, 0xf2, 0x1a, 0xfc, 0x05, 0x0a, 0x7b, 0x25, 0xfd, 0x0a, 0xfa,
	0x4b, 0x14, 0xea, 0x2a, 0x76, 0x02, 0x83, 0x15, 0x16, 0x15, 0x5e, 0x05, 0xff, 0xf6, 0x20, 0x98,
	0x35, 0xb5, 0xfd, 0xc3, 0x7f, 0x84, 0xe1, 0xb6, 0x4b, 0x45, 0xb6, 0xe1, 0xfd, 0xd3, 0xf4, 0x5c,
	0x6b, 0x7a, 0x11, 0x99, 0x5f, 0x92, 0xec, 0x0e, 0x7c, 0x4b, 0xd9, 0xe2, 0x1b, 0xf2, 0x3c, 0xee,
	0x3c, 0x4d, 0x66, 0xde, 0xde, 0xd9, 0x04, 0xa2, 0xbd, 0x56, 0xb8, 0x76, 0x28, 0x54, 0xdc, 0x27,
	0x78, 0xd4, 0xc1, 0x75, 0x64, 0x1e, 0xd6, 0x00, 0x85, 0x9f, 0x40, 0x24, 0x4e, 0xe2, 0x57, 0x03,
	0x0f, 0xfe, 0x0f, 0xd7, 0x00, 0xc1, 0x2f, 0x61, 0x70, 0xac, 0x5b, 0x88, 0x7d, 0x02, 0x1f, 0x75,
	0x20, 0x95, 0xc3, 0x9b, 0x2b, 0xfb, 0x0c, 0x60, 0x30, 0x47, 0x83, 0xb8, 0x96, 0x59, 0x1c, 0x10,
	0x3b, 0x4e, 0x9b, 0xbd, 0xa4, 0xe7, 0xbd, 0xa4, 0xf3, 0xc2, 0x7d, 0x78, 0xbf, 0x12, 0x87, 0x0a,
	0x79, 0xd4, 0xe2, 0xf3, 0xec, 0xef, 0x56, 0x9c, 0x54, 0x18, 0x87, 0x54, 0x1b, 0x6d, 0x65, 0x29,
	0x15, 0xde, 0x7f, 0x83, 0x51, 0xdb, 0xeb, 0x02, 0xcd, 0x51, 0x6e, 0x91, 0x7d, 0x82, 0x87, 0xdf,
	0xa5, 0x75, 0xad, 0x6a, 0xd9, 0x93, 0xee, 0x49, 0xed, 0xfa, 0x6e, 0x2f, 0xa4, 0x16, 0x4b, 0x1e,
	0xbc, 0xe9, 0x6d, 0x7c, 0x7a, 0xc8, 0xbb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x71, 0xa3,
	0x76, 0xef, 0x02, 0x00, 0x00,
}
