package main

import (
	"fmt"
	"log"
)

func main() {
	userIP, err := getPublicIP()
	if err != nil {
		log.Fatalln(err)
	}
	geoInfo, err := getLatLon(userIP)
	if err != nil {
		log.Fatalln(err)
	}
	currentTemp, err := getWeather(geoInfo.lat, geoInfo.lon)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Your Public address : ", userIP, "\n",
		"You're in: ", geoInfo.city, "\n",
		"Current temprature : ", currentTemp)
}
