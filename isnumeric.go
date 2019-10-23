package piscine

func IsNumeric(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if runes[i] < '0' || runes[i] > '9' {
			return false
		}
	}
	return true
}
