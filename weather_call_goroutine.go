package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)
//apiKey has been ommitted to maintain security
//Please obtain a key from OpenWeather API
const apiKey = "secret"

func fetchWeather(city string, ch1 chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

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

	ch1 <- fmt.Sprintf(" %f", data.Main.Temp)
	return data
}

func main() {
	timeStart := time.Now()
	cities := []string{"Tokyo,JP", "London,UK", "Sacramento,US", "Beijing,CN", "Rio de Janeiro,BR"}

	ch1 := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)

		go fetchWeather(city, ch1, &wg)
	}
	go func() {
		wg.Wait()
		close(ch1)
	}()
	count := 0
	for res := range ch1 {
		fmt.Printf("%s:", cities[count])
		fmt.Println(res)
		count += 1
	}

	fmt.Println("This operation took:", time.Since(timeStart))
}
