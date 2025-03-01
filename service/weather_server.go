package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	proto "github.com/weather-app/generated"
	"github.com/weather-app/monad"
	utils "github.com/weather-app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type weatherServiceServer struct {
	proto.UnimplementedWeatherServiceServer
	apiKey       string
	redisService *RedisService
}

func NewWeatherService(apiKey string, redisService *RedisService) proto.WeatherServiceServer {
	return &weatherServiceServer{
		apiKey:       apiKey,
		redisService: redisService,
	}
}

// Pure function to fetch weather data from the API
func fetchWeatherData[T any](url string) monad.IO[T] {
	return monad.IO[T]{Run: func() (T, error) {
		var target T
		resp, err := http.Get(url)
		if err != nil {
			return target, status.Errorf(codes.Internal, "failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return target, status.Errorf(codes.Internal, "unexpected status code: %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return target, status.Errorf(codes.Internal, "failed to read response body: %v", err)
		}

		if err := json.Unmarshal(body, &target); err != nil {
			return target, status.Errorf(codes.Internal, "failed to unmarshal response: %v", err)
		}

		return target, nil
	}}
}

func (s *weatherServiceServer) GetRealtimeWeather(req *proto.RealtimeWeatherRequest, stream proto.WeatherService_GetRealtimeWeatherServer) error {
	baseURL := "https://api.weatherapi.com/v1/current.json"
	params := map[string]string{
		"q":    req.Query,
		"lang": req.Lang,
	}

	for {
		url := utils.BuildRequestURL(baseURL, params, s.apiKey)

		// Wrap in IO Monad
		ioWeather := fetchWeatherData[proto.RealtimeWeatherResponse](url)

		// Execute and handle errors
		weatherResp, err := ioWeather.Run()
		if err != nil {
			return err
		}

		if err := stream.Send(&weatherResp); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}
}

func (s *weatherServiceServer) GetForecastWeather(ctx context.Context, req *proto.ForecastRequest) (*proto.ForecastResponse, error) {
	baseURL := "https://api.weatherapi.com/v1/forecast.json"
	params := map[string]string{
		"q":      req.GetQuery(),
		"days":   req.GetDays(),
		"dt":     req.GetDt(),
		"unixdt": req.GetUnixdt(),
		"hour":   req.GetHour(),
		"lang":   req.GetLang(),
		"alerts": req.GetAlerts(),
		"aqi":    req.GetAqi(),
		"tp":     req.GetTp(),
	}

	url := utils.BuildRequestURL(baseURL, params, s.apiKey)

	// Wrap in IO Monad
	ioForecast := fetchWeatherData[proto.ForecastResponse](url)

	// Execute
	forecastResp, err := ioForecast.Run()
	if err != nil {
		return nil, err
	}

	return &forecastResp, nil
}
