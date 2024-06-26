// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: models/movie/v1/movie.proto

package moviev1

import (
	v1 "github.com/nareshbhatia/movie-magic-services-go/gen/go/models/common/v1"
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

type Certificate int32

const (
	Certificate_CERTIFICATE_UNSPECIFIED Certificate = 0
	Certificate_CERTIFICATE_G           Certificate = 1
	Certificate_CERTIFICATE_NR          Certificate = 2
	Certificate_CERTIFICATE_PG_13       Certificate = 3
	Certificate_CERTIFICATE_PG          Certificate = 4
	Certificate_CERTIFICATE_R           Certificate = 5
)

// Enum value maps for Certificate.
var (
	Certificate_name = map[int32]string{
		0: "CERTIFICATE_UNSPECIFIED",
		1: "CERTIFICATE_G",
		2: "CERTIFICATE_NR",
		3: "CERTIFICATE_PG_13",
		4: "CERTIFICATE_PG",
		5: "CERTIFICATE_R",
	}
	Certificate_value = map[string]int32{
		"CERTIFICATE_UNSPECIFIED": 0,
		"CERTIFICATE_G":           1,
		"CERTIFICATE_NR":          2,
		"CERTIFICATE_PG_13":       3,
		"CERTIFICATE_PG":          4,
		"CERTIFICATE_R":           5,
	}
)

func (x Certificate) Enum() *Certificate {
	p := new(Certificate)
	*p = x
	return p
}

func (x Certificate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Certificate) Descriptor() protoreflect.EnumDescriptor {
	return file_models_movie_v1_movie_proto_enumTypes[0].Descriptor()
}

func (Certificate) Type() protoreflect.EnumType {
	return &file_models_movie_v1_movie_proto_enumTypes[0]
}

func (x Certificate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Certificate.Descriptor instead.
func (Certificate) EnumDescriptor() ([]byte, []int) {
	return file_models_movie_v1_movie_proto_rawDescGZIP(), []int{0}
}

type CastMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId   string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	Characters []string `protobuf:"bytes,2,rep,name=characters,proto3" json:"characters,omitempty"`
}

func (x *CastMember) Reset() {
	*x = CastMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_v1_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CastMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CastMember) ProtoMessage() {}

func (x *CastMember) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_v1_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CastMember.ProtoReflect.Descriptor instead.
func (*CastMember) Descriptor() ([]byte, []int) {
	return file_models_movie_v1_movie_proto_rawDescGZIP(), []int{0}
}

func (x *CastMember) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

func (x *CastMember) GetCharacters() []string {
	if x != nil {
		return x.Characters
	}
	return nil
}

type RatingsSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregateRating float64 `protobuf:"fixed64,1,opt,name=aggregate_rating,json=aggregateRating,proto3" json:"aggregate_rating,omitempty"`
	VoteCount       int32   `protobuf:"varint,2,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
}

func (x *RatingsSummary) Reset() {
	*x = RatingsSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_v1_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingsSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingsSummary) ProtoMessage() {}

func (x *RatingsSummary) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_v1_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingsSummary.ProtoReflect.Descriptor instead.
func (*RatingsSummary) Descriptor() ([]byte, []int) {
	return file_models_movie_v1_movie_proto_rawDescGZIP(), []int{1}
}

func (x *RatingsSummary) GetAggregateRating() float64 {
	if x != nil {
		return x.AggregateRating
	}
	return 0
}

func (x *RatingsSummary) GetVoteCount() int32 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description    string          `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Cast           []*CastMember   `protobuf:"bytes,4,rep,name=cast,proto3" json:"cast,omitempty"`
	Certificate    Certificate     `protobuf:"varint,5,opt,name=certificate,proto3,enum=models.movie.v1.Certificate" json:"certificate,omitempty"`
	Genres         []string        `protobuf:"bytes,6,rep,name=genres,proto3" json:"genres,omitempty"`
	Image          *v1.Image       `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
	Rank           int32           `protobuf:"varint,8,opt,name=rank,proto3" json:"rank,omitempty"`
	RatingsSummary *RatingsSummary `protobuf:"bytes,9,opt,name=ratings_summary,json=ratingsSummary,proto3" json:"ratings_summary,omitempty"`
	ReleaseYear    int32           `protobuf:"varint,10,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	Runtime        int32           `protobuf:"varint,11,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Tagline        string          `protobuf:"bytes,12,opt,name=tagline,proto3" json:"tagline,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_movie_v1_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_models_movie_v1_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_models_movie_v1_movie_proto_rawDescGZIP(), []int{2}
}

func (x *Movie) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Movie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Movie) GetCast() []*CastMember {
	if x != nil {
		return x.Cast
	}
	return nil
}

func (x *Movie) GetCertificate() Certificate {
	if x != nil {
		return x.Certificate
	}
	return Certificate_CERTIFICATE_UNSPECIFIED
}

func (x *Movie) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Movie) GetImage() *v1.Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Movie) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Movie) GetRatingsSummary() *RatingsSummary {
	if x != nil {
		return x.RatingsSummary
	}
	return nil
}

func (x *Movie) GetReleaseYear() int32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *Movie) GetRuntime() int32 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

func (x *Movie) GetTagline() string {
	if x != nil {
		return x.Tagline
	}
	return ""
}

var File_models_movie_v1_movie_proto protoreflect.FileDescriptor

var file_models_movie_v1_movie_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x0a, 0x43, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x22, 0x5a, 0x0a, 0x0e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xba, 0x03, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x04, 0x63, 0x61, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x2d, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x12, 0x48, 0x0a, 0x0f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x67, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x6c, 0x69, 0x6e,
	0x65, 0x2a, 0x8f, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x5f, 0x4e, 0x52, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49,
	0x43, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x47, 0x5f, 0x31, 0x33, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e,
	0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x47, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f,
	0x52, 0x10, 0x05, 0x42, 0xcf, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x72, 0x65, 0x73, 0x68, 0x62, 0x68, 0x61, 0x74,
	0x69, 0x61, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2d, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x4d, 0x58, 0xaa,
	0x02, 0x0f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3a, 0x3a, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_movie_v1_movie_proto_rawDescOnce sync.Once
	file_models_movie_v1_movie_proto_rawDescData = file_models_movie_v1_movie_proto_rawDesc
)

func file_models_movie_v1_movie_proto_rawDescGZIP() []byte {
	file_models_movie_v1_movie_proto_rawDescOnce.Do(func() {
		file_models_movie_v1_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_movie_v1_movie_proto_rawDescData)
	})
	return file_models_movie_v1_movie_proto_rawDescData
}

var file_models_movie_v1_movie_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_models_movie_v1_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_models_movie_v1_movie_proto_goTypes = []interface{}{
	(Certificate)(0),       // 0: models.movie.v1.Certificate
	(*CastMember)(nil),     // 1: models.movie.v1.CastMember
	(*RatingsSummary)(nil), // 2: models.movie.v1.RatingsSummary
	(*Movie)(nil),          // 3: models.movie.v1.Movie
	(*v1.Image)(nil),       // 4: models.common.v1.Image
}
var file_models_movie_v1_movie_proto_depIdxs = []int32{
	1, // 0: models.movie.v1.Movie.cast:type_name -> models.movie.v1.CastMember
	0, // 1: models.movie.v1.Movie.certificate:type_name -> models.movie.v1.Certificate
	4, // 2: models.movie.v1.Movie.image:type_name -> models.common.v1.Image
	2, // 3: models.movie.v1.Movie.ratings_summary:type_name -> models.movie.v1.RatingsSummary
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_models_movie_v1_movie_proto_init() }
func file_models_movie_v1_movie_proto_init() {
	if File_models_movie_v1_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_movie_v1_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CastMember); i {
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
		file_models_movie_v1_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingsSummary); i {
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
		file_models_movie_v1_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
			RawDescriptor: file_models_movie_v1_movie_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_movie_v1_movie_proto_goTypes,
		DependencyIndexes: file_models_movie_v1_movie_proto_depIdxs,
		EnumInfos:         file_models_movie_v1_movie_proto_enumTypes,
		MessageInfos:      file_models_movie_v1_movie_proto_msgTypes,
	}.Build()
	File_models_movie_v1_movie_proto = out.File
	file_models_movie_v1_movie_proto_rawDesc = nil
	file_models_movie_v1_movie_proto_goTypes = nil
	file_models_movie_v1_movie_proto_depIdxs = nil
}
