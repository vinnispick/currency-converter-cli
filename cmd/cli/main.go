package main

import (
	"currency-converter-cli/internal/api"
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
	c := api.GetPairConversion(url, args.From, args.To)

	res := converter.Convert(args.Amount, c.ConversionRate)
	utils.PrintConversionResult(args.Amount, res, args.From, args.To)
}
