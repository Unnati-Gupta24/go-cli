package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"github.com/fatih/color"
	"time"
	"log"
	"net/http"
	"bufio"
	"strings"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter city: ")
	q, _ := reader.ReadString('\n')
	q = strings.TrimSpace(q)

	fmt.Print("Enter country: ")
	country, _ := reader.ReadString('\n')
	country = strings.TrimSpace(country)

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	if len(os.Args) >= 3 {
		country = os.Args[2]
	}

	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=ec180872243c4f57a4f153631230105&q=%s,%s&days=1&aqi=no&alerts=no", q, country))
	if err!=nil{
		panic(err)
	}
	defer res.Body.Close()
	
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}
	
	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name, location.Country, current.TempC, current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(int64(hour.TimeEpoch), 0)
		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf(
			"%s - %.0fC, %.0f, %s",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		if hour.ChanceOfRain < 40{
			fmt.Println(message)
		}else{
			color.Red(message)
		}
	} 
}