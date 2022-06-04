package piscine

// Counts the runes of a string and returns that count
func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}
