package services

import (
	"bmkg/models"
	"bmkg/repositories"
	"bmkg/utils"
	"errors"
	"math"
)

// NewWeatherService ...
func NewWeatherService(
	repo repositories.WeatherRepositoryAPI,
	cache repositories.CacheRepositoryAPI,
	) WeatherServiceAPI {
	return &WeatherService{
		Repo:  repo,
		Cache: cache,
	}
}

// WeatherServiceAPI ...
type WeatherServiceAPI interface {
	RetrieveRegionalWeatherForecast(region string, baseCoordinate string) (models.WeatherOutput, error)
}

// WeatherService ...
type WeatherService struct {
	Repo repositories.WeatherRepositoryAPI
	Cache repositories.CacheRepositoryAPI
}

// RetrieveNationalWeatherForecast ...
func (s *WeatherService) RetrieveRegionalWeatherForecast(region string, baseCoordinate string) (models.WeatherOutput, error) {
	var output models.WeatherOutput

	if region == "" {
		region = "Indonesia"
	}

	// parse base coordinate
	xA, xB, coorErr := utils.StringToCoordinate(baseCoordinate)
	if coorErr != nil {
		return output, coorErr
	}

	// retrieve weather data
	var weather models.BaseWeather
	if cacheWeather, ok, _ := s.Cache.GetRegionWeatherCache(region); ok {
		weather = cacheWeather
	} else if sourceWeather, sourceErr := s.Repo.GetWeatherForecast(region); sourceErr != nil {
		return output, errors.New("invalid region name")
	} else {
		weather = sourceWeather
		if cacheErr := s.Cache.SetRegionWeatherCache(region, weather); cacheErr != nil {
			return output, cacheErr
		}
	}

	// calculate distance to determine closest location

	var currentArea models.Area
	currentDistance := math.MaxFloat64

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
		Data:  currentArea.GetWeather(),
	}

	return output, nil
}
