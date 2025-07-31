package api

import (
	"currency-converter-cli/pkg/models"
)

type CurrencyAPI interface {
	GetPairConversion(url, from, to string) (*models.Conversion, error)
	GetSupportedCodes() ([][]string, error)
}
