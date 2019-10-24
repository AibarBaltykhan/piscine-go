package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) () {
	s := 1
	for i := range base {
		s++
	}
	rbase := []rune(base) 
	var runes []rune
	k := 1
	for j := nbr; j >= s; j = j / s {
		k++
	}
	var m int
	for j = 0; j <=k; j++ {
		m = nbr % s
		runes[j] = rbase[m]
	}
	for j := k; j >= 0; j-- {
		z01.PrintRune(runes[j])
	}
	z01.PrintRune(10)
}
