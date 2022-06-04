package piscine

// Returns the factorial of the int passed as a parameter, errors return 0
func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb >= 25 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	return result
}
