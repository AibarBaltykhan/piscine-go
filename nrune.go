package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	for index, i := range runes {
		i = i
	}
	if (n - 1) > len(s) {
		return 0
	}
	return runes[n-1]
}
