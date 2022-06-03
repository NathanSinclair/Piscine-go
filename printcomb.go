package piscine

import "github.com/01-edu/z01"

// Prints in ascending order and on a single line:
// all unique combinations off three digits, the first is lower than the second and the second is lower than the third.
// Combinations are separated by a comma and a space.
func PrintComb() {
	addComma := false
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				if addComma {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				addComma = true
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
			}
		}
	}
	z01.PrintRune('\n')
}
