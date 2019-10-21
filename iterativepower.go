package piscine

func IterativePower(nb int, power int) int {
	s := 1
	for i := 1; i <= power; i++ {
		s = s * nb
	}
	return s
}
