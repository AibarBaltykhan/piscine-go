package piscine

func ConcatParams(args []string) string {
	var a string
	for i := range args {
		a = make(a, args[i]) + "\n"
	}
	return a
}
