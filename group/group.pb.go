// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: group/group.proto

package group

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Property   int32              `protobuf:"varint,2,opt,name=property,proto3" json:"property,omitempty"`
	Additional []*UnrepeatedGroup `protobuf:"bytes,3,rep,name=additional,proto3" json:"additional,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetProperty() int32 {
	if x != nil {
		return x.Property
	}
	return 0
}

func (x *Group) GetAdditional() []*UnrepeatedGroup {
	if x != nil {
		return x.Additional
	}
	return nil
}

type UnrepeatedGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Property int32  `protobuf:"varint,2,opt,name=property,proto3" json:"property,omitempty"`
}

func (x *UnrepeatedGroup) Reset() {
	*x = UnrepeatedGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnrepeatedGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnrepeatedGroup) ProtoMessage() {}

func (x *UnrepeatedGroup) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnrepeatedGroup.ProtoReflect.Descriptor instead.
func (*UnrepeatedGroup) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{1}
}

func (x *UnrepeatedGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UnrepeatedGroup) GetProperty() int32 {
	if x != nil {
		return x.Property
	}
	return 0
}

var file_group_group_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Group)(nil),
		Field:         1101,
		Name:          "google.api.group",
		Tag:           "bytes,1101,opt,name=group",
		Filename:      "group/group.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional google.api.Group group = 1101;
	E_Group = &file_group_group_proto_extTypes[0]
)

var File_group_group_proto protoreflect.FileDescriptor

var file_group_group_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x74, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x6e, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x41, 0x0a, 0x0f, 0x55, 0x6e, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x3a, 0x48, 0x0a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xcd, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x67, 0x64, 0x61, 0x63, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3b, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_group_proto_rawDescOnce sync.Once
	file_group_group_proto_rawDescData = file_group_group_proto_rawDesc
)

func file_group_group_proto_rawDescGZIP() []byte {
	file_group_group_proto_rawDescOnce.Do(func() {
		file_group_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_group_proto_rawDescData)
	})
	return file_group_group_proto_rawDescData
}

var file_group_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_group_group_proto_goTypes = []interface{}{
	(*Group)(nil),                      // 0: google.api.Group
	(*UnrepeatedGroup)(nil),            // 1: google.api.UnrepeatedGroup
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_group_group_proto_depIdxs = []int32{
	1, // 0: google.api.Group.additional:type_name -> google.api.UnrepeatedGroup
	2, // 1: google.api.group:extendee -> google.protobuf.MethodOptions
	0, // 2: google.api.group:type_name -> google.api.Group
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_group_group_proto_init() }
func file_group_group_proto_init() {
	if File_group_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_group_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnrepeatedGroup); i {
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
			RawDescriptor: file_group_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_group_group_proto_goTypes,
		DependencyIndexes: file_group_group_proto_depIdxs,
		MessageInfos:      file_group_group_proto_msgTypes,
		ExtensionInfos:    file_group_group_proto_extTypes,
	}.Build()
	File_group_group_proto = out.File
	file_group_group_proto_rawDesc = nil
	file_group_group_proto_goTypes = nil
	file_group_group_proto_depIdxs = nil
}
