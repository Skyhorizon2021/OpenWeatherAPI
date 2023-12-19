package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// apiKey has been ommitted to maintain security
// Please obtain a key from OpenWeather API
const apiKey = "secret"

func fetchWeather(city string) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s&units=imperial", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
	}

	return data
}

func main() {
	timeStart := time.Now()
	cities := []string{"Tokyo,JP", "London,UK", "Sacramento,US", "Beijing,CN", "Rio de Janeiro,BR"}

	for _, city := range cities {
		data := fetchWeather(city)
		fmt.Println(city, ":", data)
	}

	fmt.Println("This operation took:", time.Since(timeStart))
}
