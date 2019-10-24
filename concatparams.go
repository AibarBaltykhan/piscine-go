package piscine

func ConcatParams(args []string) string {
	var a string
	for i := range args {
		a = a + args[i] + "\n"
	}
	return a
}
