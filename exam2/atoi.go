package exam2

import "correction/utils"

func Atoi(s string) int {
	var res int
	trune := []rune(s)
	negatif := 1
	for i, c := range trune {
		if c == '-' {
			if i != 0 {
				return 0
			} else {
				negatif = -1
			}

		} else if c == '+' {
			if i != 0 {
				return 0
			}
		} else if !utils.IsDigits(c) {
			return 0
		} else {
			if i == len(trune)-1 {
				res += int(c - 48)
			} else {
				res += int(c - 48)
				res *= 10
			}

		}

	}
	return res * negatif
}
