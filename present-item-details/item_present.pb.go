// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item_present.proto

package main

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Item struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=Type,proto3" json:"Type,omitempty"`
	Rarity               string   `protobuf:"bytes,6,opt,name=Rarity,proto3" json:"Rarity,omitempty"`
	Level                uint32   `protobuf:"varint,7,opt,name=Level,proto3" json:"Level,omitempty"`
	VendorValue          uint32   `protobuf:"varint,8,opt,name=VendorValue,proto3" json:"VendorValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_12063e307637a3d8, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Item) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Item) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Item) GetRarity() string {
	if m != nil {
		return m.Rarity
	}
	return ""
}

func (m *Item) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Item) GetVendorValue() uint32 {
	if m != nil {
		return m.VendorValue
	}
	return 0
}

type ItemIdFilter struct {
	ItemId               uint32   `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemIdFilter) Reset()         { *m = ItemIdFilter{} }
func (m *ItemIdFilter) String() string { return proto.CompactTextString(m) }
func (*ItemIdFilter) ProtoMessage()    {}
func (*ItemIdFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_12063e307637a3d8, []int{1}
}

func (m *ItemIdFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemIdFilter.Unmarshal(m, b)
}
func (m *ItemIdFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemIdFilter.Marshal(b, m, deterministic)
}
func (m *ItemIdFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemIdFilter.Merge(m, src)
}
func (m *ItemIdFilter) XXX_Size() int {
	return xxx_messageInfo_ItemIdFilter.Size(m)
}
func (m *ItemIdFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemIdFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ItemIdFilter proto.InternalMessageInfo

func (m *ItemIdFilter) GetItemId() uint32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

type ItemIds struct {
	ItemId               uint32   `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemIds) Reset()         { *m = ItemIds{} }
func (m *ItemIds) String() string { return proto.CompactTextString(m) }
func (*ItemIds) ProtoMessage()    {}
func (*ItemIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_12063e307637a3d8, []int{2}
}

func (m *ItemIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemIds.Unmarshal(m, b)
}
func (m *ItemIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemIds.Marshal(b, m, deterministic)
}
func (m *ItemIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemIds.Merge(m, src)
}
func (m *ItemIds) XXX_Size() int {
	return xxx_messageInfo_ItemIds.Size(m)
}
func (m *ItemIds) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemIds.DiscardUnknown(m)
}

var xxx_messageInfo_ItemIds proto.InternalMessageInfo

func (m *ItemIds) GetItemId() uint32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "main.Item")
	proto.RegisterType((*ItemIdFilter)(nil), "main.ItemIdFilter")
	proto.RegisterType((*ItemIds)(nil), "main.ItemIds")
}

func init() { proto.RegisterFile("item_present.proto", fileDescriptor_12063e307637a3d8) }

var fileDescriptor_12063e307637a3d8 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xff, 0xe9, 0x3f, 0x4d, 0xeb, 0xd4, 0x7a, 0x18, 0xa4, 0x2e, 0xf5, 0x52, 0x72, 0xb1,
	0xa7, 0xad, 0xd4, 0x83, 0x5f, 0xa0, 0x2a, 0x0b, 0x22, 0x52, 0xa4, 0x57, 0x49, 0x9b, 0x31, 0x2c,
	0x24, 0xd9, 0xb0, 0xd9, 0x0a, 0xf1, 0xf3, 0xf9, 0xc1, 0x64, 0x76, 0x2b, 0xf6, 0xd2, 0xdb, 0xec,
	0xef, 0xbd, 0xc7, 0xec, 0x3c, 0x40, 0xed, 0xa8, 0x7a, 0x6f, 0x2c, 0xb5, 0x54, 0x3b, 0xd9, 0x58,
	0xe3, 0x0c, 0xc6, 0x55, 0xa6, 0xeb, 0xe9, 0x75, 0x61, 0x4c, 0x51, 0xd2, 0xc2, 0xb3, 0xed, 0xfe,
	0x63, 0x41, 0x55, 0xe3, 0xba, 0x60, 0x49, 0xbf, 0x23, 0x88, 0x95, 0xa3, 0x0a, 0x2f, 0xa0, 0xa7,
	0x56, 0x22, 0x9a, 0x45, 0xf3, 0xf1, 0xba, 0xa7, 0x56, 0x88, 0x10, 0xbf, 0x64, 0x15, 0x89, 0xde,
	0x2c, 0x9a, 0x9f, 0xad, 0xfd, 0xcc, 0x4c, 0xed, 0x4c, 0x2d, 0xfe, 0x07, 0xc6, 0x33, 0xce, 0x60,
	0xb4, 0xa2, 0x76, 0x67, 0x75, 0xe3, 0xb4, 0xa9, 0x45, 0xec, 0xa5, 0x63, 0xc4, 0xa9, 0xb7, 0xae,
	0x21, 0xd1, 0x0f, 0x29, 0x9e, 0x71, 0x02, 0xc9, 0x3a, 0xb3, 0xda, 0x75, 0x22, 0xf1, 0xf4, 0xf0,
	0xc2, 0x4b, 0xe8, 0x3f, 0xd3, 0x27, 0x95, 0x62, 0xe0, 0x3f, 0x12, 0x1e, 0xbc, 0x63, 0x43, 0x75,
	0x6e, 0xec, 0x26, 0x2b, 0xf7, 0x24, 0x86, 0x5e, 0x3b, 0x46, 0xe9, 0x0d, 0x9c, 0xf3, 0x15, 0x2a,
	0x7f, 0xd4, 0xa5, 0x23, 0x8b, 0x57, 0x30, 0xf0, 0x7d, 0xe8, 0xfc, 0x70, 0x52, 0xa2, 0xbd, 0x9c,
	0xa6, 0x30, 0x08, 0xc6, 0xf6, 0xa4, 0x67, 0xf9, 0x05, 0x23, 0xf6, 0xbc, 0x86, 0x2e, 0x71, 0x09,
	0xc3, 0x82, 0x1c, 0x93, 0x16, 0x27, 0x32, 0x94, 0x29, 0x7f, 0xcb, 0x94, 0x0f, 0x5c, 0xe6, 0x14,
	0x24, 0x57, 0x2d, 0xd9, 0x94, 0xfe, 0xbb, 0x8d, 0xf0, 0x1e, 0xe0, 0x29, 0x64, 0x78, 0xd3, 0xa9,
	0xd4, 0xf8, 0x2f, 0xa5, 0xf2, 0x96, 0x83, 0xdb, 0xc4, 0x5b, 0xee, 0x7e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x58, 0xf2, 0xd2, 0x9d, 0xcf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ItemPresentClient is the client API for ItemPresent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemPresentClient interface {
	// Create a new event to the event store
	GetItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ItemPresent_GetItemsClient, error)
	// Get all events for the given aggregate and event
	GetItemIds(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ItemPresent_GetItemIdsClient, error)
}

type itemPresentClient struct {
	cc *grpc.ClientConn
}

func NewItemPresentClient(cc *grpc.ClientConn) ItemPresentClient {
	return &itemPresentClient{cc}
}

func (c *itemPresentClient) GetItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ItemPresent_GetItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ItemPresent_serviceDesc.Streams[0], "/main.ItemPresent/getItems", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemPresentGetItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemPresent_GetItemsClient interface {
	Recv() (*Item, error)
	grpc.ClientStream
}

type itemPresentGetItemsClient struct {
	grpc.ClientStream
}

func (x *itemPresentGetItemsClient) Recv() (*Item, error) {
	m := new(Item)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemPresentClient) GetItemIds(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ItemPresent_GetItemIdsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ItemPresent_serviceDesc.Streams[1], "/main.ItemPresent/GetItemIds", opts...)
	if err != nil {
		return nil, err
	}
	x := &itemPresentGetItemIdsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ItemPresent_GetItemIdsClient interface {
	Recv() (*ItemIds, error)
	grpc.ClientStream
}

type itemPresentGetItemIdsClient struct {
	grpc.ClientStream
}

func (x *itemPresentGetItemIdsClient) Recv() (*ItemIds, error) {
	m := new(ItemIds)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ItemPresentServer is the server API for ItemPresent service.
type ItemPresentServer interface {
	// Create a new event to the event store
	GetItems(*empty.Empty, ItemPresent_GetItemsServer) error
	// Get all events for the given aggregate and event
	GetItemIds(*empty.Empty, ItemPresent_GetItemIdsServer) error
}

// UnimplementedItemPresentServer can be embedded to have forward compatible implementations.
type UnimplementedItemPresentServer struct {
}

func (*UnimplementedItemPresentServer) GetItems(req *empty.Empty, srv ItemPresent_GetItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (*UnimplementedItemPresentServer) GetItemIds(req *empty.Empty, srv ItemPresent_GetItemIdsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetItemIds not implemented")
}

func RegisterItemPresentServer(s *grpc.Server, srv ItemPresentServer) {
	s.RegisterService(&_ItemPresent_serviceDesc, srv)
}

func _ItemPresent_GetItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemPresentServer).GetItems(m, &itemPresentGetItemsServer{stream})
}

type ItemPresent_GetItemsServer interface {
	Send(*Item) error
	grpc.ServerStream
}

type itemPresentGetItemsServer struct {
	grpc.ServerStream
}

func (x *itemPresentGetItemsServer) Send(m *Item) error {
	return x.ServerStream.SendMsg(m)
}

func _ItemPresent_GetItemIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ItemPresentServer).GetItemIds(m, &itemPresentGetItemIdsServer{stream})
}

type ItemPresent_GetItemIdsServer interface {
	Send(*ItemIds) error
	grpc.ServerStream
}

type itemPresentGetItemIdsServer struct {
	grpc.ServerStream
}

func (x *itemPresentGetItemIdsServer) Send(m *ItemIds) error {
	return x.ServerStream.SendMsg(m)
}

var _ItemPresent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.ItemPresent",
	HandlerType: (*ItemPresentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getItems",
			Handler:       _ItemPresent_GetItems_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetItemIds",
			Handler:       _ItemPresent_GetItemIds_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "item_present.proto",
}
