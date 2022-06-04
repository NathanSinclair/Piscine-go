package piscine

// Returns true if string passed as parameter only contains uppercase characters
func IsUpper(s string) bool {
	string1 := []rune(s)
	upper := true
	for i := 0; i <= len(string1)-1; i++ {
		if string1[i] >= 65 && string1[i] <= 90 && upper {
			upper = true
		} else {
			upper = false
		}
	}
	return upper
}
