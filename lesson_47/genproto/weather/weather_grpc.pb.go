// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: weather.proto

package weather

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

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	GetCurrentWeather(ctx context.Context, in *Void, opts ...grpc.CallOption) (*WeatherDaily, error)
	GetWeatherForecast(ctx context.Context, in *Date, opts ...grpc.CallOption) (*WeatherData, error)
	ReportWeatherCondition(ctx context.Context, in *WeatherDaily, opts ...grpc.CallOption) (*Response, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetCurrentWeather(ctx context.Context, in *Void, opts ...grpc.CallOption) (*WeatherDaily, error) {
	out := new(WeatherDaily)
	err := c.cc.Invoke(ctx, "/weather.WeatherService/GetCurrentWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetWeatherForecast(ctx context.Context, in *Date, opts ...grpc.CallOption) (*WeatherData, error) {
	out := new(WeatherData)
	err := c.cc.Invoke(ctx, "/weather.WeatherService/GetWeatherForecast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) ReportWeatherCondition(ctx context.Context, in *WeatherDaily, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/weather.WeatherService/ReportWeatherCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	GetCurrentWeather(context.Context, *Void) (*WeatherDaily, error)
	GetWeatherForecast(context.Context, *Date) (*WeatherData, error)
	ReportWeatherCondition(context.Context, *WeatherDaily) (*Response, error)
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) GetCurrentWeather(context.Context, *Void) (*WeatherDaily, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentWeather not implemented")
}
func (UnimplementedWeatherServiceServer) GetWeatherForecast(context.Context, *Date) (*WeatherData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeatherForecast not implemented")
}
func (UnimplementedWeatherServiceServer) ReportWeatherCondition(context.Context, *WeatherDaily) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportWeatherCondition not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_GetCurrentWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetCurrentWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/GetCurrentWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetCurrentWeather(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetWeatherForecast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetWeatherForecast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/GetWeatherForecast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetWeatherForecast(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_ReportWeatherCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherDaily)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).ReportWeatherCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/ReportWeatherCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).ReportWeatherCondition(ctx, req.(*WeatherDaily))
	}
	return interceptor(ctx, in, info, handler)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weather.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentWeather",
			Handler:    _WeatherService_GetCurrentWeather_Handler,
		},
		{
			MethodName: "GetWeatherForecast",
			Handler:    _WeatherService_GetWeatherForecast_Handler,
		},
		{
			MethodName: "ReportWeatherCondition",
			Handler:    _WeatherService_ReportWeatherCondition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weather.proto",
}
