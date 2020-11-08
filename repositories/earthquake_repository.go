package repositories

import (
	"bmkg/models"
	"bmkg/utils"
	"encoding/xml"
)

// NewEarthquakeRepository ...
func NewEarthquakeRepository() EarthquakeRepositoryAPI {
	return &EarthquakeRepository{}
}

// EarthquakeRepositoryAPI ...
type EarthquakeRepositoryAPI interface {
	GetLastEarthquake() (models.LastEartquake, error)
	GetLatestEarthquake() (models.LatestEarthquake, error)
}

// EarthquakeRepository ...
type EarthquakeRepository struct{}

// GetLastEarthquake ...
func (r *EarthquakeRepository) GetLastEarthquake() (models.LastEartquake, error) {
	xmlBytes, err := utils.GetXMLFromURL("https://data.bmkg.go.id/autogempa.xml")

	var earthquake models.LastEartquake
	if err = xml.Unmarshal(xmlBytes, &earthquake); err != nil {
		return earthquake, err
	}

	return earthquake, err
}

// GetLatestEarthquake ...
func (r *EarthquakeRepository) GetLatestEarthquake() (models.LatestEarthquake, error) {
	xmlBytes, err := utils.GetXMLFromURL("https://data.bmkg.go.id/gempaterkini.xml")

	var earthquakes models.LatestEarthquake
	if err = xml.Unmarshal(xmlBytes, &earthquakes); err != nil {
		return earthquakes, err
	}

	return earthquakes, err
}
