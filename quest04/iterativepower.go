package piscine

// Returns the value of nb to the power of power, negative powers return 0.
// Overflows are not dealt with
func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power >= 65 {
		return 0
	}
	result := 1
	for i := 1; i <= power; i++ {
		result = result * nb
	}
	return result
}
