// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: selector/selector.proto

package selector

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

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Additional []string `protobuf:"bytes,2,rep,name=additional,proto3" json:"additional,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selector_selector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_selector_selector_proto_msgTypes[0]
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
	return file_selector_selector_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetAdditional() []string {
	if x != nil {
		return x.Additional
	}
	return nil
}

type Handlers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handlers []*Handler `protobuf:"bytes,1,rep,name=handlers,proto3" json:"handlers,omitempty"`
}

func (x *Handlers) Reset() {
	*x = Handlers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selector_selector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handlers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handlers) ProtoMessage() {}

func (x *Handlers) ProtoReflect() protoreflect.Message {
	mi := &file_selector_selector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handlers.ProtoReflect.Descriptor instead.
func (*Handlers) Descriptor() ([]byte, []int) {
	return file_selector_selector_proto_rawDescGZIP(), []int{1}
}

func (x *Handlers) GetHandlers() []*Handler {
	if x != nil {
		return x.Handlers
	}
	return nil
}

type Handler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Select   string `protobuf:"bytes,2,opt,name=select,proto3" json:"select,omitempty"`
	Property int32  `protobuf:"varint,3,opt,name=property,proto3" json:"property,omitempty"`
	Enabled  bool   `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Global   bool   `protobuf:"varint,5,opt,name=global,proto3" json:"global,omitempty"`
}

func (x *Handler) Reset() {
	*x = Handler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_selector_selector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handler) ProtoMessage() {}

func (x *Handler) ProtoReflect() protoreflect.Message {
	mi := &file_selector_selector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handler.ProtoReflect.Descriptor instead.
func (*Handler) Descriptor() ([]byte, []int) {
	return file_selector_selector_proto_rawDescGZIP(), []int{2}
}

func (x *Handler) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Handler) GetSelect() string {
	if x != nil {
		return x.Select
	}
	return ""
}

func (x *Handler) GetProperty() int32 {
	if x != nil {
		return x.Property
	}
	return 0
}

func (x *Handler) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Handler) GetGlobal() bool {
	if x != nil {
		return x.Global
	}
	return false
}

var file_selector_selector_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Group)(nil),
		Field:         1101,
		Name:          "selector.group",
		Tag:           "bytes,1101,opt,name=group",
		Filename:      "selector/selector.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Handlers)(nil),
		Field:         1102,
		Name:          "selector.handlers",
		Tag:           "bytes,1102,opt,name=handlers",
		Filename:      "selector/selector.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional selector.Group group = 1101;
	E_Group = &file_selector_selector_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional selector.Handlers handlers = 1102;
	E_Handlers = &file_selector_selector_proto_extTypes[1]
)

var File_selector_selector_proto protoreflect.FileDescriptor

var file_selector_selector_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x22, 0x39, 0x0a, 0x08, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x2d,
	0x0a, 0x08, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x52, 0x08, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x22, 0x83, 0x01,
	0x0a, 0x07, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x3a, 0x46, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcd, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x50, 0x0a, 0x08, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xce, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x73, 0x52, 0x08, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x67, 0x64,
	0x61, 0x63, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x3b, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_selector_selector_proto_rawDescOnce sync.Once
	file_selector_selector_proto_rawDescData = file_selector_selector_proto_rawDesc
)

func file_selector_selector_proto_rawDescGZIP() []byte {
	file_selector_selector_proto_rawDescOnce.Do(func() {
		file_selector_selector_proto_rawDescData = protoimpl.X.CompressGZIP(file_selector_selector_proto_rawDescData)
	})
	return file_selector_selector_proto_rawDescData
}

var file_selector_selector_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_selector_selector_proto_goTypes = []interface{}{
	(*Group)(nil),                       // 0: selector.Group
	(*Handlers)(nil),                    // 1: selector.Handlers
	(*Handler)(nil),                     // 2: selector.Handler
	(*descriptorpb.MethodOptions)(nil),  // 3: google.protobuf.MethodOptions
	(*descriptorpb.ServiceOptions)(nil), // 4: google.protobuf.ServiceOptions
}
var file_selector_selector_proto_depIdxs = []int32{
	2, // 0: selector.Handlers.handlers:type_name -> selector.Handler
	3, // 1: selector.group:extendee -> google.protobuf.MethodOptions
	4, // 2: selector.handlers:extendee -> google.protobuf.ServiceOptions
	0, // 3: selector.group:type_name -> selector.Group
	1, // 4: selector.handlers:type_name -> selector.Handlers
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_selector_selector_proto_init() }
func file_selector_selector_proto_init() {
	if File_selector_selector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_selector_selector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_selector_selector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handlers); i {
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
		file_selector_selector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handler); i {
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
			RawDescriptor: file_selector_selector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_selector_selector_proto_goTypes,
		DependencyIndexes: file_selector_selector_proto_depIdxs,
		MessageInfos:      file_selector_selector_proto_msgTypes,
		ExtensionInfos:    file_selector_selector_proto_extTypes,
	}.Build()
	File_selector_selector_proto = out.File
	file_selector_selector_proto_rawDesc = nil
	file_selector_selector_proto_goTypes = nil
	file_selector_selector_proto_depIdxs = nil
}