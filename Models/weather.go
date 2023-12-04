package Models

import (
	"time"
)

type WeatherData struct {
	Data     Data     `json:"data"`
	Location Location `json:"location"`
}

type Data struct {
	Time   time.Time `json:"time"`
	Values Values    `json:"values"`
}

type Values struct {
	CloudBase                float64 `json:"cloudBase"`
	CloudCeiling             float64 `json:"cloudCeiling"`
	CloudCover               int     `json:"cloudCover"`
	DewPoint                 float64 `json:"dewPoint"`
	FreezingRainIntensity    int     `json:"freezingRainIntensity"`
	Humidity                 int     `json:"humidity"`
	PrecipitationProbability int     `json:"precipitationProbability"`
	PressureSurfaceLevel     float64 `json:"pressureSurfaceLevel"`
	RainIntensity            int     `json:"rainIntensity"`
	SleetIntensity           int     `json:"sleetIntensity"`
	SnowIntensity            int     `json:"snowIntensity"`
	Temperature              float64 `json:"temperature"`
	TemperatureApparent      float64 `json:"temperatureApparent"`
	UVHealthConcern          int     `json:"uvHealthConcern"`
	UVIndex                  int     `json:"uvIndex"`
	Visibility               int     `json:"visibility"`
	WeatherCode              int     `json:"weatherCode"`
	WindDirection            float64 `json:"windDirection"`
	WindGust                 float64 `json:"windGust"`
	WindSpeed                float64 `json:"windSpeed"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
	Type string  `json:"type"`
}
