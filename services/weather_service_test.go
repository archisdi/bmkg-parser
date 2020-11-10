package services

import (
	"bmkg/mocks"
	"bmkg/models"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRetrieveRegionalWeatherForecastFromSourceSuccess(t *testing.T) {
	weatherRepoMock := mocks.WeatherRepositoryAPI{}
	cacheRepoMock := mocks.CacheRepositoryAPI{}

	regionMock := "SulawesiSelatan"
	coordinateMock := "123.3, 321.1"

	baseWeatherMock := models.BaseWeather{
		XMLName:          xml.Name{},
		Text:             "-",
		Source:           "bmkg",
		Productioncenter: "jakarta",
		Forecast: models.Forecast{
			Text:   "-",
			Domain: "-",
			Issue: models.Issue{
				Text:      "-",
				Timestamp: "-",
				Year:      "-",
				Month:     "-",
				Day:       "-",
				Hour:      "-",
				Minute:    "-",
				Second:    "-",
			} ,
			Area:   []models.Area{
				{
					Text:        "-",
					ID:          "-",
					Latitude:    "123.3",
					Longitude:   "321.1",
					Coordinate:  "123.3, 321.1",
					Type:        "-",
					Region:      "-",
					Level:       "-",
					Description: "-",
					Domain:      "-",
					Tags:        "-",
					Name: []models.AreaName{
						{Text: "wow", Lang: "ID"},
					},
					Parameter: []models.Parameter{
						{
							ID:          "weather",
							Description: "-",
							Type:        "-",
							Timerange: []models.Timerange{
								{
									Type:     "",
									H:        "",
									Datetime: "",
									Day:      "",
									Value:    []models.TimerangeVal{ {Text: "1", Unit: "icon" }, },
												},
											},
										},
									},
								},
			},
		},
	}

	cacheRepoMock.On("GetRegionWeatherCache", regionMock).Return(models.BaseWeather{}, false, nil).Once()
	cacheRepoMock.On("SetRegionWeatherCache", regionMock, baseWeatherMock).Return(nil)

	weatherRepoMock.On("GetWeatherForecast", regionMock).Return(baseWeatherMock, nil).Once()

	service := NewWeatherService(&weatherRepoMock, &cacheRepoMock)
	result, err := service.RetrieveRegionalWeatherForecast(regionMock, coordinateMock)

	assert.Equal(t, err, nil)
	assert.Equal(t,  result.Data[0].Weather.Value, "1")
}
