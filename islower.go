package piscine

func IsLower(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if (runes[i] >= 33 && runes[i] <= 47) || (runes[i] >= 58 && runes[i] <= 64) || (runes[i] >= 91 && runes[i] <= 96) || (runes[i] >= 123 && runes[i] <= 126) || runes[i] == ' ' || (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= '0' && runes[i] <= '9') {
			return false
		}
	}
	return true
}
