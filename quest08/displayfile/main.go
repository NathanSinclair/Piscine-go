package main

import (
	"fmt"
	"os"
)

// Displays on the standard output, the content of a file given as an argument
func main() {
	a := len(os.Args)
	if a == 1 {
		fmt.Println("File name missing")
	} else if a == 2 {
		fmt.Println("Almost there!!")
	} else {
		fmt.Println("Too many arguments")
	}
}
