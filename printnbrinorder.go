package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var a []int
	if n < 0 {
		return
	}
	if n == 0 {
		a = append(a, 0)
	}
	for n != 0 {
		a = append(a, n%10)
		n = n / 10
	}
	SortIntegerTable(a)
	for z := range a {
		z01.PrintRune(rune('0' + a[z]))
	}
}
