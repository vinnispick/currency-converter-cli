package api

import "currency-converter-cli/pkg/models"

type LiveCurrencyAPI struct {
	BaseURL string
}

func (api *LiveCurrencyAPI) GetSupportedCodes() ([][]string, error) {
	return GetSupportedCodes(api.BaseURL)
}
func (api *LiveCurrencyAPI) GetPairConversion(from, to string) (*models.Conversion, error) {
	return GetPairConversion(api.BaseURL, from, to)
}
