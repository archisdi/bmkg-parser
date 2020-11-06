package controllers

import (
	"bmkg/mocks"
	"bmkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetLastEarthquakeSuccess(t *testing.T) {
	service := &mocks.EarthquakeServiceApi{}

	earthquakeMock := models.Earthquake{
		Coordinates: "123,321",
		Magnitude:   10,
		Depth:       69,
		Timestamp:   time.Time{},
	}
	service.On("RetrieveLastEarthquake").Return(earthquakeMock, nil).Once()

	controller := &EarthquakeController{Service: service}
	earthquake, _ := controller.GetEarthquakesLast()

	assert.Equal(t, earthquake.(models.Earthquake).Depth, earthquakeMock.Depth)
}