// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: movie/v1/movie.proto

package moviev1

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

type SortParam int32

const (
	SortParam_SORT_PARAM_UNSPECIFIED SortParam = 0
	SortParam_SORT_PARAM_RANK_ASC    SortParam = 1
)

// Enum value maps for SortParam.
var (
	SortParam_name = map[int32]string{
		0: "SORT_PARAM_UNSPECIFIED",
		1: "SORT_PARAM_RANK_ASC",
	}
	SortParam_value = map[string]int32{
		"SORT_PARAM_UNSPECIFIED": 0,
		"SORT_PARAM_RANK_ASC":    1,
	}
)

func (x SortParam) Enum() *SortParam {
	p := new(SortParam)
	*p = x
	return p
}

func (x SortParam) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortParam) Descriptor() protoreflect.EnumDescriptor {
	return file_movie_v1_movie_proto_enumTypes[0].Descriptor()
}

func (SortParam) Type() protoreflect.EnumType {
	return &file_movie_v1_movie_proto_enumTypes[0]
}

func (x SortParam) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortParam.Descriptor instead.
func (SortParam) EnumDescriptor() ([]byte, []int) {
	return file_movie_v1_movie_proto_rawDescGZIP(), []int{0}
}

type FilterParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certs []v1.Certificate `protobuf:"varint,1,rep,packed,name=certs,proto3,enum=models.movie.v1.Certificate" json:"certs,omitempty"`
}

func (x *FilterParams) Reset() {
	*x = FilterParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_v1_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterParams) ProtoMessage() {}

func (x *FilterParams) ProtoReflect() protoreflect.Message {
	mi := &file_movie_v1_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterParams.ProtoReflect.Descriptor instead.
func (*FilterParams) Descriptor() ([]byte, []int) {
	return file_movie_v1_movie_proto_rawDescGZIP(), []int{0}
}

func (x *FilterParams) GetCerts() []v1.Certificate {
	if x != nil {
		return x.Certs
	}
	return nil
}

type PageSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *PageSpec) Reset() {
	*x = PageSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_v1_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageSpec) ProtoMessage() {}

func (x *PageSpec) ProtoReflect() protoreflect.Message {
	mi := &file_movie_v1_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageSpec.ProtoReflect.Descriptor instead.
func (*PageSpec) Descriptor() ([]byte, []int) {
	return file_movie_v1_movie_proto_rawDescGZIP(), []int{1}
}

func (x *PageSpec) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageSpec) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters  *FilterParams `protobuf:"bytes,1,opt,name=filters,proto3" json:"filters,omitempty"`
	Sort     SortParam     `protobuf:"varint,2,opt,name=sort,proto3,enum=movie.v1.SortParam" json:"sort,omitempty"`
	PageSpec *PageSpec     `protobuf:"bytes,3,opt,name=page_spec,json=pageSpec,proto3" json:"page_spec,omitempty"`
}

func (x *ListMoviesRequest) Reset() {
	*x = ListMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_v1_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesRequest) ProtoMessage() {}

func (x *ListMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_v1_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesRequest.ProtoReflect.Descriptor instead.
func (*ListMoviesRequest) Descriptor() ([]byte, []int) {
	return file_movie_v1_movie_proto_rawDescGZIP(), []int{2}
}

func (x *ListMoviesRequest) GetFilters() *FilterParams {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListMoviesRequest) GetSort() SortParam {
	if x != nil {
		return x.Sort
	}
	return SortParam_SORT_PARAM_UNSPECIFIED
}

func (x *ListMoviesRequest) GetPageSpec() *PageSpec {
	if x != nil {
		return x.PageSpec
	}
	return nil
}

type PaginationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// total number of pages
	TotalPages int32 `protobuf:"varint,1,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	// total number of items
	TotalItems int32 `protobuf:"varint,2,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	// current page number
	Page int32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	// number of items per page
	PerPage int32 `protobuf:"varint,4,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	// when paginating forwards, are there more items?
	HasNextPage bool `protobuf:"varint,5,opt,name=has_next_page,json=hasNextPage,proto3" json:"has_next_page,omitempty"`
	// when paginating backwards, are there more items?
	HasPreviousPage bool `protobuf:"varint,6,opt,name=has_previous_page,json=hasPreviousPage,proto3" json:"has_previous_page,omitempty"`
}

func (x *PaginationInfo) Reset() {
	*x = PaginationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_v1_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationInfo) ProtoMessage() {}

func (x *PaginationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_movie_v1_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationInfo.ProtoReflect.Descriptor instead.
func (*PaginationInfo) Descriptor() ([]byte, []int) {
	return file_movie_v1_movie_proto_rawDescGZIP(), []int{3}
}

func (x *PaginationInfo) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *PaginationInfo) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *PaginationInfo) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationInfo) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *PaginationInfo) GetHasNextPage() bool {
	if x != nil {
		return x.HasNextPage
	}
	return false
}

func (x *PaginationInfo) GetHasPreviousPage() bool {
	if x != nil {
		return x.HasPreviousPage
	}
	return false
}

type ListMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies   []*v1.Movie     `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	PageInfo *PaginationInfo `protobuf:"bytes,2,opt,name=page_info,json=pageInfo,proto3" json:"page_info,omitempty"`
}

