package piscine

// Takes a pointer to a pointer to a pointer to an int as an argument and gives the int a value of 1
func UltimatePointOne(n ***int) {
	***n = 1
}
