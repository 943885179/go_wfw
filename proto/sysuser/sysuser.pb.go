// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sysuser.proto

package sysuser

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type LoginType int32

const (
	LoginType_NAME  LoginType = 0
	LoginType_PHONE LoginType = 1
	LoginType_EMAIL LoginType = 2
)

var LoginType_name = map[int32]string{
	0: "NAME",
	1: "PHONE",
	2: "EMAIL",
}

var LoginType_value = map[string]int32{
	"NAME":  0,
	"PHONE": 1,
	"EMAIL": 2,
}

func (x LoginType) String() string {
	return proto.EnumName(LoginType_name, int32(x))
}

func (LoginType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d42e5eade424557, []int{0}
}

type UserType int32

const (
	UserType_ADMIN    UserType = 0
	UserType_PLATFORM UserType = 1
	UserType_SHOP     UserType = 2
	UserType_CUSTOMER UserType = 3
)

var UserType_name = map[int32]string{
	0: "ADMIN",
	1: "PLATFORM",
	2: "SHOP",
	3: "CUSTOMER",
}

var UserType_value = map[string]int32{
	"ADMIN":    0,
	"PLATFORM": 1,
	"SHOP":     2,
	"CUSTOMER": 3,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}

func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d42e5eade424557, []int{1}
}

