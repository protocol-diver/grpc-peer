// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: msg/msg.proto

package msg

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

type MessageSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   int32  `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver int32  `protobuf:"varint,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageSendRequest) Reset() {
	*x = MessageSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendRequest) ProtoMessage() {}

func (x *MessageSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendRequest.ProtoReflect.Descriptor instead.
func (*MessageSendRequest) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{0}
}

func (x *MessageSendRequest) GetSender() int32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *MessageSendRequest) GetReceiver() int32 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

func (x *MessageSendRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender   int32  `protobuf:"varint,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver int32  `protobuf:"varint,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Message  string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Id       int32  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Error    string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MessageSendResponse) Reset() {
	*x = MessageSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendResponse) ProtoMessage() {}

func (x *MessageSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_msg_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendResponse.ProtoReflect.Descriptor instead.
func (*MessageSendResponse) Descriptor() ([]byte, []int) {
	return file_msg_msg_proto_rawDescGZIP(), []int{1}
}

func (x *MessageSendResponse) GetSender() int32 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *MessageSendResponse) GetReceiver() int32 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

func (x *MessageSendResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MessageSendResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageSendResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_msg_msg_proto protoreflect.FileDescriptor

var file_msg_msg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x62, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x32, 0x4b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x40, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x17,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x64, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_msg_proto_rawDescOnce sync.Once
	file_msg_msg_proto_rawDescData = file_msg_msg_proto_rawDesc
)

func file_msg_msg_proto_rawDescGZIP() []byte {
	file_msg_msg_proto_rawDescOnce.Do(func() {
		file_msg_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_msg_proto_rawDescData)
	})
	return file_msg_msg_proto_rawDescData
}

var file_msg_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msg_msg_proto_goTypes = []interface{}{
	(*MessageSendRequest)(nil),  // 0: msg.MessageSendRequest
	(*MessageSendResponse)(nil), // 1: msg.MessageSendResponse
}
var file_msg_msg_proto_depIdxs = []int32{
	0, // 0: msg.Message.MessageSend:input_type -> msg.MessageSendRequest
	1, // 1: msg.Message.MessageSend:output_type -> msg.MessageSendResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_msg_proto_init() }
func file_msg_msg_proto_init() {
	if File_msg_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSendRequest); i {
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
		file_msg_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSendResponse); i {
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
			RawDescriptor: file_msg_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_msg_proto_goTypes,
		DependencyIndexes: file_msg_msg_proto_depIdxs,
		MessageInfos:      file_msg_msg_proto_msgTypes,
	}.Build()
	File_msg_msg_proto = out.File
	file_msg_msg_proto_rawDesc = nil
	file_msg_msg_proto_goTypes = nil
	file_msg_msg_proto_depIdxs = nil
}
