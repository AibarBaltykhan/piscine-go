package piscine

func IsAlpha(str string) bool {
	runes := []rune(str)
	for i := range r {
		if (runes[i] >= 33 && runes[i] <= 47) || (runes[i] >= 58 && runes[i] <= 64) || (runes[i] >= 91 && runes[i] <= 96) || (runes[i] >= 123 && runes[i] <= 126) || runes[i] == ' ' {
			return false
		}
	}
	return true
}
