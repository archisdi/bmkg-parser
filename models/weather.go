package models

import "encoding/xml"

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
	Area   []area `xml:"area"`
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

type area struct {
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
	Parameter []parameter `xml:"parameter"`
}

type parameter struct {
	Text        string      `xml:",chardata"`
	ID          string      `xml:"id,attr"`
	Description string      `xml:"description,attr"`
	Type        string      `xml:"type,attr"`
	Timerange   []timerange `xml:"timerange"`
}

type timerange struct {
	Text     string `xml:",chardata"`
	Type     string `xml:"type,attr"`
	H        string `xml:"h,attr"`
	Datetime string `xml:"datetime,attr"`
	Day      string `xml:"day,attr"`
	Value    []struct {
		Text string `xml:",chardata"`
		Unit string `xml:"unit,attr"`
	} `xml:"value"`
}
