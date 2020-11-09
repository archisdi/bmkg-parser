package models

import (
	"bmkg/utils"
	"encoding/xml"
	"time"
)


func generateEmptyPredictions() []Prediction {
	var predictions []Prediction
	for i := 0; i < 12; i++ {
		predictions = append(predictions, Prediction{})
	}
	return predictions
}

// WeatherOutput ...
type WeatherOutput struct {
	Location 	GeoLocation `json:"location"`
	Data		[]Prediction 	`json:"data"`
}

// GeoLocation ...
type GeoLocation struct {
	Name       string `json:"name"`
	Coordinate string `json:"coordinate"`
}

// Prediction ...
type Prediction struct {
	Weather 		PredictionParam `json:"weather"`
	Temperature 	PredictionParam `json:"temperature"`
	Humidity		PredictionParam `json:"humidity"`
	WindDirection	PredictionParam `json:"wind_direction"`
	WindSpeed		PredictionParam `json:"wind_speed"`
	Timestamp   	time.Time 		`json:"timestamp"`
}

// PredictionParam ...
type PredictionParam struct {
	Unit 		string 	  `json:"unit"`
	Value       string    `json:"value"`
	Description string    `json:"description"`
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
func (a *Area) GetWeather() []Prediction {
	predictions := generateEmptyPredictions()
	for _, params := range a.Parameter {
		switch params.ID {
			case "weather": {
				for i, timeRange := range params.Timerange {
					predictions[i].Weather = timeRange.ToWeather()
					predictions[i].Timestamp = timeRange.GetDatetime()
				}
			}
			case "t": {
				for i, timeRange := range params.Timerange {
					predictions[i].Temperature = timeRange.ToBaseData()
				}
			}
			case "hu": {
				for i, timeRange := range params.Timerange {
					predictions[i].Humidity = timeRange.ToBaseData()
				}
			}
			case "wd": {
				for i, timeRange := range params.Timerange {
					predictions[i].WindDirection = timeRange.ToWindDirection()
				}
			}
			case "ws": {
				for i, timeRange := range params.Timerange {
					predictions[i].WindSpeed = timeRange.ToWindSpeed()
				}
			}
		}
	}
	return predictions
}

// GetValue ...
func (t *Timerange) GetValue() (string, string) {
	return t.Value[0].Text, t.Value[0].Unit
}

// ToWeather ...
func (t *Timerange) ToWeather() PredictionParam {
	val, unit := t.GetValue()
	return PredictionParam{
		Unit:        unit,
		Value:       val,
		Description: utils.Constant.WeatherCode[val],
	}
}

// ToWindDirection ...
func (t *Timerange) ToWindDirection() PredictionParam {
	return PredictionParam{
		Unit:        t.Value[0].Unit,
		Value:       t.Value[0].Text,
		Description: t.Value[1].Text,
	}
}

// ToWindSpeed ...
func (t *Timerange) ToWindSpeed() PredictionParam {
	return PredictionParam{
		Unit:        t.Value[2].Unit,
		Value:       t.Value[2].Text,
		Description: "-",
	}
}

// ToTemperature ...
func (t *Timerange) ToBaseData() PredictionParam {
	val, unit := t.GetValue()
	return PredictionParam{
		Unit:        unit,
		Value:       val,
		Description: "-",
	}
}

// GetDatetime ...
func (t *Timerange) GetDatetime() time.Time {
	parsedTime, _ := time.Parse("200601021504", t.Datetime)
	return parsedTime
}
