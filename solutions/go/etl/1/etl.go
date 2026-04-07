package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}

    for key, values := range in {
        for _, letter := range  values {
            out[strings.ToLower(letter)] = key
        }
    }

    return out
}
