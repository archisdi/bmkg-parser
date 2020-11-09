package controllers

import (
	"bmkg/models"
	"bmkg/services"
)

// WeatherController ...
type WeatherController struct {
	Service services.WeatherServiceAPI
}

// GetWeather ...
func (c *WeatherController) GetWeather(request models.WeatherRequest) (models.Response, error) {
	return c.Service.RetrieveRegionalWeatherForecast(request.Region, request.Coordinate)
}
