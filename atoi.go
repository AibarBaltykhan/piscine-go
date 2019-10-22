package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	result := 0
	isneg := 0
	runes := []rune(s)
	if runes[0] == '-' {
		isneg = 1
		runes[0] = '0'
	} else if runes[0] == '+' {
		runes[0] = '0'
	}
	for _, j := range runes {
		result *= 10
		if j >= '0' && j <= '9' {
			result = result + 'j'-'0'
			return 0
		}
	}
	if isneg == 1 {
		result *= -1
	}
	return result
}