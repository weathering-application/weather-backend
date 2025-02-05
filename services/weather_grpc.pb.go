// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: weather.proto

package services

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
	GetRealtimeWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error)
	GetForecastWeather(ctx context.Context, in *ForecastRequest, opts ...grpc.CallOption) (*WeatherResponse, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetRealtimeWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error) {
	out := new(WeatherResponse)
	err := c.cc.Invoke(ctx, "/weather.WeatherService/GetRealtimeWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherServiceClient) GetForecastWeather(ctx context.Context, in *ForecastRequest, opts ...grpc.CallOption) (*WeatherResponse, error) {
	out := new(WeatherResponse)
	err := c.cc.Invoke(ctx, "/weather.WeatherService/GetForecastWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	GetRealtimeWeather(context.Context, *WeatherRequest) (*WeatherResponse, error)
	GetForecastWeather(context.Context, *ForecastRequest) (*WeatherResponse, error)
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) GetRealtimeWeather(context.Context, *WeatherRequest) (*WeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealtimeWeather not implemented")
}
func (UnimplementedWeatherServiceServer) GetForecastWeather(context.Context, *ForecastRequest) (*WeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForecastWeather not implemented")
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

func _WeatherService_GetRealtimeWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetRealtimeWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/GetRealtimeWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetRealtimeWeather(ctx, req.(*WeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherService_GetForecastWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForecastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).GetForecastWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/GetForecastWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).GetForecastWeather(ctx, req.(*ForecastRequest))
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
			MethodName: "GetRealtimeWeather",
			Handler:    _WeatherService_GetRealtimeWeather_Handler,
		},
		{
			MethodName: "GetForecastWeather",
			Handler:    _WeatherService_GetForecastWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weather.proto",
}
