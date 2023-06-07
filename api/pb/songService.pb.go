// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: songService.proto

package pb

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

type RawSong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Init   int32  `protobuf:"varint,1,opt,name=Init,proto3" json:"Init,omitempty"`
	End    int32  `protobuf:"varint,2,opt,name=End,proto3" json:"End,omitempty"`
	Buffer []byte `protobuf:"bytes,3,opt,name=Buffer,proto3" json:"Buffer,omitempty"`
}

func (x *RawSong) Reset() {
	*x = RawSong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawSong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawSong) ProtoMessage() {}

func (x *RawSong) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawSong.ProtoReflect.Descriptor instead.
func (*RawSong) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{0}
}

func (x *RawSong) GetInit() int32 {
	if x != nil {
		return x.Init
	}
	return 0
}

func (x *RawSong) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *RawSong) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

type SongMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album       *string `protobuf:"bytes,1,opt,name=Album,proto3,oneof" json:"Album,omitempty"`
	AlbumArtist *string `protobuf:"bytes,2,opt,name=AlbumArtist,proto3,oneof" json:"AlbumArtist,omitempty"`
	Artist      *string `protobuf:"bytes,3,opt,name=Artist,proto3,oneof" json:"Artist,omitempty"`
	Comment     *string `protobuf:"bytes,4,opt,name=Comment,proto3,oneof" json:"Comment,omitempty"`
	Composer    *string `protobuf:"bytes,5,opt,name=Composer,proto3,oneof" json:"Composer,omitempty"`
	FileType    *string `protobuf:"bytes,6,opt,name=FileType,proto3,oneof" json:"FileType,omitempty"`
	Format      *string `protobuf:"bytes,7,opt,name=Format,proto3,oneof" json:"Format,omitempty"`
	Genre       *string `protobuf:"bytes,8,opt,name=Genre,proto3,oneof" json:"Genre,omitempty"`
	Lyrics      *string `protobuf:"bytes,10,opt,name=Lyrics,proto3,oneof" json:"Lyrics,omitempty"`
	Title       *string `protobuf:"bytes,11,opt,name=Title,proto3,oneof" json:"Title,omitempty"`
	Year        *int32  `protobuf:"varint,12,opt,name=Year,proto3,oneof" json:"Year,omitempty"`
}

func (x *SongMetadata) Reset() {
	*x = SongMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongMetadata) ProtoMessage() {}

func (x *SongMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongMetadata.ProtoReflect.Descriptor instead.
func (*SongMetadata) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{1}
}

func (x *SongMetadata) GetAlbum() string {
	if x != nil && x.Album != nil {
		return *x.Album
	}
	return ""
}

func (x *SongMetadata) GetAlbumArtist() string {
	if x != nil && x.AlbumArtist != nil {
		return *x.AlbumArtist
	}
	return ""
}

func (x *SongMetadata) GetArtist() string {
	if x != nil && x.Artist != nil {
		return *x.Artist
	}
	return ""
}

func (x *SongMetadata) GetComment() string {
	if x != nil && x.Comment != nil {
		return *x.Comment
	}
	return ""
}

func (x *SongMetadata) GetComposer() string {
	if x != nil && x.Composer != nil {
		return *x.Composer
	}
	return ""
}

func (x *SongMetadata) GetFileType() string {
	if x != nil && x.FileType != nil {
		return *x.FileType
	}
	return ""
}

func (x *SongMetadata) GetFormat() string {
	if x != nil && x.Format != nil {
		return *x.Format
	}
	return ""
}

func (x *SongMetadata) GetGenre() string {
	if x != nil && x.Genre != nil {
		return *x.Genre
	}
	return ""
}

func (x *SongMetadata) GetLyrics() string {
	if x != nil && x.Lyrics != nil {
		return *x.Lyrics
	}
	return ""
}

func (x *SongMetadata) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *SongMetadata) GetYear() int32 {
	if x != nil && x.Year != nil {
		return *x.Year
	}
	return 0
}

type SongId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *SongId) Reset() {
	*x = SongId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongId) ProtoMessage() {}

func (x *SongId) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongId.ProtoReflect.Descriptor instead.
func (*SongId) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{2}
}

func (x *SongId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *SongId       `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	RawSong  *RawSong      `protobuf:"bytes,2,opt,name=RawSong,proto3" json:"RawSong,omitempty"`
	Metadata *SongMetadata `protobuf:"bytes,3,opt,name=Metadata,proto3,oneof" json:"Metadata,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{3}
}

func (x *Song) GetId() *SongId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Song) GetRawSong() *RawSong {
	if x != nil {
		return x.RawSong
	}
	return nil
}

