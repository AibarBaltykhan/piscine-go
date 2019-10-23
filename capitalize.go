package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	l := 0
	for a, y := range runes {
		l++
		a = a
		y = y
	}
	for i := l-1; i >= 0; i-- {
		if runes[0] >= 'a' && runes[0] <= 'z' {
			runes[0] = runes[0] - 32
		}
		if (runes[i-1] >= 33 && runes[i-1] <= 47) || (runes[i-1] >= 58 && runes[i-1] <= 64) || (runes[i-1] >= 91 && runes[i-1] <= 96) || (runes[i-1] >= 123 && runes[i-1] <= 126) || runes[i-1] == ' ' {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = runes[i] - 32
			}
		}
		if (runes[i-1] >= 48 && runes[i-1] <= 57) || (runes[i-1] >= 65 && runes[i-1] <= 90) || (runes[i-1] >= 97 && runes[i-1] <= 122) {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = runes[i] + 32
			}
		}
	}
	s = string(runes)
	return s
}
