package piscine

func Index(s string, toFind string) int {
	w1 := []rune(s)
	w2 := []rune(toFind)
	k := 0
	l := 0
	t := 1
	for index, i := range w2 {
		k++
		i = i
		index = index
	}
	for inde, j := range w1 {
		l++
		j = j
		inde = inde
		if w1[inde] == w2[0] {
			for m := l; m <= l+k-2; m++ {
				if w1[m] == w2[m-l+1] {
					t++
				}
			}
			if t == k {
				return l - 1
			}
		}
	}
	return -1
}
