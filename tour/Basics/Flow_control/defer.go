package main

import "fmt"

func foo() string {
	fmt.Println("AAA.")
	return "foo"
}

func foo2() string {
	defer fmt.Println("BBB.")
	return "foo2"
}

// defer stack, LIFO
// https://blog.go-zh.org/defer-panic-and-recover
func multiDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Println("CCC.", i)
	}
}

func main() {
	// 提供参数的 foo 会即刻，而defer后面的语句延迟到函数即将返回时才调用。
	defer fmt.Println("DDD.", foo(), foo2())

	multiDefer()

	fmt.Println("EEE.")
}
