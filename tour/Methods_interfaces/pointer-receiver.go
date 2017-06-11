package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// receiver can be mutated within methods with pointer receiver.
func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scalePointer(v *vertex, scale float64) {
	v.X = v.X * scale
	v.Y = v.Y * scale
}

func scaleValue(v vertex, scale float64) {
	v.X = v.X * scale
	v.Y = v.Y * scale
}

func main() {
	v := vertex{3, 4}
	v.Scale(10) // no error, go automatically interprets this statement as (&v).Scale(10)
	v.Abs()     // no error
	fmt.Println(v.Abs())

	// {30, 40}
	scaleValue(v, 10.0)
	fmt.Println(v) // still {30, 40}

	// scalePointer(v, 10.0) // compile error

	scalePointer(&v, 10.0)
	fmt.Println(v) // {300, 400}
}
