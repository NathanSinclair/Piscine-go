package piscine

// Returns the last rune of a string
func LastRune(s string) rune {
	sentence := []rune(s)
	return sentence[len(sentence)-1]
}
