// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/result/result.proto

package statistico_data

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

type ResultRequest struct {
	// The Team ID that the Result set relates to
	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The number of results to return. If limit is not set the whole Result set for the Team
	// will be returned
	Limit                *wrappers.Int32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResultRequest) Reset()         { *m = ResultRequest{} }
func (m *ResultRequest) String() string { return proto.CompactTextString(m) }
func (*ResultRequest) ProtoMessage()    {}
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_5ba2b741c50e2aef, []int{0}
}
func (m *ResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultRequest.Unmarshal(m, b)
}
func (m *ResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultRequest.Marshal(b, m, deterministic)
}
func (dst *ResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultRequest.Merge(dst, src)
}
func (m *ResultRequest) XXX_Size() int {
	return xxx_messageInfo_ResultRequest.Size(m)
}
func (m *ResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResultRequest proto.InternalMessageInfo

func (m *ResultRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *ResultRequest) GetLimit() *wrappers.Int32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

type Result struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Competition          *Competition `protobuf:"bytes,2,opt,name=competition,proto3" json:"competition,omitempty"`
	Season               *Season      `protobuf:"bytes,3,opt,name=season,proto3" json:"season,omitempty"`
	Venue                *Venue       `protobuf:"bytes,4,opt,name=venue,proto3" json:"venue,omitempty"`
	MatchData            *MatchData   `protobuf:"bytes,5,opt,name=match_data,json=matchData,proto3" json:"match_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_5ba2b741c50e2aef, []int{1}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Result) GetCompetition() *Competition {
	if m != nil {
		return m.Competition
	}
	return nil
}

func (m *Result) GetSeason() *Season {
	if m != nil {
		return m.Season
	}
	return nil
}

func (m *Result) GetVenue() *Venue {
	if m != nil {
		return m.Venue
	}
	return nil
}

func (m *Result) GetMatchData() *MatchData {
	if m != nil {
		return m.MatchData
	}
	return nil
}

type MatchData struct {
	HomeTeam             *Team       `protobuf:"bytes,1,opt,name=home_team,json=homeTeam,proto3" json:"home_team,omitempty"`
	AwayTeam             *Team       `protobuf:"bytes,2,opt,name=away_team,json=awayTeam,proto3" json:"away_team,omitempty"`
	Stats                *MatchStats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MatchData) Reset()         { *m = MatchData{} }
func (m *MatchData) String() string { return proto.CompactTextString(m) }
func (*MatchData) ProtoMessage()    {}
func (*MatchData) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_5ba2b741c50e2aef, []int{2}
}
func (m *MatchData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchData.Unmarshal(m, b)
}
func (m *MatchData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchData.Marshal(b, m, deterministic)
}
func (dst *MatchData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchData.Merge(dst, src)
}
func (m *MatchData) XXX_Size() int {
	return xxx_messageInfo_MatchData.Size(m)
}
func (m *MatchData) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchData.DiscardUnknown(m)
}

var xxx_messageInfo_MatchData proto.InternalMessageInfo

func (m *MatchData) GetHomeTeam() *Team {
	if m != nil {
		return m.HomeTeam
	}
	return nil
}

func (m *MatchData) GetAwayTeam() *Team {
	if m != nil {
		return m.AwayTeam
	}
	return nil
}

func (m *MatchData) GetStats() *MatchStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type MatchStats struct {
	Pitch                *wrappers.StringValue `protobuf:"bytes,1,opt,name=pitch,proto3" json:"pitch,omitempty"`
	HomeFormation        *wrappers.StringValue `protobuf:"bytes,2,opt,name=home_formation,json=homeFormation,proto3" json:"home_formation,omitempty"`
	AwayFormation        *wrappers.StringValue `protobuf:"bytes,3,opt,name=away_formation,json=awayFormation,proto3" json:"away_formation,omitempty"`
	HomeScore            int32                 `protobuf:"varint,4,opt,name=home_score,json=homeScore,proto3" json:"home_score,omitempty"`
	AwayScore            int32                 `protobuf:"varint,5,opt,name=away_score,json=awayScore,proto3" json:"away_score,omitempty"`
	HomePenScore         int32                 `protobuf:"varint,6,opt,name=home_pen_score,json=homePenScore,proto3" json:"home_pen_score,omitempty"`
	AwayPenScore         int32                 `protobuf:"varint,7,opt,name=away_pen_score,json=awayPenScore,proto3" json:"away_pen_score,omitempty"`
	HalfTimeScore        int32                 `protobuf:"varint,8,opt,name=half_time_score,json=halfTimeScore,proto3" json:"half_time_score,omitempty"`
	FullTimeScore        int32                 `protobuf:"varint,9,opt,name=full_time_score,json=fullTimeScore,proto3" json:"full_time_score,omitempty"`
	ExtraTimeScore       int32                 `protobuf:"varint,10,opt,name=extra_time_score,json=extraTimeScore,proto3" json:"extra_time_score,omitempty"`
	HomeLeaguePosition   *wrappers.Int32Value  `protobuf:"bytes,11,opt,name=home_league_position,json=homeLeaguePosition,proto3" json:"home_league_position,omitempty"`
	AwayLeaguePosition   *wrappers.Int32Value  `protobuf:"bytes,12,opt,name=away_league_position,json=awayLeaguePosition,proto3" json:"away_league_position,omitempty"`
	Minutes              *wrappers.Int32Value  `protobuf:"bytes,13,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Seconds              *wrappers.Int32Value  `protobuf:"bytes,14,opt,name=seconds,proto3" json:"seconds,omitempty"`
	AddedTime            *wrappers.Int32Value  `protobuf:"bytes,15,opt,name=added_time,json=addedTime,proto3" json:"added_time,omitempty"`
	ExtraTime            *wrappers.Int32Value  `protobuf:"bytes,16,opt,name=extra_time,json=extraTime,proto3" json:"extra_time,omitempty"`
	InjuryTime           *wrappers.Int32Value  `protobuf:"bytes,17,opt,name=injury_time,json=injuryTime,proto3" json:"injury_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MatchStats) Reset()         { *m = MatchStats{} }
func (m *MatchStats) String() string { return proto.CompactTextString(m) }
func (*MatchStats) ProtoMessage()    {}
func (*MatchStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_5ba2b741c50e2aef, []int{3}
}
func (m *MatchStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchStats.Unmarshal(m, b)
}
func (m *MatchStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchStats.Marshal(b, m, deterministic)
}
func (dst *MatchStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchStats.Merge(dst, src)
}
func (m *MatchStats) XXX_Size() int {
	return xxx_messageInfo_MatchStats.Size(m)
}
func (m *MatchStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchStats.DiscardUnknown(m)
}

var xxx_messageInfo_MatchStats proto.InternalMessageInfo

func (m *MatchStats) GetPitch() *wrappers.StringValue {
	if m != nil {
		return m.Pitch
	}
	return nil
}

func (m *MatchStats) GetHomeFormation() *wrappers.StringValue {
	if m != nil {
		return m.HomeFormation
	}
	return nil
}

func (m *MatchStats) GetAwayFormation() *wrappers.StringValue {
	if m != nil {
		return m.AwayFormation
	}
	return nil
}

func (m *MatchStats) GetHomeScore() int32 {
	if m != nil {
		return m.HomeScore
	}
	return 0
}

func (m *MatchStats) GetAwayScore() int32 {
	if m != nil {
		return m.AwayScore
	}
	return 0
}

func (m *MatchStats) GetHomePenScore() int32 {
	if m != nil {
		return m.HomePenScore
	}
	return 0
}

func (m *MatchStats) GetAwayPenScore() int32 {
	if m != nil {
		return m.AwayPenScore
	}
	return 0
}

func (m *MatchStats) GetHalfTimeScore() int32 {
	if m != nil {
		return m.HalfTimeScore
	}
	return 0
}

func (m *MatchStats) GetFullTimeScore() int32 {
	if m != nil {
		return m.FullTimeScore
	}
	return 0
}

func (m *MatchStats) GetExtraTimeScore() int32 {
	if m != nil {
		return m.ExtraTimeScore
	}
	return 0
}

func (m *MatchStats) GetHomeLeaguePosition() *wrappers.Int32Value {
	if m != nil {
		return m.HomeLeaguePosition
	}
	return nil
}

func (m *MatchStats) GetAwayLeaguePosition() *wrappers.Int32Value {
	if m != nil {
		return m.AwayLeaguePosition
	}
	return nil
}

func (m *MatchStats) GetMinutes() *wrappers.Int32Value {
	if m != nil {
		return m.Minutes
	}
	return nil
}

func (m *MatchStats) GetSeconds() *wrappers.Int32Value {
	if m != nil {
		return m.Seconds
	}
	return nil
}

func (m *MatchStats) GetAddedTime() *wrappers.Int32Value {
	if m != nil {
		return m.AddedTime
	}
	return nil
}

func (m *MatchStats) GetExtraTime() *wrappers.Int32Value {
	if m != nil {
		return m.ExtraTime
	}
	return nil
}

func (m *MatchStats) GetInjuryTime() *wrappers.Int32Value {
	if m != nil {
		return m.InjuryTime
	}
	return nil
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
	return fileDescriptor_result_5ba2b741c50e2aef, []int{4}
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

type Competition struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsCup                int32    `protobuf:"varint,3,opt,name=is_cup,json=isCup,proto3" json:"is_cup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Competition) Reset()         { *m = Competition{} }
func (m *Competition) String() string { return proto.CompactTextString(m) }
func (*Competition) ProtoMessage()    {}
func (*Competition) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_5ba2b741c50e2aef, []int{5}
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

func (m *Competition) GetIsCup() int32 {
	if m != nil {
		return m.IsCup
	}
	return 0
}

type Season struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsCurrent            int32    `protobuf:"varint,3,opt,name=is_current,json=isCurrent,proto3" json:"is_current,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Season) Reset()         { *m = Season{} }
func (m *Season) String() string { return proto.CompactTextString(m) }
func (*Season) ProtoMessage()    {}
func (*Season) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_5ba2b741c50e2aef, []int{6}
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

func (m *Season) GetIsCurrent() int32 {
	if m != nil {
		return m.IsCurrent
	}
	return 0
}

type Venue struct {
	Id                   *wrappers.Int64Value  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Venue) Reset()         { *m = Venue{} }
func (m *Venue) String() string { return proto.CompactTextString(m) }
func (*Venue) ProtoMessage()    {}
func (*Venue) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_5ba2b741c50e2aef, []int{7}
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

func (m *Venue) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Venue) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*ResultRequest)(nil), "statistico_data.ResultRequest")
	proto.RegisterType((*Result)(nil), "statistico_data.Result")
	proto.RegisterType((*MatchData)(nil), "statistico_data.MatchData")
	proto.RegisterType((*MatchStats)(nil), "statistico_data.MatchStats")
	proto.RegisterType((*Team)(nil), "statistico_data.Team")
	proto.RegisterType((*Competition)(nil), "statistico_data.Competition")
	proto.RegisterType((*Season)(nil), "statistico_data.Season")
	proto.RegisterType((*Venue)(nil), "statistico_data.Venue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResultServiceClient is the client API for ResultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResultServiceClient interface {
	GetResults(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (ResultService_GetResultsClient, error)
}

type resultServiceClient struct {
	cc *grpc.ClientConn
}

func NewResultServiceClient(cc *grpc.ClientConn) ResultServiceClient {
	return &resultServiceClient{cc}
}

func (c *resultServiceClient) GetResults(ctx context.Context, in *ResultRequest, opts ...grpc.CallOption) (ResultService_GetResultsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ResultService_serviceDesc.Streams[0], "/statistico_data.ResultService/GetResults", opts...)
	if err != nil {
		return nil, err
	}
	x := &resultServiceGetResultsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResultService_GetResultsClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type resultServiceGetResultsClient struct {
	grpc.ClientStream
}

func (x *resultServiceGetResultsClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResultServiceServer is the server API for ResultService service.
type ResultServiceServer interface {
	GetResults(*ResultRequest, ResultService_GetResultsServer) error
}

func RegisterResultServiceServer(s *grpc.Server, srv ResultServiceServer) {
	s.RegisterService(&_ResultService_serviceDesc, srv)
}

func _ResultService_GetResults_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResultRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResultServiceServer).GetResults(m, &resultServiceGetResultsServer{stream})
}

type ResultService_GetResultsServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type resultServiceGetResultsServer struct {
	grpc.ServerStream
}

func (x *resultServiceGetResultsServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

var _ResultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "statistico_data.ResultService",
	HandlerType: (*ResultServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetResults",
			Handler:       _ResultService_GetResults_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/result/result.proto",
}

func init() { proto.RegisterFile("proto/result/result.proto", fileDescriptor_result_5ba2b741c50e2aef) }

var fileDescriptor_result_5ba2b741c50e2aef = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x4e, 0x1b, 0x47,
	0x14, 0xc7, 0x6b, 0xc3, 0xda, 0xf8, 0x18, 0xdb, 0x74, 0x54, 0xca, 0x96, 0x2f, 0xa1, 0x55, 0x85,
	0x50, 0x5b, 0xd9, 0x60, 0xda, 0x4a, 0x89, 0xa2, 0xdc, 0x10, 0x25, 0x41, 0x09, 0x12, 0x5a, 0x23,
	0x2e, 0x92, 0x0b, 0x6b, 0xd8, 0x1d, 0xdb, 0x13, 0xed, 0x57, 0x76, 0x66, 0x21, 0xbc, 0x45, 0x9e,
	0x22, 0x0f, 0x98, 0x27, 0x88, 0xce, 0x99, 0xb5, 0xd7, 0xd8, 0x01, 0xfb, 0x8a, 0xe5, 0x9c, 0xdf,
	0xef, 0x3f, 0x5f, 0x9a, 0x31, 0xfc, 0x91, 0xa4, 0xb1, 0x8e, 0x3b, 0xa9, 0x50, 0x59, 0xa0, 0xf3,
	0x3f, 0x6d, 0xaa, 0xb1, 0x96, 0xd2, 0x5c, 0x4b, 0xa5, 0xa5, 0x17, 0xf7, 0x7d, 0xae, 0xf9, 0xf6,
	0xfe, 0x30, 0x8e, 0x87, 0x81, 0xe8, 0x50, 0xfb, 0x26, 0x1b, 0x74, 0xee, 0x52, 0x9e, 0x24, 0x22,
	0x55, 0x46, 0x70, 0x3e, 0x42, 0xc3, 0xa5, 0x00, 0x57, 0x7c, 0xce, 0x84, 0xd2, 0x6c, 0x0b, 0xaa,
	0x5a, 0xf0, 0xb0, 0x2f, 0x7d, 0xbb, 0x74, 0x50, 0x3a, 0x5a, 0x71, 0x2b, 0xf8, 0xef, 0xb9, 0xcf,
	0x4e, 0xc0, 0x0a, 0x64, 0x28, 0xb5, 0x5d, 0x3e, 0x28, 0x1d, 0xd5, 0xbb, 0x3b, 0x6d, 0x93, 0xdc,
	0x1e, 0x27, 0xb7, 0xcf, 0x23, 0x7d, 0xda, 0xbd, 0xe6, 0x41, 0x26, 0x5c, 0x43, 0x3a, 0xdf, 0x4b,
	0x50, 0x31, 0xe9, 0xac, 0x09, 0xe5, 0x49, 0x62, 0x59, 0xfa, 0xec, 0x25, 0xd4, 0xbd, 0x38, 0x4c,
	0x84, 0x96, 0x5a, 0xc6, 0x51, 0x9e, 0xb9, 0xdb, 0x9e, 0x99, 0x7e, 0xfb, 0xac, 0x60, 0xdc, 0x69,
	0x81, 0x75, 0xa0, 0xa2, 0x04, 0x57, 0x71, 0x64, 0xaf, 0x90, 0xba, 0x35, 0xa7, 0xf6, 0xa8, 0xed,
	0xe6, 0x18, 0xfb, 0x07, 0xac, 0x5b, 0x11, 0x65, 0xc2, 0x5e, 0x25, 0xfe, 0xf7, 0x39, 0xfe, 0x1a,
	0xbb, 0xae, 0x81, 0xd8, 0x33, 0x80, 0x90, 0x6b, 0x6f, 0x44, 0x2d, 0xdb, 0x22, 0x65, 0x7b, 0x4e,
	0xb9, 0x40, 0xe4, 0x15, 0xd7, 0xdc, 0xad, 0x85, 0xe3, 0x4f, 0xe7, 0x5b, 0x09, 0x6a, 0x93, 0x06,
	0xeb, 0x42, 0x6d, 0x14, 0x87, 0xa2, 0x8f, 0x9b, 0x48, 0xcb, 0xaf, 0x77, 0x37, 0xe7, 0x72, 0xae,
	0x04, 0x0f, 0xdd, 0x35, 0xe4, 0xf0, 0x0b, 0x1d, 0x7e, 0xc7, 0xef, 0x8d, 0x53, 0x7e, 0xd2, 0x41,
	0x8e, 0x9c, 0x13, 0xb0, 0x90, 0x50, 0xf9, 0x76, 0xec, 0xfc, 0x7c, 0xae, 0x3d, 0x44, 0x5c, 0x43,
	0x3a, 0x5f, 0xab, 0x00, 0x45, 0x95, 0x75, 0xc1, 0x4a, 0xa4, 0xf6, 0x46, 0xf9, 0x2c, 0x77, 0xe7,
	0xce, 0xb7, 0xa7, 0x53, 0x19, 0x0d, 0xf3, 0x03, 0x26, 0x94, 0x9d, 0x41, 0x93, 0x56, 0x37, 0x88,
	0xd3, 0x90, 0x3f, 0x38, 0xc8, 0xa7, 0xe4, 0x06, 0x3a, 0xaf, 0xc7, 0x0a, 0x86, 0xd0, 0x72, 0x8b,
	0x90, 0x95, 0x65, 0x42, 0xd0, 0x29, 0x42, 0xf6, 0x00, 0x68, 0x26, 0xca, 0x8b, 0x53, 0x73, 0xc6,
	0x96, 0x4b, 0x3b, 0xdf, 0xc3, 0x02, 0xb6, 0x69, 0x0c, 0xd3, 0xb6, 0x4c, 0x1b, 0x2b, 0xa6, 0xfd,
	0x67, 0xbe, 0x8e, 0x44, 0x44, 0x39, 0x52, 0x21, 0x64, 0x1d, 0xab, 0x97, 0x22, 0x9a, 0x50, 0x14,
	0x52, 0x50, 0x55, 0x43, 0x61, 0x75, 0x42, 0x1d, 0x42, 0x6b, 0xc4, 0x83, 0x41, 0x5f, 0xcb, 0xc9,
	0x74, 0xd6, 0x08, 0x6b, 0x60, 0xf9, 0x4a, 0x8e, 0xa7, 0x74, 0x08, 0xad, 0x41, 0x16, 0x04, 0xd3,
	0x5c, 0xcd, 0x70, 0x58, 0x2e, 0xb8, 0x23, 0xd8, 0x10, 0x5f, 0x74, 0xca, 0xa7, 0x41, 0x20, 0xb0,
	0x49, 0xf5, 0x82, 0xbc, 0x80, 0xdf, 0x68, 0x15, 0x81, 0xe0, 0xc3, 0x4c, 0xf4, 0x93, 0x58, 0x99,
	0xcb, 0x55, 0x5f, 0x7c, 0x61, 0x19, 0x8a, 0xef, 0xc9, 0xbb, 0xcc, 0x35, 0x8c, 0xa3, 0xe5, 0xce,
	0xc6, 0xad, 0x2f, 0x11, 0x87, 0xe2, 0x4c, 0xdc, 0x7f, 0x50, 0x0d, 0x65, 0x94, 0x69, 0xa1, 0xec,
	0xc6, 0xe2, 0x84, 0x31, 0x8b, 0x9a, 0x12, 0x5e, 0x1c, 0xf9, 0xca, 0x6e, 0x2e, 0xa1, 0xe5, 0x2c,
	0x7b, 0x0e, 0xc0, 0x7d, 0x5f, 0xf8, 0xb4, 0x6b, 0x76, 0x6b, 0xb1, 0x59, 0x23, 0x1c, 0x37, 0x13,
	0xdd, 0x62, 0xc7, 0xed, 0x8d, 0x25, 0xdc, 0xc9, 0x41, 0xb0, 0x17, 0x50, 0x97, 0xd1, 0xa7, 0x2c,
	0xbd, 0x37, 0xf2, 0xaf, 0x8b, 0x65, 0x30, 0x3c, 0xda, 0xce, 0x5f, 0xb0, 0x4a, 0xb7, 0x79, 0xf6,
	0xb5, 0x64, 0xb0, 0x1a, 0xf1, 0x50, 0xd0, 0xed, 0xaa, 0xb9, 0xf4, 0xed, 0xbc, 0x85, 0xfa, 0xd4,
	0xeb, 0xb8, 0x8c, 0xc2, 0x36, 0xa1, 0x22, 0x55, 0xdf, 0xcb, 0x12, 0xba, 0x61, 0x96, 0x6b, 0x49,
	0x75, 0x96, 0x25, 0xce, 0x3b, 0xa8, 0x98, 0xc7, 0x72, 0xa9, 0x90, 0x3d, 0x00, 0x0a, 0x49, 0x53,
	0x11, 0xe9, 0x3c, 0xa8, 0x86, 0x41, 0x54, 0x70, 0x06, 0x60, 0xd1, 0x4b, 0xca, 0xfe, 0x9e, 0x64,
	0x3d, 0xb2, 0x01, 0xff, 0xff, 0x6b, 0x36, 0x00, 0x07, 0x3a, 0x9e, 0x1a, 0x68, 0xd1, 0xcd, 0x27,
	0xb2, 0xfb, 0x61, 0xfc, 0xc3, 0xd5, 0x13, 0xe9, 0xad, 0xf4, 0x04, 0x3b, 0x07, 0x78, 0x23, 0xb4,
	0xa9, 0x29, 0xb6, 0x3f, 0xf7, 0x00, 0x3e, 0xf8, 0x99, 0xdb, 0xde, 0x7a, 0xa4, 0xef, 0xfc, 0x72,
	0x5c, 0xba, 0xa9, 0xd0, 0xb8, 0xa7, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xf0, 0xcf, 0x05,
	0x69, 0x07, 0x00, 0x00,
}