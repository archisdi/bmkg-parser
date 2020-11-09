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
	region := request.Region
	if region == "" {
		region = "Indonesia"
	}
	return c.Service.RetrieveRegionalWeatherForecast(region, request.Coordinate)
}
