package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func main() {
	apiResponse, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=52.3738&longitude=4.8910&hourly=temperature_2m&current_weather=true")
	if err != nil {
		log.Fatalln(err)
	}
	responseBody, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}
	currentTemp := gjson.Get(string(responseBody), "current_weather.temperature")
	lat, lon := getLatLon(getPublicIP())
	fmt.Println(lat, lon)
	fmt.Println("Current temprature is: ", currentTemp)

}
