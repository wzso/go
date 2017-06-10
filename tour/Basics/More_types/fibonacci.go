package main

import "fmt"

/////// recursive approach   \\\\\\\\\
/////////////////\\\\\\\\\\\\\\\\\\\\\
func fibonacci(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func printFibonacci(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%v ", fibonacci(i))
	}
}

func main() {
	printFibonacci(10)
}
