package piscine

// Applies the Map function of this type (int) bool on each element of that slice
// Returns a slice of all returned values
func Map(f func(int) bool, a []int) []bool {
	num := make([]bool, len(a))

	for i, j := range a {
		num[i] = f(j)
	}
	return num
}
