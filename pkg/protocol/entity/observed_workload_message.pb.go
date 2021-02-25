// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/protoc/entity/observed_workload_message.proto

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

// ObservedWorkload is entity with interesting workload specification, from Runner point of view. Listener watches
// market, and if workload matching our capabilities was emitted, the ObserverWorkload is persisted and watched
// for future request.
type ObservedWorkload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required,uuid"
	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validate:"required,uuid"`
	State *State `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	// @inject_tag: validate:"required"
	WorkloadSpecificationMarketId string `protobuf:"bytes,3,opt,name=workloadSpecificationMarketId,proto3" json:"workloadSpecificationMarketId,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	Spec *WorkloadSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	PrincipalProposal   *ContractProposal                     `protobuf:"bytes,5,opt,name=principalProposal,proto3" json:"principalProposal,omitempty" validate:"required"`
	PrincipalAcceptance *ObservedWorkload_PrincipalAcceptance `protobuf:"bytes,6,opt,name=principalAcceptance,proto3" json:"principalAcceptance,omitempty"`
}

func (x *ObservedWorkload) Reset() {
	*x = ObservedWorkload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_entity_observed_workload_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservedWorkload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservedWorkload) ProtoMessage() {}

func (x *ObservedWorkload) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_entity_observed_workload_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservedWorkload.ProtoReflect.Descriptor instead.
func (*ObservedWorkload) Descriptor() ([]byte, []int) {
	return file_api_protoc_entity_observed_workload_message_proto_rawDescGZIP(), []int{0}
}

func (x *ObservedWorkload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ObservedWorkload) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *ObservedWorkload) GetWorkloadSpecificationMarketId() string {
	if x != nil {
		return x.WorkloadSpecificationMarketId
	}
	return ""
}

func (x *ObservedWorkload) GetSpec() *WorkloadSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ObservedWorkload) GetPrincipalProposal() *ContractProposal {
	if x != nil {
		return x.PrincipalProposal
	}
	return nil
}

func (x *ObservedWorkload) GetPrincipalAcceptance() *ObservedWorkload_PrincipalAcceptance {
	if x != nil {
		return x.PrincipalAcceptance
	}
	return nil
}

type ObservedWorkload_PrincipalAcceptance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: validate:"required"
	ContractAcceptMarketId string `protobuf:"bytes,1,opt,name=contractAcceptMarketId,proto3" json:"contractAcceptMarketId,omitempty" validate:"required"`
	// @inject_tag: validate:"required"
	Accept *ContractAccept `protobuf:"bytes,2,opt,name=accept,proto3" json:"accept,omitempty" validate:"required"`
}

func (x *ObservedWorkload_PrincipalAcceptance) Reset() {
	*x = ObservedWorkload_PrincipalAcceptance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protoc_entity_observed_workload_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservedWorkload_PrincipalAcceptance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservedWorkload_PrincipalAcceptance) ProtoMessage() {}

func (x *ObservedWorkload_PrincipalAcceptance) ProtoReflect() protoreflect.Message {
	mi := &file_api_protoc_entity_observed_workload_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservedWorkload_PrincipalAcceptance.ProtoReflect.Descriptor instead.
func (*ObservedWorkload_PrincipalAcceptance) Descriptor() ([]byte, []int) {
	return file_api_protoc_entity_observed_workload_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ObservedWorkload_PrincipalAcceptance) GetContractAcceptMarketId() string {
	if x != nil {
		return x.ContractAcceptMarketId
	}
	return ""
}

func (x *ObservedWorkload_PrincipalAcceptance) GetAccept() *ContractAccept {
	if x != nil {
		return x.Accept
	}
	return nil
}

var File_api_protoc_entity_observed_workload_message_proto protoreflect.FileDescriptor

var file_api_protoc_entity_observed_workload_message_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x69, 0x6f, 0x2e, 0x6e, 0x61, 0x66, 0x74, 0x69, 0x73, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x91, 0x04, 0x0a, 0x10, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x69, 0x6f, 0x2e, 0x6e, 0x61, 0x66, 0x74, 0x69, 0x73, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x1d, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6f, 0x2e, 0x6e, 0x61, 0x66,
	0x74, 0x69, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x50, 0x0a,
	0x11, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6f, 0x2e, 0x6e, 0x61,
	0x66, 0x74, 0x69, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x11, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12,
	0x68, 0x0a, 0x13, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x69,
	0x6f, 0x2e, 0x6e, 0x61, 0x66, 0x74, 0x69, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x13, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x87, 0x01, 0x0a, 0x13, 0x50, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x36, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6f, 0x2e, 0x6e,
	0x61, 0x66, 0x74, 0x69, 0x73, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x52, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x61, 0x66, 0x74, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6e, 0x61, 0x66,
	0x74, 0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protoc_entity_observed_workload_message_proto_rawDescOnce sync.Once
	file_api_protoc_entity_observed_workload_message_proto_rawDescData = file_api_protoc_entity_observed_workload_message_proto_rawDesc
)

func file_api_protoc_entity_observed_workload_message_proto_rawDescGZIP() []byte {
	file_api_protoc_entity_observed_workload_message_proto_rawDescOnce.Do(func() {
		file_api_protoc_entity_observed_workload_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protoc_entity_observed_workload_message_proto_rawDescData)
	})
	return file_api_protoc_entity_observed_workload_message_proto_rawDescData
}

var file_api_protoc_entity_observed_workload_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_protoc_entity_observed_workload_message_proto_goTypes = []interface{}{
	(*ObservedWorkload)(nil),                     // 0: io.naftis.entity.ObservedWorkload
	(*ObservedWorkload_PrincipalAcceptance)(nil), // 1: io.naftis.entity.ObservedWorkload.PrincipalAcceptance
	(*State)(nil),                                // 2: io.naftis.entity.State
	(*WorkloadSpec)(nil),                         // 3: io.naftis.entity.WorkloadSpec
	(*ContractProposal)(nil),                     // 4: io.naftis.entity.ContractProposal
	(*ContractAccept)(nil),                       // 5: io.naftis.entity.ContractAccept
}
var file_api_protoc_entity_observed_workload_message_proto_depIdxs = []int32{
	2, // 0: io.naftis.entity.ObservedWorkload.state:type_name -> io.naftis.entity.State
	3, // 1: io.naftis.entity.ObservedWorkload.spec:type_name -> io.naftis.entity.WorkloadSpec
	4, // 2: io.naftis.entity.ObservedWorkload.principalProposal:type_name -> io.naftis.entity.ContractProposal
	1, // 3: io.naftis.entity.ObservedWorkload.principalAcceptance:type_name -> io.naftis.entity.ObservedWorkload.PrincipalAcceptance
	5, // 4: io.naftis.entity.ObservedWorkload.PrincipalAcceptance.accept:type_name -> io.naftis.entity.ContractAccept
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_protoc_entity_observed_workload_message_proto_init() }
func file_api_protoc_entity_observed_workload_message_proto_init() {
	if File_api_protoc_entity_observed_workload_message_proto != nil {
		return
	}
	file_api_protoc_entity_workload_spec_message_proto_init()
	file_api_protoc_entity_contract_proposal_message_proto_init()
	file_api_protoc_entity_contract_accept_message_proto_init()
	file_api_protoc_entity_state_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_protoc_entity_observed_workload_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservedWorkload); i {
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
		file_api_protoc_entity_observed_workload_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservedWorkload_PrincipalAcceptance); i {
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
			RawDescriptor: file_api_protoc_entity_observed_workload_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_protoc_entity_observed_workload_message_proto_goTypes,
		DependencyIndexes: file_api_protoc_entity_observed_workload_message_proto_depIdxs,
		MessageInfos:      file_api_protoc_entity_observed_workload_message_proto_msgTypes,
	}.Build()
	File_api_protoc_entity_observed_workload_message_proto = out.File
	file_api_protoc_entity_observed_workload_message_proto_rawDesc = nil
	file_api_protoc_entity_observed_workload_message_proto_goTypes = nil
	file_api_protoc_entity_observed_workload_message_proto_depIdxs = nil
}
