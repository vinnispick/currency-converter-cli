package cli

import (
	"currency-converter-cli/internal/cache"
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
	cache := cache.NewFileCache(tmpFile.Name())
	err = Run(mockAPI, args, cache)
	if err != nil {
		t.Fatalf("run failed %v", err)
	}
}

func TestRunWithMockCache(t *testing.T) {
	mockAPI := &mocks.MockCurrencyAPI{
		Codes:          [][]string{{"USD", "United States Dollar"}, {"EUR", "Euro"}},
		ConversionRate: 1.0,
	}

	args := &models.Args{
		Amount: 100,
		From:   "USD",
		To:     "EUR",
	}

	cache := mocks.NewMockCache()
	err := Run(mockAPI, args, cache)
	if err != nil {
		t.Fatalf("run failed %v", err)
	}

	val, ok := cache.Data["USD_EUR"]
	if !ok {
		t.Fatalf("expected cache to contain USD_EUR")
	}
	if val != 1.0 {
		t.Errorf("expected cache value 1.0, got %f", val)
	}
}
