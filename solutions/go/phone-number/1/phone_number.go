package phonenumber

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`\D`)

func Number(phoneNumber string) (string, error) {
	digits := re.ReplaceAllString(phoneNumber, "")

	// Handle country code
	if len(digits) == 11 {
		if digits[0] != '1' {
			return "", fmt.Errorf("invalid country code")
		}
		digits = digits[1:]
	}

	// Must be exactly 10 digits
	if len(digits) != 10 {
		return "", fmt.Errorf("invalid length")
	}

	// NANP rules
	if digits[0] < '2' || digits[0] > '9' {
		return "", fmt.Errorf("invalid area code")
	}
	if digits[3] < '2' || digits[3] > '9' {
		return "", fmt.Errorf("invalid exchange code")
	}

	return digits, nil
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s",
		num[0:3],
		num[3:6],
		num[6:],
	), nil
}