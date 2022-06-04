package main

import "github.com/01-edu/z01"

//Prints the decimal digits in ascending order on a single line
func main() {
	i := '0'
	for i <= '9' {

		z01.PrintRune(i)
		i++
	}
	z01.PrintRune('\n')
}
