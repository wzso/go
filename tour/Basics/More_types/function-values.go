package main

import "fmt"

func main() {
	wordCount := func(s string) int {
		slice := []rune(s)
		fmt.Println(slice)
		return len(slice)
	}
	byteCount := func(s string) int {
		slice := []byte(s)
		fmt.Println(slice)
		return len(slice)
	}

	str := "哈哈哈"
	fmt.Println(wordCount(str))
	fmt.Println(byteCount(str))
	fmt.Printf("%T(%v): %p", wordCount, wordCount, &wordCount)
}
