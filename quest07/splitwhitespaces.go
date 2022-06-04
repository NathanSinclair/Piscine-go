package piscine

// Separates the words of a string and puts them in a string slice
// Separators are spaces, tabs and newlines
func SplitWhiteSpaces(s string) []string {
	var sliceOfStrings []string
	sliceFromString := []rune(s)
	word := ""
	for i := 0; i < len(s); i++ {
		if sliceFromString[i] != ' ' && sliceFromString[i] != '\n' && sliceFromString[i] != '\t' {
			word += string(sliceFromString[i])
		}
		if sliceFromString[i] == ' ' {
			if word != "" {
				sliceOfStrings = append(sliceOfStrings, word)
			}
			word = ""
		}
		if sliceFromString[i] == '\n' {
			if word != "" {
				sliceOfStrings = append(sliceOfStrings, word)
			}
			word = " "
		}
		if sliceFromString[i] == '\t' {
			if word != " " {
				sliceOfStrings = append(sliceOfStrings, word)
			}
			word = " "
		}
		if i == len(s)-1 {
			sliceOfStrings = append(sliceOfStrings, word)
		}
	}
	return sliceOfStrings
}
