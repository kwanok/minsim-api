// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: server/minsim/minsim.proto

package minsim

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

type NewMinsimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Positive  float32 `protobuf:"fixed32,2,opt,name=positive,proto3" json:"positive,omitempty"`
	Negative  float32 `protobuf:"fixed32,3,opt,name=negative,proto3" json:"negative,omitempty"`
	Neutral   float32 `protobuf:"fixed32,4,opt,name=neutral,proto3" json:"neutral,omitempty"`
	Content   string  `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	User      string  `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Url       string  `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	Id        string  `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int64   `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *NewMinsimRequest) Reset() {
	*x = NewMinsimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_minsim_minsim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMinsimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMinsimRequest) ProtoMessage() {}

func (x *NewMinsimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_minsim_minsim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMinsimRequest.ProtoReflect.Descriptor instead.
func (*NewMinsimRequest) Descriptor() ([]byte, []int) {
	return file_server_minsim_minsim_proto_rawDescGZIP(), []int{0}
}

func (x *NewMinsimRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NewMinsimRequest) GetPositive() float32 {
	if x != nil {
		return x.Positive
	}
	return 0
}

func (x *NewMinsimRequest) GetNegative() float32 {
	if x != nil {
		return x.Negative
	}
	return 0
}

func (x *NewMinsimRequest) GetNeutral() float32 {
	if x != nil {
		return x.Neutral
	}
	return 0
}

func (x *NewMinsimRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NewMinsimRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *NewMinsimRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NewMinsimRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewMinsimRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type NewMinsimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewMinsimResponse) Reset() {
	*x = NewMinsimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_minsim_minsim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMinsimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMinsimResponse) ProtoMessage() {}

func (x *NewMinsimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_minsim_minsim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMinsimResponse.ProtoReflect.Descriptor instead.
func (*NewMinsimResponse) Descriptor() ([]byte, []int) {
	return file_server_minsim_minsim_proto_rawDescGZIP(), []int{1}
}

func (x *NewMinsimResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_server_minsim_minsim_proto protoreflect.FileDescriptor

var file_server_minsim_minsim_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x2f,
	0x6d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x22, 0xe7, 0x01, 0x0a, 0x10,
	0x4e, 0x65, 0x77, 0x4d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6e,
	0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x4d, 0x69, 0x6e, 0x73,
	0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x56, 0x0a, 0x06, 0x4d, 0x69,
	0x6e, 0x73, 0x69, 0x6d, 0x12, 0x4c, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x2e,
	0x4e, 0x65, 0x77, 0x4d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x6d,
	0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x77, 0x61, 0x6e, 0x6f, 0x6b, 0x2f, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x6d, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x69, 0x6e, 0x73, 0x69, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_minsim_minsim_proto_rawDescOnce sync.Once
	file_server_minsim_minsim_proto_rawDescData = file_server_minsim_minsim_proto_rawDesc
)

func file_server_minsim_minsim_proto_rawDescGZIP() []byte {
	file_server_minsim_minsim_proto_rawDescOnce.Do(func() {
		file_server_minsim_minsim_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_minsim_minsim_proto_rawDescData)
	})
	return file_server_minsim_minsim_proto_rawDescData
}

var file_server_minsim_minsim_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_minsim_minsim_proto_goTypes = []interface{}{
	(*NewMinsimRequest)(nil),  // 0: server.minsim.NewMinsimRequest
	(*NewMinsimResponse)(nil), // 1: server.minsim.NewMinsimResponse
}
var file_server_minsim_minsim_proto_depIdxs = []int32{
	0, // 0: server.minsim.Minsim.NewPost:input_type -> server.minsim.NewMinsimRequest
	1, // 1: server.minsim.Minsim.NewPost:output_type -> server.minsim.NewMinsimResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_minsim_minsim_proto_init() }
func file_server_minsim_minsim_proto_init() {
	if File_server_minsim_minsim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_minsim_minsim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMinsimRequest); i {
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
		file_server_minsim_minsim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMinsimResponse); i {
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
			RawDescriptor: file_server_minsim_minsim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_minsim_minsim_proto_goTypes,
		DependencyIndexes: file_server_minsim_minsim_proto_depIdxs,
		MessageInfos:      file_server_minsim_minsim_proto_msgTypes,
	}.Build()
	File_server_minsim_minsim_proto = out.File
	file_server_minsim_minsim_proto_rawDesc = nil
	file_server_minsim_minsim_proto_goTypes = nil
	file_server_minsim_minsim_proto_depIdxs = nil
}