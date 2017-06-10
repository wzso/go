package main

import "fmt"

func main() {
	m := map[int]int{
		1: 100,
		2: 110,
		3: 150,
	}
	// insert key-value pair
	m[4] = 220
	fmt.Println(m)

	// delete key-value pair
	delete(m, 3)
	fmt.Println(m)

	// detect key-value
	v, exsists := m[10]

	const f = "%T(%v)\n"
	fmt.Printf(f, exsists, exsists)
	fmt.Printf(f, v, v) // when exsists is false, v == 0

	m2 := map[int]string{
		2: "2",
		4: "",
	}
	v2, exsists2 := m2[5]
	fmt.Printf(f, exsists2, exsists2)
	fmt.Printf(f, v2, v2) // when exsists is false, v == ""

	_, exsists3 := m2[4]
	fmt.Println(exsists3)
}
