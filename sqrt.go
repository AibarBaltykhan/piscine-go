package piscine

func Sqrt(nb int) int {
	for i := 2; i <= nb/2; i ++ {
		if nb/i == i && nb % i == 0 {
			return i
		}
	}
	return 0
}
