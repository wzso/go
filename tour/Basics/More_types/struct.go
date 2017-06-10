package main

import "fmt"

type name struct {
	firstName  string
	middleName string
	familyName string
}

func main() {
	// 省略所有初始值
	n1 := name{}
	fmt.Println(n1)
	// 省略部分参数，需要给传参加上标签，顺序无所谓
	n2 := name{middleName: "Jack", firstName: "Yo"}
	fmt.Println(n2)

	n := name{"Tom", "", "Smith"}
	n.middleName = "Padding"
	fmt.Println(n)

	// pointer to struct
	p := &n
	p.firstName = "Bob"
	fmt.Println(n)
	fmt.Printf("%T\n", p) // *main.name
}
