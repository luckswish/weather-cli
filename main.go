package main

import (
	"fmt"
	"github.com/luckswish/weather-cli/config"
)

func main() {
	config.LoadEnv()
	apiKey := config.GetApiKey()
	fmt.Println(apiKey)
}
