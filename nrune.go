package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	s := 0
	for index, i := range runes {
			i = i
			s = index
	}
	if (n-1) > s {
		return 0
	}
	return runes[n-1]
}
