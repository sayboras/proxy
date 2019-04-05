// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/mysql_proxy/v1alpha1/mysql_proxy.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
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

// [#protodoc-title: MySQL proxy]
// MySQL Proxy :ref:`configuration overview <config_network_filters_mysql_proxy>`.
type MySQLProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mysql_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// [#not-implemented-hide:] The optional path to use for writing MySQL access logs.
	// If the access log field is empty, access logs will not be written.
	AccessLog            string   `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLProxy) Reset()         { *m = MySQLProxy{} }
func (m *MySQLProxy) String() string { return proto.CompactTextString(m) }
func (*MySQLProxy) ProtoMessage()    {}
func (*MySQLProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4bac5cccef760ed, []int{0}
}

func (m *MySQLProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySQLProxy.Unmarshal(m, b)
}
func (m *MySQLProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySQLProxy.Marshal(b, m, deterministic)
}
func (m *MySQLProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLProxy.Merge(m, src)
}
func (m *MySQLProxy) XXX_Size() int {
	return xxx_messageInfo_MySQLProxy.Size(m)
}
func (m *MySQLProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLProxy.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLProxy proto.InternalMessageInfo

func (m *MySQLProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MySQLProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func init() {
	proto.RegisterType((*MySQLProxy)(nil), "envoy.config.filter.network.mysql_proxy.v1alpha1.MySQLProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/mysql_proxy/v1alpha1/mysql_proxy.proto", fileDescriptor_c4bac5cccef760ed)
}

var fileDescriptor_c4bac5cccef760ed = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0xcf, 0xad, 0x2c, 0x2e, 0xcc, 0x89, 0x2f, 0x28,
	0xca, 0xaf, 0xa8, 0xd4, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x44, 0x16, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x00, 0x9b, 0xa1, 0x07, 0x31, 0x43, 0x0f, 0x62, 0x86, 0x1e,
	0xd4, 0x0c, 0x3d, 0x64, 0xe5, 0x30, 0x33, 0xa4, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b,
	0x52, 0xf5, 0x61, 0x0c, 0x88, 0x51, 0x4a, 0xe1, 0x5c, 0x5c, 0xbe, 0x95, 0xc1, 0x81, 0x3e, 0x01,
	0x20, 0xf5, 0x42, 0x5a, 0x5c, 0xdc, 0xc5, 0x25, 0x89, 0x25, 0xf1, 0x05, 0x45, 0xa9, 0x69, 0x99,
	0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x9c, 0xbb, 0x5e, 0x1e, 0x60, 0x66, 0x29, 0x62,
	0x52, 0x60, 0x0c, 0xe2, 0x02, 0xc9, 0x06, 0x80, 0x25, 0x85, 0x64, 0xb9, 0xb8, 0x12, 0x93, 0x93,
	0x53, 0x8b, 0x8b, 0xe3, 0x73, 0xf2, 0xd3, 0x25, 0x98, 0x40, 0x4a, 0x83, 0x38, 0x21, 0x22, 0x3e,
	0xf9, 0xe9, 0x4e, 0xb1, 0x5c, 0x76, 0x99, 0xf9, 0x7a, 0x60, 0x87, 0x42, 0xdc, 0x42, 0xaa, 0x9b,
	0x9d, 0xf8, 0x7d, 0x41, 0xa2, 0x60, 0x87, 0x05, 0x80, 0xdc, 0x1a, 0xc0, 0x18, 0xc5, 0x01, 0x93,
	0x4c, 0x62, 0x03, 0x3b, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x32, 0xb8, 0xb4, 0x4f,
	0x01, 0x00, 0x00,
}