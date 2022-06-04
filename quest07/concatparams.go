package piscine

// Takes the arguments received in the parameters and returns them as a string
// The string si the result of all the arguments concatenated with a newline in between
func ConcatParams(args []string) string {
	var str []rune
	for i := 0; i < len(args); i++ {
		for x := 0; x < len(args[i]); x++ {
			string := []rune(args[i])
			str = append(str, rune(string[x]))
		}
		if i < len(args)-1 {
			str = append(str, '\n')
		}
	}
	return string(str)
}
