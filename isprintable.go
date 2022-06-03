package piscine

// Returns true if string passed as parameter only contains printable characters
func IsPrintable(s string) bool {
	printable := true
	for _, letter := range s {
		if letter >= 31 && printable {
			printable = true
		} else {
			printable = false
		}
	}
	return printable
}
