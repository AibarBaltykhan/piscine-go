package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	l := 0
	for a, y := range runes {
		l++
		a = a
		y = y
	}
	for i := 0; i <= l; i++ {
		if runes[i] == ' ' {
			if runes[a+1] >= 'a' && runes[a+1] <= 'z' {
				runes[i+1] = runes[i+1] - 32
			}
		}
	}
	s := string(runes)
	return s
}
