package allyourbase

import (
	"errors"
	"fmt"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
	}

	decimal := convertToDecimal(inputDigits, inputBase)
    fmt.Println("convertToDecimal", decimal)
	return convertFromDecimal(decimal, outputBase), nil
}

// convertToDecimal interprets digits as a number in the given base
// and returns its decimal (base-10) integer value.
func convertToDecimal(digits []int, base int) int {
	result := 0
	for _, d := range digits {
		result = result*base + d
	}
	return result
}

// convertFromDecimal converts a decimal integer into a slice of digits
// in the given base, most-significant digit first.
func convertFromDecimal(n, base int) []int {
	if n == 0 {
		return []int{0}
	}
	var digits []int
	for n > 0 {
		digits = append(digits, n%base)
		n /= base
	}
	// reverse so most-significant digit comes first
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}