// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.29.2
// source: envoy/extensions/filters/http/header_mutation/v3/header_mutation.proto

package header_mutationv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/common/mutation_rules/v3"
	v31 "github.com/cilium/proxy/go/envoy/config/core/v3"
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

type Mutations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request mutations are applied before the request is forwarded to the upstream cluster.
	RequestMutations []*v3.HeaderMutation `protobuf:"bytes,1,rep,name=request_mutations,json=requestMutations,proto3" json:"request_mutations,omitempty"`
	// The “path“ header query parameter mutations are applied after “request_mutations“ and before the request
	// is forwarded to the next filter in the filter chain.
	QueryParameterMutations []*v31.KeyValueMutation `protobuf:"bytes,3,rep,name=query_parameter_mutations,json=queryParameterMutations,proto3" json:"query_parameter_mutations,omitempty"`
	// The response mutations are applied before the response is sent to the downstream client.
	ResponseMutations []*v3.HeaderMutation `protobuf:"bytes,2,rep,name=response_mutations,json=responseMutations,proto3" json:"response_mutations,omitempty"`
}

func (x *Mutations) Reset() {
	*x = Mutations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mutations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mutations) ProtoMessage() {}

func (x *Mutations) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mutations.ProtoReflect.Descriptor instead.
func (*Mutations) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescGZIP(), []int{0}
}

func (x *Mutations) GetRequestMutations() []*v3.HeaderMutation {
	if x != nil {
		return x.RequestMutations
	}
	return nil
}

func (x *Mutations) GetQueryParameterMutations() []*v31.KeyValueMutation {
	if x != nil {
		return x.QueryParameterMutations
	}
	return nil
}

func (x *Mutations) GetResponseMutations() []*v3.HeaderMutation {
	if x != nil {
		return x.ResponseMutations
	}
	return nil
}

// Per route configuration for the header mutation filter.
type HeaderMutationPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mutations *Mutations `protobuf:"bytes,1,opt,name=mutations,proto3" json:"mutations,omitempty"`
}

func (x *HeaderMutationPerRoute) Reset() {
	*x = HeaderMutationPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderMutationPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderMutationPerRoute) ProtoMessage() {}

func (x *HeaderMutationPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderMutationPerRoute.ProtoReflect.Descriptor instead.
func (*HeaderMutationPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescGZIP(), []int{1}
}

func (x *HeaderMutationPerRoute) GetMutations() *Mutations {
	if x != nil {
		return x.Mutations
	}
	return nil
}

// Configuration for the header mutation filter. The mutation rules in the filter configuration will
// always be applied first and then the per-route mutation rules, if both are specified.
type HeaderMutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mutations *Mutations `protobuf:"bytes,1,opt,name=mutations,proto3" json:"mutations,omitempty"`
	// If per route HeaderMutationPerRoute config is configured at multiple route levels, header mutations
	// at all specified levels are evaluated. By default, the order is from most specific (i.e. route entry level)
	// to least specific (i.e. route configuration level). Later header mutations may override earlier mutations.
	//
	// This order can be reversed by setting this field to true. In other words, most specific level mutation
	// is evaluated last.
	MostSpecificHeaderMutationsWins bool `protobuf:"varint,2,opt,name=most_specific_header_mutations_wins,json=mostSpecificHeaderMutationsWins,proto3" json:"most_specific_header_mutations_wins,omitempty"`
}

func (x *HeaderMutation) Reset() {
	*x = HeaderMutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderMutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderMutation) ProtoMessage() {}

func (x *HeaderMutation) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderMutation.ProtoReflect.Descriptor instead.
func (*HeaderMutation) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescGZIP(), []int{2}
}

func (x *HeaderMutation) GetMutations() *Mutations {
	if x != nil {
		return x.Mutations
	}
	return nil
}

func (x *HeaderMutation) GetMostSpecificHeaderMutationsWins() bool {
	if x != nil {
		return x.MostSpecificHeaderMutationsWins
	}
	return false
}

var File_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDesc = []byte{
	0x0a, 0x46, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x33, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x3a, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76,
	0x33, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x09, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x62, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x62, 0x0a, 0x19, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x17, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x64, 0x0a, 0x12,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x73, 0x0a, 0x16, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x59, 0x0a, 0x09,
	0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb9, 0x01, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x09, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x6d, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x23, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1f, 0x6d, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57,
	0x69, 0x6e, 0x73, 0x42, 0xca, 0x01, 0x0a, 0x3e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescData = file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDesc
)

func file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDescData
}

var file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_goTypes = []interface{}{
	(*Mutations)(nil),              // 0: envoy.extensions.filters.http.header_mutation.v3.Mutations
	(*HeaderMutationPerRoute)(nil), // 1: envoy.extensions.filters.http.header_mutation.v3.HeaderMutationPerRoute
	(*HeaderMutation)(nil),         // 2: envoy.extensions.filters.http.header_mutation.v3.HeaderMutation
	(*v3.HeaderMutation)(nil),      // 3: envoy.config.common.mutation_rules.v3.HeaderMutation
	(*v31.KeyValueMutation)(nil),   // 4: envoy.config.core.v3.KeyValueMutation
}
var file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.filters.http.header_mutation.v3.Mutations.request_mutations:type_name -> envoy.config.common.mutation_rules.v3.HeaderMutation
	4, // 1: envoy.extensions.filters.http.header_mutation.v3.Mutations.query_parameter_mutations:type_name -> envoy.config.core.v3.KeyValueMutation
	3, // 2: envoy.extensions.filters.http.header_mutation.v3.Mutations.response_mutations:type_name -> envoy.config.common.mutation_rules.v3.HeaderMutation
	0, // 3: envoy.extensions.filters.http.header_mutation.v3.HeaderMutationPerRoute.mutations:type_name -> envoy.extensions.filters.http.header_mutation.v3.Mutations
	0, // 4: envoy.extensions.filters.http.header_mutation.v3.HeaderMutation.mutations:type_name -> envoy.extensions.filters.http.header_mutation.v3.Mutations
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_init() }
func file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_init() {
	if File_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mutations); i {
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
		file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderMutationPerRoute); i {
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
		file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderMutation); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto = out.File
	file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_rawDesc = nil
	file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_goTypes = nil
	file_envoy_extensions_filters_http_header_mutation_v3_header_mutation_proto_depIdxs = nil
}
