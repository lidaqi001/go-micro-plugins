// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: basic/event/event_hello/demo.proto

package event_hello

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

type HelloEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Who string `protobuf:"bytes,1,opt,name=who,proto3" json:"who,omitempty"`
}

func (x *HelloEvent) Reset() {
	*x = HelloEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_event_event_hello_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloEvent) ProtoMessage() {}

func (x *HelloEvent) ProtoReflect() protoreflect.Message {
	mi := &file_basic_event_event_hello_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloEvent.ProtoReflect.Descriptor instead.
func (*HelloEvent) Descriptor() ([]byte, []int) {
	return file_basic_event_event_hello_demo_proto_rawDescGZIP(), []int{0}
}

func (x *HelloEvent) GetWho() string {
	if x != nil {
		return x.Who
	}
	return ""
}

type BaseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseInt int32            `protobuf:"varint,1,opt,name=baseInt,proto3" json:"baseInt,omitempty"`
	Data    map[int32]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BaseMessage) Reset() {
	*x = BaseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_event_event_hello_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseMessage) ProtoMessage() {}

func (x *BaseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_basic_event_event_hello_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseMessage.ProtoReflect.Descriptor instead.
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return file_basic_event_event_hello_demo_proto_rawDescGZIP(), []int{1}
}

func (x *BaseMessage) GetBaseInt() int32 {
	if x != nil {
		return x.BaseInt
	}
	return 0
}

func (x *BaseMessage) GetData() map[int32]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_basic_event_event_hello_demo_proto protoreflect.FileDescriptor

var file_basic_event_event_hello_demo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x22, 0x1e, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x77, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x68,
	0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3d, 0x5a, 0x3b,
	0x63, 0x6f, 0x64, 0x65, 0x75, 0x70, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x68, 0x61, 0x6e, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_basic_event_event_hello_demo_proto_rawDescOnce sync.Once
	file_basic_event_event_hello_demo_proto_rawDescData = file_basic_event_event_hello_demo_proto_rawDesc
)

func file_basic_event_event_hello_demo_proto_rawDescGZIP() []byte {
	file_basic_event_event_hello_demo_proto_rawDescOnce.Do(func() {
		file_basic_event_event_hello_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_event_event_hello_demo_proto_rawDescData)
	})
	return file_basic_event_event_hello_demo_proto_rawDescData
}

var file_basic_event_event_hello_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_basic_event_event_hello_demo_proto_goTypes = []interface{}{
	(*HelloEvent)(nil),  // 0: event_hello.HelloEvent
	(*BaseMessage)(nil), // 1: event_hello.BaseMessage
	nil,                 // 2: event_hello.BaseMessage.DataEntry
}
var file_basic_event_event_hello_demo_proto_depIdxs = []int32{
	2, // 0: event_hello.BaseMessage.data:type_name -> event_hello.BaseMessage.DataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_basic_event_event_hello_demo_proto_init() }
func file_basic_event_event_hello_demo_proto_init() {
	if File_basic_event_event_hello_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basic_event_event_hello_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloEvent); i {
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
		file_basic_event_event_hello_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseMessage); i {
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
			RawDescriptor: file_basic_event_event_hello_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basic_event_event_hello_demo_proto_goTypes,
		DependencyIndexes: file_basic_event_event_hello_demo_proto_depIdxs,
		MessageInfos:      file_basic_event_event_hello_demo_proto_msgTypes,
	}.Build()
	File_basic_event_event_hello_demo_proto = out.File
	file_basic_event_event_hello_demo_proto_rawDesc = nil
	file_basic_event_event_hello_demo_proto_goTypes = nil
	file_basic_event_event_hello_demo_proto_depIdxs = nil
}
