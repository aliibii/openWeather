package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func getPublicIP() string {
	req, err := http.Get("https://api.bigdatacloud.net/data/client-ip")
	if err != nil {
		log.Fatalln(err)
	}
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	clientPublicIP := gjson.Get(string(res), "ipString").String()
	return clientPublicIP
}

func getLatLon(ip string) (string, string) {
	req, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	lat := gjson.Get(string(res), "lat").String()
	lon := gjson.Get(string(res), "lon").String()
	return lat, lon
}

// func main() {
// 	lat, lon := getLatLon(getPublicIP())
// 	fmt.Println(lat, lon)
// }
