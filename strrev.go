package piscine

func StrRev(s string) string {
	runes := []rune(s)
	var a rune
	b := len(s)
	for i := 0; i < b/2; i++ {
		a = runes[i]
		runes[i] = runes[b-i-1]
		runes[b-i-1] = a
	}
	return string(runes)
}
