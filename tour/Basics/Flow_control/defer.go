package main

import "fmt"

func foo() string {
	fmt.Printf("\\\\\\\\\\\n")
	return "bluez"
}

// defer stack, LIFO
// https://blog.go-zh.org/defer-panic-and-recover
func multiDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Println("for-loop", i)
	}
}

func main() {
	// 提供参数的 foo 会即刻，而defer后面的语句延迟到函数即将返回时才调用。
	defer fmt.Println("I am in a defer scope.", foo())

	multiDefer()

	fmt.Println("Function is going to end.")
}
