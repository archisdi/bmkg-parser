package repositories

import (
	"bmkg/models"
	"bmkg/utils"
	"encoding/xml"
)

// WeatherRepositoryApi ...
type WeatherRepositoryApi interface {
	GetLocationWeatherForecast(resource string) (models.BaseWeather, error)
}

// WeatherRepository ...
type WeatherRepository struct{}

// GetLocationWeatherForecast ...
func (*WeatherRepository) GetLocationWeatherForecast(resource string) (models.BaseWeather, error) {
	xmlBytes, err := utils.GetXMLFromURL("https://data.bmkg.go.id/datamkg/MEWS/DigitalForecast/DigitalForecast-" + resource + ".xml")

	var weather models.BaseWeather
	if err = xml.Unmarshal(xmlBytes, &weather); err != nil {
		return weather, err
	}

	return weather, nil
}
