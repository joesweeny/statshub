// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/result/result.proto

package result // import "github.com/statistico/statistico-data/proto/result"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import competition "github.com/statistico/statistico-data/proto/competition"
import season "github.com/statistico/statistico-data/proto/season"
import team "github.com/statistico/statistico-data/proto/team"
import venue "github.com/statistico/statistico-data/proto/venue"

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

type TeamRequest struct {
	// The Team ID that the Result set relates to
	TeamId int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// The number of results to return. If limit is not set the whole Result set for the Team
	// will be returned
	Limit *wrappers.Int32Value `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// A filter to return Results before a specific date
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	DateBefore           string   `protobuf:"bytes,3,opt,name=date_before,json=dateBefore,proto3" json:"date_before,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamRequest) Reset()         { *m = TeamRequest{} }
func (m *TeamRequest) String() string { return proto.CompactTextString(m) }
func (*TeamRequest) ProtoMessage()    {}
func (*TeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_9de89dc56915fb36, []int{0}
}
func (m *TeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamRequest.Unmarshal(m, b)
}
func (m *TeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamRequest.Marshal(b, m, deterministic)
}
func (dst *TeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamRequest.Merge(dst, src)
}
func (m *TeamRequest) XXX_Size() int {
	return xxx_messageInfo_TeamRequest.Size(m)
}
func (m *TeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeamRequest proto.InternalMessageInfo

func (m *TeamRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *TeamRequest) GetLimit() *wrappers.Int32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *TeamRequest) GetDateBefore() string {
	if m != nil {
		return m.DateBefore
	}
	return ""
}

type Result struct {
	Id                   int64                    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Competition          *competition.Competition `protobuf:"bytes,2,opt,name=competition,proto3" json:"competition,omitempty"`
	Season               *season.Season           `protobuf:"bytes,3,opt,name=season,proto3" json:"season,omitempty"`
	Venue                *venue.Venue             `protobuf:"bytes,4,opt,name=venue,proto3" json:"venue,omitempty"`
	RefereeId            *wrappers.Int64Value     `protobuf:"bytes,5,opt,name=referee_id,json=refereeId,proto3" json:"referee_id,omitempty"`
	DateTime             int64                    `protobuf:"varint,6,opt,name=date_time,json=dateTime,proto3" json:"date_time,omitempty"`
	MatchData            *MatchData               `protobuf:"bytes,7,opt,name=match_data,json=matchData,proto3" json:"match_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_9de89dc56915fb36, []int{1}
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

func (m *Result) GetCompetition() *competition.Competition {
	if m != nil {
		return m.Competition
	}
	return nil
}

func (m *Result) GetSeason() *season.Season {
	if m != nil {
		return m.Season
	}
	return nil
}

func (m *Result) GetVenue() *venue.Venue {
	if m != nil {
		return m.Venue
	}
	return nil
}

func (m *Result) GetRefereeId() *wrappers.Int64Value {
	if m != nil {
		return m.RefereeId
	}
	return nil
}

func (m *Result) GetDateTime() int64 {
	if m != nil {
		return m.DateTime
	}
	return 0
}

func (m *Result) GetMatchData() *MatchData {
	if m != nil {
		return m.MatchData
	}
	return nil
}

type MatchData struct {
	HomeTeam             *team.Team  `protobuf:"bytes,1,opt,name=home_team,json=homeTeam,proto3" json:"home_team,omitempty"`
	AwayTeam             *team.Team  `protobuf:"bytes,2,opt,name=away_team,json=awayTeam,proto3" json:"away_team,omitempty"`
	Stats                *MatchStats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MatchData) Reset()         { *m = MatchData{} }
func (m *MatchData) String() string { return proto.CompactTextString(m) }
func (*MatchData) ProtoMessage()    {}
func (*MatchData) Descriptor() ([]byte, []int) {
	return fileDescriptor_result_9de89dc56915fb36, []int{2}
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

func (m *MatchData) GetHomeTeam() *team.Team {
	if m != nil {
		return m.HomeTeam
	}
	return nil
}

func (m *MatchData) GetAwayTeam() *team.Team {
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
	HomeScore            *wrappers.Int32Value  `protobuf:"bytes,4,opt,name=home_score,json=homeScore,proto3" json:"home_score,omitempty"`
	AwayScore            *wrappers.Int32Value  `protobuf:"bytes,5,opt,name=away_score,json=awayScore,proto3" json:"away_score,omitempty"`
	HomePenScore         *wrappers.Int32Value  `protobuf:"bytes,6,opt,name=home_pen_score,json=homePenScore,proto3" json:"home_pen_score,omitempty"`
	AwayPenScore         *wrappers.Int32Value  `protobuf:"bytes,7,opt,name=away_pen_score,json=awayPenScore,proto3" json:"away_pen_score,omitempty"`
	HalfTimeScore        *wrappers.StringValue `protobuf:"bytes,8,opt,name=half_time_score,json=halfTimeScore,proto3" json:"half_time_score,omitempty"`
	FullTimeScore        *wrappers.StringValue `protobuf:"bytes,9,opt,name=full_time_score,json=fullTimeScore,proto3" json:"full_time_score,omitempty"`
	ExtraTimeScore       *wrappers.StringValue `protobuf:"bytes,10,opt,name=extra_time_score,json=extraTimeScore,proto3" json:"extra_time_score,omitempty"`
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
	return fileDescriptor_result_9de89dc56915fb36, []int{3}
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

func (m *MatchStats) GetHomeScore() *wrappers.Int32Value {
	if m != nil {
		return m.HomeScore
	}
	return nil
}

func (m *MatchStats) GetAwayScore() *wrappers.Int32Value {
	if m != nil {
		return m.AwayScore
	}
	return nil
}

func (m *MatchStats) GetHomePenScore() *wrappers.Int32Value {
	if m != nil {
		return m.HomePenScore
	}
	return nil
}

func (m *MatchStats) GetAwayPenScore() *wrappers.Int32Value {
	if m != nil {
		return m.AwayPenScore
	}
	return nil
}

func (m *MatchStats) GetHalfTimeScore() *wrappers.StringValue {
	if m != nil {
		return m.HalfTimeScore
	}
	return nil
}

func (m *MatchStats) GetFullTimeScore() *wrappers.StringValue {
	if m != nil {
		return m.FullTimeScore
	}
	return nil
}

func (m *MatchStats) GetExtraTimeScore() *wrappers.StringValue {
	if m != nil {
		return m.ExtraTimeScore
	}
	return nil
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

func init() {
	proto.RegisterType((*TeamRequest)(nil), "result.TeamRequest")
	proto.RegisterType((*Result)(nil), "result.Result")
	proto.RegisterType((*MatchData)(nil), "result.MatchData")
	proto.RegisterType((*MatchStats)(nil), "result.MatchStats")
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
	GetResultsForTeam(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (ResultService_GetResultsForTeamClient, error)
}

type resultServiceClient struct {
	cc *grpc.ClientConn
}

func NewResultServiceClient(cc *grpc.ClientConn) ResultServiceClient {
	return &resultServiceClient{cc}
}

func (c *resultServiceClient) GetResultsForTeam(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (ResultService_GetResultsForTeamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ResultService_serviceDesc.Streams[0], "/result.ResultService/GetResultsForTeam", opts...)
	if err != nil {
		return nil, err
	}
	x := &resultServiceGetResultsForTeamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResultService_GetResultsForTeamClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type resultServiceGetResultsForTeamClient struct {
	grpc.ClientStream
}

func (x *resultServiceGetResultsForTeamClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResultServiceServer is the server API for ResultService service.
type ResultServiceServer interface {
	GetResultsForTeam(*TeamRequest, ResultService_GetResultsForTeamServer) error
}

func RegisterResultServiceServer(s *grpc.Server, srv ResultServiceServer) {
	s.RegisterService(&_ResultService_serviceDesc, srv)
}

func _ResultService_GetResultsForTeam_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TeamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResultServiceServer).GetResultsForTeam(m, &resultServiceGetResultsForTeamServer{stream})
}

type ResultService_GetResultsForTeamServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type resultServiceGetResultsForTeamServer struct {
	grpc.ServerStream
}

func (x *resultServiceGetResultsForTeamServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

var _ResultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "result.ResultService",
	HandlerType: (*ResultServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetResultsForTeam",
			Handler:       _ResultService_GetResultsForTeam_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/result/result.proto",
}

func init() { proto.RegisterFile("proto/result/result.proto", fileDescriptor_result_9de89dc56915fb36) }

var fileDescriptor_result_9de89dc56915fb36 = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xf3, 0x44,
	0x10, 0x25, 0x69, 0xf3, 0x37, 0x6e, 0xd3, 0x76, 0x01, 0xd5, 0xb4, 0x08, 0xaa, 0x20, 0x41, 0x6e,
	0x70, 0x4a, 0x5a, 0xb8, 0x88, 0xb8, 0xa1, 0xad, 0x8a, 0x2a, 0x51, 0xa9, 0x72, 0xaa, 0x5e, 0x70,
	0x13, 0x6d, 0xe2, 0x49, 0xb2, 0xc8, 0x3f, 0xc1, 0x5e, 0xb7, 0x94, 0x17, 0xe0, 0x61, 0x78, 0x27,
	0x9e, 0x05, 0xcd, 0xec, 0x26, 0x36, 0xad, 0xf8, 0xe2, 0x1b, 0x7b, 0x3d, 0x73, 0xce, 0xd9, 0x9d,
	0xb3, 0xa3, 0x31, 0x7c, 0xb6, 0x4a, 0x13, 0x9d, 0x0c, 0x52, 0xcc, 0xf2, 0x50, 0xdb, 0x97, 0xc7,
	0x31, 0xd1, 0x34, 0x5f, 0x27, 0x5f, 0x2c, 0x92, 0x64, 0x11, 0xe2, 0x80, 0xa3, 0xd3, 0x7c, 0x3e,
	0x78, 0x49, 0xe5, 0x6a, 0x85, 0x69, 0x66, 0x70, 0x27, 0x5f, 0x19, 0x89, 0x59, 0x12, 0xad, 0x50,
	0x2b, 0xad, 0x92, 0xb8, 0xbc, 0xb6, 0x20, 0xbb, 0x4f, 0x86, 0x32, 0x4b, 0x62, 0xfb, 0xb2, 0xa9,
	0x4f, 0x4d, 0x4a, 0xa3, 0x8c, 0xf8, 0x61, 0xc3, 0xc7, 0x26, 0xfc, 0x8c, 0x71, 0x8e, 0xe6, 0x69,
	0x12, 0xbd, 0x3f, 0xc1, 0x79, 0x44, 0x19, 0xf9, 0xf8, 0x7b, 0x8e, 0x99, 0x16, 0xc7, 0xd0, 0x22,
	0xd6, 0x44, 0x05, 0x6e, 0xed, 0xac, 0xd6, 0xdf, 0xf1, 0x9b, 0xf4, 0x79, 0x17, 0x88, 0xef, 0xa0,
	0x11, 0xaa, 0x48, 0x69, 0xb7, 0x7e, 0x56, 0xeb, 0x3b, 0xc3, 0x53, 0xcf, 0xd4, 0xe1, 0xad, 0xeb,
	0xf0, 0xee, 0x62, 0x7d, 0x31, 0x7c, 0x92, 0x61, 0x8e, 0xbe, 0x41, 0x8a, 0x2f, 0xc1, 0x09, 0xa4,
	0xc6, 0xc9, 0x14, 0xe7, 0x49, 0x8a, 0xee, 0xce, 0x59, 0xad, 0xdf, 0xf1, 0x81, 0x42, 0x57, 0x1c,
	0xe9, 0xfd, 0x5d, 0x87, 0xa6, 0xcf, 0xb6, 0x88, 0x2e, 0xd4, 0x37, 0x5b, 0xd6, 0x55, 0x20, 0x46,
	0xe0, 0x94, 0xca, 0xb6, 0x9b, 0xba, 0x5e, 0xd9, 0x8a, 0xeb, 0x62, 0xed, 0x97, 0xc1, 0xe2, 0x6b,
	0x68, 0x1a, 0x4b, 0x78, 0x4b, 0x67, 0xd8, 0xf5, 0xac, 0x43, 0x63, 0x7e, 0xf9, 0x36, 0x2b, 0x7a,
	0xd0, 0x60, 0x27, 0xdc, 0x5d, 0x86, 0xed, 0x79, 0xc6, 0x97, 0x27, 0x7a, 0xfa, 0x26, 0x25, 0x46,
	0x00, 0x29, 0xce, 0x31, 0x45, 0x24, 0x4b, 0x1a, 0xff, 0x5f, 0xfb, 0x0f, 0x97, 0xa6, 0xf6, 0x8e,
	0x85, 0xdf, 0x05, 0xe2, 0x14, 0x3a, 0x5c, 0xbf, 0x56, 0x11, 0xba, 0x4d, 0x2e, 0xad, 0x4d, 0x81,
	0x47, 0x15, 0xa1, 0x38, 0x07, 0x88, 0xa4, 0x9e, 0x2d, 0x27, 0x81, 0xd4, 0xd2, 0x6d, 0xb1, 0xf0,
	0x91, 0x67, 0x5b, 0xe6, 0x9e, 0x32, 0x37, 0x52, 0x4b, 0xbf, 0x13, 0xad, 0x97, 0xbd, 0xbf, 0x6a,
	0xd0, 0xd9, 0x24, 0xc4, 0x37, 0xd0, 0x59, 0x26, 0x11, 0x4e, 0xe8, 0x7a, 0xd8, 0x37, 0x67, 0x08,
	0x1e, 0x5f, 0x38, 0x5f, 0x67, 0x9b, 0x92, 0xb4, 0x22, 0xa0, 0x7c, 0x91, 0xaf, 0x06, 0x58, 0x7f,
	0x0f, 0xa4, 0x24, 0x03, 0xfb, 0xd0, 0xc8, 0xb4, 0xd4, 0x99, 0x75, 0x4d, 0xfc, 0xe7, 0x30, 0x63,
	0xca, 0xf8, 0x06, 0xd0, 0xfb, 0xa7, 0x0d, 0x50, 0x44, 0xc5, 0x10, 0x1a, 0x2b, 0xa5, 0x67, 0x4b,
	0x7b, 0x8c, 0xcf, 0xdf, 0xd9, 0x33, 0xd6, 0xa9, 0x8a, 0x17, 0xb6, 0x37, 0x18, 0x2a, 0xae, 0xa1,
	0xcb, 0xc7, 0x9f, 0x27, 0x69, 0x24, 0x4b, 0x57, 0xfc, 0x61, 0xf2, 0x3e, 0x71, 0x6e, 0xd7, 0x14,
	0x12, 0xe1, 0xd2, 0x0a, 0x91, 0x9d, 0x2a, 0x22, 0xc4, 0x29, 0x44, 0x46, 0x00, 0x7c, 0x92, 0x6c,
	0x46, 0x4d, 0xba, 0xbb, 0xbd, 0xbb, 0xd9, 0xf7, 0x31, 0xa1, 0x89, 0xcb, 0x07, 0x30, 0xdc, 0x46,
	0x05, 0x2e, 0xc1, 0x0d, 0xf7, 0x27, 0xeb, 0xc0, 0x0a, 0x63, 0xcb, 0x6f, 0x6e, 0xe7, 0xef, 0x11,
	0xe5, 0x01, 0xe3, 0x8d, 0x04, 0x6f, 0x5f, 0x48, 0xb4, 0x2a, 0x48, 0x10, 0x65, 0x23, 0x71, 0x03,
	0x07, 0x4b, 0x19, 0xce, 0xb9, 0x47, 0xad, 0x46, 0xbb, 0xd2, 0x45, 0xc8, 0x70, 0x4e, 0x7d, 0xbc,
	0x51, 0x99, 0xe7, 0x61, 0x58, 0x56, 0xe9, 0x54, 0x51, 0x21, 0x52, 0xa1, 0x72, 0x0b, 0x87, 0xf8,
	0x87, 0x4e, 0x65, 0x59, 0x06, 0x2a, 0xc8, 0x74, 0x99, 0x55, 0xe8, 0xdc, 0xc3, 0x27, 0xec, 0x6c,
	0x88, 0x72, 0x91, 0xe3, 0x64, 0x95, 0x64, 0x66, 0x88, 0x38, 0xdb, 0xcd, 0x11, 0x44, 0xfc, 0x85,
	0x79, 0x0f, 0x96, 0x46, 0x72, 0xec, 0xf2, 0x5b, 0xb9, 0xbd, 0x0a, 0x72, 0x44, 0x7c, 0x23, 0xf7,
	0x3d, 0xb4, 0x22, 0x15, 0xe7, 0x1a, 0x33, 0x77, 0x7f, 0xbb, 0xc2, 0x1a, 0x4b, 0xb4, 0x0c, 0x67,
	0x49, 0x1c, 0x64, 0x6e, 0xb7, 0x02, 0xcd, 0x62, 0xb9, 0x43, 0x83, 0x00, 0x03, 0x33, 0x84, 0x0e,
	0xaa, 0x74, 0x28, 0xc1, 0x79, 0x44, 0x8d, 0x00, 0x8a, 0xfb, 0x70, 0x0f, 0x2b, 0x70, 0x37, 0x17,
	0x21, 0x7e, 0x04, 0x47, 0xc5, 0xbf, 0xe5, 0xe9, 0xab, 0x21, 0x1f, 0x6d, 0x27, 0x83, 0xc1, 0x13,
	0x7b, 0x78, 0x0f, 0xfb, 0xe6, 0xbf, 0x30, 0xc6, 0xf4, 0x59, 0xcd, 0x48, 0xee, 0xe8, 0x67, 0xd4,
	0x26, 0x96, 0xdd, 0x26, 0x29, 0x0f, 0xac, 0x8f, 0xd7, 0x13, 0xaa, 0xf4, 0x03, 0x3b, 0xe9, 0xae,
	0x83, 0x06, 0xdc, 0xfb, 0xe8, 0xbc, 0x76, 0x75, 0xf9, 0xeb, 0x70, 0xa1, 0xf4, 0x32, 0x9f, 0xd2,
	0xff, 0x63, 0x40, 0x33, 0x4c, 0x65, 0x5a, 0xcd, 0x92, 0xd2, 0xf2, 0x5b, 0x1a, 0xc3, 0x83, 0xf2,
	0xef, 0x7b, 0xda, 0xe4, 0xaf, 0x8b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x89, 0x47, 0x98,
	0xd5, 0x07, 0x00, 0x00,
}