package controllers

import (
	"bmkg/mocks"
	"bmkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetWeatherSuccess(t *testing.T) {
	service := &mocks.WeatherServiceAPI{}

	weatherMock := models.WeatherOutput{
		Location: models.GeoLocation{
			Name:       "Makassar",
			Coordinate: "123,321",
		},
		Data: []models.Prediction{
			{
				Weather: models.PredictionParam{
					Unit:        "icon",
					Value:       "69",
					Description: "its wednesday my dudes",
				},
				Temperature: models.PredictionParam{},
				Humidity: models.PredictionParam{},
				WindDirection: models.PredictionParam{},
				WindSpeed: models.PredictionParam{},
				Timestamp: time.Now(),
			},
		},
	}

	requestMock := models.WeatherRequest{
		Coordinate: "321,123",
		Region:     "Indonesia",
	}

	service.On("RetrieveRegionalWeatherForecast", requestMock.Region, requestMock.Coordinate).Return(weatherMock, nil).Once()

	controller := &WeatherController{Service: service}
	weather, _ := controller.GetWeather(requestMock)

	assert.Equal(t, weather.(models.WeatherOutput), weatherMock)
}
