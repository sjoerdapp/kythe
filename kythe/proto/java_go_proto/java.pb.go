// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/java.proto

package java_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	storage_go_proto "kythe.io/kythe/proto/storage_go_proto"
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

type JarDetails struct {
	Jar                  []*JarDetails_Jar `protobuf:"bytes,1,rep,name=jar,proto3" json:"jar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JarDetails) Reset()         { *m = JarDetails{} }
func (m *JarDetails) String() string { return proto.CompactTextString(m) }
func (*JarDetails) ProtoMessage()    {}
func (*JarDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5838ccc9a1d029, []int{0}
}

func (m *JarDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JarDetails.Unmarshal(m, b)
}
func (m *JarDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JarDetails.Marshal(b, m, deterministic)
}
func (m *JarDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JarDetails.Merge(m, src)
}
func (m *JarDetails) XXX_Size() int {
	return xxx_messageInfo_JarDetails.Size(m)
}
func (m *JarDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_JarDetails.DiscardUnknown(m)
}

var xxx_messageInfo_JarDetails proto.InternalMessageInfo

func (m *JarDetails) GetJar() []*JarDetails_Jar {
	if m != nil {
		return m.Jar
	}
	return nil
}

type JarDetails_Jar struct {
	VName                *storage_go_proto.VName `protobuf:"bytes,1,opt,name=v_name,json=vName,proto3" json:"v_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *JarDetails_Jar) Reset()         { *m = JarDetails_Jar{} }
func (m *JarDetails_Jar) String() string { return proto.CompactTextString(m) }
func (*JarDetails_Jar) ProtoMessage()    {}
func (*JarDetails_Jar) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5838ccc9a1d029, []int{0, 0}
}

func (m *JarDetails_Jar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JarDetails_Jar.Unmarshal(m, b)
}
func (m *JarDetails_Jar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JarDetails_Jar.Marshal(b, m, deterministic)
}
func (m *JarDetails_Jar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JarDetails_Jar.Merge(m, src)
}
func (m *JarDetails_Jar) XXX_Size() int {
	return xxx_messageInfo_JarDetails_Jar.Size(m)
}
func (m *JarDetails_Jar) XXX_DiscardUnknown() {
	xxx_messageInfo_JarDetails_Jar.DiscardUnknown(m)
}

var xxx_messageInfo_JarDetails_Jar proto.InternalMessageInfo

func (m *JarDetails_Jar) GetVName() *storage_go_proto.VName {
	if m != nil {
		return m.VName
	}
	return nil
}

