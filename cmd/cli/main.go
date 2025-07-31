package main

import (
	"currency-converter-cli/internal/api"
	"currency-converter-cli/internal/cache"
	"currency-converter-cli/internal/config"
	"currency-converter-cli/internal/converter"
	"currency-converter-cli/internal/utils"
	"log"
)

func main() {
	config.LoadEnv()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	apiKey := config.GetEnv("API_KEY")
	baseUrl := config.GetEnv("BASE_URL")
	cacheFile := config.GetEnv("CACHE_FILE")
	url := baseUrl + "/" + apiKey
	args, err := utils.ArgParse()
	if err != nil {
		return err
	}
	if args.List {
		codes, err := api.GetSupportedCodes(url)
		if err != nil {
			return err
		}
		utils.PrintSupportedCodes(codes)
		return nil
	}
	cacheKey := cache.GetCacheKey(args)
	cacheCurrency, err := cache.GetCacheValue(cacheFile, cacheKey)
	if err != nil {
		return err
	}
	if cacheCurrency != nil {
		handleConversion(args.Amount, *cacheCurrency, args.From, args.To)
		return nil
	} else {
		con, err := api.GetPairConversion(url, args.From, args.To)
		if err != nil {
			return err
		}
		cache.SetCacheValue(cacheFile, cacheKey, con.ConversionRate)
		handleConversion(args.Amount, con.ConversionRate, args.From, args.To)
	}
	return nil
}

func handleConversion(amount float64, rate float64, from string, to string) {
	result := converter.Convert(amount, rate)
	utils.PrintConversionResult(amount, result, from, to)
}
