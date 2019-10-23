package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	a := os.Args
	l := -1
	for i := range a {
		l++
		i = i
	}
	for i := l; i >= 1; i-- {
		s := []rune(a[i])
		for j := range s {
			z01.PrintRune(s[j])
		}
		z01.PrintRune(10)
	}
}
