// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

package main

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
	return fileDescriptor_6007f868cf6553df, []int{1}
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
	proto.RegisterType((*ItemIds)(nil), "main.ItemIds")
}

func init() { proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df) }

var fileDescriptor_6007f868cf6553df = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4a, 0xc6, 0x30,
	0x14, 0x85, 0x49, 0xcd, 0x9f, 0xea, 0x15, 0x1d, 0x2e, 0xa2, 0x77, 0x2c, 0x9d, 0xfe, 0xc9, 0xc5,
	0x57, 0xe8, 0x12, 0x10, 0x87, 0x22, 0x5d, 0x25, 0xb6, 0x77, 0x08, 0x34, 0x49, 0x49, 0xa3, 0xd0,
	0xf7, 0xf3, 0xc1, 0x24, 0x69, 0x87, 0x6e, 0xe7, 0x7e, 0x87, 0x43, 0xf2, 0x01, 0xd8, 0xc4, 0xee,
	0x75, 0x89, 0x21, 0x05, 0x94, 0xce, 0x58, 0xdf, 0xfe, 0x09, 0x90, 0x3a, 0xb1, 0xc3, 0x47, 0xa8,
	0x74, 0x47, 0xa2, 0x11, 0xd7, 0x87, 0xbe, 0xd2, 0x1d, 0x22, 0xc8, 0x0f, 0xe3, 0x98, 0xaa, 0x46,
	0x5c, 0xef, 0xfa, 0x92, 0x33, 0xd3, 0x63, 0xf0, 0x74, 0xb3, 0xb3, 0x9c, 0xb1, 0x81, 0xfb, 0x8e,
	0xd7, 0x31, 0xda, 0x25, 0xd9, 0xe0, 0x49, 0x96, 0xea, 0x8c, 0xf2, 0xea, 0x73, 0x5b, 0x98, 0x2e,
	0xfb, 0x2a, 0x67, 0x7c, 0x06, 0xd5, 0x9b, 0x68, 0xd3, 0x46, 0xaa, 0xd0, 0xe3, 0xc2, 0x27, 0xb8,
	0xbc, 0xf3, 0x2f, 0xcf, 0x54, 0x97, 0x8f, 0xec, 0x47, 0x7e, 0x63, 0x60, 0x3f, 0x85, 0x38, 0x98,
	0xf9, 0x87, 0xe9, 0xb6, 0x74, 0x67, 0xd4, 0xb6, 0x50, 0x67, 0x0b, 0x3d, 0xad, 0xf8, 0x02, 0x75,
	0xb6, 0xfc, 0xb2, 0xd3, 0x61, 0xa3, 0x6c, 0x69, 0xbe, 0x55, 0xf1, 0x7e, 0xfb, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x7b, 0x38, 0x4c, 0xad, 0x05, 0x01, 0x00, 0x00,
}
