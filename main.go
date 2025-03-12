package main

import (
	"flag"
	"github.com/luckswish/weather-cli/config"
	format "github.com/luckswish/weather-cli/internal/utils"
	"github.com/luckswish/weather-cli/internal/weather"
	"log"
)

func main() {
	city := flag.String("city", "Moscow", "City name for weather info")
	flag.Parse()
	config.LoadEnv()
	apiKey := config.GetApiKey()
	w, err := weather.GetWeather(*city, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	format.FormatWeather(w)
}
