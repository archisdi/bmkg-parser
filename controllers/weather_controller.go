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
func (c *WeatherController) GetWeathersBy(location string) (models.Response, error) {

	c.service.RetrieveLocationWeatherForecast("DKIJakarta")

	return map[string]string{
		"messsage": "WOW",
	}, nil
}
