package utils

func IsDigits(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