func (x *Song) GetMetadata() *SongMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type UpdatedSong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *SongId       `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Metadata *SongMetadata `protobuf:"bytes,3,opt,name=Metadata,proto3,oneof" json:"Metadata,omitempty"`
}

func (x *UpdatedSong) Reset() {
	*x = UpdatedSong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedSong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedSong) ProtoMessage() {}

func (x *UpdatedSong) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedSong.ProtoReflect.Descriptor instead.
func (*UpdatedSong) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatedSong) GetId() *SongId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdatedSong) GetMetadata() *SongMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_songService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_songService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_songService_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_songService_proto protoreflect.FileDescriptor

var file_songService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x6f, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x22, 0x47, 0x0a, 0x07, 0x52, 0x61,
	0x77, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x42,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x22, 0xd5, 0x03, 0x0a, 0x0c, 0x53, 0x6f, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x0b, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x41, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x07, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x06,
	0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x0a, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x59, 0x65, 0x61, 0x72, 0x22, 0x18, 0x0a, 0x06, 0x53,
	0x6f, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1d,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x6f, 0x6e,
	0x67, 0x73, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x02, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x07, 0x52, 0x61, 0x77, 0x53, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x61, 0x77, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x07,
	0x52, 0x61, 0x77, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x6f, 0x6e, 0x67,
	0x73, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6f, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x6f,
	0x6e, 0x67, 0x49, 0x64, 0x52, 0x02, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x6f, 0x6e,
	0x67, 0x73, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x8b, 0x02, 0x0a, 0x0b, 0x53, 0x6f, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x31, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12,
	0x0e, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x61, 0x77, 0x53, 0x6f, 0x6e, 0x67, 0x1a,
	0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x33, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f,
	0x6e, 0x67, 0x12, 0x12, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x1a, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x53, 0x6f, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73,
	0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e,
	0x53, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x53, 0x6f, 0x6e, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x6f,
	0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0b, 0x2e, 0x73, 0x6f, 0x6e,
	0x67, 0x73, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x32, 0x0a, 0x0e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e,
	0x73, 0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x73,
	0x6f, 0x6e, 0x67, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x2d, 0x61, 0x72, 0x74, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_songService_proto_rawDescOnce sync.Once
	file_songService_proto_rawDescData = file_songService_proto_rawDesc
)

func file_songService_proto_rawDescGZIP() []byte {
	file_songService_proto_rawDescOnce.Do(func() {
		file_songService_proto_rawDescData = protoimpl.X.CompressGZIP(file_songService_proto_rawDescData)
	})
	return file_songService_proto_rawDescData
}

var file_songService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_songService_proto_goTypes = []interface{}{
	(*RawSong)(nil),      // 0: songs.RawSong
	(*SongMetadata)(nil), // 1: songs.SongMetadata
	(*SongId)(nil),       // 2: songs.SongId
	(*Song)(nil),         // 3: songs.Song
	(*UpdatedSong)(nil),  // 4: songs.UpdatedSong
	(*Response)(nil),     // 5: songs.Response
}
var file_songService_proto_depIdxs = []int32{
	2,  // 0: songs.Song.Id:type_name -> songs.SongId
	0,  // 1: songs.Song.RawSong:type_name -> songs.RawSong
	1,  // 2: songs.Song.Metadata:type_name -> songs.SongMetadata
	2,  // 3: songs.UpdatedSong.Id:type_name -> songs.SongId
	1,  // 4: songs.UpdatedSong.Metadata:type_name -> songs.SongMetadata
	0,  // 5: songs.SongService.CreateSong:input_type -> songs.RawSong
	4,  // 6: songs.SongService.UpdateSong:input_type -> songs.UpdatedSong
	2,  // 7: songs.SongService.GetSongById:input_type -> songs.SongId
	1,  // 8: songs.SongService.SongFilter:input_type -> songs.SongMetadata
	2,  // 9: songs.SongService.RemoveSongById:input_type -> songs.SongId
	5,  // 10: songs.SongService.CreateSong:output_type -> songs.Response
	5,  // 11: songs.SongService.UpdateSong:output_type -> songs.Response
	3,  // 12: songs.SongService.GetSongById:output_type -> songs.Song
	3,  // 13: songs.SongService.SongFilter:output_type -> songs.Song
	5,  // 14: songs.SongService.RemoveSongById:output_type -> songs.Response
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_songService_proto_init() }
func file_songService_proto_init() {
	if File_songService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_songService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawSong); i {
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
		file_songService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongMetadata); i {
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
		file_songService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongId); i {
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
		file_songService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Song); i {
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
		file_songService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedSong); i {
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
		file_songService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_songService_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_songService_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_songService_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_songService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_songService_proto_goTypes,
		DependencyIndexes: file_songService_proto_depIdxs,
		MessageInfos:      file_songService_proto_msgTypes,
	}.Build()
	File_songService_proto = out.File
	file_songService_proto_rawDesc = nil
	file_songService_proto_goTypes = nil
	file_songService_proto_depIdxs = nil
}
