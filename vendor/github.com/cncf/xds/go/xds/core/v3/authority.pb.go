// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.5
// source: xds/core/v3/authority.proto

package v3

import (
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Authority struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Authority) Reset() {
	*x = Authority{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_core_v3_authority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authority) ProtoMessage() {}

func (x *Authority) ProtoReflect() protoreflect.Message {
	mi := &file_xds_core_v3_authority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authority.ProtoReflect.Descriptor instead.
func (*Authority) Descriptor() ([]byte, []int) {
	return file_xds_core_v3_authority_proto_rawDescGZIP(), []int{0}
}

func (x *Authority) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_xds_core_v3_authority_proto protoreflect.FileDescriptor

var file_xds_core_v3_authority_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x78, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x78,
	0x64, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x56,
	0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x42, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_core_v3_authority_proto_rawDescOnce sync.Once
	file_xds_core_v3_authority_proto_rawDescData = file_xds_core_v3_authority_proto_rawDesc
)

func file_xds_core_v3_authority_proto_rawDescGZIP() []byte {
	file_xds_core_v3_authority_proto_rawDescOnce.Do(func() {
		file_xds_core_v3_authority_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_core_v3_authority_proto_rawDescData)
	})
	return file_xds_core_v3_authority_proto_rawDescData
}

var file_xds_core_v3_authority_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_xds_core_v3_authority_proto_goTypes = []interface{}{
	(*Authority)(nil), // 0: xds.core.v3.Authority
}
var file_xds_core_v3_authority_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_xds_core_v3_authority_proto_init() }
func file_xds_core_v3_authority_proto_init() {
	if File_xds_core_v3_authority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xds_core_v3_authority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authority); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xds_core_v3_authority_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_core_v3_authority_proto_goTypes,
		DependencyIndexes: file_xds_core_v3_authority_proto_depIdxs,
		MessageInfos:      file_xds_core_v3_authority_proto_msgTypes,
	}.Build()
	File_xds_core_v3_authority_proto = out.File
	file_xds_core_v3_authority_proto_rawDesc = nil
	file_xds_core_v3_authority_proto_goTypes = nil
	file_xds_core_v3_authority_proto_depIdxs = nil
}
