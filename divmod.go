package piscine

// Divides the int a and b
// The result of the division stores in the int pointed by the div
// The remainder of the division is stored in the int pointed by the mod.
func DivMod(a int, b int, div *int, mod *int) {
	division := a / b
	module := a % b
	*div = division
	*mod = module
}
