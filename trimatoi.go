package piscine

func TrimAtoi(s string) int {
	if s == "" {
		return 0
	}
	result := 0
	isneg := 0
	runes := []rune(s)

	for index, i := range runes {
		if runes[index] == '-' && runes[index+1] >= '0' && runes[index+1] < '9' {
			isneg++
			i = i
		}
	}

	for _, j := range runes {
		if j >= '0' && j < '9' {
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
			}
		}
	}
	if isneg >= 1 {
		result *= -1
	}
	return result
}
