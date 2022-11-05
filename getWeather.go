package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func getWeather(lat float64, lon float64) string {
	req := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&hourly=temperature_2m&current_weather=true", lat, lon)
	res, err := http.Get(req)
	if err != nil {
		log.Fatalln(err)
	}
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	currentTemp := gjson.Get(string(responseBody), "current_weather.temperature").String()

	return currentTemp
}
