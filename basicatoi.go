package piscine

func BasicAtoi(s string) int {
	b := 0
	for _, n := range s {
		a := 0	
		for  i := '1'; i <= n; i++ { 
			a++
		}
		b = b*10 + a
	}
	return b
}