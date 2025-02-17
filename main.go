package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	proto "github.com/weather-app/generated"
	"github.com/weather-app/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func startGRPCServer() {
	grpcServer := grpc.NewServer()
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatalf("WEATHER_API_KEY is not set in the environment variables")
	}

	weatherService := services.NewWeatherService(apiKey)
	proto.RegisterWeatherServiceServer(grpcServer, weatherService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "[::]:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	loadEnv()
	startGRPCServer()
}
