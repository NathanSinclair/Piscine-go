package piscine

// Returns the nth rune of a string. If not possible, retuns 0
func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	sentence := []rune(s)
	return sentence[n-1]
}
