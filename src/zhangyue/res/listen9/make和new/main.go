package main

import "fmt"

func testNew() {
	var a *int = new(int)
	*a = 100
	fmt.Printf("a=%v, a的内存地址%p, *a= v%  *a的内存地址%p\n", a, a, *a, &(*a))

	var b *[]int = new([]int)
	*b = make([]int, 5, 10)
	(*b)[0] = 10
	fmt.Printf("b=%v, b的内存地址%p, *b= v%  *b的内存地址%p\n", b, b, *b, &(*b))
}

func main() {
	testNew()
}
