// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: auth.proto

package genproto

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*Register, error)
	LoginUser(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*Token, error)
	GetUser(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Register, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*Register, error)
	UpdateUserType(ctx context.Context, in *UserTypeReq, opts ...grpc.CallOption) (*UserTypeRes, error)
	GetAllUsers(ctx context.Context, in *PageLimit, opts ...grpc.CallOption) (*GetAllUsersRes, error)
	DeleteUser(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Success, error)
	ResetPassword(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*Success, error)
	RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*Token, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutRes, error)
	GetEcoPoints(ctx context.Context, in *ById, opts ...grpc.CallOption) (*EcoPointsBalance, error)
	AddEcoPoints(ctx context.Context, in *AddEcoPointsReq, opts ...grpc.CallOption) (*AddEcoPointsRes, error)
	GetEcoPointsHistory(ctx context.Context, in *PageLimit, opts ...grpc.CallOption) (*EcoPointsHistoryRes, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*Register, error) {
	out := new(Register)
	err := c.cc.Invoke(ctx, "/protos.AuthService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) LoginUser(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/protos.AuthService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUser(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Register, error) {
	out := new(Register)
	err := c.cc.Invoke(ctx, "/protos.AuthService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*Register, error) {
	out := new(Register)
	err := c.cc.Invoke(ctx, "/protos.AuthService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUserType(ctx context.Context, in *UserTypeReq, opts ...grpc.CallOption) (*UserTypeRes, error) {
	out := new(UserTypeRes)
	err := c.cc.Invoke(ctx, "/protos.AuthService/UpdateUserType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAllUsers(ctx context.Context, in *PageLimit, opts ...grpc.CallOption) (*GetAllUsersRes, error) {
	out := new(GetAllUsersRes)
	err := c.cc.Invoke(ctx, "/protos.AuthService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteUser(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/protos.AuthService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *ByEmail, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/protos.AuthService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/protos.AuthService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutRes, error) {
	out := new(LogoutRes)
	err := c.cc.Invoke(ctx, "/protos.AuthService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetEcoPoints(ctx context.Context, in *ById, opts ...grpc.CallOption) (*EcoPointsBalance, error) {
	out := new(EcoPointsBalance)
	err := c.cc.Invoke(ctx, "/protos.AuthService/GetEcoPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AddEcoPoints(ctx context.Context, in *AddEcoPointsReq, opts ...grpc.CallOption) (*AddEcoPointsRes, error) {
	out := new(AddEcoPointsRes)
	err := c.cc.Invoke(ctx, "/protos.AuthService/AddEcoPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetEcoPointsHistory(ctx context.Context, in *PageLimit, opts ...grpc.CallOption) (*EcoPointsHistoryRes, error) {
	out := new(EcoPointsHistoryRes)
	err := c.cc.Invoke(ctx, "/protos.AuthService/GetEcoPointsHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	RegisterUser(context.Context, *RegisterUserReq) (*Register, error)
	LoginUser(context.Context, *LoginReq) (*Token, error)
	GetUser(context.Context, *ById) (*Register, error)
	UpdateUser(context.Context, *UpdateUserReq) (*Register, error)
	UpdateUserType(context.Context, *UserTypeReq) (*UserTypeRes, error)
	GetAllUsers(context.Context, *PageLimit) (*GetAllUsersRes, error)
	DeleteUser(context.Context, *ById) (*Success, error)
	ResetPassword(context.Context, *ByEmail) (*Success, error)
	RefreshToken(context.Context, *RefreshTokenReq) (*Token, error)
	Logout(context.Context, *LogoutReq) (*LogoutRes, error)
	GetEcoPoints(context.Context, *ById) (*EcoPointsBalance, error)
	AddEcoPoints(context.Context, *AddEcoPointsReq) (*AddEcoPointsRes, error)
	GetEcoPointsHistory(context.Context, *PageLimit) (*EcoPointsHistoryRes, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) RegisterUser(context.Context, *RegisterUserReq) (*Register, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAuthServiceServer) LoginUser(context.Context, *LoginReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAuthServiceServer) GetUser(context.Context, *ById) (*Register, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUser(context.Context, *UpdateUserReq) (*Register, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUserType(context.Context, *UserTypeReq) (*UserTypeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserType not implemented")
}
func (UnimplementedAuthServiceServer) GetAllUsers(context.Context, *PageLimit) (*GetAllUsersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedAuthServiceServer) DeleteUser(context.Context, *ById) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *ByEmail) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) RefreshToken(context.Context, *RefreshTokenReq) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *LogoutReq) (*LogoutRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceServer) GetEcoPoints(context.Context, *ById) (*EcoPointsBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcoPoints not implemented")
}
func (UnimplementedAuthServiceServer) AddEcoPoints(context.Context, *AddEcoPointsReq) (*AddEcoPointsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEcoPoints not implemented")
}
func (UnimplementedAuthServiceServer) GetEcoPointsHistory(context.Context, *PageLimit) (*EcoPointsHistoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEcoPointsHistory not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RegisterUser(ctx, req.(*RegisterUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).LoginUser(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUser(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUserType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUserType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/UpdateUserType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUserType(ctx, req.(*UserTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageLimit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAllUsers(ctx, req.(*PageLimit))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteUser(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*ByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshToken(ctx, req.(*RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetEcoPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetEcoPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/GetEcoPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetEcoPoints(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AddEcoPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEcoPointsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AddEcoPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/AddEcoPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AddEcoPoints(ctx, req.(*AddEcoPointsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetEcoPointsHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageLimit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetEcoPointsHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.AuthService/GetEcoPointsHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetEcoPointsHistory(ctx, req.(*PageLimit))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _AuthService_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _AuthService_LoginUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _AuthService_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserType",
			Handler:    _AuthService_UpdateUserType_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _AuthService_GetAllUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthService_DeleteUser_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthService_RefreshToken_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "GetEcoPoints",
			Handler:    _AuthService_GetEcoPoints_Handler,
		},
		{
			MethodName: "AddEcoPoints",
			Handler:    _AuthService_AddEcoPoints_Handler,
		},
		{
			MethodName: "GetEcoPointsHistory",
			Handler:    _AuthService_GetEcoPointsHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
