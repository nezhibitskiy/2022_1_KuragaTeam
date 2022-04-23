// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: authorization.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationClient interface {
	SignUp(ctx context.Context, in *SignUpData, opts ...grpc.CallOption) (*Cookie, error)
	LogIn(ctx context.Context, in *LogInData, opts ...grpc.CallOption) (*Cookie, error)
	LogOut(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Empty, error)
	CheckAuthorization(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*UserID, error)
}

type authorizationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationClient(cc grpc.ClientConnInterface) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) SignUp(ctx context.Context, in *SignUpData, opts ...grpc.CallOption) (*Cookie, error) {
	out := new(Cookie)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) LogIn(ctx context.Context, in *LogInData, opts ...grpc.CallOption) (*Cookie, error) {
	out := new(Cookie)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) LogOut(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/LogOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) CheckAuthorization(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/CheckAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
// All implementations should embed UnimplementedAuthorizationServer
// for forward compatibility
type AuthorizationServer interface {
	SignUp(context.Context, *SignUpData) (*Cookie, error)
	LogIn(context.Context, *LogInData) (*Cookie, error)
	LogOut(context.Context, *Cookie) (*Empty, error)
	CheckAuthorization(context.Context, *Cookie) (*UserID, error)
}

// UnimplementedAuthorizationServer should be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (UnimplementedAuthorizationServer) SignUp(context.Context, *SignUpData) (*Cookie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthorizationServer) LogIn(context.Context, *LogInData) (*Cookie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedAuthorizationServer) LogOut(context.Context, *Cookie) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedAuthorizationServer) CheckAuthorization(context.Context, *Cookie) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthorization not implemented")
}

// UnsafeAuthorizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServer will
// result in compilation errors.
type UnsafeAuthorizationServer interface {
	mustEmbedUnimplementedAuthorizationServer()
}

func RegisterAuthorizationServer(s grpc.ServiceRegistrar, srv AuthorizationServer) {
	s.RegisterService(&Authorization_ServiceDesc, srv)
}

func _Authorization_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).SignUp(ctx, req.(*SignUpData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).LogIn(ctx, req.(*LogInData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/LogOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).LogOut(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_CheckAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).CheckAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/CheckAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).CheckAuthorization(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

// Authorization_ServiceDesc is the grpc.ServiceDesc for Authorization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authorization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Authorization_SignUp_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _Authorization_LogIn_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _Authorization_LogOut_Handler,
		},
		{
			MethodName: "CheckAuthorization",
			Handler:    _Authorization_CheckAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization.proto",
}
