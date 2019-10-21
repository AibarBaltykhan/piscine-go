package piscine

func FindNextPrime(nb int) int {
	s := nb
	if nb < 2 {
		return 2
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return s
}
