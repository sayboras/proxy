// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.5
// source: udpa/annotations/status.proto

package annotations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PackageVersionStatus int32

const (
	PackageVersionStatus_UNKNOWN                      PackageVersionStatus = 0
	PackageVersionStatus_FROZEN                       PackageVersionStatus = 1
	PackageVersionStatus_ACTIVE                       PackageVersionStatus = 2
	PackageVersionStatus_NEXT_MAJOR_VERSION_CANDIDATE PackageVersionStatus = 3
)

// Enum value maps for PackageVersionStatus.
var (
	PackageVersionStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "FROZEN",
		2: "ACTIVE",
		3: "NEXT_MAJOR_VERSION_CANDIDATE",
	}
	PackageVersionStatus_value = map[string]int32{
		"UNKNOWN":                      0,
		"FROZEN":                       1,
		"ACTIVE":                       2,
		"NEXT_MAJOR_VERSION_CANDIDATE": 3,
	}
)

func (x PackageVersionStatus) Enum() *PackageVersionStatus {
	p := new(PackageVersionStatus)
	*p = x
	return p
}

func (x PackageVersionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackageVersionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_udpa_annotations_status_proto_enumTypes[0].Descriptor()
}

func (PackageVersionStatus) Type() protoreflect.EnumType {
	return &file_udpa_annotations_status_proto_enumTypes[0]
}

func (x PackageVersionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageVersionStatus.Descriptor instead.
func (PackageVersionStatus) EnumDescriptor() ([]byte, []int) {
	return file_udpa_annotations_status_proto_rawDescGZIP(), []int{0}
}

type StatusAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkInProgress       bool                 `protobuf:"varint,1,opt,name=work_in_progress,json=workInProgress,proto3" json:"work_in_progress,omitempty"`
	PackageVersionStatus PackageVersionStatus `protobuf:"varint,2,opt,name=package_version_status,json=packageVersionStatus,proto3,enum=udpa.annotations.PackageVersionStatus" json:"package_version_status,omitempty"`
}

func (x *StatusAnnotation) Reset() {
	*x = StatusAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_udpa_annotations_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusAnnotation) ProtoMessage() {}

func (x *StatusAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_udpa_annotations_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusAnnotation.ProtoReflect.Descriptor instead.
func (*StatusAnnotation) Descriptor() ([]byte, []int) {
	return file_udpa_annotations_status_proto_rawDescGZIP(), []int{0}
}

func (x *StatusAnnotation) GetWorkInProgress() bool {
	if x != nil {
		return x.WorkInProgress
	}
	return false
}

func (x *StatusAnnotation) GetPackageVersionStatus() PackageVersionStatus {
	if x != nil {
		return x.PackageVersionStatus
	}
	return PackageVersionStatus_UNKNOWN
}

var file_udpa_annotations_status_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*StatusAnnotation)(nil),
		Field:         222707719,
		Name:          "udpa.annotations.file_status",
		Tag:           "bytes,222707719,opt,name=file_status",
		Filename:      "udpa/annotations/status.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional udpa.annotations.StatusAnnotation file_status = 222707719;
	E_FileStatus = &file_udpa_annotations_status_proto_extTypes[0]
)

var File_udpa_annotations_status_proto protoreflect.FileDescriptor

var file_udpa_annotations_status_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x5c, 0x0a, 0x16, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x26, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x14, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x5d, 0x0a, 0x14, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x4f, 0x5a, 0x45, 0x4e, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x20, 0x0a,
	0x1c, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x4d, 0x41, 0x4a, 0x4f, 0x52, 0x5f, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x3a,
	0x64, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x87, 0x80, 0x99,
	0x6a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_udpa_annotations_status_proto_rawDescOnce sync.Once
	file_udpa_annotations_status_proto_rawDescData = file_udpa_annotations_status_proto_rawDesc
)

func file_udpa_annotations_status_proto_rawDescGZIP() []byte {
	file_udpa_annotations_status_proto_rawDescOnce.Do(func() {
		file_udpa_annotations_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_udpa_annotations_status_proto_rawDescData)
	})
	return file_udpa_annotations_status_proto_rawDescData
}

var file_udpa_annotations_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_udpa_annotations_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_udpa_annotations_status_proto_goTypes = []interface{}{
	(PackageVersionStatus)(0),        // 0: udpa.annotations.PackageVersionStatus
	(*StatusAnnotation)(nil),         // 1: udpa.annotations.StatusAnnotation
	(*descriptorpb.FileOptions)(nil), // 2: google.protobuf.FileOptions
}
var file_udpa_annotations_status_proto_depIdxs = []int32{
	0, // 0: udpa.annotations.StatusAnnotation.package_version_status:type_name -> udpa.annotations.PackageVersionStatus
	2, // 1: udpa.annotations.file_status:extendee -> google.protobuf.FileOptions
	1, // 2: udpa.annotations.file_status:type_name -> udpa.annotations.StatusAnnotation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_udpa_annotations_status_proto_init() }
func file_udpa_annotations_status_proto_init() {
	if File_udpa_annotations_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_udpa_annotations_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusAnnotation); i {
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
			RawDescriptor: file_udpa_annotations_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_udpa_annotations_status_proto_goTypes,
		DependencyIndexes: file_udpa_annotations_status_proto_depIdxs,
		EnumInfos:         file_udpa_annotations_status_proto_enumTypes,
		MessageInfos:      file_udpa_annotations_status_proto_msgTypes,
		ExtensionInfos:    file_udpa_annotations_status_proto_extTypes,
	}.Build()
	File_udpa_annotations_status_proto = out.File
	file_udpa_annotations_status_proto_rawDesc = nil
	file_udpa_annotations_status_proto_goTypes = nil
	file_udpa_annotations_status_proto_depIdxs = nil
}
