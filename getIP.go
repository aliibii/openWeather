package main

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func getPublicIP() (string, error) {
	req, err := http.Get("https://api.bigdatacloud.net/data/client-ip")
	if err != nil {
		return "", err
	}
	res, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	clientPublicIP := gjson.Get(string(res), "ipString").String()
	return clientPublicIP, nil
}
