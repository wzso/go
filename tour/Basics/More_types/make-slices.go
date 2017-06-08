package main

import "fmt"

func printSlice(name string, slice []int) {
	fmt.Printf("%s (%p): len == %d, cap == %d. %v.\n", name, &slice, len(slice), cap(slice), slice)
}

func main() {
	s1 := make([]int, 5) // all zero
	printSlice("s1", s1)

	// len is 2, cap is 10
	s2 := make([]int, 3, 10)
	printSlice("s2", s2)

	s3 := s2[2:]
	printSlice("s3", s3)

	s4 := s3[:2]
	printSlice("s4", s4)

	s2[2] = 100
	printSlice("s3", s3)
	printSlice("s4", s4)

	// for more, visit: http://www.tuicool.com/articles/QrymYz
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	as := arr[:3]
	fmt.Printf("%p, %p", &arr, &as)
	as[0] = 9
	fmt.Println(arr, as)
}
