package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Returns even or odd string if the remainder of an int is equal to zero
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	if isEven(len(os.Args)) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
