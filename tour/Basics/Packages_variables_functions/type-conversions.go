package main

import (
	"fmt"
)

func main() {
	i := 14
	f := float32(i) // 与 C 语言不通，不同类型的数值的转换，必须是显式的
	// c := complex128(i) // can't convert int to complex128
	s := string(i)
	const format = "%T(%v)\n"
	fmt.Printf(format, f, f) // float32(14)
	fmt.Printf(format, s, s) // string()

	// cannot use "bluez" (type string) as type float32 in assignment
	// f = "bluez"
	// fmt.Printf(format, f, f)

	s = "bluez"
	fmt.Printf(format, s, s)
}
