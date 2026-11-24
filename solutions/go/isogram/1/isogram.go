package isogram

import "strings"

func IsIsogram(word string) bool {
    seen := make(map[rune]bool)

    word = strings.ToLower(word)
    word = strings.ReplaceAll(word, " ", "")
    word = strings.ReplaceAll(word, "-", "")

    for _, char := range word {
        if seen[char] {
            return false
        }

        seen[char] = true
    }

    return true
}