package main

import "fmt"

// 函数作为参数

func add(a, b int) int {
	return a + b
}

func testfunc(a, b int, fun func(int, int) int) int {
	sum := fun(a, b)
	return sum
}

func main() {
	sum := testfunc(2, 3, add)
	fmt.Println(sum)
}
