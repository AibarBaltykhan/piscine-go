package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	l := -1
	for a, y := range runes {
		l++
		a = a
		y = y
	}
	for i := 0; i <= l; i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = runes[i] - 32
		}
	}
	s = string(runes)
	return s
}
