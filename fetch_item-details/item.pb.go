// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

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
	return fileDescriptor_6007f868cf6553df, []int{0}
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

type ItemIdsRequest struct {
	BatchUUID            string   `protobuf:"bytes,1,opt,name=batchUUID,proto3" json:"batchUUID,omitempty"`
	ItemIds              []uint32 `protobuf:"varint,2,rep,packed,name=itemIds,proto3" json:"itemIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemIdsRequest) Reset()         { *m = ItemIdsRequest{} }
func (m *ItemIdsRequest) String() string { return proto.CompactTextString(m) }
func (*ItemIdsRequest) ProtoMessage()    {}
func (*ItemIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

func (m *ItemIdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemIdsRequest.Unmarshal(m, b)
}
func (m *ItemIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemIdsRequest.Marshal(b, m, deterministic)
}
func (m *ItemIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemIdsRequest.Merge(m, src)
}
func (m *ItemIdsRequest) XXX_Size() int {
	return xxx_messageInfo_ItemIdsRequest.Size(m)
}
func (m *ItemIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ItemIdsRequest proto.InternalMessageInfo

func (m *ItemIdsRequest) GetBatchUUID() string {
	if m != nil {
		return m.BatchUUID
	}
	return ""
}

func (m *ItemIdsRequest) GetItemIds() []uint32 {
	if m != nil {
		return m.ItemIds
	}
	return nil
}

type ItemBatch struct {
	BatchUUID            string   `protobuf:"bytes,1,opt,name=batchUUID,proto3" json:"batchUUID,omitempty"`
	ItemData             []byte   `protobuf:"bytes,2,opt,name=itemData,proto3" json:"itemData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemBatch) Reset()         { *m = ItemBatch{} }
func (m *ItemBatch) String() string { return proto.CompactTextString(m) }
func (*ItemBatch) ProtoMessage()    {}
func (*ItemBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{2}
}

func (m *ItemBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemBatch.Unmarshal(m, b)
}
func (m *ItemBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemBatch.Marshal(b, m, deterministic)
}
func (m *ItemBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemBatch.Merge(m, src)
}
func (m *ItemBatch) XXX_Size() int {
	return xxx_messageInfo_ItemBatch.Size(m)
}
func (m *ItemBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ItemBatch proto.InternalMessageInfo

func (m *ItemBatch) GetBatchUUID() string {
	if m != nil {
		return m.BatchUUID
	}
	return ""
}

func (m *ItemBatch) GetItemData() []byte {
	if m != nil {
		return m.ItemData
	}
	return nil
}

type ItemResponse struct {
	ItemId               uint32   `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemResponse) Reset()         { *m = ItemResponse{} }
func (m *ItemResponse) String() string { return proto.CompactTextString(m) }
func (*ItemResponse) ProtoMessage()    {}
func (*ItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{3}
}

func (m *ItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemResponse.Unmarshal(m, b)
}
func (m *ItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemResponse.Marshal(b, m, deterministic)
}
func (m *ItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemResponse.Merge(m, src)
}
func (m *ItemResponse) XXX_Size() int {
	return xxx_messageInfo_ItemResponse.Size(m)
}
func (m *ItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ItemResponse proto.InternalMessageInfo

func (m *ItemResponse) GetItemId() uint32 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "messages.Item")
	proto.RegisterType((*ItemIdsRequest)(nil), "messages.ItemIdsRequest")
	proto.RegisterType((*ItemBatch)(nil), "messages.ItemBatch")
	proto.RegisterType((*ItemResponse)(nil), "messages.ItemResponse")
}

func init() { proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df) }

var fileDescriptor_6007f868cf6553df = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0x5a, 0xa0, 0xb4, 0xf3, 0x01, 0x87, 0x8d, 0xc1, 0x0d, 0x7a, 0x68, 0x7a, 0x91, 0x53,
	0x49, 0xf4, 0xea, 0xc9, 0x94, 0x68, 0x13, 0xe3, 0x61, 0x15, 0xae, 0x66, 0x81, 0x11, 0x9b, 0xd0,
	0x6e, 0xed, 0x2e, 0x26, 0xfc, 0x3f, 0x7f, 0x98, 0x99, 0x5d, 0xd0, 0xde, 0xbc, 0xcd, 0x7b, 0xfb,
	0x66, 0xde, 0xce, 0x3c, 0x80, 0xc2, 0x60, 0x99, 0xd6, 0x8d, 0x32, 0x8a, 0x85, 0x25, 0x6a, 0x2d,
	0xb7, 0xa8, 0x27, 0x17, 0x5b, 0xa5, 0xb6, 0x3b, 0x9c, 0x59, 0x7e, 0xb5, 0x7f, 0x9b, 0x61, 0x59,
	0x9b, 0x83, 0x93, 0x25, 0x5f, 0x1e, 0x74, 0x73, 0x83, 0x25, 0x1b, 0x81, 0x9f, 0x67, 0xdc, 0x8b,
	0xbd, 0xe9, 0x50, 0xf8, 0x79, 0xc6, 0x18, 0x74, 0x9f, 0x64, 0x89, 0xdc, 0x8f, 0xbd, 0x69, 0x24,
	0x6c, 0x4d, 0x5c, 0xbe, 0x56, 0x15, 0xef, 0x38, 0x8e, 0x6a, 0x16, 0xc3, 0xff, 0x0c, 0xf5, 0xba,
	0x29, 0x6a, 0x53, 0xa8, 0x8a, 0x77, 0xed, 0x53, 0x9b, 0xa2, 0xae, 0x97, 0x43, 0x8d, 0xbc, 0xe7,
	0xba, 0xa8, 0x66, 0x63, 0x08, 0x84, 0x6c, 0x0a, 0x73, 0xe0, 0x81, 0x65, 0x8f, 0x88, 0x9d, 0x41,
	0xef, 0x11, 0x3f, 0x71, 0xc7, 0xfb, 0xf6, 0x23, 0x0e, 0x90, 0xc7, 0x12, 0xab, 0x8d, 0x6a, 0x96,
	0x72, 0xb7, 0x47, 0x1e, 0xda, 0xb7, 0x36, 0x95, 0x3c, 0xc0, 0x88, 0xb6, 0xc8, 0x37, 0x5a, 0xe0,
	0xc7, 0x1e, 0xb5, 0x61, 0x97, 0x10, 0xad, 0xa4, 0x59, 0xbf, 0x2f, 0x16, 0xc7, 0xb5, 0x22, 0xf1,
	0x4b, 0x30, 0x0e, 0xfd, 0xc2, 0xe9, 0xb9, 0x1f, 0x77, 0xa6, 0x43, 0x71, 0x82, 0xc9, 0x1c, 0x22,
	0x9a, 0x74, 0x47, 0xd2, 0x3f, 0x86, 0x4c, 0x20, 0xa4, 0xae, 0x4c, 0x1a, 0x69, 0xcf, 0x34, 0x10,
	0x3f, 0x38, 0xb9, 0x82, 0x01, 0x8d, 0x11, 0xa8, 0x6b, 0x55, 0x69, 0x64, 0xe7, 0xce, 0xf0, 0xb5,
	0xd8, 0x1c, 0x6f, 0x1c, 0x38, 0xc3, 0xeb, 0xdc, 0xf9, 0x3d, 0x1b, 0xd5, 0x20, 0xbb, 0x85, 0xf0,
	0x1e, 0x0d, 0x61, 0xcd, 0xc6, 0xa9, 0xcb, 0x2d, 0x3d, 0xe5, 0x96, 0xce, 0x29, 0xb7, 0xc9, 0x38,
	0x3d, 0x25, 0x9b, 0xb6, 0x1d, 0x92, 0x7f, 0xab, 0xc0, 0x2a, 0x6f, 0xbe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xb3, 0x48, 0x64, 0x7d, 0x07, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ItemStoreClient is the client API for ItemStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemStoreClient interface {
	// Get all events for the given aggregate and event
	GetItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ItemResponse, error)
}

type itemStoreClient struct {
	cc *grpc.ClientConn
}

func NewItemStoreClient(cc *grpc.ClientConn) ItemStoreClient {
	return &itemStoreClient{cc}
}

func (c *itemStoreClient) GetItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ItemResponse, error) {
	out := new(ItemResponse)
	err := c.cc.Invoke(ctx, "/messages.ItemStore/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemStoreServer is the server API for ItemStore service.
type ItemStoreServer interface {
	// Get all events for the given aggregate and event
	GetItems(context.Context, *empty.Empty) (*ItemResponse, error)
}

// UnimplementedItemStoreServer can be embedded to have forward compatible implementations.
type UnimplementedItemStoreServer struct {
}

func (*UnimplementedItemStoreServer) GetItems(ctx context.Context, req *empty.Empty) (*ItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}

func RegisterItemStoreServer(s *grpc.Server, srv ItemStoreServer) {
	s.RegisterService(&_ItemStore_serviceDesc, srv)
}

func _ItemStore_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemStoreServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.ItemStore/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemStoreServer).GetItems(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.ItemStore",
	HandlerType: (*ItemStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItems",
			Handler:    _ItemStore_GetItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "item.proto",
}
