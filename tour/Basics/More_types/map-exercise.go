package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWords(s string) map[string]int {
	// `rune`, `byte`, `string`
	// 访问 unicode 编码时，需要转换为 `rune` 切片
	f := func(char rune) bool {
		return !unicode.IsLetter(char)
	}
	words := strings.FieldsFunc(s, f)
	fmt.Println(words)
	result := make(map[string]int)
	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	fmt.Println(countWords("we are who we are, always. If not, who else can we be?"))
}
