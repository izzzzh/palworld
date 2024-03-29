// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: pal_mate.proto

package pal_mate

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

type ListPalMateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentOne int64 `protobuf:"varint,1,opt,name=parent_one,json=parentOne,proto3" json:"parent_one,omitempty"`
	ParentTwo int64 `protobuf:"varint,2,opt,name=parent_two,json=parentTwo,proto3" json:"parent_two,omitempty"`
	Result    int64 `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	Page      int32 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListPalMateReq) Reset() {
	*x = ListPalMateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pal_mate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPalMateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPalMateReq) ProtoMessage() {}

func (x *ListPalMateReq) ProtoReflect() protoreflect.Message {
	mi := &file_pal_mate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPalMateReq.ProtoReflect.Descriptor instead.
func (*ListPalMateReq) Descriptor() ([]byte, []int) {
	return file_pal_mate_proto_rawDescGZIP(), []int{0}
}

func (x *ListPalMateReq) GetParentOne() int64 {
	if x != nil {
		return x.ParentOne
	}
	return 0
}

func (x *ListPalMateReq) GetParentTwo() int64 {
	if x != nil {
		return x.ParentTwo
	}
	return 0
}

func (x *ListPalMateReq) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ListPalMateReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Pal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *Pal) Reset() {
	*x = Pal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pal_mate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pal) ProtoMessage() {}

func (x *Pal) ProtoReflect() protoreflect.Message {
	mi := &file_pal_mate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pal.ProtoReflect.Descriptor instead.
func (*Pal) Descriptor() ([]byte, []int) {
	return file_pal_mate_proto_rawDescGZIP(), []int{1}
}

func (x *Pal) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pal) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type PalMate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentOne *Pal `protobuf:"bytes,1,opt,name=parent_one,json=parentOne,proto3" json:"parent_one,omitempty"`
	ParentTwo *Pal `protobuf:"bytes,2,opt,name=parent_two,json=parentTwo,proto3" json:"parent_two,omitempty"`
	Result    *Pal `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PalMate) Reset() {
	*x = PalMate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pal_mate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PalMate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PalMate) ProtoMessage() {}

func (x *PalMate) ProtoReflect() protoreflect.Message {
	mi := &file_pal_mate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PalMate.ProtoReflect.Descriptor instead.
func (*PalMate) Descriptor() ([]byte, []int) {
	return file_pal_mate_proto_rawDescGZIP(), []int{2}
}

func (x *PalMate) GetParentOne() *Pal {
	if x != nil {
		return x.ParentOne
	}
	return nil
}

func (x *PalMate) GetParentTwo() *Pal {
	if x != nil {
		return x.ParentTwo
	}
	return nil
}

func (x *PalMate) GetResult() *Pal {
	if x != nil {
		return x.Result
	}
	return nil
}

type ListPalMateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*PalMate `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListPalMateResp) Reset() {
	*x = ListPalMateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pal_mate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPalMateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPalMateResp) ProtoMessage() {}

func (x *ListPalMateResp) ProtoReflect() protoreflect.Message {
	mi := &file_pal_mate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPalMateResp.ProtoReflect.Descriptor instead.
func (*ListPalMateResp) Descriptor() ([]byte, []int) {
	return file_pal_mate_proto_rawDescGZIP(), []int{3}
}

func (x *ListPalMateResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListPalMateResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ListPalMateResp) GetData() []*PalMate {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pal_mate_proto protoreflect.FileDescriptor

var file_pal_mate_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x22, 0x7a, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x77, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x3d, 0x0a, 0x03, 0x50, 0x61, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x8c, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x6c, 0x4d, 0x61, 0x74,
	0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x50, 0x61, 0x6c, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x65, 0x12,
	0x2c, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x50,
	0x61, 0x6c, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x77, 0x6f, 0x12, 0x25, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x6c, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x66, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x6c, 0x4d,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x50,
	0x61, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x53, 0x0a, 0x0d,
	0x50, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70,
	0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x6c, 0x4d,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x61, 0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pal_mate_proto_rawDescOnce sync.Once
	file_pal_mate_proto_rawDescData = file_pal_mate_proto_rawDesc
)

func file_pal_mate_proto_rawDescGZIP() []byte {
	file_pal_mate_proto_rawDescOnce.Do(func() {
		file_pal_mate_proto_rawDescData = protoimpl.X.CompressGZIP(file_pal_mate_proto_rawDescData)
	})
	return file_pal_mate_proto_rawDescData
}

var file_pal_mate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pal_mate_proto_goTypes = []interface{}{
	(*ListPalMateReq)(nil),  // 0: pal_mate.ListPalMateReq
	(*Pal)(nil),             // 1: pal_mate.Pal
	(*PalMate)(nil),         // 2: pal_mate.PalMate
	(*ListPalMateResp)(nil), // 3: pal_mate.ListPalMateResp
}
var file_pal_mate_proto_depIdxs = []int32{
	1, // 0: pal_mate.PalMate.parent_one:type_name -> pal_mate.Pal
	1, // 1: pal_mate.PalMate.parent_two:type_name -> pal_mate.Pal
	1, // 2: pal_mate.PalMate.result:type_name -> pal_mate.Pal
	2, // 3: pal_mate.ListPalMateResp.data:type_name -> pal_mate.PalMate
	0, // 4: pal_mate.PalMateServer.ListPalMate:input_type -> pal_mate.ListPalMateReq
	3, // 5: pal_mate.PalMateServer.ListPalMate:output_type -> pal_mate.ListPalMateResp
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pal_mate_proto_init() }
func file_pal_mate_proto_init() {
	if File_pal_mate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pal_mate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPalMateReq); i {
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
		file_pal_mate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pal); i {
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
		file_pal_mate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PalMate); i {
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
		file_pal_mate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPalMateResp); i {
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
			RawDescriptor: file_pal_mate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pal_mate_proto_goTypes,
		DependencyIndexes: file_pal_mate_proto_depIdxs,
		MessageInfos:      file_pal_mate_proto_msgTypes,
	}.Build()
	File_pal_mate_proto = out.File
	file_pal_mate_proto_rawDesc = nil
	file_pal_mate_proto_goTypes = nil
	file_pal_mate_proto_depIdxs = nil
}
