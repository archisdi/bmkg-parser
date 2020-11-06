package controllers

import (
	"bmkg/models"
	"bmkg/services"
)

// EarthquakeController ...
type EarthquakeController struct {
	Service services.EarthquakeServiceApi
}

// GetEarthquakes ...
func (c *EarthquakeController) GetEarthquakes() (models.Response, error) {
	return c.Service.RetrieveLatestEarthquakes()
}

// GetEarthquakesLast ...
func (c *EarthquakeController) GetEarthquakesLast() (models.Response, error) {
	return c.Service.RetrieveLastEarthquake()
}
