package main

import "fmt"

func functionGenerator() func(int) int {
	n := 100
	fmt.Printf("n = %v, addr:%p\n", n, &n)
	result := func(num int) int {
		n += num
		fmt.Printf("Inner----n = %v, addr:%p\n", n, &n)
		return n
	}
	return result
}

func main() {
	f := functionGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println(f(i))
	}
	sh := functionGenerator()
	sh(10)
}
