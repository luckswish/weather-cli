package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/luckswish/weather-cli/internal/weather"
)

func FormatWeather(w *weather.WeatherResponse) {
	city := w.Name
	description := w.Weather[0].Description
	temp := w.Main.Temp
	feels := w.Main.FeelsLike
	color.Cyan(fmt.Sprint("Here is weather forecast for ", city))
	color.Blue(fmt.Sprintf("Description: %s", description))
	color.Magenta(fmt.Sprintf("Temp: %.2f\nFeels like: %.2f", temp, feels))
}
