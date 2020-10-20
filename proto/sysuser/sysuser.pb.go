// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sysuser/sysuser.proto

package sysuser

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	dbmodel "qshapi/proto/dbmodel"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//用户登录
type LoginReq struct {
	LoginType              dbmodel.LoginType `protobuf:"varint,1,opt,name=login_type,json=loginType,proto3,enum=dbmodel.LoginType" json:"login_type,omitempty"`
	UserNameOrPhoneOrEmail string            `protobuf:"bytes,2,opt,name=user_name_or_phone_or_email,json=userNameOrPhoneOrEmail,proto3" json:"user_name_or_phone_or_email,omitempty"`
	UserPasswordOrCode     string            `protobuf:"bytes,3,opt,name=user_password_or_code,json=userPasswordOrCode,proto3" json:"user_password_or_code,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4a741cbfe12ba90, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetLoginType() dbmodel.LoginType {
	if m != nil {
		return m.LoginType
	}
	return dbmodel.LoginType_NAME
}

func (m *LoginReq) GetUserNameOrPhoneOrEmail() string {
	if m != nil {
		return m.UserNameOrPhoneOrEmail
	}
	return ""
}

func (m *LoginReq) GetUserPasswordOrCode() string {
	if m != nil {
		return m.UserPasswordOrCode
	}
	return ""
}

//用户登录返回
type LoginResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserImg              string   `protobuf:"bytes,3,opt,name=user_img,json=userImg,proto3" json:"user_img,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4a741cbfe12ba90, []int{1}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResp) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginResp) GetUserImg() string {
	if m != nil {
		return m.UserImg
	}
	return ""
}

