// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wire.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	wire.proto

It has these top-level messages:
	Contact
	Store
	Login
	PlainText
	Wire
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Identifies which field is filled in
type Wire_Which int32

const (
	Wire_LOGIN               Wire_Which = 0
	Wire_CONTACTS            Wire_Which = 1
	Wire_PRESENCE            Wire_Which = 2
	Wire_STORE               Wire_Which = 3
	Wire_LOAD                Wire_Which = 4
	Wire_PUBLIC_KEY          Wire_Which = 5
	Wire_PUBLIC_KEY_RESPONSE Wire_Which = 6
	Wire_HANDSHAKE           Wire_Which = 7
	Wire_PAYLOAD             Wire_Which = 8
	Wire_LOGIN_RESPONSE      Wire_Which = 9
	Wire_PLAIN_TEXT          Wire_Which = 10
)

var Wire_Which_name = map[int32]string{
	0:  "LOGIN",
	1:  "CONTACTS",
	2:  "PRESENCE",
	3:  "STORE",
	4:  "LOAD",
	5:  "PUBLIC_KEY",
	6:  "PUBLIC_KEY_RESPONSE",
	7:  "HANDSHAKE",
	8:  "PAYLOAD",
	9:  "LOGIN_RESPONSE",
	10: "PLAIN_TEXT",
}
var Wire_Which_value = map[string]int32{
	"LOGIN":               0,
	"CONTACTS":            1,
	"PRESENCE":            2,
	"STORE":               3,
	"LOAD":                4,
	"PUBLIC_KEY":          5,
	"PUBLIC_KEY_RESPONSE": 6,
	"HANDSHAKE":           7,
	"PAYLOAD":             8,
	"LOGIN_RESPONSE":      9,
	"PLAIN_TEXT":          10,
}

func (x Wire_Which) String() string {
	return proto.EnumName(Wire_Which_name, int32(x))
}
func (Wire_Which) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Contact struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Online      bool   `protobuf:"varint,3,opt,name=online" json:"online,omitempty"`
	DeviceToken string `protobuf:"bytes,4,opt,name=deviceToken" json:"deviceToken,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Contact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

func (m *Contact) GetDeviceToken() string {
	if m != nil {
		return m.DeviceToken
	}
	return ""
}

type Store struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Store) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type Login struct {
	Type        uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty"`
	AuthenToken string `protobuf:"bytes,3,opt,name=authenToken" json:"authenToken,omitempty"`
	DeviceToken string `protobuf:"bytes,4,opt,name=deviceToken" json:"deviceToken,omitempty"`
}

func (m *Login) Reset()                    { *m = Login{} }
func (m *Login) String() string            { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()               {}
func (*Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Login) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Login) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Login) GetAuthenToken() string {
	if m != nil {
		return m.AuthenToken
	}
	return ""
}

func (m *Login) GetDeviceToken() string {
	if m != nil {
		return m.DeviceToken
	}
	return ""
}

type PlainText struct {
	Type    uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *PlainText) Reset()                    { *m = PlainText{} }
