package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	l := -1
	for a, y := range runes {
		l++
		a = a
		y = y
	}
	for i := 0; i <= l; i++ {
		if runes[0] >= 'a' && runes[0] <= 'z' {
			runes[0] = runes[0] - 32
		}
		if i < l {
			if (runes[i] < 'a' || runes[i] > 'z') && (runes[i] < 'A' || runes[i] > 'Z') {
				if runes[i+1] >= 'a' && runes[i+1] <= 'z' {
					runes[i+1] = runes[i+1] - 32
				}
			} else {
				if (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= '0' && runes[i] <= '9') {
					runes[i+1] = runes[i+1] + 32
				}
			}
		}
	}
	s = string(runes)
	return s
}
