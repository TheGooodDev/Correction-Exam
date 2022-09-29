package exam3

func Max(a []int) int {
	res := 0
	for _, nb := range a {
		if nb > res {
			res = nb
		}
	}
	return res
}
