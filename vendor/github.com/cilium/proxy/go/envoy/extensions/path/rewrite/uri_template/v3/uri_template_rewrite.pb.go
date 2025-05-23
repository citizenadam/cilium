// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.29.2
// source: envoy/extensions/path/rewrite/uri_template/v3/uri_template_rewrite.proto

package uri_templatev3

import (
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

// Indicates that during forwarding, portions of the path that match the
// pattern should be rewritten, even allowing the substitution of variables
// from the match pattern into the new path as specified by the rewrite template.
// This is useful to allow application paths to be
// rewritten in a way that is aware of segments with variable content like
// identifiers. The router filter will place the original path as it was
// before the rewrite into the :ref:`x-envoy-original-path
// <config_http_filters_router_x-envoy-original-path>` header.
//
// Only one of :ref:`prefix_rewrite <envoy_v3_api_field_config.route.v3.RouteAction.prefix_rewrite>`,
// :ref:`regex_rewrite <envoy_v3_api_field_config.route.v3.RouteAction.regex_rewrite>`,
// or *path_template_rewrite* may be specified.
//
// Template pattern matching types:
//
// * “*“ : Matches a single path component, up to the next path separator: /
//
// * “**“ : Matches zero or more path segments. If present, must be the last operator.
//
// * “{name} or {name=*}“ :  A named variable matching one path segment up to the next path separator: /.
//
//   - “{name=videos/*}“ : A named variable matching more than one path segment.
//     The path component matching videos/* is captured as the named variable.
//
// * “{name=**}“ : A named variable matching zero or more path segments.
//
// Only named matches can be used to perform rewrites.
//
// Examples using path_template_rewrite:
//
//   - The pattern “/{one}/{two}“ paired with a substitution string of “/{two}/{one}“ would
//     transform “/cat/dog“ into “/dog/cat“.
//
//   - The pattern “/videos/{language=lang/*}/*“ paired with a substitution string of
//     “/{language}“ would transform “/videos/lang/en/video.m4s“ into “lang/en“.
//
//   - The path pattern “/content/{format}/{lang}/{id}/{file}.vtt“ paired with a substitution
//     string of “/{lang}/{format}/{file}.vtt“ would transform “/content/hls/en-us/12345/en_193913.vtt“
//     into “/en-us/hls/en_193913.vtt“.
type UriTemplateRewriteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PathTemplateRewrite string `protobuf:"bytes,1,opt,name=path_template_rewrite,json=pathTemplateRewrite,proto3" json:"path_template_rewrite,omitempty"`
}

func (x *UriTemplateRewriteConfig) Reset() {
	*x = UriTemplateRewriteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UriTemplateRewriteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UriTemplateRewriteConfig) ProtoMessage() {}

func (x *UriTemplateRewriteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UriTemplateRewriteConfig.ProtoReflect.Descriptor instead.
func (*UriTemplateRewriteConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDescGZIP(), []int{0}
}

func (x *UriTemplateRewriteConfig) GetPathTemplateRewrite() string {
	if x != nil {
		return x.PathTemplateRewrite
	}
	return ""
}

var File_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto protoreflect.FileDescriptor

var file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDesc = []byte{
	0x0a, 0x48, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2f,
	0x75, 0x72, 0x69, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x75, 0x72, 0x69, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x61, 0x74,
	0x68, 0x2e, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x75, 0x72, 0x69, 0x5f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5a, 0x0a, 0x18, 0x55, 0x72, 0x69, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a,
	0x15, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x02, 0x52, 0x13, 0x70, 0x61, 0x74, 0x68, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0xc5, 0x01,
	0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x75, 0x72,
	0x69, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x17, 0x55,
	0x72, 0x69, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x75,
	0x72, 0x69, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x75,
	0x72, 0x69, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x76, 0x33, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDescOnce sync.Once
	file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDescData = file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDesc
)

func file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDescGZIP() []byte {
	file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDescData)
	})
	return file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDescData
}

var file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_goTypes = []interface{}{
	(*UriTemplateRewriteConfig)(nil), // 0: envoy.extensions.path.rewrite.uri_template.v3.UriTemplateRewriteConfig
}
var file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_init() }
func file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_init() {
	if File_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UriTemplateRewriteConfig); i {
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
			RawDescriptor: file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_msgTypes,
	}.Build()
	File_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto = out.File
	file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_rawDesc = nil
	file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_goTypes = nil
	file_envoy_extensions_path_rewrite_uri_template_v3_uri_template_rewrite_proto_depIdxs = nil
}
