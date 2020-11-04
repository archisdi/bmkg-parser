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

	if cachedEarthquake, ok, _ := s.cache.GetLastEarthquakeCache(); ok {
		return cachedEarthquake, nil
	}

	// if cache not found, retrieve from source and cache
	lastEarthquake, _ := s.repo.GetLastEarthquake()
	earthquake = lastEarthquake.Gempa.ToEarthquake()
	if err := s.cache.SetLastEarthquakeCache(earthquake); err != nil {
		return earthquake ,err
	}

	return earthquake, nil
}

// RetrieveLatestEarthquakes ...
func (s *EarthquakeService) RetrieveLatestEarthquakes() ([]models.Earthquake, error) {
	var earthquakes []models.Earthquake

	if cachedEarthquakes, ok, _ := s.cache.GetLatestEarthquakeCache(); ok {
		return cachedEarthquakes, nil
	}

	// if cache not found, retrieve from source and cache
	latestEarthquake, _ := s.repo.GetLatestEarthquake()
	earthquakes = latestEarthquake.ToEarthquakeList()
	if err := s.cache.SetLatestEarthquakeCache(earthquakes); err != nil {
		return earthquakes, err
	}

	return earthquakes, nil
}
