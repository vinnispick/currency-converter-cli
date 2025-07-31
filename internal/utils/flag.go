package utils

import (
	"currency-converter-cli/pkg/models"
	"errors"
	"flag"
)

func FlagParse() (*models.Args, error) {
	flag.Parse()
	var args models.Args

	flag.Float64Var(&args.Amount, "amount", 0, "Amount of currency")
	flag.StringVar(&args.From, "from", "", "From currency")
	flag.StringVar(&args.To, "to", "", "To currency")

	if args.Amount == 0 {
		return nil, errors.New("zero value of currency")
	}
	if args.From == "" || args.To == "" {
		return nil, errors.New("empty string of currency converts")
	}

	return &args, nil
}
