package piscine

// Takes two pointers to an int(*int) nd swaps their contents
func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
