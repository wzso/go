package main

import (
	"fmt"
	"math/cmplx"
)

var (
	a bool       = false
	b string     = "b"
	c complex128 = cmplx.Sqrt(-5 + 12i) // complex number 复数，完全忘记这个概念了 😶
	d complex64  = -5 + 82i
)

func main() {
	const f = "%T(%v)\n" // no new line appended. do it myself.
	fmt.Printf(f, a, a)
	fmt.Printf(f, c, c)
	fmt.Printf(f, d, d)
}

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 代表一个Unicode码

float32 float64

complex64 complex128
*/
