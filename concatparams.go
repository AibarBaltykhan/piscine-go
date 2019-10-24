package piscine

func ConcatParams(args []string) string {
	var a string
	l := -1
	for i := range args {
		i = i
		l++
	}
	for i := range args {
		if i <= l-2 {
			a = a + args[i] + "\n"
		} else {
			a = a + args[i]
		}
	}
	return a
}
