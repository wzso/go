package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
		fmt.Println(i, sum)
	}

	fmt.Println("-------")
	// i 的操作是可省的，相当于 while 语句了
	// for is go's while
	for sum > 100 {
		sum /= 2
		fmt.Println(sum)
	}

	// infinit loop
	for {
		fmt.Println("I can print forever.")
	}
}
