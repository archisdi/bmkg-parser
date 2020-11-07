package services

import (
	"bmkg/repositories"
	"fmt"
)

// WeatherServiceApi ...
type WeatherServiceApi interface {
}

// WeatherService ...
type WeatherService struct {
	repo repositories.WeatherRepository
}

// RetrieveLocationWeatherForecast ...
func (s *WeatherService) RetrieveLocationWeatherForecast(location string) {
	weather, _ := s.repo.GetLocationWeatherForecast(location)

	for _, area := range weather.Forecast.Area {
		fmt.Println(area.Name[0].Text)
	}

}
