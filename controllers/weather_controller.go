package controllers

import (
	"bmkg/models"
	"bmkg/services"
)

// WeatherController ...
type WeatherController struct {
	service services.WeatherService
}

// GetWeather ...
func (c *WeatherController) GetWeather(request models.WeatherRequest) (models.Response, error) {
	region := request.Region
	if region == "" {
		region = "Indonesia"
	}
	return c.service.RetrieveRegionalWeatherForecast(region, request.Coordinate)
}
