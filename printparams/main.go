package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Prints the arguments received in the command line
func main() {
	argument := os.Args
	for index, bc := range argument {
		if index != 0 {
			for _, i := range bc {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')
		}
	}
}
