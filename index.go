package piscine

// Custom Index function
func Index(s string, toFind string) int {
	if len(s) <= 0 || len(toFind) <= 0 {
		return 0
	}
	runes := []rune(toFind)
	for index, i := range s {
		if len(runes) > 1 {
			if i == runes[0] {
				for index, i = range s {
					if i == runes[1] {
						return index - 1
					}
				}
			}
		} else {
			if i == runes[0] {
				return index
			}
		}
	}
	return -1
}
