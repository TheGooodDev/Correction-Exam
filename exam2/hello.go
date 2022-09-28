package exam2

import "github.com/01-edu/z01"

func Hello() {
	for _, c := range "Hello World!" {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
