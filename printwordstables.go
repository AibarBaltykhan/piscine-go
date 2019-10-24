package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for i := range table {
		PrintStr(table[i])
		z01.PrintRune('\n')
	}
}
