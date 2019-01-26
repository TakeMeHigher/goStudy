package main

import "fmt"

func test() {
	var a []int = []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
	fmt.Printf("c=%v\n", c)
}

func main() {
	test()
}
