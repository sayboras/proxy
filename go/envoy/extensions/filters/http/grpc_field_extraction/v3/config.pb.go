// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.24.4
// source: envoy/extensions/filters/http/grpc_field_extraction/v3/config.proto

package grpc_field_extractionv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type GrpcFieldExtractionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The proto descriptor set binary for the gRPC services.
	//
	// It could be passed by a local file through “Datasource.filename“ or embedded in the
	// “Datasource.inline_bytes“.
	DescriptorSet *v3.DataSource `protobuf:"bytes,1,opt,name=descriptor_set,json=descriptorSet,proto3" json:"descriptor_set,omitempty"`
	// Specify the extraction info.
	// The key is the fully qualified gRPC method name.
	// “${package}.${Service}.${Method}“, like
	// “endpoints.examples.bookstore.BookStore.GetShelf“
	//
	// The value is the field extractions for individual gRPC method.
	ExtractionsByMethod map[string]*FieldExtractions `protobuf:"bytes,2,rep,name=extractions_by_method,json=extractionsByMethod,proto3" json:"extractions_by_method,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GrpcFieldExtractionConfig) Reset() {
	*x = GrpcFieldExtractionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcFieldExtractionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcFieldExtractionConfig) ProtoMessage() {}

func (x *GrpcFieldExtractionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcFieldExtractionConfig.ProtoReflect.Descriptor instead.
func (*GrpcFieldExtractionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcFieldExtractionConfig) GetDescriptorSet() *v3.DataSource {
	if x != nil {
		return x.DescriptorSet
	}
	return nil
}

func (x *GrpcFieldExtractionConfig) GetExtractionsByMethod() map[string]*FieldExtractions {
	if x != nil {
		return x.ExtractionsByMethod
	}
	return nil
}

// This message can be used to support per route config approach later even
// though the Istio doesn't support that so far.
type FieldExtractions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field extractions for requests.
	// The key is the field path within the grpc request.
	// For example, we can define “foo.bar.name“ if we want to extract
	// “Request.foo.bar.name“.
	//
	// .. code-block:: proto
	//
	//	message Request {
	//	  Foo foo = 1;
	//	}
	//
	//	message Foo {
	//	  Bar bar = 1;
	//	}
	//
	//	message Bar {
	//	  string name = 1;
	//	}
	RequestFieldExtractions map[string]*RequestFieldValueDisposition `protobuf:"bytes,1,rep,name=request_field_extractions,json=requestFieldExtractions,proto3" json:"request_field_extractions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FieldExtractions) Reset() {
	*x = FieldExtractions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldExtractions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldExtractions) ProtoMessage() {}

func (x *FieldExtractions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldExtractions.ProtoReflect.Descriptor instead.
func (*FieldExtractions) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescGZIP(), []int{1}
}

func (x *FieldExtractions) GetRequestFieldExtractions() map[string]*RequestFieldValueDisposition {
	if x != nil {
		return x.RequestFieldExtractions
	}
	return nil
}

type RequestFieldValueDisposition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Disposition:
	//
	//	*RequestFieldValueDisposition_DynamicMetadata
	Disposition isRequestFieldValueDisposition_Disposition `protobuf_oneof:"disposition"`
}

func (x *RequestFieldValueDisposition) Reset() {
	*x = RequestFieldValueDisposition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFieldValueDisposition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFieldValueDisposition) ProtoMessage() {}

func (x *RequestFieldValueDisposition) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFieldValueDisposition.ProtoReflect.Descriptor instead.
func (*RequestFieldValueDisposition) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescGZIP(), []int{2}
}

func (m *RequestFieldValueDisposition) GetDisposition() isRequestFieldValueDisposition_Disposition {
	if m != nil {
		return m.Disposition
	}
	return nil
}

func (x *RequestFieldValueDisposition) GetDynamicMetadata() string {
	if x, ok := x.GetDisposition().(*RequestFieldValueDisposition_DynamicMetadata); ok {
		return x.DynamicMetadata
	}
	return ""
}

