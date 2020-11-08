package services

import (
	"bmkg/models"
	"bmkg/repositories"
)

func NewEarthquakeService(
	repo repositories.EarthquakeRepositoryAPI, cache repositories.CacheRepositoryAPI) EarthquakeServiceAPI {
	return &EarthquakeService{
		Repo: repo,
		Cache: cache,
	}
}

// EarthquakeServiceAPI
type EarthquakeServiceAPI interface {
	RetrieveLastEarthquake() (models.Earthquake, error)
	RetrieveLatestEarthquakes() ([]models.Earthquake, error)
}

// EarthquakeService ...
type EarthquakeService struct {
	Repo  repositories.EarthquakeRepositoryAPI
	Cache repositories.CacheRepositoryAPI
}

// RetrieveLastEarthquake ...
func (s *EarthquakeService) RetrieveLastEarthquake() (models.Earthquake, error) {
	var earthquake models.Earthquake

	if cachedEarthquake, ok, _ := s.Cache.GetLastEarthquakeCache(); ok {
		return cachedEarthquake, nil
	}

	// if Cache not found, retrieve from source and Cache
	lastEarthquake, _ := s.Repo.GetLastEarthquake()
	earthquake = lastEarthquake.Gempa.ToEarthquake()
	if err := s.Cache.SetLastEarthquakeCache(earthquake); err != nil {
		return earthquake ,err
	}

	return earthquake, nil
}

// RetrieveLatestEarthquakes ...
func (s *EarthquakeService) RetrieveLatestEarthquakes() ([]models.Earthquake, error) {
	var earthquakes []models.Earthquake

	if cachedEarthquakes, ok, _ := s.Cache.GetLatestEarthquakeCache(); ok {
		return cachedEarthquakes, nil
	}

	// if Cache not found, retrieve from source and Cache
	latestEarthquake, _ := s.Repo.GetLatestEarthquake()
	earthquakes = latestEarthquake.ToEarthquakeList()
	if err := s.Cache.SetLatestEarthquakeCache(earthquakes); err != nil {
		return earthquakes, err
	}

	return earthquakes, nil
}
