// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: microservices/compilations/proto/compilations.proto

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

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{0}
}

func (x *Genre) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MovieInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Genre   []*Genre `protobuf:"bytes,3,rep,name=Genre,proto3" json:"Genre,omitempty"`
	Picture string   `protobuf:"bytes,4,opt,name=Picture,proto3" json:"Picture,omitempty"`
}

func (x *MovieInfo) Reset() {
	*x = MovieInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieInfo) ProtoMessage() {}

func (x *MovieInfo) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieInfo.ProtoReflect.Descriptor instead.
func (*MovieInfo) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{1}
}

func (x *MovieInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *MovieInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MovieInfo) GetGenre() []*Genre {
	if x != nil {
		return x.Genre
	}
	return nil
}

func (x *MovieInfo) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

type MovieCompilation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string       `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Movies []*MovieInfo `protobuf:"bytes,2,rep,name=Movies,proto3" json:"Movies,omitempty"`
}

func (x *MovieCompilation) Reset() {
	*x = MovieCompilation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieCompilation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieCompilation) ProtoMessage() {}

func (x *MovieCompilation) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieCompilation.ProtoReflect.Descriptor instead.
func (*MovieCompilation) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{2}
}

func (x *MovieCompilation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MovieCompilation) GetMovies() []*MovieInfo {
	if x != nil {
		return x.Movies
	}
	return nil
}

type MovieCompilationsArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieCompilations []*MovieCompilation `protobuf:"bytes,1,rep,name=MovieCompilations,proto3" json:"MovieCompilations,omitempty"`
}

func (x *MovieCompilationsArr) Reset() {
	*x = MovieCompilationsArr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieCompilationsArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieCompilationsArr) ProtoMessage() {}

func (x *MovieCompilationsArr) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieCompilationsArr.ProtoReflect.Descriptor instead.
func (*MovieCompilationsArr) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{3}
}

func (x *MovieCompilationsArr) GetMovieCompilations() []*MovieCompilation {
	if x != nil {
		return x.MovieCompilations
	}
	return nil
}

type GetMainCompilationsOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMainCompilationsOptions) Reset() {
	*x = GetMainCompilationsOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMainCompilationsOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMainCompilationsOptions) ProtoMessage() {}

func (x *GetMainCompilationsOptions) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMainCompilationsOptions.ProtoReflect.Descriptor instead.
func (*GetMainCompilationsOptions) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{4}
}

type GetByIDOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetByIDOptions) Reset() {
	*x = GetByIDOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDOptions) ProtoMessage() {}

func (x *GetByIDOptions) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDOptions.ProtoReflect.Descriptor instead.
func (*GetByIDOptions) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{5}
}

func (x *GetByIDOptions) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetByIDOptions) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetByIDOptions) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetCompilationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetCompilationOptions) Reset() {
	*x = GetCompilationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompilationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompilationOptions) ProtoMessage() {}

func (x *GetCompilationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompilationOptions.ProtoReflect.Descriptor instead.
func (*GetCompilationOptions) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{6}
}

func (x *GetCompilationOptions) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetCompilationOptions) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetFavoritesOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFavoritesOptions) Reset() {
	*x = GetFavoritesOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFavoritesOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFavoritesOptions) ProtoMessage() {}

func (x *GetFavoritesOptions) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFavoritesOptions.ProtoReflect.Descriptor instead.
func (*GetFavoritesOptions) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{7}
}

func (x *GetFavoritesOptions) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type SearchText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	IsMovie bool   `protobuf:"varint,2,opt,name=is_movie,json=isMovie,proto3" json:"is_movie,omitempty"`
}

func (x *SearchText) Reset() {
	*x = SearchText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchText) ProtoMessage() {}

func (x *SearchText) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchText.ProtoReflect.Descriptor instead.
func (*SearchText) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{8}
}

func (x *SearchText) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SearchText) GetIsMovie() bool {
	if x != nil {
		return x.IsMovie
	}
	return false
}

type PersonInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Photo    string   `protobuf:"bytes,3,opt,name=Photo,proto3" json:"Photo,omitempty"`
	Position []string `protobuf:"bytes,4,rep,name=position,proto3" json:"position,omitempty"`
}

func (x *PersonInfo) Reset() {
	*x = PersonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonInfo) ProtoMessage() {}

func (x *PersonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonInfo.ProtoReflect.Descriptor instead.
func (*PersonInfo) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{9}
}

func (x *PersonInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PersonInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonInfo) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *PersonInfo) GetPosition() []string {
	if x != nil {
		return x.Position
	}
	return nil
}

type PersonCompilation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*PersonInfo `protobuf:"bytes,1,rep,name=Persons,proto3" json:"Persons,omitempty"`
}

func (x *PersonCompilation) Reset() {
	*x = PersonCompilation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonCompilation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonCompilation) ProtoMessage() {}

func (x *PersonCompilation) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonCompilation.ProtoReflect.Descriptor instead.
func (*PersonCompilation) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{10}
}

func (x *PersonCompilation) GetPersons() []*PersonInfo {
	if x != nil {
		return x.Persons
	}
	return nil
}

type SearchCompilation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies  []*MovieInfo  `protobuf:"bytes,1,rep,name=Movies,proto3" json:"Movies,omitempty"`
	Series  []*MovieInfo  `protobuf:"bytes,2,rep,name=Series,proto3" json:"Series,omitempty"`
	Persons []*PersonInfo `protobuf:"bytes,3,rep,name=Persons,proto3" json:"Persons,omitempty"`
}

func (x *SearchCompilation) Reset() {
	*x = SearchCompilation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCompilation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCompilation) ProtoMessage() {}

func (x *SearchCompilation) ProtoReflect() protoreflect.Message {
	mi := &file_microservices_compilations_proto_compilations_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCompilation.ProtoReflect.Descriptor instead.
func (*SearchCompilation) Descriptor() ([]byte, []int) {
	return file_microservices_compilations_proto_compilations_proto_rawDescGZIP(), []int{11}
}

func (x *SearchCompilation) GetMovies() []*MovieInfo {
	if x != nil {
		return x.Movies
	}
	return nil
}

func (x *SearchCompilation) GetSeries() []*MovieInfo {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *SearchCompilation) GetPersons() []*PersonInfo {
	if x != nil {
		return x.Persons
	}
	return nil
}

var File_microservices_compilations_proto_compilations_proto protoreflect.FileDescriptor

var file_microservices_compilations_proto_compilations_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x05,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6d, 0x0a, 0x09, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x50, 0x0a, 0x10, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x14, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41,
	0x72, 0x72, 0x12, 0x45, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x25,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x22, 0x62, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x06, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x06, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x32,
	0xa1, 0x07, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x47,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x72, 0x72, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x42, 0x79, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x12,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x72,
	0x72, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x04,
	0x46, 0x69, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x78, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservices_compilations_proto_compilations_proto_rawDescOnce sync.Once
	file_microservices_compilations_proto_compilations_proto_rawDescData = file_microservices_compilations_proto_compilations_proto_rawDesc
)

func file_microservices_compilations_proto_compilations_proto_rawDescGZIP() []byte {
	file_microservices_compilations_proto_compilations_proto_rawDescOnce.Do(func() {
		file_microservices_compilations_proto_compilations_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservices_compilations_proto_compilations_proto_rawDescData)
	})
	return file_microservices_compilations_proto_compilations_proto_rawDescData
}

var file_microservices_compilations_proto_compilations_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_microservices_compilations_proto_compilations_proto_goTypes = []interface{}{
	(*Genre)(nil),                      // 0: proto.Genre
	(*MovieInfo)(nil),                  // 1: proto.MovieInfo
	(*MovieCompilation)(nil),           // 2: proto.MovieCompilation
	(*MovieCompilationsArr)(nil),       // 3: proto.MovieCompilationsArr
	(*GetMainCompilationsOptions)(nil), // 4: proto.GetMainCompilationsOptions
	(*GetByIDOptions)(nil),             // 5: proto.GetByIDOptions
	(*GetCompilationOptions)(nil),      // 6: proto.GetCompilationOptions
	(*GetFavoritesOptions)(nil),        // 7: proto.GetFavoritesOptions
	(*SearchText)(nil),                 // 8: proto.SearchText
	(*PersonInfo)(nil),                 // 9: proto.PersonInfo
	(*PersonCompilation)(nil),          // 10: proto.PersonCompilation
	(*SearchCompilation)(nil),          // 11: proto.SearchCompilation
}
var file_microservices_compilations_proto_compilations_proto_depIdxs = []int32{
	0,  // 0: proto.MovieInfo.Genre:type_name -> proto.Genre
	1,  // 1: proto.MovieCompilation.Movies:type_name -> proto.MovieInfo
	2,  // 2: proto.MovieCompilationsArr.MovieCompilations:type_name -> proto.MovieCompilation
	9,  // 3: proto.PersonCompilation.Persons:type_name -> proto.PersonInfo
	1,  // 4: proto.SearchCompilation.Movies:type_name -> proto.MovieInfo
	1,  // 5: proto.SearchCompilation.Series:type_name -> proto.MovieInfo
	9,  // 6: proto.SearchCompilation.Persons:type_name -> proto.PersonInfo
	6,  // 7: proto.MovieCompilations.GetAllMovies:input_type -> proto.GetCompilationOptions
	6,  // 8: proto.MovieCompilations.GetAllSeries:input_type -> proto.GetCompilationOptions
	4,  // 9: proto.MovieCompilations.GetMainCompilations:input_type -> proto.GetMainCompilationsOptions
	5,  // 10: proto.MovieCompilations.GetByGenre:input_type -> proto.GetByIDOptions
	5,  // 11: proto.MovieCompilations.GetByCountry:input_type -> proto.GetByIDOptions
	5,  // 12: proto.MovieCompilations.GetByMovie:input_type -> proto.GetByIDOptions
	5,  // 13: proto.MovieCompilations.GetByPerson:input_type -> proto.GetByIDOptions
	5,  // 14: proto.MovieCompilations.GetTopByYear:input_type -> proto.GetByIDOptions
	6,  // 15: proto.MovieCompilations.GetTop:input_type -> proto.GetCompilationOptions
	7,  // 16: proto.MovieCompilations.GetFavorites:input_type -> proto.GetFavoritesOptions
	7,  // 17: proto.MovieCompilations.GetFavoritesFilms:input_type -> proto.GetFavoritesOptions
	7,  // 18: proto.MovieCompilations.GetFavoritesSeries:input_type -> proto.GetFavoritesOptions
	8,  // 19: proto.MovieCompilations.Find:input_type -> proto.SearchText
	2,  // 20: proto.MovieCompilations.GetAllMovies:output_type -> proto.MovieCompilation
	2,  // 21: proto.MovieCompilations.GetAllSeries:output_type -> proto.MovieCompilation
	3,  // 22: proto.MovieCompilations.GetMainCompilations:output_type -> proto.MovieCompilationsArr
	2,  // 23: proto.MovieCompilations.GetByGenre:output_type -> proto.MovieCompilation
	2,  // 24: proto.MovieCompilations.GetByCountry:output_type -> proto.MovieCompilation
	2,  // 25: proto.MovieCompilations.GetByMovie:output_type -> proto.MovieCompilation
	2,  // 26: proto.MovieCompilations.GetByPerson:output_type -> proto.MovieCompilation
	2,  // 27: proto.MovieCompilations.GetTopByYear:output_type -> proto.MovieCompilation
	2,  // 28: proto.MovieCompilations.GetTop:output_type -> proto.MovieCompilation
	3,  // 29: proto.MovieCompilations.GetFavorites:output_type -> proto.MovieCompilationsArr
	2,  // 30: proto.MovieCompilations.GetFavoritesFilms:output_type -> proto.MovieCompilation
	2,  // 31: proto.MovieCompilations.GetFavoritesSeries:output_type -> proto.MovieCompilation
	11, // 32: proto.MovieCompilations.Find:output_type -> proto.SearchCompilation
	20, // [20:33] is the sub-list for method output_type
	7,  // [7:20] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_microservices_compilations_proto_compilations_proto_init() }
func file_microservices_compilations_proto_compilations_proto_init() {
	if File_microservices_compilations_proto_compilations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservices_compilations_proto_compilations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieInfo); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieCompilation); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieCompilationsArr); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMainCompilationsOptions); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDOptions); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompilationOptions); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFavoritesOptions); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchText); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonInfo); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonCompilation); i {
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
		file_microservices_compilations_proto_compilations_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCompilation); i {
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
			RawDescriptor: file_microservices_compilations_proto_compilations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_microservices_compilations_proto_compilations_proto_goTypes,
		DependencyIndexes: file_microservices_compilations_proto_compilations_proto_depIdxs,
		MessageInfos:      file_microservices_compilations_proto_compilations_proto_msgTypes,
	}.Build()
	File_microservices_compilations_proto_compilations_proto = out.File
	file_microservices_compilations_proto_compilations_proto_rawDesc = nil
	file_microservices_compilations_proto_compilations_proto_goTypes = nil
	file_microservices_compilations_proto_compilations_proto_depIdxs = nil
}
