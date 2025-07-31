package cli

import (
	"currency-converter-cli/internal/api"
	"currency-converter-cli/internal/cache"
	"currency-converter-cli/internal/converter"
	"currency-converter-cli/internal/utils"
	"currency-converter-cli/pkg/models"
)

func Run(apiClient api.CurrencyAPI, args *models.Args, cacheFile string) error {
	if args.List {
		codes, err := apiClient.GetSupportedCodes()
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
	if cacheCurrency != nil && !args.Refresh {
		handleConversion(args.Amount, *cacheCurrency, args.From, args.To)
		return nil
	} else {
		con, err := apiClient.GetPairConversion(args.From, args.To)
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
