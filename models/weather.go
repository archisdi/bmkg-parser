package models

import (
	"bmkg/utils"
	"encoding/xml"
)

// BaseWeather ...
type BaseWeather struct {
	XMLName          xml.Name `xml:"data"`
	Text             string   `xml:",chardata"`
	Source           string   `xml:"source,attr"`
	Productioncenter string   `xml:"productioncenter,attr"`
	Forecast         forecast `xml:"forecast"`
}

type forecast struct {
	Text   string `xml:",chardata"`
	Domain string `xml:"domain,attr"`
	Issue  issue  `xml:"issue"`
	Area   []Area `xml:"area"`
}

type issue struct {
	Text      string `xml:",chardata"`
	Timestamp string `xml:"timestamp"`
	Year      string `xml:"year"`
	Month     string `xml:"month"`
	Day       string `xml:"day"`
	Hour      string `xml:"hour"`
	Minute    string `xml:"minute"`
	Second    string `xml:"second"`
}

type Area struct {
	Text        string `xml:",chardata"`
	ID          string `xml:"id,attr"`
	Latitude    string `xml:"latitude,attr"`
	Longitude   string `xml:"longitude,attr"`
	Coordinate  string `xml:"coordinate,attr"`
	Type        string `xml:"type,attr"`
	Region      string `xml:"region,attr"`
	Level       string `xml:"level,attr"`
	Description string `xml:"description,attr"`
	Domain      string `xml:"domain,attr"`
	Tags        string `xml:"tags,attr"`
	Name        []struct {
		Text string `xml:",chardata"`
		Lang string `xml:"lang,attr"`
	} `xml:"name"`
	Parameter []Parameter `xml:"parameter"`
}

// Parameter ...
type Parameter struct {
	ID          string      `xml:"id,attr" json:"id"`
	Description string      `xml:"description,attr" json:"description"`
	Type        string      `xml:"type,attr" json:"type"`
	Timerange   []Timerange `xml:"timerange" json:"time_range"`
}

// Timerange ...
type Timerange struct {
	Type     string `xml:"type,attr" json:"type"`
	H        string `xml:"h,attr" json:"H"`
	Datetime string `xml:"datetime,attr" json:"date_time"`
	Day      string `xml:"day,attr" json:"day"`
	Value    []struct {
		Text string `xml:",chardata" json:"text"`
		Unit string `xml:"unit,attr" json:"unit"`
	} `xml:"value" json: "value"`
}

// GetCoordinates ...
func (a *Area) GetCoordinates() string {
	return a.Latitude + "," + a.Longitude
}

// GetName ...
func (a *Area) GetName() string {
	return a.Name[0].Text
}

// GetDomain ...
func (a *Area) GetDomain() string {
	return utils.SpaceFieldsJoin(a.Domain)
}

// GetWeather ...
func (a *Area) GetWeather() Parameter {
	for _, params := range a.Parameter {
		if params.ID == "weather" {
			return params
		}
	}
	return Parameter{}
}
