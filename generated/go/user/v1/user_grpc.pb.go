// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_Login_FullMethodName           = "/user_v1.AuthService/Login"
	AuthService_Register_FullMethodName        = "/user_v1.AuthService/Register"
	AuthService_GetRefreshToken_FullMethodName = "/user_v1.AuthService/GetRefreshToken"
	AuthService_GetAccessToken_FullMethodName  = "/user_v1.AuthService/GetAccessToken"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetRefreshToken(ctx context.Context, in *GetRefreshTokenRequest, opts ...grpc.CallOption) (*GetRefreshTokenResponse, error)
	GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, AuthService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetRefreshToken(ctx context.Context, in *GetRefreshTokenRequest, opts ...grpc.CallOption) (*GetRefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRefreshTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_GetRefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAccessTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_GetAccessToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetRefreshToken(context.Context, *GetRefreshTokenRequest) (*GetRefreshTokenResponse, error)
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) GetRefreshToken(context.Context, *GetRefreshTokenRequest) (*GetRefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefreshToken not implemented")
}
func (UnimplementedAuthServiceServer) GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetRefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetRefreshToken(ctx, req.(*GetRefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAccessToken(ctx, req.(*GetAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "GetRefreshToken",
			Handler:    _AuthService_GetRefreshToken_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _AuthService_GetAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

const (
	UserService_GetUser_FullMethodName               = "/user_v1.UserService/GetUser"
	UserService_GetProfile_FullMethodName            = "/user_v1.UserService/GetProfile"
	UserService_GetPublicProfile_FullMethodName      = "/user_v1.UserService/GetPublicProfile"
	UserService_GetBatchProfiles_FullMethodName      = "/user_v1.UserService/GetBatchProfiles"
	UserService_UpdateProfile_FullMethodName         = "/user_v1.UserService/UpdateProfile"
	UserService_UpdateUserPreferences_FullMethodName = "/user_v1.UserService/UpdateUserPreferences"
	UserService_AddAvatar_FullMethodName             = "/user_v1.UserService/AddAvatar"
	UserService_DeleteAvatar_FullMethodName          = "/user_v1.UserService/DeleteAvatar"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	GetPublicProfile(ctx context.Context, in *GetPublicProfileRequest, opts ...grpc.CallOption) (*GetPublicProfileResponse, error)
	GetBatchProfiles(ctx context.Context, in *GetBatchProfilesRequest, opts ...grpc.CallOption) (*GetBatchProfilesResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
	UpdateUserPreferences(ctx context.Context, in *UpdateUserPreferencesRequest, opts ...grpc.CallOption) (*UpdateUserPreferencesResponse, error)
	AddAvatar(ctx context.Context, in *AddAvatarRequest, opts ...grpc.CallOption) (*AddAvatarResponse, error)
	DeleteAvatar(ctx context.Context, in *DeleteAvatarRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, UserService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPublicProfile(ctx context.Context, in *GetPublicProfileRequest, opts ...grpc.CallOption) (*GetPublicProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPublicProfileResponse)
	err := c.cc.Invoke(ctx, UserService_GetPublicProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetBatchProfiles(ctx context.Context, in *GetBatchProfilesRequest, opts ...grpc.CallOption) (*GetBatchProfilesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBatchProfilesResponse)
	err := c.cc.Invoke(ctx, UserService_GetBatchProfiles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProfileResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserPreferences(ctx context.Context, in *UpdateUserPreferencesRequest, opts ...grpc.CallOption) (*UpdateUserPreferencesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserPreferencesResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUserPreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddAvatar(ctx context.Context, in *AddAvatarRequest, opts ...grpc.CallOption) (*AddAvatarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddAvatarResponse)
	err := c.cc.Invoke(ctx, UserService_AddAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteAvatar(ctx context.Context, in *DeleteAvatarRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	GetPublicProfile(context.Context, *GetPublicProfileRequest) (*GetPublicProfileResponse, error)
	GetBatchProfiles(context.Context, *GetBatchProfilesRequest) (*GetBatchProfilesResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	UpdateUserPreferences(context.Context, *UpdateUserPreferencesRequest) (*UpdateUserPreferencesResponse, error)
	AddAvatar(context.Context, *AddAvatarRequest) (*AddAvatarResponse, error)
	DeleteAvatar(context.Context, *DeleteAvatarRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServiceServer) GetPublicProfile(context.Context, *GetPublicProfileRequest) (*GetPublicProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicProfile not implemented")
}
func (UnimplementedUserServiceServer) GetBatchProfiles(context.Context, *GetBatchProfilesRequest) (*GetBatchProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatchProfiles not implemented")
}
func (UnimplementedUserServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserPreferences(context.Context, *UpdateUserPreferencesRequest) (*UpdateUserPreferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPreferences not implemented")
}
func (UnimplementedUserServiceServer) AddAvatar(context.Context, *AddAvatarRequest) (*AddAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAvatar not implemented")
}
func (UnimplementedUserServiceServer) DeleteAvatar(context.Context, *DeleteAvatarRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAvatar not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPublicProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPublicProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetPublicProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPublicProfile(ctx, req.(*GetPublicProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetBatchProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatchProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetBatchProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetBatchProfiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetBatchProfiles(ctx, req.(*GetBatchProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserPreferences(ctx, req.(*UpdateUserPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddAvatar(ctx, req.(*AddAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteAvatar(ctx, req.(*DeleteAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _UserService_GetProfile_Handler,
		},
		{
			MethodName: "GetPublicProfile",
			Handler:    _UserService_GetPublicProfile_Handler,
		},
		{
			MethodName: "GetBatchProfiles",
			Handler:    _UserService_GetBatchProfiles_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserService_UpdateProfile_Handler,
		},
		{
			MethodName: "UpdateUserPreferences",
			Handler:    _UserService_UpdateUserPreferences_Handler,
		},
		{
			MethodName: "AddAvatar",
			Handler:    _UserService_AddAvatar_Handler,
		},
		{
			MethodName: "DeleteAvatar",
			Handler:    _UserService_DeleteAvatar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
