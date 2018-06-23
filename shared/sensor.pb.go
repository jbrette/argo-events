// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared/sensor.proto

package shared // import "github.com/argoproj/argo-events/shared"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1alpha1 "github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
import empty "github.com/golang/protobuf/ptypes/empty"

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

// TriggerRequest contains the trigger and the corresponding dependent event(s)
type TriggerRequest struct {
	Trigger              *v1alpha1.Trigger `protobuf:"bytes,1,opt,name=trigger" json:"trigger,omitempty"`
	Events               []*v1alpha1.Event `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TriggerRequest) Reset()         { *m = TriggerRequest{} }
func (m *TriggerRequest) String() string { return proto.CompactTextString(m) }
func (*TriggerRequest) ProtoMessage()    {}
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sensor_7a9e27428984a6b3, []int{0}
}
func (m *TriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerRequest.Unmarshal(m, b)
}
func (m *TriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerRequest.Marshal(b, m, deterministic)
}
func (dst *TriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerRequest.Merge(dst, src)
}
func (m *TriggerRequest) XXX_Size() int {
	return xxx_messageInfo_TriggerRequest.Size(m)
}
func (m *TriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerRequest proto.InternalMessageInfo

func (m *TriggerRequest) GetTrigger() *v1alpha1.Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *TriggerRequest) GetEvents() []*v1alpha1.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*TriggerRequest)(nil), "shared.TriggerRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Signal service

type SignalClient interface {
	// Start listening. Events are streamed back to the client.
	Start(ctx context.Context, in *v1alpha1.Signal, opts ...grpc.CallOption) (Signal_StartClient, error)
	// Stop listening. This should terminate the above event stream.
	Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type signalClient struct {
	cc *grpc.ClientConn
}

func NewSignalClient(cc *grpc.ClientConn) SignalClient {
	return &signalClient{cc}
}

func (c *signalClient) Start(ctx context.Context, in *v1alpha1.Signal, opts ...grpc.CallOption) (Signal_StartClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Signal_serviceDesc.Streams[0], c.cc, "/shared.Signal/Start", opts...)
	if err != nil {
		return nil, err
	}
	x := &signalStartClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Signal_StartClient interface {
	Recv() (*v1alpha1.Event, error)
	grpc.ClientStream
}

type signalStartClient struct {
	grpc.ClientStream
}

func (x *signalStartClient) Recv() (*v1alpha1.Event, error) {
	m := new(v1alpha1.Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *signalClient) Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/shared.Signal/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Signal service

type SignalServer interface {
	// Start listening. Events are streamed back to the client.
	Start(*v1alpha1.Signal, Signal_StartServer) error
	// Stop listening. This should terminate the above event stream.
	Stop(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterSignalServer(s *grpc.Server, srv SignalServer) {
	s.RegisterService(&_Signal_serviceDesc, srv)
}

func _Signal_Start_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(v1alpha1.Signal)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SignalServer).Start(m, &signalStartServer{stream})
}

type Signal_StartServer interface {
	Send(*v1alpha1.Event) error
	grpc.ServerStream
}

type signalStartServer struct {
	grpc.ServerStream
}

func (x *signalStartServer) Send(m *v1alpha1.Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Signal_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.Signal/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServer).Stop(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Signal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shared.Signal",
	HandlerType: (*SignalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stop",
			Handler:    _Signal_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Start",
			Handler:       _Signal_Start_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "shared/sensor.proto",
}

// Client API for Trigger service

type TriggerClient interface {
	Fire(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type triggerClient struct {
	cc *grpc.ClientConn
}

func NewTriggerClient(cc *grpc.ClientConn) TriggerClient {
	return &triggerClient{cc}
}

func (c *triggerClient) Fire(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/shared.Trigger/Fire", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Trigger service

type TriggerServer interface {
	Fire(context.Context, *TriggerRequest) (*empty.Empty, error)
}

func RegisterTriggerServer(s *grpc.Server, srv TriggerServer) {
	s.RegisterService(&_Trigger_serviceDesc, srv)
}

func _Trigger_Fire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServer).Fire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.Trigger/Fire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServer).Fire(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Trigger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shared.Trigger",
	HandlerType: (*TriggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fire",
			Handler:    _Trigger_Fire_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shared/sensor.proto",
}

func init() { proto.RegisterFile("shared/sensor.proto", fileDescriptor_sensor_7a9e27428984a6b3) }

var fileDescriptor_sensor_7a9e27428984a6b3 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x59, 0xff, 0xa4, 0xb0, 0x05, 0x0f, 0x2b, 0x94, 0x12, 0x2f, 0xc5, 0x83, 0xf4, 0xe2,
	0xac, 0x8d, 0xe0, 0x55, 0x2b, 0x54, 0x3c, 0x37, 0x82, 0xa0, 0x07, 0xd9, 0xd8, 0x71, 0x13, 0x9b,
	0x66, 0xd7, 0xdd, 0x4d, 0xc1, 0x77, 0xf0, 0xb9, 0xbc, 0xfa, 0x4a, 0x92, 0x6c, 0x16, 0xf1, 0x50,
	0x84, 0x7a, 0x9b, 0xcc, 0x30, 0x1f, 0x5f, 0xf6, 0x37, 0xf4, 0xd0, 0xe6, 0xc2, 0xe0, 0x82, 0x5b,
	0xac, 0xac, 0x32, 0xa0, 0x8d, 0x72, 0x8a, 0x45, 0xbe, 0x19, 0xdf, 0xca, 0xc2, 0xe5, 0x75, 0x06,
	0xcf, 0x6a, 0xc5, 0x85, 0x91, 0x4a, 0x1b, 0xf5, 0xda, 0x16, 0xa7, 0xb8, 0xc6, 0xca, 0x59, 0xae,
	0x97, 0x92, 0x0b, 0x5d, 0xd8, 0x6e, 0x9d, 0xaf, 0x27, 0xa2, 0xd4, 0xb9, 0x98, 0x70, 0x89, 0x15,
	0x1a, 0xe1, 0x70, 0xe1, 0x89, 0xf1, 0x91, 0x54, 0x4a, 0x96, 0xc8, 0xdb, 0xaf, 0xac, 0x7e, 0xe1,
	0xb8, 0xd2, 0xee, 0xdd, 0x0f, 0x8f, 0x3f, 0x09, 0x3d, 0xb8, 0x33, 0x85, 0x94, 0x68, 0xe6, 0xf8,
	0x56, 0xa3, 0x75, 0xec, 0x91, 0xf6, 0x9c, 0xef, 0x0c, 0xc9, 0x88, 0x8c, 0xfb, 0xc9, 0x14, 0x7e,
	0x5c, 0x20, 0xb8, 0xb4, 0xc5, 0x93, 0x77, 0x01, 0xbd, 0x94, 0xd0, 0xb8, 0x40, 0xf7, 0x2b, 0xc1,
	0x05, 0x02, 0x3a, 0x10, 0xd9, 0x3d, 0x8d, 0xfc, 0xce, 0x70, 0x67, 0xb4, 0x3b, 0xee, 0x27, 0x97,
	0xdb, 0xb3, 0x67, 0xcd, 0x7c, 0xde, 0xe1, 0x92, 0x2f, 0x42, 0xa3, 0xb4, 0x90, 0x95, 0x28, 0xd9,
	0x07, 0xa1, 0xfb, 0xa9, 0x13, 0xc6, 0xb1, 0xab, 0xed, 0xe9, 0x9e, 0x15, 0xff, 0xd7, 0xef, 0x8c,
	0xb0, 0x0b, 0xba, 0x97, 0x3a, 0xa5, 0xd9, 0x00, 0x7c, 0x10, 0x10, 0x82, 0x80, 0x59, 0x13, 0x44,
	0xbc, 0xa1, 0x9f, 0x4c, 0x69, 0xaf, 0x7b, 0xbe, 0x06, 0x71, 0x53, 0x18, 0x64, 0x03, 0xf0, 0xd7,
	0x01, 0xbf, 0x23, 0xdb, 0x84, 0xb8, 0x1e, 0x3f, 0x9c, 0xfc, 0x75, 0x46, 0x1e, 0x98, 0x45, 0xed,
	0xe6, 0xf9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xf4, 0x77, 0xce, 0x94, 0x02, 0x00, 0x00,
}