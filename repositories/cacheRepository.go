package repositories

import (
	"bmkg/models"
	"bmkg/modules"
)

// CacheRepository ...
type CacheRepository struct{}

// GetLastEarthquakeCache ...
func (*CacheRepository) GetLastEarthquakeCache() (models.Earthquake, bool, error) {
	var earthquake models.Earthquake
	modules.Redis.GetCache("last_earthquake", &earthquake)
	return earthquake, earthquake != (models.Earthquake{}), nil
}

// SetLastEarthquakeCache ...
func (*CacheRepository) SetLastEarthquakeCache(earthquake models.Earthquake) error {
	return modules.Redis.SetCache("last_earthquake", earthquake)
}

// GetLatestEarthquakeCache ...
func (*CacheRepository) GetLatestEarthquakeCache() ([]models.Earthquake, bool, error) {
	var earthquakes []models.Earthquake
	modules.Redis.GetCache("latest_earthquake", &earthquakes)
	return earthquakes, len(earthquakes) != 0, nil
}

// SetLatestEarthquakeCache ...
func (*CacheRepository) SetLatestEarthquakeCache(earthquakes []models.Earthquake) error {
	return modules.Redis.SetCache("latest_earthquake", earthquakes)
}
