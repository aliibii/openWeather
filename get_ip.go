package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

func getPublicIP() (string, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := client.Get("https://api.bigdatacloud.net/data/client-ip")
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
