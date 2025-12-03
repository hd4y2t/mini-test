package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"io"
)

type ForecastResponse struct {
	List []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	} `json:"list"`
}

type GeoResponse struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

func main() {
	// apiKey := "76687ba93a49cc8decb2a7d8718ed39d"
	apiKey := "Token"
	city := "Jakarta"

	lat, lon, err := getLatLon(city, apiKey)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&units=metric&appid=%s", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data ForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	tempsByDay := make(map[string]float64)
	dayObjects := make(map[string]time.Time)

	for _, item := range data.List {
		t := time.Unix(item.Dt, 0).UTC().Add(7 * time.Hour)
		day := t.Format("2006-01-02")

		if _, exists := tempsByDay[day]; !exists {
			tempsByDay[day] = item.Main.Temp
			dayObjects[day] = t
		}
	}

	fmt.Println("Weather Forecast:")
	count := 0
	for day, temp := range tempsByDay {
		if count == 5 {
			break
		}

		t := dayObjects[day]
		dayLabel := t.Format("Mon, 02 Jan 2006")

		fmt.Printf("%s: %.2f°C\n", dayLabel, temp)
		count++
	}
}

func getLatLon(city, apiKey string) (float64, float64, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var geo []GeoResponse
	json.Unmarshal(body, &geo)

	if len(geo) == 0 {
		return 0, 0, fmt.Errorf("lokasi tidak ditemukan")
	}

	return geo[0].Lat, geo[0].Lon, nil
}
