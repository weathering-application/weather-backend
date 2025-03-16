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
	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Initialize Redis service
	cacheService := service.NewCacheService(db.ConnectRedis("redis.weather.svc.cluster.local", "", 0))

	// Initialize weather service
	weatherService := service.NewWeatherService(weatherAPIKey, cacheService)

	// Register services
	proto.RegisterWeatherServiceServer(grpcServer, weatherService)
	reflection.Register(grpcServer)
	// Return Result wrapping the grpcServer
	return monad.Ok(grpcServer)
}

func startGRPCServer(grpcServer *grpc.Server) monad.Result[*grpc.Server] {
	// Try to listen on TCP port 50051
	listener, err := net.Listen("tcp", "[::]:50051")
	if err != nil {
		// Return error result if listener fails
		return monad.Err[*grpc.Server](err)
	}

	log.Println("gRPC server running on port 50051")
	// Start the gRPC server and check for errors
	if err := grpcServer.Serve(listener); err != nil {
		// Return error result if server fails to serve
		return monad.Err[*grpc.Server](err)
	}

	// Return success result when server starts successfully
	return monad.Ok(grpcServer)
}

func main() {
	// Load environment variables
	config.LoadEnv()

	// Get the weather API key from environment
	weatherAPIKey := config.GetWeatherAPIKey()
	if weatherAPIKey == "" {
		log.Fatalf("Weather API key is missing in the environment")
	}

	// Setup the gRPC server and handle the result using Bind for chaining
	result := setupGRPCServer(weatherAPIKey).Bind(func(grpcServer *grpc.Server) monad.Result[*grpc.Server] {
		// Start the gRPC server and return the result
		return startGRPCServer(grpcServer)
	})

	// Check if there was an error in starting the gRPC server
	if result.Err != nil {
		log.Fatalf("Failed to start gRPC server: %v", result.Err)
	}
}
