package main

import "fmt"

func main() {
	// 跟 C 语言指针不同，Go 指针没有指针运算
	i, j := 22, 33
	p := &i
	fmt.Println(*p)
	*p = j
	fmt.Println(i)
}
