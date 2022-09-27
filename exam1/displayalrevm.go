package exam1

import "github.com/01-edu/z01"

func Displayalrevm() {
	for i, j := 1, 'z'; j >= 'a'; i, j = i+1, j-1 {
		if i%2 == 0 {
			z01.PrintRune(j - 32)
		} else {
			z01.PrintRune(j)
		}
	}
	z01.PrintRune('\n')
}
