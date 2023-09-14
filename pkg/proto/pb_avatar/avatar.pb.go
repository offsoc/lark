// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pb_avatar/avatar.proto

package pb_avatar

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pb_enum "lark/pkg/proto/pb_enum"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetAvatarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId      int64                `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerType    pb_enum.AVATAR_OWNER `protobuf:"varint,2,opt,name=owner_type,json=ownerType,proto3,enum=pb_enum.AVATAR_OWNER" json:"owner_type,omitempty"`
	AvatarSmall  string               `protobuf:"bytes,3,opt,name=avatar_small,json=avatarSmall,proto3" json:"avatar_small,omitempty"`
	AvatarMedium string               `protobuf:"bytes,4,opt,name=avatar_medium,json=avatarMedium,proto3" json:"avatar_medium,omitempty"`
	AvatarLarge  string               `protobuf:"bytes,5,opt,name=avatar_large,json=avatarLarge,proto3" json:"avatar_large,omitempty"`
}

func (x *SetAvatarReq) Reset() {
	*x = SetAvatarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_avatar_avatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvatarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvatarReq) ProtoMessage() {}

func (x *SetAvatarReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_avatar_avatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvatarReq.ProtoReflect.Descriptor instead.
func (*SetAvatarReq) Descriptor() ([]byte, []int) {
	return file_pb_avatar_avatar_proto_rawDescGZIP(), []int{0}
}

func (x *SetAvatarReq) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *SetAvatarReq) GetOwnerType() pb_enum.AVATAR_OWNER {
	if x != nil {
		return x.OwnerType
	}
	return pb_enum.AVATAR_OWNER(0)
}

func (x *SetAvatarReq) GetAvatarSmall() string {
	if x != nil {
		return x.AvatarSmall
	}
	return ""
}

func (x *SetAvatarReq) GetAvatarMedium() string {
	if x != nil {
		return x.AvatarMedium
	}
	return ""
}

func (x *SetAvatarReq) GetAvatarLarge() string {
	if x != nil {
		return x.AvatarLarge
	}
	return ""
}

type SetAvatarResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Avatar *AvatarInfo `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"` // 头像
}

func (x *SetAvatarResp) Reset() {
	*x = SetAvatarResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_avatar_avatar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvatarResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvatarResp) ProtoMessage() {}

func (x *SetAvatarResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_avatar_avatar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvatarResp.ProtoReflect.Descriptor instead.
func (*SetAvatarResp) Descriptor() ([]byte, []int) {
	return file_pb_avatar_avatar_proto_rawDescGZIP(), []int{1}
}

func (x *SetAvatarResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SetAvatarResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SetAvatarResp) GetAvatar() *AvatarInfo {
	if x != nil {
		return x.Avatar
	}
	return nil
}

type AvatarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId      int64                `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerType    pb_enum.AVATAR_OWNER `protobuf:"varint,2,opt,name=owner_type,json=ownerType,proto3,enum=pb_enum.AVATAR_OWNER" json:"owner_type,omitempty"`
	AvatarSmall  string               `protobuf:"bytes,3,opt,name=avatar_small,json=avatarSmall,proto3" json:"avatar_small,omitempty"`
	AvatarMedium string               `protobuf:"bytes,4,opt,name=avatar_medium,json=avatarMedium,proto3" json:"avatar_medium,omitempty"`
	AvatarLarge  string               `protobuf:"bytes,5,opt,name=avatar_large,json=avatarLarge,proto3" json:"avatar_large,omitempty"`
}

func (x *AvatarInfo) Reset() {
	*x = AvatarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_avatar_avatar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarInfo) ProtoMessage() {}

func (x *AvatarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_avatar_avatar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarInfo.ProtoReflect.Descriptor instead.
func (*AvatarInfo) Descriptor() ([]byte, []int) {
	return file_pb_avatar_avatar_proto_rawDescGZIP(), []int{2}
}

func (x *AvatarInfo) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *AvatarInfo) GetOwnerType() pb_enum.AVATAR_OWNER {
	if x != nil {
		return x.OwnerType
	}
	return pb_enum.AVATAR_OWNER(0)
}

func (x *AvatarInfo) GetAvatarSmall() string {
	if x != nil {
		return x.AvatarSmall
	}
	return ""
}

func (x *AvatarInfo) GetAvatarMedium() string {
	if x != nil {
		return x.AvatarMedium
	}
	return ""
}

func (x *AvatarInfo) GetAvatarLarge() string {
	if x != nil {
		return x.AvatarLarge
	}
	return ""
}

var File_pb_avatar_avatar_proto protoreflect.FileDescriptor

var file_pb_avatar_avatar_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2f, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x62, 0x5f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x1a, 0x12, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x52, 0x09,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x65, 0x64, 0x69, 0x75,
	0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x61, 0x72, 0x67,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c,
	0x61, 0x72, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62,
	0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x52,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x73, 0x6d, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4d, 0x65, 0x64, 0x69,
	0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x61, 0x72,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x4c, 0x61, 0x72, 0x67, 0x65, 0x32, 0x48, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x3e, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x24, 0x5a, 0x22, 0x6c, 0x61, 0x72, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x62, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x3b, 0x70, 0x62, 0x5f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_avatar_avatar_proto_rawDescOnce sync.Once
	file_pb_avatar_avatar_proto_rawDescData = file_pb_avatar_avatar_proto_rawDesc
)

func file_pb_avatar_avatar_proto_rawDescGZIP() []byte {
	file_pb_avatar_avatar_proto_rawDescOnce.Do(func() {
		file_pb_avatar_avatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_avatar_avatar_proto_rawDescData)
	})
	return file_pb_avatar_avatar_proto_rawDescData
}

var file_pb_avatar_avatar_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_avatar_avatar_proto_goTypes = []interface{}{
	(*SetAvatarReq)(nil),      // 0: pb_avatar.SetAvatarReq
	(*SetAvatarResp)(nil),     // 1: pb_avatar.SetAvatarResp
	(*AvatarInfo)(nil),        // 2: pb_avatar.AvatarInfo
	(pb_enum.AVATAR_OWNER)(0), // 3: pb_enum.AVATAR_OWNER
}
var file_pb_avatar_avatar_proto_depIdxs = []int32{
	3, // 0: pb_avatar.SetAvatarReq.owner_type:type_name -> pb_enum.AVATAR_OWNER
	2, // 1: pb_avatar.SetAvatarResp.avatar:type_name -> pb_avatar.AvatarInfo
	3, // 2: pb_avatar.AvatarInfo.owner_type:type_name -> pb_enum.AVATAR_OWNER
	0, // 3: pb_avatar.Avatar.SetAvatar:input_type -> pb_avatar.SetAvatarReq
	1, // 4: pb_avatar.Avatar.SetAvatar:output_type -> pb_avatar.SetAvatarResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_avatar_avatar_proto_init() }
func file_pb_avatar_avatar_proto_init() {
	if File_pb_avatar_avatar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_avatar_avatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAvatarReq); i {
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
		file_pb_avatar_avatar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAvatarResp); i {
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
		file_pb_avatar_avatar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarInfo); i {
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
			RawDescriptor: file_pb_avatar_avatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_avatar_avatar_proto_goTypes,
		DependencyIndexes: file_pb_avatar_avatar_proto_depIdxs,
		MessageInfos:      file_pb_avatar_avatar_proto_msgTypes,
	}.Build()
	File_pb_avatar_avatar_proto = out.File
	file_pb_avatar_avatar_proto_rawDesc = nil
	file_pb_avatar_avatar_proto_goTypes = nil
	file_pb_avatar_avatar_proto_depIdxs = nil
}
