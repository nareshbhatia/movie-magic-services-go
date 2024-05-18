// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: seed/v1/seed.proto

package seedv1

import (
	v1 "github.com/code-shaper/movie-magic/gen/go/models/movie/v1"
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

type Wrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*v1.Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
}

func (x *Wrapper) Reset() {
	*x = Wrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_seed_v1_seed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrapper) ProtoMessage() {}

func (x *Wrapper) ProtoReflect() protoreflect.Message {
	mi := &file_seed_v1_seed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrapper.ProtoReflect.Descriptor instead.
func (*Wrapper) Descriptor() ([]byte, []int) {
	return file_seed_v1_seed_proto_rawDescGZIP(), []int{0}
}

func (x *Wrapper) GetMovies() []*v1.Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

var File_seed_v1_seed_proto protoreflect.FileDescriptor

var file_seed_v1_seed_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x07, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x42, 0x8f, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65,
	0x65, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x53, 0x65, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x2d, 0x73, 0x68, 0x61, 0x70, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2d, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65,
	0x65, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x65, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53,
	0x58, 0x58, 0xaa, 0x02, 0x07, 0x53, 0x65, 0x65, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x53,
	0x65, 0x65, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x53, 0x65, 0x65, 0x64, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x53,
	0x65, 0x65, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_seed_v1_seed_proto_rawDescOnce sync.Once
	file_seed_v1_seed_proto_rawDescData = file_seed_v1_seed_proto_rawDesc
)

func file_seed_v1_seed_proto_rawDescGZIP() []byte {
	file_seed_v1_seed_proto_rawDescOnce.Do(func() {
		file_seed_v1_seed_proto_rawDescData = protoimpl.X.CompressGZIP(file_seed_v1_seed_proto_rawDescData)
	})
	return file_seed_v1_seed_proto_rawDescData
}

var file_seed_v1_seed_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_seed_v1_seed_proto_goTypes = []interface{}{
	(*Wrapper)(nil),  // 0: seed.v1.Wrapper
	(*v1.Movie)(nil), // 1: models.movie.v1.Movie
}
var file_seed_v1_seed_proto_depIdxs = []int32{
	1, // 0: seed.v1.Wrapper.movies:type_name -> models.movie.v1.Movie
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_seed_v1_seed_proto_init() }
func file_seed_v1_seed_proto_init() {
	if File_seed_v1_seed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_seed_v1_seed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrapper); i {
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
			RawDescriptor: file_seed_v1_seed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_seed_v1_seed_proto_goTypes,
		DependencyIndexes: file_seed_v1_seed_proto_depIdxs,
		MessageInfos:      file_seed_v1_seed_proto_msgTypes,
	}.Build()
	File_seed_v1_seed_proto = out.File
	file_seed_v1_seed_proto_rawDesc = nil
	file_seed_v1_seed_proto_goTypes = nil
	file_seed_v1_seed_proto_depIdxs = nil
}