package piscine

func Fibonacci(index int) int {
	if index == 1 {
		return 0
	}
	if index == 2 {
		return 1
	}
	if index >= 3 {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
	return -1
}
