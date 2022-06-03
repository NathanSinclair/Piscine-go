package piscine

import "github.com/01-edu/z01"

// Prints characters of a string one by one.
func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
}
