package mocks

import "currency-converter-cli/pkg/models"

type MockCurrencyAPI struct {
}

func (m *MockCurrencyAPI) GetPairConversion(url, from, to string) (*models.Conversion, error) {
	return &models.Conversion{ConversionRate: 1.0}, nil // Mock conversion rate
}
func (m *MockCurrencyAPI) GetSupportedCodes() ([][]string, error) {
	return [][]string{{"USD", "United States Dollar"}, {"EUR", "Euro"}}, nil // Mock supported codes
}
