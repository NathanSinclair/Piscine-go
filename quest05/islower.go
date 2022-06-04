package piscine

// Returns true if string passed as parameter only contains lowercase characters
func IsLower(s string) bool {
	string1 := []rune(s)
	upper := true
	for i := 0; i <= len(string1)-1; i++ {
		if string1[i] >= 97 && string1[i] <= 122 && upper {
			upper = true
		} else {
			upper = false
		}
	}
	return upper
}
