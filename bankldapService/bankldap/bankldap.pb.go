// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bankldap.proto

/*
Package bankldap is a generated protocol buffer package.

It is generated from these files:
	bankldap.proto

It has these top-level messages:
	Credentials
	Authenticated
	User
	Email
	UserSummary
	Health
	Empty
*/
package bankldap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Credentials struct {
	Anumber  string `protobuf:"bytes,1,opt,name=anumber" json:"anumber,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Credentials) Reset()                    { *m = Credentials{} }
func (m *Credentials) String() string            { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()               {}
func (*Credentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Credentials) GetAnumber() string {
	if m != nil {
		return m.Anumber
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Authenticated struct {
	Authenticated bool `protobuf:"varint,1,opt,name=authenticated" json:"authenticated,omitempty"`
}

func (m *Authenticated) Reset()                    { *m = Authenticated{} }
func (m *Authenticated) String() string            { return proto.CompactTextString(m) }
func (*Authenticated) ProtoMessage()               {}
func (*Authenticated) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Authenticated) GetAuthenticated() bool {
	if m != nil {
		return m.Authenticated
	}
	return false
}

type User struct {
	Anumber string `protobuf:"bytes,1,opt,name=anumber" json:"anumber,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetAnumber() string {
	if m != nil {
		return m.Anumber
	}
	return ""
}

type Email struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Email) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserSummary struct {
	Username    string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	FirstName   string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
}

