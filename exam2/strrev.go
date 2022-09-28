package exam2

func StrRev(s string) string {
	trune := []rune(s)

	for i, j := 0, len(trune)-1; i < j; i, j = i+1, j-1 {
		trune[i], trune[j] = trune[j], trune[i]
	}
	return string(trune)
}
