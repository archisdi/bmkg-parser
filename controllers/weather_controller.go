package controllers

import (
	"bmkg/models"
	"bmkg/services"
)

// WeatherController ...
type WeatherController struct {
	service services.WeatherService
}

// GetWeathersBy ...
func (c *WeatherController) GetWeathersBy(coordinate string) (models.Response, error) {
	return c.service.RetrieveNationalWeatherForecast(coordinate)
}
