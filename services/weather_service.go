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
func (s *WeatherService) RetrieveNationalWeatherForecast(baseCoordinate string) (models.WeatherOutput, error) {
	var output models.WeatherOutput

	// parse base coordinate
	xA, xB, coorErr := utils.StringToCoordinate(baseCoordinate)
	if coorErr != nil {
		return output, coorErr
	}

	// get weather data from BMKG
	weather, _ := s.repo.GetWeatherForecast("Indonesia")

	var currentArea models.Area
	currentDistance := math.MaxFloat64

	// calculate distance to determine closest location
	for _, area := range weather.Forecast.Area {
		if yA, yB, coorErr := utils.StringToCoordinate(area.GetCoordinates()); coorErr != nil {
			return output, coorErr
		} else {
			distance := utils.CalculateEuclideanDistance(xA, xB, yA, yB)
			if distance < currentDistance {
				currentDistance = distance
				currentArea = area
			}
		}
	}

	// parse data
	output = models.WeatherOutput{
		Location: models.GeoLocation{
			Name:       currentArea.GetName(),
			Coordinate: currentArea.GetCoordinates(),
		},
		Weather:  currentArea.GetWeather(),
	}

	return output, nil
}
