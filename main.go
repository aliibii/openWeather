package main

import (
	"fmt"
)

func main() {
	lat, lon, city := getLatLon(getPublicIP())
	fmt.Println("Your Public address : ", getPublicIP(), "\n",
		"You're in: ", city, "\n",
		"Current temprature : ", getWeather(lat, lon))
}
