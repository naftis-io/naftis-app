// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/protoc/entity/node_selector_message.proto

package entity

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

// NodeSelector is entity that contains node label and constraint enforcement.
type NodeSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	Label *NodeLabel `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	IsConstraint bool `protobuf:"varint,2,opt,name=isConstraint,proto3" json:"isConstraint,omitempty" validate:"required"`
}

func (x *NodeSelector) Reset() {
	*x = NodeSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_entity_node_selector_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSelector) ProtoMessage() {}

func (x *NodeSelector) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_entity_node_selector_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSelector.ProtoReflect.Descriptor instead.
func (*NodeSelector) Descriptor() ([]byte, []int) {
	return file_api_protoc_entity_node_selector_message_proto_rawDescGZIP(), []int{0}
}

func (x *NodeSelector) GetLabel() *NodeLabel {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *NodeSelector) GetIsConstraint() bool {
	if x != nil {
		return x.IsConstraint
	}
	return false
}

var File_api_protoc_entity_node_selector_message_proto protoreflect.FileDescriptor

var file_api_protoc_entity_node_selector_message_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x69, 0x6f, 0x2e, 0x6e, 0x61, 0x66, 0x74, 0x69, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x1a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a,
	0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x31, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69,
	0x6f, 0x2e, 0x6e, 0x61, 0x66, 0x74, 0x69, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x74, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x66, 0x74, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6e, 0x61,
	0x66, 0x74, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protoc_entity_node_selector_message_proto_rawDescOnce sync.Once
	file_api_protoc_entity_node_selector_message_proto_rawDescData = file_api_protoc_entity_node_selector_message_proto_rawDesc
)

func file_api_protoc_entity_node_selector_message_proto_rawDescGZIP() []byte {
	file_api_protoc_entity_node_selector_message_proto_rawDescOnce.Do(func() {
		file_api_protoc_entity_node_selector_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protoc_entity_node_selector_message_proto_rawDescData)
	})
	return file_api_protoc_entity_node_selector_message_proto_rawDescData
}

var file_api_protoc_entity_node_selector_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_protoc_entity_node_selector_message_proto_goTypes = []interface{}{
	(*NodeSelector)(nil), // 0: io.naftis.entity.NodeSelector
	(*NodeLabel)(nil),    // 1: io.naftis.entity.NodeLabel
}
var file_api_protoc_entity_node_selector_message_proto_depIdxs = []int32{
	1, // 0: io.naftis.entity.NodeSelector.label:type_name -> io.naftis.entity.NodeLabel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_protoc_entity_node_selector_message_proto_init() }
func file_api_protoc_entity_node_selector_message_proto_init() {
	if File_api_protoc_entity_node_selector_message_proto != nil {
		return
	}
	file_api_protoc_entity_node_label_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_protoc_entity_node_selector_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSelector); i {
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
			RawDescriptor: file_api_protoc_entity_node_selector_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_protoc_entity_node_selector_message_proto_goTypes,
		DependencyIndexes: file_api_protoc_entity_node_selector_message_proto_depIdxs,
		MessageInfos:      file_api_protoc_entity_node_selector_message_proto_msgTypes,
	}.Build()
	File_api_protoc_entity_node_selector_message_proto = out.File
	file_api_protoc_entity_node_selector_message_proto_rawDesc = nil
	file_api_protoc_entity_node_selector_message_proto_goTypes = nil
	file_api_protoc_entity_node_selector_message_proto_depIdxs = nil
}
