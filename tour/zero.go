package main

import "fmt"

// 未初始化的变量，都会被置为零值
func main() {
	var i int
	var f float32
	var s string
	var b bool
	// verbs list: https://golang.org/pkg/fmt/
	fmt.Printf("%v %v %v %v\n", i, f, s, b)  // 0 0  false
	fmt.Printf("%#v %f %q %v\n", i, f, s, b) // 0 0.000000 "" false
}
