package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"5", "6", "7", "8"}
	fmt.Println(strings.Join(slice, "*"))

	// slice a slice 起点是闭区间，终点是开区间
	fmt.Println(slice[1:1]) // slice from the 2nd to the 1st, so gets nothing.
	fmt.Println(slice[1:2])
	// 起点的缺省值为0，终点的缺省值为 len(slice)
	fmt.Println(slice[:])
	fmt.Println(slice[:len(slice)])

	fmt.Println(slice[:2]) // [0, 1]
	fmt.Println(slice[3:])
}
