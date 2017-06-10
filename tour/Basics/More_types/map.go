package main

import "fmt"

type Location struct {
	Latitude, Longitude float32
}

func main() {
	m := make(map[string]Location)

	// set value for key
	m["Shenzhen, CN"] = Location{12.0, 73.6}
	fmt.Println(m)

	//
	m2 := map[int]Location{
		1: Location{1, 1},
		2: Location{2, 2}, // not duplicate key allowed
		// "1": Location{3, 3}, // key type must be as declared
	}
	m2[3] = Location{3, 3}
	fmt.Println(m2)

	// type `Location` can be eliminated
	m3 := map[float32]Location{
		2.4:  {24, 24},
		2.41: {241, 241},
	}
	fmt.Println(m3)
}
