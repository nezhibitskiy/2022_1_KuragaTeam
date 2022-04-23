// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: movie.proto

package proto

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

type GetMovieOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieID int64 `protobuf:"varint,1,opt,name=MovieID,proto3" json:"MovieID,omitempty"`
}

func (x *GetMovieOptions) Reset() {
	*x = GetMovieOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieOptions) ProtoMessage() {}

func (x *GetMovieOptions) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieOptions.ProtoReflect.Descriptor instead.
func (*GetMovieOptions) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *GetMovieOptions) GetMovieID() int64 {
	if x != nil {
		return x.MovieID
	}
	return 0
}

type GetRandomOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetRandomOptions) Reset() {
	*x = GetRandomOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRandomOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRandomOptions) ProtoMessage() {}

func (x *GetRandomOptions) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRandomOptions.ProtoReflect.Descriptor instead.
func (*GetRandomOptions) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *GetRandomOptions) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetRandomOptions) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetMainMovieOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMainMovieOptions) Reset() {
	*x = GetMainMovieOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMainMovieOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMainMovieOptions) ProtoMessage() {}

func (x *GetMainMovieOptions) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMainMovieOptions.ProtoReflect.Descriptor instead.
func (*GetMainMovieOptions) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{2}
}

type PersonInMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Photo    string `protobuf:"bytes,3,opt,name=Photo,proto3" json:"Photo,omitempty"`
	Position string `protobuf:"bytes,4,opt,name=Position,proto3" json:"Position,omitempty"`
}

func (x *PersonInMovie) Reset() {
	*x = PersonInMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonInMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonInMovie) ProtoMessage() {}

func (x *PersonInMovie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonInMovie.ProtoReflect.Descriptor instead.
func (*PersonInMovie) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{3}
}

