package piscine

// Returns true if the slice of int is sorted
func IsSorted(f func(a, b int) int, a []int) bool {
	isAscending := false
	isDescending := false
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			isDescending = true
		} else if f(a[i-1], a[i]) < 0 {
			isAscending = true
		}
	}
	if isAscending && isDescending {
		return false
	} else {
		return true
	}
}
