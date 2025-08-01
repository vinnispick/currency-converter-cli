package utils

import "currency-converter-cli/pkg/models"

func GetCacheKey(args *models.Args) string {
	return args.From + "_" + args.To
}
