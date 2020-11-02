package services

import (
	"bmkg/models"
	"bmkg/repositories"
)

// EarthquakeService ...
type EarthquakeService struct {
	repo  repositories.EarthquakeRepository
	cache repositories.CacheRepository
}

// RetrieveLastEarthquake ...
func (s *EarthquakeService) RetrieveLastEarthquake() (models.Earthquake, error) {
	var earthquake models.Earthquake

	if cacheEarthquake, ok, _ := s.cache.GetLastEarthquakeCache(); ok {
		earthquake = cacheEarthquake
	} else {
		lastEarthquake, _ := s.repo.GetLastEarthquake()
		earthquake = lastEarthquake.Gempa.ToEarthquake()
		s.cache.SetLastEarthquakeCache(earthquake)
	}

	return earthquake, nil
}

// RetrieveLatestEarthquakes ...
func (s *EarthquakeService) RetrieveLatestEarthquakes() ([]models.Earthquake, error) {
	earthquakes, _ := s.repo.GetLatestEarthquake()
	return earthquakes.ToEarthquakeList(), nil
}
