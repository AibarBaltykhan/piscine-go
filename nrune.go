package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	k := 0
	for index, i := range runes {
		k++
		i = i
		index = index
	}
	if n > k {
		return -1
	}
	return runes[n-1]
}
