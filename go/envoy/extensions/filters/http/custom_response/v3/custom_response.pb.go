// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.29.2
// source: envoy/extensions/filters/http/custom_response/v3/custom_response.proto

package custom_responsev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/cncf/xds/go/xds/type/matcher/v3"
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

// The filter configuration is a collection of custom response
// policies in a matcher tree. The configuration can be defined at the filter,
// virtual host or route level. The response will be matched against the most
// specific to the least specific config, till a match is found.
type CustomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matcher to match against the original response to select a
	// :ref:`Custom Response Policy <extension_category_envoy.http.custom_response>`
	// that will override the original response. The matching is done by matching
	// against :ref:`response header values<extension_category_envoy.matching.http.input>`
	// Example:
	//
	// .. validated-code-block:: yaml
	//
	//	:type-name: xds.type.matcher.v3.Matcher
	//
	//	matcher_list:
	//	  matchers:
	//	    # Apply a locally stored custom response to any 4xx response.
	//	  - predicate:
	//	      single_predicate:
	//	        input:
	//	          name: 4xx_response
	//	          typed_config:
	//	            "@type": type.googleapis.com/envoy.type.matcher.v3.HttpResponseStatusCodeClassMatchInput
	//	        value_match:
	//	          exact: "4xx"
	//	    on_match:
	//	      action:
	//	        name: action
	//	        typed_config:
	//	          "@type": type.googleapis.com/envoy.extensions.http.custom_response.local_response_policy.v3.LocalResponsePolicy
	//	          status_code: 499
	//	          body:
	//	            inline_string: "not allowed"
	//	          body_format:
	//	            json_format:
	//	              status: "%RESPONSE_CODE%"
	//	              message: "%LOCAL_REPLY_BODY%"
	//	          response_headers_to_add:
	//	          - header:
	//	              key: "foo"
	//	              value: "x-bar"
	//	    # Redirect to different upstream if the status code is one of 502, 503 or 504.
	//	  - predicate:
	//	      or_matcher:
	//	        predicate:
	//	        - single_predicate:
	//	            input:
	//	              name: "502_response"
	//	              typed_config:
	//	                "@type": type.googleapis.com/envoy.type.matcher.v3.HttpResponseStatusCodeMatchInput
	//	            value_match:
	//	              exact: "502"
	//	        - single_predicate:
	//	            input:
	//	              name: "503_response"
	//	              typed_config:
	//	                "@type": type.googleapis.com/envoy.type.matcher.v3.HttpResponseStatusCodeMatchInput
	//	            value_match:
	//	              exact: "503"
	//	        - single_predicate:
	//	            input:
	//	              name: "504_response"
	//	              typed_config:
	//	                "@type": type.googleapis.com/envoy.type.matcher.v3.HttpResponseStatusCodeMatchInput
	//	            value_match:
	//	              exact: "504"
	//	    on_match:
	//	      action:
	//	        name: action
	//	        typed_config:
	//	          "@type": type.googleapis.com/envoy.extensions.http.custom_response.redirect_policy.v3.RedirectPolicy
	//	          status_code: 299
	//	          uri: "https://foo.example/gateway_error"
	//	          response_headers_to_add:
	//	          - header:
	//	              key: "foo2"
	//	              value: "x-bar2"
	//
	// -- attention::
	//
	//	The first matched policy wins. Once the response is matched, matcher
	//	evaluations end.
	//
	// Refer to :ref:`Unified Matcher API <envoy_v3_api_msg_.xds.type.matcher.v3.Matcher>`
	// documentation for more information on the matcher trees.
	// [#extension-category: envoy.http.custom_response]
	CustomResponseMatcher *v3.Matcher `protobuf:"bytes,1,opt,name=custom_response_matcher,json=customResponseMatcher,proto3" json:"custom_response_matcher,omitempty"`
}

func (x *CustomResponse) Reset() {
	*x = CustomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomResponse) ProtoMessage() {}

func (x *CustomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomResponse.ProtoReflect.Descriptor instead.
func (*CustomResponse) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDescGZIP(), []int{0}
}

func (x *CustomResponse) GetCustomResponseMatcher() *v3.Matcher {
	if x != nil {
		return x.CustomResponseMatcher
	}
	return nil
}

var File_envoy_extensions_filters_http_custom_response_v3_custom_response_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDesc = []byte{
	0x0a, 0x46, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x78, 0x64, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a,
	0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x15,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0xd2, 0x01, 0x0a, 0x3e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDescData = file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDesc
)

func file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDescData
}

var file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_goTypes = []interface{}{
	(*CustomResponse)(nil), // 0: envoy.extensions.filters.http.custom_response.v3.CustomResponse
	(*v3.Matcher)(nil),     // 1: xds.type.matcher.v3.Matcher
}
var file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.custom_response.v3.CustomResponse.custom_response_matcher:type_name -> xds.type.matcher.v3.Matcher
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_init() }
func file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_init() {
	if File_envoy_extensions_filters_http_custom_response_v3_custom_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomResponse); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_custom_response_v3_custom_response_proto = out.File
	file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_rawDesc = nil
	file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_goTypes = nil
	file_envoy_extensions_filters_http_custom_response_v3_custom_response_proto_depIdxs = nil
}
