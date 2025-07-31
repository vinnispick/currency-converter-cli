package utils

import (
	"flag"
	"os"
	"testing"
)

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestArgsParse_ResultIsCorrect(t *testing.T) {
	resetFlags()

	os.Args = []string{
		"cmd",
		"--amount=100.0",
		"--from=USD",
		"--to=EUR",
	}

	args, err := ArgParse()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if args.Amount != 100.0 {
		t.Errorf("Expected amount 100.0, got %f", args.Amount)
	}
	if args.From != "USD" {
		t.Errorf("Expected from currency USD, got %s", args.From)
	}
	if args.To != "EUR" {
		t.Errorf("Expected to currency EUR, got %s", args.To)
	}
}

func TestArgsParse_MissingAmount(t *testing.T) {
	resetFlags()

	os.Args = []string{
		"cmd",
		"--from=USD",
		"--to=EUR",
	}

	_, err := ArgParse()
	if err == nil {
		t.Fatal("Expected error for missing amount, got nil")
	}
	if err.Error() != "zero value of currency" {
		t.Errorf("Expected error 'zero value of currency', got %v", err)
	}
}

func TestArgsParse_MissingFrom(t *testing.T) {
	resetFlags()

	os.Args = []string{
		"cmd",
		"--amount=100.0",
		"--to=EUR",
	}

	_, err := ArgParse()
	if err == nil {
		t.Fatal("Expected error for missing from currency, got nil")
	}
	if err.Error() != "empty string of currency converts" {
		t.Errorf("Expected error 'empty string of currency converts', got %v", err)
	}
}

func TestArgsParse_MissingTo(t *testing.T) {
	resetFlags()

	os.Args = []string{
		"cmd",
		"--amount=100.0",
		"--from=USD",
	}

	_, err := ArgParse()
	if err == nil {
		t.Fatal("Expected error for missing to currency, got nil")
	}
	if err.Error() != "empty string of currency converts" {
		t.Errorf("Expected error 'empty string of currency converts', got %v", err)
	}
}

func TestArgsParse_ListFlag(t *testing.T) {
	resetFlags()

	os.Args = []string{
		"cmd",
		"--list=true",
	}

	args, err := ArgParse()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !args.List {
		t.Error("Expected list flag to be true")
	}
	if args.Amount != 0 || args.From != "" || args.To != "" {
		t.Error("Expected amount, from, and to to be zero or empty when list is true")
	}
}
