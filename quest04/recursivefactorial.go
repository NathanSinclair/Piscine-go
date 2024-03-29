package piscine

// Returns the factorial of the int passed as a parameter with the use of for loops, errors return 0
func RecursiveFactorial(nb int) int {
	if nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb == 1 {
		return 1
	}
	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}
