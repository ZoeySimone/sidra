// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: options/sidra.proto

package options

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

type FileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// define Terraform provider name
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_sidra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_sidra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
func (*FileOptions) Descriptor() ([]byte, []int) {
	return file_options_sidra_proto_rawDescGZIP(), []int{0}
}

func (x *FileOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// operation type, one of: Create, Read, Update, Delete, Exists
	Operation string `protobuf:"bytes,1,opt,name=Operation,proto3" json:"Operation,omitempty"`
	// resource the method is operating on
	Resource string `protobuf:"bytes,2,opt,name=Resource,proto3" json:"Resource,omitempty"`
}

func (x *MethodOptions) Reset() {
	*x = MethodOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_sidra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodOptions) ProtoMessage() {}

func (x *MethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_sidra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodOptions.ProtoReflect.Descriptor instead.
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return file_options_sidra_proto_rawDescGZIP(), []int{1}
}

func (x *MethodOptions) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *MethodOptions) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// custom resource name
	ResourceName string `protobuf:"bytes,1,opt,name=ResourceName,proto3" json:"ResourceName,omitempty"`
	// prevents a resource from being generated for a message
	Ignore bool `protobuf:"varint,2,opt,name=Ignore,proto3" json:"Ignore,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_sidra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_sidra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_options_sidra_proto_rawDescGZIP(), []int{2}
}

func (x *MessageOptions) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MessageOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

type FieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// avoids adding field to the resource
	Ignore bool `protobuf:"varint,1,opt,name=Ignore,proto3" json:"Ignore,omitempty"`
	// marks field in Terraform as sensitive
	Sensitive bool `protobuf:"varint,2,opt,name=Sensitive,proto3" json:"Sensitive,omitempty"`
	// define which field to use as an ID for the resource, only one field may be an ID
	ID bool `protobuf:"varint,3,opt,name=ID,proto3" json:"ID,omitempty"`
	// define whether value is computed, required or optional
	Behaviour string `protobuf:"bytes,4,opt,name=Behaviour,proto3" json:"Behaviour,omitempty"`
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_sidra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_options_sidra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_options_sidra_proto_rawDescGZIP(), []int{3}
}

func (x *FieldOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *FieldOptions) GetSensitive() bool {
	if x != nil {
		return x.Sensitive
	}
	return false
}

func (x *FieldOptions) GetID() bool {
	if x != nil {
		return x.ID
	}
	return false
}

func (x *FieldOptions) GetBehaviour() string {
	if x != nil {
		return x.Behaviour
	}
	return ""
}

var file_options_sidra_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileOptions)(nil),
		Field:         1234,
		Name:          "sidra.Provider",
		Tag:           "bytes,1234,opt,name=Provider",
		Filename:      "options/sidra.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         1234,
		Name:          "sidra.Resource",
		Tag:           "bytes,1234,opt,name=Resource",
		Filename:      "options/sidra.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         1234,
		Name:          "sidra.Field",
		Tag:           "bytes,1234,opt,name=Field",
		Filename:      "options/sidra.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodOptions)(nil),
		Field:         1234,
		Name:          "sidra.Method",
		Tag:           "bytes,1234,opt,name=Method",
		Filename:      "options/sidra.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional sidra.FileOptions Provider = 1234;
	E_Provider = &file_options_sidra_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional sidra.MessageOptions Resource = 1234;
	E_Resource = &file_options_sidra_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional sidra.FieldOptions Field = 1234;
	E_Field = &file_options_sidra_proto_extTypes[2]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional sidra.MethodOptions Method = 1234;
	E_Method = &file_options_sidra_proto_extTypes[3]
)

var File_options_sidra_proto protoreflect.FileDescriptor

var file_options_sidra_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x72, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x69, 0x64, 0x72, 0x61, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21,
	0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x49, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x0e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x22, 0x72, 0x0a, 0x0c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x75, 0x72, 0x3a, 0x50,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x69, 0x64, 0x72, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01,
	0x3a, 0x56, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x69, 0x64, 0x72, 0x61, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x4c, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd2, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x69, 0x64, 0x72, 0x61, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x50, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd2, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x69, 0x64, 0x72, 0x61, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_sidra_proto_rawDescOnce sync.Once
	file_options_sidra_proto_rawDescData = file_options_sidra_proto_rawDesc
)

func file_options_sidra_proto_rawDescGZIP() []byte {
	file_options_sidra_proto_rawDescOnce.Do(func() {
		file_options_sidra_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_sidra_proto_rawDescData)
	})
	return file_options_sidra_proto_rawDescData
}

var file_options_sidra_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_options_sidra_proto_goTypes = []interface{}{
	(*FileOptions)(nil),                 // 0: sidra.FileOptions
	(*MethodOptions)(nil),               // 1: sidra.MethodOptions
	(*MessageOptions)(nil),              // 2: sidra.MessageOptions
	(*FieldOptions)(nil),                // 3: sidra.FieldOptions
	(*descriptorpb.FileOptions)(nil),    // 4: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 5: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 6: google.protobuf.FieldOptions
	(*descriptorpb.MethodOptions)(nil),  // 7: google.protobuf.MethodOptions
}
var file_options_sidra_proto_depIdxs = []int32{
	4, // 0: sidra.Provider:extendee -> google.protobuf.FileOptions
	5, // 1: sidra.Resource:extendee -> google.protobuf.MessageOptions
	6, // 2: sidra.Field:extendee -> google.protobuf.FieldOptions
	7, // 3: sidra.Method:extendee -> google.protobuf.MethodOptions
	0, // 4: sidra.Provider:type_name -> sidra.FileOptions
	2, // 5: sidra.Resource:type_name -> sidra.MessageOptions
	3, // 6: sidra.Field:type_name -> sidra.FieldOptions
	1, // 7: sidra.Method:type_name -> sidra.MethodOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	4, // [4:8] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_options_sidra_proto_init() }
func file_options_sidra_proto_init() {
	if File_options_sidra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_sidra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOptions); i {
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
		file_options_sidra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodOptions); i {
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
		file_options_sidra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_options_sidra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOptions); i {
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
			RawDescriptor: file_options_sidra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_options_sidra_proto_goTypes,
		DependencyIndexes: file_options_sidra_proto_depIdxs,
		MessageInfos:      file_options_sidra_proto_msgTypes,
		ExtensionInfos:    file_options_sidra_proto_extTypes,
	}.Build()
	File_options_sidra_proto = out.File
	file_options_sidra_proto_rawDesc = nil
	file_options_sidra_proto_goTypes = nil
	file_options_sidra_proto_depIdxs = nil
}