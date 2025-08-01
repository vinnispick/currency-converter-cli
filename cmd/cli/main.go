package main

import (
	"currency-converter-cli/internal/api"
	"currency-converter-cli/internal/app/cli"
	"currency-converter-cli/internal/cache"
	"currency-converter-cli/internal/config"
	"currency-converter-cli/internal/utils"
	"log"
)

func main() {
	config.LoadEnv()

	apiKey := config.GetEnv("API_KEY")
	baseURL := config.GetEnv("BASE_URL")
	cachePath := config.GetEnv("CACHE_FILE")

	args, err := utils.ArgParse()

	if err != nil {
		log.Fatal(err)
	}

	liveAPI := &api.LiveCurrencyAPI{BaseURL: baseURL + "/" + apiKey}
	cacheFile := cache.NewFileCache(cachePath)

	if err := cli.Run(liveAPI, args, cacheFile); err != nil {
		log.Fatal(err)
	}
}
