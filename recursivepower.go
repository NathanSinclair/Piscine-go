package piscine

// Returns the value of nb to the power of power without the use of for loops, negative powers return 0.
// Overflows are not dealt with
func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	if power > 0 {
		return nb * RecursivePower(nb, power-1)
	}
	return 0
}
