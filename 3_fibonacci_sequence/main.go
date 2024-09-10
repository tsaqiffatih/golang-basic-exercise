package main

import "fmt"

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a, b := 1, 1

	for i := 2; i < n; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
