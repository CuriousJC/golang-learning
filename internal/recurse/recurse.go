package recurse

import (
	"fmt"
)

func Factorial(p int) int {
	fmt.Println("Running factorial for", p)

	if p == 0 || p == 1 {
		return 1
	}
	return p * Factorial(p-1)

}

func Fibonacci(p int) int {
	fmt.Println("Calculating fibonacci for", p)
	if p == 0 {
		return 0
	}

	if p == 1 {
		return 1
	}

	return Fibonacci(p-1) + Fibonacci(p-2)
}
