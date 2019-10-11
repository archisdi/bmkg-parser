package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Eartquake struct {
	magnitude float32
	depth     float32
	latitude  float32
	longitude float32
	date      time.Time
}

type Infogempa struct {
	XMLName xml.Name `xml:"Infogempa"`
	Text    string   `xml:",chardata"`
	Gempa   struct {
		Text    string `xml:",chardata"`
		Tanggal string `xml:"Tanggal"`
		Jam     string `xml:"Jam"`
		Point   struct {
			Text        string `xml:",chardata"`
			Coordinates string `xml:"coordinates"`
		} `xml:"point"`
		Lintang   string `xml:"Lintang"`
		Bujur     string `xml:"Bujur"`
		Magnitude string `xml:"Magnitude"`
		Kedalaman string `xml:"Kedalaman"`
		Symbol    string `xml:"_symbol"`
		Wilayah1  string `xml:"Wilayah1"`
		Wilayah2  string `xml:"Wilayah2"`
		Wilayah3  string `xml:"Wilayah3"`
		Wilayah4  string `xml:"Wilayah4"`
		Wilayah5  string `xml:"Wilayah5"`
		Potensi   string `xml:"Potensi"`
	} `xml:"gempa"`
}

func (e Eartquake) String() string {
	return ("coordinates: " + fmt.Sprintf("%f", e.latitude))
}

func (e Infogempa) String() string {
	return ("coordinates: " + e.Gempa.Point.Coordinates + "; magnitude: " + e.Gempa.Magnitude + "; time: " + e.Gempa.Tanggal + " " + e.Gempa.Jam)
}
