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
