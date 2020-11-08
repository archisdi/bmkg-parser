package services

import (
	"bmkg/models"
	"bmkg/repositories"
	"bmkg/utils"
	"math"
)

// WeatherServiceAPI ...
type WeatherServiceAPI interface {
}

// WeatherService ...
type WeatherService struct {
	repo repositories.WeatherRepository
}

// RetrieveNationalWeatherForecast ...
func (s *WeatherService) RetrieveNationalWeatherForecast(baseCoordinate string) []models.Weather {
	weather, _ := s.repo.GetWeatherForecast("Indonesia")

	var currentArea models.Area
	currentDistance := math.MaxFloat64

	// calculate distance to determine closest location
	for _, area := range weather.Forecast.Area {
		coordinate := area.GetCoordinates()
		distance := utils.CalculateEuclideanDistance(baseCoordinate, coordinate)

		if distance < currentDistance {
			currentDistance = distance
			currentArea = area
		}

	}

	return currentArea.GetWeather()
}
