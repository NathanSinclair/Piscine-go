package piscine

// Returns true if string passed as parameter only contains alphanumerical characters or is empty
func IsAlpha(s string) bool {
	alpha := true
	for _, letter := range s {
		if letter >= 97 && letter <= 122 && alpha || letter >= 65 && letter <= 90 && alpha || letter >= 48 && letter <= 57 && alpha {
			alpha = true
		} else {
			alpha = false
		}
	}
	return alpha
}
