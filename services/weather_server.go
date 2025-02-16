package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	proto "github.com/weather-app/generated"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type weatherServiceServer struct {
	proto.UnimplementedWeatherServiceServer
	apiKey string
}

func NewWeatherService(apiKey string) proto.WeatherServiceServer {
	return &weatherServiceServer{apiKey: apiKey}
}

func (s *weatherServiceServer) GetRealtimeWeather(ctx context.Context, req *proto.WeatherRequest) (*proto.WeatherResponse, error) {
	baseURL := "https://api.weatherapi.com/v1/current.json"
	params := url.Values{}
	params.Add("q", req.Query)
	params.Add("lang", req.Lang)
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

	var weatherResp proto.WeatherResponse
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmarshal response: %v", err)
	}
	return &weatherResp, nil
}

func (s *weatherServiceServer) GetForecastWeather(ctx context.Context, req *proto.ForecastRequest) (*proto.ForecastResponse, error) {
	baseURL := "https://api.weatherapi.com/v1/forecast.json"
	params := url.Values{}
	params.Add("q", req.Query)
	params.Add("days", req.Days)
	if req.Dt != "" {
		params.Add("dt", req.Dt)
	}
	if req.Unixdt != "" {
		params.Add("unixdt", req.Unixdt)
	}
	if req.Hour != "" {
		params.Add("hour", req.Hour)
	}
	if req.Lang != "" {
		params.Add("lang", req.Lang)
	}
	if req.Alerts != "" {
		params.Add("alerts", req.Alerts)
	}
	if req.Aqi != "" {
		params.Add("aqi", req.Aqi)
	}
	params.Add("key", s.apiKey)

	requestURL := baseURL + "?" + params.Encode()
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var forecastResp proto.ForecastResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &forecastResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return &forecastResp, nil
}
