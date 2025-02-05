package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type weatherService struct {
	apiKey string
}

func NewWeatherService(apiKey string) WeatherServiceServer {
	return &weatherService{apiKey: apiKey}
}

func (s *weatherService) GetRealtimeWeather(ctx context.Context, req *WeatherRequest) (*WeatherResponse, error) {
	baseURL := "https://api.weatherapi.com/v1/current.json"
	params := url.Values{}
	params.Add("q", req.Query)
	params.Add("lang", req.Lang)
	params.Add("key", s.apiKey)

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherResp WeatherResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &weatherResp, nil
}

func (s *weatherService) GetForecastWeather(ctx context.Context, req *ForecastRequest) (*WeatherResponse, error) {
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

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherResp WeatherResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return &weatherResp, nil
}

func (s *weatherService) mustEmbedUnimplementedWeatherServiceServer() {}