func (x *ListMoviesResponse) Reset() {
	*x = ListMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_v1_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMoviesResponse) ProtoMessage() {}

func (x *ListMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_v1_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMoviesResponse.ProtoReflect.Descriptor instead.
func (*ListMoviesResponse) Descriptor() ([]byte, []int) {
	return file_movie_v1_movie_proto_rawDescGZIP(), []int{4}
}

func (x *ListMoviesResponse) GetMovies() []*v1.Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

func (x *ListMoviesResponse) GetPageInfo() *PaginationInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

var File_movie_v1_movie_proto protoreflect.FileDescriptor

var file_movie_v1_movie_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x32, 0x0a,
	0x05, 0x63, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x05, 0x63, 0x65, 0x72, 0x74,
	0x73, 0x22, 0x39, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x9f, 0x01, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x2f, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x22, 0xd1,
	0x01, 0x0a, 0x0e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61, 0x73, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x68, 0x61, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x50, 0x61,
	0x67, 0x65, 0x22, 0x7b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2a,
	0x40, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x4f, 0x52, 0x54, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x4f, 0x52, 0x54,
	0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x41, 0x53, 0x43, 0x10,
	0x01, 0x32, 0x59, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x49, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12,
	0x1b, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x97, 0x01, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2d, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x08,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_v1_movie_proto_rawDescOnce sync.Once
	file_movie_v1_movie_proto_rawDescData = file_movie_v1_movie_proto_rawDesc
)

func file_movie_v1_movie_proto_rawDescGZIP() []byte {
	file_movie_v1_movie_proto_rawDescOnce.Do(func() {
		file_movie_v1_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_v1_movie_proto_rawDescData)
	})
	return file_movie_v1_movie_proto_rawDescData
}

var file_movie_v1_movie_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_movie_v1_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_movie_v1_movie_proto_goTypes = []interface{}{
	(SortParam)(0),             // 0: movie.v1.SortParam
	(*FilterParams)(nil),       // 1: movie.v1.FilterParams
	(*PageSpec)(nil),           // 2: movie.v1.PageSpec
	(*ListMoviesRequest)(nil),  // 3: movie.v1.ListMoviesRequest
	(*PaginationInfo)(nil),     // 4: movie.v1.PaginationInfo
	(*ListMoviesResponse)(nil), // 5: movie.v1.ListMoviesResponse
	(v1.Certificate)(0),        // 6: models.movie.v1.Certificate
	(*v1.Movie)(nil),           // 7: models.movie.v1.Movie
}
var file_movie_v1_movie_proto_depIdxs = []int32{
	6, // 0: movie.v1.FilterParams.certs:type_name -> models.movie.v1.Certificate
	1, // 1: movie.v1.ListMoviesRequest.filters:type_name -> movie.v1.FilterParams
	0, // 2: movie.v1.ListMoviesRequest.sort:type_name -> movie.v1.SortParam
	2, // 3: movie.v1.ListMoviesRequest.page_spec:type_name -> movie.v1.PageSpec
	7, // 4: movie.v1.ListMoviesResponse.movies:type_name -> models.movie.v1.Movie
	4, // 5: movie.v1.ListMoviesResponse.page_info:type_name -> movie.v1.PaginationInfo
	3, // 6: movie.v1.MovieService.ListMovies:input_type -> movie.v1.ListMoviesRequest
	5, // 7: movie.v1.MovieService.ListMovies:output_type -> movie.v1.ListMoviesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_movie_v1_movie_proto_init() }
func file_movie_v1_movie_proto_init() {
	if File_movie_v1_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_v1_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterParams); i {
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
		file_movie_v1_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageSpec); i {
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
		file_movie_v1_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMoviesRequest); i {
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
		file_movie_v1_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationInfo); i {
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
		file_movie_v1_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMoviesResponse); i {
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
			RawDescriptor: file_movie_v1_movie_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_v1_movie_proto_goTypes,
		DependencyIndexes: file_movie_v1_movie_proto_depIdxs,
		EnumInfos:         file_movie_v1_movie_proto_enumTypes,
		MessageInfos:      file_movie_v1_movie_proto_msgTypes,
	}.Build()
	File_movie_v1_movie_proto = out.File
	file_movie_v1_movie_proto_rawDesc = nil
	file_movie_v1_movie_proto_goTypes = nil
	file_movie_v1_movie_proto_depIdxs = nil
}
