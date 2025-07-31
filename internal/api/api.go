package api

import (
	"currency-converter-cli/pkg/models"
)

type CurrencyAPI interface {
	GetPairConversion(from, to string) (*models.Conversion, error)
	GetSupportedCodes() ([][]string, error)
}
