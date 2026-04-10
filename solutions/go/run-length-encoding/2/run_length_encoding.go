package runlengthencoding

import (
    "strconv"
    "unicode"
    "strings"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			if count > 1 {
				result.WriteString(strconv.Itoa(count))
			}
			result.WriteByte(input[i-1])
			count = 1
		}
	}

	// handle last group
	if count > 1 {
		result.WriteString(strconv.Itoa(count))
	}
	result.WriteByte(input[len(input)-1])

	return result.String()
}

func RunLengthDecode(input string) string {
	var result strings.Builder
	count := 0

	for _, ch := range input {
		if unicode.IsDigit(ch) {
			count = count*10 + int(ch-'0') // build number
		} else {
			if count == 0 {
				count = 1
			}

			for i := 0; i < count; i++ {
				result.WriteRune(ch)
			}

			count = 0
		}
	}

	return result.String()
}
