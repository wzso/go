package main

import (
	"fmt"
	"math"
	"math/rand"
)

// if with short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		// v is a local variable within if-else clause
		return v
	} else {
		fmt.Printf("%v is bigger than %v.\n", v, lim)
	}
	return lim
}

func main() {
	i := rand.Int()
	fmt.Println(i)

	// plain if clause
	if i < 100 {
		fmt.Println("Small i")
	}

	fmt.Println(
		pow(2, 3, 16),
		pow(3, 5, 20), // comma is necessary
	)

	fmt.Println(pow(2, 4, 15), pow(3, 5, 36))
}