type isRequestFieldValueDisposition_Disposition interface {
	isRequestFieldValueDisposition_Disposition()
}

type RequestFieldValueDisposition_DynamicMetadata struct {
	// The dynamic metadata namespace. If empty, "envoy.filters.http.grpc_field_extraction" will be used by default.
	//
	// Unimplemented. Uses "envoy.filters.http.grpc_field_extraction" for now.
	DynamicMetadata string `protobuf:"bytes,1,opt,name=dynamic_metadata,json=dynamicMetadata,proto3,oneof"`
}

func (*RequestFieldValueDisposition_DynamicMetadata) isRequestFieldValueDisposition_Disposition() {}

var File_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDesc = []byte{
	0x0a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x03, 0x0a, 0x19, 0x47, 0x72, 0x70, 0x63, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x51, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x15, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x6a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x13, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x90, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd9, 0x02, 0x0a, 0x10,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0xa1, 0x01, 0x0a, 0x19, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x65, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x17, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xa0, 0x01, 0x0a, 0x1c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x6a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5a, 0x0a, 0x1c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x69, 0x73, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x10, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0xd4, 0x01, 0x0a, 0x44, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x75, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescData = file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDesc
)

func file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDescData
}

var file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_goTypes = []interface{}{
	(*GrpcFieldExtractionConfig)(nil),    // 0: envoy.extensions.filters.http.grpc_field_extraction.v3.GrpcFieldExtractionConfig
	(*FieldExtractions)(nil),             // 1: envoy.extensions.filters.http.grpc_field_extraction.v3.FieldExtractions
	(*RequestFieldValueDisposition)(nil), // 2: envoy.extensions.filters.http.grpc_field_extraction.v3.RequestFieldValueDisposition
	nil,                                  // 3: envoy.extensions.filters.http.grpc_field_extraction.v3.GrpcFieldExtractionConfig.ExtractionsByMethodEntry
	nil,                                  // 4: envoy.extensions.filters.http.grpc_field_extraction.v3.FieldExtractions.RequestFieldExtractionsEntry
	(*v3.DataSource)(nil),                // 5: envoy.config.core.v3.DataSource
}
var file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_depIdxs = []int32{
	5, // 0: envoy.extensions.filters.http.grpc_field_extraction.v3.GrpcFieldExtractionConfig.descriptor_set:type_name -> envoy.config.core.v3.DataSource
	3, // 1: envoy.extensions.filters.http.grpc_field_extraction.v3.GrpcFieldExtractionConfig.extractions_by_method:type_name -> envoy.extensions.filters.http.grpc_field_extraction.v3.GrpcFieldExtractionConfig.ExtractionsByMethodEntry
	4, // 2: envoy.extensions.filters.http.grpc_field_extraction.v3.FieldExtractions.request_field_extractions:type_name -> envoy.extensions.filters.http.grpc_field_extraction.v3.FieldExtractions.RequestFieldExtractionsEntry
	1, // 3: envoy.extensions.filters.http.grpc_field_extraction.v3.GrpcFieldExtractionConfig.ExtractionsByMethodEntry.value:type_name -> envoy.extensions.filters.http.grpc_field_extraction.v3.FieldExtractions
	2, // 4: envoy.extensions.filters.http.grpc_field_extraction.v3.FieldExtractions.RequestFieldExtractionsEntry.value:type_name -> envoy.extensions.filters.http.grpc_field_extraction.v3.RequestFieldValueDisposition
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_init() }
func file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_init() {
	if File_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcFieldExtractionConfig); i {
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
		file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldExtractions); i {
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
		file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFieldValueDisposition); i {
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
	file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*RequestFieldValueDisposition_DynamicMetadata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto = out.File
	file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_rawDesc = nil
	file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_goTypes = nil
	file_envoy_extensions_filters_http_grpc_field_extraction_v3_config_proto_depIdxs = nil
}
