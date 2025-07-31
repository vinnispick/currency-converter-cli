package converter

import "testing"

func TestConverter_ResultIsCorrect(t *testing.T) {
	amount := 100.0
	rate := 1.2
	expected := 120.0

	result := Convert(amount, rate)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