//用户注册请求
type RegistryReq struct {
	UserType             dbmodel.UserType `protobuf:"varint,6,opt,name=user_type,json=userType,proto3,enum=dbmodel.UserType" json:"user_type,omitempty"`
	UserName             string           `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPassword         string           `protobuf:"bytes,2,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
	UserPasswordAgain    string           `protobuf:"bytes,3,opt,name=user_password_again,json=userPasswordAgain,proto3" json:"user_password_again,omitempty"`
	UserPhone            string           `protobuf:"bytes,4,opt,name=user_phone,json=userPhone,proto3" json:"user_phone,omitempty"`
	UserPhoneCode        string           `protobuf:"bytes,5,opt,name=user_phone_code,json=userPhoneCode,proto3" json:"user_phone_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RegistryReq) Reset()         { *m = RegistryReq{} }
func (m *RegistryReq) String() string { return proto.CompactTextString(m) }
func (*RegistryReq) ProtoMessage()    {}
func (*RegistryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4a741cbfe12ba90, []int{2}
}

func (m *RegistryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryReq.Unmarshal(m, b)
}
func (m *RegistryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryReq.Marshal(b, m, deterministic)
}
func (m *RegistryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryReq.Merge(m, src)
}
func (m *RegistryReq) XXX_Size() int {
	return xxx_messageInfo_RegistryReq.Size(m)
}
func (m *RegistryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryReq proto.InternalMessageInfo

func (m *RegistryReq) GetUserType() dbmodel.UserType {
	if m != nil {
		return m.UserType
	}
	return dbmodel.UserType_ADMIN
}

func (m *RegistryReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegistryReq) GetUserPassword() string {
	if m != nil {
		return m.UserPassword
	}
	return ""
}

func (m *RegistryReq) GetUserPasswordAgain() string {
	if m != nil {
		return m.UserPasswordAgain
	}
	return ""
}

func (m *RegistryReq) GetUserPhone() string {
	if m != nil {
		return m.UserPhone
	}
	return ""
}

func (m *RegistryReq) GetUserPhoneCode() string {
	if m != nil {
		return m.UserPhoneCode
	}
	return ""
}

//修改密码
type ChangePasswordReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserPassword         string   `protobuf:"bytes,2,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
	UserPasswordAgain    string   `protobuf:"bytes,3,opt,name=user_password_again,json=userPasswordAgain,proto3" json:"user_password_again,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordReq) Reset()         { *m = ChangePasswordReq{} }
func (m *ChangePasswordReq) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordReq) ProtoMessage()    {}
func (*ChangePasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4a741cbfe12ba90, []int{3}
}

func (m *ChangePasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordReq.Unmarshal(m, b)
}
func (m *ChangePasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordReq.Marshal(b, m, deterministic)
}
func (m *ChangePasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordReq.Merge(m, src)
}
func (m *ChangePasswordReq) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordReq.Size(m)
}
func (m *ChangePasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordReq proto.InternalMessageInfo

func (m *ChangePasswordReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChangePasswordReq) GetUserPassword() string {
	if m != nil {
		return m.UserPassword
	}
	return ""
}

func (m *ChangePasswordReq) GetUserPasswordAgain() string {
	if m != nil {
		return m.UserPasswordAgain
	}
	return ""
}

type UserInfoListReq struct {
	PageReq              *dbmodel.PageReq `protobuf:"bytes,1,opt,name=page_req,json=pageReq,proto3" json:"page_req,omitempty"`
	UserName             string           `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPhone            string           `protobuf:"bytes,4,opt,name=user_phone,json=userPhone,proto3" json:"user_phone,omitempty"`
	UserEmail            string           `protobuf:"bytes,5,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UserInfoListReq) Reset()         { *m = UserInfoListReq{} }
func (m *UserInfoListReq) String() string { return proto.CompactTextString(m) }
func (*UserInfoListReq) ProtoMessage()    {}
func (*UserInfoListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4a741cbfe12ba90, []int{4}
}

func (m *UserInfoListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoListReq.Unmarshal(m, b)
}
func (m *UserInfoListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoListReq.Marshal(b, m, deterministic)
}
func (m *UserInfoListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoListReq.Merge(m, src)
}
func (m *UserInfoListReq) XXX_Size() int {
	return xxx_messageInfo_UserInfoListReq.Size(m)
}
func (m *UserInfoListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoListReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoListReq proto.InternalMessageInfo

func (m *UserInfoListReq) GetPageReq() *dbmodel.PageReq {
	if m != nil {
		return m.PageReq
	}
	return nil
}

func (m *UserInfoListReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfoListReq) GetUserPhone() string {
	if m != nil {
		return m.UserPhone
	}
	return ""
}

func (m *UserInfoListReq) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "sysuser.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "sysuser.LoginResp")
	proto.RegisterType((*RegistryReq)(nil), "sysuser.RegistryReq")
	proto.RegisterType((*ChangePasswordReq)(nil), "sysuser.ChangePasswordReq")
	proto.RegisterType((*UserInfoListReq)(nil), "sysuser.UserInfoListReq")
}

func init() { proto.RegisterFile("proto/sysuser/sysuser.proto", fileDescriptor_d4a741cbfe12ba90) }

var fileDescriptor_d4a741cbfe12ba90 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0x86, 0x31, 0x7c, 0xc1, 0xce, 0x40, 0xc8, 0x97, 0x2d, 0xad, 0x42, 0x50, 0x25, 0x94, 0xfe,
	0x88, 0x16, 0x29, 0x10, 0x2a, 0xf5, 0x04, 0xf5, 0x20, 0x05, 0x54, 0x21, 0xd1, 0x82, 0x1c, 0x7a,
	0xd0, 0x9e, 0x44, 0xa6, 0x9e, 0x1a, 0xab, 0x8e, 0x77, 0x59, 0x07, 0xda, 0x5c, 0x4a, 0x2f, 0xa2,
	0xd7, 0xd1, 0x3b, 0xea, 0x71, 0x35, 0xe3, 0x9f, 0xd8, 0x26, 0x22, 0x02, 0xa9, 0x47, 0xd9, 0x9d,
	0x79, 0x67, 0xe7, 0xd9, 0x77, 0xc7, 0x0a, 0xac, 0x2b, 0x2d, 0x47, 0x72, 0x3b, 0x1a, 0x47, 0x57,
	0x11, 0xea, 0xf4, 0xb7, 0xc3, 0x51, 0x61, 0x26, 0xdb, 0x56, 0xa2, 0x72, 0xcf, 0x87, 0xd2, 0xc5,
	0x20, 0xfd, 0x8d, 0x55, 0xad, 0xb5, 0x62, 0x32, 0x97, 0x6a, 0xff, 0x32, 0xc0, 0x3a, 0x96, 0x9e,
	0x1f, 0xda, 0x78, 0x29, 0xba, 0x00, 0x01, 0xad, 0x07, 0xa3, 0xb1, 0xc2, 0xa6, 0xb1, 0x61, 0x6c,
	0xae, 0xec, 0x8a, 0x4e, 0x7a, 0x16, 0xcb, 0xce, 0xc6, 0x0a, 0xed, 0x6a, 0x90, 0x2e, 0xc5, 0x1e,
	0xac, 0x53, 0xff, 0x41, 0xe8, 0x0c, 0x71, 0x20, 0xf5, 0x40, 0x5d, 0xc8, 0x90, 0x17, 0x38, 0x74,
	0xfc, 0xa0, 0x39, 0xbf, 0x61, 0x6c, 0x56, 0xed, 0x47, 0x24, 0xf9, 0xe0, 0x0c, 0xf1, 0x44, 0x9f,
	0x52, 0xfe, 0x44, 0x1f, 0x52, 0x56, 0x74, 0xe1, 0x21, 0x17, 0x2b, 0x27, 0x8a, 0xbe, 0x4b, 0xed,
	0x52, 0xdd, 0x17, 0xe9, 0x62, 0x73, 0x81, 0xcb, 0x04, 0x25, 0x4f, 0x93, 0xdc, 0x89, 0xde, 0x97,
	0x2e, 0xb6, 0x3f, 0x41, 0x35, 0xc1, 0x8d, 0x94, 0x58, 0x85, 0xca, 0x48, 0x7e, 0xc3, 0x90, 0x51,
	0xab, 0x76, 0xbc, 0x11, 0xeb, 0x50, 0xcd, 0x90, 0x12, 0x00, 0x2b, 0x05, 0x10, 0x6b, 0xc0, 0xeb,
	0x81, 0x3f, 0xf4, 0x92, 0x2e, 0x26, 0xed, 0x8f, 0x86, 0x5e, 0xfb, 0x8f, 0x01, 0x4b, 0x36, 0x7a,
	0x7e, 0x34, 0xd2, 0x63, 0x72, 0xa3, 0x93, 0x9c, 0xc3, 0x66, 0x2c, 0xb2, 0x19, 0x8d, 0xcc, 0x8c,
	0x8f, 0x11, 0x6a, 0xf6, 0x82, 0x8f, 0x63, 0x2b, 0x0a, 0x7d, 0x8d, 0x52, 0xdf, 0x27, 0x50, 0x2b,
	0x5c, 0x35, 0x01, 0x5b, 0xce, 0x5f, 0x51, 0x74, 0xe0, 0x41, 0xd1, 0x0f, 0xc7, 0x73, 0xfc, 0x30,
	0xe1, 0x6c, 0xe4, 0xa5, 0x3d, 0x4a, 0x88, 0xc7, 0x00, 0xb1, 0x9e, 0x4c, 0x6d, 0xfe, 0xc7, 0x32,
	0x66, 0x60, 0x97, 0xc5, 0x73, 0xa8, 0x4f, 0xd2, 0xb1, 0xb1, 0x15, 0xd6, 0xd4, 0x32, 0x0d, 0x7b,
	0xfa, 0x03, 0x1a, 0xfb, 0x17, 0x4e, 0xe8, 0x61, 0x7a, 0x3a, 0xdd, 0x7e, 0x05, 0xe6, 0x7d, 0x97,
	0xaf, 0xb1, 0x60, 0xcf, 0xfb, 0xee, 0x3f, 0xb9, 0x40, 0xfb, 0xa7, 0x01, 0x75, 0x72, 0xf2, 0x28,
	0xfc, 0x2a, 0x8f, 0xfd, 0x68, 0x44, 0x8d, 0xb7, 0xc0, 0x52, 0x8e, 0x87, 0x03, 0x8d, 0x97, 0xdc,
	0x7e, 0x69, 0xf7, 0xff, 0xcc, 0xf5, 0x53, 0xc7, 0x43, 0x1b, 0x2f, 0x6d, 0x53, 0xc5, 0x8b, 0xa2,
	0xe7, 0x0b, 0x25, 0xcf, 0x67, 0xd8, 0x93, 0xa6, 0xe3, 0x49, 0xad, 0x4c, 0xd2, 0x3c, 0x9c, 0xbb,
	0xbf, 0x2d, 0x30, 0x89, 0xad, 0xaf, 0xaf, 0xc5, 0x0e, 0x54, 0x78, 0xea, 0x44, 0xa3, 0x93, 0x7e,
	0x7f, 0xe9, 0x47, 0xd3, 0x12, 0xe5, 0x50, 0xa4, 0xda, 0x73, 0x62, 0x1b, 0xac, 0x74, 0x96, 0xc4,
	0x6a, 0xa6, 0xc8, 0x8d, 0x57, 0x6b, 0x29, 0xbb, 0xd5, 0x91, 0xdb, 0x9e, 0x13, 0x7b, 0xb0, 0x52,
	0x7c, 0x04, 0xd1, 0xca, 0xca, 0x6e, 0xbc, 0x4e, 0xb9, 0xf8, 0x0d, 0x2c, 0xe7, 0x6d, 0x14, 0xcd,
	0xac, 0xb4, 0xe4, 0x6e, 0xab, 0x51, 0xf2, 0x92, 0x61, 0xb7, 0xc0, 0x3a, 0x74, 0xfd, 0x11, 0x69,
	0xc5, 0xc4, 0xec, 0xfe, 0x38, 0xa2, 0x48, 0xb9, 0x57, 0x22, 0xb6, 0x65, 0x80, 0x45, 0x31, 0x45,
	0xca, 0xe2, 0x67, 0x60, 0x1e, 0x60, 0xc0, 0xda, 0x7c, 0xa6, 0x2c, 0xeb, 0x82, 0x45, 0x1a, 0x66,
	0xbf, 0xf1, 0xda, 0xd3, 0x99, 0xbb, 0x50, 0x4b, 0x99, 0xdf, 0x69, 0x79, 0xa5, 0x44, 0x23, 0xcf,
	0xc2, 0xa1, 0x72, 0x97, 0x97, 0xb0, 0x7c, 0x80, 0xc1, 0xa4, 0xe2, 0x36, 0xa2, 0xd7, 0x50, 0xcb,
	0x84, 0x77, 0xc1, 0x4a, 0xdc, 0x79, 0x8f, 0xe1, 0x55, 0xd1, 0x1d, 0x8a, 0x4c, 0x77, 0x87, 0xb5,
	0x33, 0xdc, 0x21, 0xcd, 0x5d, 0x30, 0x5e, 0x80, 0x49, 0x18, 0x3d, 0xe5, 0x8b, 0x7a, 0x9e, 0xa2,
	0xa7, 0xfc, 0xf2, 0xe9, 0x4f, 0x61, 0xf1, 0x00, 0x03, 0x52, 0xde, 0xc6, 0xb0, 0x03, 0x66, 0x4f,
	0xf9, 0xf7, 0x40, 0xa0, 0xcf, 0xa7, 0x80, 0xd0, 0xd7, 0xd7, 0xd3, 0x11, 0x48, 0x39, 0x03, 0xa1,
	0xaf, 0xaf, 0xef, 0xf1, 0x18, 0x67, 0x1a, 0x4b, 0xa3, 0x4a, 0x91, 0xe9, 0x8f, 0xc1, 0xda, 0x19,
	0x8f, 0x41, 0x9a, 0x3b, 0x60, 0xbc, 0xad, 0x7f, 0xae, 0x15, 0xfe, 0xc3, 0xcf, 0x17, 0x79, 0xfb,
	0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xb8, 0xf3, 0xe9, 0xdb, 0x07, 0x00, 0x00,
}
