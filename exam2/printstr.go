package exam2

import "github.com/01-edu/z01"

func PrintStr(s string) {
	trune := []rune(s)
	for _, c := range trune {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
