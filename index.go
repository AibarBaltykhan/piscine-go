package piscine

func Index(s string, toFind string) int {
	w1 := []rune(s)
	w2 := []rune(toFind)
	k := 0
	l := -1
	t := 1
	for index, i := range w2 {
		k++
		i = i
		index = index
	}
	if k == 0 {
		return -1
	}
	for inde, j := range w1 {
		l++
		j = j
		inde = inde
		if w1[inde] == w2[0] {
			for m := l + 1; m <= l+k-1; m++ {
				if w1[m] == w2[m-l] {
					t++
				}
			}
			if t == k {
				return l
			}
		}
	}
	return -1
}
