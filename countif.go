package piscine

// Returns the number  of elements of a string slice when the f function returns true
func CountIf(f func(string) bool, tab []string) int {
	num := 0
	for _, a := range tab {
		if f(a) {
			num++
		}
	}
	return num
}
