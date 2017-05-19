/*
the magic iota
from: http://studygolang.com/articles/2192

iota只能在常量的表达式中使用。

fmt.Println(iota)
编译错误： undefined: iota

每次 const 出现时，都会让 iota 初始化为0.
*/
package main

import "fmt"

// enum-like
type fakeEnum int

const (
	a fakeEnum = iota
	b          // 1
	c          // 2
	_
	d // 4
)

// bit mask
const (
	aa fakeEnum = 1 << iota
	bb          // 1<<1
	cc          // 1<<2
	_
	dd // 1<<4
)

func main() {
	fmt.Println(d, dd)
}
