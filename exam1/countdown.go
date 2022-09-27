package exam1

import "github.com/01-edu/z01"

func Countdown() {
	for c := '9'; c >= '0'; c-- {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
