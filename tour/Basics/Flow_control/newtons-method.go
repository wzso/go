package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	abs := x
	ret := 1.0
	for i := 1.0; i < x; i++ {
		v := math.Abs((i*i - x) / 2 * i)
		fmt.Println(i, v, abs)
		if v == 0 {
			ret = i
			return ret
		}
		if v < abs {
			abs = v
			ret = i
		}
	}
	fmt.Println("-------")
	for i := ret; i < ret+1; i += 0.1 {
		v := math.Abs((i*i - x) / 2 * i)
		fmt.Println(i, v, abs)
		if v < abs {
			abs = v
			ret = i
		}
	}
	return ret
}

func main() {
	fmt.Println(sqrt(1))
	fmt.Println(math.Sqrt(17))
}
