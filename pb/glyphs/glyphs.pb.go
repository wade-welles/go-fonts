// Code generated by protoc-gen-go. DO NOT EDIT.
// source: glyphs.proto

package glyphs

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

type Glyphs struct {
	Glyphs               []*Glyph `protobuf:"bytes,1,rep,name=glyphs,proto3" json:"glyphs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Glyphs) Reset()         { *m = Glyphs{} }
func (m *Glyphs) String() string { return proto.CompactTextString(m) }
func (*Glyphs) ProtoMessage()    {}
func (*Glyphs) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bbe9e0d5eab4d4a, []int{0}
}

func (m *Glyphs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Glyphs.Unmarshal(m, b)
}
func (m *Glyphs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Glyphs.Marshal(b, m, deterministic)
}
func (m *Glyphs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Glyphs.Merge(m, src)
}
func (m *Glyphs) XXX_Size() int {
	return xxx_messageInfo_Glyphs.Size(m)
}
func (m *Glyphs) XXX_DiscardUnknown() {
	xxx_messageInfo_Glyphs.DiscardUnknown(m)
}

var xxx_messageInfo_Glyphs proto.InternalMessageInfo

func (m *Glyphs) GetGlyphs() []*Glyph {
	if m != nil {
		return m.Glyphs
	}
	return nil
}

type Glyph struct {
	HorizAdvX            float64     `protobuf:"fixed64,1,opt,name=horiz_adv_x,json=horizAdvX,proto3" json:"horiz_adv_x,omitempty"`
	Unicode              string      `protobuf:"bytes,2,opt,name=unicode,proto3" json:"unicode,omitempty"`
	GerberLP             string      `protobuf:"bytes,3,opt,name=gerber_l_p,json=gerberLP,proto3" json:"gerber_l_p,omitempty"`
	PathSteps            []*PathStep `protobuf:"bytes,4,rep,name=path_steps,json=pathSteps,proto3" json:"path_steps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Glyph) Reset()         { *m = Glyph{} }
func (m *Glyph) String() string { return proto.CompactTextString(m) }
func (*Glyph) ProtoMessage()    {}
func (*Glyph) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bbe9e0d5eab4d4a, []int{1}
}

func (m *Glyph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Glyph.Unmarshal(m, b)
}
func (m *Glyph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Glyph.Marshal(b, m, deterministic)
}
func (m *Glyph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Glyph.Merge(m, src)
}
func (m *Glyph) XXX_Size() int {
	return xxx_messageInfo_Glyph.Size(m)
}
func (m *Glyph) XXX_DiscardUnknown() {
	xxx_messageInfo_Glyph.DiscardUnknown(m)
}

var xxx_messageInfo_Glyph proto.InternalMessageInfo

func (m *Glyph) GetHorizAdvX() float64 {
	if m != nil {
		return m.HorizAdvX
	}
	return 0
}

func (m *Glyph) GetUnicode() string {
	if m != nil {
		return m.Unicode
	}
	return ""
}

func (m *Glyph) GetGerberLP() string {
	if m != nil {
		return m.GerberLP
	}
	return ""
}

func (m *Glyph) GetPathSteps() []*PathStep {
	if m != nil {
		return m.PathSteps
	}
	return nil
}

type PathStep struct {
	C                    uint32    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	P                    []float64 `protobuf:"fixed64,2,rep,packed,name=p,proto3" json:"p,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PathStep) Reset()         { *m = PathStep{} }
func (m *PathStep) String() string { return proto.CompactTextString(m) }
func (*PathStep) ProtoMessage()    {}
func (*PathStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bbe9e0d5eab4d4a, []int{2}
}

func (m *PathStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathStep.Unmarshal(m, b)
}
func (m *PathStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathStep.Marshal(b, m, deterministic)
}
func (m *PathStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathStep.Merge(m, src)
}
func (m *PathStep) XXX_Size() int {
	return xxx_messageInfo_PathStep.Size(m)
}
func (m *PathStep) XXX_DiscardUnknown() {
	xxx_messageInfo_PathStep.DiscardUnknown(m)
}

var xxx_messageInfo_PathStep proto.InternalMessageInfo

func (m *PathStep) GetC() uint32 {
	if m != nil {
		return m.C
	}
	return 0
}

func (m *PathStep) GetP() []float64 {
	if m != nil {
		return m.P
	}
	return nil
}

func init() {
	proto.RegisterType((*Glyphs)(nil), "glyphs.Glyphs")
	proto.RegisterType((*Glyph)(nil), "glyphs.Glyph")
	proto.RegisterType((*PathStep)(nil), "glyphs.PathStep")
}

func init() { proto.RegisterFile("glyphs.proto", fileDescriptor_6bbe9e0d5eab4d4a) }

var fileDescriptor_6bbe9e0d5eab4d4a = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x99, 0xae, 0xae, 0xdd, 0x69, 0x0b, 0x92, 0x53, 0x0e, 0x22, 0xa1, 0xa0, 0xe4, 0xd4,
	0x82, 0x3e, 0x81, 0x27, 0x2f, 0x1e, 0x4a, 0xbc, 0x78, 0x0b, 0xdb, 0xdd, 0xd0, 0x14, 0x16, 0x33,
	0x24, 0x71, 0x51, 0x9f, 0xc2, 0x47, 0x96, 0xce, 0x6e, 0x6e, 0xf3, 0x7f, 0x3f, 0x3f, 0x7c, 0x83,
	0xeb, 0xd3, 0xf0, 0x43, 0x3e, 0xed, 0x28, 0x86, 0x1c, 0x44, 0x3d, 0xa5, 0xed, 0x1e, 0xeb, 0x57,
	0xbe, 0xc4, 0x03, 0xce, 0x4c, 0x82, 0xaa, 0xf4, 0xea, 0x69, 0xb3, 0x9b, 0x07, 0xdc, 0x9b, 0x32,
	0xf8, 0x03, 0xbc, 0x66, 0x22, 0xee, 0x71, 0xe5, 0x43, 0x3c, 0xff, 0xda, 0xb6, 0x1f, 0xed, 0xb7,
	0x04, 0x05, 0x1a, 0x4c, 0xc3, 0xe8, 0xa5, 0x1f, 0x3f, 0x84, 0xc4, 0x9b, 0xaf, 0xcf, 0x73, 0x17,
	0x7a, 0x27, 0x17, 0x0a, 0x74, 0x63, 0x4a, 0x14, 0x77, 0x88, 0x27, 0x17, 0x8f, 0x2e, 0xda, 0xc1,
	0x92, 0xac, 0xb8, 0x5c, 0x4e, 0xe4, 0xed, 0x20, 0xf6, 0x88, 0xd4, 0x66, 0x6f, 0x53, 0x76, 0x94,
	0xe4, 0x15, 0xcb, 0xdc, 0x16, 0x99, 0x43, 0x9b, 0xfd, 0x7b, 0x76, 0x64, 0x1a, 0x9a, 0xaf, 0xb4,
	0x7d, 0xc4, 0x65, 0xc1, 0x62, 0x8d, 0xd0, 0xb1, 0xca, 0xc6, 0x40, 0x77, 0x49, 0x24, 0x17, 0xaa,
	0xd2, 0x60, 0x80, 0x8e, 0x35, 0xbf, 0xfe, 0xfc, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x5b, 0x0f,
	0x6a, 0x0a, 0x01, 0x00, 0x00,
}