package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type weatherResp struct {
	List []struct {
		Main struct {
			Temp float32 `json:"temp"`
		} `json:"main"`
	} `json:"list"`
}

const (
	url = "http://api.openweathermap.org/data/2.5/forecast"
)

var apiKey = os.Getenv("API_KEY")

func getWeatherByCity(city string) (float32, error) {
	url := fmt.Sprintf("%s?q=%s&APPID=%s&units=metric", url, city, apiKey)
	fmt.Println(url)
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	res, err := client.Get(url)
	if err != nil {
		return 0, err
	}

	defer func(){
		_ = res.Body.Close()
	}()

	d := weatherResp{}
	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		return 0, err
	}
	if len(d.List) == 0 {
		return 0, errors.New("no results")
	}

	return d.List[0].Main.Temp, nil
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("bad usage of cmd")
	}

	city := os.Args[1]
	temprature, err := getWeatherByCity(city)
	if err != nil {
		log.Fatalf("failure: %v", err)
	}
	if temprature >= float32(15) {
		fmt.Println("no hoodie")
	} else {
		fmt.Println("hoodie")
	}
}
