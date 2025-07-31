package main

import (
	"currency-converter-cli/internal/api"
	"currency-converter-cli/internal/app/cli"
	"currency-converter-cli/internal/config"
	"currency-converter-cli/internal/utils"
	"log"
)

func main() {
	config.LoadEnv()

	apiKey := config.GetEnv("API_KEY")
	baseURL := config.GetEnv("BASE_URL")
	cacheFile := config.GetEnv("CACHE_FILE")

	args, err := utils.ArgParse()

	if err != nil {
		log.Fatal(err)
	}

	liveAPI := &api.LiveCurrencyAPI{BaseURL: baseURL + "/" + apiKey}

	if err := cli.Run(liveAPI, args, cacheFile); err != nil {
		log.Fatal(err)
	}
}
