package piscine

func RecursiveFactorial(nb int) int {
	s := 1
	if nb == 0 {
		return 1
	}
	if nb >= 1 && nb < 20 {
			s = s * nb
			return RecursiveFactorial(nb - 1)
		}
		return s
	} else {
		return 0
	}

}
