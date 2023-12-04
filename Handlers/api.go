package Handlers

import (
	"encoding/json"
	"example.com/Models"
	"io"
	"log"
	"net/http"
)

func GetData() int {
	url := "https://api.tomorrow.io/v4/weather/realtime?location=glen%20mills&apikey=DcIy5mM8YHNwH94Ow4FsQTLZjKaoMcxo"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	var weatherResp Models.WeatherData
	body, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(body, &weatherResp)
	if err != nil {
		log.Fatal("Error decoding JSON", err)
		return -1
	}

	return weatherResp.Data.Values.PrecipitationProbability
}
