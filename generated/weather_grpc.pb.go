// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: protos/weather.proto

package generated

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
	GetRealtimeWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (WeatherService_GetRealtimeWeatherClient, error)
	GetForecastWeather(ctx context.Context, in *ForecastRequest, opts ...grpc.CallOption) (WeatherService_GetForecastWeatherClient, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) GetRealtimeWeather(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (WeatherService_GetRealtimeWeatherClient, error) {
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[0], "/weather.WeatherService/GetRealtimeWeather", opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherServiceGetRealtimeWeatherClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WeatherService_GetRealtimeWeatherClient interface {
	Recv() (*WeatherResponse, error)
	grpc.ClientStream
}

type weatherServiceGetRealtimeWeatherClient struct {
	grpc.ClientStream
}

func (x *weatherServiceGetRealtimeWeatherClient) Recv() (*WeatherResponse, error) {
	m := new(WeatherResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *weatherServiceClient) GetForecastWeather(ctx context.Context, in *ForecastRequest, opts ...grpc.CallOption) (WeatherService_GetForecastWeatherClient, error) {
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[1], "/weather.WeatherService/GetForecastWeather", opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherServiceGetForecastWeatherClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WeatherService_GetForecastWeatherClient interface {
	Recv() (*ForecastResponse, error)
	grpc.ClientStream
}

type weatherServiceGetForecastWeatherClient struct {
	grpc.ClientStream
}

func (x *weatherServiceGetForecastWeatherClient) Recv() (*ForecastResponse, error) {
	m := new(ForecastResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	GetRealtimeWeather(*WeatherRequest, WeatherService_GetRealtimeWeatherServer) error
	GetForecastWeather(*ForecastRequest, WeatherService_GetForecastWeatherServer) error
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) GetRealtimeWeather(*WeatherRequest, WeatherService_GetRealtimeWeatherServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRealtimeWeather not implemented")
}
func (UnimplementedWeatherServiceServer) GetForecastWeather(*ForecastRequest, WeatherService_GetForecastWeatherServer) error {
	return status.Errorf(codes.Unimplemented, "method GetForecastWeather not implemented")
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

func _WeatherService_GetRealtimeWeather_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WeatherRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServiceServer).GetRealtimeWeather(m, &weatherServiceGetRealtimeWeatherServer{stream})
}

type WeatherService_GetRealtimeWeatherServer interface {
	Send(*WeatherResponse) error
	grpc.ServerStream
}

type weatherServiceGetRealtimeWeatherServer struct {
	grpc.ServerStream
}

func (x *weatherServiceGetRealtimeWeatherServer) Send(m *WeatherResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WeatherService_GetForecastWeather_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ForecastRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServiceServer).GetForecastWeather(m, &weatherServiceGetForecastWeatherServer{stream})
}

type WeatherService_GetForecastWeatherServer interface {
	Send(*ForecastResponse) error
	grpc.ServerStream
}

type weatherServiceGetForecastWeatherServer struct {
	grpc.ServerStream
}

func (x *weatherServiceGetForecastWeatherServer) Send(m *ForecastResponse) error {
	return x.ServerStream.SendMsg(m)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weather.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRealtimeWeather",
			Handler:       _WeatherService_GetRealtimeWeather_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetForecastWeather",
			Handler:       _WeatherService_GetForecastWeather_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/weather.proto",
}
