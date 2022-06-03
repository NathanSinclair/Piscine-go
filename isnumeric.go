package piscine

// Returns true if string passed as parameter only contains numerical characters
func IsNumeric(s string) bool {
	numeric := true
	for _, letter := range s {
		if letter >= 48 && letter <= 57 && numeric {
			numeric = true
		} else {
			numeric = false
		}
	}
	return numeric
}
