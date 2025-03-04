package services

import (
	"context"
	"time"

	"github.com/weather-app/generated"
	"github.com/weather-app/util"
)

type weatherServiceServer struct {
	generated.UnimplementedWeatherServiceServer
	apiKey       string
	cacheService CacheService
}

func NewWeatherService(apiKey string, cacheService CacheService) generated.WeatherServiceServer {
	return &weatherServiceServer{
		apiKey:       apiKey,
		cacheService: cacheService,
	}
}

// Pure
func buildForecastQueryParams(req *generated.ForecastRequest) map[string]string {
	return map[string]string{
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
}

// Pure
func buildRealtimeWeatherQueryParams(req *generated.RealtimeWeatherRequest) map[string]string {
	return map[string]string{
		"q":    req.GetQuery(),
		"lang": req.GetLang(),
	}
}

// Impure
func (s *weatherServiceServer) GetRealtimeWeather(req *generated.RealtimeWeatherRequest, stream generated.WeatherService_GetRealtimeWeatherServer) error {
	baseURL := "https://api.weatherapi.com/v1/current.json"
	realtimeWeatherRequestParams := buildRealtimeWeatherQueryParams(req)
	url := util.BuildRequestURLWithAPIKey(baseURL, realtimeWeatherRequestParams, s.apiKey)
	for {
		ioWeather := util.FetchData[generated.RealtimeWeatherResponse](url)

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

// Impure
func (s *weatherServiceServer) GetForecastWeather(ctx context.Context, req *generated.ForecastRequest) (*generated.ForecastResponse, error) {
	baseURL := "https://api.weatherapi.com/v1/forecast.json"
	forecastQueryParams := buildForecastQueryParams(req)
	url := util.BuildRequestURLWithAPIKey(baseURL, forecastQueryParams, s.apiKey)

	cachedResp, err := GetCachedData[*generated.ForecastResponse](ctx, s.cacheService, url).Run()
	if err == nil {
		return cachedResp, nil
	}

	forecastResp, err := util.FetchData[*generated.ForecastResponse](url).Run()
	if err != nil {
		return nil, err
	}

	_, err = CacheData(ctx, s.cacheService, url, forecastResp, time.Hour).Run()
	if err != nil {
		return nil, err
	}

	return forecastResp, nil
}
