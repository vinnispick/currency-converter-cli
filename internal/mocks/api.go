package mocks

import "currency-converter-cli/pkg/models"

type MockCurrencyAPI struct {
	Codes          [][]string
	ConversionRate float64
}

func (m *MockCurrencyAPI) GetPairConversion(from, to string) (*models.Conversion, error) {
	return &models.Conversion{ConversionRate: m.ConversionRate}, nil // Mock conversion rate
}
func (m *MockCurrencyAPI) GetSupportedCodes() ([][]string, error) {
	return m.Codes, nil // Mock supported codes
}
