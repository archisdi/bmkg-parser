package controllers

import (
	"bmkg/models"
	"bmkg/repositories"
)

// EarthquakeController ...
type EarthquakeController struct {
	repo repositories.EarthquakeRepository
}

// GetEarthquakes ...
func (c *EarthquakeController) GetEarthquakes() (models.Response, error) {
	earthquake, err := c.repo.GetLatestEarthquake()
	return earthquake.ToEarthquakeList(), err
}

// GetEarthquakesLast ...
func (c *EarthquakeController) GetEarthquakesLast() (models.Response, error) {
	earthquake, err := c.repo.GetLastEarthquake()
	return earthquake.Gempa.ToEarthquake(), err
}
