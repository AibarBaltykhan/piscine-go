package piscine

func AlphaCount(str string) int {
	count := 0
	for _, n := range str {
		if (n >= 'a' && n <= 'z') || (n >= 'A' && n <= 'Z') {
			count++
		}
	}
	return count
}
