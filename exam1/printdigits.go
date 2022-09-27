package exam1

import "github.com/01-edu/z01"

func Printdigits() {
	for c := '0'; c <= '9'; c++ {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
