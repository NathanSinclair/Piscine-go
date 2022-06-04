package piscine

// Returns true for a string slice passed through the f function when at least one element returns true
func Any(f func(string) bool, a []string) bool {
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			return true
		}
	}
	return false
}
