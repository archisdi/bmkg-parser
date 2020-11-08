package utils

// Constant ...
var Constant constant

type constant struct {
	WeatherCode map[string]string
}

// InitializeConstant ...
func InitializeConstant() {
	Constant = constant{
		WeatherCode: map[string]string{
			"0":  "Clear Skies",
			"1":  "Partly Cloudy",
			"2":  "Partly Cloudy",
			"3":  "Mostly Cloudy",
			"4":  "Overcast",
			"5":  "Haze",
			"10": "Smoke",
			"45": "Fog",
			"60": "Light Rain",
			"61": "Rain",
			"63": "Heavy Rain",
			"80": "Isolated Shower",
			"95": "Severe Thunderstorm",
			"97": "Severe Thunderstorm",
		},
	}
}