func (m *PlainText) String() string            { return proto.CompactTextString(m) }
func (*PlainText) ProtoMessage()               {}
func (*PlainText) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PlainText) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PlainText) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Wire struct {
	Version   uint32     `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	SessionId string     `protobuf:"bytes,2,opt,name=sessionId" json:"sessionId,omitempty"`
	From      string     `protobuf:"bytes,3,opt,name=from" json:"from,omitempty"`
	To        string     `protobuf:"bytes,4,opt,name=to" json:"to,omitempty"`
	Which     Wire_Which `protobuf:"varint,5,opt,name=which,enum=Wire_Which" json:"which,omitempty"`
	// One of the following will be filled in
	Login     *Login     `protobuf:"bytes,101,opt,name=login" json:"login,omitempty"`
	Contacts  []*Contact `protobuf:"bytes,102,rep,name=contacts" json:"contacts,omitempty"`
	Store     *Store     `protobuf:"bytes,104,opt,name=store" json:"store,omitempty"`
	Payload   []byte     `protobuf:"bytes,106,opt,name=payload,proto3" json:"payload,omitempty"`
	PlainText *PlainText `protobuf:"bytes,107,opt,name=plainText" json:"plainText,omitempty"`
}

func (m *Wire) Reset()                    { *m = Wire{} }
func (m *Wire) String() string            { return proto.CompactTextString(m) }
func (*Wire) ProtoMessage()               {}
func (*Wire) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Wire) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Wire) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Wire) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Wire) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Wire) GetWhich() Wire_Which {
	if m != nil {
		return m.Which
	}
	return Wire_LOGIN
}

func (m *Wire) GetLogin() *Login {
	if m != nil {
		return m.Login
	}
	return nil
}

func (m *Wire) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *Wire) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *Wire) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Wire) GetPlainText() *PlainText {
	if m != nil {
		return m.PlainText
	}
	return nil
}

func init() {
	proto.RegisterType((*Contact)(nil), "Contact")
	proto.RegisterType((*Store)(nil), "Store")
	proto.RegisterType((*Login)(nil), "Login")
	proto.RegisterType((*PlainText)(nil), "PlainText")
	proto.RegisterType((*Wire)(nil), "Wire")
	proto.RegisterEnum("Wire_Which", Wire_Which_name, Wire_Which_value)
}

func init() { proto.RegisterFile("wire.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0xb1, 0x9d, 0xd8, 0x93, 0x36, 0xb2, 0x16, 0x09, 0x16, 0x54, 0x09, 0x13, 0x21, 0xe1,
	0x93, 0x0f, 0xe5, 0xc4, 0xd1, 0x4d, 0x2d, 0x1a, 0x35, 0x72, 0xa2, 0xb5, 0x51, 0x29, 0x97, 0xc8,
	0xb5, 0xb7, 0xcd, 0x92, 0x64, 0x37, 0xb2, 0x37, 0x2d, 0x11, 0xbf, 0xc4, 0x17, 0xf1, 0x35, 0x68,
	0xd7, 0x4e, 0xd2, 0x03, 0x07, 0x6e, 0xf3, 0xde, 0xec, 0xcc, 0x7b, 0x9e, 0x27, 0x03, 0x3c, 0xb1,
	0x8a, 0x86, 0x9b, 0x4a, 0x48, 0x31, 0x7c, 0x80, 0xde, 0x48, 0x70, 0x99, 0x17, 0x12, 0x0d, 0xa0,
	0xc3, 0x4a, 0x6c, 0xf8, 0x46, 0xe0, 0x92, 0x0e, 0x2b, 0x11, 0x02, 0x8b, 0xe7, 0x6b, 0x8a, 0x3b,
	0x9a, 0xd1, 0x35, 0x7a, 0x05, 0x5d, 0xc1, 0x57, 0x8c, 0x53, 0x6c, 0xfa, 0x46, 0xe0, 0x90, 0x16,
	0x21, 0x1f, 0xfa, 0x25, 0x7d, 0x64, 0x05, 0xcd, 0xc4, 0x92, 0x72, 0x6c, 0xe9, 0x91, 0xe7, 0xd4,
	0xf0, 0x0d, 0xd8, 0xa9, 0x14, 0x15, 0x45, 0x1e, 0x98, 0x4b, 0xba, 0xd3, 0x3a, 0x27, 0x44, 0x95,
	0xc3, 0x5f, 0x60, 0x4f, 0xc4, 0x03, 0xe3, 0x4a, 0x51, 0xee, 0x36, 0x54, 0xf7, 0x4e, 0x89, 0xae,
	0xd1, 0x5b, 0x70, 0xb6, 0x35, 0xad, 0x92, 0xa3, 0x93, 0x03, 0x56, 0xaa, 0xf9, 0x56, 0x2e, 0x28,
	0x6f, 0x54, 0xcd, 0x46, 0xf5, 0x19, 0xf5, 0x1f, 0xbe, 0x3e, 0x83, 0x3b, 0x5b, 0xe5, 0x8c, 0x67,
	0xf4, 0xa7, 0xfc, 0xa7, 0x01, 0x0c, 0xbd, 0x42, 0x70, 0x49, 0xb9, 0x6c, 0xf5, 0xf7, 0x70, 0xf8,
	0xc7, 0x04, 0xeb, 0x86, 0x55, 0xfa, 0xc9, 0x23, 0xad, 0x6a, 0x26, 0x78, 0x3b, 0xb9, 0x87, 0xe8,
	0x0c, 0xdc, 0x9a, 0xd6, 0xaa, 0x1c, 0x97, 0xed, 0xf8, 0x91, 0x50, 0x72, 0xf7, 0x95, 0x58, 0xb7,
	0xc6, 0x75, 0xad, 0x52, 0x90, 0xa2, 0x35, 0xda, 0x91, 0x02, 0xbd, 0x07, 0xfb, 0x69, 0xc1, 0x8a,
	0x05, 0xb6, 0x7d, 0x23, 0x18, 0x9c, 0xf7, 0x43, 0xa5, 0x18, 0xde, 0x28, 0x8a, 0x34, 0x1d, 0x74,
	0x06, 0xf6, 0x4a, 0xdd, 0x0f, 0x53, 0xdf, 0x08, 0xfa, 0xe7, 0xdd, 0x50, 0x5f, 0x93, 0x34, 0x24,
	0xfa, 0x00, 0x4e, 0xd1, 0x24, 0x5c, 0xe3, 0x7b, 0xdf, 0x0c, 0xfa, 0xe7, 0x4e, 0xd8, 0x46, 0x4e,
	0x0e, 0x1d, 0xb5, 0xa3, 0x56, 0xf1, 0xe0, 0x45, 0xbb, 0x43, 0x87, 0x45, 0x1a, 0x52, 0x7d, 0xe0,
	0x26, 0xdf, 0xad, 0x44, 0x5e, 0xe2, 0x1f, 0x3a, 0xb7, 0x3d, 0x44, 0x01, 0xb8, 0x9b, 0xfd, 0xf9,
	0xf0, 0x52, 0xcf, 0x42, 0x78, 0x38, 0x28, 0x39, 0x36, 0x87, 0xbf, 0x0d, 0xb0, 0xb5, 0x6d, 0xe4,
	0x82, 0x3d, 0x99, 0x7e, 0x19, 0x27, 0xde, 0x0b, 0x74, 0x02, 0xce, 0x68, 0x9a, 0x64, 0xd1, 0x28,
	0x4b, 0x3d, 0x43, 0xa1, 0x19, 0x89, 0xd3, 0x38, 0x19, 0xc5, 0x5e, 0x47, 0x3d, 0x4b, 0xb3, 0x29,
	0x89, 0x3d, 0x13, 0x39, 0x60, 0x4d, 0xa6, 0xd1, 0xa5, 0x67, 0xa1, 0x01, 0xc0, 0xec, 0xeb, 0xc5,
	0x64, 0x3c, 0x9a, 0x5f, 0xc7, 0xb7, 0x9e, 0x8d, 0x5e, 0xc3, 0xcb, 0x23, 0x9e, 0x93, 0x38, 0x9d,
	0x4d, 0x93, 0x34, 0xf6, 0xba, 0xe8, 0x14, 0xdc, 0xab, 0x28, 0xb9, 0x4c, 0xaf, 0xa2, 0xeb, 0xd8,
	0xeb, 0xa1, 0x3e, 0xf4, 0x66, 0xd1, 0xad, 0x5e, 0xe2, 0x20, 0x04, 0x03, 0x6d, 0xe0, 0xf8, 0xde,
	0xd5, 0x8b, 0x27, 0xd1, 0x38, 0x99, 0x67, 0xf1, 0xb7, 0xcc, 0x83, 0x8b, 0x8f, 0xf0, 0xae, 0xa2,
	0x65, 0x28, 0xe9, 0x2a, 0x2c, 0x16, 0xb9, 0x0c, 0x1f, 0x28, 0xa7, 0x55, 0x2e, 0x69, 0x39, 0xd7,
	0x3f, 0xce, 0xdd, 0xf6, 0xfe, 0xbb, 0xb5, 0xce, 0x19, 0xbf, 0xeb, 0x6a, 0xfc, 0xe9, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xce, 0x7b, 0x8b, 0xcf, 0x56, 0x03, 0x00, 0x00,
}