func (x *PersonInMovie) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PersonInMovie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonInMovie) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *PersonInMovie) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              int64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name            string           `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	NamePicture     string           `protobuf:"bytes,3,opt,name=NamePicture,proto3" json:"NamePicture,omitempty"`
	Year            int32            `protobuf:"varint,4,opt,name=Year,proto3" json:"Year,omitempty"`
	Duration        string           `protobuf:"bytes,5,opt,name=Duration,proto3" json:"Duration,omitempty"`
	AgeLimit        int32            `protobuf:"varint,6,opt,name=AgeLimit,proto3" json:"AgeLimit,omitempty"`
	Description     string           `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`
	KinopoiskRating float32          `protobuf:"fixed32,8,opt,name=KinopoiskRating,proto3" json:"KinopoiskRating,omitempty"`
	Rating          float32          `protobuf:"fixed32,9,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Tagline         string           `protobuf:"bytes,10,opt,name=Tagline,proto3" json:"Tagline,omitempty"`
	Picture         string           `protobuf:"bytes,11,opt,name=Picture,proto3" json:"Picture,omitempty"`
	Video           string           `protobuf:"bytes,12,opt,name=Video,proto3" json:"Video,omitempty"`
	Trailer         string           `protobuf:"bytes,13,opt,name=Trailer,proto3" json:"Trailer,omitempty"`
	Country         []string         `protobuf:"bytes,14,rep,name=Country,proto3" json:"Country,omitempty"`
	Genre           []string         `protobuf:"bytes,15,rep,name=Genre,proto3" json:"Genre,omitempty"`
	Staff           []*PersonInMovie `protobuf:"bytes,16,rep,name=Staff,proto3" json:"Staff,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[4]
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
	return file_movie_proto_rawDescGZIP(), []int{4}
}

func (x *Movie) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Movie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Movie) GetNamePicture() string {
	if x != nil {
		return x.NamePicture
	}
	return ""
}

func (x *Movie) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Movie) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Movie) GetAgeLimit() int32 {
	if x != nil {
		return x.AgeLimit
	}
	return 0
}

func (x *Movie) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Movie) GetKinopoiskRating() float32 {
	if x != nil {
		return x.KinopoiskRating
	}
	return 0
}

func (x *Movie) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Movie) GetTagline() string {
	if x != nil {
		return x.Tagline
	}
	return ""
}

func (x *Movie) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Movie) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

func (x *Movie) GetTrailer() string {
	if x != nil {
		return x.Trailer
	}
	return ""
}

func (x *Movie) GetCountry() []string {
	if x != nil {
		return x.Country
	}
	return nil
}

func (x *Movie) GetGenre() []string {
	if x != nil {
		return x.Genre
	}
	return nil
}

func (x *Movie) GetStaff() []*PersonInMovie {
	if x != nil {
		return x.Staff
	}
	return nil
}

type MoviesArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie []*Movie `protobuf:"bytes,1,rep,name=Movie,proto3" json:"Movie,omitempty"`
}

func (x *MoviesArr) Reset() {
	*x = MoviesArr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoviesArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoviesArr) ProtoMessage() {}

func (x *MoviesArr) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoviesArr.ProtoReflect.Descriptor instead.
func (*MoviesArr) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{5}
}

func (x *MoviesArr) GetMovie() []*Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type MainMovie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	NamePicture string `protobuf:"bytes,2,opt,name=NamePicture,proto3" json:"NamePicture,omitempty"`
	Tagline     string `protobuf:"bytes,3,opt,name=Tagline,proto3" json:"Tagline,omitempty"`
	Picture     string `protobuf:"bytes,4,opt,name=Picture,proto3" json:"Picture,omitempty"`
}

func (x *MainMovie) Reset() {
	*x = MainMovie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainMovie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainMovie) ProtoMessage() {}

func (x *MainMovie) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainMovie.ProtoReflect.Descriptor instead.
func (*MainMovie) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{6}
}

func (x *MainMovie) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *MainMovie) GetNamePicture() string {
	if x != nil {
		return x.NamePicture
	}
	return ""
}

func (x *MainMovie) GetTagline() string {
	if x != nil {
		return x.Tagline
	}
	return ""
}

func (x *MainMovie) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x15, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x65, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb7, 0x03, 0x0a, 0x05, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65,
	0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e,
	0x61, 0x6d, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65,
	0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x67,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x41, 0x67,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x4b, 0x69, 0x6e, 0x6f,
	0x70, 0x6f, 0x69, 0x73, 0x6b, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0f, 0x4b, 0x69, 0x6e, 0x6f, 0x70, 0x6f, 0x69, 0x73, 0x6b, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61,
	0x67, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x61, 0x67,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x24,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x22, 0x29, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x41, 0x72,
	0x72, 0x12, 0x1c, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x22,
	0x71, 0x0a, 0x09, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x4e, 0x61, 0x6d, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x54, 0x61, 0x67, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x54, 0x61, 0x67, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x32, 0x91, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x06, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x0a, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x41, 0x72, 0x72,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x0a, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_movie_proto_goTypes = []interface{}{
	(*GetMovieOptions)(nil),     // 0: GetMovieOptions
	(*GetRandomOptions)(nil),    // 1: GetRandomOptions
	(*GetMainMovieOptions)(nil), // 2: GetMainMovieOptions
	(*PersonInMovie)(nil),       // 3: PersonInMovie
	(*Movie)(nil),               // 4: Movie
	(*MoviesArr)(nil),           // 5: MoviesArr
	(*MainMovie)(nil),           // 6: MainMovie
}
var file_movie_proto_depIdxs = []int32{
	3, // 0: Movie.Staff:type_name -> PersonInMovie
	4, // 1: MoviesArr.Movie:type_name -> Movie
	0, // 2: Movies.GetByID:input_type -> GetMovieOptions
	1, // 3: Movies.GetRandom:input_type -> GetRandomOptions
	2, // 4: Movies.GetMainMovie:input_type -> GetMainMovieOptions
	4, // 5: Movies.GetByID:output_type -> Movie
	5, // 6: Movies.GetRandom:output_type -> MoviesArr
	6, // 7: Movies.GetMainMovie:output_type -> MainMovie
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieOptions); i {
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
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRandomOptions); i {
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
		file_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMainMovieOptions); i {
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
		file_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonInMovie); i {
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
		file_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoviesArr); i {
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
		file_movie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainMovie); i {
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
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}
