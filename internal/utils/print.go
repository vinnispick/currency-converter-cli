package utils

import (
	"fmt"
	"strings"
)

func PrintSupportedCodes(codes [][]string) {
	if len(codes) == 0 {
		fmt.Println("No supported codes found.")
		return
	}

	println("Supported Codes:")
	for _, code := range codes {
		fmt.Println(strings.Join(code, " | "))
	}
}

func PrintConversionResult(amount, result float64, from, to string) {
	fmt.Printf("Converted %.2f %s to %.2f %s\n", amount, from, result, to)
}
