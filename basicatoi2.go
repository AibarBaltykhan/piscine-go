package piscine

func BasicAtoi2(s string) int {
	b := 0
	for _, n := range s {
		a := 0
		if n < '0' || n > '9' {
			return 0
		}
		for i := '1'; i <= n; i++ {
			a++
		}
		b = b*10 + a
	}
	return b
}
