package main

import (
	"log"
	"net"

	"github.com/weather-app/config"
	"github.com/weather-app/db"
	proto "github.com/weather-app/generated"
	"github.com/weather-app/monad"
	service "github.com/weather-app/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func setupGRPCServer(weatherAPIKey string) monad.Result[*grpc.Server] {
	grpcServer := grpc.NewServer()

	cacheService := service.NewCacheService(db.ConnectRedis("localhost:6379", "", 0))

	weatherService := service.NewWeatherService(weatherAPIKey, cacheService)

	proto.RegisterWeatherServiceServer(grpcServer, weatherService)
	reflection.Register(grpcServer)
	return monad.Ok(grpcServer)
}

func startGRPCServer(grpcServer *grpc.Server) monad.Result[*grpc.Server] {
	listener, err := net.Listen("tcp", "[::]:50051")
	if err != nil {
		return monad.Err[*grpc.Server](err)
	}

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		return monad.Err[*grpc.Server](err)
	}

	return monad.Ok(grpcServer)
}

func main() {
	config.LoadEnv()

	weatherAPIKey := config.GetWeatherAPIKey()
	if weatherAPIKey == "" {
		log.Fatalf("Weather API key is missing in the environment")
	}

	result := setupGRPCServer(weatherAPIKey).Bind(func(grpcServer *grpc.Server) monad.Result[*grpc.Server] {
		return startGRPCServer(grpcServer)
	})

	if result.Err != nil {
		log.Fatalf("Failed to start gRPC server: %v", result.Err)
	}
}
