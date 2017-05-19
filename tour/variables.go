package main

import (
	"fmt"
)

var packageLevelVariable = 5
var a, b, c bool // initial value is false

func printASeriesOfVariables() {
	var d, e int
	fmt.Println(d, e)       // 0 0
	var f, g = "hello", 123 // 有初始值时，类型是可以推断的，因此可省略
	fmt.Println(f, g)
}

// :=
// 申明并初始化变量
// 只能在函数里面，函数外的每个语句都必须以关键字开始（var, func 等）
func syntaxSugar() {
	a := "bluez framework"
	fmt.Println(a)
}

// 与导入语句类似，变量也可以“打包”在一块儿
var (
	floatVar   float64
	str        string = "a block of vars [grin]"
	booleanVar        = true
)

func main() {
	var functionLevelVariable int
	var localString string
	fmt.Println(a)                     // false
	fmt.Println(functionLevelVariable) // 0
	fmt.Println(localString)           // 空串

	printASeriesOfVariables()
	syntaxSugar()

	fmt.Println(str, "\n", booleanVar)
}
