package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/weather-app/services"
)

type WeatherHandler interface {
	GetRealtimeWeather(*gin.Context)
	GetForecastWeather(*gin.Context)
}

type weatherHandler struct {
	service services.WeatherServiceServer
}

func NewWeatherHandler(service services.WeatherServiceServer) WeatherHandler {
	return &weatherHandler{service: service}
}

func (h *weatherHandler) GetRealtimeWeather(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		c.JSON(400, nil)
		return
	}

	lang := c.DefaultQuery("lang", "en")

	weatherRequest := &services.WeatherRequest{
		Query: q,
		Lang:  lang,
	}
	response, err := h.service.GetRealtimeWeather(c.Request.Context(), weatherRequest)
	if err != nil {
		c.JSON(500, nil)
		return
	}

	c.JSON(200, response)
}

func (h *weatherHandler) GetForecastWeather(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	if q == "" {
		c.JSON(400, nil)
		return
	}

	days := c.DefaultQuery("days", "")
	if days == "" {
		c.JSON(400, nil)
		return
	}

	dt := c.DefaultQuery("dt", "")
	unixdt := c.DefaultQuery("unixdt", "")
	hour := c.DefaultQuery("hour", "")
	lang := c.DefaultQuery("lang", "en")
	alerts := c.DefaultQuery("alerts", "")
	aqi := c.DefaultQuery("aqi", "")

	weatherRequest := &services.ForecastRequest{
		Query:  q,
		Days:   days,
		Dt:     dt,
		Unixdt: unixdt,
		Hour:   hour,
		Lang:   lang,
		Alerts: alerts,
		Aqi:    aqi,
	}
	response, err := h.service.GetForecastWeather(c.Request.Context(), weatherRequest)
	if err != nil {
		c.JSON(500, nil)
		return
	}
	c.JSON(200, response)
}
