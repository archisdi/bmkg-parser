package controllers

import (
	"bmkg/mocks"
	"bmkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetLastEarthquakeSuccess(t *testing.T) {
	service := &mocks.EarthquakeServiceAPI{}

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

func TestGetLatestEarthquakeSuccess(t *testing.T) {
	service := &mocks.EarthquakeServiceAPI{}
	earthquakesMock := []models.Earthquake{
		{"123,321", 10, 69, time.Time{} },
		{"321,123", 69, 10, time.Time{} },
	}
	service.On("RetrieveLatestEarthquakes").Return(earthquakesMock, nil).Once()

	controller := &EarthquakeController{Service: service}
	earthquakes, _ := controller.GetEarthquakes()

	assert.Equal(t, len(earthquakes.([]models.Earthquake)), len(earthquakesMock))
}