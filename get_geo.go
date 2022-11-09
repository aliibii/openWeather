package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

type geo struct {
	lat  float64
	lon  float64
	city string
}

var geoInfo geo

func getLatLon(ip string) (geo, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := client.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return geoInfo, err
	}
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return geoInfo, err
	}
	geoInfo.lat = gjson.Get(string(res), "lat").Float()
	geoInfo.lon = gjson.Get(string(res), "lon").Float()
	geoInfo.city = gjson.Get(string(res), "city").String()
	return geoInfo, nil
}
