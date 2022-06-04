package piscine

// Counts the letters of a string and returns the count
func AlphaCount(s string) int {
	counter := 0
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' || letter >= 'a' && letter <= 'z' {
			counter++
		}
	}
	return counter
}
