// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: user_model.proto

package serviceUser

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Accessibility struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DarkMode      bool                   `protobuf:"varint,2,opt,name=dark_mode,json=darkMode,proto3" json:"dark_mode,omitempty"`
	Transparency  bool                   `protobuf:"varint,3,opt,name=transparency,proto3" json:"transparency,omitempty"`
	AccentColorId uint64                 `protobuf:"varint,4,opt,name=accent_color_id,json=accentColorId,proto3" json:"accent_color_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,91,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,92,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Relasi ke Color
	AccentColor *Color `protobuf:"bytes,100,opt,name=accent_color,json=accentColor,proto3" json:"accent_color,omitempty"`
}

func (x *Accessibility) Reset() {
	*x = Accessibility{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accessibility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accessibility) ProtoMessage() {}

func (x *Accessibility) ProtoReflect() protoreflect.Message {
	mi := &file_user_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accessibility.ProtoReflect.Descriptor instead.
func (*Accessibility) Descriptor() ([]byte, []int) {
	return file_user_model_proto_rawDescGZIP(), []int{0}
}

func (x *Accessibility) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Accessibility) GetDarkMode() bool {
	if x != nil {
		return x.DarkMode
	}
	return false
}

func (x *Accessibility) GetTransparency() bool {
	if x != nil {
		return x.Transparency
	}
	return false
}

func (x *Accessibility) GetAccentColorId() uint64 {
	if x != nil {
		return x.AccentColorId
	}
	return 0
}

func (x *Accessibility) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Accessibility) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Accessibility) GetAccentColor() *Color {
	if x != nil {
		return x.AccentColor
	}
	return nil
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code      string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,91,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,92,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Relasi ke Accessibility sebagai AccentColor
	Accessibilities []*Accessibility `protobuf:"bytes,100,rep,name=accessibilities,proto3" json:"accessibilities,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_user_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_user_model_proto_rawDescGZIP(), []int{1}
}

func (x *Color) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Color) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Color) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Color) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Color) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Color) GetAccessibilities() []*Accessibility {
	if x != nil {
		return x.Accessibilities
	}
	return nil
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code      string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,91,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,92,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Relasi ke NotificationsFeatures
	NotificationsFeatures []*NotificationsFeature `protobuf:"bytes,100,rep,name=notifications_features,json=notificationsFeatures,proto3" json:"notifications_features,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_user_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_user_model_proto_rawDescGZIP(), []int{2}
}

func (x *Feature) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Feature) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Feature) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Feature) GetNotificationsFeatures() []*NotificationsFeature {
	if x != nil {
		return x.NotificationsFeatures
	}
	return nil
}

type NotificationsFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FeatureId uint64                 `protobuf:"varint,2,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	SettingId uint64                 `protobuf:"varint,3,opt,name=setting_id,json=settingId,proto3" json:"setting_id,omitempty"`
	IsEnabled bool                   `protobuf:"varint,4,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,91,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,92,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Relasi ke Feature dan Setting
	Feature *Feature `protobuf:"bytes,100,opt,name=feature,proto3" json:"feature,omitempty"`
	Setting *Setting `protobuf:"bytes,101,opt,name=setting,proto3" json:"setting,omitempty"`
}

func (x *NotificationsFeature) Reset() {
	*x = NotificationsFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationsFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationsFeature) ProtoMessage() {}

func (x *NotificationsFeature) ProtoReflect() protoreflect.Message {
	mi := &file_user_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationsFeature.ProtoReflect.Descriptor instead.
func (*NotificationsFeature) Descriptor() ([]byte, []int) {
	return file_user_model_proto_rawDescGZIP(), []int{3}
}

func (x *NotificationsFeature) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NotificationsFeature) GetFeatureId() uint64 {
	if x != nil {
		return x.FeatureId
	}
	return 0
}

func (x *NotificationsFeature) GetSettingId() uint64 {
	if x != nil {
		return x.SettingId
	}
	return 0
}

func (x *NotificationsFeature) GetIsEnabled() bool {
	if x != nil {
		return x.IsEnabled
	}
	return false
}

func (x *NotificationsFeature) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *NotificationsFeature) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *NotificationsFeature) GetFeature() *Feature {
	if x != nil {
		return x.Feature
	}
	return nil
}

func (x *NotificationsFeature) GetSetting() *Setting {
	if x != nil {
		return x.Setting
	}
	return nil
}

type Setting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username        string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	AccessibilityId uint64                 `protobuf:"varint,3,opt,name=accessibility_id,json=accessibilityId,proto3" json:"accessibility_id,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,91,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,92,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Relasi ke Accessibility
	Accessibility *Accessibility `protobuf:"bytes,100,opt,name=accessibility,proto3" json:"accessibility,omitempty"`
	// Relasi ke NotificationsFeatures sebagai Setting
	NotificationsFeatures []*NotificationsFeature `protobuf:"bytes,101,rep,name=notifications_features,json=notificationsFeatures,proto3" json:"notifications_features,omitempty"`
}

func (x *Setting) Reset() {
	*x = Setting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_user_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_user_model_proto_rawDescGZIP(), []int{4}
}

func (x *Setting) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Setting) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Setting) GetAccessibilityId() uint64 {
	if x != nil {
		return x.AccessibilityId
	}
	return 0
}

func (x *Setting) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Setting) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Setting) GetAccessibility() *Accessibility {
	if x != nil {
		return x.Accessibility
	}
	return nil
}

func (x *Setting) GetNotificationsFeatures() []*NotificationsFeature {
	if x != nil {
		return x.NotificationsFeatures
	}
	return nil
}

