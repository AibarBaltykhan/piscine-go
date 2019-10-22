package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	k := 0
	for index, i := range runes {
		k++
		i = i
		index = index
	}
	return runes[k-1]
}
