package exam2

func Nrune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	trune := []rune(s)
	if len(trune) == 0 || len(trune) < n {
		return 0
	}

	return trune[n-1]

}
