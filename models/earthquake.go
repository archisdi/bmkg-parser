package models

import (
	"strconv"
	"strings"
	"time"
)

// Earthquake ...
type Earthquake struct {
	Coordinates string    `json:"coordinates"`
	Magnitude   float64   `json:"magnitude"`
	Depth       float64   `json:"depth"`
	Timestamp   time.Time `json:"timestamp"`
}

type baseXMLEarthquake struct {
	Tanggal string `xml:"Tanggal" json:"date"`
	Jam     string `xml:"Jam" json:"tome"`
	Point   struct {
		Coordinates string `xml:"coordinates" json:"coordinates"`
	} `xml:"point" json:"point"`
	Lintang   string `xml:"Lintang" json:"latitude"`
	Bujur     string `xml:"Bujur" json:"longitude"`
	Magnitude string `xml:"Magnitude" json:"magnitude"`
	Kedalaman string `xml:"Kedalaman" json:"depth"`
	Symbol    string `xml:"_symbol" json:"symbol"`
}

// LastEartquake ...
type LastEartquake struct {
	Gempa struct {
		baseXMLEarthquake
		Wilayah1 string `xml:"Wilayah1" json:"area_1"`
		Wilayah2 string `xml:"Wilayah2" json:"area_2"`
		Wilayah3 string `xml:"Wilayah3" json:"area_3"`
		Wilayah4 string `xml:"Wilayah4" json:"area_4"`
		Potensi  string `xml:"Potensi" json:"description"`
	} `xml:"gempa" json:"gempa"`
}

// LatestEarthquake ...
type LatestEarthquake struct {
	Gempa []struct {
		baseXMLEarthquake
		Wilayah string `xml:"Wilayah" json:"area"`
	} `xml:"gempa"`
}

// ToEarthquakeList ...
func (e *LatestEarthquake) ToEarthquakeList() []Earthquake {
	var earthquakes []Earthquake

	for _, earthquake := range e.Gempa {
		earthquakes = append(earthquakes, earthquake.ToEarthquake())
	}

	return earthquakes
}

func (e baseXMLEarthquake) String() string {
	return "coordinates: " + e.Point.Coordinates + "; magnitude: " + e.Kedalaman + "; time: " + e.Tanggal + " " + e.Jam
}

func (e *baseXMLEarthquake) ToEarthquake() Earthquake {

	parsedDate, _ := time.Parse("02-Jan-06 15:04:05 MST", e.Tanggal+" "+e.Jam)

	splDepth := strings.Split(e.Kedalaman, " ")
	depth, _ := strconv.ParseFloat(splDepth[0], 2)

	splMag := strings.Split(e.Magnitude, " ")
	magnitude, _ := strconv.ParseFloat(splMag[0], 2)

	coordinates := strings.Split(e.Point.Coordinates, ",")

	return Earthquake{
		Coordinates: coordinates[1] + "," + coordinates[0],
		Depth:       depth,
		Magnitude:   magnitude,
		Timestamp:   parsedDate.UTC(),
	}
}
