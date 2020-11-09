package repositories

import (
	"bmkg/models"
	"bmkg/modules"
)

// NewCacheRepository ...
func NewCacheRepository() CacheRepositoryAPI {
	return &CacheRepository{}
}

// CacheRepositoryAPI ...
type CacheRepositoryAPI interface {
	GetLastEarthquakeCache() (models.Earthquake, bool, error)
	SetLastEarthquakeCache(earthquake models.Earthquake) error
	GetLatestEarthquakeCache() ([]models.Earthquake, bool, error)
	SetLatestEarthquakeCache(earthquakes []models.Earthquake) error
	GetRegionWeatherCache(region string) (models.BaseWeather, bool, error)
	SetRegionWeatherCache(region string, weather models.BaseWeather) error
}

// CacheRepository ...
type CacheRepository struct{}

// GetLastEarthquakeCache ...
func (*CacheRepository) GetLastEarthquakeCache() (models.Earthquake, bool, error) {
	var earthquake models.Earthquake

	if err := modules.Redis.GetCache("last_earthquake", &earthquake); err != nil {
		return earthquake, false, err
	}

	return earthquake, earthquake != (models.Earthquake{}), nil
}

// SetLastEarthquakeCache ...
func (*CacheRepository) SetLastEarthquakeCache(earthquake models.Earthquake) error {
	return modules.Redis.SetCache("last_earthquake", earthquake)
}

// GetLatestEarthquakeCache ...
func (*CacheRepository) GetLatestEarthquakeCache() ([]models.Earthquake, bool, error) {
	var earthquakes []models.Earthquake

	if err := modules.Redis.GetCache("latest_earthquake", &earthquakes); err != nil {
		return earthquakes, false, err
	}

	return earthquakes, len(earthquakes) != 0, nil
}

// SetLatestEarthquakeCache ...
func (*CacheRepository) SetLatestEarthquakeCache(earthquakes []models.Earthquake) error {
	return modules.Redis.SetCache("latest_earthquake", earthquakes)
}

// GetRegionWeatherCache ...
func (*CacheRepository) GetRegionWeatherCache(region string) (models.BaseWeather, bool, error) {
	var weather models.BaseWeather

	if err := modules.Redis.GetCache("weather_" + region, &weather); err != nil {
		return weather, false, err
	}

	return weather, weather.Source != "", nil
}

// SetRegionWeatherCache ...
func (*CacheRepository) SetRegionWeatherCache(region string, weather models.BaseWeather) error {
	return modules.Redis.SetCache("weather_" + region, weather)
}