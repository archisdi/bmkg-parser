package models

type Request struct {}

type WeatherRequest struct {
	Coordinate string `json:"coordinate" validate:"required"`
	Region 	   string `json:"region"`
}