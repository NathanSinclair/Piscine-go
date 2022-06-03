package main

import "github.com/01-edu/z01"

//Prints the latin alphabet in lowercase in reverse order on a single line
func main() {
	i := 'z'
	for i >= 'a' {

		z01.PrintRune(i)
		i--
	}
	z01.PrintRune('\n')
}
