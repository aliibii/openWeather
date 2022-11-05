package main

import (
	"fmt"
)

func main() {
	lat, lon := getLatLon(getPublicIP())
	fmt.Println("Your Public address : ", getPublicIP(), "\n", "Current temprature : ", getWeather(lat, lon))
}
