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
		if j == '0' {
			result += 0
		} else if j == '1' {
			result += 1
		} else if j == '2' {
			result += 2
		} else if j == '3' {
			result += 3
		} else if j == '4' {
			result += 4
		} else if j == '5' {
			result += 5
		} else if j == '6' {
			result += 6
		} else if j == '7' {
			result += 7
		} else if j == '8' {
			result += 8
		} else if j == '9' {
			result += 9
		} else {
			return 0
		}
	}
	if isneg == 1 {
		result *= -1
	}
	return result
}
