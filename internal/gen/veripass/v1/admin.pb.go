// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: veripass/v1/admin.proto

package veripassv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Admin struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hostel        string                 `protobuf:"bytes,3,opt,name=hostel,proto3" json:"hostel,omitempty"`
	CanAddPass    bool                   `protobuf:"varint,4,opt,name=can_add_pass,json=canAddPass,proto3" json:"can_add_pass,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Admin) Reset() {
	*x = Admin{}
	mi := &file_veripass_v1_admin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_veripass_v1_admin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_veripass_v1_admin_proto_rawDescGZIP(), []int{0}
}

func (x *Admin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Admin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Admin) GetHostel() string {
	if x != nil {
		return x.Hostel
	}
	return ""
}

func (x *Admin) GetCanAddPass() bool {
	if x != nil {
		return x.CanAddPass
	}
	return false
}

type GetAdminRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAdminRequest) Reset() {
	*x = GetAdminRequest{}
	mi := &file_veripass_v1_admin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminRequest) ProtoMessage() {}

func (x *GetAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_veripass_v1_admin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminRequest.ProtoReflect.Descriptor instead.
func (*GetAdminRequest) Descriptor() ([]byte, []int) {
	return file_veripass_v1_admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_veripass_v1_admin_proto protoreflect.FileDescriptor

const file_veripass_v1_admin_proto_rawDesc = "" +
	"\n" +
	"\x17veripass/v1/admin.proto\x12\vveripass.v1\"k\n" +
	"\x05Admin\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x16\n" +
	"\x06hostel\x18\x03 \x01(\tR\x06hostel\x12 \n" +
	"\fcan_add_pass\x18\x04 \x01(\bR\n" +
	"canAddPass\"'\n" +
	"\x0fGetAdminRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email2L\n" +
	"\fAdminService\x12<\n" +
	"\bGetAdmin\x12\x1c.veripass.v1.GetAdminRequest\x1a\x12.veripass.v1.AdminB\xae\x01\n" +
	"\x0fcom.veripass.v1B\n" +
	"AdminProtoP\x01ZBgithub.com/chetan0402/veripass/internal/gen/veripass/v1;veripassv1\xa2\x02\x03VXX\xaa\x02\vVeripass.V1\xca\x02\vVeripass\\V1\xe2\x02\x17Veripass\\V1\\GPBMetadata\xea\x02\fVeripass::V1b\x06proto3"

var (
	file_veripass_v1_admin_proto_rawDescOnce sync.Once
	file_veripass_v1_admin_proto_rawDescData []byte
)

func file_veripass_v1_admin_proto_rawDescGZIP() []byte {
	file_veripass_v1_admin_proto_rawDescOnce.Do(func() {
		file_veripass_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_veripass_v1_admin_proto_rawDesc), len(file_veripass_v1_admin_proto_rawDesc)))
	})
	return file_veripass_v1_admin_proto_rawDescData
}

var file_veripass_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_veripass_v1_admin_proto_goTypes = []any{
	(*Admin)(nil),           // 0: veripass.v1.Admin
	(*GetAdminRequest)(nil), // 1: veripass.v1.GetAdminRequest
}
var file_veripass_v1_admin_proto_depIdxs = []int32{
	1, // 0: veripass.v1.AdminService.GetAdmin:input_type -> veripass.v1.GetAdminRequest
	0, // 1: veripass.v1.AdminService.GetAdmin:output_type -> veripass.v1.Admin
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_veripass_v1_admin_proto_init() }
func file_veripass_v1_admin_proto_init() {
	if File_veripass_v1_admin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_veripass_v1_admin_proto_rawDesc), len(file_veripass_v1_admin_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_veripass_v1_admin_proto_goTypes,
		DependencyIndexes: file_veripass_v1_admin_proto_depIdxs,
		MessageInfos:      file_veripass_v1_admin_proto_msgTypes,
	}.Build()
	File_veripass_v1_admin_proto = out.File
	file_veripass_v1_admin_proto_goTypes = nil
	file_veripass_v1_admin_proto_depIdxs = nil
}
