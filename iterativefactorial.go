package piscine

func IterativeFactorial(nb int) int {

	s := 1
	if nb == 0 {
		return 1
	}
	if nb >= 1 && nb < 20 {
		for i := 1; i <= nb; i++ {
			s = s * i
		}
		return s
	} else {
		return 0
	}
}
