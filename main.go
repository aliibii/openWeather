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
	lat, lon, city, err := getLatLon(userIP)
	if err != nil {
		log.Fatalln(err)
	}
	currentTemp, err := getWeather(lat, lon)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Your Public address : ", userIP, "\n",
		"You're in: ", city, "\n",
		"Current temprature : ", currentTemp)
}
