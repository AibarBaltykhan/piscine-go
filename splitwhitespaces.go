package piscine

func SplitWhiteSpaces(str string) []string {
	len := 0
	strlLen := 0
	start := 0
	j := 0
	b := false
	for _, v := range str {
		if (v == '\n' || v == '\t' || v == ' ') && wasLetter {
			len++
			b = false
		} else if v != '\n' && v != '\t' && v != ' ' {
			b = true
		}
		strlLen++
	}
	if b {
		len++
		b = false
	}
	arr := make([]string, len)
	for i, v := range str {
		if (v == '\n' || v == '\t' || v == ' ') && b {
			arr[j] = str[start:i]
			j++
			b = false
		} else if v != '\n' && v != '\t' && v != ' ' {
			if !b {
				start = i
			}
			b = true
		}
	}
	if b {
		arr[j] = str[start:strlLen]
	}
	return arr
}