var File_user_model_proto protoreflect.FileDescriptor

var file_user_model_proto_rawDesc = []byte{
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x03, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28,
	0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x64, 0x61, 0x72, 0x6b, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0f, 0xba, 0xb9, 0x19, 0x0b,
	0x0a, 0x09, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x40, 0x01, 0x52, 0x08, 0x64, 0x61, 0x72,
	0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0e, 0xba, 0xb9, 0x19,
	0x0a, 0x0a, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x40, 0x01, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x33, 0x0a, 0x0f, 0x61, 0x63, 0x63,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x0b, 0xe0, 0x41, 0x02, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52,
	0x0d, 0x61, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x5b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x5c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x42,
	0x15, 0xba, 0xb9, 0x19, 0x11, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x49, 0x44, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x3a, 0x17, 0xba, 0xb9, 0x19, 0x13, 0x08, 0x01, 0x1a, 0x0f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0xb2, 0x02, 0x0a,
	0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x5b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x5c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x5c, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x15, 0xba, 0xb9, 0x19, 0x11, 0x2a,
	0x0f, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x49, 0x44,
	0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x3a, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x08, 0x01, 0x1a, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x73, 0x22, 0xe0, 0x02, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe0, 0x41, 0x02, 0xba,
	0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe0, 0x41, 0x02,
	0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x5b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x5c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x6c, 0x0a, 0x16, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x64,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x2a, 0x0b,
	0x0a, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x52, 0x15, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x3a, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x08, 0x01, 0x1a, 0x08, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x22, 0xdb, 0x03, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba, 0xb9,
	0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x0a,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x0b, 0xe0, 0x41, 0x02, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0b, 0xe0, 0x41,
	0x02, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0f, 0xba, 0xb9, 0x19, 0x0b, 0x0a, 0x09,
	0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x40, 0x01, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x5b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x5c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x07, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x22, 0x0b, 0x0a, 0x09, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x49, 0x44, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x42,
	0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x22, 0x0b, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x3a, 0x1e, 0xba, 0xb9, 0x19, 0x1a, 0x08, 0x01, 0x1a, 0x16, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x22, 0xce, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xe0, 0x41, 0x03, 0xba,
	0xb9, 0x19, 0x06, 0x0a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x10, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x0b, 0xe0, 0x41, 0x02, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01,
	0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x5b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x5c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x5a, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x17, 0xba, 0xb9, 0x19,
	0x13, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x6c, 0x0a, 0x16, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x65, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x2a, 0x0b, 0x0a,
	0x09, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x52, 0x15, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x3a, 0x10, 0xba, 0xb9, 0x19, 0x0c, 0x08, 0x01, 0x1a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x61, 0x2d, 0x59, 0x6f, 0x67, 0x79, 0x61, 0x6b,
	0x61, 0x72, 0x74, 0x61, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_model_proto_rawDescOnce sync.Once
	file_user_model_proto_rawDescData = file_user_model_proto_rawDesc
)

func file_user_model_proto_rawDescGZIP() []byte {
	file_user_model_proto_rawDescOnce.Do(func() {
		file_user_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_model_proto_rawDescData)
	})
	return file_user_model_proto_rawDescData
}

var file_user_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_model_proto_goTypes = []interface{}{
	(*Accessibility)(nil),         // 0: user.service.Accessibility
	(*Color)(nil),                 // 1: user.service.Color
	(*Feature)(nil),               // 2: user.service.Feature
	(*NotificationsFeature)(nil),  // 3: user.service.NotificationsFeature
	(*Setting)(nil),               // 4: user.service.Setting
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_user_model_proto_depIdxs = []int32{
	5,  // 0: user.service.Accessibility.created_at:type_name -> google.protobuf.Timestamp
	5,  // 1: user.service.Accessibility.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: user.service.Accessibility.accent_color:type_name -> user.service.Color
	5,  // 3: user.service.Color.created_at:type_name -> google.protobuf.Timestamp
	5,  // 4: user.service.Color.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 5: user.service.Color.accessibilities:type_name -> user.service.Accessibility
	5,  // 6: user.service.Feature.created_at:type_name -> google.protobuf.Timestamp
	5,  // 7: user.service.Feature.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 8: user.service.Feature.notifications_features:type_name -> user.service.NotificationsFeature
	5,  // 9: user.service.NotificationsFeature.created_at:type_name -> google.protobuf.Timestamp
	5,  // 10: user.service.NotificationsFeature.updated_at:type_name -> google.protobuf.Timestamp
	2,  // 11: user.service.NotificationsFeature.feature:type_name -> user.service.Feature
	4,  // 12: user.service.NotificationsFeature.setting:type_name -> user.service.Setting
	5,  // 13: user.service.Setting.created_at:type_name -> google.protobuf.Timestamp
	5,  // 14: user.service.Setting.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 15: user.service.Setting.accessibility:type_name -> user.service.Accessibility
	3,  // 16: user.service.Setting.notifications_features:type_name -> user.service.NotificationsFeature
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_user_model_proto_init() }
func file_user_model_proto_init() {
	if File_user_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accessibility); i {
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
		file_user_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
		file_user_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_user_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationsFeature); i {
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
		file_user_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Setting); i {
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
			RawDescriptor: file_user_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_model_proto_goTypes,
		DependencyIndexes: file_user_model_proto_depIdxs,
		MessageInfos:      file_user_model_proto_msgTypes,
	}.Build()
	File_user_model_proto = out.File
	file_user_model_proto_rawDesc = nil
	file_user_model_proto_goTypes = nil
	file_user_model_proto_depIdxs = nil
}
