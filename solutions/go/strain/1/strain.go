package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

//keep function which returns elemnets where predicate is true

func Keep[T any](s []T, predicate func(T) bool) []T {
	result := []T{}

	for _, e := range s {
		if predicate(e) {
			result = append(result, e)
		}
	}

	return result
}

//discard function which returns elemnets where predicate is false
func Discard[T any](s []T, predicate func(T) bool) []T {
    result := []T{}

    for _, e := range s {
        if !predicate(e) {
            result = append(result, e)
        }
    }

    return result
}