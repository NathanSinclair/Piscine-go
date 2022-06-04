package piscine

// Applies the ForEach function on each element of a slice
func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}

func PrintNbr(num int) {
}
