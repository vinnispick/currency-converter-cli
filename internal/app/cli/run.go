package cli

import (
	"currency-converter-cli/internal/api"
	"currency-converter-cli/internal/cache"
	"currency-converter-cli/internal/converter"
	"currency-converter-cli/internal/utils"
	"currency-converter-cli/pkg/models"
)

func Run(apiClient api.CurrencyAPI, args *models.Args, cache cache.Cache) error {
	if args.List {
		codes, err := apiClient.GetSupportedCodes()
		if err != nil {
			return err
		}
		utils.PrintSupportedCodes(codes)
		return nil
	}
	cacheKey := utils.GetCacheKey(args)
	cacheCurrency, err := cache.Get(cacheKey)
	if err != nil {
		return err
	}
	if cacheCurrency != nil && !args.Refresh {
		handleConversion(args.Amount, *cacheCurrency, args.From, args.To)
		return nil
	}
	con, err := apiClient.GetPairConversion(args.From, args.To)
	if err != nil {
		return err
	}
	err = cache.Set(cacheKey, con.ConversionRate)
	if err != nil {
		return err
	}
	handleConversion(args.Amount, con.ConversionRate, args.From, args.To)
	return nil
}

func handleConversion(amount float64, rate float64, from string, to string) {
	result := converter.Convert(amount, rate)
	utils.PrintConversionResult(amount, result, from, to)
}
