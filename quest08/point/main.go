package main

import "github.com/01-edu/z01"

// Converts the pointed string to x and y and prints that string
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func setPoint(x, y *string) {
	*x = "42"
	*y = "21"
}

func main() {
	var x, y string

	setPoint(&x, &y)

	printStr("x = " + x + ", y = " + y)
}
