package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Prints the name of a program
func main() {
	args := os.Args[0]
	ar := []rune(args)
	for i := 2; i < len(ar); i++ {
		z01.PrintRune(ar[i])
	}
	z01.PrintRune('\n')
}