//用户登录
type LoginReq struct {
	LoginType              LoginType `protobuf:"varint,1,opt,name=login_type,json=loginType,proto3,enum=sysuser.LoginType" json:"login_type,omitempty"`
	UserNameOrPhoneOrEmail string    `protobuf:"bytes,2,opt,name=user_name_or_phone_or_email,json=userNameOrPhoneOrEmail,proto3" json:"user_name_or_phone_or_email,omitempty"`
	UserPasswordOrCode     string    `protobuf:"bytes,3,opt,name=user_password_or_code,json=userPasswordOrCode,proto3" json:"user_password_or_code,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}  `json:"-"`
	XXX_unrecognized       []byte    `json:"-"`
	XXX_sizecache          int32     `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d42e5eade424557, []int{0}
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

func (m *LoginReq) GetLoginType() LoginType {
	if m != nil {
		return m.LoginType
	}
	return LoginType_NAME
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
	return fileDescriptor_0d42e5eade424557, []int{1}
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

type RegistryReq struct {
	UserType             UserType `protobuf:"varint,6,opt,name=user_type,json=userType,proto3,enum=sysuser.UserType" json:"user_type,omitempty"`
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPassword         string   `protobuf:"bytes,2,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
	UserPasswordAgain    string   `protobuf:"bytes,3,opt,name=user_password_again,json=userPasswordAgain,proto3" json:"user_password_again,omitempty"`
	UserPhone            string   `protobuf:"bytes,4,opt,name=user_phone,json=userPhone,proto3" json:"user_phone,omitempty"`
	UserPhoneCode        string   `protobuf:"bytes,5,opt,name=user_phone_code,json=userPhoneCode,proto3" json:"user_phone_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryReq) Reset()         { *m = RegistryReq{} }
func (m *RegistryReq) String() string { return proto.CompactTextString(m) }
func (*RegistryReq) ProtoMessage()    {}
func (*RegistryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d42e5eade424557, []int{2}
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

func (m *RegistryReq) GetUserType() UserType {
	if m != nil {
		return m.UserType
	}
	return UserType_ADMIN
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

//用户登录返回
type RegistryResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserImg              string   `protobuf:"bytes,3,opt,name=user_img,json=userImg,proto3" json:"user_img,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryResp) Reset()         { *m = RegistryResp{} }
func (m *RegistryResp) String() string { return proto.CompactTextString(m) }
func (*RegistryResp) ProtoMessage()    {}
func (*RegistryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d42e5eade424557, []int{3}
}

func (m *RegistryResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryResp.Unmarshal(m, b)
}
func (m *RegistryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryResp.Marshal(b, m, deterministic)
}
func (m *RegistryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryResp.Merge(m, src)
}
func (m *RegistryResp) XXX_Size() int {
	return xxx_messageInfo_RegistryResp.Size(m)
}
func (m *RegistryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryResp proto.InternalMessageInfo

func (m *RegistryResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegistryResp) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegistryResp) GetUserImg() string {
	if m != nil {
		return m.UserImg
	}
	return ""
}

func init() {
	proto.RegisterEnum("sysuser.LoginType", LoginType_name, LoginType_value)
	proto.RegisterEnum("sysuser.UserType", UserType_name, UserType_value)
	proto.RegisterType((*LoginReq)(nil), "sysuser.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "sysuser.LoginResp")
	proto.RegisterType((*RegistryReq)(nil), "sysuser.RegistryReq")
	proto.RegisterType((*RegistryResp)(nil), "sysuser.RegistryResp")
}

func init() { proto.RegisterFile("sysuser.proto", fileDescriptor_0d42e5eade424557) }

var fileDescriptor_0d42e5eade424557 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6f, 0xd3, 0x40,
	0x10, 0x85, 0xe3, 0xb4, 0x69, 0xec, 0x21, 0xa1, 0xce, 0xd0, 0xa2, 0xd0, 0x0a, 0xa9, 0x0a, 0x12,
	0xaa, 0x8a, 0x14, 0x68, 0x39, 0xa1, 0x9e, 0x42, 0x31, 0x6a, 0xa4, 0x38, 0xb6, 0x9c, 0xf4, 0x00,
	0x42, 0xb2, 0x0c, 0x59, 0x19, 0x8b, 0xd8, 0xeb, 0xee, 0x86, 0xa2, 0xfc, 0x28, 0xfe, 0x1e, 0x67,
	0x34, 0x93, 0x75, 0x9a, 0x54, 0x1c, 0xb9, 0xcd, 0xce, 0x7b, 0xb3, 0xf3, 0xf6, 0x8b, 0x03, 0x6d,
	0xbd, 0xd4, 0x3f, 0xb5, 0x50, 0xfd, 0x52, 0xc9, 0x85, 0xc4, 0xa6, 0x39, 0xf6, 0x7e, 0x5b, 0x60,
	0x8f, 0x64, 0x9a, 0x15, 0x91, 0xb8, 0xc5, 0x73, 0x80, 0x39, 0xd5, 0xf1, 0x62, 0x59, 0x8a, 0xae,
	0x75, 0x62, 0x9d, 0x3e, 0xbe, 0xc0, 0x7e, 0x35, 0xc9, 0xb6, 0xe9, 0xb2, 0x14, 0x91, 0x33, 0xaf,
	0x4a, 0xbc, 0x84, 0x63, 0x12, 0xe3, 0x22, 0xc9, 0x45, 0x2c, 0x55, 0x5c, 0x7e, 0x97, 0x05, 0x17,
	0x22, 0x4f, 0xb2, 0x79, 0xb7, 0x7e, 0x62, 0x9d, 0x3a, 0xd1, 0x53, 0xb2, 0x8c, 0x93, 0x5c, 0x04,
	0x2a, 0x24, 0x3d, 0x50, 0x1e, 0xa9, 0x78, 0x0e, 0x87, 0x3c, 0x5c, 0x26, 0x5a, 0xff, 0x92, 0x6a,
	0x46, 0x73, 0xdf, 0xe4, 0x4c, 0x74, 0x77, 0x78, 0x0c, 0x49, 0x0c, 0x8d, 0x16, 0xa8, 0x2b, 0x39,
	0x13, 0xbd, 0x4f, 0xe0, 0x98, 0xb8, 0xba, 0xc4, 0x03, 0x68, 0x2c, 0xe4, 0x0f, 0x51, 0x70, 0x54,
	0x27, 0x5a, 0x1d, 0xf0, 0x18, 0x9c, 0x75, 0x24, 0x13, 0xc0, 0xae, 0x02, 0xe0, 0x33, 0xe0, 0x3a,
	0xce, 0xf2, 0xd4, 0x6c, 0x69, 0xd2, 0x79, 0x98, 0xa7, 0xbd, 0x3f, 0x16, 0x3c, 0x8a, 0x44, 0x9a,
	0xe9, 0x85, 0x5a, 0x12, 0x8d, 0xbe, 0xb9, 0x87, 0x61, 0xec, 0x31, 0x8c, 0xce, 0x1a, 0xc6, 0x8d,
	0x16, 0x8a, 0x59, 0xf0, 0x75, 0x8c, 0x62, 0x6b, 0xaf, 0xf5, 0x60, 0xef, 0x0b, 0x68, 0x6f, 0x3d,
	0xd5, 0x04, 0x6b, 0x6d, 0x3e, 0x11, 0xfb, 0xf0, 0x64, 0x9b, 0x47, 0x92, 0x26, 0x59, 0x61, 0x72,
	0x76, 0x36, 0xad, 0x03, 0x12, 0xf0, 0x39, 0xc0, 0xca, 0x4f, 0x50, 0xbb, 0xbb, 0x6c, 0xe3, 0x0c,
	0x4c, 0x19, 0x5f, 0xc2, 0xfe, 0xbd, 0xbc, 0x02, 0xdb, 0x60, 0x4f, 0x7b, 0xed, 0x61, 0xa6, 0x5f,
	0xa0, 0x75, 0xff, 0xee, 0xff, 0x8d, 0xf5, 0xec, 0x95, 0xf9, 0xc5, 0x98, 0x91, 0x0d, 0xbb, 0xe3,
	0x81, 0xef, 0xb9, 0x35, 0x74, 0xa0, 0x11, 0x5e, 0x07, 0x63, 0xcf, 0xb5, 0xa8, 0xf4, 0xfc, 0xc1,
	0x70, 0xe4, 0xd6, 0xcf, 0x2e, 0xc1, 0xae, 0xc8, 0x52, 0x7b, 0xf0, 0xc1, 0x1f, 0x8e, 0xdd, 0x1a,
	0xb6, 0xc0, 0x0e, 0x47, 0x83, 0xe9, 0xc7, 0x20, 0xf2, 0x5d, 0x8b, 0x2e, 0x99, 0x5c, 0x07, 0xa1,
	0x5b, 0xa7, 0xfe, 0xd5, 0xcd, 0x64, 0x1a, 0xf8, 0x5e, 0xe4, 0xee, 0x5c, 0xdc, 0x41, 0x93, 0x86,
	0x27, 0xea, 0x0e, 0xdf, 0x40, 0x83, 0x97, 0x62, 0x67, 0xfb, 0xf3, 0x8d, 0xc4, 0xed, 0x11, 0x3e,
	0x6c, 0xe9, 0xb2, 0x57, 0xc3, 0x77, 0x60, 0x57, 0x10, 0xf0, 0x60, 0xed, 0xd8, 0xf8, 0x1e, 0x8e,
	0x0e, 0xff, 0xd1, 0xa5, 0xd1, 0xf7, 0xfb, 0x9f, 0xdb, 0xfc, 0xaf, 0x7a, 0x6d, 0xf4, 0xaf, 0x7b,
	0x7c, 0x7c, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x94, 0x35, 0x1d, 0x12, 0x75, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserSrvClient is the client API for UserSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserSrvClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*RegistryResp, error)
}

type userSrvClient struct {
	cc *grpc.ClientConn
}

func NewUserSrvClient(cc *grpc.ClientConn) UserSrvClient {
	return &userSrvClient{cc}
}

func (c *userSrvClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/sysuser.UserSrv/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvClient) Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*RegistryResp, error) {
	out := new(RegistryResp)
	err := c.cc.Invoke(ctx, "/sysuser.UserSrv/Registry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSrvServer is the server API for UserSrv service.
type UserSrvServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Registry(context.Context, *RegistryReq) (*RegistryResp, error)
}

// UnimplementedUserSrvServer can be embedded to have forward compatible implementations.
type UnimplementedUserSrvServer struct {
}

func (*UnimplementedUserSrvServer) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserSrvServer) Registry(ctx context.Context, req *RegistryReq) (*RegistryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registry not implemented")
}

func RegisterUserSrvServer(s *grpc.Server, srv UserSrvServer) {
	s.RegisterService(&_UserSrv_serviceDesc, srv)
}

func _UserSrv_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysuser.UserSrv/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSrv_Registry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).Registry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sysuser.UserSrv/Registry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).Registry(ctx, req.(*RegistryReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sysuser.UserSrv",
	HandlerType: (*UserSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserSrv_Login_Handler,
		},
		{
			MethodName: "Registry",
			Handler:    _UserSrv_Registry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sysuser.proto",
}
