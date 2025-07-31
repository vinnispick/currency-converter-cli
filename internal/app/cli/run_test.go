package cli

import (
	"currency-converter-cli/internal/mocks"
	"currency-converter-cli/pkg/models"
	"os"
	"testing"
)

func TestRunWithMockAPI(t *testing.T) {
	mockAPI := &mocks.MockCurrencyAPI{
		Codes:          [][]string{{"USD", "United States Dollar"}, {"EUR", "Euro"}},
		ConversionRate: 1.0, // Example conversion rate
	}

	args := &models.Args{
		Amount: 100,
		From:   "USD",
		To:     "EUR",
	}

	tmpFile, err := os.CreateTemp("", "currency_cache_test_*.json")
	tmpFile.WriteString(`{}`)
	tmpFile.Seek(0, 0)

	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	err = Run(mockAPI, args, tmpFile.Name())
	if err != nil {
		t.Fatalf("run failed %v", err)
	}
}