func (m *UserSummary) Reset()                    { *m = UserSummary{} }
func (m *UserSummary) String() string            { return proto.CompactTextString(m) }
func (*UserSummary) ProtoMessage()               {}
func (*UserSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserSummary) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserSummary) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *UserSummary) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserSummary) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserSummary) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Health struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Health) Reset()                    { *m = Health{} }
func (m *Health) String() string            { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()               {}
func (*Health) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Health) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Credentials)(nil), "bankldap.Credentials")
	proto.RegisterType((*Authenticated)(nil), "bankldap.Authenticated")
	proto.RegisterType((*User)(nil), "bankldap.User")
	proto.RegisterType((*Email)(nil), "bankldap.Email")
	proto.RegisterType((*UserSummary)(nil), "bankldap.UserSummary")
	proto.RegisterType((*Health)(nil), "bankldap.Health")
	proto.RegisterType((*Empty)(nil), "bankldap.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BankLdap service

type BankLdapClient interface {
	Authenticate(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Authenticated, error)
	Lookup(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserSummary, error)
	Anumber(ctx context.Context, in *Email, opts ...grpc.CallOption) (*User, error)
	Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Health, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type bankLdapClient struct {
	cc *grpc.ClientConn
}

func NewBankLdapClient(cc *grpc.ClientConn) BankLdapClient {
	return &bankLdapClient{cc}
}

func (c *bankLdapClient) Authenticate(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Authenticated, error) {
	out := new(Authenticated)
	err := grpc.Invoke(ctx, "/bankldap.BankLdap/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankLdapClient) Lookup(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserSummary, error) {
	out := new(UserSummary)
	err := grpc.Invoke(ctx, "/bankldap.BankLdap/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankLdapClient) Anumber(ctx context.Context, in *Email, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/bankldap.BankLdap/Anumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankLdapClient) Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := grpc.Invoke(ctx, "/bankldap.BankLdap/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankLdapClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/bankldap.BankLdap/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BankLdap service

type BankLdapServer interface {
	Authenticate(context.Context, *Credentials) (*Authenticated, error)
	Lookup(context.Context, *User) (*UserSummary, error)
	Anumber(context.Context, *Email) (*User, error)
	Status(context.Context, *Empty) (*Health, error)
	Ping(context.Context, *Empty) (*Empty, error)
}

func RegisterBankLdapServer(s *grpc.Server, srv BankLdapServer) {
	s.RegisterService(&_BankLdap_serviceDesc, srv)
}

func _BankLdap_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankLdapServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bankldap.BankLdap/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankLdapServer).Authenticate(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankLdap_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankLdapServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bankldap.BankLdap/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankLdapServer).Lookup(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankLdap_Anumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankLdapServer).Anumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bankldap.BankLdap/Anumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankLdapServer).Anumber(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankLdap_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankLdapServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bankldap.BankLdap/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankLdapServer).Status(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankLdap_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankLdapServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bankldap.BankLdap/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankLdapServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BankLdap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bankldap.BankLdap",
	HandlerType: (*BankLdapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _BankLdap_Authenticate_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _BankLdap_Lookup_Handler,
		},
		{
			MethodName: "Anumber",
			Handler:    _BankLdap_Anumber_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _BankLdap_Status_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _BankLdap_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bankldap.proto",
}

func init() { proto.RegisterFile("bankldap.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5b, 0x4b, 0xe3, 0x40,
	0x14, 0x4e, 0xbb, 0x6d, 0x9a, 0x9e, 0x5e, 0x76, 0x19, 0xb6, 0xbb, 0x25, 0x4b, 0xa1, 0x3b, 0xf8,
	0x20, 0xa2, 0x05, 0x15, 0xdf, 0xad, 0x45, 0xf0, 0xa1, 0x88, 0xb4, 0xf8, 0x2c, 0xa7, 0x66, 0xb4,
	0xa1, 0xb9, 0x0c, 0x33, 0x13, 0x24, 0xff, 0xc0, 0x5f, 0xe1, 0x6f, 0x95, 0x99, 0xa4, 0xb9, 0x54,
	0x7c, 0xcb, 0x77, 0xe1, 0xcb, 0x39, 0xdf, 0x19, 0x18, 0x6e, 0x30, 0xda, 0x05, 0x1e, 0xf2, 0x19,
	0x17, 0xb1, 0x8a, 0x89, 0xb3, 0xc7, 0x74, 0x01, 0xbd, 0x85, 0x60, 0x1e, 0x8b, 0x94, 0x8f, 0x81,
	0x24, 0x63, 0xe8, 0x60, 0x94, 0x84, 0x1b, 0x26, 0xc6, 0x8d, 0x69, 0xe3, 0xb8, 0xbb, 0xda, 0x43,
	0xe2, 0x82, 0xc3, 0x51, 0xca, 0xb7, 0x58, 0x78, 0xe3, 0xa6, 0x91, 0x0a, 0x4c, 0xaf, 0x60, 0x30,
	0x4f, 0xd4, 0x56, 0x87, 0x3c, 0xa3, 0x62, 0x1e, 0x39, 0x82, 0x01, 0x56, 0x09, 0x13, 0xe6, 0xac,
	0xea, 0x24, 0x9d, 0x42, 0xeb, 0x51, 0x32, 0xf1, 0xfd, 0x4f, 0xe9, 0x04, 0xda, 0xb7, 0x21, 0xfa,
	0x01, 0xf9, 0x0d, 0x6d, 0xa6, 0x3f, 0x72, 0x43, 0x06, 0xe8, 0x47, 0x03, 0x7a, 0x3a, 0x61, 0x9d,
	0x84, 0x21, 0x8a, 0x54, 0xcf, 0x98, 0x48, 0x26, 0x22, 0x0c, 0x59, 0x6e, 0x2c, 0x30, 0xf9, 0x0f,
	0x7d, 0xcf, 0x97, 0x3c, 0xc0, 0xf4, 0xc9, 0xe8, 0xd9, 0x0e, 0xbd, 0x9c, 0xbb, 0xd7, 0x96, 0x09,
	0xc0, 0x8b, 0x2f, 0xa4, 0xca, 0x0c, 0x3f, 0x8c, 0xa1, 0x6b, 0x18, 0x23, 0xff, 0x83, 0x6e, 0x80,
	0x7b, 0xb5, 0x95, 0xc5, 0x6b, 0xc2, 0x88, 0xc5, 0x80, 0xed, 0xea, 0x80, 0x53, 0xb0, 0xef, 0x18,
	0x06, 0x6a, 0x4b, 0xfe, 0x80, 0x2d, 0x15, 0xaa, 0x44, 0xe6, 0x83, 0xe5, 0x88, 0x76, 0xf4, 0x86,
	0x5c, 0xa5, 0x17, 0xef, 0x4d, 0x70, 0x6e, 0x30, 0xda, 0x2d, 0x3d, 0xe4, 0xe4, 0x1a, 0xfa, 0xd5,
	0x42, 0xc9, 0x68, 0x56, 0x1c, 0xb0, 0x72, 0x2d, 0xf7, 0x6f, 0x49, 0xd7, 0xfa, 0xa7, 0x16, 0x39,
	0x07, 0x7b, 0x19, 0xc7, 0xbb, 0x84, 0x93, 0x61, 0x69, 0xd2, 0x5d, 0xb9, 0xa3, 0x3a, 0xce, 0xbb,
	0xa3, 0x16, 0x39, 0x85, 0xce, 0x3c, 0x3f, 0xf6, 0xcf, 0xd2, 0x63, 0xfa, 0x77, 0x0f, 0x42, 0xa8,
	0x45, 0xce, 0xc0, 0x5e, 0x9b, 0x15, 0xea, 0x66, 0xae, 0x52, 0xf7, 0x57, 0x49, 0x64, 0xdb, 0x53,
	0x8b, 0x9c, 0x40, 0xeb, 0xc1, 0x8f, 0x5e, 0xbf, 0x9a, 0x0f, 0x09, 0x6a, 0x6d, 0x6c, 0xf3, 0x48,
	0x2f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0xd5, 0xa7, 0xde, 0xb6, 0x02, 0x00, 0x00,
}