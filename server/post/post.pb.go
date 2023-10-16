// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: server/post/post.proto

package post

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

type NewPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Author    string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Subreddit string `protobuf:"bytes,4,opt,name=subreddit,proto3" json:"subreddit,omitempty"`
	CreatedAt int64  `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Url       string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Id        string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewPostRequest) Reset() {
	*x = NewPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPostRequest) ProtoMessage() {}

func (x *NewPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPostRequest.ProtoReflect.Descriptor instead.
func (*NewPostRequest) Descriptor() ([]byte, []int) {
	return file_server_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *NewPostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewPostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NewPostRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *NewPostRequest) GetSubreddit() string {
	if x != nil {
		return x.Subreddit
	}
	return ""
}

func (x *NewPostRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NewPostRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NewPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NewPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewPostResponse) Reset() {
	*x = NewPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_post_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPostResponse) ProtoMessage() {}

func (x *NewPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_post_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPostResponse.ProtoReflect.Descriptor instead.
func (*NewPostResponse) Descriptor() ([]byte, []int) {
	return file_server_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *NewPostResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_server_post_post_proto protoreflect.FileDescriptor

var file_server_post_post_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x22, 0xb7, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x21, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0x4c, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x07, 0x4e, 0x65,
	0x77, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x77, 0x61, 0x6e, 0x6f, 0x6b, 0x2f, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_post_post_proto_rawDescOnce sync.Once
	file_server_post_post_proto_rawDescData = file_server_post_post_proto_rawDesc
)

func file_server_post_post_proto_rawDescGZIP() []byte {
	file_server_post_post_proto_rawDescOnce.Do(func() {
		file_server_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_post_post_proto_rawDescData)
	})
	return file_server_post_post_proto_rawDescData
}

var file_server_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_post_post_proto_goTypes = []interface{}{
	(*NewPostRequest)(nil),  // 0: server.post.NewPostRequest
	(*NewPostResponse)(nil), // 1: server.post.NewPostResponse
}
var file_server_post_post_proto_depIdxs = []int32{
	0, // 0: server.post.Post.NewPost:input_type -> server.post.NewPostRequest
	1, // 1: server.post.Post.NewPost:output_type -> server.post.NewPostResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_post_post_proto_init() }
func file_server_post_post_proto_init() {
	if File_server_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPostRequest); i {
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
		file_server_post_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPostResponse); i {
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
			RawDescriptor: file_server_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_post_post_proto_goTypes,
		DependencyIndexes: file_server_post_post_proto_depIdxs,
		MessageInfos:      file_server_post_post_proto_msgTypes,
	}.Build()
	File_server_post_post_proto = out.File
	file_server_post_post_proto_rawDesc = nil
	file_server_post_post_proto_goTypes = nil
	file_server_post_post_proto_depIdxs = nil
}
