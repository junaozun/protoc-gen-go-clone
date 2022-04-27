// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: example.proto

package example

import (
	common1 "gitlab.uuzu.com/war/pbtool/cmd/protoc-gen-clone/example/common1"
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

//专属装备抽卡
// |DrawOnlyEquipBack
type DrawOnlyEquipBack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stu              map[int32]*StudentMc                  `protobuf:"bytes,1,rep,name=stu,proto3" json:"stu,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DrawData         map[int32]*DrawOnlyEquipBack_DrawData `protobuf:"bytes,2,rep,name=drawData,proto3" json:"drawData,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SpecialDrawCount map[int32]string                      `protobuf:"bytes,10,rep,name=specialDrawCount,proto3" json:"specialDrawCount,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DrawOnlyEquipBack) Reset() {
	*x = DrawOnlyEquipBack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawOnlyEquipBack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawOnlyEquipBack) ProtoMessage() {}

func (x *DrawOnlyEquipBack) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawOnlyEquipBack.ProtoReflect.Descriptor instead.
func (*DrawOnlyEquipBack) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *DrawOnlyEquipBack) GetStu() map[int32]*StudentMc {
	if x != nil {
		return x.Stu
	}
	return nil
}

func (x *DrawOnlyEquipBack) GetDrawData() map[int32]*DrawOnlyEquipBack_DrawData {
	if x != nil {
		return x.DrawData
	}
	return nil
}

func (x *DrawOnlyEquipBack) GetSpecialDrawCount() map[int32]string {
	if x != nil {
		return x.SpecialDrawCount
	}
	return nil
}

type StudentMc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StudentMc) Reset() {
	*x = StudentMc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentMc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentMc) ProtoMessage() {}

func (x *StudentMc) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentMc.ProtoReflect.Descriptor instead.
func (*StudentMc) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *StudentMc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// |User
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *common1.User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetUser() *common1.User {
	if x != nil {
		return x.User
	}
	return nil
}

type DrawOnlyEquipBack_DrawData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpecialDrawCount map[string]int32 `protobuf:"bytes,1,rep,name=specialDrawCount,proto3" json:"specialDrawCount,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 特殊规则
	HistoryDrawCount int32            `protobuf:"varint,10,opt,name=historyDrawCount,proto3" json:"historyDrawCount,omitempty"`                                                                                        //历史抽卡总次数
	ResId            int32            `protobuf:"varint,11,opt,name=resId,proto3" json:"resId,omitempty"`                                                                                                              //根据资源Id重制一些变量
}

func (x *DrawOnlyEquipBack_DrawData) Reset() {
	*x = DrawOnlyEquipBack_DrawData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrawOnlyEquipBack_DrawData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrawOnlyEquipBack_DrawData) ProtoMessage() {}

func (x *DrawOnlyEquipBack_DrawData) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrawOnlyEquipBack_DrawData.ProtoReflect.Descriptor instead.
func (*DrawOnlyEquipBack_DrawData) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DrawOnlyEquipBack_DrawData) GetSpecialDrawCount() map[string]int32 {
	if x != nil {
		return x.SpecialDrawCount
	}
	return nil
}

func (x *DrawOnlyEquipBack_DrawData) GetHistoryDrawCount() int32 {
	if x != nil {
		return x.HistoryDrawCount
	}
	return 0
}

func (x *DrawOnlyEquipBack_DrawData) GetResId() int32 {
	if x != nil {
		return x.ResId
	}
	return 0
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x05, 0x0a, 0x11, 0x44, 0x72, 0x61, 0x77,
	0x4f, 0x6e, 0x6c, 0x79, 0x45, 0x71, 0x75, 0x69, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x35, 0x0a,
	0x03, 0x73, 0x74, 0x75, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x4f, 0x6e, 0x6c, 0x79, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x75, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x03, 0x73, 0x74, 0x75, 0x12, 0x44, 0x0a, 0x08, 0x64, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x44, 0x72, 0x61, 0x77, 0x4f, 0x6e, 0x6c, 0x79, 0x45, 0x71, 0x75, 0x69, 0x70, 0x42, 0x61,
	0x63, 0x6b, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x64, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x10, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x44, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44,
	0x72, 0x61, 0x77, 0x4f, 0x6e, 0x6c, 0x79, 0x45, 0x71, 0x75, 0x69, 0x70, 0x42, 0x61, 0x63, 0x6b,
	0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x44, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x44,
	0x72, 0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0xf8, 0x01, 0x0a, 0x08, 0x44, 0x72, 0x61,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x65, 0x0a, 0x10, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x44, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x72, 0x61, 0x77, 0x4f, 0x6e,
	0x6c, 0x79, 0x45, 0x71, 0x75, 0x69, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x44, 0x72, 0x61, 0x77,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x44, 0x72, 0x61, 0x77,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x44, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44,
	0x72, 0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x49,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65, 0x73, 0x49, 0x64, 0x1a, 0x43,
	0x0a, 0x15, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x44, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x4a, 0x0a, 0x08, 0x53, 0x74, 0x75, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x4d, 0x63, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x60, 0x0a, 0x0d, 0x44, 0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x72, 0x61, 0x77,
	0x4f, 0x6e, 0x6c, 0x79, 0x45, 0x71, 0x75, 0x69, 0x70, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x44, 0x72,
	0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x43, 0x0a, 0x15, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x44, 0x72, 0x61, 0x77,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x4d, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x75, 0x75, 0x7a,
	0x75, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x72, 0x2f, 0x70, 0x62, 0x74, 0x6f, 0x6f, 0x6c,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_example_proto_goTypes = []interface{}{
	(*DrawOnlyEquipBack)(nil),          // 0: example.DrawOnlyEquipBack
	(*StudentMc)(nil),                  // 1: example.StudentMc
	(*User)(nil),                       // 2: example.User
	(*DrawOnlyEquipBack_DrawData)(nil), // 3: example.DrawOnlyEquipBack.DrawData
	nil,                                // 4: example.DrawOnlyEquipBack.StuEntry
	nil,                                // 5: example.DrawOnlyEquipBack.DrawDataEntry
	nil,                                // 6: example.DrawOnlyEquipBack.SpecialDrawCountEntry
	nil,                                // 7: example.DrawOnlyEquipBack.DrawData.SpecialDrawCountEntry
	(*common1.User)(nil),               // 8: common1.User
}
var file_example_proto_depIdxs = []int32{
	4, // 0: example.DrawOnlyEquipBack.stu:type_name -> example.DrawOnlyEquipBack.StuEntry
	5, // 1: example.DrawOnlyEquipBack.drawData:type_name -> example.DrawOnlyEquipBack.DrawDataEntry
	6, // 2: example.DrawOnlyEquipBack.specialDrawCount:type_name -> example.DrawOnlyEquipBack.SpecialDrawCountEntry
	8, // 3: example.User.User:type_name -> common1.User
	7, // 4: example.DrawOnlyEquipBack.DrawData.specialDrawCount:type_name -> example.DrawOnlyEquipBack.DrawData.SpecialDrawCountEntry
	1, // 5: example.DrawOnlyEquipBack.StuEntry.value:type_name -> example.StudentMc
	3, // 6: example.DrawOnlyEquipBack.DrawDataEntry.value:type_name -> example.DrawOnlyEquipBack.DrawData
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawOnlyEquipBack); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentMc); i {
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
		file_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrawOnlyEquipBack_DrawData); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}
