// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: models/common/v1/common.proto

package commonv1

import (
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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url    string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width  int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_models_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_models_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Image) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Image) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_models_common_v1_common_proto protoreflect.FileDescriptor

var file_models_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x22, 0x47, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0xd7, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x61, 0x72, 0x65, 0x73, 0x68, 0x62, 0x68, 0x61, 0x74, 0x69, 0x61, 0x2f, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2d, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x43, 0x58, 0xaa, 0x02, 0x10, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x12, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_common_v1_common_proto_rawDescOnce sync.Once
	file_models_common_v1_common_proto_rawDescData = file_models_common_v1_common_proto_rawDesc
)

func file_models_common_v1_common_proto_rawDescGZIP() []byte {
	file_models_common_v1_common_proto_rawDescOnce.Do(func() {
		file_models_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_common_v1_common_proto_rawDescData)
	})
	return file_models_common_v1_common_proto_rawDescData
}

var file_models_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_common_v1_common_proto_goTypes = []interface{}{
	(*Image)(nil), // 0: models.common.v1.Image
}
var file_models_common_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_common_v1_common_proto_init() }
func file_models_common_v1_common_proto_init() {
	if File_models_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
			RawDescriptor: file_models_common_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_common_v1_common_proto_goTypes,
		DependencyIndexes: file_models_common_v1_common_proto_depIdxs,
		MessageInfos:      file_models_common_v1_common_proto_msgTypes,
	}.Build()
	File_models_common_v1_common_proto = out.File
	file_models_common_v1_common_proto_rawDesc = nil
	file_models_common_v1_common_proto_goTypes = nil
	file_models_common_v1_common_proto_depIdxs = nil
}
