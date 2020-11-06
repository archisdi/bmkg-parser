package controllers

import (
	"bmkg/models"
	"bmkg/services"

	"github.com/kataras/iris/v12"
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

// GetEarthquakesLastGif ...
func (EarthquakeController) GetEarthquakesLastGif(ctx iris.Context) {
	ctx.Redirect("https://data.bmkg.go.id/eqmap.gif", iris.StatusPermanentRedirect)
}
