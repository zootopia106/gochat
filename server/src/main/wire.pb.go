// Code generated by protoc-gen-go.
// source: wire.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	wire.proto

It has these top-level messages:
	Login
	Contact
	Text
	File
	Time
	Timestamp
	Image
	FormatDescription
	VideoSample
	AudioSample
	Av
	Haber
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
type Haber_Which int32

const (
	Haber_LOGIN    Haber_Which = 0
	Haber_CONTACTS Haber_Which = 1
	Haber_PRESENCE Haber_Which = 2
	Haber_TEXT     Haber_Which = 3
	Haber_FILE     Haber_Which = 4
	Haber_AV       Haber_Which = 5
)

var Haber_Which_name = map[int32]string{
	0: "LOGIN",
	1: "CONTACTS",
	2: "PRESENCE",
	3: "TEXT",
	4: "FILE",
	5: "AV",
}
var Haber_Which_value = map[string]int32{
	"LOGIN":    0,
	"CONTACTS": 1,
	"PRESENCE": 2,
	"TEXT":     3,
	"FILE":     4,
	"AV":       5,
}

func (x Haber_Which) String() string {
	return proto.EnumName(Haber_Which_name, int32(x))
}
func (Haber_Which) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

type Login struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *Login) Reset()                    { *m = Login{} }
func (m *Login) String() string            { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()               {}
func (*Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Login) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Contact struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Online bool   `protobuf:"varint,2,opt,name=online" json:"online,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

type Text struct {
	Body string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
}

func (m *Text) Reset()                    { *m = Text{} }
func (m *Text) String() string            { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()               {}
func (*Text) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Text) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type File struct {
	Key  string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *File) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Time struct {
	Value int64  `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Scale int32  `protobuf:"varint,2,opt,name=scale" json:"scale,omitempty"`
	Flags uint32 `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	Epoch int64  `protobuf:"varint,4,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Time) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Time) GetScale() int32 {
	if m != nil {
		return m.Scale
	}
	return 0
}

func (m *Time) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *Time) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

type Timestamp struct {
	Duration     *Time `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
	Presentation *Time `protobuf:"bytes,2,opt,name=presentation" json:"presentation,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Timestamp) GetDuration() *Time {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Timestamp) GetPresentation() *Time {
	if m != nil {
		return m.Presentation
	}
	return nil
}

type Image struct {
	Width       int64             `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height      int64             `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Format      uint32            `protobuf:"varint,3,opt,name=format" json:"format,omitempty"`
	Attachments map[string]string `protobuf:"bytes,4,rep,name=attachments" json:"attachments,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Data        []byte            `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Image) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetFormat() uint32 {
	if m != nil {
		return m.Format
	}
	return 0
}

