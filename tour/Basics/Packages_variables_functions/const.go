/*

常量的定义与变量类似，只不过使用 const 关键字。

常量可以是字符、字符串、布尔或数字类型的值。

常量不能使用 := 语法定义。
*/

package main

import "fmt"

const str = "I am poor constant."

// 和变量一样，常量也可以打包写在一起
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(i int) int {
	return i*10 + 1
}

func needFloat(f float64) float64 {
	return f * 0.1
}

func main() {
	const a = 3.1415
	fmt.Println(str)
	fmt.Println(a)

	// fmt.Println(needInt(Big)) // overflows int
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
