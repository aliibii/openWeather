package main

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func getLatLon(ip string) (float64, float64, string, error) {
	req, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return 0, 0, "none", err
	}
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return 0, 0, "none", err
	}
	lat := gjson.Get(string(res), "lat").Float()
	lon := gjson.Get(string(res), "lon").Float()
	city := gjson.Get(string(res), "city").String()
	return lat, lon, city, nil
}
