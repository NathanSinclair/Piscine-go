package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x >= 1 {
		z01.PrintRune('/')
		if x >= 2 {
			for a := 2; a < x; a++ {
				z01.PrintRune('*')
			}
			if x >= 3 {
				z01.PrintRune('\\')
			}
		}
		z01.PrintRune('\n')
		for b := 2; b < y; b++ {
			if b >= 2 {
				z01.PrintRune('*')
				for a := 2; a < x; a++ {
					z01.PrintRune(' ')
				}
				if x >= 3 {
					z01.PrintRune('*')
				}
				z01.PrintRune('\n')
			}
		}
		if x >= 1 && y > 2 {
			z01.PrintRune('\\')
			if x >= 2 {
				for a := 2; a < x; a++ {
					z01.PrintRune('*')
				}
				if x >= 3 {
					z01.PrintRune('/')
				}
				z01.PrintRune('\n')
			}
		}
	}
	z01.PrintRune('\n')
}
