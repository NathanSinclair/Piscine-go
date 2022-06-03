package piscine

// Divides the int a and b
// The result of the division stores in the int pointed by a
// The remainder of the division is stored in the int pointed by b.
func UltimateDivMod(a *int, b *int) {
	division := *a / *b
	module := *a % *b
	*a = division
	*b = module
}
