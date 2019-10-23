package piscine

func Compare(a, b string) int {
	l1 := -1
	l2 := -1
	n := 0
	for a, y := range a {
		l1++
		a = a
		y = y
	}
	for b, z := range b {
		l2++
		b = b
		z = z
	}
		if l1 > l2 {
		n = l2
	} else {
		n = l1
 	}
  	r1 := []rune(a)
	r2 := []rune(b)
  	for i := 0; i <= l1; i++ {
    	if r1[i] > r2[i] {
      		return 1
    	}
    	if r1[i] < r2[i] {
      		return -1
    	}
  	}
  	if l1 == l2 {
    	return 0
  	}
  	if l1 > l2 {
		return 1
		}
  	if l1 < l2 {
		return -1
	}
}
