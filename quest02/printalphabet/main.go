package main

import "github.com/01-edu/z01"

//Prints the latin alphabet in lower case on a single line
func main() {
	i := 'a'
	for i <= 'z' {

		z01.PrintRune(i)
		i++
	}
	z01.PrintRune('\n')
}
