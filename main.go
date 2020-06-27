package cmd

import (
	"gh-actions-experiment/fibonacci"
	"fmt"
)

func main() {
	fmt.Printf("Fibonacci(10) = %d\n", fibonacci.Fibonacci(10))
}
