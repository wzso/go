package main

import (
	"fmt"
	"math"
)

/*
You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method.
You cannot declare a method with a receiver whose type is defined in another package
(which includes the built-in types such as int).
*/

type MyFloat float64

func (f MyFloat) abs() float64 {
	return math.Abs(float64(f))
}

func main() {
	f := MyFloat(-123.2)
	fmt.Println(f.abs())
}