func (m *Image) GetAttachments() map[string]string {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FormatDescription struct {
	MediaType    uint32            `protobuf:"varint,1,opt,name=mediaType" json:"mediaType,omitempty"`
	MediaSubtype uint32            `protobuf:"varint,2,opt,name=mediaSubtype" json:"mediaSubtype,omitempty"`
	Extensions   map[string]string `protobuf:"bytes,3,rep,name=extensions" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FormatDescription) Reset()                    { *m = FormatDescription{} }
func (m *FormatDescription) String() string            { return proto.CompactTextString(m) }
func (*FormatDescription) ProtoMessage()               {}
func (*FormatDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FormatDescription) GetMediaType() uint32 {
	if m != nil {
		return m.MediaType
	}
	return 0
}

func (m *FormatDescription) GetMediaSubtype() uint32 {
	if m != nil {
		return m.MediaSubtype
	}
	return 0
}

func (m *FormatDescription) GetExtensions() map[string]string {
	if m != nil {
		return m.Extensions
	}
	return nil
}

type VideoSample struct {
	Timestamp *Timestamp         `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Image     *Image             `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	Format    *FormatDescription `protobuf:"bytes,3,opt,name=format" json:"format,omitempty"`
}

func (m *VideoSample) Reset()                    { *m = VideoSample{} }
func (m *VideoSample) String() string            { return proto.CompactTextString(m) }
func (*VideoSample) ProtoMessage()               {}
func (*VideoSample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *VideoSample) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *VideoSample) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *VideoSample) GetFormat() *FormatDescription {
	if m != nil {
		return m.Format
	}
	return nil
}

type AudioSample struct {
}

func (m *AudioSample) Reset()                    { *m = AudioSample{} }
func (m *AudioSample) String() string            { return proto.CompactTextString(m) }
func (*AudioSample) ProtoMessage()               {}
func (*AudioSample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type Av struct {
	Video *VideoSample `protobuf:"bytes,1,opt,name=video" json:"video,omitempty"`
	Audio *AudioSample `protobuf:"bytes,2,opt,name=audio" json:"audio,omitempty"`
}

func (m *Av) Reset()                    { *m = Av{} }
func (m *Av) String() string            { return proto.CompactTextString(m) }
func (*Av) ProtoMessage()               {}
func (*Av) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Av) GetVideo() *VideoSample {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *Av) GetAudio() *AudioSample {
	if m != nil {
		return m.Audio
	}
	return nil
}

type Haber struct {
	Version   uint32      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	SessionId string      `protobuf:"bytes,2,opt,name=sessionId" json:"sessionId,omitempty"`
	From      string      `protobuf:"bytes,3,opt,name=from" json:"from,omitempty"`
	To        string      `protobuf:"bytes,4,opt,name=to" json:"to,omitempty"`
	Which     Haber_Which `protobuf:"varint,5,opt,name=which,enum=Haber_Which" json:"which,omitempty"`
	// One of the following will be filled in
	Login    *Login     `protobuf:"bytes,101,opt,name=login" json:"login,omitempty"`
	Contacts []*Contact `protobuf:"bytes,102,rep,name=contacts" json:"contacts,omitempty"`
	Text     *Text      `protobuf:"bytes,104,opt,name=text" json:"text,omitempty"`
	Av       *Av        `protobuf:"bytes,105,opt,name=av" json:"av,omitempty"`
	File     *File      `protobuf:"bytes,106,opt,name=file" json:"file,omitempty"`
}

func (m *Haber) Reset()                    { *m = Haber{} }
func (m *Haber) String() string            { return proto.CompactTextString(m) }
func (*Haber) ProtoMessage()               {}
func (*Haber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Haber) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Haber) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Haber) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Haber) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Haber) GetWhich() Haber_Which {
	if m != nil {
		return m.Which
	}
	return Haber_LOGIN
}

func (m *Haber) GetLogin() *Login {
	if m != nil {
		return m.Login
	}
	return nil
}

func (m *Haber) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *Haber) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Haber) GetAv() *Av {
	if m != nil {
		return m.Av
	}
	return nil
}

func (m *Haber) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*Login)(nil), "Login")
	proto.RegisterType((*Contact)(nil), "Contact")
	proto.RegisterType((*Text)(nil), "Text")
	proto.RegisterType((*File)(nil), "File")
	proto.RegisterType((*Time)(nil), "Time")
	proto.RegisterType((*Timestamp)(nil), "Timestamp")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*FormatDescription)(nil), "FormatDescription")
	proto.RegisterType((*VideoSample)(nil), "VideoSample")
	proto.RegisterType((*AudioSample)(nil), "AudioSample")
	proto.RegisterType((*Av)(nil), "Av")
	proto.RegisterType((*Haber)(nil), "Haber")
	proto.RegisterEnum("Haber_Which", Haber_Which_name, Haber_Which_value)
}

func init() { proto.RegisterFile("wire.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x8f, 0xd3, 0x3a,
	0x10, 0x7e, 0x49, 0x93, 0x6e, 0x3b, 0x6d, 0xf7, 0xe5, 0xf9, 0x21, 0x08, 0xd5, 0x1e, 0x8a, 0xe1,
	0x50, 0x10, 0xca, 0xa1, 0x08, 0x09, 0x90, 0x40, 0xea, 0x96, 0x2e, 0x14, 0x55, 0xbb, 0xc8, 0xad,
	0x96, 0x1f, 0x07, 0x24, 0xb7, 0x71, 0x1b, 0x43, 0x7e, 0x54, 0x89, 0xdb, 0xed, 0x1e, 0xf9, 0x2f,
	0x39, 0x70, 0xe2, 0x2f, 0x41, 0xb6, 0x93, 0xb4, 0xbb, 0xcb, 0x85, 0xdb, 0x7c, 0xdf, 0x8c, 0x3d,
	0x9f, 0x67, 0xc6, 0x03, 0x70, 0xc1, 0x53, 0xe6, 0xad, 0xd2, 0x44, 0x24, 0xf8, 0x3e, 0xd8, 0xe3,
	0x64, 0xc9, 0x63, 0xd4, 0x86, 0xda, 0x3a, 0x63, 0x69, 0x4c, 0x23, 0xe6, 0x1a, 0x1d, 0xa3, 0x5b,
	0x27, 0x25, 0xc6, 0x4f, 0xe1, 0x60, 0x90, 0xc4, 0x82, 0xce, 0x05, 0x42, 0x60, 0xed, 0x85, 0x28,
	0x1b, 0xdd, 0x86, 0x6a, 0x12, 0x87, 0x3c, 0x66, 0xae, 0xd9, 0x31, 0xba, 0x35, 0x92, 0x23, 0xdc,
	0x06, 0x6b, 0xca, 0xb6, 0xea, 0xcc, 0x2c, 0xf1, 0x2f, 0x8b, 0x33, 0xd2, 0xc6, 0x8f, 0xc1, 0x3a,
	0xe1, 0x21, 0x43, 0x0e, 0x54, 0xbe, 0xb1, 0xc2, 0x25, 0x4d, 0x19, 0xed, 0x53, 0x41, 0xd5, 0x5d,
	0x4d, 0xa2, 0x6c, 0xfc, 0x05, 0xac, 0x29, 0x8f, 0x18, 0xba, 0x05, 0xf6, 0x86, 0x86, 0x6b, 0x9d,
	0xbe, 0x42, 0x34, 0x90, 0x6c, 0x36, 0xa7, 0xa1, 0x4e, 0x6f, 0x13, 0x0d, 0x24, 0xbb, 0x08, 0xe9,
	0x32, 0x73, 0x2b, 0x1d, 0xa3, 0xdb, 0x22, 0x1a, 0x48, 0x96, 0xad, 0x92, 0x79, 0xe0, 0x5a, 0xfa,
	0x06, 0x05, 0xf0, 0x27, 0xa8, 0xcb, 0xfb, 0x33, 0x41, 0xa3, 0x15, 0xba, 0x07, 0x35, 0x7f, 0x9d,
	0x52, 0xc1, 0x93, 0x58, 0xe5, 0x69, 0xf4, 0x6c, 0x4f, 0x7a, 0x49, 0x49, 0xa3, 0x87, 0xd0, 0x5c,
	0xa5, 0x2c, 0x63, 0xb1, 0xd0, 0x61, 0xe6, 0x7e, 0xd8, 0x15, 0x17, 0xfe, 0x69, 0x80, 0x3d, 0x8a,
	0xe8, 0x52, 0x09, 0xba, 0xe0, 0xbe, 0x08, 0x0a, 0xf1, 0x0a, 0xc8, 0xe2, 0x05, 0x8c, 0x2f, 0x03,
	0xa1, 0x2e, 0xa9, 0x90, 0x1c, 0x49, 0x7e, 0x91, 0xa4, 0x11, 0x15, 0xb9, 0xfe, 0x1c, 0xa1, 0xe7,
	0xd0, 0xa0, 0x42, 0xd0, 0x79, 0x10, 0xb1, 0x58, 0x64, 0xae, 0xd5, 0xa9, 0x74, 0x1b, 0xbd, 0x3b,
	0x9e, 0x4a, 0xe1, 0xf5, 0x77, 0x9e, 0x61, 0x2c, 0xd2, 0x4b, 0xb2, 0x1f, 0x5b, 0x56, 0xd6, 0xde,
	0x55, 0xb6, 0xfd, 0x0a, 0x9c, 0xeb, 0x87, 0xfe, 0xd0, 0x93, 0xb2, 0xee, 0xa6, 0xe2, 0x34, 0x78,
	0x61, 0x3e, 0x33, 0xf0, 0x0f, 0x03, 0xfe, 0x3b, 0x51, 0xca, 0x5e, 0xb3, 0x6c, 0x9e, 0xf2, 0x95,
	0xaa, 0xcf, 0x11, 0xd4, 0x23, 0xe6, 0x73, 0x3a, 0xbd, 0x5c, 0xe9, 0x5e, 0xb5, 0xc8, 0x8e, 0x40,
	0x18, 0x9a, 0x0a, 0x4c, 0xd6, 0x33, 0x21, 0x03, 0x4c, 0x15, 0x70, 0x85, 0x43, 0xc7, 0x00, 0x6c,
	0x2b, 0x58, 0x9c, 0xf1, 0x24, 0x96, 0x2d, 0x94, 0xaf, 0xc4, 0xde, 0x8d, 0x4c, 0xde, 0xb0, 0x0c,
	0xd2, 0x0f, 0xde, 0x3b, 0xd5, 0x7e, 0x09, 0xff, 0x5e, 0x73, 0xff, 0xd5, 0xd3, 0xbe, 0x1b, 0xd0,
	0x38, 0xe7, 0x3e, 0x4b, 0x26, 0x34, 0x5a, 0x85, 0x0c, 0x75, 0xa1, 0x2e, 0x8a, 0x21, 0xc9, 0x07,
	0x03, 0xbc, 0x72, 0x6c, 0xc8, 0xce, 0x89, 0x8e, 0xc0, 0xe6, 0xb2, 0x1f, 0xf9, 0x5c, 0x54, 0x75,
	0x77, 0x88, 0x26, 0xd1, 0xa3, 0x2b, 0x9d, 0x6d, 0xf4, 0xd0, 0xcd, 0x67, 0x15, 0xdd, 0xc6, 0x2d,
	0x68, 0xf4, 0xd7, 0x3e, 0xcf, 0x25, 0xe0, 0x31, 0x98, 0xfd, 0x0d, 0xc2, 0x60, 0x6f, 0xa4, 0xae,
	0x5c, 0x44, 0xd3, 0xdb, 0x53, 0x49, 0xb4, 0x4b, 0xc6, 0x50, 0x79, 0x30, 0x97, 0xd0, 0xf4, 0xf6,
	0xae, 0x21, 0xda, 0x85, 0x7f, 0x99, 0x60, 0xbf, 0xa5, 0x33, 0x96, 0x22, 0x17, 0x0e, 0x36, 0x2c,
	0xcd, 0x8a, 0x89, 0x6f, 0x91, 0x02, 0xca, 0x4e, 0x66, 0x2c, 0x93, 0xe6, 0xc8, 0xcf, 0x4b, 0xb4,
	0x23, 0xe4, 0x44, 0x2d, 0xd2, 0x24, 0x52, 0x0f, 0xa9, 0x13, 0x65, 0xa3, 0x43, 0x30, 0x45, 0xa2,
	0xbe, 0x57, 0x9d, 0x98, 0x42, 0x29, 0xb9, 0x08, 0xf8, 0x3c, 0x50, 0x63, 0x77, 0xd8, 0x6b, 0x7a,
	0x2a, 0xa5, 0xf7, 0x41, 0x72, 0x44, 0xbb, 0x64, 0xc1, 0x42, 0xb9, 0x85, 0x5c, 0x96, 0x17, 0x4c,
	0xed, 0x24, 0xa2, 0x49, 0xf4, 0x00, 0x6a, 0x73, 0xbd, 0x7e, 0x32, 0x77, 0xa1, 0x26, 0xa1, 0xe6,
	0xe5, 0xfb, 0x88, 0x94, 0x1e, 0x74, 0x17, 0x2c, 0xc1, 0xb6, 0xc2, 0x0d, 0x8a, 0xbf, 0xc8, 0xb6,
	0x82, 0x28, 0x0a, 0xfd, 0x0f, 0x26, 0xdd, 0xb8, 0x5c, 0x39, 0x2a, 0x5e, 0x7f, 0x43, 0x4c, 0xba,
	0x91, 0xf1, 0x0b, 0x1e, 0x32, 0xf7, 0x6b, 0x1e, 0x2f, 0xd7, 0x11, 0x51, 0x14, 0x7e, 0x07, 0xb6,
	0x92, 0x87, 0xea, 0x60, 0x8f, 0xcf, 0xde, 0x8c, 0x4e, 0x9d, 0x7f, 0x50, 0x13, 0x6a, 0x83, 0xb3,
	0xd3, 0x69, 0x7f, 0x30, 0x9d, 0x38, 0x86, 0x44, 0xef, 0xc9, 0x70, 0x32, 0x3c, 0x1d, 0x0c, 0x1d,
	0x13, 0xd5, 0xc0, 0x9a, 0x0e, 0x3f, 0x4e, 0x9d, 0x8a, 0xb4, 0x4e, 0x46, 0xe3, 0xa1, 0x63, 0xa1,
	0x2a, 0x98, 0xfd, 0x73, 0xc7, 0x3e, 0xae, 0x7e, 0xb6, 0x22, 0xca, 0xe3, 0x59, 0x55, 0xed, 0xdb,
	0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x78, 0x88, 0x4d, 0x29, 0x7d, 0x05, 0x00, 0x00,
}
