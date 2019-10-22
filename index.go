package piscine

func Index(s string, toFind string) int {
	w1 := []rune(s)
	w2 := []rune(toFind)
	if w2[0] == "" {
		return 0
	}
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
			for m := 1; m <= k-1; m++ {
				if w1[inde+m] == w2[m] {
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
