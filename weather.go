package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

type conf struct {
	APIKEY string
}

func getKey() string {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	conf := conf{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error", err)
	}
	return conf.APIKEY
}

func query(city string) (weatherData, error) {
	key := getKey()
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + key + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}
