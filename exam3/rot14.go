package exam3

import (
	"correction/utils"
)

func Rot14(s string) string {
	trune := []rune(s)
	for i, c := range trune {
		if utils.IsAlphabetics(c) {
			if utils.IsUpper(c) {
				temp := c - 'A'
				if temp+14 > 26 {
					trune[i] = 'A' + (temp + 14 - 26)
				} else {
					trune[i] += 14
				}
			} else {
				temp := c - 'a'
				if temp+14 > 26 {
					trune[i] = 'a' + (temp + 14 - 26)
				} else {
					trune[i] += 14
				}
			}
		}
	}

	return string(trune)
}
