package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func getLatLon(ip string) (float64, float64, string) {
	req, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		log.Fatalln(err)
	}
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	lat := gjson.Get(string(res), "lat").Float()
	lon := gjson.Get(string(res), "lon").Float()
	city := gjson.Get(string(res), "city").String()
	return lat, lon, city
}
