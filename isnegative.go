package piscine

import "github.com/01-edu/z01"

// Prints 'T'(true) on a single line if the int parameter is negative, otherwise it prints 'F'(false)
func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	}
}
