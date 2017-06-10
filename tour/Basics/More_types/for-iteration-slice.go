package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 7, 9, 11, 13, 15}
	for i, v := range slice {
		fmt.Println("i =", i, ",", "v =", v)
	}

	// _
	for _, v := range slice {
		fmt.Println(v)
	}

}
