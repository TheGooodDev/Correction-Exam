package utils

func IsAlphabetics(c rune) bool {
	if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
		return true
	}
	return false
}
