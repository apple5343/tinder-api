// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: swipe.proto

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
	SwipeService_Swipe_FullMethodName          = "/swipe_v1.SwipeService/Swipe"
	SwipeService_GetSwipes_FullMethodName      = "/swipe_v1.SwipeService/GetSwipes"
	SwipeService_GetSwipedUsers_FullMethodName = "/swipe_v1.SwipeService/GetSwipedUsers"
	SwipeService_GetLiked_FullMethodName       = "/swipe_v1.SwipeService/GetLiked"
)

// SwipeServiceClient is the client API for SwipeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SwipeServiceClient interface {
	Swipe(ctx context.Context, in *SwipeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSwipes(ctx context.Context, in *GetSwipesRequest, opts ...grpc.CallOption) (*GetSwipesResponse, error)
	GetSwipedUsers(ctx context.Context, in *GetSwipedUsersRequest, opts ...grpc.CallOption) (*GetSwipedUsersResponse, error)
	GetLiked(ctx context.Context, in *GetLikedRequest, opts ...grpc.CallOption) (*GetLikedResponse, error)
}

type swipeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSwipeServiceClient(cc grpc.ClientConnInterface) SwipeServiceClient {
	return &swipeServiceClient{cc}
}

func (c *swipeServiceClient) Swipe(ctx context.Context, in *SwipeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SwipeService_Swipe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swipeServiceClient) GetSwipes(ctx context.Context, in *GetSwipesRequest, opts ...grpc.CallOption) (*GetSwipesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSwipesResponse)
	err := c.cc.Invoke(ctx, SwipeService_GetSwipes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swipeServiceClient) GetSwipedUsers(ctx context.Context, in *GetSwipedUsersRequest, opts ...grpc.CallOption) (*GetSwipedUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSwipedUsersResponse)
	err := c.cc.Invoke(ctx, SwipeService_GetSwipedUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swipeServiceClient) GetLiked(ctx context.Context, in *GetLikedRequest, opts ...grpc.CallOption) (*GetLikedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLikedResponse)
	err := c.cc.Invoke(ctx, SwipeService_GetLiked_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwipeServiceServer is the server API for SwipeService service.
// All implementations must embed UnimplementedSwipeServiceServer
// for forward compatibility.
type SwipeServiceServer interface {
	Swipe(context.Context, *SwipeRequest) (*emptypb.Empty, error)
	GetSwipes(context.Context, *GetSwipesRequest) (*GetSwipesResponse, error)
	GetSwipedUsers(context.Context, *GetSwipedUsersRequest) (*GetSwipedUsersResponse, error)
	GetLiked(context.Context, *GetLikedRequest) (*GetLikedResponse, error)
	mustEmbedUnimplementedSwipeServiceServer()
}

// UnimplementedSwipeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSwipeServiceServer struct{}

func (UnimplementedSwipeServiceServer) Swipe(context.Context, *SwipeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Swipe not implemented")
}
func (UnimplementedSwipeServiceServer) GetSwipes(context.Context, *GetSwipesRequest) (*GetSwipesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwipes not implemented")
}
func (UnimplementedSwipeServiceServer) GetSwipedUsers(context.Context, *GetSwipedUsersRequest) (*GetSwipedUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwipedUsers not implemented")
}
func (UnimplementedSwipeServiceServer) GetLiked(context.Context, *GetLikedRequest) (*GetLikedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiked not implemented")
}
func (UnimplementedSwipeServiceServer) mustEmbedUnimplementedSwipeServiceServer() {}
func (UnimplementedSwipeServiceServer) testEmbeddedByValue()                      {}

// UnsafeSwipeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SwipeServiceServer will
// result in compilation errors.
type UnsafeSwipeServiceServer interface {
	mustEmbedUnimplementedSwipeServiceServer()
}

func RegisterSwipeServiceServer(s grpc.ServiceRegistrar, srv SwipeServiceServer) {
	// If the following call pancis, it indicates UnimplementedSwipeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SwipeService_ServiceDesc, srv)
}

func _SwipeService_Swipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwipeServiceServer).Swipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwipeService_Swipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwipeServiceServer).Swipe(ctx, req.(*SwipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwipeService_GetSwipes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSwipesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwipeServiceServer).GetSwipes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwipeService_GetSwipes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwipeServiceServer).GetSwipes(ctx, req.(*GetSwipesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwipeService_GetSwipedUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSwipedUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwipeServiceServer).GetSwipedUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwipeService_GetSwipedUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwipeServiceServer).GetSwipedUsers(ctx, req.(*GetSwipedUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwipeService_GetLiked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwipeServiceServer).GetLiked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwipeService_GetLiked_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwipeServiceServer).GetLiked(ctx, req.(*GetLikedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SwipeService_ServiceDesc is the grpc.ServiceDesc for SwipeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SwipeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "swipe_v1.SwipeService",
	HandlerType: (*SwipeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Swipe",
			Handler:    _SwipeService_Swipe_Handler,
		},
		{
			MethodName: "GetSwipes",
			Handler:    _SwipeService_GetSwipes_Handler,
		},
		{
			MethodName: "GetSwipedUsers",
			Handler:    _SwipeService_GetSwipedUsers_Handler,
		},
		{
			MethodName: "GetLiked",
			Handler:    _SwipeService_GetLiked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "swipe.proto",
}
