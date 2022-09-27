package exam1

import "github.com/01-edu/z01"

func PrintrevComb() {
	for nb1 := '9'; nb1 >= '0'; nb1-- {
		for nb2 := '9'; nb2 >= '0'; nb2-- {
			for nb3 := '9'; nb3 >= '0'; nb3-- {
				if nb1 > nb2 && nb2 > nb3 {
					z01.PrintRune(nb1)
					z01.PrintRune(nb2)
					z01.PrintRune(nb3)
					if nb1 != '2' || nb2 != '1' || nb3 != '0' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
