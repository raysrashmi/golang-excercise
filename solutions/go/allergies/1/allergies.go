package allergies

var allergenOrder = []struct {
	name  string
	value uint
}{
	{"eggs", 1},
	{"peanuts", 2},
	{"shellfish", 4},
	{"strawberries", 8},
	{"tomatoes", 16},
	{"chocolate", 32},
	{"pollen", 64},
	{"cats", 128},
}

// Returns all allergies for given score
func Allergies(score uint) []string {
	var result []string

	for _, a := range allergenOrder {
		if score&a.value != 0 {
			result = append(result, a.name)
		}
	}

	return result
}

// Checks if allergic to a specific allergen
func AllergicTo(score uint, allergen string) bool {
	for _, a := range allergenOrder {
		if a.name == allergen {
			return score&a.value != 0
		}
	}
	return false
}