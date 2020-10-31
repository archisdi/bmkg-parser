package controllers

import (
	"bmkg/models"
	"bmkg/repositories"

	"github.com/kataras/iris/v12"
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

// GetEarthquakesLastGif ...
func (EarthquakeController) GetEarthquakesLastGif(ctx iris.Context) {
	ctx.Redirect("https://data.bmkg.go.id/eqmap.gif", iris.StatusPermanentRedirect)
}
