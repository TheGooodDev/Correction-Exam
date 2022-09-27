package exam1

import "github.com/01-edu/z01"

func Printreversealphabet() {
	for c := 'z'; c >= 'a'; c-- {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
