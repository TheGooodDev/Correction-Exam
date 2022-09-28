package utils

func IsUpper(c rune) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	} else {
		return false
	}
}
