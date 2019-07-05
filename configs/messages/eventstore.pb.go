// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eventstore.proto

package messages

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Event struct {
	EventId              string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventType            string   `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	AggregateId          string   `protobuf:"bytes,3,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateType        string   `protobuf:"bytes,4,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	EventData            string   `protobuf:"bytes,5,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	Channel              string   `protobuf:"bytes,6,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_00783576c07562d6, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *Event) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

func (m *Event) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type Response struct {
	IsSuccess            bool     `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_00783576c07562d6, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type EventFilter struct {
	EventId              string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	AggregateId          string   `protobuf:"bytes,2,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_00783576c07562d6, []int{2}
}

func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFilter.Unmarshal(m, b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
}
func (m *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(m, src)
}
func (m *EventFilter) XXX_Size() int {
	return xxx_messageInfo_EventFilter.Size(m)
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EventFilter) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

type EventResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00783576c07562d6, []int{3}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "messages.Event")
	proto.RegisterType((*Response)(nil), "messages.Response")
	proto.RegisterType((*EventFilter)(nil), "messages.EventFilter")
	proto.RegisterType((*EventResponse)(nil), "messages.EventResponse")
}

func init() { proto.RegisterFile("eventstore.proto", fileDescriptor_00783576c07562d6) }

var fileDescriptor_00783576c07562d6 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0xad, 0x6d, 0x93, 0x89, 0x55, 0x59, 0x14, 0xd7, 0x82, 0x50, 0x03, 0x62, 0x4f, 0x39,
	0x54, 0x0f, 0x5e, 0xc4, 0x83, 0x5f, 0x14, 0x6f, 0xa9, 0xf7, 0xb2, 0x26, 0x43, 0x0c, 0xd4, 0x24,
	0xec, 0xac, 0x42, 0x8f, 0xfe, 0x37, 0x7f, 0x98, 0x64, 0x36, 0x26, 0x12, 0xc1, 0xe3, 0xbc, 0x79,
	0xf3, 0x66, 0xde, 0xbe, 0x85, 0x7d, 0xfc, 0xc0, 0xdc, 0x90, 0x29, 0x34, 0x86, 0xa5, 0x2e, 0x4c,
	0x21, 0xdc, 0x37, 0x24, 0x52, 0x29, 0x52, 0xf0, 0xe5, 0xc0, 0xe0, 0xbe, 0x6a, 0x8b, 0x63, 0x70,
	0x99, 0xb7, 0xca, 0x12, 0xe9, 0x4c, 0x9d, 0x99, 0x17, 0x8d, 0xb8, 0x5e, 0x24, 0xe2, 0x04, 0xc0,
	0xb6, 0xcc, 0xa6, 0x44, 0xd9, 0xe3, 0xa6, 0xc7, 0xc8, 0xf3, 0xa6, 0x44, 0x71, 0x0a, 0x3b, 0x2a,
	0x4d, 0x35, 0xa6, 0xca, 0x60, 0x35, 0xdd, 0x67, 0x82, 0xdf, 0x60, 0x8b, 0x44, 0x9c, 0xc1, 0x6e,
	0x4b, 0x61, 0x95, 0x6d, 0x26, 0x8d, 0x1b, 0x94, 0x95, 0x9a, 0x45, 0x89, 0x32, 0x4a, 0x0e, 0x7e,
	0x2d, 0xba, 0x53, 0x46, 0x09, 0x09, 0xa3, 0xf8, 0x55, 0xe5, 0x39, 0xae, 0xe5, 0xd0, 0x5e, 0x58,
	0x97, 0xc1, 0x0d, 0xb8, 0x11, 0x52, 0x59, 0xe4, 0xc4, 0x22, 0x19, 0xad, 0xe8, 0x3d, 0x8e, 0x91,
	0x88, 0xad, 0xb8, 0x91, 0x97, 0xd1, 0xd2, 0x02, 0xe2, 0x00, 0x06, 0xa8, 0x75, 0xa1, 0x6b, 0x1f,
	0xb6, 0x08, 0x9e, 0xc0, 0xe7, 0x67, 0x78, 0xc8, 0xd6, 0x06, 0xf5, 0x7f, 0x8f, 0xd1, 0x75, 0xdb,
	0xfb, 0xe3, 0x36, 0xb8, 0x82, 0x31, 0x8b, 0x35, 0x27, 0x9d, 0xc3, 0xd0, 0x66, 0x20, 0x9d, 0x69,
	0x7f, 0xe6, 0xcf, 0xf7, 0xc2, 0x9f, 0x00, 0x42, 0x4b, 0xac, 0xdb, 0xf3, 0x4f, 0x07, 0x80, 0x91,
	0x65, 0x95, 0x96, 0xb8, 0x06, 0xef, 0x11, 0x0d, 0x03, 0x24, 0x0e, 0x3b, 0x43, 0xf6, 0xd4, 0xc9,
	0x51, 0x57, 0xab, 0x5e, 0x1a, 0x6c, 0x89, 0x4b, 0xf0, 0x6f, 0x35, 0x2a, 0x83, 0x36, 0xe1, 0xee,
	0xd6, 0x89, 0x68, 0x81, 0x76, 0xea, 0x65, 0xc8, 0x7f, 0xe4, 0xe2, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xe9, 0x3b, 0x01, 0x04, 0x37, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventStoreClient is the client API for EventStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventStoreClient interface {
	// Get all events for the given aggregate and event
	GetEvents(ctx context.Context, in *EventFilter, opts ...grpc.CallOption) (*EventResponse, error)
	// Create a new event to the event store
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error)
}

type eventStoreClient struct {
	cc *grpc.ClientConn
}

func NewEventStoreClient(cc *grpc.ClientConn) EventStoreClient {
	return &eventStoreClient{cc}
}

func (c *eventStoreClient) GetEvents(ctx context.Context, in *EventFilter, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/messages.EventStore/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStoreClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/messages.EventStore/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventStoreServer is the server API for EventStore service.
type EventStoreServer interface {
	// Get all events for the given aggregate and event
	GetEvents(context.Context, *EventFilter) (*EventResponse, error)
	// Create a new event to the event store
	CreateEvent(context.Context, *Event) (*Response, error)
}

// UnimplementedEventStoreServer can be embedded to have forward compatible implementations.
type UnimplementedEventStoreServer struct {
}

func (*UnimplementedEventStoreServer) GetEvents(ctx context.Context, req *EventFilter) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (*UnimplementedEventStoreServer) CreateEvent(ctx context.Context, req *Event) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}

func RegisterEventStoreServer(s *grpc.Server, srv EventStoreServer) {
	s.RegisterService(&_EventStore_serviceDesc, srv)
}

func _EventStore_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.EventStore/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).GetEvents(ctx, req.(*EventFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStore_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.EventStore/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.EventStore",
	HandlerType: (*EventStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _EventStore_GetEvents_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _EventStore_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventstore.proto",
}
