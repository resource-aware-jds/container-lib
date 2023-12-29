// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/container.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TaskAttributes []byte `protobuf:"bytes,2,opt,name=TaskAttributes,proto3" json:"TaskAttributes,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_container_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_container_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Task) GetTaskAttributes() []byte {
	if x != nil {
		return x.TaskAttributes
	}
	return nil
}

var File_proto_container_proto protoreflect.FileDescriptor

var file_proto_container_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3e, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x32,
	0x4c, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x77, 0x61, 0x72, 0x65, 0x2d, 0x6a, 0x64, 0x73, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_container_proto_rawDescOnce sync.Once
	file_proto_container_proto_rawDescData = file_proto_container_proto_rawDesc
)

func file_proto_container_proto_rawDescGZIP() []byte {
	file_proto_container_proto_rawDescOnce.Do(func() {
		file_proto_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_container_proto_rawDescData)
	})
	return file_proto_container_proto_rawDescData
}

var file_proto_container_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_container_proto_goTypes = []interface{}{
	(*Task)(nil),          // 0: Container.Task
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_proto_container_proto_depIdxs = []int32{
	0, // 0: Container.ContainerTaskRunner.SendTask:input_type -> Container.Task
	1, // 1: Container.ContainerTaskRunner.SendTask:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_container_proto_init() }
func file_proto_container_proto_init() {
	if File_proto_container_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_container_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
			RawDescriptor: file_proto_container_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_container_proto_goTypes,
		DependencyIndexes: file_proto_container_proto_depIdxs,
		MessageInfos:      file_proto_container_proto_msgTypes,
	}.Build()
	File_proto_container_proto = out.File
	file_proto_container_proto_rawDesc = nil
	file_proto_container_proto_goTypes = nil
	file_proto_container_proto_depIdxs = nil
}
