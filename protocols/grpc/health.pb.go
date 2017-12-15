// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: health.proto

package grpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc1 "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorHealth, []int{1, 0}
}

type CheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{0} }

func (m *CheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=grpc.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{1} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

type VersionCheckResponse struct {
	GrpcVersion  string `protobuf:"bytes,1,opt,name=grpc_version,json=grpcVersion,proto3" json:"grpc_version,omitempty"`
	AgentVersion string `protobuf:"bytes,2,opt,name=agent_version,json=agentVersion,proto3" json:"agent_version,omitempty"`
}

func (m *VersionCheckResponse) Reset()                    { *m = VersionCheckResponse{} }
func (m *VersionCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionCheckResponse) ProtoMessage()               {}
func (*VersionCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{2} }

func (m *VersionCheckResponse) GetGrpcVersion() string {
	if m != nil {
		return m.GrpcVersion
	}
	return ""
}

func (m *VersionCheckResponse) GetAgentVersion() string {
	if m != nil {
		return m.AgentVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "grpc.CheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "grpc.HealthCheckResponse")
	proto.RegisterType((*VersionCheckResponse)(nil), "grpc.VersionCheckResponse")
	proto.RegisterEnum("grpc.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}
func (this *CheckRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CheckRequest)
	if !ok {
		that2, ok := that.(CheckRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Service != that1.Service {
		return false
	}
	return true
}
func (this *HealthCheckResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheckResponse)
	if !ok {
		that2, ok := that.(HealthCheckResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (this *VersionCheckResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VersionCheckResponse)
	if !ok {
		that2, ok := that.(VersionCheckResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.GrpcVersion != that1.GrpcVersion {
		return false
	}
	if this.AgentVersion != that1.AgentVersion {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc1.CallOption) (*HealthCheckResponse, error)
	Version(ctx context.Context, in *CheckRequest, opts ...grpc1.CallOption) (*VersionCheckResponse, error)
}

type healthClient struct {
	cc *grpc1.ClientConn
}

func NewHealthClient(cc *grpc1.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc1.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc1.Invoke(ctx, "/grpc.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) Version(ctx context.Context, in *CheckRequest, opts ...grpc1.CallOption) (*VersionCheckResponse, error) {
	out := new(VersionCheckResponse)
	err := grpc1.Invoke(ctx, "/grpc.Health/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *CheckRequest) (*HealthCheckResponse, error)
	Version(context.Context, *CheckRequest) (*VersionCheckResponse, error)
}

func RegisterHealthServer(s *grpc1.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Version(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Health/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Version(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Health_Version_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "health.proto",
}

func NewPopulatedCheckRequest(r randyHealth, easy bool) *CheckRequest {
	this := &CheckRequest{}
	this.Service = string(randStringHealth(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedHealthCheckResponse(r randyHealth, easy bool) *HealthCheckResponse {
	this := &HealthCheckResponse{}
	this.Status = HealthCheckResponse_ServingStatus([]int32{0, 1, 2}[r.Intn(3)])
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedVersionCheckResponse(r randyHealth, easy bool) *VersionCheckResponse {
	this := &VersionCheckResponse{}
	this.GrpcVersion = string(randStringHealth(r))
	this.AgentVersion = string(randStringHealth(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyHealth interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHealth(r randyHealth) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHealth(r randyHealth) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneHealth(r)
	}
	return string(tmps)
}
func randUnrecognizedHealth(r randyHealth, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHealth(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHealth(dAtA []byte, r randyHealth, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHealth(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHealth(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("health.proto", fileDescriptorHealth) }

var fileDescriptorHealth = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2f, 0x2a, 0x48, 0x96, 0xd2,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7,
	0x07, 0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4, 0xa4, 0xc1, 0xc5,
	0xe3, 0x9c, 0x91, 0x9a, 0x9c, 0x1d, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc1, 0xc5,
	0x5e, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3,
	0x2a, 0x4d, 0x62, 0xe4, 0x12, 0xf6, 0x00, 0xdb, 0x07, 0xd5, 0x50, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x2a, 0x64, 0xcf, 0xc5, 0x56, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x0c, 0xd6, 0xc0, 0x67, 0xa4, 0xae,
	0x07, 0x72, 0x87, 0x1e, 0x16, 0xa5, 0x7a, 0xc1, 0x20, 0xa3, 0xf2, 0xd2, 0x83, 0xc1, 0xca, 0x83,
	0xa0, 0xda, 0x94, 0xac, 0xb8, 0x78, 0x51, 0x24, 0x84, 0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd,
	0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x40, 0x9c, 0x60, 0xd7, 0xa0, 0x30, 0x4f, 0x3f, 0x77, 0x01, 0x46,
	0x21, 0x7e, 0x2e, 0x6e, 0x3f, 0xff, 0x90, 0x78, 0x98, 0x00, 0x93, 0x52, 0x1c, 0x97, 0x48, 0x58,
	0x6a, 0x51, 0x71, 0x66, 0x7e, 0x1e, 0xaa, 0xa3, 0x14, 0xb9, 0x78, 0x40, 0xae, 0x88, 0x2f, 0x83,
	0x48, 0x42, 0xfd, 0xc2, 0x0d, 0x12, 0x83, 0xaa, 0x17, 0x52, 0xe6, 0xe2, 0x4d, 0x4c, 0x4f, 0xcd,
	0x2b, 0x81, 0xab, 0x61, 0x02, 0xab, 0xe1, 0x01, 0x0b, 0x42, 0x15, 0x19, 0x55, 0x73, 0xb1, 0x41,
	0x3c, 0x22, 0x64, 0xc6, 0xc5, 0x0a, 0xb6, 0x42, 0x48, 0x08, 0xe2, 0x3f, 0xe4, 0x50, 0x93, 0x92,
	0xc4, 0xe9, 0x67, 0x21, 0x4b, 0x2e, 0x76, 0x98, 0x8d, 0xd8, 0x74, 0x4a, 0x41, 0xc4, 0xb0, 0x79,
	0xc2, 0x49, 0xe0, 0xc7, 0x43, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x77, 0x3c, 0x92, 0x63, 0x3c,
	0xf0, 0x48, 0x8e, 0x31, 0x89, 0x0d, 0x1c, 0x69, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e,
	0x1c, 0xfd, 0x3a, 0xf9, 0x01, 0x00, 0x00,
}
