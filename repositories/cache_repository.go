package repositories

import (
	"bmkg/models"
	"bmkg/modules"
)

// NewCacheRepository ...
func NewCacheRepository() CacheRepositoryApi {
	return &CacheRepository{}
}

// CacheRepositoryApi ...
type CacheRepositoryApi interface {
	GetLastEarthquakeCache() (models.Earthquake, bool, error)
	SetLastEarthquakeCache(earthquake models.Earthquake) error
	GetLatestEarthquakeCache() ([]models.Earthquake, bool, error)
	SetLatestEarthquakeCache(earthquakes []models.Earthquake) error
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
