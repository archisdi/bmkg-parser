package models

import (
	"bmkg/utils"
	"encoding/xml"
	"time"
)

// GeoLocation ...
type GeoLocation struct {
	Province struct {
		Name       string `json:"name"`
		Coordinate string `json:"coordinate"`
	} `json:"province"`
	Region struct {
		Name       string `json:"name"`
		Coordinate string `json:"coordinate"`
	} `json:"region"`
}

// Weather ...
type Weather struct {
	Value       string    `json:"value"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

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

// Area ...
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
	Type     string         `xml:"type,attr" json:"type"`
	H        string         `xml:"h,attr" json:"H"`
	Datetime string         `xml:"datetime,attr" json:"date_time"`
	Day      string         `xml:"day,attr" json:"day"`
	Value    []TimerangeVal `xml:"value" json: "value"`
}

// TimerangeVal ...
type TimerangeVal struct {
	Text string `xml:",chardata" json:"text"`
	Unit string `xml:"unit,attr" json:"unit"`
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
func (a *Area) GetWeather() []Weather {
	parameter := Parameter{}
	for _, params := range a.Parameter {
		if params.ID == "weather" {
			parameter = params
		}
	}

	var weathers []Weather
	for _, timerange := range parameter.Timerange {
		weathers = append(weathers, timerange.ToWeather())
	}

	return weathers
}

// GetValue ...
func (t *Timerange) GetValue() string {
	return t.Value[0].Text
}

// ToWeather ...
func (t *Timerange) ToWeather() Weather {
	val := t.GetValue()
	time, _ := time.Parse("200601021504", t.Datetime)
	return Weather{
		Value:       val,
		Description: utils.Constant.WeatherCode[val],
		Timestamp:   time,
	}
}
