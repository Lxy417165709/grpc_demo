// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: common/pb/email.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SendEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To      string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Subject string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *SendEmailReq) Reset() {
	*x = SendEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailReq) ProtoMessage() {}

func (x *SendEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailReq.ProtoReflect.Descriptor instead.
func (*SendEmailReq) Descriptor() ([]byte, []int) {
	return file_common_pb_email_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailReq) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendEmailReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendEmailReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type SendEmailRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendEmailRes) Reset() {
	*x = SendEmailRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_pb_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRes) ProtoMessage() {}

func (x *SendEmailRes) ProtoReflect() protoreflect.Message {
	mi := &file_common_pb_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRes.ProtoReflect.Descriptor instead.
func (*SendEmailRes) Descriptor() ([]byte, []int) {
	return file_common_pb_email_proto_rawDescGZIP(), []int{1}
}

var File_common_pb_email_proto protoreflect.FileDescriptor

var file_common_pb_email_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x52, 0x0a, 0x0c, 0x53,
	0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x32,
	0x33, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_pb_email_proto_rawDescOnce sync.Once
	file_common_pb_email_proto_rawDescData = file_common_pb_email_proto_rawDesc
)

func file_common_pb_email_proto_rawDescGZIP() []byte {
	file_common_pb_email_proto_rawDescOnce.Do(func() {
		file_common_pb_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_pb_email_proto_rawDescData)
	})
	return file_common_pb_email_proto_rawDescData
}

var file_common_pb_email_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_pb_email_proto_goTypes = []interface{}{
	(*SendEmailReq)(nil), // 0: pb.SendEmailReq
	(*SendEmailRes)(nil), // 1: pb.SendEmailRes
}
var file_common_pb_email_proto_depIdxs = []int32{
	0, // 0: pb.Email.Send:input_type -> pb.SendEmailReq
	1, // 1: pb.Email.Send:output_type -> pb.SendEmailRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_pb_email_proto_init() }
func file_common_pb_email_proto_init() {
	if File_common_pb_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_pb_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailReq); i {
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
		file_common_pb_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRes); i {
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
			RawDescriptor: file_common_pb_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_pb_email_proto_goTypes,
		DependencyIndexes: file_common_pb_email_proto_depIdxs,
		MessageInfos:      file_common_pb_email_proto_msgTypes,
	}.Build()
	File_common_pb_email_proto = out.File
	file_common_pb_email_proto_rawDesc = nil
	file_common_pb_email_proto_goTypes = nil
	file_common_pb_email_proto_depIdxs = nil
}
