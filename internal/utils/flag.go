package utils

import (
	"currency-converter-cli/pkg/models"
	"errors"
	"flag"
)

func ArgParse() (*models.Args, error) {
	var args models.Args

	flag.Float64Var(&args.Amount, "amount", 0, "Amount of currency")
	flag.StringVar(&args.From, "from", "", "From currency")
	flag.StringVar(&args.To, "to", "", "To currency")
	flag.BoolVar(&args.List, "list", false, "List all available currencies")
	flag.BoolVar(&args.Refresh, "refresh", false, "Refresh cache")

	flag.Parse()

	if args.List {
		return &args, nil
	}

	if args.Amount == 0 {
		return nil, errors.New("zero value of currency")
	}
	if args.From == "" || args.To == "" {
		return nil, errors.New("empty string of currency converts")
	}

	return &args, nil
}
