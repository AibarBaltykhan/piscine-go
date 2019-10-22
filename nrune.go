package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	k := 0
	for index, i := range runes {
		k++
		i = i
		index = index
	}
	if (n - 1) > k {
		return 0
	}
	return runes[n-1]
}
