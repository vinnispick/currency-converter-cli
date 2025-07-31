package main

import (
	"currency-converter-cli/internal/api"
	"currency-converter-cli/internal/cache"
	"currency-converter-cli/internal/config"
	"currency-converter-cli/internal/converter"
	"currency-converter-cli/internal/utils"
	"log"
)

func init() {
	config.LoadEnv()
}

func main() {
	apiKey := config.GetEnv("API_KEY")
	baseUrl := config.GetEnv("BASE_URL")
	cacheFile := config.GetEnv("CACHE_FILE")
	url := baseUrl + "/" + apiKey
	args, err := utils.ArgParse()
	if err != nil {
		log.Fatalf("args with no arguments!")
	}
	if args.List {
		codes := api.GetSupportedCodes(url)
		utils.PrintSupportedCodes(codes)
		return
	}
	cv := cache.GetCacheKey(args)
	cacheCurrency := cache.GetCacheValue(cacheFile, cv)
	if cacheCurrency != nil {
		res := converter.Convert(args.Amount, *cacheCurrency)
		utils.PrintConversionResult(args.Amount, res, args.From, args.To)
		return
	} else {
		c := api.GetPairConversion(url, args.From, args.To)
		cache.SetCacheValue(cacheFile, cv, c.ConversionRate)
		res := converter.Convert(args.Amount, c.ConversionRate)
		utils.PrintConversionResult(args.Amount, res, args.From, args.To)
	}

}
