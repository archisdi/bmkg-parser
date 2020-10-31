package repositories

import (
	"bmkg/models"
	"bmkg/utils"
	"encoding/xml"
	"log"
	"os"
)

// EarthquakeRepository ...
type EarthquakeRepository struct{}

// GetLastEarthquake ...
func (r *EarthquakeRepository) GetLastEarthquake() (models.LastEartquake, error) {
	xmlBytes, err := utils.GetXMLFromURL("https://data.bmkg.go.id/autogempa.xml")

	var earthquake models.LastEartquake
	if err = xml.Unmarshal(xmlBytes, &earthquake); err != nil {
		log.Printf("Failed Parse Data: %v", err)
		os.Exit(1)
	}

	return earthquake, err
}

// GetLatestEarthquake ...
func (r *EarthquakeRepository) GetLatestEarthquake() (models.LatestEarthquake, error) {
	xmlBytes, err := utils.GetXMLFromURL("https://data.bmkg.go.id/gempaterkini.xml")

	var earthquakes models.LatestEarthquake
	if err = xml.Unmarshal(xmlBytes, &earthquakes); err != nil {
		log.Printf("Failed Parse Data: %v", err)
		os.Exit(1)
	}

	return earthquakes, err
}
