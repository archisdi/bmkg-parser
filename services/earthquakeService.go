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
	earthquake, _ := s.repo.GetLastEarthquake()
	return earthquake.Gempa.ToEarthquake(), nil
}

// RetrieveLatestEarthquakes ...
func (s *EarthquakeService) RetrieveLatestEarthquakes() ([]models.Earthquake, error) {
	earthquakes, _ := s.repo.GetLatestEarthquake()
	return earthquakes.ToEarthquakeList(), nil
}
