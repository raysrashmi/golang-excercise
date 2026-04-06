package pangram

import "strings"

func IsPangram(input string) bool {
	m := make(map[rune]bool)
    input = strings.ToLower(input)
    input = strings.ReplaceAll(input, " ", "")
    for _, char := range input {
        if (char >= 'a' && char <= 'z') {
        	m[char] = true
        }
    }
	
    return len(m) == 26
    
}