type JarEntryDetails struct {
	JarContainer         int32    `protobuf:"varint,1,opt,name=jar_container,json=jarContainer,proto3" json:"jar_container,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JarEntryDetails) Reset()         { *m = JarEntryDetails{} }
func (m *JarEntryDetails) String() string { return proto.CompactTextString(m) }
func (*JarEntryDetails) ProtoMessage()    {}
func (*JarEntryDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5838ccc9a1d029, []int{1}
}

func (m *JarEntryDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JarEntryDetails.Unmarshal(m, b)
}
func (m *JarEntryDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JarEntryDetails.Marshal(b, m, deterministic)
}
func (m *JarEntryDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JarEntryDetails.Merge(m, src)
}
func (m *JarEntryDetails) XXX_Size() int {
	return xxx_messageInfo_JarEntryDetails.Size(m)
}
func (m *JarEntryDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_JarEntryDetails.DiscardUnknown(m)
}

var xxx_messageInfo_JarEntryDetails proto.InternalMessageInfo

func (m *JarEntryDetails) GetJarContainer() int32 {
	if m != nil {
		return m.JarContainer
	}
	return 0
}

type JavaDetails struct {
	Classpath            []string `protobuf:"bytes,1,rep,name=classpath,proto3" json:"classpath,omitempty"`
	Sourcepath           []string `protobuf:"bytes,2,rep,name=sourcepath,proto3" json:"sourcepath,omitempty"`
	Bootclasspath        []string `protobuf:"bytes,3,rep,name=bootclasspath,proto3" json:"bootclasspath,omitempty"`
	ExtraJavacopts       []string `protobuf:"bytes,10,rep,name=extra_javacopts,json=extraJavacopts,proto3" json:"extra_javacopts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JavaDetails) Reset()         { *m = JavaDetails{} }
func (m *JavaDetails) String() string { return proto.CompactTextString(m) }
func (*JavaDetails) ProtoMessage()    {}
func (*JavaDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5838ccc9a1d029, []int{2}
}

func (m *JavaDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JavaDetails.Unmarshal(m, b)
}
func (m *JavaDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JavaDetails.Marshal(b, m, deterministic)
}
func (m *JavaDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JavaDetails.Merge(m, src)
}
func (m *JavaDetails) XXX_Size() int {
	return xxx_messageInfo_JavaDetails.Size(m)
}
func (m *JavaDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_JavaDetails.DiscardUnknown(m)
}

var xxx_messageInfo_JavaDetails proto.InternalMessageInfo

func (m *JavaDetails) GetClasspath() []string {
	if m != nil {
		return m.Classpath
	}
	return nil
}

func (m *JavaDetails) GetSourcepath() []string {
	if m != nil {
		return m.Sourcepath
	}
	return nil
}

func (m *JavaDetails) GetBootclasspath() []string {
	if m != nil {
		return m.Bootclasspath
	}
	return nil
}

func (m *JavaDetails) GetExtraJavacopts() []string {
	if m != nil {
		return m.ExtraJavacopts
	}
	return nil
}

func init() {
	proto.RegisterType((*JarDetails)(nil), "kythe.proto.JarDetails")
	proto.RegisterType((*JarDetails_Jar)(nil), "kythe.proto.JarDetails.Jar")
	proto.RegisterType((*JarEntryDetails)(nil), "kythe.proto.JarEntryDetails")
	proto.RegisterType((*JavaDetails)(nil), "kythe.proto.JavaDetails")
}

func init() { proto.RegisterFile("kythe/proto/java.proto", fileDescriptor_4e5838ccc9a1d029) }

var fileDescriptor_4e5838ccc9a1d029 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4f, 0xb3, 0x40,
	0x10, 0xc6, 0x43, 0x79, 0xdb, 0xbc, 0x1d, 0xac, 0x35, 0x7b, 0x30, 0xb5, 0x1a, 0x6d, 0xd0, 0xc4,
	0x7a, 0x90, 0x9a, 0x9a, 0xf4, 0x03, 0xf8, 0xe7, 0xb2, 0x07, 0x0f, 0x1c, 0x3c, 0x78, 0x21, 0x53,
	0xdc, 0xd0, 0x22, 0x30, 0x64, 0x76, 0x25, 0xf6, 0xab, 0xf8, 0x69, 0x0d, 0xab, 0x04, 0x7a, 0x9a,
	0x99, 0xdf, 0xf3, 0xcc, 0xb3, 0x9b, 0x81, 0xe3, 0x8f, 0x9d, 0xd9, 0xa8, 0x45, 0xc9, 0x64, 0x68,
	0x91, 0x62, 0x85, 0x81, 0x6d, 0x85, 0x67, 0xf9, 0xef, 0x30, 0x3d, 0xe9, 0x9a, 0xb4, 0x21, 0xc6,
	0xe4, 0x4f, 0xf2, 0x35, 0x80, 0x44, 0x7e, 0x52, 0x06, 0xb7, 0x99, 0x16, 0xb7, 0xe0, 0xa6, 0xc8,
	0x13, 0x67, 0xe6, 0xce, 0xbd, 0xe5, 0x69, 0xd0, 0xc9, 0x08, 0x5a, 0x57, 0xdd, 0x86, 0xb5, 0x6f,
	0xba, 0x02, 0x57, 0x22, 0x8b, 0x1b, 0x18, 0x54, 0x51, 0x81, 0xb9, 0x9a, 0x38, 0x33, 0x67, 0xee,
	0x2d, 0xc5, 0xde, 0xe2, 0xeb, 0x0b, 0xe6, 0x2a, 0xec, 0x57, 0x75, 0x91, 0xff, 0xfe, 0xf7, 0x8e,
	0x5c, 0x7f, 0x05, 0x63, 0x89, 0xfc, 0x5c, 0x18, 0xde, 0x35, 0x2f, 0x5f, 0xc2, 0x28, 0x45, 0x8e,
	0x62, 0x2a, 0x0c, 0x6e, 0x0b, 0xc5, 0x36, 0xaa, 0x1f, 0x1e, 0xa4, 0xc8, 0x8f, 0x0d, 0xf3, 0xbf,
	0x1d, 0xf0, 0x24, 0x56, 0xd8, 0x2c, 0x9d, 0xc1, 0x30, 0xce, 0x50, 0xeb, 0x12, 0xcd, 0xc6, 0x7e,
	0x7a, 0x18, 0xb6, 0x40, 0x9c, 0x03, 0x68, 0xfa, 0xe4, 0x58, 0x59, 0xb9, 0x67, 0xe5, 0x0e, 0x11,
	0x57, 0x30, 0x5a, 0x13, 0x99, 0x36, 0xc1, 0xb5, 0x96, 0x7d, 0x28, 0xae, 0x61, 0xac, 0xbe, 0x0c,
	0x63, 0x54, 0x1f, 0x37, 0xa6, 0xd2, 0xe8, 0x09, 0x58, 0xdf, 0xa1, 0xc5, 0xb2, 0xa1, 0x0f, 0x77,
	0x70, 0x11, 0x53, 0x1e, 0x24, 0x44, 0x49, 0xa6, 0x82, 0x77, 0x55, 0x19, 0xa2, 0x4c, 0x77, 0x4f,
	0xf1, 0x36, 0xaa, 0x33, 0xa2, 0x84, 0x22, 0x3b, 0xae, 0x07, 0xb6, 0xdc, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xe6, 0x14, 0x80, 0x5d, 0xc4, 0x01, 0x00, 0x00,
}
