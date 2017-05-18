package main

import (
	"fmt"
)

// 参数类型放在参数名后
func add(a int, b int) int {
	return a + b
}

// 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
func sum(a, b, c int) int {
	return a + b + c
}

// 多返回值
func swap(a, b string) (string, string) {
	return b, a
}

// named-return value
func namedResult() (x int, y string) {
	x = 3
	y = "haha, you don't have to return me directly."
	return
}

func main() {
	fmt.Println("oops, look this Go function! weirdo.")
	fmt.Println(add(2, 9))
	fmt.Println(sum(1, 2, 3))
	// :
	a, b := swap("hello", "world")
	a, b = b, a
	fmt.Println(a, b)

	// named return value
	fmt.Println(namedResult())
}
