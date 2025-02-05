package main

import (
	"log"
	"net"
	"os"

	"github.com/weather-app/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startGRPCServer() {
	grpcServer := grpc.NewServer()
	apiKey := os.Getenv("WEATHER_API_KEY")
	weatherService := services.NewWeatherService(apiKey)
	services.RegisterWeatherServiceServer(grpcServer, weatherService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func setupHandlers() handlers.WeatherHandler {
// 	apiKey := os.Getenv("WEATHER_API_KEY")
// 	weatherService := services.NewWeatherService(apiKey)
// 	return handlers.NewWeatherHandler(weatherService)
// }

// func startHTTPServer() {
// 	router := gin.Default()
// 	weatherHandler := setupHandlers()
// 	router.GET("/weathers/current", weatherHandler.GetRealtimeWeather)
// 	router.GET("/weathers/forecast", weatherHandler.GetForecastWeather)
// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatalf("failed to run HTTP server: %v", err)
// 	}
// }

func main() {
	startGRPCServer()
	// startHTTPServer()
}
