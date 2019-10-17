package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	runes := []rune(str)
	for i := range runes {
		z01.PrintRune(runes[i])
	}
}
