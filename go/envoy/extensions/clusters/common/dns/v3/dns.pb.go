// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.29.2
// source: envoy/extensions/clusters/common/dns/v3/dns.proto

package dnsv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DnsLookupFamily int32

const (
	DnsLookupFamily_UNSPECIFIED  DnsLookupFamily = 0
	DnsLookupFamily_AUTO         DnsLookupFamily = 1
	DnsLookupFamily_V4_ONLY      DnsLookupFamily = 2
	DnsLookupFamily_V6_ONLY      DnsLookupFamily = 3
	DnsLookupFamily_V4_PREFERRED DnsLookupFamily = 4
	DnsLookupFamily_ALL          DnsLookupFamily = 5
)

// Enum value maps for DnsLookupFamily.
var (
	DnsLookupFamily_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "AUTO",
		2: "V4_ONLY",
		3: "V6_ONLY",
		4: "V4_PREFERRED",
		5: "ALL",
	}
	DnsLookupFamily_value = map[string]int32{
		"UNSPECIFIED":  0,
		"AUTO":         1,
		"V4_ONLY":      2,
		"V6_ONLY":      3,
		"V4_PREFERRED": 4,
		"ALL":          5,
	}
)

func (x DnsLookupFamily) Enum() *DnsLookupFamily {
	p := new(DnsLookupFamily)
	*p = x
	return p
}

func (x DnsLookupFamily) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DnsLookupFamily) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_clusters_common_dns_v3_dns_proto_enumTypes[0].Descriptor()
}

func (DnsLookupFamily) Type() protoreflect.EnumType {
	return &file_envoy_extensions_clusters_common_dns_v3_dns_proto_enumTypes[0]
}

func (x DnsLookupFamily) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DnsLookupFamily.Descriptor instead.
func (DnsLookupFamily) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDescGZIP(), []int{0}
}

var File_envoy_extensions_clusters_common_dns_v3_dns_proto protoreflect.FileDescriptor

var file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDesc = []byte{
	0x0a, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x61, 0x0a, 0x0f, 0x44,
	0x6e, 0x73, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x34, 0x5f,
	0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x36, 0x5f, 0x4f, 0x4e, 0x4c,
	0x59, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x34, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x45, 0x52,
	0x52, 0x45, 0x44, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x05, 0x42, 0xa1,
	0x01, 0x0a, 0x35, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x08, 0x44, 0x6e, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x6e,
	0x73, 0x2f, 0x76, 0x33, 0x3b, 0x64, 0x6e, 0x73, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDescOnce sync.Once
	file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDescData = file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDesc
)

func file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDescGZIP() []byte {
	file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDescData)
	})
	return file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDescData
}

var file_envoy_extensions_clusters_common_dns_v3_dns_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_clusters_common_dns_v3_dns_proto_goTypes = []interface{}{
	(DnsLookupFamily)(0), // 0: envoy.extensions.clusters.common.dns.v3.DnsLookupFamily
}
var file_envoy_extensions_clusters_common_dns_v3_dns_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_clusters_common_dns_v3_dns_proto_init() }
func file_envoy_extensions_clusters_common_dns_v3_dns_proto_init() {
	if File_envoy_extensions_clusters_common_dns_v3_dns_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_clusters_common_dns_v3_dns_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_clusters_common_dns_v3_dns_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_clusters_common_dns_v3_dns_proto_enumTypes,
	}.Build()
	File_envoy_extensions_clusters_common_dns_v3_dns_proto = out.File
	file_envoy_extensions_clusters_common_dns_v3_dns_proto_rawDesc = nil
	file_envoy_extensions_clusters_common_dns_v3_dns_proto_goTypes = nil
	file_envoy_extensions_clusters_common_dns_v3_dns_proto_depIdxs = nil
}
