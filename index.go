package piscine

func Index(s string, toFind string) int {
	w1 := []rune(s)
	w2 := []rune(toFind)
	k := -1
	l := -1
	t := 0
	for index, i := range w2 {
		k++
		i = i
		index = index
	}
	for inde, j := range w1 {
		l++
		j = j
		inde = inde
		if j == w2[0] {
			for m := l; m < l+k; l++ {
				if w1[m] == w2[k] {
					t++
				}

			}
			if t == k+1 {
				return l
			}
		}
	}
	return -1
}
