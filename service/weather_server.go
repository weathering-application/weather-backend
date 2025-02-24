package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	proto "github.com/weather-app/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type weatherServiceServer struct {
	proto.UnimplementedWeatherServiceServer
	apiKey string
}

func NewWeatherService(apiKey string) proto.WeatherServiceServer {
	return &weatherServiceServer{apiKey: apiKey}
}

func (s *weatherServiceServer) GetRealtimeWeather(req *proto.RealtimeWeatherRequest, stream proto.WeatherService_GetRealtimeWeatherServer) error {
	baseURL := "https://api.weatherapi.com/v1/current.json"
	params := url.Values{}
	params.Add("q", req.Query)
	params.Add("lang", req.Lang)
	params.Add("key", s.apiKey)

	requestURL := baseURL + "?" + params.Encode()

	for {
		fmt.Println("GetRealtimeWeather")

		resp, err := http.Get(requestURL)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return status.Errorf(codes.Internal, "unexpected status code: %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to read response body: %v", err)
		}

		var weatherResp proto.RealtimeWeatherResponse
		if err := json.Unmarshal(body, &weatherResp); err != nil {
			return status.Errorf(codes.Internal, "failed to unmarshal response: %v", err)
		}
		if err := stream.Send(&weatherResp); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *weatherServiceServer) GetForecastWeather(ctx context.Context, req *proto.ForecastRequest) (*proto.ForecastResponse, error) {
	baseURL := "https://api.weatherapi.com/v1/forecast.json"
	params := url.Values{}
	params.Add("q", req.GetQuery())
	params.Add("days", req.GetDays())
	params.Add("key", s.apiKey)

	requestURL := baseURL + "?" + params.Encode()

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.Internal, "unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read response body: %v", err)
	}

	var forecastResp proto.ForecastResponse
	if err := json.Unmarshal(body, &forecastResp); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal response: %v", err)
	}

	return &forecastResp, nil
}
