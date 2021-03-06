package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	expected := [][]int{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, v := range expected {
		if Fibonacci(v[0]) != v[1] {
			t.Errorf("Fibonacci(%d) is %d, should be %d\n", v[0], Fibonacci(v[0]), v[1])
		}
	}
}
