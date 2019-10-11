package main

import (
	"encoding/xml"
)

type Earthquake struct {
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
	Wilayah   string `xml:"Wilayah"`
}

type LatestEartquake struct {
	XMLName xml.Name   `xml:"Infogempa"`
	Text    string     `xml:",chardata"`
	Gempa   Earthquake `xml:"gempa"`
}

type EarthquakeList struct {
	XMLName xml.Name     `xml:"Infogempa"`
	Text    string       `xml:",chardata"`
	Gempa   []Earthquake `xml:"gempa"`
}

func (e Earthquake) String() string {
	return ("coordinates: " + e.Point.Coordinates + "; magnitude: " + e.Kedalaman + "; time: " + e.Tanggal + " " + e.Jam)
}
