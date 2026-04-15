package largestseriesproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if span > len(digits) {
		return 0, errors.New("span must not exceed string length")
	}
	if span == 0 {
		return 1, nil
	}

	max := int64(0)

	for i := 0; i <= len(digits)-span; i++ {
		product := int64(1)

		for j := 0; j < span; j++ {
			ch := digits[i+j]

			if ch < '0' || ch > '9' {
				return 0, errors.New("digits input must only contain digits")
			}

			product *= int64(ch - '0')
		}

		if product > max {
			max = product
		}
	}

	return max, nil
}