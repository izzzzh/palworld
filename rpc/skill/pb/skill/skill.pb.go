// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: skill.proto

package skill

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

type GetSkillReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSkillReq) Reset() {
	*x = GetSkillReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSkillReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkillReq) ProtoMessage() {}

func (x *GetSkillReq) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkillReq.ProtoReflect.Descriptor instead.
func (*GetSkillReq) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{0}
}

func (x *GetSkillReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Ct          int32  `protobuf:"varint,4,opt,name=ct,proto3" json:"ct,omitempty"`
	Power       int32  `protobuf:"varint,5,opt,name=power,proto3" json:"power,omitempty"`
	AttributeId int32  `protobuf:"varint,6,opt,name=attribute_id,json=attributeId,proto3" json:"attribute_id,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{1}
}

func (x *Skill) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Skill) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Skill) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Skill) GetCt() int32 {
	if x != nil {
		return x.Ct
	}
	return 0
}

func (x *Skill) GetPower() int32 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *Skill) GetAttributeId() int32 {
	if x != nil {
		return x.AttributeId
	}
	return 0
}

type GetSkillResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Skill `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSkillResp) Reset() {
	*x = GetSkillResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSkillResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSkillResp) ProtoMessage() {}

func (x *GetSkillResp) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSkillResp.ProtoReflect.Descriptor instead.
func (*GetSkillResp) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{2}
}

func (x *GetSkillResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetSkillResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetSkillResp) GetData() *Skill {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListSkillReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttributeId int32 `protobuf:"varint,1,opt,name=attribute_id,json=attributeId,proto3" json:"attribute_id,omitempty"`
}

func (x *ListSkillReq) Reset() {
	*x = ListSkillReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSkillReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSkillReq) ProtoMessage() {}

func (x *ListSkillReq) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSkillReq.ProtoReflect.Descriptor instead.
func (*ListSkillReq) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{3}
}

func (x *ListSkillReq) GetAttributeId() int32 {
	if x != nil {
		return x.AttributeId
	}
	return 0
}

type ListSkillResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Skill `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListSkillResp) Reset() {
	*x = ListSkillResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skill_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSkillResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSkillResp) ProtoMessage() {}

func (x *ListSkillResp) ProtoReflect() protoreflect.Message {
	mi := &file_skill_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSkillResp.ProtoReflect.Descriptor instead.
func (*ListSkillResp) Descriptor() ([]byte, []int) {
	return file_skill_proto_rawDescGZIP(), []int{4}
}

func (x *ListSkillResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListSkillResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListSkillResp) GetData() []*Skill {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_skill_proto protoreflect.FileDescriptor

var file_skill_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x22, 0x1d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x22,
	0x5f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x7a, 0x0a, 0x0b, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x33, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x12, 0x13, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skill_proto_rawDescOnce sync.Once
	file_skill_proto_rawDescData = file_skill_proto_rawDesc
)

func file_skill_proto_rawDescGZIP() []byte {
	file_skill_proto_rawDescOnce.Do(func() {
		file_skill_proto_rawDescData = protoimpl.X.CompressGZIP(file_skill_proto_rawDescData)
	})
	return file_skill_proto_rawDescData
}

var file_skill_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_skill_proto_goTypes = []interface{}{
	(*GetSkillReq)(nil),   // 0: skill.GetSkillReq
	(*Skill)(nil),         // 1: skill.Skill
	(*GetSkillResp)(nil),  // 2: skill.GetSkillResp
	(*ListSkillReq)(nil),  // 3: skill.ListSkillReq
	(*ListSkillResp)(nil), // 4: skill.ListSkillResp
}
var file_skill_proto_depIdxs = []int32{
	1, // 0: skill.GetSkillResp.data:type_name -> skill.Skill
	1, // 1: skill.ListSkillResp.data:type_name -> skill.Skill
	0, // 2: skill.SkillServer.GetSkill:input_type -> skill.GetSkillReq
	3, // 3: skill.SkillServer.ListSkill:input_type -> skill.ListSkillReq
	2, // 4: skill.SkillServer.GetSkill:output_type -> skill.GetSkillResp
	4, // 5: skill.SkillServer.ListSkill:output_type -> skill.ListSkillResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_skill_proto_init() }
func file_skill_proto_init() {
	if File_skill_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSkillReq); i {
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
		file_skill_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
		file_skill_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSkillResp); i {
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
		file_skill_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSkillReq); i {
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
		file_skill_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSkillResp); i {
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
			RawDescriptor: file_skill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_skill_proto_goTypes,
		DependencyIndexes: file_skill_proto_depIdxs,
		MessageInfos:      file_skill_proto_msgTypes,
	}.Build()
	File_skill_proto = out.File
	file_skill_proto_rawDesc = nil
	file_skill_proto_goTypes = nil
	file_skill_proto_depIdxs = nil
}
