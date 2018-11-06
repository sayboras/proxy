// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/rbac/v2/rbac.proto

package v2

import (
	fmt "fmt"
	v2alpha "github.com/cilium/proxy/go/envoy/config/rbac/v2alpha"
	_ "github.com/gogo/protobuf/gogoproto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RBAC filter config.
type RBAC struct {
	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	Rules *v2alpha.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	// Shadow rules are not enforced by the filter (i.e., returning a 403)
	// but will emit stats and logs and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	ShadowRules          *v2alpha.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_15d628c6558085a7, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v2alpha.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v2alpha.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

type RBACPerRoute struct {
	// Override the global configuration of the filter with this new config.
	// If absent, the global RBAC policy will be disabled for this route.
	Rbac                 *RBAC    `protobuf:"bytes,2,opt,name=rbac,proto3" json:"rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBACPerRoute) Reset()         { *m = RBACPerRoute{} }
func (m *RBACPerRoute) String() string { return proto.CompactTextString(m) }
func (*RBACPerRoute) ProtoMessage()    {}
func (*RBACPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_15d628c6558085a7, []int{1}
}

func (m *RBACPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBACPerRoute.Unmarshal(m, b)
}
func (m *RBACPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBACPerRoute.Marshal(b, m, deterministic)
}
func (m *RBACPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBACPerRoute.Merge(m, src)
}
func (m *RBACPerRoute) XXX_Size() int {
	return xxx_messageInfo_RBACPerRoute.Size(m)
}
func (m *RBACPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RBACPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RBACPerRoute proto.InternalMessageInfo

func (m *RBACPerRoute) GetRbac() *RBAC {
	if m != nil {
		return m.Rbac
	}
	return nil
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.http.rbac.v2.RBAC")
	proto.RegisterType((*RBACPerRoute)(nil), "envoy.config.filter.http.rbac.v2.RBACPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/rbac/v2/rbac.proto", fileDescriptor_15d628c6558085a7)
}

var fileDescriptor_15d628c6558085a7 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0xb9, 0xe3, 0x94, 0x92, 0x76, 0x90, 0x20, 0x28, 0x5d, 0x2c, 0x45, 0x44, 0x10, 0x5e,
	0x20, 0xe2, 0xe2, 0xe6, 0xb9, 0x39, 0x49, 0xc6, 0x2e, 0x92, 0x6b, 0xd2, 0xbb, 0x40, 0xe8, 0x3b,
	0x72, 0x69, 0xc4, 0xd1, 0xff, 0x5c, 0xf2, 0x43, 0xb0, 0x93, 0x9d, 0xee, 0xcb, 0xbb, 0xcf, 0xe7,
	0x9b, 0x97, 0x90, 0x07, 0xbd, 0x0f, 0xf8, 0xc5, 0xb6, 0xb8, 0xdf, 0x99, 0x9e, 0xed, 0x8c, 0xf5,
	0xda, 0xb1, 0xc1, 0xfb, 0x91, 0xb9, 0x4e, 0x6e, 0x59, 0xe0, 0xe9, 0x0b, 0xa3, 0x43, 0x8f, 0x74,
	0x95, 0x60, 0xc8, 0x30, 0x64, 0x18, 0x22, 0x0c, 0x09, 0x0a, 0x7c, 0x79, 0x7b, 0x54, 0x57, 0x2a,
	0xa4, 0x1d, 0x07, 0xf9, 0xa7, 0x67, 0x79, 0x15, 0xa4, 0x35, 0x4a, 0x7a, 0xcd, 0x7e, 0x43, 0xf9,
	0x71, 0xd9, 0x63, 0x8f, 0x29, 0xb2, 0x98, 0xf2, 0x74, 0xfd, 0x5d, 0x91, 0x46, 0xb4, 0x2f, 0xaf,
	0xf4, 0x89, 0x9c, 0xb9, 0x83, 0xd5, 0xd3, 0x75, 0xb5, 0xaa, 0xee, 0xe7, 0xfc, 0x06, 0x8e, 0xf6,
	0x29, 0x3b, 0xa4, 0xd3, 0x20, 0xf2, 0x22, 0xd3, 0xb4, 0x25, 0x8b, 0x69, 0x90, 0x0a, 0x3f, 0x3f,
	0xb2, 0x5d, 0x9f, 0x66, 0xcf, 0xb3, 0x24, 0xa2, 0xb3, 0xde, 0x90, 0x45, 0x1c, 0xbe, 0x6b, 0x27,
	0xf0, 0xe0, 0x35, 0x7d, 0x26, 0x4d, 0x34, 0x4a, 0xd7, 0x1d, 0xfc, 0xf7, 0x32, 0xb9, 0x32, 0x39,
	0x6f, 0xcd, 0xac, 0xba, 0xa8, 0xc5, 0x4c, 0x99, 0x49, 0x76, 0x56, 0xab, 0xb6, 0xd9, 0xd4, 0x81,
	0x77, 0xe7, 0xe9, 0xb2, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0xa3, 0xe0, 0x0b, 0x92,
	0x01, 0x00, 0x00,
}