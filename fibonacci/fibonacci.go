package fibonacci

// Fibonacci computes the nth Fibonacci number.
//
// No care is being taken about integer overflows.
func Fibonacci(n int) int {
	a, b := 1, 1

	for ; n > 0; n-- {
		a, b = b, a+b
	}

	return a
}
