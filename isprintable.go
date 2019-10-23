package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if runes[i] < 33 || runes[i] > 126 {
			return false
		}
	}
	return true
}
