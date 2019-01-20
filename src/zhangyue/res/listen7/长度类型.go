package main

import "fmt"

func test1() {
	// var a [3]int;
	// b := [4]int{1,2,3,4}
	// a = b   报错 不是 同类型的数组 不能直接 赋值

	var c [4]int
	d := [4]int{1, 2, 3, 4}
	c = d
	fmt.Printf("c=%v\n", c)
	fmt.Printf("d=%v\n", d)
}

func main() {
	test1()
}
