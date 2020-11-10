package services

import (
	"bmkg/mocks"
	"bmkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetrieveLastEarthquakeFromSourceSuccess(t *testing.T) {
	earthquakeRepoMock := mocks.EarthquakeRepositoryAPI{}
	cacheRepoMock := mocks.CacheRepositoryAPI{}

	earthquakeMock := models.LastEartquake{
		Gempa: models.LastEarthquakeDetail{
			BaseXMLEarthquake: models.BaseXMLEarthquake{
				Tanggal: "02-Jan-06",
				Jam:     "15:04:05 WIB",
				Point: struct {
					Coordinates string `xml:"coordinates" json:"coordinates"`
				}{
					Coordinates: "321, 123",
				},
				Lintang:   "321",
				Bujur:     "123",
				Magnitude: "5",
				Kedalaman: "50",
				Symbol:    "",
			},
			Wilayah1: "A",
			Wilayah2: "B",
			Wilayah3: "C",
			Wilayah4: "D",
			Potensi:  "tidak",
		},
	}

	cacheRepoMock.On("GetLastEarthquakeCache").Return(models.Earthquake{}, false, nil)
	cacheRepoMock.On("SetLastEarthquakeCache", earthquakeMock.Gempa.ToEarthquake()).Return(nil)

	earthquakeRepoMock.On("GetLastEarthquake").Return(earthquakeMock, nil)

	service := NewEarthquakeService(&earthquakeRepoMock, &cacheRepoMock)
	result, err := service.RetrieveLastEarthquake()

	assert.Equal(t, err, nil)
	assert.Equal(t, result, earthquakeMock.Gempa.ToEarthquake())
}

func TestRetrieveLatestEarthquakesFromSourceSuccess(t *testing.T) {
	earthquakeRepoMock := mocks.EarthquakeRepositoryAPI{}
	cacheRepoMock := mocks.CacheRepositoryAPI{}

	earthquakesMock := models.LatestEarthquake{
		Gempa: []models.LatestEartquakeDetail{
			models.LatestEartquakeDetail{
				BaseXMLEarthquake: models.BaseXMLEarthquake{
					Tanggal: "02-Jan-06",
					Jam:     "15:04:05 WIB",
					Point: struct {
						Coordinates string `xml:"coordinates" json:"coordinates"`
					}{
						Coordinates: "321, 123",
					},
					Lintang:   "321",
					Bujur:     "123",
					Magnitude: "5",
					Kedalaman: "50",
					Symbol:    "",
				},
				Wilayah: "A",
			},
		},
	}

	cacheRepoMock.On("GetLatestEarthquakeCache").Return([]models.Earthquake{}, false, nil)
	cacheRepoMock.On("SetLatestEarthquakeCache", earthquakesMock.ToEarthquakeList()).Return(nil)

	earthquakeRepoMock.On("GetLatestEarthquake").Return(earthquakesMock, nil)

	service := NewEarthquakeService(&earthquakeRepoMock, &cacheRepoMock)
	result, err := service.RetrieveLatestEarthquakes()

	assert.Equal(t, err, nil)
	assert.Equal(t, result, earthquakesMock.ToEarthquakeList())
}