package rotationalcipher

func RotationalCipher(input string, shift int) string {
	var result []rune

	for _, ch := range input {
		switch {
		case ch >= 'a' && ch <= 'z':
			// shift lowercase
			newChar := 'a' + (ch-'a'+rune(shift))%26
			result = append(result, newChar)

		case ch >= 'A' && ch <= 'Z':
			// shift uppercase
			newChar := 'A' + (ch-'A'+rune(shift))%26
			result = append(result, newChar)

		default:
			// keep punctuation, spaces unchanged
			result = append(result, ch)
		}
	}

	return string(result)
}