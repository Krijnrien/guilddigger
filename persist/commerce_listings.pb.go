// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commerce_listings.proto

package persist

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

type Listings struct {
	Id                   int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Buys                 []*Listings_BUYS  `protobuf:"bytes,2,rep,name=buys,proto3" json:"buys,omitempty"`
	Sells                []*Listings_SELLS `protobuf:"bytes,3,rep,name=sells,proto3" json:"sells,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Listings) Reset()         { *m = Listings{} }
func (m *Listings) String() string { return proto.CompactTextString(m) }
func (*Listings) ProtoMessage()    {}
func (*Listings) Descriptor() ([]byte, []int) {
	return fileDescriptor_472404e8ab389c17, []int{0}
}

func (m *Listings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listings.Unmarshal(m, b)
}
func (m *Listings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listings.Marshal(b, m, deterministic)
}
func (m *Listings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listings.Merge(m, src)
}
func (m *Listings) XXX_Size() int {
	return xxx_messageInfo_Listings.Size(m)
}
func (m *Listings) XXX_DiscardUnknown() {
	xxx_messageInfo_Listings.DiscardUnknown(m)
}

var xxx_messageInfo_Listings proto.InternalMessageInfo

func (m *Listings) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Listings) GetBuys() []*Listings_BUYS {
	if m != nil {
		return m.Buys
	}
	return nil
}

func (m *Listings) GetSells() []*Listings_SELLS {
	if m != nil {
		return m.Sells
	}
	return nil
}

type Listings_BUYS struct {
	Listings             int32    `protobuf:"varint,1,opt,name=listings,proto3" json:"listings,omitempty"`
	UnitPrice            int32    `protobuf:"varint,2,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listings_BUYS) Reset()         { *m = Listings_BUYS{} }
func (m *Listings_BUYS) String() string { return proto.CompactTextString(m) }
func (*Listings_BUYS) ProtoMessage()    {}
func (*Listings_BUYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_472404e8ab389c17, []int{0, 0}
}

func (m *Listings_BUYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listings_BUYS.Unmarshal(m, b)
}
func (m *Listings_BUYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listings_BUYS.Marshal(b, m, deterministic)
}
func (m *Listings_BUYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listings_BUYS.Merge(m, src)
}
func (m *Listings_BUYS) XXX_Size() int {
	return xxx_messageInfo_Listings_BUYS.Size(m)
}
func (m *Listings_BUYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Listings_BUYS.DiscardUnknown(m)
}

var xxx_messageInfo_Listings_BUYS proto.InternalMessageInfo

func (m *Listings_BUYS) GetListings() int32 {
	if m != nil {
		return m.Listings
	}
	return 0
}

func (m *Listings_BUYS) GetUnitPrice() int32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *Listings_BUYS) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type Listings_SELLS struct {
	Listings             int32    `protobuf:"varint,1,opt,name=listings,proto3" json:"listings,omitempty"`
	UnitPrice            int32    `protobuf:"varint,2,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listings_SELLS) Reset()         { *m = Listings_SELLS{} }
func (m *Listings_SELLS) String() string { return proto.CompactTextString(m) }
func (*Listings_SELLS) ProtoMessage()    {}
func (*Listings_SELLS) Descriptor() ([]byte, []int) {
	return fileDescriptor_472404e8ab389c17, []int{0, 1}
}

func (m *Listings_SELLS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listings_SELLS.Unmarshal(m, b)
}
func (m *Listings_SELLS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listings_SELLS.Marshal(b, m, deterministic)
}
func (m *Listings_SELLS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listings_SELLS.Merge(m, src)
}
func (m *Listings_SELLS) XXX_Size() int {
	return xxx_messageInfo_Listings_SELLS.Size(m)
}
func (m *Listings_SELLS) XXX_DiscardUnknown() {
	xxx_messageInfo_Listings_SELLS.DiscardUnknown(m)
}

var xxx_messageInfo_Listings_SELLS proto.InternalMessageInfo

func (m *Listings_SELLS) GetListings() int32 {
	if m != nil {
		return m.Listings
	}
	return 0
}

func (m *Listings_SELLS) GetUnitPrice() int32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *Listings_SELLS) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func init() {
	proto.RegisterType((*Listings)(nil), "messages.Listings")
	proto.RegisterType((*Listings_BUYS)(nil), "messages.Listings.BUYS")
	proto.RegisterType((*Listings_SELLS)(nil), "messages.Listings.SELLS")
}

func init() { proto.RegisterFile("commerce_listings.proto", fileDescriptor_472404e8ab389c17) }

var fileDescriptor_472404e8ab389c17 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0xcf, 0xcd,
	0x4d, 0x2d, 0x4a, 0x4e, 0x8d, 0xcf, 0xc9, 0x2c, 0x2e, 0xc9, 0xcc, 0x4b, 0x2f, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x56, 0xda, 0xc6,
	0xc4, 0xc5, 0xe1, 0x03, 0x95, 0x14, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0xd2, 0xe6, 0x62, 0x49, 0x2a, 0xad, 0x2c, 0x96, 0x60, 0x52,
	0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd7, 0x83, 0xe9, 0xd2, 0x83, 0xe9, 0xd0, 0x73, 0x0a, 0x8d, 0x0c,
	0x0e, 0x02, 0x2b, 0x12, 0xd2, 0xe3, 0x62, 0x2d, 0x4e, 0xcd, 0xc9, 0x29, 0x96, 0x60, 0x06, 0xab,
	0x96, 0xc0, 0xa2, 0x3a, 0xd8, 0xd5, 0xc7, 0x27, 0x38, 0x08, 0xa2, 0x4c, 0x2a, 0x96, 0x8b, 0x05,
	0xa4, 0x5b, 0x48, 0x8a, 0x8b, 0x03, 0xe6, 0x3a, 0xa8, 0xd5, 0x70, 0xbe, 0x90, 0x2c, 0x17, 0x57,
	0x69, 0x5e, 0x66, 0x49, 0x7c, 0x41, 0x51, 0x66, 0x72, 0xaa, 0x04, 0x13, 0x58, 0x96, 0x13, 0x24,
	0x12, 0x00, 0x12, 0x00, 0x69, 0x2d, 0x2c, 0x4d, 0xcc, 0x2b, 0xc9, 0x2c, 0xa9, 0x94, 0x60, 0x86,
	0x68, 0x85, 0xf1, 0xa5, 0xe2, 0xb8, 0x58, 0xc1, 0xd6, 0xd1, 0xc8, 0xfc, 0x24, 0x36, 0x70, 0x48,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x76, 0x1c, 0xad, 0x64, 0x01, 0x00, 0x00,
}