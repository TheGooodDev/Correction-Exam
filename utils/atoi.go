package utils

import "errors"

func Atoi(s string) (int, error) {
	var res int
	trune := []rune(s)
	negatif := 1
	for i, c := range trune {
		if c == '-' {
			if i != 0 {
				return 0, nil
			} else {
				negatif = -1
			}

		} else if c == '+' {
			if i != 0 {
				return 0, nil
			}
		} else if !IsDigits(c) {
			return 0, errors.New("None digits value")
		} else {
			if i == len(trune)-1 {
				res += int(c - 48)
			} else {
				res += int(c - 48)
				res *= 10
			}

		}

	}
	return res * negatif, nil
}
