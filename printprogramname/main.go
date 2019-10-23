package piscine

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	a := os.Args
	name := []rune(a[0])
	for i := range name {
		z01.PrintRune(name[i])
	}
	z01.PrintRune(10)
}
