package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 类型 [n]T 是一个有 n 个类型为 T 的值的数组
	// 数组的长度是其类型的一部分，因此数组不能改变大小
	var arr [10]string
	arr[0] = "hello"
	for i := 1; i < 10; i++ {
		arr[i] = strconv.Itoa(i)
	}
	fmt.Println(arr, "len is", len(arr))
	fmt.Println(len(arr))

	// []T 是 slice
	// slice
	slice := []int{2, 3, 4, 5, 6}
	fmt.Println(len(slice))
}
